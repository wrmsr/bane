package arm64

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/wrmsr/bane/pkg/util/check"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

// Ext. register name -> int. name.
var map_archdef = map[string]string{"xzr": "@x31", "wzr": "@w31", "lr": "x30"}

// Int. register name -> ext. name.
var map_reg_rev = map[string]string{"@x31": "xzr", "@w31": "wzr", "x30": "lr"}

var map_shift = map[string]int{"lsl": 0, "lsr": 1, "asr": 2}

var map_extend = map[string]int{
	"uxtb": 0, "uxth": 1, "uxtw": 2, "uxtx": 3,
	"sxtb": 4, "sxth": 5, "sxtw": 6, "sxtx": 7,
}

var map_cond = map[string]int{
	"eq": 0,
	"ne": 1,
	"cs": 2,
	"cc": 3,
	"mi": 4,
	"pl": 5,
	"vs": 6,
	"vc": 7,
	"hi": 8,
	"ls": 9,
	"ge": 10,
	"lt": 11,
	"gt": 12,
	"le": 13,
	"al": 14,
	"hs": 2,
	"lo": 3,
}

// Template strings for ARM instructions.
var map_op = map[string]any{
	// Basic data processing instructions.
	"add_3":  "0b000000DNMg|11000000pDpNIg|8b206000pDpNMx",
	"add_4":  "0b000000DNMSg|0b200000DNMXg|8b200000pDpNMXx|8b200000pDpNxMwX",
	"adds_3": "2b000000DNMg|31000000DpNIg|ab206000DpNMx",
	"adds_4": "2b000000DNMSg|2b200000DNMXg|ab200000DpNMXx|ab200000DpNxMwX",
	"cmn_2":  "2b00001fNMg|3100001fpNIg|ab20601fpNMx",
	"cmn_3":  "2b00001fNMSg|2b20001fNMXg|ab20001fpNMXx|ab20001fpNxMwX",

	"sub_3":  "4b000000DNMg|51000000pDpNIg|cb206000pDpNMx",
	"sub_4":  "4b000000DNMSg|4b200000DNMXg|cb200000pDpNMXx|cb200000pDpNxMwX",
	"subs_3": "6b000000DNMg|71000000DpNIg|eb206000DpNMx",
	"subs_4": "6b000000DNMSg|6b200000DNMXg|eb200000DpNMXx|eb200000DpNxMwX",
	"cmp_2":  "6b00001fNMg|7100001fpNIg|eb20601fpNMx",
	"cmp_3":  "6b00001fNMSg|6b20001fNMXg|eb20001fpNMXx|eb20001fpNxMwX",

	"neg_2":  "4b0003e0DMg",
	"neg_3":  "4b0003e0DMSg",
	"negs_2": "6b0003e0DMg",
	"negs_3": "6b0003e0DMSg",

	"adc_3":  "1a000000DNMg",
	"adcs_3": "3a000000DNMg",
	"sbc_3":  "5a000000DNMg",
	"sbcs_3": "7a000000DNMg",
	"ngc_2":  "5a0003e0DMg",
	"ngcs_2": "7a0003e0DMg",

	"and_3":  "0a000000DNMg|12000000pDNig",
	"and_4":  "0a000000DNMSg",
	"orr_3":  "2a000000DNMg|32000000pDNig",
	"orr_4":  "2a000000DNMSg",
	"eor_3":  "4a000000DNMg|52000000pDNig",
	"eor_4":  "4a000000DNMSg",
	"ands_3": "6a000000DNMg|72000000DNig",
	"ands_4": "6a000000DNMSg",
	"tst_2":  "6a00001fNMg|7200001fNig",
	"tst_3":  "6a00001fNMSg",

	"bic_3":  "0a200000DNMg",
	"bic_4":  "0a200000DNMSg",
	"orn_3":  "2a200000DNMg",
	"orn_4":  "2a200000DNMSg",
	"eon_3":  "4a200000DNMg",
	"eon_4":  "4a200000DNMSg",
	"bics_3": "6a200000DNMg",
	"bics_4": "6a200000DNMSg",

	"movn_2": "12800000DWg",
	"movn_3": "12800000DWRg",
	"movz_2": "52800000DWg",
	"movz_3": "52800000DWRg",
	"movk_2": "72800000DWg",
	"movk_3": "72800000DWRg",

	// TODO: this doesn't cover all valid immediates for mov reg, #imm.
	"mov_2": "2a0003e0DMg|52800000DW|320003e0pDig|11000000pDpNg",
	"mov_3": "2a0003e0DMSg",
	"mvn_2": "2a2003e0DMg",
	"mvn_3": "2a2003e0DMSg",

	"adr_2":  "10000000DBx",
	"adrp_2": "90000000DBx",

	"csel_4":  "1a800000DNMCg",
	"csinc_4": "1a800400DNMCg",
	"csinv_4": "5a800000DNMCg",
	"csneg_4": "5a800400DNMCg",
	"cset_2":  "1a9f07e0Dcg",
	"csetm_2": "5a9f03e0Dcg",
	"cinc_3":  "1a800400DNmcg",
	"cinv_3":  "5a800000DNmcg",
	"cneg_3":  "5a800400DNmcg",

	"ccmn_4": "3a400000NMVCg|3a400800N5VCg",
	"ccmp_4": "7a400000NMVCg|7a400800N5VCg",

	"madd_4": "1b000000DNMAg",
	"msub_4": "1b008000DNMAg",
	"mul_3":  "1b007c00DNMg",
	"mneg_3": "1b00fc00DNMg",

	"smaddl_4": "9b200000DxNMwAx",
	"smsubl_4": "9b208000DxNMwAx",
	"smull_3":  "9b207c00DxNMw",
	"smnegl_3": "9b20fc00DxNMw",
	"smulh_3":  "9b407c00DNMx",
	"umaddl_4": "9ba00000DxNMwAx",
	"umsubl_4": "9ba08000DxNMwAx",
	"umull_3":  "9ba07c00DxNMw",
	"umnegl_3": "9ba0fc00DxNMw",
	"umulh_3":  "9bc07c00DNMx",

	"udiv_3": "1ac00800DNMg",
	"sdiv_3": "1ac00c00DNMg",

	// Bit operations.
	"sbfm_4": "13000000DN12w|93400000DN12x",
	"bfm_4":  "33000000DN12w|b3400000DN12x",
	"ubfm_4": "53000000DN12w|d3400000DN12x",
	"extr_4": "13800000DNM2w|93c00000DNM2x",

	"sxtb_2": "13001c00DNw|93401c00DNx",
	"sxth_2": "13003c00DNw|93403c00DNx",
	"sxtw_2": "93407c00DxNw",
	"uxtb_2": "53001c00DNw",
	"uxth_2": "53003c00DNw",

	// sbfx_4  = op_alias("sbfm_4", alias_bfx),
	// bfxil_4 = op_alias("bfm_4", alias_bfx),
	// ubfx_4  = op_alias("ubfm_4", alias_bfx),
	// sbfiz_4 = op_alias("sbfm_4", alias_bfiz),
	// bfi_4   = op_alias("bfm_4", alias_bfiz),
	// ubfiz_4 = op_alias("ubfm_4", alias_bfiz),

	// lsl_3  = function(params, nparams)
	//   if params and params[3]:byte() == 35 then
	//     return alias_lslimm(params, nparams)
	//   else
	//     return op_template(params, "1ac02000DNMg", nparams)
	//   end
	// end,
	"lsr_3": "1ac02400DNMg|53007c00DN1w|d340fc00DN1x",
	"asr_3": "1ac02800DNMg|13007c00DN1w|9340fc00DN1x",
	"ror_3": "1ac02c00DNMg|13800000DNm2w|93c00000DNm2x",

	"clz_2":   "5ac01000DNg",
	"cls_2":   "5ac01400DNg",
	"rbit_2":  "5ac00000DNg",
	"rev_2":   "5ac00800DNw|dac00c00DNx",
	"rev16_2": "5ac00400DNg",
	"rev32_2": "dac00800DNx",

	// Loads and stores.
	"strb_*":  "38000000DwL",
	"ldrb_*":  "38400000DwL",
	"ldrsb_*": "38c00000DwL|38800000DxL",
	"strh_*":  "78000000DwL",
	"ldrh_*":  "78400000DwL",
	"ldrsh_*": "78c00000DwL|78800000DxL",
	"str_*":   "b8000000DwL|f8000000DxL|bc000000DsL|fc000000DdL",
	"ldr_*":   "18000000DwB|58000000DxB|1c000000DsB|5c000000DdB|b8400000DwL|f8400000DxL|bc400000DsL|fc400000DdL",
	"ldrsw_*": "98000000DxB|b8800000DxL",
	// NOTE: ldur etc. are handled by ldr et al.

	"stp_*":   "28000000DAwP|a8000000DAxP|2c000000DAsP|6c000000DAdP",
	"ldp_*":   "28400000DAwP|a8400000DAxP|2c400000DAsP|6c400000DAdP",
	"ldpsw_*": "68400000DAxP",

	// Branches.
	"b_1":   "14000000B",
	"bl_1":  "94000000B",
	"blr_1": "d63f0000Nx",
	"br_1":  "d61f0000Nx",
	"ret_0": "d65f03c0",
	"ret_1": "d65f0000Nx",
	// b.cond is added below.
	"cbz_2":  "34000000DBg",
	"cbnz_2": "35000000DBg",
	"tbz_3":  "36000000DTBw|36000000DTBx",
	"tbnz_3": "37000000DTBw|37000000DTBx",

	// Miscellaneous instructions.
	// TODO: hlt, hvc, smc, svc, eret, dcps[123], drps, mrs, msr
	// TODO: sys, sysl, ic, dc, at, tlbi
	// TODO: hint, yield, wfe, wfi, sev, sevl
	// TODO: clrex, dsb, dmb, isb
	"nop_0": "d503201f",
	"brk_0": "d4200000",
	"brk_1": "d4200000W",

	// Floating point instructions.
	"fmov_2":  "1e204000DNf|1e260000DwNs|1e270000DsNw|9e660000DxNd|9e670000DdNx|1e201000DFf",
	"fabs_2":  "1e20c000DNf",
	"fneg_2":  "1e214000DNf",
	"fsqrt_2": "1e21c000DNf",

	"fcvt_2": "1e22c000DdNs|1e624000DsNd",

	// TODO: half-precision and fixed-point conversions.
	"fcvtas_2": "1e240000DwNs|9e240000DxNs|1e640000DwNd|9e640000DxNd",
	"fcvtau_2": "1e250000DwNs|9e250000DxNs|1e650000DwNd|9e650000DxNd",
	"fcvtms_2": "1e300000DwNs|9e300000DxNs|1e700000DwNd|9e700000DxNd",
	"fcvtmu_2": "1e310000DwNs|9e310000DxNs|1e710000DwNd|9e710000DxNd",
	"fcvtns_2": "1e200000DwNs|9e200000DxNs|1e600000DwNd|9e600000DxNd",
	"fcvtnu_2": "1e210000DwNs|9e210000DxNs|1e610000DwNd|9e610000DxNd",
	"fcvtps_2": "1e280000DwNs|9e280000DxNs|1e680000DwNd|9e680000DxNd",
	"fcvtpu_2": "1e290000DwNs|9e290000DxNs|1e690000DwNd|9e690000DxNd",
	"fcvtzs_2": "1e380000DwNs|9e380000DxNs|1e780000DwNd|9e780000DxNd",
	"fcvtzu_2": "1e390000DwNs|9e390000DxNs|1e790000DwNd|9e790000DxNd",

	"scvtf_2": "1e220000DsNw|9e220000DsNx|1e620000DdNw|9e620000DdNx",
	"ucvtf_2": "1e230000DsNw|9e230000DsNx|1e630000DdNw|9e630000DdNx",

	"frintn_2": "1e244000DNf",
	"frintp_2": "1e24c000DNf",
	"frintm_2": "1e254000DNf",
	"frintz_2": "1e25c000DNf",
	"frinta_2": "1e264000DNf",
	"frintx_2": "1e274000DNf",
	"frinti_2": "1e27c000DNf",

	"fadd_3":  "1e202800DNMf",
	"fsub_3":  "1e203800DNMf",
	"fmul_3":  "1e200800DNMf",
	"fnmul_3": "1e208800DNMf",
	"fdiv_3":  "1e201800DNMf",

	"fmadd_4":  "1f000000DNMAf",
	"fmsub_4":  "1f008000DNMAf",
	"fnmadd_4": "1f200000DNMAf",
	"fnmsub_4": "1f208000DNMAf",

	"fmax_3":   "1e204800DNMf",
	"fmaxnm_3": "1e206800DNMf",
	"fmin_3":   "1e205800DNMf",
	"fminnm_3": "1e207800DNMf",

	"fcmp_2":  "1e202000NMf|1e202008NZf",
	"fcmpe_2": "1e202010NMf|1e202018NZf",

	"fccmp_4":  "1e200400NMVCf",
	"fccmpe_4": "1e200410NMVCf",

	"fcsel_4": "1e200c00DNMCf",

	// TODO: crc32*, aes*, sha*, pmull
	// TODO: SIMD instructions.
}

