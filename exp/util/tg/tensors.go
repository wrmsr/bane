package tg

type Tensor struct {
	data *LazyBuffer

	//grad         *LazyBuffer
	//requiresGrad bool
}

func NewTensor(data *LazyBuffer) *Tensor {
	return &Tensor{
		data: data,
	}
}
