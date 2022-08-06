package behave

import (
	"reflect"
	"time"

	opt "github.com/wrmsr/bane/pkg/util/optional"
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
	telegram  Telegram[T]
	timestamp time.Time
}

//

/*

type Dispatcher[T any] struct {
	queue PriorityQueue
	listeners SimpleRegistry[
}

var _ Telegraph[any] = &Dispatcher[any]{}

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