func tohex(i int) string {
	return strconv.FormatInt(int64(i), 16)
}

func init() {
	for cond, c := range map_cond {
		map_op["b"+cond+"_1"] = tohex(0x54000000+c) + "B"
	}
}

var parse_reg_type string

type user_type struct {
	ctype    string
	ctypefmt string
	reg      string
}

var map_type = make(map[string]*user_type)

// Add action to list with optional arg. Advance buffer pos, too.
func waction(action, val, a, num any) {
	// w := assert(map_action[action], "bad action name `" .. action .. "'")
	// wputxw(w * 0x10000 + (val or 0))
	// if a {
	//     actargs[#actargs + 1] = a
	// }
	// if a || num {
	//     secpos = secpos + (num or 1)
	// }
}

func parse_reg(expr string, shift int, no_vreg bool) (int, any) {
	if expr == "" {
		panic("expected register name")
	}

	var tname, ovreg string
	m := regexp.MustCompile(`^([%w_]+):(@?%l%d+)$`).FindStringSubmatch(expr)
	if m != nil {
		tname, ovreg = m[1], m[2]
	} else {
		m = regexp.MustCompile(`^([%w_]+):(R[xwqdshb]%b())$`).FindStringSubmatch(expr)
		tname, ovreg = m[1], m[2]
	}

	tp := map_type[bt.Coalesce(tname, expr)]
	if tp != nil {
		reg := bt.Coalesce(ovreg, tp.reg)
		if reg == "" {
			panic(fmt.Errorf("type `%s` needs a register override", bt.Coalesce(tname, expr)))
		}
		expr = reg
	}

	m = regexp.MustCompile(`^(@?)([xwqdshb])([123]?[0-9])$`).FindStringSubmatch(expr)
	ok31, rt, r := m[1], m[2], m[3]
	if r != "" {
		r := check.Must1(strconv.Atoi(r))
		if r <= 30 || (r == 31 && ok31 != "" || (rt != "w" && rt != "x")) {
			if parse_reg_type == "" {
				parse_reg_type = rt
			} else if parse_reg_type != rt {
				panic("register size mismatch")
			}
			return r << shift, tp
		}
	}

	m = regexp.MustCompile(`^R([xwqdshb])(%b())$`).FindStringSubmatch(expr)
	vrt, vreg := m[1], m[2]
	if vreg != "" {
		if parse_reg_type == "" {
			parse_reg_type = vrt
		} else if parse_reg_type != vrt {
			panic("register size mismatch")
		}
		if !no_vreg {
			waction("VREG", shift, vreg, nil)
		}
		return 0, nil
	}

	panic(fmt.Errorf("bad register name `%v`", expr))
}

