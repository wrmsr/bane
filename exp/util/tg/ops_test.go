package tg

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/slices"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

func TestOps1(t *testing.T) {
	//a := -.5
	//b := 20.
	//rand.Seed(0)

	for _, sh := range []Shape{
		{9, 1},
		{3, 3},
	} {
		xt := NewTensor(MakeLoadBuffer(BufferOf(sh, bt.RangeTo[float32](9.).Slice()), sh), true)
		yt := NewTensor(MakeLoadBuffer(BufferOf(sh, bt.RangeOf[float32](10., 19., 1.).Slice()), sh), true)

		zt := xt.Add(yt)
		fmt.Println(zt)

		fmt.Println(zt.Data().Realize())

		zt.Mean(nil, false).Backward()

		zg := zt.grad.Data()

		zgt := zg.Realize()
		fmt.Println(zgt)
	}
}

func TestOpsRelu(t *testing.T) {
	for _, sh := range []Shape{
		{9, 1},
		{3, 3},
	} {
		xs := BufferOf(sh, slices.Map(func(v float32) float32 { return (v / 9) - .5 }, bt.RangeTo[float32](9.).Slice()))

		xt := NewTensor(MakeLoadBuffer(xs, sh), true)

		zt := xt.Relu()
		fmt.Println(zt)

		fmt.Println(zt.Data().Realize())

		zt.Mean(nil, false).Backward()

		zg := zt.grad.Data()

		zgt := zg.Realize()
		fmt.Println(zgt)
	}
}

