package behave

import (
	"fmt"
	"reflect"
	"time"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/log"
	"github.com/wrmsr/bane/pkg/util/slices"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type Telegram[T any] struct {
	sender   Telegraph[T]
	receiver Telegraph[T]

	message T

	needsReturnReceipt bool
}

type ReturnReceipt[T any] struct {
	telegram Telegram[T]
}

type Telegraph[T any] interface {
	Identity() uintptr
	HandleMessage(telegram Telegram[T]) bool
}

type TelegramProvider[T any] interface {
	Identity() uintptr
	ProvideMessage(receiver Telegraph[T], messageType reflect.Type) bt.Optional[T]
}

type Delayed[T any] struct {
	telegram Telegram[T]
	deadline time.Time
}

//

type Dispatcher[T any] struct {
	generatedIdentity

	listeners Registry[reflect.Type, Telegraph[T]]
	providers Registry[reflect.Type, TelegramProvider[T]]

	queue slices.PriorityQueue[Delayed[T]]

	log log.DefaultGlobalLogger
}

func NewDispatcher[T any]() *Dispatcher[T] {
	d := &Dispatcher[T]{}
	d.listeners = NewTypeRegistry[Telegraph[T]](d.onAddListener)
	d.providers = NewTypeRegistry[TelegramProvider[T]](nil)
	return d
}

func (d *Dispatcher[T]) Listeners() Registry[reflect.Type, Telegraph[T]]        { return d.listeners }
func (d *Dispatcher[T]) Providers() Registry[reflect.Type, TelegramProvider[T]] { return d.providers }

func (d *Dispatcher[T]) onAddListener(receiver Telegraph[T], messageType reflect.Type) {
	s := d.providers.Get(messageType)
	if s != nil {
		for _, provider := range s {
			provider.ProvideMessage(receiver, messageType).IfPresent(func(message T) {
				var sender Telegraph[T]
				if p, ok := provider.(Telegraph[T]); ok {
					sender = p
				}
				d.Dispatch(sender, receiver, message, bt.None[time.Duration](), false)
			})
		}
	}
}

func (d *Dispatcher[T]) Dispatch(
	sender Telegraph[T],
	receiver Telegraph[T],
	message T,
	delay bt.Optional[time.Duration],
	needsReturnReceipt bool,
) {
	if needsReturnReceipt {
		check.NotNil(sender)
	}

	telegram := Telegram[T]{
		sender:             sender,
		receiver:           receiver,
		message:            message,
		needsReturnReceipt: needsReturnReceipt,
	}

	if delay.Present() {
		dl := time.Now().Add(delay.Value())
		d.queue.Push(
			Delayed[T]{
				telegram: telegram,
				deadline: dl,
			},
			float32(dl.UnixMilli()),
		)
	} else {
		d.discharge(telegram)
	}
}

var _ Telegraph[any] = &Dispatcher[any]{}

func (d *Dispatcher[T]) HandleMessage(telegram Telegram[T]) bool {
	return false
}

func (d *Dispatcher[T]) discharge(telegram Telegram[T]) {
	if telegram.receiver != nil {
		if !telegram.receiver.HandleMessage(telegram) {
			d.log.Error(fmt.Sprintf("Message not handled: %v", telegram))
		}

	} else {
		numHandled := 0
		if ls := d.listeners.Get(reflect.TypeOf(telegram.message)); ls != nil {
			for _, listener := range ls {
				if listener.HandleMessage(telegram) {
					numHandled += 1
				}
			}
		}
		if numHandled < 1 {
			d.log.Error(fmt.Sprintf("Message not handled: %v", telegram))
		}
	}

	if telegram.needsReturnReceipt {
		receiptTelegram := Telegram[T]{
			sender:   d,
			receiver: telegram.sender,
			message:  telegram.message,
		}
		d.discharge(receiptTelegram)
	}
}

func (d *Dispatcher[T]) update() {
	for d.queue.Len() > 0 {
		cur := d.queue.Pop()
		if cur.Value.deadline.Before(time.Now()) {
			d.queue.Push(cur.Value, cur.Priority())
			break
		}
		d.discharge(cur.Value.telegram)
	}
}