func parse_reg_base(expr string) (int, any) {
	if expr == "sp" {
		return 0x3e0, nil
	}
	base, tp := parse_reg(expr, 5, false)
	if parse_reg_type != "x" {
		panic("bad register type")
	}
	parse_reg_type = ""
	return base, tp
}

func parse_number(s string) bt.Optional[int] {
	if strings.HasPrefix(s, "0x") || strings.HasSuffix(s, "h") {
		s = strings.TrimPrefix(s, "0x")
		s = strings.TrimSuffix(s, "h")
		if n, err := strconv.ParseInt(s, 16, 64); err == nil {
			return bt.Just(int(n))
		}
	} else {
		if n, err := strconv.Atoi(s); err == nil {
			return bt.Just(n)
		}
	}
	return bt.None[int]()
}

func parse_imm(imm string, bits, shift, scale int, signed bool) any {
	m := regexp.MustCompile(`^#(.*)$`).FindStringSubmatch(imm)
	imm = m[1]
	if imm == "" {
		panic("expected immediate operand")
	}
	n := parse_number(imm)
	if n.Present() {
		n := n.Value()
		m := n >> scale
		if (m << scale) == n {
			if signed {
				s := m >> (bits - 1)
				if s == 0 {
					return m << shift
				} else if s == -1 {
					return (m + (1 << bits)) << shift
				}
			} else {
				if (m >> bits) == 0 {
					return m << shift
				}
			}
		}
		panic(fmt.Errorf("out of range immediate `%v`", imm))
	} else {
		waction("IMM", bt.Choose(signed, 32768, 0)+scale*1024+bits*32+shift, imm, nil)
		return 0
	}
}

