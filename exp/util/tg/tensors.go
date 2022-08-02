package tg

type Tensor struct {
	data *LazyBuffer

	grad         *LazyBuffer
	requiresGrad bool

	ctx *FuncContext
}

func NewTensor(data *LazyBuffer, requiresGrad bool) *Tensor {
	return &Tensor{
		data: data,

		requiresGrad: requiresGrad,
	}
}

func (t *Tensor) Data() *LazyBuffer  { return t.data }
func (t *Tensor) RequiresGrad() bool { return t.requiresGrad }

func (t *Tensor) Shape() Shape {
	return t.data.st.Shape()
}

func BroadcastedTensor(fn func(x, y *Tensor) *Tensor, x, y *Tensor) *Tensor {
	// nDims := bt.Max(len(x.Shape()), len(y.Shape()))

	// if len(x.shape) != n_dims {
	// 	x = x.reshape(list(x.shape) + [1] * (n_dims - len(x.shape)))
	// }
	// if len(y.shape) != n_dims {
	// 	y = y.reshape(list(y.shape) + [1] * (n_dims - len(y.shape)))
	// }

	// shapeRet = tuple([max(sx, sy) for sx, sy in zip(x.shape, y.shape)])
	// if x.shape != shapeRet {
	//	x = x.expand(shape_ret)
	// }
	// if y.shape != shapeRet {
	//	y = y.expand(shape_ret)
	// }

	return fn(x, y)
}
