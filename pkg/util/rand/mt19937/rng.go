/*
# MT19937

Copyright (c) 2003-2005, Jean-Sebastien Roy (js@jeannot.org)

The rk_random and rk_seed functions algorithms and the original design of the Mersenne Twister RNG:

	Copyright (C) 1997 - 2002, Makoto Matsumoto and Takuji Nishimura, All rights reserved.

	Redistribution and use in source and binary forms, with or without modification, are permitted provided that the
	following conditions are met:

	1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following
	   disclaimer.

	2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following
	   disclaimer in the documentation and/or other materials provided with the distribution.

	3. The names of its contributors may not be used to endorse or promote products derived from this software without
	   specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES,
INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED.  IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

Original algorithm for the implementation of rk_interval function from Richard J. Wagner's implementation of the
Mersenne Twister RNG, optimised by Magnus Jonsson.

Constants used in the rk_double implementation by Isaku Wada.

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
documentation files (the "Software"), to deal in the Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit
persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the
Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
package mt19937

// parameters for computing Jump
const W_SIZE = 32 // size of uint64
const MEXP = 19937
const P_SIZE = ((MEXP / W_SIZE) + 1)
const LSB = 0x00000001
const QQ = 7
const LL = 128 // LL = 2^(QQ)

// 2**128 step polynomial produced using the file mt19937-generate-jump-poly.c
// (randomgen) which is a modified version of minipoly_mt19937.c as distributed
// in
// http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/JUMP/jump_ahead_1.02.tar.gz
var poly_coef = [624]uint64{
	1927166307, 3044056772, 2284297142, 2820929765, 651705945,
	69149273, 3892165397, 2337412983, 1219880790, 3207074517,
	3836784057, 189286826, 1049791363, 3916249550, 2942382547,
	166392552, 861176918, 3246476411, 2302311555, 4273801148,
	29196903, 1363664063, 3802562022, 2600400244, 3090369801,
	4040416970, 1432485208, 3632558139, 4015816763, 3013316418,
	551532385, 3592224467, 3479125595, 1195467127, 2391032553,
	2393493419, 1482493632, 1625159565, 748389672, 4042774030,
	2998615036, 3393119101, 2177492569, 2265897321, 2507383006,
	3461498961, 2003319700, 1942857197, 1455226044, 4097545580,
	529653268, 3204756480, 2486748289, 495294513, 3396001954,
	2643963605, 2655404568, 3881604377, 624710790, 3443737948,
	1941294296, 2139259604, 3368734020, 422436761, 3602810182,
	1384691081, 3035786407, 2551797119, 537227499, 65486120,
	642436100, 2023822537, 2515598203, 1122953367, 2882306242,
	1743213032, 321965189, 336496623, 2436602518, 3556266590,
	1055117829, 463541647, 743234441, 527083645, 2606668346,
	2274046499, 2761475053, 2760669048, 2538258534, 487125077,
	3365962306, 3604906217, 2714700608, 680709708, 2217161159,
	1614899374, 3710119533, 3201300658, 3752620679, 2755041105,
	3129723037, 1247297753, 2812642690, 4114340845, 3485092247,
	2752814364, 3586551747, 4073138437, 3462966585, 2924318358,
	4061374901, 3314086806, 2640385723, 744590670, 3007586513,
	3959120371, 997207767, 3420235506, 2092400998, 3190305685,
	60965738, 549507222, 3784354415, 3209279509, 1238863299,
	2605037827, 178570440, 1743491299, 4079686640, 2136795825,
	3435430548, 1679732443, 1835708342, 2159367000, 1924487218,
	4059723674, 996192116, 2308091645, 1336281586, 674600050,
	1642572529, 1383973289, 2202960007, 3165481279, 3385474038,
	2501318550, 2671842890, 3084085109, 3475033915, 1551329147,
	4101397249, 1205851807, 3641536021, 3607635071, 1609126163,
	2910426664, 3324508658, 4244311266, 254034382, 1258304384,
	1914048768, 1358592011, 527610138, 3072108727, 4289413885,
	1417001678, 2445445945, 896462712, 339855811, 3699378285,
	2529457297, 3049459401, 2723472429, 2838633181, 2520397330,
	3272339035, 1667003847, 3742634787, 942706520, 2301027215,
	1907791250, 2306299096, 1021173342, 1539334516, 2907834628,
	3199959207, 1556251860, 3642580275, 2355865416, 285806145,
	867932457, 1177354172, 3291107470, 4022765061, 1613380116,
	588147929, 650574324, 1236855601, 1371354511, 2085218212,
	1203081931, 420526905, 1022192219, 2903287064, 2470845899,
	3649873273, 2502333582, 3972385637, 4246356763, 199084157,
	1567178788, 2107121836, 4293612856, 1902910177, 332397359,
	83422598, 3614961721, 456321943, 2277615967, 2302518510,
	3258315116, 2521897172, 3900282042, 4186973154, 3146532165,
	2299685029, 3889120948, 1293301857, 187455105, 3395849230,
	913321567, 3093513909, 1440944571, 1923481911, 338680924,
	1204882963, 2739724491, 2886241328, 2408907774, 1299817192,
	2474012871, 45400213, 553186784, 134558656, 2180943666,
	2870807589, 76511085, 3053566760, 2516601415, 4172865902,
	1751297915, 1251975234, 2964780642, 1412975316, 2739978478,
	2171013719, 637935041, 975972384, 3044407449, 3111425639,
	1938684970, 2860857400, 13419586, 2772079268, 3484375614,
	3184054178, 159924837, 1386213021, 2765617231, 2523689118,
	1283505218, 3510789588, 4125878259, 2990287597, 2152014833,
	3084155970, 2815101609, 1932985704, 114887365, 1712687646,
	2550515629, 3299051916, 2022747614, 2143630992, 2244188960,
	3309469192, 3234358520, 800720365, 3278176634, 554357439,
	2415629802, 1620877315, 2389462898, 2229691332, 1007748450,
	1966873768, 2264971043, 1214524156, 346854700, 3471905342,
	3984889660, 4034246840, 216712649, 4027196762, 3754772604,
	2121785562, 2347070732, 7457687, 1443375102, 683948143,
	2940226032, 3211475670, 2836507357, 774899409, 1588968308,
	780438009, 3278878781, 2217181540, 2184194887, 1642129086,
	69346830, 297114710, 3841068188, 2631265450, 4167492314,
	2613519651, 1388582503, 2171556668, 1201873758, 2698772382,
	207791958, 3936134563, 3725025702, 3306317801, 1055730422,
	4069230694, 1767821343, 4252407395, 2422583118, 3158834399,
	3754582617, 1112422556, 376187931, 3137549150, 712221089,
	3300799453, 3868250200, 1165257666, 2494837767, 131304831,
	1619349427, 1958236644, 3678218946, 3651007751, 2261987899,
	1567368524, 2193599522, 3034394674, 2994602555, 3072727647,
	889094521, 1089692095, 1822324824, 3876999182, 1703361286,
	902229515, 4213728487, 3838170364, 672727494, 2240733828,
	3858539469, 1149254245, 4166055926, 4193525313, 1709921593,
	2278290377, 3190784116, 2919588882, 1012709717, 3640562031,
	2931984863, 3515665246, 250577343, 1147230194, 1183856202,
	3734511989, 3243867808, 3499383067, 2985115159, 2036821626,
	3298159553, 2726542838, 1686910320, 1778823772, 965412224,
	233509772, 3843098861, 1312622954, 500855830, 2950562091,
	1915683607, 3405781138, 596073719, 2195150546, 3381728478,
	546426436, 3527890868, 2324975353, 2241074266, 3992514859,
	2576108287, 4077653225, 2632319392, 3127212632, 917000669,
	2498161805, 3980835128, 2259526768, 1083920509, 1187452089,
	97018536, 3056075838, 2059706760, 2373335692, 182196406,
	2136713111, 1762080153, 1572125803, 1145919955, 1023966754,
	3921694345, 1632005969, 1418372326, 354407429, 2438288265,
	1620072033, 1586320921, 1044153697, 969324572, 613487980,
	4230993062, 397726764, 2194259193, 735511759, 2066049260,
	88093248, 1562536153, 2114157419, 3630951546, 589238503,
	3120654384, 2521793793, 2746692127, 2557723425, 889897693,
	2778878177, 643269509, 3342389831, 19218890, 3442706236,
	3314581273, 3503147052, 1546343434, 1448529060, 529038801,
	2748942264, 2213019208, 111314040, 2488697563, 1180642808,
	2605272289, 4207476668, 1502558669, 2972370981, 4204339995,
	1046225278, 992840610, 3847290298, 2387673094, 2221565747,
	1045901716, 3997739302, 1556952765, 1103336648, 279418400,
	2711316466, 2336215718, 2317900806, 974624729, 909575434,
	1675610631, 1922393214, 2054896570, 3197007361, 3932554569,
	1008619802, 3349254938, 113511461, 932630384, 2098759268,
	3436837432, 3119972401, 1612590197, 2281609013, 4174211248,
	4016332246, 2097525539, 1398632760, 1543697535, 2419227174,
	1676465074, 2882923045, 23216933, 808195649, 3690720147,
	484419260, 2254772642, 2975434733, 288528113, 204598404,
	589968818, 3021152400, 2463155141, 1397846755, 157285579,
	4230258857, 2469135246, 625357422, 3435224647, 465239124,
	1022535736, 2823317040, 274194469, 2214966446, 3661001613,
	518802547, 2293436304, 1335881988, 2247010176, 1856732584,
	1088028094, 1877563709, 1015352636, 1700817932, 2960695857,
	1882229300, 1666906557, 1838841022, 3983797810, 1667630361,
	385998221, 241341791, 403550441, 2629200403, 3552759102,
	2029750442, 2247999048, 2726665298, 2507798776, 2419064129,
	1266444923, 526255242, 2384866697, 1886200981, 3954956408,
	2171436866, 2295200753, 1047315850, 1967809707, 2860382973,
	3918334466, 3057439479, 952682588, 1925559679, 3112119050,
	3833190964, 1430139895, 2089165610, 3009202424, 3989186157,
	3395807230, 347600520, 120428923, 3017004655, 1384933954,
	303039929, 234010146, 2278760249, 315514836, 3987659575,
	1239335668, 2387869477, 3885908826, 1983922602, 698609264,
	3009002846, 1520611399, 809159940, 3089771783, 374838722,
	2789914419, 2500831937, 3751970335, 4279852547, 2362894437,
	1588814060, 1671213155, 434218829, 2126587176, 2002526422,
	2756464095, 141700479, 2965974322, 2211530172, 992085992,
	1943691492, 2705131817, 2519208889, 1938768395, 3949294294,
	354046666, 2158272751, 602858583, 0,
}

const RK_STATE_LEN = 624

const N = 624
const M = 397
const MATRIX_A = 0x9908b0df
const UPPER_MASK = uint32(0x80000000)
const LOWER_MASK = uint32(0x7fffffff)

type mt19937_state struct {
	key [RK_STATE_LEN]uint32
	pos int
}

// 32-bits function
// return the i-th coefficient of the polynomial pf
func get_coef(pf []uint64, deg uint) uint64 {
	if (pf[deg>>5] & (LSB << (deg & 0x1f))) != 0 {
		return 1
	}
	return 0
}

func copy_state(target_state *mt19937_state, state *mt19937_state) {
	var i int

	for i = 0; i < N; i++ {
		target_state.key[i] = state.key[i]
	}

	target_state.pos = state.pos
}

var mag02 = [2]uint64{0x0, MATRIX_A}

// next state generating function
func gen_next(state *mt19937_state) {
	var num int
	var y uint64

	num = state.pos
	if num < N-M {
		y = uint64(state.key[num]&UPPER_MASK) | uint64(state.key[num+1]&LOWER_MASK)
		state.key[num] = uint32(uint64(state.key[num+M]) ^ (y >> 1) ^ mag02[y%2])
		state.pos++
	} else if num < N-1 {
		y = uint64(state.key[num]&UPPER_MASK) | uint64(state.key[num+1]&LOWER_MASK)
		state.key[num] = uint32(uint64(state.key[num+(M-N)]) ^ (y >> 1) ^ mag02[y%2])
		state.pos++
	} else if num == N-1 {
		y = uint64(state.key[N-1]&UPPER_MASK) | uint64(state.key[0]&LOWER_MASK)
		state.key[N-1] = uint32(uint64(state.key[M-1]) ^ (y >> 1) ^ mag02[y%2])
		state.pos = 0
	}
}

func add_state(state1 *mt19937_state, state2 *mt19937_state) {
	i := 0
	pt1, pt2 := state1.pos, state2.pos

	if pt2-pt1 >= 0 {
		for i = 0; i < N-pt2; i++ {
			state1.key[i+pt1] ^= state2.key[i+pt2]
		}
		for ; i < N-pt1; i++ {
			state1.key[i+pt1] ^= state2.key[i+(pt2-N)]
		}
		for ; i < N; i++ {
			state1.key[i+(pt1-N)] ^= state2.key[i+(pt2-N)]
		}
	} else {
		for i = 0; i < N-pt1; i++ {
			state1.key[i+pt1] ^= state2.key[i+pt2]
		}
		for ; i < N-pt2; i++ {
			state1.key[i+(pt1-N)] ^= state2.key[i+pt2]
		}
		for ; i < N; i++ {
			state1.key[i+(pt1-N)] ^= state2.key[i+(pt2-N)]
		}
	}
}

// compute pf(ss) using standard Horner method
func horner1(pf []uint64, state *mt19937_state) {
	i := MEXP - 1
	var temp *mt19937_state

	temp = &mt19937_state{}

	for get_coef(pf, uint(i)) == 0 {
		i--
	}

	if i > 0 {
		copy_state(temp, state)
		gen_next(temp)
		i--
		for ; i > 0; i-- {
			if get_coef(pf, uint(i)) != 0 {
				add_state(temp, state)
			}
			gen_next(temp)
		}
		if get_coef(pf, 0) != 0 {
			add_state(temp, state)
		}
	} else if i == 0 {
		copy_state(temp, state)
	}

	copy_state(state, temp)
}

func mt19937_jump_state(state *mt19937_state) {
	var pf []uint64
	var i int

	pf = make([]uint64, P_SIZE)
	for i = 0; i < P_SIZE; i++ {
		pf[i] = poly_coef[i]
	}

	if state.pos >= N {
		state.pos = 0
	}

	horner1(pf, state)
}

/*
// Slightly optimized reference implementation of the Mersenne Twister
uint32_t mt19937_next(mt19937_state *state) {
  uint32_t y;

  if (state.pos == RK_STATE_LEN) {
    // Move to function to help inlining
    mt19937_gen(state);
  }
  y = state.key[state.pos++];

  // Tempering
  y ^= (y >> 11);
  y ^= (y << 7) & 0x9d2c5680UL;
  y ^= (y << 15) & 0xefc60000UL;
  y ^= (y >> 18);

  return y;
}

uint64 mt19937_next64(mt19937_state *state) {
  return (uint64)mt19937_next(state) << 32 | mt19937_next(state);
}

uint32_t mt19937_next32(mt19937_state *state) {
  return mt19937_next(state);
}

double mt19937_next_double(mt19937_state *state) {
  int32_t a = mt19937_next(state) >> 5, b = mt19937_next(state) >> 6;
  return (a * 67108864.0 + b) / 9007199254740992.0;
}

func mt19937_seed(mt19937_state *state, uint32_t seed) {
  int pos;
  seed &= 0xffffffffUL;

  // Knuth's PRNG as used in the Mersenne Twister reference implementation
  for (pos = 0; pos < RK_STATE_LEN; pos++) {
    state.key[pos] = seed;
    seed = (1812433253UL * (seed ^ (seed >> 30)) + pos + 1) & 0xffffffffUL;
  }
  state.pos = RK_STATE_LEN;
}

// initializes mt[RK_STATE_LEN] with a seed
func init_genrand(mt19937_state *state, uint32_t s) {
  int mti;
  uint32_t *mt = state.key;

  mt[0] = s & 0xffffffffUL;
  for (mti = 1; mti < RK_STATE_LEN; mti++) {
    // See Knuth TAOCP Vol2. 3rd Ed. P.106 for multiplier.
    // In the previous versions, MSBs of the seed affect
    // only MSBs of the array mt[].
    // 2002/01/09 modified by Makoto Matsumoto
    mt[mti] = (1812433253UL * (mt[mti - 1] ^ (mt[mti - 1] >> 30)) + mti);
    // for > 32 bit machines
    mt[mti] &= 0xffffffffUL;
  }
  state.pos = mti;
  return;
}

// initialize by an array with array-length
// init_key is the array for initializing keys
// key_length is its length
func mt19937_init_by_array(mt19937_state *state, uint32_t *init_key, int key_length) {
  // was signed in the original code. RDH 12/16/2002
  int i = 1;
  int j = 0;
  uint32_t *mt = state.key;
  int k;

  init_genrand(state, 19650218UL);
  k = (RK_STATE_LEN > key_length ? RK_STATE_LEN : key_length);
  for (; k; k--) {
    // non linear
    mt[i] = (mt[i] ^ ((mt[i - 1] ^ (mt[i - 1] >> 30)) * 1664525UL)) + init_key[j] + j;
    // for > 32 bit machines
    mt[i] &= 0xffffffffUL;
    i++;
    j++;
    if (i >= RK_STATE_LEN) {
      mt[0] = mt[RK_STATE_LEN - 1];
      i = 1;
    }
    if (j >= key_length) {
      j = 0;
    }
  }
  for (k = RK_STATE_LEN - 1; k; k--) {
    mt[i] = (mt[i] ^ ((mt[i - 1] ^ (mt[i - 1] >> 30)) * 1566083941UL)) - i; // non linear
    mt[i] &= 0xffffffffUL; // for WORDSIZE > 32 machines
    i++;
    if (i >= RK_STATE_LEN) {
      mt[0] = mt[RK_STATE_LEN - 1];
      i = 1;
    }
  }

  mt[0] = 0x80000000UL; // MSB is 1; assuring non-zero initial array
}

func mt19937_gen(mt19937_state *state) {
  uint32_t y;
  int i;

  for (i = 0; i < N - M; i++) {
    y = (state.key[i] & UPPER_MASK) | (state.key[i + 1] & LOWER_MASK);
    state.key[i] = state.key[i + M] ^ (y >> 1) ^ (-(y & 1) & MATRIX_A);
  }
  for (; i < N - 1; i++) {
    y = (state.key[i] & UPPER_MASK) | (state.key[i + 1] & LOWER_MASK);
    state.key[i] = state.key[i + (M - N)] ^ (y >> 1) ^ (-(y & 1) & MATRIX_A);
  }
  y = (state.key[N - 1] & UPPER_MASK) | (state.key[0] & LOWER_MASK);
  state.key[N - 1] = state.key[M - 1] ^ (y >> 1) ^ (-(y & 1) & MATRIX_A);

  state.pos = 0;
}

func mt19937_jump(mt19937_state *state) { mt19937_jump_state(state); }
*/
