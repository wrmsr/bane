package rand

import (
	"fmt"
	"testing"
)

func TestPcg64(t *testing.T) {
	type tc struct {
		seed     uint64
		seed_gen []uint64
		expected []uint64
	}

	for _, c := range []tc{
		{
			seed: 4,
			seed_gen: []uint64{
				0x6c73632b58de42c6,
				0x23188e641dec4cb6,
				0x68acb4c4bbbea956,
				0x5476d87a1286562c,
			},
			expected: []uint64{
				0xf16c1ffbb9f8274e,
				0x82e65ccce1a329b5,
				0xf9eb1b84f0c08ec3,
				0x14b1ab6ef864ac36,
				0x9b7babfb741a13b3,
				0x60616cbf487b3d06,
				0xcd4965c2a081097b,
				0x2caddade94b61004,
				0xdf237d44ad0cd991,
				0x8b3fbe5f383b4614,
			},
		},

		{
			seed: 1337,
			seed_gen: []uint64{
				0xbc7e664d90aa98cc,
				0x83e7a3f304ca3335,
				0x959ba968f3163f00,
				0x8b7089338ccd3d3,
			},
			expected: []uint64{
				0xe0cb49408bdbe881,
				0x2f7ec2afba1d1a31,
				0xebc021d88a4ae104,
				0xf25223f4643aebe2,
				0xdfdfc3204a1d02e8,
				0x1da151557631462d,
				0x319865f20f4c9da9,
				0x577c16c72a3dcee8,
				0x7eec270646ccb3eb,
				0xe5fba83890a0281d,
			},
		},
	} {
		fmt.Printf("seed %x\n", c.seed)

		state := pcg64_state{
			pcg_state: &pcg64_random_t{},
		}

		pcg64_set_seed(&state, c.seed_gen[:2], c.seed_gen[2:])
		fmt.Printf("%+v\n", state)
		fmt.Printf("state.state: %s\n", state.pcg_state.state)
		fmt.Printf("state.inc: %s\n", state.pcg_state.inc)

		for _, e := range c.expected {
			a := pcg64_next64(&state)
			fmt.Printf("a %x e %x\n", a, e)
		}
	}
}