/*
[<LoadOps.FROMCPU: 1>] :  -> (9,)
[0. 1. 2. 3. 4. 5. 6. 7. 8.]
[<MovementOps.RESHAPE: 1>] : (9,) -> (3, 3)
[[0. 1. 2.]
 [3. 4. 5.]
 [6. 7. 8.]]
[<MovementOps.PERMUTE: 2>] : (3, 3) -> (3, 3)
[[0. 3. 6.]
 [1. 4. 7.]
 [2. 5. 8.]]
[<MovementOps.RESHAPE: 1>] : (3, 3) -> (1, 3, 3, 1)
[[[[0.]
   [3.]
   [6.]]
  [[1.]
   [4.]
   [7.]]
  [[2.]
   [5.]
   [8.]]]]
[<LoadOps.FROMCPU: 1>] :  -> (9,)
[10. 11. 12. 13. 14. 15. 16. 17. 18.]
[<MovementOps.RESHAPE: 1>] : (9,) -> (3, 3)
[[10. 11. 12.]
 [13. 14. 15.]
 [16. 17. 18.]]
[<MovementOps.PERMUTE: 2>] : (3, 3) -> (3, 3)
[[10. 13. 16.]
 [11. 14. 17.]
 [12. 15. 18.]]
[<MovementOps.RESHAPE: 1>] : (3, 3) -> (3, 3, 1, 1)
[[[[10.]]
  [[13.]]
  [[16.]]]
 [[[11.]]
  [[14.]]
  [[17.]]]
 [[[12.]]
  [[15.]]
  [[18.]]]]
[<ProcessingOps.CONV: 1>] : (1, 3, 3, 1), (3, 3, 1, 1) -> (1, 3, 3, 1)
[[[[ 45.]
   [162.]
   [279.]]
  [[ 48.]
   [174.]
   [300.]]
  [[ 51.]
   [186.]
   [321.]]]]
[<MovementOps.RESHAPE: 1>] : (1, 3, 3, 1) -> (3, 3)
[[ 45. 162. 279.]
 [ 48. 174. 300.]
 [ 51. 186. 321.]]
[<MovementOps.PERMUTE: 2>] : (3, 3) -> (3, 3)
[[ 45.  48.  51.]
 [162. 174. 186.]
 [279. 300. 321.]]
[<UnaryOps.RELU: 3>] : (3, 3) -> (3, 3)
[[ 45.  48.  51.]
 [162. 174. 186.]
 [279. 300. 321.]]
==== WRONG NOW ====
[<MovementOps.PERMUTE: 2>] : (3, 3) -> (3, 3)
[[ 45. 162. 279.]
 [ 48. 174. 300.]
 [ 51. 186. 321.]]
[<MovementOps.RESHAPE: 1>] : (3, 3) -> (1, 3, 3, 1)
[[[[ 45.]
   [162.]
   [279.]]
  [[ 48.]
   [174.]
   [300.]]
  [[ 51.]
   [186.]
   [321.]]]]
[<LoadOps.FROMCPU: 1>] :  -> (9,)
[20. 21. 22. 23. 24. 25. 26. 27. 28.]
[<MovementOps.RESHAPE: 1>] : (9,) -> (3, 3)
[[20. 21. 22.]
 [23. 24. 25.]
 [26. 27. 28.]]
[<MovementOps.PERMUTE: 2>] : (3, 3) -> (3, 3)
[[20. 23. 26.]
 [21. 24. 27.]
 [22. 25. 28.]]
[<MovementOps.RESHAPE: 1>] : (3, 3) -> (3, 3, 1, 1)
[[[[20.]]
  [[23.]]
  [[26.]]]
 [[[21.]]
  [[24.]]
  [[27.]]]
 [[[22.]]
  [[25.]]
  [[28.]]]]
[<ProcessingOps.CONV: 1>] : (1, 3, 3, 1), (3, 3, 1, 1) -> (1, 3, 3, 1)
[[[[ 3330.]
   [12078.]
   [20826.]]
  [[ 3474.]
   [12600.]
   [21726.]]
  [[ 3618.]
   [13122.]
   [22626.]]]]
[<MovementOps.RESHAPE: 1>] : (1, 3, 3, 1) -> (3, 3)
[[ 3330. 12078. 20826.]
 [ 3474. 12600. 21726.]
 [ 3618. 13122. 22626.]]
[<MovementOps.PERMUTE: 2>] : (3, 3) -> (3, 3)
[[ 3330.  3474.  3618.]
 [12078. 12600. 13122.]
 [20826. 21726. 22626.]]
[<ReduceOps.MAX: 2>] : (3, 3) -> (3, 1)
[[ 3618.]
 [13122.]
 [22626.]]
[<MovementOps.EXPAND: 4>] : (3, 1) -> (3, 3)
[[ 3618.  3618.  3618.]
 [13122. 13122. 13122.]
 [22626. 22626. 22626.]]
[<BinaryOps.SUB: 2>] : (3, 3), (3, 3) -> (3, 3)
[[ -288.  -144.     0.]
 [-1044.  -522.     0.]
 [-1800.  -900.     0.]]
[<UnaryOps.EXP: 4>, <BinaryOps.SUB: 2>] : (3, 3), (3, 3) -> (3, 3)
[[0. 0. 1.]
 [0. 0. 1.]
 [0. 0. 1.]]
[<ReduceOps.SUM: 1>] : (3, 3) -> (3, 1)
[[1.]
 [1.]
 [1.]]
[<UnaryOps.LOG: 5>] : (3, 1) -> (3, 1)
[[0.]
 [0.]
 [0.]]
[<MovementOps.EXPAND: 4>] : (3, 1) -> (3, 3)
[[0. 0. 0.]
 [0. 0. 0.]
 [0. 0. 0.]]
[<BinaryOps.SUB: 2>] : (3, 3), (3, 3) -> (3, 3)
[[ -288.  -144.     0.]
 [-1044.  -522.     0.]
 [-1800.  -900.     0.]]
[[ -288.  -144.     0.]
 [-1044.  -522.     0.]
 [-1800.  -900.     0.]]
*/

func TestBobNet(t *testing.T) {
	sh := Shape{3, 3}

	xt := NewTensor(MakeLoadBuffer(BufferOf(sh, bt.RangeTo[float32](9.).Slice()), Shape{9}), true).Reshape(sh)
	//l1t := NewTensor(MakeLoadBuffer(BufferOf(sh, bt.RangeOf[float32](10., 19., 1.).Slice()), Shape{9}), true).Reshape(sh)
	//l2t := NewTensor(MakeLoadBuffer(BufferOf(sh, bt.RangeOf[float32](20., 29., 1.).Slice()), Shape{9}), true).Reshape(sh)

	//zt := xt.Dot(l1t).Relu().Dot(l2t).LogSoftmax()
	zt := xt.LogSoftmax()
	fmt.Println("====")

	var rec func(any, string)
	rec = func(o any, p string) {
		switch o := o.(type) {
		case *Tensor:
			rec(o.data, p)
		case *LazyBuffer:
			rec(o.op, p)
		case *LazyOp:
			fmt.Printf("%s%v\n", p, o.op)
			for _, c := range o.srcs {
				rec(c, "  "+p)
			}
		default:
			panic(o)
		}
	}
	rec(zt, "")

	fmt.Println(zt.Data().Realize())
}