func parse_imm12(imm string) int {
	m := regexp.MustCompile(`^#(.*)$`).FindStringSubmatch(imm)
	imm = m[1]
	if imm == "" {
		panic("expected immediate operand")
	}
	n := parse_number(imm)
	if n.Present() {
		n := n.Value()
		if (n >> 12) == 0 {
			return n << 10
		} else if (n & 0xff000fff) == 0 {
			return (n >> 2) + 0x00400000
		}
		panic(fmt.Errorf("out of range immediate `%v`", imm))
	} else {
		waction("IMM12", 0, imm, nil)
		return 0
	}
}

/*
func parse_imm13(imm string) int {
	m := regexp.MustCompile(`^#(.*)$`).FindStringSubmatch(imm)
	imm = m[1]
    if imm == "" {
        panic("expected immediate operand")
    }
    n := parse_number(imm)
    r64 := parse_reg_type == "x"
    if n && n % 1 == 0 && n >= 0 && n <= 0xffffffff {
        inv := false
        if (n & 1) == 1 {
            n = ^n
            inv = true
        }
        t := make(map[int]int)
        for i := 1; i <= 32; i++ {
            t[i] = n & 1
            n = n >> 1
        }
        local b = table.concat(t)
        b = b .. (r64 and (inv and "1" or "0"):rep(32) or b)
        local p0, p1, p0a, p1a = b:match("^(0+)(1+)(0*)(1*)")
        if p0 {
            w := p1a == "" && bt.Choose(r64, 64, 32) or #p1 + #p0a
            if band(w, w - 1) == 0 and b == b:sub(1, w):rep(64 / w) {
                s := ((-2 * w) & 0x3f) - 1
                if w == 64 {
                    s = s + 0x1000
                }
                if inv {
                    return ((w - #p1 - #p0) << 16) + ((s + w - #p1) << 10)
                } else {
                    return ((w - #p0) << 16) + ((s + #p1) << 10)
                }
            }
        }
        panic(fmt.Errorf("out of range immediate `%v`", imm))
    } else if r64 {
        waction("IMM13X", 0, fmt.Sprintf("(unsigned int)(%s)", imm), nil)
        //actargs[#actargs + 1] = fmt.Sprintf("(unsigned int)((unsigned long long)(%s)>>32)", imm)
        return 0
    } else {
        waction("IMM13W", 0, imm, nil)
        return 0
    }
}

func parse_imm6(imm string) int {
	m := regexp.MustCompile(`^#(.*)$`).FindStringSubmatch(imm)
	imm = m[1]
    if imm == "" {
        panic("expected immediate operand")
    }
    local n = parse_number(imm)
    if n {
        if n >= 0 and n <= 63 then
            return shl(band(n, 0x1f), 19) + (n >= 32 and 0x80000000 or 0)
        end
        panic(fmt.Errorf("out of range immediate `%v`", imm))
    } else {
        waction("IMM6", 0, imm)
        return 0
    }
}
*/
