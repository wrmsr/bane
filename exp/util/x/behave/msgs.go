package behave

import (
	"reflect"
	"time"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/log"
	opt "github.com/wrmsr/bane/pkg/util/optional"
	"github.com/wrmsr/bane/pkg/util/slices"
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
	HandleMessage(telegram Telegram[T]) bool
}

type TelegramProvider[T any] interface {
	ProvideMessage(receiver Telegraph[T], messageType reflect.Type) opt.Optional[T]
}

type Delayed[T any] struct {
	telegram Telegram[T]
	deadline time.Time
}

//

type Dispatcher[T any] struct {
	listeners Registry[reflect.Type, Telegraph[T]]
	providers Registry[reflect.Type, TelegramProvider[T]]

	queue slices.PriorityQueue[Delayed[T]]

	log log.DefaultLogger
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
		for provider := range s {
			provider.ProvideMessage(receiver, messageType).IfPresent(func(message T) {
				var sender Telegraph[T]
				if p, ok := provider.(Telegraph[T]); ok {
					sender = p
				}
				d.Dispatch(sender, receiver, message, opt.None[time.Duration](), false)
			})
		}
	}
}

func (d *Dispatcher[T]) Dispatch(
	sender Telegraph[T],
	receiver Telegraph[T],
	message T,
	delay opt.Optional[time.Duration],
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
		d.queue.Push(Delayed[T]{
			telegram: telegram,
			deadline: time.Now().Add(delay.Value()),
		})
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
		           d.log.Error('Message not handled: {telegram}')
				}

		   } else {
		       numHandled := 0
				ls := d.listeners[type(telegram.message)]
		       for listener := range self._listeners[reflect.TypeOf(telegram.message)] {
		           if listener.handle_message(telegram) {
		               numHandled += 1
					}
				}
		       if numHandled < 1 {
		           d.log.error(f'Message not handled: {telegram}')
				}
			}


/*
   if telegram.needs_return_receipt:
		       receipt_telegram = Telegram(
		           self,
		           telegram.sender,
		           telegram,
		       )
		       self._discharge(receipt_telegram)
	*/
}

func (d *Dispatcher[T]) update() {
	/*
	   while not self._queue.empty():
	       cur: Delayed = self._queue.get(block=False)
	       if cur.timestamp > time.time():
	           self._queue.put(cur)
	           break
	       self._discharge(cur.telegram)
	*/
}
