package behave

import (
	"reflect"
	"time"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/maps"
	opt "github.com/wrmsr/bane/pkg/util/optional"
)

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
	telegram  Telegram[T]
	timestamp time.Time
}

type SimpleRegistry[K, O comparable] struct {
	setsByKey   map[K]maps.Set[O]
	addCallback func(O, K)
}

func NewSimpleRegistry[K, O comparable](addCallback func(O, K)) *SimpleRegistry[K, O] {
	return &SimpleRegistry[K, O]{
		setsByKey:   make(map[K]maps.Set[O]),
		addCallback: addCallback,
	}
}

func (r *SimpleRegistry[K, O]) Get(key K) maps.Set[O] {
	check.NotNil(key)
	return r.setsByKey[key]
}

func (r *SimpleRegistry[K, O]) Add(obj O, key K) {
	maps.ComputeDefault(r.setsByKey, key, maps.MakeSet[O]).Add(obj)
	if r.addCallback != nil {
		r.addCallback(obj, key)
	}
}

func (r *SimpleRegistry[K, O]) Remove(obj O, key K) {
	if key != nil {
		if s, ok := r.setsByKey[key]; ok {
			s.Remove(obj)
		}
	} else {
		for _, s := range r.setsByKey {
			s.Remove(obj)
		}
	}
}

func (r *SimpleRegistry[K, O]) Clear(key K) {
	if key != nil {
		delete(r.setsByKey, key)
	} else {
		r.setsByKey = make(map[K]maps.Set[O])
	}
}

/*

   class Dispatcher(Telegraph):

       def __init__(self) -> None:
           super().__init__()

           self._queue = queue.PriorityQueue()
           self._listeners: SimpleRegistry[type, Telegraph] = SimpleRegistry(add_callback=self._on_add_listener)
           self._providers: SimpleRegistry[type, TelegramProvider] = SimpleRegistry()

       @property
       def listeners(self) -> SimpleRegistry[type, Telegraph]:
           return self._listeners

       @property
       def providers(self) -> SimpleRegistry[type, TelegramProvider]:
           return self._providers

       def _on_add_listener(self, receiver: Telegraph, message_type: type) -> None:
           for provider in self._providers.get(message_type):
               message = provider.provide_message(receiver, message_type)
               if message is not None:
                   self.dispatch(
                       provider if isinstance(provider, Telegraph) else None,
                       receiver,
                       message,
                   )

       def dispatch(
               self,
               sender: ta.Optional[Telegraph],
               receiver: ta.Optional[Telegraph],
               message: T,
               *,
               delay: float = None,
               needs_return_receipt: bool = False,
       ) -> None:
           if needs_return_receipt:
               check.not_none(sender)

           telegram = Telegram(
               sender,
               receiver,
               message,
               needs_return_receipt=needs_return_receipt,
           )

           if not delay or delay <= 0.:
               self._discharge(telegram)
           else:
               self._queue.put(Delayed(telegram, time.time() + delay))

       def _discharge(self, telegram: Telegram) -> None:
           if telegram.receiver is not None:
               if not telegram.receiver.handle_message(telegram):
                   log.error(f'Message not handled: {telegram}')

           else:
               num_handled = 0
               for listener in self._listeners[type(telegram.message)]:
                   if listener.handle_message(telegram):
                       num_handled += 1
               if not num_handled:
                   log.error(f'Message not handled: {telegram}')

           if telegram.needs_return_receipt:
               receipt_telegram = Telegram(
                   self,
                   telegram.sender,
                   telegram,
               )
               self._discharge(receipt_telegram)

       def handle_message(self, telegram: Telegram) -> bool:
           return False

       def update(self) -> None:
           while not self._queue.empty():
               cur: Delayed = self._queue.get(block=False)
               if cur.timestamp > time.time():
                   self._queue.put(cur)
                   break
               self._discharge(cur.telegram)

*/
