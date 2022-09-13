import gzip

import numpy as np
import tinygrad as tg
import tinygrad.helpers
import tinygrad.nn.optim


def get_parameters(obj):
    parameters = []
    if isinstance(obj, tg.tensor.Tensor):
        if obj.requires_grad:
            parameters.append(obj)
    elif isinstance(obj, list) or isinstance(obj, tuple):
        for x in obj:
            parameters.extend(get_parameters(x))
    elif hasattr(obj, '__dict__'):
        for v in obj.__dict__.values():
            parameters.extend(get_parameters(v))
    return parameters


def sparse_categorical_crossentropy(out, y):
    num_classes = out.shape[-1]
    yf = y.flatten()
    yz = np.zeros((yf.shape[0], num_classes), np.float32)
    # correct loss for NLL, torch NLL loss returns one per row
    yz[range(yz.shape[0]), yf] = -1.0 * num_classes
    yz = yz.reshape(list(y.shape) + [num_classes])
    yz = tg.tensor.Tensor(yz)
    return out.mul(yz).mean()


def train(
        model,
        x_train,
        y_train,
        optim,
        steps,
        *,
        bs=128,
        lossfn=sparse_categorical_crossentropy,
        transform=lambda x: x,
        target_transform=lambda x: x,
):
    tg.tensor.Tensor.training = True
    losses, accuracies = [], []
    for i in (t := range(steps)):
        samp = np.random.randint(0, x_train.shape[0], size=bs)
        x = tg.tensor.Tensor(transform(x_train[samp]), requires_grad=False)
        y = target_transform(y_train[samp])

        out = model.forward(x)

        loss = lossfn(out, y)
        optim.zero_grad()
        loss.backward()
        optim.step()

        cat = np.argmax(out.cpu().data, axis=-1)
        accuracy = (cat == y).mean()

        loss = loss.detach().cpu().data
        losses.append(loss)
        accuracies.append(accuracy)
        print("loss %f accuracy %f" % (loss, accuracy))


def evaluate(
        model,
        x_test,
        y_test,
        *,
        num_classes=None,
        bs=128,
        return_predict=False,
        transform=lambda x: x,
        target_transform=lambda y: y,
):
    tg.tensor.Tensor.training = False

    def numpy_eval(y_test, num_classes):
        y_test_preds_out = np.zeros(list(y_test.shape) + [num_classes])
        for i in range((len(y_test) - 1) // bs + 1):
            x = tg.tensor.Tensor(transform(x_test[i * bs:(i + 1) * bs]))
            y_test_preds_out[i * bs:(i + 1) * bs] = model.forward(x).cpu().data
        y_test_preds = np.argmax(y_test_preds_out, axis=-1)
        y_test = target_transform(y_test)
        return (y_test == y_test_preds).mean(), y_test_preds

    if num_classes is None:
        num_classes = y_test.max().astype(int) + 1
    acc, y_test_pred = numpy_eval(y_test, num_classes)
    print("test set accuracy is %f" % acc)
    return (acc, y_test_pred) if return_predict else acc


def uniform(*shape, **kwargs):
    return tg.tensor.Tensor(
        (np.random.uniform(-1., 1., size=shape)/np.sqrt(tg.helpers.prod(shape))).astype(np.float32),
        **kwargs,
    )


class TinyBobNet:
    def __init__(self):
        self.l1 = uniform(784, 128)
        self.l2 = uniform(128, 10)

    def parameters(self):
        return get_parameters(self)

    def forward(self, x):
        return x.dot(self.l1).relu().dot(self.l2).logsoftmax()


def fetch_mnist():
    def parse(file):
        return np.frombuffer(gzip.open(file).read(), dtype=np.uint8).copy()

    x_train = parse(".cache/datasets/mnist/train-images-idx3-ubyte.gz")[0x10:].reshape((-1, 28 * 28)).astype(np.float32)
    y_train = parse(".cache/datasets/mnist/train-labels-idx1-ubyte.gz")[8:]
    x_test = parse(".cache/datasets/mnist/t10k-images-idx3-ubyte.gz")[0x10:].reshape((-1, 28 * 28)).astype(np.float32)
    y_test = parse(".cache/datasets/mnist/t10k-labels-idx1-ubyte.gz")[8:]

    return x_train, y_train, x_test, y_test


def _main():
    x_train, y_train, x_test, y_test = fetch_mnist()
    np.random.seed(1337)
    model = TinyBobNet()
    optimizer = tg.nn.optim.SGD(model.parameters(), lr=0.001)
    train(model, x_train, y_train, optimizer, bs=69, steps=10)
    for p in model.parameters():
        p.realize()


if __name__ == '__main__':
    _main()
