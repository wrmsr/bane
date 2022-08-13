package tensor

import (
	"fmt"
	"strings"
	"testing"
)

/*
a = np.arange(60.).reshape(3,4,5)
b = np.arange(24.).reshape(4,3,2)
c = np.tensordot(a,b, axes=([1,0],[0,1]))
c.shape
(5, 2)
c
array([[4400., 4730.],
	   [4532., 4874.],
	   [4664., 5018.],
	   [4796., 5162.],
	   [4928., 5306.]])
# A slower but equivalent way of computing the same...
d = np.zeros((5,2))
for i in range(5):
  for j in range(2):
	for k in range(3):
	  for n in range(4):
		d[i,j] += a[k,n,i] * b[n,k,j]
*/

func TestTensorDot(t *testing.T) {
	a := [][][]int{
		{
			{1, 2},
			{3, 4},
		},
		{
			{5, 6},
			{7, 8},
		},
	}

	b := [][]string{
		{"a", "b"},
		{"c", "d"},
	}

	var c [][]string = make([][]string, 3)
	for i := 0; i < 3; i++ {
		c[i] = make([]string, 3)
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			var o strings.Builder
			for k := 0; k < 2; k++ {
				o.WriteString(strings.Repeat(b[k][j], a[0][i][k]))
			}
			c[i][j] = o.String()
		}
	}

	fmt.Println(c)

	/*
		axes = 0 : tensor product
		axes = 1 : tensor dot product
		axes = 2 : (default) tensor double contraction
	*/

	/*
		np.tensordot(a, A) # third argument default is 2 for double-contraction
		array(['abbcccdddd', 'aaaaabbbbbbcccccccdddddddd'])

		np.tensordot(a, A, 1)
		array([[['acc', 'bdd'],
				['aaacccc', 'bbbdddd']],
			   [['aaaaacccccc', 'bbbbbdddddd'],
				['aaaaaaacccccccc', 'bbbbbbbdddddddd']]])

		np.tensordot(a, A, 0) # tensor product
		array([[[[['a', 'b'],
				  ['c', 'd']],

				 [['aa', 'bb'],
				  ['cc', 'dd']]],


				[[['aaa', 'bbb'],
				  ['ccc', 'ddd']],

				 [['aaaa', 'bbbb'],
				  ['cccc', 'dddd']]]],


			   [[[['aaaaa', 'bbbbb'],
				  ['ccccc', 'ddddd']],

				 [['aaaaaa', 'bbbbbb'],
				  ['cccccc', 'dddddd']]],


				[[['aaaaaaa', 'bbbbbbb'],
				  ['ccccccc', 'ddddddd']],

				 [['aaaaaaaa', 'bbbbbbbb'],
				  ['cccccccc', 'dddddddd']]]]])

		np.tensordot(a, A, (0, 1))
		array([[['abbbbb', 'cddddd'],
				['aabbbbbb', 'ccdddddd']],
			   [['aaabbbbbbb', 'cccddddddd'],
				['aaaabbbbbbbb', 'ccccdddddddd']]])

		np.tensordot(a, A, (2, 1))
		array([[['abb', 'cdd'],
				['aaabbbb', 'cccdddd']],
			   [['aaaaabbbbbb', 'cccccdddddd'],
				['aaaaaaabbbbbbbb', 'cccccccdddddddd']]])

		np.tensordot(a, A, ((0, 1), (0, 1)))
		array(['abbbcccccddddddd', 'aabbbbccccccdddddddd'])

		np.tensordot(a, A, ((1, 0), (0, 1)))
		array(['abbbbbcccddddddd', 'aabbbbbbccccdddddddd'], dtype=object)

			np.tensordot(a, A, ((2, 1), (1, 0)))
		array(['acccbbdddd', 'aaaaacccccccbbbbbbdddddddd'])
	*/
}
