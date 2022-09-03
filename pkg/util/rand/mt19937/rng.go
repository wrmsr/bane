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

/*
#pragma once
#include "mt19937.h"
#include <stdlib.h>

// parameters for computing Jump
#define W_SIZE 32 // size of unsigned long
#define MEXP 19937
#define P_SIZE ((MEXP / W_SIZE) + 1)
#define LSB 0x00000001UL
#define QQ 7
#define LL 128 // LL = 2^(QQ)

void mt19937_jump_state(mt19937_state *state);

void set_coef(unsigned long *pf, unsigned int deg, unsigned long v);

// 2**128 step polynomial produced using the file mt19937-generate-jump-poly.c
// (randomgen) which is a modified version of minipoly_mt19937.c as distributed
// in
// http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/JUMP/jump_ahead_1.02.tar.gz
//
// These files are not part of NumPy.

static const unsigned long poly_coef[624] = {
    1927166307UL, 3044056772UL, 2284297142UL, 2820929765UL, 651705945UL,
    69149273UL,   3892165397UL, 2337412983UL, 1219880790UL, 3207074517UL,
    3836784057UL, 189286826UL,  1049791363UL, 3916249550UL, 2942382547UL,
    166392552UL,  861176918UL,  3246476411UL, 2302311555UL, 4273801148UL,
    29196903UL,   1363664063UL, 3802562022UL, 2600400244UL, 3090369801UL,
    4040416970UL, 1432485208UL, 3632558139UL, 4015816763UL, 3013316418UL,
    551532385UL,  3592224467UL, 3479125595UL, 1195467127UL, 2391032553UL,
    2393493419UL, 1482493632UL, 1625159565UL, 748389672UL,  4042774030UL,
    2998615036UL, 3393119101UL, 2177492569UL, 2265897321UL, 2507383006UL,
    3461498961UL, 2003319700UL, 1942857197UL, 1455226044UL, 4097545580UL,
    529653268UL,  3204756480UL, 2486748289UL, 495294513UL,  3396001954UL,
    2643963605UL, 2655404568UL, 3881604377UL, 624710790UL,  3443737948UL,
    1941294296UL, 2139259604UL, 3368734020UL, 422436761UL,  3602810182UL,
    1384691081UL, 3035786407UL, 2551797119UL, 537227499UL,  65486120UL,
    642436100UL,  2023822537UL, 2515598203UL, 1122953367UL, 2882306242UL,
    1743213032UL, 321965189UL,  336496623UL,  2436602518UL, 3556266590UL,
    1055117829UL, 463541647UL,  743234441UL,  527083645UL,  2606668346UL,
    2274046499UL, 2761475053UL, 2760669048UL, 2538258534UL, 487125077UL,
    3365962306UL, 3604906217UL, 2714700608UL, 680709708UL,  2217161159UL,
    1614899374UL, 3710119533UL, 3201300658UL, 3752620679UL, 2755041105UL,
    3129723037UL, 1247297753UL, 2812642690UL, 4114340845UL, 3485092247UL,
    2752814364UL, 3586551747UL, 4073138437UL, 3462966585UL, 2924318358UL,
    4061374901UL, 3314086806UL, 2640385723UL, 744590670UL,  3007586513UL,
    3959120371UL, 997207767UL,  3420235506UL, 2092400998UL, 3190305685UL,
    60965738UL,   549507222UL,  3784354415UL, 3209279509UL, 1238863299UL,
    2605037827UL, 178570440UL,  1743491299UL, 4079686640UL, 2136795825UL,
    3435430548UL, 1679732443UL, 1835708342UL, 2159367000UL, 1924487218UL,
    4059723674UL, 996192116UL,  2308091645UL, 1336281586UL, 674600050UL,
    1642572529UL, 1383973289UL, 2202960007UL, 3165481279UL, 3385474038UL,
    2501318550UL, 2671842890UL, 3084085109UL, 3475033915UL, 1551329147UL,
    4101397249UL, 1205851807UL, 3641536021UL, 3607635071UL, 1609126163UL,
    2910426664UL, 3324508658UL, 4244311266UL, 254034382UL,  1258304384UL,
    1914048768UL, 1358592011UL, 527610138UL,  3072108727UL, 4289413885UL,
    1417001678UL, 2445445945UL, 896462712UL,  339855811UL,  3699378285UL,
    2529457297UL, 3049459401UL, 2723472429UL, 2838633181UL, 2520397330UL,
    3272339035UL, 1667003847UL, 3742634787UL, 942706520UL,  2301027215UL,
    1907791250UL, 2306299096UL, 1021173342UL, 1539334516UL, 2907834628UL,
    3199959207UL, 1556251860UL, 3642580275UL, 2355865416UL, 285806145UL,
    867932457UL,  1177354172UL, 3291107470UL, 4022765061UL, 1613380116UL,
    588147929UL,  650574324UL,  1236855601UL, 1371354511UL, 2085218212UL,
    1203081931UL, 420526905UL,  1022192219UL, 2903287064UL, 2470845899UL,
    3649873273UL, 2502333582UL, 3972385637UL, 4246356763UL, 199084157UL,
    1567178788UL, 2107121836UL, 4293612856UL, 1902910177UL, 332397359UL,
    83422598UL,   3614961721UL, 456321943UL,  2277615967UL, 2302518510UL,
    3258315116UL, 2521897172UL, 3900282042UL, 4186973154UL, 3146532165UL,
    2299685029UL, 3889120948UL, 1293301857UL, 187455105UL,  3395849230UL,
    913321567UL,  3093513909UL, 1440944571UL, 1923481911UL, 338680924UL,
    1204882963UL, 2739724491UL, 2886241328UL, 2408907774UL, 1299817192UL,
    2474012871UL, 45400213UL,   553186784UL,  134558656UL,  2180943666UL,
    2870807589UL, 76511085UL,   3053566760UL, 2516601415UL, 4172865902UL,
    1751297915UL, 1251975234UL, 2964780642UL, 1412975316UL, 2739978478UL,
    2171013719UL, 637935041UL,  975972384UL,  3044407449UL, 3111425639UL,
    1938684970UL, 2860857400UL, 13419586UL,   2772079268UL, 3484375614UL,
    3184054178UL, 159924837UL,  1386213021UL, 2765617231UL, 2523689118UL,
    1283505218UL, 3510789588UL, 4125878259UL, 2990287597UL, 2152014833UL,
    3084155970UL, 2815101609UL, 1932985704UL, 114887365UL,  1712687646UL,
    2550515629UL, 3299051916UL, 2022747614UL, 2143630992UL, 2244188960UL,
    3309469192UL, 3234358520UL, 800720365UL,  3278176634UL, 554357439UL,
    2415629802UL, 1620877315UL, 2389462898UL, 2229691332UL, 1007748450UL,
    1966873768UL, 2264971043UL, 1214524156UL, 346854700UL,  3471905342UL,
    3984889660UL, 4034246840UL, 216712649UL,  4027196762UL, 3754772604UL,
    2121785562UL, 2347070732UL, 7457687UL,    1443375102UL, 683948143UL,
    2940226032UL, 3211475670UL, 2836507357UL, 774899409UL,  1588968308UL,
    780438009UL,  3278878781UL, 2217181540UL, 2184194887UL, 1642129086UL,
    69346830UL,   297114710UL,  3841068188UL, 2631265450UL, 4167492314UL,
    2613519651UL, 1388582503UL, 2171556668UL, 1201873758UL, 2698772382UL,
    207791958UL,  3936134563UL, 3725025702UL, 3306317801UL, 1055730422UL,
    4069230694UL, 1767821343UL, 4252407395UL, 2422583118UL, 3158834399UL,
    3754582617UL, 1112422556UL, 376187931UL,  3137549150UL, 712221089UL,
    3300799453UL, 3868250200UL, 1165257666UL, 2494837767UL, 131304831UL,
    1619349427UL, 1958236644UL, 3678218946UL, 3651007751UL, 2261987899UL,
    1567368524UL, 2193599522UL, 3034394674UL, 2994602555UL, 3072727647UL,
    889094521UL,  1089692095UL, 1822324824UL, 3876999182UL, 1703361286UL,
    902229515UL,  4213728487UL, 3838170364UL, 672727494UL,  2240733828UL,
    3858539469UL, 1149254245UL, 4166055926UL, 4193525313UL, 1709921593UL,
    2278290377UL, 3190784116UL, 2919588882UL, 1012709717UL, 3640562031UL,
    2931984863UL, 3515665246UL, 250577343UL,  1147230194UL, 1183856202UL,
    3734511989UL, 3243867808UL, 3499383067UL, 2985115159UL, 2036821626UL,
    3298159553UL, 2726542838UL, 1686910320UL, 1778823772UL, 965412224UL,
    233509772UL,  3843098861UL, 1312622954UL, 500855830UL,  2950562091UL,
    1915683607UL, 3405781138UL, 596073719UL,  2195150546UL, 3381728478UL,
    546426436UL,  3527890868UL, 2324975353UL, 2241074266UL, 3992514859UL,
    2576108287UL, 4077653225UL, 2632319392UL, 3127212632UL, 917000669UL,
    2498161805UL, 3980835128UL, 2259526768UL, 1083920509UL, 1187452089UL,
    97018536UL,   3056075838UL, 2059706760UL, 2373335692UL, 182196406UL,
    2136713111UL, 1762080153UL, 1572125803UL, 1145919955UL, 1023966754UL,
    3921694345UL, 1632005969UL, 1418372326UL, 354407429UL,  2438288265UL,
    1620072033UL, 1586320921UL, 1044153697UL, 969324572UL,  613487980UL,
    4230993062UL, 397726764UL,  2194259193UL, 735511759UL,  2066049260UL,
    88093248UL,   1562536153UL, 2114157419UL, 3630951546UL, 589238503UL,
    3120654384UL, 2521793793UL, 2746692127UL, 2557723425UL, 889897693UL,
    2778878177UL, 643269509UL,  3342389831UL, 19218890UL,   3442706236UL,
    3314581273UL, 3503147052UL, 1546343434UL, 1448529060UL, 529038801UL,
    2748942264UL, 2213019208UL, 111314040UL,  2488697563UL, 1180642808UL,
    2605272289UL, 4207476668UL, 1502558669UL, 2972370981UL, 4204339995UL,
    1046225278UL, 992840610UL,  3847290298UL, 2387673094UL, 2221565747UL,
    1045901716UL, 3997739302UL, 1556952765UL, 1103336648UL, 279418400UL,
    2711316466UL, 2336215718UL, 2317900806UL, 974624729UL,  909575434UL,
    1675610631UL, 1922393214UL, 2054896570UL, 3197007361UL, 3932554569UL,
    1008619802UL, 3349254938UL, 113511461UL,  932630384UL,  2098759268UL,
    3436837432UL, 3119972401UL, 1612590197UL, 2281609013UL, 4174211248UL,
    4016332246UL, 2097525539UL, 1398632760UL, 1543697535UL, 2419227174UL,
    1676465074UL, 2882923045UL, 23216933UL,   808195649UL,  3690720147UL,
    484419260UL,  2254772642UL, 2975434733UL, 288528113UL,  204598404UL,
    589968818UL,  3021152400UL, 2463155141UL, 1397846755UL, 157285579UL,
    4230258857UL, 2469135246UL, 625357422UL,  3435224647UL, 465239124UL,
    1022535736UL, 2823317040UL, 274194469UL,  2214966446UL, 3661001613UL,
    518802547UL,  2293436304UL, 1335881988UL, 2247010176UL, 1856732584UL,
    1088028094UL, 1877563709UL, 1015352636UL, 1700817932UL, 2960695857UL,
    1882229300UL, 1666906557UL, 1838841022UL, 3983797810UL, 1667630361UL,
    385998221UL,  241341791UL,  403550441UL,  2629200403UL, 3552759102UL,
    2029750442UL, 2247999048UL, 2726665298UL, 2507798776UL, 2419064129UL,
    1266444923UL, 526255242UL,  2384866697UL, 1886200981UL, 3954956408UL,
    2171436866UL, 2295200753UL, 1047315850UL, 1967809707UL, 2860382973UL,
    3918334466UL, 3057439479UL, 952682588UL,  1925559679UL, 3112119050UL,
    3833190964UL, 1430139895UL, 2089165610UL, 3009202424UL, 3989186157UL,
    3395807230UL, 347600520UL,  120428923UL,  3017004655UL, 1384933954UL,
    303039929UL,  234010146UL,  2278760249UL, 315514836UL,  3987659575UL,
    1239335668UL, 2387869477UL, 3885908826UL, 1983922602UL, 698609264UL,
    3009002846UL, 1520611399UL, 809159940UL,  3089771783UL, 374838722UL,
    2789914419UL, 2500831937UL, 3751970335UL, 4279852547UL, 2362894437UL,
    1588814060UL, 1671213155UL, 434218829UL,  2126587176UL, 2002526422UL,
    2756464095UL, 141700479UL,  2965974322UL, 2211530172UL, 992085992UL,
    1943691492UL, 2705131817UL, 2519208889UL, 1938768395UL, 3949294294UL,
    354046666UL,  2158272751UL, 602858583UL,  0UL};

// 32-bits function
// return the i-th coefficient of the polynomial pf
unsigned long get_coef(unsigned long *pf, unsigned int deg) {
  if ((pf[deg >> 5] & (LSB << (deg & 0x1ful))) != 0)
    return (1);
  else
    return (0);
}

void copy_state(mt19937_state *target_state, mt19937_state *state) {
  int i;

  for (i = 0; i < N; i++)
    target_state->key[i] = state->key[i];

  target_state->pos = state->pos;
}

// next state generating function
void gen_next(mt19937_state *state) {
  int num;
  unsigned long y;
  static unsigned long mag02[2] = {0x0ul, MATRIX_A};

  num = state->pos;
  if (num < N - M) {
    y = (state->key[num] & UPPER_MASK) | (state->key[num + 1] & LOWER_MASK);
    state->key[num] = state->key[num + M] ^ (y >> 1) ^ mag02[y % 2];
    state->pos++;
  } else if (num < N - 1) {
    y = (state->key[num] & UPPER_MASK) | (state->key[num + 1] & LOWER_MASK);
    state->key[num] = state->key[num + (M - N)] ^ (y >> 1) ^ mag02[y % 2];
    state->pos++;
  } else if (num == N - 1) {
    y = (state->key[N - 1] & UPPER_MASK) | (state->key[0] & LOWER_MASK);
    state->key[N - 1] = state->key[M - 1] ^ (y >> 1) ^ mag02[y % 2];
    state->pos = 0;
  }
}

void add_state(mt19937_state *state1, mt19937_state *state2) {
  int i, pt1 = state1->pos, pt2 = state2->pos;

  if (pt2 - pt1 >= 0) {
    for (i = 0; i < N - pt2; i++)
      state1->key[i + pt1] ^= state2->key[i + pt2];
    for (; i < N - pt1; i++)
      state1->key[i + pt1] ^= state2->key[i + (pt2 - N)];
    for (; i < N; i++)
      state1->key[i + (pt1 - N)] ^= state2->key[i + (pt2 - N)];
  } else {
    for (i = 0; i < N - pt1; i++)
      state1->key[i + pt1] ^= state2->key[i + pt2];
    for (; i < N - pt2; i++)
      state1->key[i + (pt1 - N)] ^= state2->key[i + pt2];
    for (; i < N; i++)
      state1->key[i + (pt1 - N)] ^= state2->key[i + (pt2 - N)];
  }
}

// compute pf(ss) using standard Horner method
void horner1(unsigned long *pf, mt19937_state *state) {
  int i = MEXP - 1;
  mt19937_state *temp;

  temp = (mt19937_state *)calloc(1, sizeof(mt19937_state));

  while (get_coef(pf, i) == 0)
    i--;

  if (i > 0) {
    copy_state(temp, state);
    gen_next(temp);
    i--;
    for (; i > 0; i--) {
      if (get_coef(pf, i) != 0)
        add_state(temp, state);
      else
        ;
      gen_next(temp);
    }
    if (get_coef(pf, 0) != 0)
      add_state(temp, state);
    else
      ;
  } else if (i == 0)
    copy_state(temp, state);
  else
    ;

  copy_state(state, temp);
  free(temp);
}

void mt19937_jump_state(mt19937_state *state) {
  unsigned long *pf;
  int i;

  pf = (unsigned long *)calloc(P_SIZE, sizeof(unsigned long));
  for (i = 0; i<P_SIZE; i++) {
    pf[i] = poly_coef[i];
  }

  if (state->pos >= N) {
    state->pos = 0;
  }

  horner1(pf, state);

  free(pf);
}

#pragma once
#include <math.h>
#include <stdint.h>

#if defined(_WIN32) && !defined (__MINGW32__)
#define inline __forceinline
#endif

#define RK_STATE_LEN 624

#define N 624
#define M 397
#define MATRIX_A 0x9908b0dfUL
#define UPPER_MASK 0x80000000UL
#define LOWER_MASK 0x7fffffffUL

typedef struct s_mt19937_state {
  uint32_t key[RK_STATE_LEN];
  int pos;
} mt19937_state;

extern void mt19937_seed(mt19937_state *state, uint32_t seed);

extern void mt19937_gen(mt19937_state *state);

// Slightly optimized reference implementation of the Mersenne Twister
static inline uint32_t mt19937_next(mt19937_state *state) {
  uint32_t y;

  if (state->pos == RK_STATE_LEN) {
    // Move to function to help inlining
    mt19937_gen(state);
  }
  y = state->key[state->pos++];

  // Tempering
  y ^= (y >> 11);
  y ^= (y << 7) & 0x9d2c5680UL;
  y ^= (y << 15) & 0xefc60000UL;
  y ^= (y >> 18);

  return y;
}

extern void mt19937_init_by_array(mt19937_state *state, uint32_t *init_key,
                                  int key_length);

static inline uint64_t mt19937_next64(mt19937_state *state) {
  return (uint64_t)mt19937_next(state) << 32 | mt19937_next(state);
}

static inline uint32_t mt19937_next32(mt19937_state *state) {
  return mt19937_next(state);
}

static inline double mt19937_next_double(mt19937_state *state) {
  int32_t a = mt19937_next(state) >> 5, b = mt19937_next(state) >> 6;
  return (a * 67108864.0 + b) / 9007199254740992.0;
}

void mt19937_jump(mt19937_state *state);

#include "mt19937.h"
#include "mt19937-jump.h"

void mt19937_seed(mt19937_state *state, uint32_t seed) {
  int pos;
  seed &= 0xffffffffUL;

  // Knuth's PRNG as used in the Mersenne Twister reference implementation
  for (pos = 0; pos < RK_STATE_LEN; pos++) {
    state->key[pos] = seed;
    seed = (1812433253UL * (seed ^ (seed >> 30)) + pos + 1) & 0xffffffffUL;
  }
  state->pos = RK_STATE_LEN;
}

// initializes mt[RK_STATE_LEN] with a seed
static void init_genrand(mt19937_state *state, uint32_t s) {
  int mti;
  uint32_t *mt = state->key;

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
  state->pos = mti;
  return;
}

// initialize by an array with array-length
// init_key is the array for initializing keys
// key_length is its length
void mt19937_init_by_array(mt19937_state *state, uint32_t *init_key,
                           int key_length) {
  // was signed in the original code. RDH 12/16/2002
  int i = 1;
  int j = 0;
  uint32_t *mt = state->key;
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
    mt[i] = (mt[i] ^ ((mt[i - 1] ^ (mt[i - 1] >> 30)) * 1566083941UL)) -
            i;             // non linear
    mt[i] &= 0xffffffffUL; // for WORDSIZE > 32 machines
    i++;
    if (i >= RK_STATE_LEN) {
      mt[0] = mt[RK_STATE_LEN - 1];
      i = 1;
    }
  }

  mt[0] = 0x80000000UL; // MSB is 1; assuring non-zero initial array
}

void mt19937_gen(mt19937_state *state) {
  uint32_t y;
  int i;

  for (i = 0; i < N - M; i++) {
    y = (state->key[i] & UPPER_MASK) | (state->key[i + 1] & LOWER_MASK);
    state->key[i] = state->key[i + M] ^ (y >> 1) ^ (-(y & 1) & MATRIX_A);
  }
  for (; i < N - 1; i++) {
    y = (state->key[i] & UPPER_MASK) | (state->key[i + 1] & LOWER_MASK);
    state->key[i] = state->key[i + (M - N)] ^ (y >> 1) ^ (-(y & 1) & MATRIX_A);
  }
  y = (state->key[N - 1] & UPPER_MASK) | (state->key[0] & LOWER_MASK);
  state->key[N - 1] = state->key[M - 1] ^ (y >> 1) ^ (-(y & 1) & MATRIX_A);

  state->pos = 0;
}

extern inline uint64_t mt19937_next64(mt19937_state *state);

extern inline uint32_t mt19937_next32(mt19937_state *state);

extern inline double mt19937_next_double(mt19937_state *state);

void mt19937_jump(mt19937_state *state) { mt19937_jump_state(state); }
*/
