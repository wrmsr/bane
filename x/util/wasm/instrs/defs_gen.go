// Code generated from defs.json. DO NOT EDIT.

package instrs

import wt "github.com/wrmsr/bane/x/util/wasm/types"

const (
	_ Instr = 0

	Unreachable                = 1
	Nop                        = 2
	Block                      = 3
	Loop                       = 4
	If                         = 5
	Else                       = 6
	End                        = 7
	Br                         = 8
	BrIf                       = 9
	BrTable                    = 10
	Return                     = 11
	Call                       = 12
	CallIndirect               = 13
	Drop                       = 14
	Select                     = 15
	SelectType                 = 16
	LocalGet                   = 17
	LocalSet                   = 18
	LocalTee                   = 19
	GlobalGet                  = 20
	GlobalSet                  = 21
	TableGet                   = 22
	TableSet                   = 23
	TableInit                  = 24
	ElemDrop                   = 25
	TableCopy                  = 26
	TableGrow                  = 27
	TableSize                  = 28
	TableFill                  = 29
	Load_I32                   = 30
	Load_I64                   = 31
	Load_F32                   = 32
	Load_F64                   = 33
	Load8S_I32                 = 34
	Load8U_I32                 = 35
	Load16S_I32                = 36
	Load16U_I32                = 37
	Load8S_I64                 = 38
	Load8U_I64                 = 39
	Load16S_I64                = 40
	Load16U_I64                = 41
	Load32S_I64                = 42
	Load32U_I64                = 43
	Store_I32                  = 44
	Store_I64                  = 45
	Store_F32                  = 46
	Store_F64                  = 47
	Store8_I32                 = 48
	Store16_I32                = 49
	Store8_I64                 = 50
	Store16_I64                = 51
	Store32_I64                = 52
	MemorySize                 = 53
	MemoryGrow                 = 54
	MemoryInit                 = 55
	DataDrop                   = 56
	MemoryCopy                 = 57
	MemoryFill                 = 58
	Const_I32                  = 59
	Const_I64                  = 60
	Const_F32                  = 61
	Const_F64                  = 62
	Eqz_I32                    = 63
	Eq_I32                     = 64
	Ne_I32                     = 65
	LtS_I32                    = 66
	LtU_I32                    = 67
	GtS_I32                    = 68
	GtU_I32                    = 69
	LeS_I32                    = 70
	LeU_I32                    = 71
	GeS_I32                    = 72
	GeU_I32                    = 73
	Eqz_I64                    = 74
	Eq_I64                     = 75
	Ne_I64                     = 76
	LtS_I64                    = 77
	LtU_I64                    = 78
	GtS_I64                    = 79
	GtU_I64                    = 80
	LeS_I64                    = 81
	LeU_I64                    = 82
	GeS_I64                    = 83
	GeU_I64                    = 84
	Eq_F32                     = 85
	Ne_F32                     = 86
	Lt_F32                     = 87
	Gt_F32                     = 88
	Le_F32                     = 89
	Ge_F32                     = 90
	Eq_F64                     = 91
	Ne_F64                     = 92
	Lt_F64                     = 93
	Gt_F64                     = 94
	Le_F64                     = 95
	Ge_F64                     = 96
	Clz_I32                    = 97
	Ctz_I32                    = 98
	Popcnt_I32                 = 99
	Add_I32                    = 100
	Sub_I32                    = 101
	Mul_I32                    = 102
	DivS_I32                   = 103
	DivU_I32                   = 104
	RemS_I32                   = 105
	RemU_I32                   = 106
	And_I32                    = 107
	Or_I32                     = 108
	Xor_I32                    = 109
	Shl_I32                    = 110
	ShrS_I32                   = 111
	ShrU_I32                   = 112
	RotL_I32                   = 113
	RotR_I32                   = 114
	Clz_I64                    = 115
	Ctz_I64                    = 116
	Popcnt_I64                 = 117
	Add_I64                    = 118
	Sub_I64                    = 119
	Mul_I64                    = 120
	DivS_I64                   = 121
	DivU_I64                   = 122
	RemS_I64                   = 123
	RemU_I64                   = 124
	And_I64                    = 125
	Or_I64                     = 126
	Xor_I64                    = 127
	Shl_I64                    = 128
	ShrS_I64                   = 129
	ShrU_I64                   = 130
	RotL_I64                   = 131
	RotR_I64                   = 132
	Abs_F32                    = 133
	Neg_F32                    = 134
	Ceil_F32                   = 135
	Floor_F32                  = 136
	Trunc_F32                  = 137
	Nearest_F32                = 138
	Sqrt_F32                   = 139
	Add_F32                    = 140
	Sub_F32                    = 141
	Mul_F32                    = 142
	Div_F32                    = 143
	Min_F32                    = 144
	Max_F32                    = 145
	Copysign_F32               = 146
	Abs_F64                    = 147
	Neg_F64                    = 148
	Ceil_F64                   = 149
	Floor_F64                  = 150
	Trunc_F64                  = 151
	Nearest_F64                = 152
	Sqrt_F64                   = 153
	Add_F64                    = 154
	Sub_F64                    = 155
	Mul_F64                    = 156
	Div_F64                    = 157
	Min_F64                    = 158
	Max_F64                    = 159
	Copysign_F64               = 160
	WrapI64_I32                = 161
	TruncF32S_I32              = 162
	TruncF32U_I32              = 163
	TruncF64S_I32              = 164
	TruncF64U_I32              = 165
	ExtendI32S_I64             = 166
	ExtendI32U_I64             = 167
	TruncF32S_I64              = 168
	TruncF32U_I64              = 169
	TruncF64S_I64              = 170
	TruncF64U_I64              = 171
	ConvertI32S_F32            = 172
	ConvertI32U_F32            = 173
	ConvertI64S_F32            = 174
	ConvertI64U_F32            = 175
	DemoteF64_F32              = 176
	ConvertI32S_F64            = 177
	ConvertI32U_F64            = 178
	ConvertI64S_F64            = 179
	ConvertI64U_F64            = 180
	PromoteF32_F64             = 181
	ReinterpretF32_I32         = 182
	ReinterpretF64_I64         = 183
	ReinterpretI32_F32         = 184
	ReinterpretI64_F64         = 185
	Extend8S_I32               = 186
	Extend16S_I32              = 187
	Extend8S_I64               = 188
	Extend16S_I64              = 189
	Extend32S_I64              = 190
	TruncSatF32S_I32           = 191
	TruncSatF32U_I32           = 192
	TruncSatF64S_I32           = 193
	TruncSatF64U_I32           = 194
	TruncSatF32S_I64           = 195
	TruncSatF32U_I64           = 196
	TruncSatF64S_I64           = 197
	TruncSatF64U_I64           = 198
	Load_V128                  = 199
	Load8x8S_V128              = 200
	Load8x8U_V128              = 201
	Load16x4S_V128             = 202
	Load16x4U_V128             = 203
	Load32x2S_V128             = 204
	Load32x2U_V128             = 205
	Load8Splat_V128            = 206
	Load16Splat_V128           = 207
	Load32Splat_V128           = 208
	Load64Splat_V128           = 209
	Load32Zero_V128            = 210
	Load64Zero_V128            = 211
	Store_V128                 = 212
	Load8Lane_V128             = 213
	Load16Lane_V128            = 214
	Load32Lane_V128            = 215
	Load64Lane_V128            = 216
	Store8Lane_V128            = 217
	Store16Lane_V128           = 218
	Store32Lane_V128           = 219
	Store64Lane_V128           = 220
	Const_V128                 = 221
	Shuffle_I8x16              = 222
	ExtractLaneS_I8x16         = 223
	ExtractLaneU_I8x16         = 224
	ReplaceLane_I8x16          = 225
	ExtractLaneS_I16x8         = 226
	ExtractLaneU_I16x8         = 227
	ReplaceLane_I16x8          = 228
	ExtractLane_I32x4          = 229
	ReplaceLane_I32x4          = 230
	ExtractLane_I64x2          = 231
	ReplaceLane_I64x2          = 232
	ExtractLane_F32x4          = 233
	ReplaceLane_F32x4          = 234
	ExtractLane_F64x2          = 235
	ReplaceLane_F64x2          = 236
	Swizzle_I8x16              = 237
	Splat_I8x16                = 238
	Splat_I16x8                = 239
	Splat_I32x4                = 240
	Splat_I64x2                = 241
	Splat_F32x4                = 242
	Splat_F64x2                = 243
	Eq_I8x16                   = 244
	Ne_I8x16                   = 245
	LtS_I8x16                  = 246
	LtU_I8x16                  = 247
	GtS_I8x16                  = 248
	GtU_I8x16                  = 249
	LeS_I8x16                  = 250
	LeU_I8x16                  = 251
	GeS_I8x16                  = 252
	GeU_I8x16                  = 253
	Eq_I16x8                   = 254
	Ne_I16x8                   = 255
	LtS_I16x8                  = 256
	LtU_I16x8                  = 257
	GtS_I16x8                  = 258
	GtU_I16x8                  = 259
	LeS_I16x8                  = 260
	LeU_I16x8                  = 261
	GeS_I16x8                  = 262
	GeU_I16x8                  = 263
	Eq_I32x4                   = 264
	Ne_I32x4                   = 265
	LtS_I32x4                  = 266
	LtU_I32x4                  = 267
	GtS_I32x4                  = 268
	GtU_I32x4                  = 269
	LeS_I32x4                  = 270
	LeU_I32x4                  = 271
	GeS_I32x4                  = 272
	GeU_I32x4                  = 273
	Eq_I64x2                   = 274
	Ne_I64x2                   = 275
	LtS_I64x2                  = 276
	GtS_I64x2                  = 277
	LeS_I64x2                  = 278
	GeS_I64x2                  = 279
	Eq_F32x4                   = 280
	Ne_F32x4                   = 281
	Lt_F32x4                   = 282
	Gt_F32x4                   = 283
	Le_F32x4                   = 284
	Ge_F32x4                   = 285
	Eq_F64x2                   = 286
	Ne_F64x2                   = 287
	Lt_F64x2                   = 288
	Gt_F64x2                   = 289
	Le_F64x2                   = 290
	Ge_F64x2                   = 291
	Not_V128                   = 292
	And_V128                   = 293
	Andnot_V128                = 294
	Or_V128                    = 295
	Xor_V128                   = 296
	Bitselect_V128             = 297
	AnyTrue_V128               = 298
	Abs_I8x16                  = 299
	Neg_I8x16                  = 300
	Popcnt_I8x16               = 301
	AllTrue_I8x16              = 302
	Bitmask_I8x16              = 303
	NarrowI16x8S_I8x16         = 304
	NarrowI16x8U_I8x16         = 305
	Shl_I8x16                  = 306
	ShrS_I8x16                 = 307
	ShrU_I8x16                 = 308
	Add_I8x16                  = 309
	AddSatS_I8x16              = 310
	AddSatU_I8x16              = 311
	Sub_I8x16                  = 312
	SubSatS_I8x16              = 313
	SubSatU_I8x16              = 314
	MinS_I8x16                 = 315
	MinU_I8x16                 = 316
	MaxS_I8x16                 = 317
	MaxU_I8x16                 = 318
	AvgrU_I8x16                = 319
	ExtaddPairwiseI8x16S_I16x8 = 320
	ExtaddPairwiseI8x16U_I16x8 = 321
	Abs_I16x8                  = 322
	Neg_I16x8                  = 323
	Q15mulrSatS_I16x8          = 324
	Alltrue_I16x8              = 325
	Bitmask_I16x8              = 326
	NarrowI32x4S_I16x8         = 327
	NarrowI32x4U_I16x8         = 328
	ExtendLowI8x16S_I16x8      = 329
	ExtendHighI8x16S_I16x8     = 330
	ExtendLowI8x16U_I16x8      = 331
	ExtendHighI8x16U_I16x8     = 332
	Shl_I16x8                  = 333
	ShrS_I16x8                 = 334
	ShrU_I16x8                 = 335
	Add_I16x8                  = 336
	AddSatS_I16x8              = 337
	AddSatU_I16x8              = 338
	Sub_I16x8                  = 339
	SubSatS_I16x8              = 340
	SubSatU_I16x8              = 341
	Mul_I16x8                  = 342
	MinS_I16x8                 = 343
	MinU_I16x8                 = 344
	MaxS_I16x8                 = 345
	MaxU_I16x8                 = 346
	AvgrU_I16x8                = 347
	ExtmulLowI8x16S_I16x8      = 348
	ExtmulHighI8x16S_I16x8     = 349
	ExtmulLowI8x16U_I16x8      = 350
	ExtmulHighI8x16U_I16x8     = 351
	ExtaddPairwiseI16x8S_I32x4 = 352
	ExtaddPairwiseI16x8U_I32x4 = 353
	Abs_I32x4                  = 354
	Neg_I32x4                  = 355
	Alltrue_I32x4              = 356
	Bitmask_I32x4              = 357
	ExtendLowI16x8S_I32x4      = 358
	ExtendHighI16x8S_I32x4     = 359
	ExtendLowI16x8U_I32x4      = 360
	ExtendHighI16x8U_I32x4     = 361
	Shl_I32x4                  = 362
	ShrS_I32x4                 = 363
	ShrU_I32x4                 = 364
	Add_I32x4                  = 365
	Sub_I32x4                  = 366
	Mul_I32x4                  = 367
	MinS_I32x4                 = 368
	MinU_I32x4                 = 369
	MaxS_I32x4                 = 370
	MaxU_I32x4                 = 371
	DotI16x8S_I32x4            = 372
	ExtmulLowI16x8S_I32x4      = 373
	ExtmulHighI16x8S_I32x4     = 374
	ExtmulLowI16x8U_I32x4      = 375
	ExtmulHighI16x8U_I32x4     = 376
	Abs_I64x2                  = 377
	Neg_I64x2                  = 378
	Alltrue_I64x2              = 379
	Bitmask_I64x2              = 380
	ExtendLowI32x4S_I64x2      = 381
	ExtendHighI32x4S_I64x2     = 382
	ExtendLowI32x4U_I64x2      = 383
	ExtendHighI32x4U_I64x2     = 384
	Shl_I64x2                  = 385
	ShrS_I64x2                 = 386
	ShrU_I64x2                 = 387
	Add_I64x2                  = 388
	Sub_I64x2                  = 389
	Mul_I64x2                  = 390
	ExtmulLowI32x4S_I64x2      = 391
	ExtmulHighI32x4S_I64x2     = 392
	ExtmulLowI32x4U_I64x2      = 393
	ExtmulHighI32x4U_I64x2     = 394
	Ceil_F32x4                 = 395
	Floor_F32x4                = 396
	Trunc_F32x4                = 397
	Nearest_F32x4              = 398
	Abs_F32x4                  = 399
	Neg_F32x4                  = 400
	Sqrt_F32x4                 = 401
	Add_F32x4                  = 402
	Sub_F32x4                  = 403
	Mul_F32x4                  = 404
	Div_F32x4                  = 405
	Min_F32x4                  = 406
	Max_F32x4                  = 407
	Pmin_F32x4                 = 408
	Pmax_F32x4                 = 409
	Ceil_F64x2                 = 410
	Floor_F64x2                = 411
	Trunc_F64x2                = 412
	Nearest_F64x2              = 413
	Abs_F64x2                  = 414
	Neg_F64x2                  = 415
	Sqrt_F64x2                 = 416
	Add_F64x2                  = 417
	Sub_F64x2                  = 418
	Mul_F64x2                  = 419
	Div_F64x2                  = 420
	Min_F64x2                  = 421
	Max_F64x2                  = 422
	Pmin_F64x2                 = 423
	Pmax_F64x2                 = 424
	TruncSatF32x4S_I32x4       = 425
	TruncSatF32x4U_I32x4       = 426
	ConvertI32x4S_F32x4        = 427
	ConvertI32x4U_F32x4        = 428
	TruncSatF64x2SZero_I32x4   = 429
	TruncSatF64x2UZero_I32x4   = 430
	ConvertLowI32x4S_F64x2     = 431
	ConvertLowI32x4U_F64x2     = 432
	DemoteF64x2Zero_F32x4      = 433
	PromoteLowF32x4_F64x2      = 434
)

var defs = []Def{
	{name: "invalid", name2: "Invalid"},
	{i: Unreachable, class: Control, name: "unreachable", name2: "Unreachable", op: uint8(0x0)},
	{i: Nop, class: Control, name: "nop", name2: "Nop", op: uint8(0x1)},
	{i: Block, class: Control, name: "block", name2: "Block", op: uint8(0x2)},
	{i: Loop, class: Control, name: "loop", name2: "Loop", op: uint8(0x3)},
	{i: If, class: Control, name: "if", name2: "If", op: uint8(0x4)},
	{i: Else, class: Control, name: "else", name2: "Else", op: uint8(0x5)},
	{i: End, class: Control, name: "end", name2: "End", op: uint8(0xB)},
	{i: Br, class: Control, name: "br", name2: "Br", op: uint8(0xC)},
	{i: BrIf, class: Control, name: "br_if", name2: "BrIf", op: uint8(0xD), a: wt.I32},
	{i: BrTable, class: Control, name: "br_table", name2: "BrTable", op: uint8(0xE), a: wt.I32},
	{i: Return, class: Control, name: "return", name2: "Return", op: uint8(0xF)},
	{i: Call, class: Control, name: "call", name2: "Call", op: uint8(0x10)},
	{i: CallIndirect, class: Control, name: "call_indirect", name2: "CallIndirect", op: uint8(0x11)},
	{i: Drop, class: Parametric, name: "drop", name2: "Drop", op: uint8(0x1A)},
	{i: Select, class: Parametric, name: "select", name2: "Select", op: uint8(0x1B), c: wt.I32},
	{i: SelectType, class: Parametric, name: "select_type", name2: "SelectType", op: uint8(0x1C), c: wt.I32},
	{i: LocalGet, class: Variable, name: "local_get", name2: "LocalGet", op: uint8(0x20)},
	{i: LocalSet, class: Variable, name: "local_set", name2: "LocalSet", op: uint8(0x21)},
	{i: LocalTee, class: Variable, name: "local_tee", name2: "LocalTee", op: uint8(0x22)},
	{i: GlobalGet, class: Variable, name: "global_get", name2: "GlobalGet", op: uint8(0x23)},
	{i: GlobalSet, class: Variable, name: "global_set", name2: "GlobalSet", op: uint8(0x24)},
	{i: TableGet, class: Table, name: "table_get", name2: "TableGet", op: uint8(0x25), a: wt.I32},
	{i: TableSet, class: Table, name: "table_set", name2: "TableSet", op: uint8(0x26), a: wt.I32},
	{i: TableInit, class: Table, name: "table_init", name2: "TableInit", opPfx: uint8(0xFC), op: uint8(0xC), a: wt.I32, b: wt.I32, c: wt.I32},
	{i: ElemDrop, class: Table, name: "elem_drop", name2: "ElemDrop", opPfx: uint8(0xFC), op: uint8(0xD)},
	{i: TableCopy, class: Table, name: "table_copy", name2: "TableCopy", opPfx: uint8(0xFC), op: uint8(0xE), a: wt.I32, b: wt.I32, c: wt.I32},
	{i: TableGrow, class: Table, name: "table_grow", name2: "TableGrow", opPfx: uint8(0xFC), op: uint8(0xF), b: wt.I32},
	{i: TableSize, class: Table, name: "table_size", name2: "TableSize", opPfx: uint8(0xFC), op: uint8(0x10)},
	{i: TableFill, class: Table, name: "table_fill", name2: "TableFill", opPfx: uint8(0xFC), op: uint8(0x11), a: wt.I32, c: wt.I32},
	{i: Load_I32, class: Memory, name: "load.i32", name2: "Load_I32", op: uint8(0x28), t: wt.I32, a: wt.I32, ma: Load, mz: 4},
	{i: Load_I64, class: Memory, name: "load.i64", name2: "Load_I64", op: uint8(0x29), t: wt.I64, a: wt.I32, ma: Load, mz: 8},
	{i: Load_F32, class: Memory, name: "load.f32", name2: "Load_F32", op: uint8(0x2A), t: wt.F32, a: wt.I32, ma: Load, mz: 4},
	{i: Load_F64, class: Memory, name: "load.f64", name2: "Load_F64", op: uint8(0x2B), t: wt.F64, a: wt.I32, ma: Load, mz: 8},
	{i: Load8S_I32, class: Memory, name: "load8_s.i32", name2: "Load8S_I32", op: uint8(0x2C), t: wt.I32, a: wt.I32, ma: Load, mz: 1},
	{i: Load8U_I32, class: Memory, name: "load8_u.i32", name2: "Load8U_I32", op: uint8(0x2D), t: wt.I32, a: wt.I32, ma: Load, mz: 1},
	{i: Load16S_I32, class: Memory, name: "load16_s.i32", name2: "Load16S_I32", op: uint8(0x2E), t: wt.I32, a: wt.I32, ma: Load, mz: 2},
	{i: Load16U_I32, class: Memory, name: "load16_u.i32", name2: "Load16U_I32", op: uint8(0x2F), t: wt.I32, a: wt.I32, ma: Load, mz: 2},
	{i: Load8S_I64, class: Memory, name: "load8_s.i64", name2: "Load8S_I64", op: uint8(0x30), t: wt.I64, a: wt.I32, ma: Load, mz: 1},
	{i: Load8U_I64, class: Memory, name: "load8_u.i64", name2: "Load8U_I64", op: uint8(0x31), t: wt.I64, a: wt.I32, ma: Load, mz: 1},
	{i: Load16S_I64, class: Memory, name: "load16_s.i64", name2: "Load16S_I64", op: uint8(0x32), t: wt.I64, a: wt.I32, ma: Load, mz: 2},
	{i: Load16U_I64, class: Memory, name: "load16_u.i64", name2: "Load16U_I64", op: uint8(0x33), t: wt.I64, a: wt.I32, ma: Load, mz: 2},
	{i: Load32S_I64, class: Memory, name: "load32_s.i64", name2: "Load32S_I64", op: uint8(0x34), t: wt.I64, a: wt.I32, ma: Load, mz: 4},
	{i: Load32U_I64, class: Memory, name: "load32_u.i64", name2: "Load32U_I64", op: uint8(0x35), t: wt.I64, a: wt.I32, ma: Load, mz: 4},
	{i: Store_I32, class: Memory, name: "store.i32", name2: "Store_I32", op: uint8(0x36), a: wt.I32, b: wt.I32, ma: Store, mz: 4},
	{i: Store_I64, class: Memory, name: "store.i64", name2: "Store_I64", op: uint8(0x37), a: wt.I32, b: wt.I64, ma: Store, mz: 8},
	{i: Store_F32, class: Memory, name: "store.f32", name2: "Store_F32", op: uint8(0x38), a: wt.I32, b: wt.F32, ma: Store, mz: 4},
	{i: Store_F64, class: Memory, name: "store.f64", name2: "Store_F64", op: uint8(0x39), a: wt.I32, b: wt.F64, ma: Store, mz: 8},
	{i: Store8_I32, class: Memory, name: "store8.i32", name2: "Store8_I32", op: uint8(0x3A), a: wt.I32, b: wt.I32, ma: Store, mz: 1},
	{i: Store16_I32, class: Memory, name: "store16.i32", name2: "Store16_I32", op: uint8(0x3B), a: wt.I32, b: wt.I32, ma: Store, mz: 2},
	{i: Store8_I64, class: Memory, name: "store8.i64", name2: "Store8_I64", op: uint8(0x3C), a: wt.I32, b: wt.I64, ma: Store, mz: 1},
	{i: Store16_I64, class: Memory, name: "store16.i64", name2: "Store16_I64", op: uint8(0x3D), a: wt.I32, b: wt.I64, ma: Store, mz: 2},
	{i: Store32_I64, class: Memory, name: "store32.i64", name2: "Store32_I64", op: uint8(0x3E), a: wt.I32, b: wt.I64, ma: Store, mz: 4},
	{i: MemorySize, class: Memory, name: "memory_size", name2: "MemorySize", op: uint8(0x3F), t: wt.I32},
	{i: MemoryGrow, class: Memory, name: "memory_grow", name2: "MemoryGrow", op: uint8(0x40), t: wt.I32, a: wt.I32},
	{i: MemoryInit, class: Memory, name: "memory_init", name2: "MemoryInit", opPfx: uint8(0xFC), op: uint8(0x8), a: wt.I32, b: wt.I32, c: wt.I32},
	{i: DataDrop, class: Memory, name: "data_drop", name2: "DataDrop", opPfx: uint8(0xFC), op: uint8(0x9)},
	{i: MemoryCopy, class: Memory, name: "memory_copy", name2: "MemoryCopy", opPfx: uint8(0xFC), op: uint8(0xA), a: wt.I32, b: wt.I32, c: wt.I32},
	{i: MemoryFill, class: Memory, name: "memory_fill", name2: "MemoryFill", opPfx: uint8(0xFC), op: uint8(0xB), a: wt.I32, b: wt.I32, c: wt.I32},
	{i: Const_I32, class: Numeric, name: "const.i32", name2: "Const_I32", op: uint8(0x41), t: wt.I32},
	{i: Const_I64, class: Numeric, name: "const.i64", name2: "Const_I64", op: uint8(0x42), t: wt.F32},
	{i: Const_F32, class: Numeric, name: "const.f32", name2: "Const_F32", op: uint8(0x43), t: wt.I64},
	{i: Const_F64, class: Numeric, name: "const.f64", name2: "Const_F64", op: uint8(0x44), t: wt.F64},
	{i: Eqz_I32, class: Numeric, name: "eqz.i32", name2: "Eqz_I32", op: uint8(0x45), t: wt.I32, a: wt.I32},
	{i: Eq_I32, class: Numeric, name: "eq.i32", name2: "Eq_I32", op: uint8(0x46), t: wt.I32, a: wt.I32, b: wt.I32},
	{i: Ne_I32, class: Numeric, name: "ne.i32", name2: "Ne_I32", op: uint8(0x47), t: wt.I32, a: wt.I32, b: wt.I32},
	{i: LtS_I32, class: Numeric, name: "lt_s.i32", name2: "LtS_I32", op: uint8(0x48), t: wt.I32, a: wt.I32, b: wt.I32},
	{i: LtU_I32, class: Numeric, name: "lt_u.i32", name2: "LtU_I32", op: uint8(0x49), t: wt.I32, a: wt.I32, b: wt.I32},
	{i: GtS_I32, class: Numeric, name: "gt_s.i32", name2: "GtS_I32", op: uint8(0x4A), t: wt.I32, a: wt.I32, b: wt.I32},
	{i: GtU_I32, class: Numeric, name: "gt_u.i32", name2: "GtU_I32", op: uint8(0x4B), t: wt.I32, a: wt.I32, b: wt.I32},
	{i: LeS_I32, class: Numeric, name: "le_s.i32", name2: "LeS_I32", op: uint8(0x4C), t: wt.I32, a: wt.I32, b: wt.I32},
	{i: LeU_I32, class: Numeric, name: "le_u.i32", name2: "LeU_I32", op: uint8(0x4D), t: wt.I32, a: wt.I32, b: wt.I32},
	{i: GeS_I32, class: Numeric, name: "ge_s.i32", name2: "GeS_I32", op: uint8(0x4E), t: wt.I32, a: wt.I32, b: wt.I32},
	{i: GeU_I32, class: Numeric, name: "ge_u.i32", name2: "GeU_I32", op: uint8(0x4F), t: wt.I32, a: wt.I32, b: wt.I32},
	{i: Eqz_I64, class: Numeric, name: "eqz.i64", name2: "Eqz_I64", op: uint8(0x50), t: wt.I32, a: wt.I64},
	{i: Eq_I64, class: Numeric, name: "eq.i64", name2: "Eq_I64", op: uint8(0x51), t: wt.I32, a: wt.I64, b: wt.I64},
	{i: Ne_I64, class: Numeric, name: "ne.i64", name2: "Ne_I64", op: uint8(0x52), t: wt.I32, a: wt.I64, b: wt.I64},
	{i: LtS_I64, class: Numeric, name: "lt_s.i64", name2: "LtS_I64", op: uint8(0x53), t: wt.I32, a: wt.I64, b: wt.I64},
	{i: LtU_I64, class: Numeric, name: "lt_u.i64", name2: "LtU_I64", op: uint8(0x54), t: wt.I32, a: wt.I64, b: wt.I64},
	{i: GtS_I64, class: Numeric, name: "gt_s.i64", name2: "GtS_I64", op: uint8(0x55), t: wt.I32, a: wt.I64, b: wt.I64},
	{i: GtU_I64, class: Numeric, name: "gt_u.i64", name2: "GtU_I64", op: uint8(0x56), t: wt.I32, a: wt.I64, b: wt.I64},
	{i: LeS_I64, class: Numeric, name: "le_s.i64", name2: "LeS_I64", op: uint8(0x57), t: wt.I32, a: wt.I64, b: wt.I64},
	{i: LeU_I64, class: Numeric, name: "le_u.i64", name2: "LeU_I64", op: uint8(0x58), t: wt.I32, a: wt.I64, b: wt.I64},
	{i: GeS_I64, class: Numeric, name: "ge_s.i64", name2: "GeS_I64", op: uint8(0x59), t: wt.I32, a: wt.I64, b: wt.I64},
	{i: GeU_I64, class: Numeric, name: "ge_u.i64", name2: "GeU_I64", op: uint8(0x5A), t: wt.I32, a: wt.I64, b: wt.I64},
	{i: Eq_F32, class: Numeric, name: "eq.f32", name2: "Eq_F32", op: uint8(0x5B), t: wt.I32, a: wt.F32},
	{i: Ne_F32, class: Numeric, name: "ne.f32", name2: "Ne_F32", op: uint8(0x5C), t: wt.I32, a: wt.F32, b: wt.F32},
	{i: Lt_F32, class: Numeric, name: "lt.f32", name2: "Lt_F32", op: uint8(0x5D), t: wt.I32, a: wt.F32, b: wt.F32},
	{i: Gt_F32, class: Numeric, name: "gt.f32", name2: "Gt_F32", op: uint8(0x5E), t: wt.I32, a: wt.F32, b: wt.F32},
	{i: Le_F32, class: Numeric, name: "le.f32", name2: "Le_F32", op: uint8(0x5F), t: wt.I32, a: wt.F32, b: wt.F32},
	{i: Ge_F32, class: Numeric, name: "ge.f32", name2: "Ge_F32", op: uint8(0x60), t: wt.I32, a: wt.F32, b: wt.F32},
	{i: Eq_F64, class: Numeric, name: "eq.f64", name2: "Eq_F64", op: uint8(0x61), t: wt.I64, a: wt.F64},
	{i: Ne_F64, class: Numeric, name: "ne.f64", name2: "Ne_F64", op: uint8(0x62), t: wt.I64, a: wt.F64, b: wt.F64},
	{i: Lt_F64, class: Numeric, name: "lt.f64", name2: "Lt_F64", op: uint8(0x63), t: wt.I64, a: wt.F64, b: wt.F64},
	{i: Gt_F64, class: Numeric, name: "gt.f64", name2: "Gt_F64", op: uint8(0x64), t: wt.I64, a: wt.F64, b: wt.F64},
	{i: Le_F64, class: Numeric, name: "le.f64", name2: "Le_F64", op: uint8(0x65), t: wt.I64, a: wt.F64, b: wt.F64},
	{i: Ge_F64, class: Numeric, name: "ge.f64", name2: "Ge_F64", op: uint8(0x66), t: wt.I64, a: wt.F64, b: wt.F64},
	{i: Clz_I32, class: Numeric, name: "clz.i32", name2: "Clz_I32", op: uint8(0x67), t: wt.I32, a: wt.I32},
	{i: Ctz_I32, class: Numeric, name: "ctz.i32", name2: "Ctz_I32", op: uint8(0x68), t: wt.I32, a: wt.I32},
	{i: Popcnt_I32, class: Numeric, name: "popcnt.i32", name2: "Popcnt_I32", op: uint8(0x69), t: wt.I32, a: wt.I32},
	{i: Add_I32, class: Numeric, name: "add.i32", name2: "Add_I32", op: uint8(0x6A), t: wt.I32, a: wt.I32, b: wt.I32},
	{i: Sub_I32, class: Numeric, name: "sub.i32", name2: "Sub_I32", op: uint8(0x6B), t: wt.I32, a: wt.I32, b: wt.I32},
	{i: Mul_I32, class: Numeric, name: "mul.i32", name2: "Mul_I32", op: uint8(0x6C), t: wt.I32, a: wt.I32, b: wt.I32},
	{i: DivS_I32, class: Numeric, name: "div_s.i32", name2: "DivS_I32", op: uint8(0x6D), t: wt.I32, a: wt.I32, b: wt.I32},
	{i: DivU_I32, class: Numeric, name: "div_u.i32", name2: "DivU_I32", op: uint8(0x6E), t: wt.I32, a: wt.I32, b: wt.I32},
	{i: RemS_I32, class: Numeric, name: "rem_s.i32", name2: "RemS_I32", op: uint8(0x6F), t: wt.I32, a: wt.I32, b: wt.I32},
	{i: RemU_I32, class: Numeric, name: "rem_u.i32", name2: "RemU_I32", op: uint8(0x70), t: wt.I32, a: wt.I32, b: wt.I32},
	{i: And_I32, class: Numeric, name: "and.i32", name2: "And_I32", op: uint8(0x71), t: wt.I32, a: wt.I32, b: wt.I32},
	{i: Or_I32, class: Numeric, name: "or.i32", name2: "Or_I32", op: uint8(0x72), t: wt.I32, a: wt.I32, b: wt.I32},
	{i: Xor_I32, class: Numeric, name: "xor.i32", name2: "Xor_I32", op: uint8(0x73), t: wt.I32, a: wt.I32, b: wt.I32},
	{i: Shl_I32, class: Numeric, name: "shl.i32", name2: "Shl_I32", op: uint8(0x74), t: wt.I32, a: wt.I32, b: wt.I32},
	{i: ShrS_I32, class: Numeric, name: "shr_s.i32", name2: "ShrS_I32", op: uint8(0x75), t: wt.I32, a: wt.I32, b: wt.I32},
	{i: ShrU_I32, class: Numeric, name: "shr_u.i32", name2: "ShrU_I32", op: uint8(0x76), t: wt.I32, a: wt.I32, b: wt.I32},
	{i: RotL_I32, class: Numeric, name: "rot_l.i32", name2: "RotL_I32", op: uint8(0x77), t: wt.I32, a: wt.I32, b: wt.I32},
	{i: RotR_I32, class: Numeric, name: "rot_r.i32", name2: "RotR_I32", op: uint8(0x78), t: wt.I32, a: wt.I32, b: wt.I32},
	{i: Clz_I64, class: Numeric, name: "clz.i64", name2: "Clz_I64", op: uint8(0x79), t: wt.I64, a: wt.I64},
	{i: Ctz_I64, class: Numeric, name: "ctz.i64", name2: "Ctz_I64", op: uint8(0x7A), t: wt.I64, a: wt.I64},
	{i: Popcnt_I64, class: Numeric, name: "popcnt.i64", name2: "Popcnt_I64", op: uint8(0x7B), t: wt.I64, a: wt.I64},
	{i: Add_I64, class: Numeric, name: "add.i64", name2: "Add_I64", op: uint8(0x7C), t: wt.I64, a: wt.I64, b: wt.I64},
	{i: Sub_I64, class: Numeric, name: "sub.i64", name2: "Sub_I64", op: uint8(0x7D), t: wt.I64, a: wt.I64, b: wt.I64},
	{i: Mul_I64, class: Numeric, name: "mul.i64", name2: "Mul_I64", op: uint8(0x7E), t: wt.I64, a: wt.I64, b: wt.I64},
	{i: DivS_I64, class: Numeric, name: "div_s.i64", name2: "DivS_I64", op: uint8(0x7F), t: wt.I64, a: wt.I64, b: wt.I64},
	{i: DivU_I64, class: Numeric, name: "div_u.i64", name2: "DivU_I64", op: uint8(0x80), t: wt.I64, a: wt.I64, b: wt.I64},
	{i: RemS_I64, class: Numeric, name: "rem_s.i64", name2: "RemS_I64", op: uint8(0x81), t: wt.I64, a: wt.I64, b: wt.I64},
	{i: RemU_I64, class: Numeric, name: "rem_u.i64", name2: "RemU_I64", op: uint8(0x82), t: wt.I64, a: wt.I64, b: wt.I64},
	{i: And_I64, class: Numeric, name: "and.i64", name2: "And_I64", op: uint8(0x83), t: wt.I64, a: wt.I64, b: wt.I64},
	{i: Or_I64, class: Numeric, name: "or.i64", name2: "Or_I64", op: uint8(0x84), t: wt.I64, a: wt.I64, b: wt.I64},
	{i: Xor_I64, class: Numeric, name: "xor.i64", name2: "Xor_I64", op: uint8(0x85), t: wt.I64, a: wt.I64, b: wt.I64},
	{i: Shl_I64, class: Numeric, name: "shl.i64", name2: "Shl_I64", op: uint8(0x86), t: wt.I64, a: wt.I64, b: wt.I64},
	{i: ShrS_I64, class: Numeric, name: "shr_s.i64", name2: "ShrS_I64", op: uint8(0x87), t: wt.I64, a: wt.I64, b: wt.I64},
	{i: ShrU_I64, class: Numeric, name: "shr_u.i64", name2: "ShrU_I64", op: uint8(0x88), t: wt.I64, a: wt.I64, b: wt.I64},
	{i: RotL_I64, class: Numeric, name: "rot_l.i64", name2: "RotL_I64", op: uint8(0x89), t: wt.I64, a: wt.I64, b: wt.I64},
	{i: RotR_I64, class: Numeric, name: "rot_r.i64", name2: "RotR_I64", op: uint8(0x8A), t: wt.I64, a: wt.I64, b: wt.I64},
	{i: Abs_F32, class: Numeric, name: "abs.f32", name2: "Abs_F32", op: uint8(0x8B), t: wt.F32, a: wt.F32, b: wt.F32},
	{i: Neg_F32, class: Numeric, name: "neg.f32", name2: "Neg_F32", op: uint8(0x8C), t: wt.F32, a: wt.F32, b: wt.F32},
	{i: Ceil_F32, class: Numeric, name: "ceil.f32", name2: "Ceil_F32", op: uint8(0x8D), t: wt.F32, a: wt.F32, b: wt.F32},
	{i: Floor_F32, class: Numeric, name: "floor.f32", name2: "Floor_F32", op: uint8(0x8E), t: wt.F32, a: wt.F32, b: wt.F32},
	{i: Trunc_F32, class: Numeric, name: "trunc.f32", name2: "Trunc_F32", op: uint8(0x8F), t: wt.F32, a: wt.F32, b: wt.F32},
	{i: Nearest_F32, class: Numeric, name: "nearest.f32", name2: "Nearest_F32", op: uint8(0x90), t: wt.F32, a: wt.F32, b: wt.F32},
	{i: Sqrt_F32, class: Numeric, name: "sqrt.f32", name2: "Sqrt_F32", op: uint8(0x91), t: wt.F32, a: wt.F32, b: wt.F32},
	{i: Add_F32, class: Numeric, name: "add.f32", name2: "Add_F32", op: uint8(0x92), t: wt.F32, a: wt.F32, b: wt.F32},
	{i: Sub_F32, class: Numeric, name: "sub.f32", name2: "Sub_F32", op: uint8(0x93), t: wt.F32, a: wt.F32, b: wt.F32},
	{i: Mul_F32, class: Numeric, name: "mul.f32", name2: "Mul_F32", op: uint8(0x94), t: wt.F32, a: wt.F32, b: wt.F32},
	{i: Div_F32, class: Numeric, name: "div.f32", name2: "Div_F32", op: uint8(0x95), t: wt.F32, a: wt.F32, b: wt.F32},
	{i: Min_F32, class: Numeric, name: "min.f32", name2: "Min_F32", op: uint8(0x96), t: wt.F32, a: wt.F32, b: wt.F32},
	{i: Max_F32, class: Numeric, name: "max.f32", name2: "Max_F32", op: uint8(0x97), t: wt.F32, a: wt.F32, b: wt.F32},
	{i: Copysign_F32, class: Numeric, name: "copysign.f32", name2: "Copysign_F32", op: uint8(0x98), t: wt.F32, a: wt.F32, b: wt.F32},
	{i: Abs_F64, class: Numeric, name: "abs.f64", name2: "Abs_F64", op: uint8(0x99), t: wt.F64, a: wt.F64, b: wt.F64},
	{i: Neg_F64, class: Numeric, name: "neg.f64", name2: "Neg_F64", op: uint8(0x9A), t: wt.F64, a: wt.F64, b: wt.F64},
	{i: Ceil_F64, class: Numeric, name: "ceil.f64", name2: "Ceil_F64", op: uint8(0x9B), t: wt.F64, a: wt.F64, b: wt.F64},
	{i: Floor_F64, class: Numeric, name: "floor.f64", name2: "Floor_F64", op: uint8(0x9C), t: wt.F64, a: wt.F64, b: wt.F64},
	{i: Trunc_F64, class: Numeric, name: "trunc.f64", name2: "Trunc_F64", op: uint8(0x9D), t: wt.F64, a: wt.F64, b: wt.F64},
	{i: Nearest_F64, class: Numeric, name: "nearest.f64", name2: "Nearest_F64", op: uint8(0x9E), t: wt.F64, a: wt.F64, b: wt.F64},
	{i: Sqrt_F64, class: Numeric, name: "sqrt.f64", name2: "Sqrt_F64", op: uint8(0x9F), t: wt.F64, a: wt.F64, b: wt.F64},
	{i: Add_F64, class: Numeric, name: "add.f64", name2: "Add_F64", op: uint8(0xA0), t: wt.F64, a: wt.F64, b: wt.F64},
	{i: Sub_F64, class: Numeric, name: "sub.f64", name2: "Sub_F64", op: uint8(0xA1), t: wt.F64, a: wt.F64, b: wt.F64},
	{i: Mul_F64, class: Numeric, name: "mul.f64", name2: "Mul_F64", op: uint8(0xA2), t: wt.F64, a: wt.F64, b: wt.F64},
	{i: Div_F64, class: Numeric, name: "div.f64", name2: "Div_F64", op: uint8(0xA3), t: wt.F64, a: wt.F64, b: wt.F64},
	{i: Min_F64, class: Numeric, name: "min.f64", name2: "Min_F64", op: uint8(0xA4), t: wt.F64, a: wt.F64, b: wt.F64},
	{i: Max_F64, class: Numeric, name: "max.f64", name2: "Max_F64", op: uint8(0xA5), t: wt.F64, a: wt.F64, b: wt.F64},
	{i: Copysign_F64, class: Numeric, name: "copysign.f64", name2: "Copysign_F64", op: uint8(0xA6), t: wt.F64, a: wt.F64, b: wt.F64},
	{i: WrapI64_I32, class: Numeric, name: "wrap_i64.i32", name2: "WrapI64_I32", op: uint8(0xA7), t: wt.I32, a: wt.I64},
	{i: TruncF32S_I32, class: Numeric, name: "trunc_f32_s.i32", name2: "TruncF32S_I32", op: uint8(0xA8), t: wt.I32, a: wt.F32},
	{i: TruncF32U_I32, class: Numeric, name: "trunc_f32_u.i32", name2: "TruncF32U_I32", op: uint8(0xA9), t: wt.I32, a: wt.F32},
	{i: TruncF64S_I32, class: Numeric, name: "trunc_f64_s.i32", name2: "TruncF64S_I32", op: uint8(0xAA), t: wt.I32, a: wt.F64},
	{i: TruncF64U_I32, class: Numeric, name: "trunc_f64_u.i32", name2: "TruncF64U_I32", op: uint8(0xAB), t: wt.I32, a: wt.F64},
	{i: ExtendI32S_I64, class: Numeric, name: "extend_i32_s.i64", name2: "ExtendI32S_I64", op: uint8(0xAC), t: wt.I64, a: wt.I32},
	{i: ExtendI32U_I64, class: Numeric, name: "extend_i32_u.i64", name2: "ExtendI32U_I64", op: uint8(0xAD), t: wt.I64, a: wt.I32},
	{i: TruncF32S_I64, class: Numeric, name: "trunc_f32_s.i64", name2: "TruncF32S_I64", op: uint8(0xAE), t: wt.I64, a: wt.F32},
	{i: TruncF32U_I64, class: Numeric, name: "trunc_f32_u.i64", name2: "TruncF32U_I64", op: uint8(0xAF), t: wt.I64, a: wt.F32},
	{i: TruncF64S_I64, class: Numeric, name: "trunc_f64_s.i64", name2: "TruncF64S_I64", op: uint8(0xB0), t: wt.I64, a: wt.F64},
	{i: TruncF64U_I64, class: Numeric, name: "trunc_f64_u.i64", name2: "TruncF64U_I64", op: uint8(0xB1), t: wt.I64, a: wt.F64},
	{i: ConvertI32S_F32, class: Numeric, name: "convert_i32_s.f32", name2: "ConvertI32S_F32", op: uint8(0xB2), t: wt.F32, a: wt.I32},
	{i: ConvertI32U_F32, class: Numeric, name: "convert_i32_u.f32", name2: "ConvertI32U_F32", op: uint8(0xB3), t: wt.F32, a: wt.I32},
	{i: ConvertI64S_F32, class: Numeric, name: "convert_i64_s.f32", name2: "ConvertI64S_F32", op: uint8(0xB4), t: wt.F32, a: wt.I64},
	{i: ConvertI64U_F32, class: Numeric, name: "convert_i64_u.f32", name2: "ConvertI64U_F32", op: uint8(0xB5), t: wt.F32, a: wt.I64},
	{i: DemoteF64_F32, class: Numeric, name: "demote_f64.f32", name2: "DemoteF64_F32", op: uint8(0xB6), t: wt.F32, a: wt.F64},
	{i: ConvertI32S_F64, class: Numeric, name: "convert_i32_s.f64", name2: "ConvertI32S_F64", op: uint8(0xB7), t: wt.F64, a: wt.I32},
	{i: ConvertI32U_F64, class: Numeric, name: "convert_i32_u.f64", name2: "ConvertI32U_F64", op: uint8(0xB8), t: wt.F64, a: wt.I32},
	{i: ConvertI64S_F64, class: Numeric, name: "convert_i64_s.f64", name2: "ConvertI64S_F64", op: uint8(0xB9), t: wt.F64, a: wt.I64},
	{i: ConvertI64U_F64, class: Numeric, name: "convert_i64_u.f64", name2: "ConvertI64U_F64", op: uint8(0xBA), t: wt.F64, a: wt.I64},
	{i: PromoteF32_F64, class: Numeric, name: "promote_f32.f64", name2: "PromoteF32_F64", op: uint8(0xBB), t: wt.F64, a: wt.F32},
	{i: ReinterpretF32_I32, class: Numeric, name: "reinterpret_f32.i32", name2: "ReinterpretF32_I32", op: uint8(0xBC), t: wt.I32, a: wt.F32},
	{i: ReinterpretF64_I64, class: Numeric, name: "reinterpret_f64.i64", name2: "ReinterpretF64_I64", op: uint8(0xBD), t: wt.I64, a: wt.F64},
	{i: ReinterpretI32_F32, class: Numeric, name: "reinterpret_i32.f32", name2: "ReinterpretI32_F32", op: uint8(0xBE), t: wt.F32, a: wt.I32},
	{i: ReinterpretI64_F64, class: Numeric, name: "reinterpret_i64.f64", name2: "ReinterpretI64_F64", op: uint8(0xBF), t: wt.F64, a: wt.I64},
	{i: Extend8S_I32, class: Numeric, name: "extend8_s.i32", name2: "Extend8S_I32", op: uint8(0xC0), t: wt.I32, a: wt.I32},
	{i: Extend16S_I32, class: Numeric, name: "extend16_s.i32", name2: "Extend16S_I32", op: uint8(0xC1), t: wt.I32, a: wt.I32},
	{i: Extend8S_I64, class: Numeric, name: "extend8_s.i64", name2: "Extend8S_I64", op: uint8(0xC2), t: wt.I64, a: wt.I64},
	{i: Extend16S_I64, class: Numeric, name: "extend16_s.i64", name2: "Extend16S_I64", op: uint8(0xC3), t: wt.I64, a: wt.I64},
	{i: Extend32S_I64, class: Numeric, name: "extend32_s.i64", name2: "Extend32S_I64", op: uint8(0xC4), t: wt.I64, a: wt.I64},
	{i: TruncSatF32S_I32, class: Numeric, name: "trunc_sat_f32_s.i32", name2: "TruncSatF32S_I32", opPfx: uint8(0xFC), op: uint8(0x0), t: wt.I32, a: wt.F32},
	{i: TruncSatF32U_I32, class: Numeric, name: "trunc_sat_f32_u.i32", name2: "TruncSatF32U_I32", opPfx: uint8(0xFC), op: uint8(0x1), t: wt.I32, a: wt.F32},
	{i: TruncSatF64S_I32, class: Numeric, name: "trunc_sat_f64_s.i32", name2: "TruncSatF64S_I32", opPfx: uint8(0xFC), op: uint8(0x2), t: wt.I32, a: wt.F64},
	{i: TruncSatF64U_I32, class: Numeric, name: "trunc_sat_f64_u.i32", name2: "TruncSatF64U_I32", opPfx: uint8(0xFC), op: uint8(0x3), t: wt.I32, a: wt.F64},
	{i: TruncSatF32S_I64, class: Numeric, name: "trunc_sat_f32_s.i64", name2: "TruncSatF32S_I64", opPfx: uint8(0xFC), op: uint8(0x4), t: wt.I64, a: wt.F32},
	{i: TruncSatF32U_I64, class: Numeric, name: "trunc_sat_f32_u.i64", name2: "TruncSatF32U_I64", opPfx: uint8(0xFC), op: uint8(0x5), t: wt.I64, a: wt.F32},
	{i: TruncSatF64S_I64, class: Numeric, name: "trunc_sat_f64_s.i64", name2: "TruncSatF64S_I64", opPfx: uint8(0xFC), op: uint8(0x6), t: wt.I64, a: wt.F64},
	{i: TruncSatF64U_I64, class: Numeric, name: "trunc_sat_f64_u.i64", name2: "TruncSatF64U_I64", opPfx: uint8(0xFC), op: uint8(0x7), t: wt.I64, a: wt.F64},
	{i: Load_V128, class: Vector, name: "load.v128", name2: "Load_V128", opPfx: uint8(0xFD), op: uint8(0x0), t: wt.V128, a: wt.I32, ma: Load, mz: 16},
	{i: Load8x8S_V128, class: Vector, name: "load8x8_s.v128", name2: "Load8x8S_V128", opPfx: uint8(0xFD), op: uint8(0x1), t: wt.V128, a: wt.I32, ma: Load, mz: 8},
	{i: Load8x8U_V128, class: Vector, name: "load8x8_u.v128", name2: "Load8x8U_V128", opPfx: uint8(0xFD), op: uint8(0x2), t: wt.V128, a: wt.I32, ma: Load, mz: 8},
	{i: Load16x4S_V128, class: Vector, name: "load16x4_s.v128", name2: "Load16x4S_V128", opPfx: uint8(0xFD), op: uint8(0x3), t: wt.V128, a: wt.I32, ma: Load, mz: 8},
	{i: Load16x4U_V128, class: Vector, name: "load16x4_u.v128", name2: "Load16x4U_V128", opPfx: uint8(0xFD), op: uint8(0x4), t: wt.V128, a: wt.I32, ma: Load, mz: 8},
	{i: Load32x2S_V128, class: Vector, name: "load32x2_s.v128", name2: "Load32x2S_V128", opPfx: uint8(0xFD), op: uint8(0x5), t: wt.V128, a: wt.I32, ma: Load, mz: 8},
	{i: Load32x2U_V128, class: Vector, name: "load32x2_u.v128", name2: "Load32x2U_V128", opPfx: uint8(0xFD), op: uint8(0x6), t: wt.V128, a: wt.I32, ma: Load, mz: 8},
	{i: Load8Splat_V128, class: Vector, name: "load8_splat.v128", name2: "Load8Splat_V128", opPfx: uint8(0xFD), op: uint8(0x7), t: wt.V128, a: wt.I32, ma: Load, mz: 1},
	{i: Load16Splat_V128, class: Vector, name: "load16_splat.v128", name2: "Load16Splat_V128", opPfx: uint8(0xFD), op: uint8(0x8), t: wt.V128, a: wt.I32, ma: Load, mz: 2},
	{i: Load32Splat_V128, class: Vector, name: "load32_splat.v128", name2: "Load32Splat_V128", opPfx: uint8(0xFD), op: uint8(0x9), t: wt.V128, a: wt.I32, ma: Load, mz: 4},
	{i: Load64Splat_V128, class: Vector, name: "load64_splat.v128", name2: "Load64Splat_V128", opPfx: uint8(0xFD), op: uint8(0xA), t: wt.V128, a: wt.I32, ma: Load, mz: 8},
	{i: Load32Zero_V128, class: Vector, name: "load32_zero.v128", name2: "Load32Zero_V128", opPfx: uint8(0xFD), op: uint8(0x5C), t: wt.V128, a: wt.I32, ma: Load, mz: 4},
	{i: Load64Zero_V128, class: Vector, name: "load64_zero.v128", name2: "Load64Zero_V128", opPfx: uint8(0xFD), op: uint8(0x5D), t: wt.V128, a: wt.I32, ma: Load, mz: 8},
	{i: Store_V128, class: Vector, name: "store.v128", name2: "Store_V128", opPfx: uint8(0xFD), op: uint8(0xB), a: wt.I32, b: wt.V128, ma: Store, mz: 16},
	{i: Load8Lane_V128, class: Vector, name: "load8_lane.v128", name2: "Load8Lane_V128", opPfx: uint8(0xFD), op: uint8(0x54), t: wt.V128, a: wt.I32, b: wt.V128, ma: Load, mz: 1},
	{i: Load16Lane_V128, class: Vector, name: "load16_lane.v128", name2: "Load16Lane_V128", opPfx: uint8(0xFD), op: uint8(0x55), t: wt.V128, a: wt.I32, b: wt.V128, ma: Load, mz: 2},
	{i: Load32Lane_V128, class: Vector, name: "load32_lane.v128", name2: "Load32Lane_V128", opPfx: uint8(0xFD), op: uint8(0x56), t: wt.V128, a: wt.I32, b: wt.V128, ma: Load, mz: 4},
	{i: Load64Lane_V128, class: Vector, name: "load64_lane.v128", name2: "Load64Lane_V128", opPfx: uint8(0xFD), op: uint8(0x57), t: wt.V128, a: wt.I32, b: wt.V128, ma: Load, mz: 8},
	{i: Store8Lane_V128, class: Vector, name: "store8_lane.v128", name2: "Store8Lane_V128", opPfx: uint8(0xFD), op: uint8(0x58), a: wt.I32, b: wt.V128, ma: Store, mz: 1},
	{i: Store16Lane_V128, class: Vector, name: "store16_lane.v128", name2: "Store16Lane_V128", opPfx: uint8(0xFD), op: uint8(0x59), a: wt.I32, b: wt.V128, ma: Store, mz: 2},
	{i: Store32Lane_V128, class: Vector, name: "store32_lane.v128", name2: "Store32Lane_V128", opPfx: uint8(0xFD), op: uint8(0x5A), a: wt.I32, b: wt.V128, ma: Store, mz: 4},
	{i: Store64Lane_V128, class: Vector, name: "store64_lane.v128", name2: "Store64Lane_V128", opPfx: uint8(0xFD), op: uint8(0x5B), a: wt.I32, b: wt.V128, ma: Store, mz: 8},
	{i: Const_V128, class: Vector, name: "const.v128", name2: "Const_V128", opPfx: uint8(0xFD), op: uint8(0xC), t: wt.V128},
	{i: Shuffle_I8x16, class: Vector, name: "shuffle.i8x16", name2: "Shuffle_I8x16", opPfx: uint8(0xFD), op: uint8(0xD), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: ExtractLaneS_I8x16, class: Vector, name: "extract_lane_s.i8x16", name2: "ExtractLaneS_I8x16", opPfx: uint8(0xFD), op: uint8(0x15), t: wt.I32, a: wt.V128},
	{i: ExtractLaneU_I8x16, class: Vector, name: "extract_lane_u.i8x16", name2: "ExtractLaneU_I8x16", opPfx: uint8(0xFD), op: uint8(0x16), t: wt.I32, a: wt.V128},
	{i: ReplaceLane_I8x16, class: Vector, name: "replace_lane.i8x16", name2: "ReplaceLane_I8x16", opPfx: uint8(0xFD), op: uint8(0x17), t: wt.V128, a: wt.V128, b: wt.I32},
	{i: ExtractLaneS_I16x8, class: Vector, name: "extract_lane_s.i16x8", name2: "ExtractLaneS_I16x8", opPfx: uint8(0xFD), op: uint8(0x18), t: wt.I32, a: wt.V128},
	{i: ExtractLaneU_I16x8, class: Vector, name: "extract_lane_u.i16x8", name2: "ExtractLaneU_I16x8", opPfx: uint8(0xFD), op: uint8(0x19), t: wt.I32, a: wt.V128},
	{i: ReplaceLane_I16x8, class: Vector, name: "replace_lane.i16x8", name2: "ReplaceLane_I16x8", opPfx: uint8(0xFD), op: uint8(0x1A), t: wt.V128, a: wt.V128, b: wt.I32},
	{i: ExtractLane_I32x4, class: Vector, name: "extract_lane.i32x4", name2: "ExtractLane_I32x4", opPfx: uint8(0xFD), op: uint8(0x1B), t: wt.I32, a: wt.V128},
	{i: ReplaceLane_I32x4, class: Vector, name: "replace_lane.i32x4", name2: "ReplaceLane_I32x4", opPfx: uint8(0xFD), op: uint8(0x1C), t: wt.V128, a: wt.V128, b: wt.I32},
	{i: ExtractLane_I64x2, class: Vector, name: "extract_lane.i64x2", name2: "ExtractLane_I64x2", opPfx: uint8(0xFD), op: uint8(0x1D), t: wt.I64, a: wt.V128},
	{i: ReplaceLane_I64x2, class: Vector, name: "replace_lane.i64x2", name2: "ReplaceLane_I64x2", opPfx: uint8(0xFD), op: uint8(0x1E), t: wt.V128, a: wt.V128, b: wt.I64},
	{i: ExtractLane_F32x4, class: Vector, name: "extract_lane.f32x4", name2: "ExtractLane_F32x4", opPfx: uint8(0xFD), op: uint8(0x1F), t: wt.F32, a: wt.V128},
	{i: ReplaceLane_F32x4, class: Vector, name: "replace_lane.f32x4", name2: "ReplaceLane_F32x4", opPfx: uint8(0xFD), op: uint8(0x20), t: wt.V128, a: wt.V128, b: wt.F32},
	{i: ExtractLane_F64x2, class: Vector, name: "extract_lane.f64x2", name2: "ExtractLane_F64x2", opPfx: uint8(0xFD), op: uint8(0x21), t: wt.F64, a: wt.V128},
	{i: ReplaceLane_F64x2, class: Vector, name: "replace_lane.f64x2", name2: "ReplaceLane_F64x2", opPfx: uint8(0xFD), op: uint8(0x22), t: wt.V128, a: wt.V128, b: wt.F64},
	{i: Swizzle_I8x16, class: Vector, name: "swizzle.i8x16", name2: "Swizzle_I8x16", opPfx: uint8(0xFD), op: uint8(0xE), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Splat_I8x16, class: Vector, name: "splat.i8x16", name2: "Splat_I8x16", opPfx: uint8(0xFD), op: uint8(0xF), t: wt.V128, a: wt.I32},
	{i: Splat_I16x8, class: Vector, name: "splat.i16x8", name2: "Splat_I16x8", opPfx: uint8(0xFD), op: uint8(0x10), t: wt.V128, a: wt.I32},
	{i: Splat_I32x4, class: Vector, name: "splat.i32x4", name2: "Splat_I32x4", opPfx: uint8(0xFD), op: uint8(0x11), t: wt.V128, a: wt.I32},
	{i: Splat_I64x2, class: Vector, name: "splat.i64x2", name2: "Splat_I64x2", opPfx: uint8(0xFD), op: uint8(0x12), t: wt.V128, a: wt.I64},
	{i: Splat_F32x4, class: Vector, name: "splat.f32x4", name2: "Splat_F32x4", opPfx: uint8(0xFD), op: uint8(0x13), t: wt.V128, a: wt.F32},
	{i: Splat_F64x2, class: Vector, name: "splat.f64x2", name2: "Splat_F64x2", opPfx: uint8(0xFD), op: uint8(0x14), t: wt.V128, a: wt.F64},
	{i: Eq_I8x16, class: Vector, name: "eq.i8x16", name2: "Eq_I8x16", opPfx: uint8(0xFD), op: uint8(0x23), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Ne_I8x16, class: Vector, name: "ne.i8x16", name2: "Ne_I8x16", opPfx: uint8(0xFD), op: uint8(0x24), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: LtS_I8x16, class: Vector, name: "lt_s.i8x16", name2: "LtS_I8x16", opPfx: uint8(0xFD), op: uint8(0x25), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: LtU_I8x16, class: Vector, name: "lt_u.i8x16", name2: "LtU_I8x16", opPfx: uint8(0xFD), op: uint8(0x26), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: GtS_I8x16, class: Vector, name: "gt_s.i8x16", name2: "GtS_I8x16", opPfx: uint8(0xFD), op: uint8(0x27), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: GtU_I8x16, class: Vector, name: "gt_u.i8x16", name2: "GtU_I8x16", opPfx: uint8(0xFD), op: uint8(0x28), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: LeS_I8x16, class: Vector, name: "le_s.i8x16", name2: "LeS_I8x16", opPfx: uint8(0xFD), op: uint8(0x29), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: LeU_I8x16, class: Vector, name: "le_u.i8x16", name2: "LeU_I8x16", opPfx: uint8(0xFD), op: uint8(0x2A), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: GeS_I8x16, class: Vector, name: "ge_s.i8x16", name2: "GeS_I8x16", opPfx: uint8(0xFD), op: uint8(0x2B), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: GeU_I8x16, class: Vector, name: "ge_u.i8x16", name2: "GeU_I8x16", opPfx: uint8(0xFD), op: uint8(0x2C), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Eq_I16x8, class: Vector, name: "eq.i16x8", name2: "Eq_I16x8", opPfx: uint8(0xFD), op: uint8(0x2D), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Ne_I16x8, class: Vector, name: "ne.i16x8", name2: "Ne_I16x8", opPfx: uint8(0xFD), op: uint8(0x2E), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: LtS_I16x8, class: Vector, name: "lt_s.i16x8", name2: "LtS_I16x8", opPfx: uint8(0xFD), op: uint8(0x2F), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: LtU_I16x8, class: Vector, name: "lt_u.i16x8", name2: "LtU_I16x8", opPfx: uint8(0xFD), op: uint8(0x30), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: GtS_I16x8, class: Vector, name: "gt_s.i16x8", name2: "GtS_I16x8", opPfx: uint8(0xFD), op: uint8(0x31), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: GtU_I16x8, class: Vector, name: "gt_u.i16x8", name2: "GtU_I16x8", opPfx: uint8(0xFD), op: uint8(0x32), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: LeS_I16x8, class: Vector, name: "le_s.i16x8", name2: "LeS_I16x8", opPfx: uint8(0xFD), op: uint8(0x33), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: LeU_I16x8, class: Vector, name: "le_u.i16x8", name2: "LeU_I16x8", opPfx: uint8(0xFD), op: uint8(0x34), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: GeS_I16x8, class: Vector, name: "ge_s.i16x8", name2: "GeS_I16x8", opPfx: uint8(0xFD), op: uint8(0x35), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: GeU_I16x8, class: Vector, name: "ge_u.i16x8", name2: "GeU_I16x8", opPfx: uint8(0xFD), op: uint8(0x36), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Eq_I32x4, class: Vector, name: "eq.i32x4", name2: "Eq_I32x4", opPfx: uint8(0xFD), op: uint8(0x37), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Ne_I32x4, class: Vector, name: "ne.i32x4", name2: "Ne_I32x4", opPfx: uint8(0xFD), op: uint8(0x38), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: LtS_I32x4, class: Vector, name: "lt_s.i32x4", name2: "LtS_I32x4", opPfx: uint8(0xFD), op: uint8(0x39), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: LtU_I32x4, class: Vector, name: "lt_u.i32x4", name2: "LtU_I32x4", opPfx: uint8(0xFD), op: uint8(0x3A), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: GtS_I32x4, class: Vector, name: "gt_s.i32x4", name2: "GtS_I32x4", opPfx: uint8(0xFD), op: uint8(0x3B), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: GtU_I32x4, class: Vector, name: "gt_u.i32x4", name2: "GtU_I32x4", opPfx: uint8(0xFD), op: uint8(0x3C), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: LeS_I32x4, class: Vector, name: "le_s.i32x4", name2: "LeS_I32x4", opPfx: uint8(0xFD), op: uint8(0x3D), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: LeU_I32x4, class: Vector, name: "le_u.i32x4", name2: "LeU_I32x4", opPfx: uint8(0xFD), op: uint8(0x3E), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: GeS_I32x4, class: Vector, name: "ge_s.i32x4", name2: "GeS_I32x4", opPfx: uint8(0xFD), op: uint8(0x3F), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: GeU_I32x4, class: Vector, name: "ge_u.i32x4", name2: "GeU_I32x4", opPfx: uint8(0xFD), op: uint8(0x40), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Eq_I64x2, class: Vector, name: "eq.i64x2", name2: "Eq_I64x2", opPfx: uint8(0xFD), op: uint8(0xD6), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Ne_I64x2, class: Vector, name: "ne.i64x2", name2: "Ne_I64x2", opPfx: uint8(0xFD), op: uint8(0xD7), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: LtS_I64x2, class: Vector, name: "lt_s.i64x2", name2: "LtS_I64x2", opPfx: uint8(0xFD), op: uint8(0xD8), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: GtS_I64x2, class: Vector, name: "gt_s.i64x2", name2: "GtS_I64x2", opPfx: uint8(0xFD), op: uint8(0xD9), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: LeS_I64x2, class: Vector, name: "le_s.i64x2", name2: "LeS_I64x2", opPfx: uint8(0xFD), op: uint8(0xDA), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: GeS_I64x2, class: Vector, name: "ge_s.i64x2", name2: "GeS_I64x2", opPfx: uint8(0xFD), op: uint8(0xDB), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Eq_F32x4, class: Vector, name: "eq.f32x4", name2: "Eq_F32x4", opPfx: uint8(0xFD), op: uint8(0x41), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Ne_F32x4, class: Vector, name: "ne.f32x4", name2: "Ne_F32x4", opPfx: uint8(0xFD), op: uint8(0x42), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Lt_F32x4, class: Vector, name: "lt.f32x4", name2: "Lt_F32x4", opPfx: uint8(0xFD), op: uint8(0x43), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Gt_F32x4, class: Vector, name: "gt.f32x4", name2: "Gt_F32x4", opPfx: uint8(0xFD), op: uint8(0x44), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Le_F32x4, class: Vector, name: "le.f32x4", name2: "Le_F32x4", opPfx: uint8(0xFD), op: uint8(0x45), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Ge_F32x4, class: Vector, name: "ge.f32x4", name2: "Ge_F32x4", opPfx: uint8(0xFD), op: uint8(0x46), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Eq_F64x2, class: Vector, name: "eq.f64x2", name2: "Eq_F64x2", opPfx: uint8(0xFD), op: uint8(0x47), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Ne_F64x2, class: Vector, name: "ne.f64x2", name2: "Ne_F64x2", opPfx: uint8(0xFD), op: uint8(0x48), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Lt_F64x2, class: Vector, name: "lt.f64x2", name2: "Lt_F64x2", opPfx: uint8(0xFD), op: uint8(0x49), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Gt_F64x2, class: Vector, name: "gt.f64x2", name2: "Gt_F64x2", opPfx: uint8(0xFD), op: uint8(0x4A), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Le_F64x2, class: Vector, name: "le.f64x2", name2: "Le_F64x2", opPfx: uint8(0xFD), op: uint8(0x4B), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Ge_F64x2, class: Vector, name: "ge.f64x2", name2: "Ge_F64x2", opPfx: uint8(0xFD), op: uint8(0x4C), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Not_V128, class: Vector, name: "not.v128", name2: "Not_V128", opPfx: uint8(0xFD), op: uint8(0x4D), t: wt.V128, a: wt.V128},
	{i: And_V128, class: Vector, name: "and.v128", name2: "And_V128", opPfx: uint8(0xFD), op: uint8(0x4E), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Andnot_V128, class: Vector, name: "andnot.v128", name2: "Andnot_V128", opPfx: uint8(0xFD), op: uint8(0x4F), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Or_V128, class: Vector, name: "or.v128", name2: "Or_V128", opPfx: uint8(0xFD), op: uint8(0x50), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Xor_V128, class: Vector, name: "xor.v128", name2: "Xor_V128", opPfx: uint8(0xFD), op: uint8(0x51), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Bitselect_V128, class: Vector, name: "bitselect.v128", name2: "Bitselect_V128", opPfx: uint8(0xFD), op: uint8(0x52), t: wt.V128, a: wt.V128, b: wt.V128, c: wt.V128},
	{i: AnyTrue_V128, class: Vector, name: "any_true.v128", name2: "AnyTrue_V128", opPfx: uint8(0xFD), op: uint8(0x53), t: wt.I32, a: wt.V128},
	{i: Abs_I8x16, class: Vector, name: "abs.i8x16", name2: "Abs_I8x16", opPfx: uint8(0xFD), op: uint8(0x60), t: wt.V128, a: wt.V128},
	{i: Neg_I8x16, class: Vector, name: "neg.i8x16", name2: "Neg_I8x16", opPfx: uint8(0xFD), op: uint8(0x61), t: wt.V128, a: wt.V128},
	{i: Popcnt_I8x16, class: Vector, name: "popcnt.i8x16", name2: "Popcnt_I8x16", opPfx: uint8(0xFD), op: uint8(0x62), t: wt.V128, a: wt.V128},
	{i: AllTrue_I8x16, class: Vector, name: "all_true.i8x16", name2: "AllTrue_I8x16", opPfx: uint8(0xFD), op: uint8(0x63), t: wt.I32, a: wt.V128},
	{i: Bitmask_I8x16, class: Vector, name: "bitmask.i8x16", name2: "Bitmask_I8x16", opPfx: uint8(0xFD), op: uint8(0x64), t: wt.I32, a: wt.V128},
	{i: NarrowI16x8S_I8x16, class: Vector, name: "narrow_i16x8_s.i8x16", name2: "NarrowI16x8S_I8x16", opPfx: uint8(0xFD), op: uint8(0x65), t: wt.V128, a: wt.V128},
	{i: NarrowI16x8U_I8x16, class: Vector, name: "narrow_i16x8_u.i8x16", name2: "NarrowI16x8U_I8x16", opPfx: uint8(0xFD), op: uint8(0x66), t: wt.V128, a: wt.V128},
	{i: Shl_I8x16, class: Vector, name: "shl.i8x16", name2: "Shl_I8x16", opPfx: uint8(0xFD), op: uint8(0x6B), t: wt.V128, a: wt.V128, b: wt.I32},
	{i: ShrS_I8x16, class: Vector, name: "shr_s.i8x16", name2: "ShrS_I8x16", opPfx: uint8(0xFD), op: uint8(0x6C), t: wt.V128, a: wt.V128, b: wt.I32},
	{i: ShrU_I8x16, class: Vector, name: "shr_u.i8x16", name2: "ShrU_I8x16", opPfx: uint8(0xFD), op: uint8(0x6D), t: wt.V128, a: wt.V128, b: wt.I32},
	{i: Add_I8x16, class: Vector, name: "add.i8x16", name2: "Add_I8x16", opPfx: uint8(0xFD), op: uint8(0x6E), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: AddSatS_I8x16, class: Vector, name: "add_sat_s.i8x16", name2: "AddSatS_I8x16", opPfx: uint8(0xFD), op: uint8(0x6F), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: AddSatU_I8x16, class: Vector, name: "add_sat_u.i8x16", name2: "AddSatU_I8x16", opPfx: uint8(0xFD), op: uint8(0x70), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Sub_I8x16, class: Vector, name: "sub.i8x16", name2: "Sub_I8x16", opPfx: uint8(0xFD), op: uint8(0x71), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: SubSatS_I8x16, class: Vector, name: "sub_sat_s.i8x16", name2: "SubSatS_I8x16", opPfx: uint8(0xFD), op: uint8(0x72), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: SubSatU_I8x16, class: Vector, name: "sub_sat_u.i8x16", name2: "SubSatU_I8x16", opPfx: uint8(0xFD), op: uint8(0x73), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: MinS_I8x16, class: Vector, name: "min_s.i8x16", name2: "MinS_I8x16", opPfx: uint8(0xFD), op: uint8(0x76), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: MinU_I8x16, class: Vector, name: "min_u.i8x16", name2: "MinU_I8x16", opPfx: uint8(0xFD), op: uint8(0x77), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: MaxS_I8x16, class: Vector, name: "max_s.i8x16", name2: "MaxS_I8x16", opPfx: uint8(0xFD), op: uint8(0x78), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: MaxU_I8x16, class: Vector, name: "max_u.i8x16", name2: "MaxU_I8x16", opPfx: uint8(0xFD), op: uint8(0x79), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: AvgrU_I8x16, class: Vector, name: "avgr_u.i8x16", name2: "AvgrU_I8x16", opPfx: uint8(0xFD), op: uint8(0x7B), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: ExtaddPairwiseI8x16S_I16x8, class: Vector, name: "extadd_pairwise_i8x16_s.i16x8", name2: "ExtaddPairwiseI8x16S_I16x8", opPfx: uint8(0xFD), op: uint8(0x7C), t: wt.V128, a: wt.V128},
	{i: ExtaddPairwiseI8x16U_I16x8, class: Vector, name: "extadd_pairwise_i8x16_u.i16x8", name2: "ExtaddPairwiseI8x16U_I16x8", opPfx: uint8(0xFD), op: uint8(0x7D), t: wt.V128, a: wt.V128},
	{i: Abs_I16x8, class: Vector, name: "abs.i16x8", name2: "Abs_I16x8", opPfx: uint8(0xFD), op: uint8(0x80), t: wt.V128, a: wt.V128},
	{i: Neg_I16x8, class: Vector, name: "neg.i16x8", name2: "Neg_I16x8", opPfx: uint8(0xFD), op: uint8(0x81), t: wt.V128, a: wt.V128},
	{i: Q15mulrSatS_I16x8, class: Vector, name: "q15mulr_sat_s.i16x8", name2: "Q15mulrSatS_I16x8", opPfx: uint8(0xFD), op: uint8(0x82), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Alltrue_I16x8, class: Vector, name: "alltrue.i16x8", name2: "Alltrue_I16x8", opPfx: uint8(0xFD), op: uint8(0x83), t: wt.I32, a: wt.V128},
	{i: Bitmask_I16x8, class: Vector, name: "bitmask.i16x8", name2: "Bitmask_I16x8", opPfx: uint8(0xFD), op: uint8(0x84), t: wt.I32, a: wt.V128},
	{i: NarrowI32x4S_I16x8, class: Vector, name: "narrow_i32x4_s.i16x8", name2: "NarrowI32x4S_I16x8", opPfx: uint8(0xFD), op: uint8(0x85), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: NarrowI32x4U_I16x8, class: Vector, name: "narrow_i32x4_u.i16x8", name2: "NarrowI32x4U_I16x8", opPfx: uint8(0xFD), op: uint8(0x86), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: ExtendLowI8x16S_I16x8, class: Vector, name: "extend_low_i8x16_s.i16x8", name2: "ExtendLowI8x16S_I16x8", opPfx: uint8(0xFD), op: uint8(0x87), t: wt.V128, a: wt.V128},
	{i: ExtendHighI8x16S_I16x8, class: Vector, name: "extend_high_i8x16_s.i16x8", name2: "ExtendHighI8x16S_I16x8", opPfx: uint8(0xFD), op: uint8(0x88), t: wt.V128, a: wt.V128},
	{i: ExtendLowI8x16U_I16x8, class: Vector, name: "extend_low_i8x16_u.i16x8", name2: "ExtendLowI8x16U_I16x8", opPfx: uint8(0xFD), op: uint8(0x89), t: wt.V128, a: wt.V128},
	{i: ExtendHighI8x16U_I16x8, class: Vector, name: "extend_high_i8x16_u.i16x8", name2: "ExtendHighI8x16U_I16x8", opPfx: uint8(0xFD), op: uint8(0x8A), t: wt.V128, a: wt.V128},
	{i: Shl_I16x8, class: Vector, name: "shl.i16x8", name2: "Shl_I16x8", opPfx: uint8(0xFD), op: uint8(0x8B), t: wt.V128, a: wt.V128, b: wt.I32},
	{i: ShrS_I16x8, class: Vector, name: "shr_s.i16x8", name2: "ShrS_I16x8", opPfx: uint8(0xFD), op: uint8(0x8C), t: wt.V128, a: wt.V128, b: wt.I32},
	{i: ShrU_I16x8, class: Vector, name: "shr_u.i16x8", name2: "ShrU_I16x8", opPfx: uint8(0xFD), op: uint8(0x8D), t: wt.V128, a: wt.V128, b: wt.I32},
	{i: Add_I16x8, class: Vector, name: "add.i16x8", name2: "Add_I16x8", opPfx: uint8(0xFD), op: uint8(0x8E), t: wt.I32, a: wt.V128, b: wt.V128},
	{i: AddSatS_I16x8, class: Vector, name: "add_sat_s.i16x8", name2: "AddSatS_I16x8", opPfx: uint8(0xFD), op: uint8(0x8F), t: wt.I32, a: wt.V128, b: wt.V128},
	{i: AddSatU_I16x8, class: Vector, name: "add_sat_u.i16x8", name2: "AddSatU_I16x8", opPfx: uint8(0xFD), op: uint8(0x90), t: wt.I32, a: wt.V128, b: wt.V128},
	{i: Sub_I16x8, class: Vector, name: "sub.i16x8", name2: "Sub_I16x8", opPfx: uint8(0xFD), op: uint8(0x91), t: wt.I32, a: wt.V128, b: wt.V128},
	{i: SubSatS_I16x8, class: Vector, name: "sub_sat_s.i16x8", name2: "SubSatS_I16x8", opPfx: uint8(0xFD), op: uint8(0x92), t: wt.I32, a: wt.V128, b: wt.V128},
	{i: SubSatU_I16x8, class: Vector, name: "sub_sat_u.i16x8", name2: "SubSatU_I16x8", opPfx: uint8(0xFD), op: uint8(0x93), t: wt.I32, a: wt.V128, b: wt.V128},
	{i: Mul_I16x8, class: Vector, name: "mul.i16x8", name2: "Mul_I16x8", opPfx: uint8(0xFD), op: uint8(0x95), t: wt.I32, a: wt.V128, b: wt.V128},
	{i: MinS_I16x8, class: Vector, name: "min_s.i16x8", name2: "MinS_I16x8", opPfx: uint8(0xFD), op: uint8(0x96), t: wt.I32, a: wt.V128, b: wt.V128},
	{i: MinU_I16x8, class: Vector, name: "min_u.i16x8", name2: "MinU_I16x8", opPfx: uint8(0xFD), op: uint8(0x97), t: wt.I32, a: wt.V128, b: wt.V128},
	{i: MaxS_I16x8, class: Vector, name: "max_s.i16x8", name2: "MaxS_I16x8", opPfx: uint8(0xFD), op: uint8(0x98), t: wt.I32, a: wt.V128, b: wt.V128},
	{i: MaxU_I16x8, class: Vector, name: "max_u.i16x8", name2: "MaxU_I16x8", opPfx: uint8(0xFD), op: uint8(0x99), t: wt.I32, a: wt.V128, b: wt.V128},
	{i: AvgrU_I16x8, class: Vector, name: "avgr_u.i16x8", name2: "AvgrU_I16x8", opPfx: uint8(0xFD), op: uint8(0x9B), t: wt.I32, a: wt.V128, b: wt.V128},
	{i: ExtmulLowI8x16S_I16x8, class: Vector, name: "extmul_low_i8x16_s.i16x8", name2: "ExtmulLowI8x16S_I16x8", opPfx: uint8(0xFD), op: uint8(0x9C), t: wt.I32, a: wt.V128, b: wt.V128},
	{i: ExtmulHighI8x16S_I16x8, class: Vector, name: "extmul_high_i8x16_s.i16x8", name2: "ExtmulHighI8x16S_I16x8", opPfx: uint8(0xFD), op: uint8(0x9D), t: wt.I32, a: wt.V128, b: wt.V128},
	{i: ExtmulLowI8x16U_I16x8, class: Vector, name: "extmul_low_i8x16_u.i16x8", name2: "ExtmulLowI8x16U_I16x8", opPfx: uint8(0xFD), op: uint8(0x9E), t: wt.I32, a: wt.V128, b: wt.V128},
	{i: ExtmulHighI8x16U_I16x8, class: Vector, name: "extmul_high_i8x16_u.i16x8", name2: "ExtmulHighI8x16U_I16x8", opPfx: uint8(0xFD), op: uint8(0x9F), t: wt.I32, a: wt.V128, b: wt.V128},
	{i: ExtaddPairwiseI16x8S_I32x4, class: Vector, name: "extadd_pairwise_i16x8_s.i32x4", name2: "ExtaddPairwiseI16x8S_I32x4", opPfx: uint8(0xFD), op: uint8(0x7E), t: wt.V128, a: wt.V128},
	{i: ExtaddPairwiseI16x8U_I32x4, class: Vector, name: "extadd_pairwise_i16x8_u.i32x4", name2: "ExtaddPairwiseI16x8U_I32x4", opPfx: uint8(0xFD), op: uint8(0x7F), t: wt.V128, a: wt.V128},
	{i: Abs_I32x4, class: Vector, name: "abs.i32x4", name2: "Abs_I32x4", opPfx: uint8(0xFD), op: uint8(0xA0), t: wt.I32, a: wt.V128},
	{i: Neg_I32x4, class: Vector, name: "neg.i32x4", name2: "Neg_I32x4", opPfx: uint8(0xFD), op: uint8(0xA1), t: wt.I32, a: wt.V128},
	{i: Alltrue_I32x4, class: Vector, name: "alltrue.i32x4", name2: "Alltrue_I32x4", opPfx: uint8(0xFD), op: uint8(0xA3), t: wt.I32, a: wt.V128},
	{i: Bitmask_I32x4, class: Vector, name: "bitmask.i32x4", name2: "Bitmask_I32x4", opPfx: uint8(0xFD), op: uint8(0xA4), t: wt.I32, a: wt.V128},
	{i: ExtendLowI16x8S_I32x4, class: Vector, name: "extend_low_i16x8_s.i32x4", name2: "ExtendLowI16x8S_I32x4", opPfx: uint8(0xFD), op: uint8(0xA7), t: wt.V128, a: wt.V128},
	{i: ExtendHighI16x8S_I32x4, class: Vector, name: "extend_high_i16x8_s.i32x4", name2: "ExtendHighI16x8S_I32x4", opPfx: uint8(0xFD), op: uint8(0xA8), t: wt.V128, a: wt.V128},
	{i: ExtendLowI16x8U_I32x4, class: Vector, name: "extend_low_i16x8_u.i32x4", name2: "ExtendLowI16x8U_I32x4", opPfx: uint8(0xFD), op: uint8(0xA9), t: wt.V128, a: wt.V128},
	{i: ExtendHighI16x8U_I32x4, class: Vector, name: "extend_high_i16x8_u.i32x4", name2: "ExtendHighI16x8U_I32x4", opPfx: uint8(0xFD), op: uint8(0xAA), t: wt.V128, a: wt.V128},
	{i: Shl_I32x4, class: Vector, name: "shl.i32x4", name2: "Shl_I32x4", opPfx: uint8(0xFD), op: uint8(0xAB), t: wt.V128, a: wt.V128, b: wt.I32},
	{i: ShrS_I32x4, class: Vector, name: "shr_s.i32x4", name2: "ShrS_I32x4", opPfx: uint8(0xFD), op: uint8(0xAC), t: wt.V128, a: wt.V128, b: wt.I32},
	{i: ShrU_I32x4, class: Vector, name: "shr_u.i32x4", name2: "ShrU_I32x4", opPfx: uint8(0xFD), op: uint8(0xAD), t: wt.V128, a: wt.V128, b: wt.I32},
	{i: Add_I32x4, class: Vector, name: "add.i32x4", name2: "Add_I32x4", opPfx: uint8(0xFD), op: uint8(0xAE), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Sub_I32x4, class: Vector, name: "sub.i32x4", name2: "Sub_I32x4", opPfx: uint8(0xFD), op: uint8(0xB1), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Mul_I32x4, class: Vector, name: "mul.i32x4", name2: "Mul_I32x4", opPfx: uint8(0xFD), op: uint8(0xB5), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: MinS_I32x4, class: Vector, name: "min_s.i32x4", name2: "MinS_I32x4", opPfx: uint8(0xFD), op: uint8(0xB6), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: MinU_I32x4, class: Vector, name: "min_u.i32x4", name2: "MinU_I32x4", opPfx: uint8(0xFD), op: uint8(0xB7), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: MaxS_I32x4, class: Vector, name: "max_s.i32x4", name2: "MaxS_I32x4", opPfx: uint8(0xFD), op: uint8(0xB8), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: MaxU_I32x4, class: Vector, name: "max_u.i32x4", name2: "MaxU_I32x4", opPfx: uint8(0xFD), op: uint8(0xB9), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: DotI16x8S_I32x4, class: Vector, name: "dot_i16x8_s.i32x4", name2: "DotI16x8S_I32x4", opPfx: uint8(0xFD), op: uint8(0xBA), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: ExtmulLowI16x8S_I32x4, class: Vector, name: "extmul_low_i16x8_s.i32x4", name2: "ExtmulLowI16x8S_I32x4", opPfx: uint8(0xFD), op: uint8(0xBC), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: ExtmulHighI16x8S_I32x4, class: Vector, name: "extmul_high_i16x8_s.i32x4", name2: "ExtmulHighI16x8S_I32x4", opPfx: uint8(0xFD), op: uint8(0xBD), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: ExtmulLowI16x8U_I32x4, class: Vector, name: "extmul_low_i16x8_u.i32x4", name2: "ExtmulLowI16x8U_I32x4", opPfx: uint8(0xFD), op: uint8(0xBE), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: ExtmulHighI16x8U_I32x4, class: Vector, name: "extmul_high_i16x8_u.i32x4", name2: "ExtmulHighI16x8U_I32x4", opPfx: uint8(0xFD), op: uint8(0xBF), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Abs_I64x2, class: Vector, name: "abs.i64x2", name2: "Abs_I64x2", opPfx: uint8(0xFD), op: uint8(0xC0), t: wt.V128, a: wt.V128},
	{i: Neg_I64x2, class: Vector, name: "neg.i64x2", name2: "Neg_I64x2", opPfx: uint8(0xFD), op: uint8(0xC1), t: wt.V128, a: wt.V128},
	{i: Alltrue_I64x2, class: Vector, name: "alltrue.i64x2", name2: "Alltrue_I64x2", opPfx: uint8(0xFD), op: uint8(0xC3), t: wt.I32, a: wt.V128},
	{i: Bitmask_I64x2, class: Vector, name: "bitmask.i64x2", name2: "Bitmask_I64x2", opPfx: uint8(0xFD), op: uint8(0xC4), t: wt.I32, a: wt.V128},
	{i: ExtendLowI32x4S_I64x2, class: Vector, name: "extend_low_i32x4_s.i64x2", name2: "ExtendLowI32x4S_I64x2", opPfx: uint8(0xFD), op: uint8(0xC7), t: wt.V128, a: wt.V128},
	{i: ExtendHighI32x4S_I64x2, class: Vector, name: "extend_high_i32x4_s.i64x2", name2: "ExtendHighI32x4S_I64x2", opPfx: uint8(0xFD), op: uint8(0xC8), t: wt.V128, a: wt.V128},
	{i: ExtendLowI32x4U_I64x2, class: Vector, name: "extend_low_i32x4_u.i64x2", name2: "ExtendLowI32x4U_I64x2", opPfx: uint8(0xFD), op: uint8(0xC9), t: wt.V128, a: wt.V128},
	{i: ExtendHighI32x4U_I64x2, class: Vector, name: "extend_high_i32x4_u.i64x2", name2: "ExtendHighI32x4U_I64x2", opPfx: uint8(0xFD), op: uint8(0xCA), t: wt.V128, a: wt.V128},
	{i: Shl_I64x2, class: Vector, name: "shl.i64x2", name2: "Shl_I64x2", opPfx: uint8(0xFD), op: uint8(0xCB), t: wt.V128, a: wt.V128, b: wt.I32},
	{i: ShrS_I64x2, class: Vector, name: "shr_s.i64x2", name2: "ShrS_I64x2", opPfx: uint8(0xFD), op: uint8(0xCC), t: wt.V128, a: wt.V128, b: wt.I32},
	{i: ShrU_I64x2, class: Vector, name: "shr_u.i64x2", name2: "ShrU_I64x2", opPfx: uint8(0xFD), op: uint8(0xCD), t: wt.V128, a: wt.V128, b: wt.I32},
	{i: Add_I64x2, class: Vector, name: "add.i64x2", name2: "Add_I64x2", opPfx: uint8(0xFD), op: uint8(0xCE), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Sub_I64x2, class: Vector, name: "sub.i64x2", name2: "Sub_I64x2", opPfx: uint8(0xFD), op: uint8(0xD1), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Mul_I64x2, class: Vector, name: "mul.i64x2", name2: "Mul_I64x2", opPfx: uint8(0xFD), op: uint8(0xD5), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: ExtmulLowI32x4S_I64x2, class: Vector, name: "extmul_low_i32x4_s.i64x2", name2: "ExtmulLowI32x4S_I64x2", opPfx: uint8(0xFD), op: uint8(0xDC), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: ExtmulHighI32x4S_I64x2, class: Vector, name: "extmul_high_i32x4_s.i64x2", name2: "ExtmulHighI32x4S_I64x2", opPfx: uint8(0xFD), op: uint8(0xDD), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: ExtmulLowI32x4U_I64x2, class: Vector, name: "extmul_low_i32x4_u.i64x2", name2: "ExtmulLowI32x4U_I64x2", opPfx: uint8(0xFD), op: uint8(0xDE), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: ExtmulHighI32x4U_I64x2, class: Vector, name: "extmul_high_i32x4_u.i64x2", name2: "ExtmulHighI32x4U_I64x2", opPfx: uint8(0xFD), op: uint8(0xDF), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Ceil_F32x4, class: Vector, name: "ceil.f32x4", name2: "Ceil_F32x4", opPfx: uint8(0xFD), op: uint8(0x67), t: wt.V128, a: wt.V128},
	{i: Floor_F32x4, class: Vector, name: "floor.f32x4", name2: "Floor_F32x4", opPfx: uint8(0xFD), op: uint8(0x68), t: wt.V128, a: wt.V128},
	{i: Trunc_F32x4, class: Vector, name: "trunc.f32x4", name2: "Trunc_F32x4", opPfx: uint8(0xFD), op: uint8(0x69), t: wt.V128, a: wt.V128},
	{i: Nearest_F32x4, class: Vector, name: "nearest.f32x4", name2: "Nearest_F32x4", opPfx: uint8(0xFD), op: uint8(0x6A), t: wt.V128, a: wt.V128},
	{i: Abs_F32x4, class: Vector, name: "abs.f32x4", name2: "Abs_F32x4", opPfx: uint8(0xFD), op: uint8(0xE0), t: wt.V128, a: wt.V128},
	{i: Neg_F32x4, class: Vector, name: "neg.f32x4", name2: "Neg_F32x4", opPfx: uint8(0xFD), op: uint8(0xE1), t: wt.V128, a: wt.V128},
	{i: Sqrt_F32x4, class: Vector, name: "sqrt.f32x4", name2: "Sqrt_F32x4", opPfx: uint8(0xFD), op: uint8(0xE3), t: wt.V128, a: wt.V128},
	{i: Add_F32x4, class: Vector, name: "add.f32x4", name2: "Add_F32x4", opPfx: uint8(0xFD), op: uint8(0xE4), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Sub_F32x4, class: Vector, name: "sub.f32x4", name2: "Sub_F32x4", opPfx: uint8(0xFD), op: uint8(0xE5), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Mul_F32x4, class: Vector, name: "mul.f32x4", name2: "Mul_F32x4", opPfx: uint8(0xFD), op: uint8(0xE6), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Div_F32x4, class: Vector, name: "div.f32x4", name2: "Div_F32x4", opPfx: uint8(0xFD), op: uint8(0xE7), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Min_F32x4, class: Vector, name: "min.f32x4", name2: "Min_F32x4", opPfx: uint8(0xFD), op: uint8(0xE8), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Max_F32x4, class: Vector, name: "max.f32x4", name2: "Max_F32x4", opPfx: uint8(0xFD), op: uint8(0xE9), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Pmin_F32x4, class: Vector, name: "pmin.f32x4", name2: "Pmin_F32x4", opPfx: uint8(0xFD), op: uint8(0xEA), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Pmax_F32x4, class: Vector, name: "pmax.f32x4", name2: "Pmax_F32x4", opPfx: uint8(0xFD), op: uint8(0xEB), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Ceil_F64x2, class: Vector, name: "ceil.f64x2", name2: "Ceil_F64x2", opPfx: uint8(0xFD), op: uint8(0x74), t: wt.V128, a: wt.V128},
	{i: Floor_F64x2, class: Vector, name: "floor.f64x2", name2: "Floor_F64x2", opPfx: uint8(0xFD), op: uint8(0x75), t: wt.V128, a: wt.V128},
	{i: Trunc_F64x2, class: Vector, name: "trunc.f64x2", name2: "Trunc_F64x2", opPfx: uint8(0xFD), op: uint8(0x7A), t: wt.V128, a: wt.V128},
	{i: Nearest_F64x2, class: Vector, name: "nearest.f64x2", name2: "Nearest_F64x2", opPfx: uint8(0xFD), op: uint8(0x94), t: wt.V128, a: wt.V128},
	{i: Abs_F64x2, class: Vector, name: "abs.f64x2", name2: "Abs_F64x2", opPfx: uint8(0xFD), op: uint8(0xEC), t: wt.V128, a: wt.V128},
	{i: Neg_F64x2, class: Vector, name: "neg.f64x2", name2: "Neg_F64x2", opPfx: uint8(0xFD), op: uint8(0xED), t: wt.V128, a: wt.V128},
	{i: Sqrt_F64x2, class: Vector, name: "sqrt.f64x2", name2: "Sqrt_F64x2", opPfx: uint8(0xFD), op: uint8(0xEF), t: wt.V128, a: wt.V128},
	{i: Add_F64x2, class: Vector, name: "add.f64x2", name2: "Add_F64x2", opPfx: uint8(0xFD), op: uint8(0xF0), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Sub_F64x2, class: Vector, name: "sub.f64x2", name2: "Sub_F64x2", opPfx: uint8(0xFD), op: uint8(0xF1), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Mul_F64x2, class: Vector, name: "mul.f64x2", name2: "Mul_F64x2", opPfx: uint8(0xFD), op: uint8(0xF2), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Div_F64x2, class: Vector, name: "div.f64x2", name2: "Div_F64x2", opPfx: uint8(0xFD), op: uint8(0xF3), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Min_F64x2, class: Vector, name: "min.f64x2", name2: "Min_F64x2", opPfx: uint8(0xFD), op: uint8(0xF4), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Max_F64x2, class: Vector, name: "max.f64x2", name2: "Max_F64x2", opPfx: uint8(0xFD), op: uint8(0xF5), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Pmin_F64x2, class: Vector, name: "pmin.f64x2", name2: "Pmin_F64x2", opPfx: uint8(0xFD), op: uint8(0xF6), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: Pmax_F64x2, class: Vector, name: "pmax.f64x2", name2: "Pmax_F64x2", opPfx: uint8(0xFD), op: uint8(0xF7), t: wt.V128, a: wt.V128, b: wt.V128},
	{i: TruncSatF32x4S_I32x4, class: Vector, name: "trunc_sat_f32x4_s.i32x4", name2: "TruncSatF32x4S_I32x4", opPfx: uint8(0xFD), op: uint8(0xF8), t: wt.V128, a: wt.V128},
	{i: TruncSatF32x4U_I32x4, class: Vector, name: "trunc_sat_f32x4_u.i32x4", name2: "TruncSatF32x4U_I32x4", opPfx: uint8(0xFD), op: uint8(0xF9), t: wt.V128, a: wt.V128},
	{i: ConvertI32x4S_F32x4, class: Vector, name: "convert_i32x4_s.f32x4", name2: "ConvertI32x4S_F32x4", opPfx: uint8(0xFD), op: uint8(0xFA), t: wt.V128, a: wt.V128},
	{i: ConvertI32x4U_F32x4, class: Vector, name: "convert_i32x4_u.f32x4", name2: "ConvertI32x4U_F32x4", opPfx: uint8(0xFD), op: uint8(0xFB), t: wt.V128, a: wt.V128},
	{i: TruncSatF64x2SZero_I32x4, class: Vector, name: "trunc_sat_f64x2_s_zero.i32x4", name2: "TruncSatF64x2SZero_I32x4", opPfx: uint8(0xFD), op: uint8(0xFC), t: wt.V128, a: wt.V128},
	{i: TruncSatF64x2UZero_I32x4, class: Vector, name: "trunc_sat_f64x2_u_zero.i32x4", name2: "TruncSatF64x2UZero_I32x4", opPfx: uint8(0xFD), op: uint8(0xFD), t: wt.V128, a: wt.V128},
	{i: ConvertLowI32x4S_F64x2, class: Vector, name: "convert_low_i32x4_s.f64x2", name2: "ConvertLowI32x4S_F64x2", opPfx: uint8(0xFD), op: uint8(0xFE), t: wt.V128, a: wt.V128},
	{i: ConvertLowI32x4U_F64x2, class: Vector, name: "convert_low_i32x4_u.f64x2", name2: "ConvertLowI32x4U_F64x2", opPfx: uint8(0xFD), op: uint8(0xFF), t: wt.V128, a: wt.V128},
	{i: DemoteF64x2Zero_F32x4, class: Vector, name: "demote_f64x2_zero.f32x4", name2: "DemoteF64x2Zero_F32x4", opPfx: uint8(0xFD), op: uint8(0x5E), t: wt.V128, a: wt.V128},
	{i: PromoteLowF32x4_F64x2, class: Vector, name: "promote_low_f32x4.f64x2", name2: "PromoteLowF32x4_F64x2", opPfx: uint8(0xFD), op: uint8(0x5F), t: wt.V128, a: wt.V128},
}

var opMap = [256]*[256]Instr{
	0x00: {
		0x00: Unreachable,
		0x01: Nop,
		0x02: Block,
		0x03: Loop,
		0x04: If,
		0x05: Else,
		0x0B: End,
		0x0C: Br,
		0x0D: BrIf,
		0x0E: BrTable,
		0x0F: Return,
		0x10: Call,
		0x11: CallIndirect,
		0x1A: Drop,
		0x1B: Select,
		0x1C: SelectType,
		0x20: LocalGet,
		0x21: LocalSet,
		0x22: LocalTee,
		0x23: GlobalGet,
		0x24: GlobalSet,
		0x25: TableGet,
		0x26: TableSet,
		0x28: Load_I32,
		0x29: Load_I64,
		0x2A: Load_F32,
		0x2B: Load_F64,
		0x2C: Load8S_I32,
		0x2D: Load8U_I32,
		0x2E: Load16S_I32,
		0x2F: Load16U_I32,
		0x30: Load8S_I64,
		0x31: Load8U_I64,
		0x32: Load16S_I64,
		0x33: Load16U_I64,
		0x34: Load32S_I64,
		0x35: Load32U_I64,
		0x36: Store_I32,
		0x37: Store_I64,
		0x38: Store_F32,
		0x39: Store_F64,
		0x3A: Store8_I32,
		0x3B: Store16_I32,
		0x3C: Store8_I64,
		0x3D: Store16_I64,
		0x3E: Store32_I64,
		0x3F: MemorySize,
		0x40: MemoryGrow,
		0x41: Const_I32,
		0x42: Const_I64,
		0x43: Const_F32,
		0x44: Const_F64,
		0x45: Eqz_I32,
		0x46: Eq_I32,
		0x47: Ne_I32,
		0x48: LtS_I32,
		0x49: LtU_I32,
		0x4A: GtS_I32,
		0x4B: GtU_I32,
		0x4C: LeS_I32,
		0x4D: LeU_I32,
		0x4E: GeS_I32,
		0x4F: GeU_I32,
		0x50: Eqz_I64,
		0x51: Eq_I64,
		0x52: Ne_I64,
		0x53: LtS_I64,
		0x54: LtU_I64,
		0x55: GtS_I64,
		0x56: GtU_I64,
		0x57: LeS_I64,
		0x58: LeU_I64,
		0x59: GeS_I64,
		0x5A: GeU_I64,
		0x5B: Eq_F32,
		0x5C: Ne_F32,
		0x5D: Lt_F32,
		0x5E: Gt_F32,
		0x5F: Le_F32,
		0x60: Ge_F32,
		0x61: Eq_F64,
		0x62: Ne_F64,
		0x63: Lt_F64,
		0x64: Gt_F64,
		0x65: Le_F64,
		0x66: Ge_F64,
		0x67: Clz_I32,
		0x68: Ctz_I32,
		0x69: Popcnt_I32,
		0x6A: Add_I32,
		0x6B: Sub_I32,
		0x6C: Mul_I32,
		0x6D: DivS_I32,
		0x6E: DivU_I32,
		0x6F: RemS_I32,
		0x70: RemU_I32,
		0x71: And_I32,
		0x72: Or_I32,
		0x73: Xor_I32,
		0x74: Shl_I32,
		0x75: ShrS_I32,
		0x76: ShrU_I32,
		0x77: RotL_I32,
		0x78: RotR_I32,
		0x79: Clz_I64,
		0x7A: Ctz_I64,
		0x7B: Popcnt_I64,
		0x7C: Add_I64,
		0x7D: Sub_I64,
		0x7E: Mul_I64,
		0x7F: DivS_I64,
		0x80: DivU_I64,
		0x81: RemS_I64,
		0x82: RemU_I64,
		0x83: And_I64,
		0x84: Or_I64,
		0x85: Xor_I64,
		0x86: Shl_I64,
		0x87: ShrS_I64,
		0x88: ShrU_I64,
		0x89: RotL_I64,
		0x8A: RotR_I64,
		0x8B: Abs_F32,
		0x8C: Neg_F32,
		0x8D: Ceil_F32,
		0x8E: Floor_F32,
		0x8F: Trunc_F32,
		0x90: Nearest_F32,
		0x91: Sqrt_F32,
		0x92: Add_F32,
		0x93: Sub_F32,
		0x94: Mul_F32,
		0x95: Div_F32,
		0x96: Min_F32,
		0x97: Max_F32,
		0x98: Copysign_F32,
		0x99: Abs_F64,
		0x9A: Neg_F64,
		0x9B: Ceil_F64,
		0x9C: Floor_F64,
		0x9D: Trunc_F64,
		0x9E: Nearest_F64,
		0x9F: Sqrt_F64,
		0xA0: Add_F64,
		0xA1: Sub_F64,
		0xA2: Mul_F64,
		0xA3: Div_F64,
		0xA4: Min_F64,
		0xA5: Max_F64,
		0xA6: Copysign_F64,
		0xA7: WrapI64_I32,
		0xA8: TruncF32S_I32,
		0xA9: TruncF32U_I32,
		0xAA: TruncF64S_I32,
		0xAB: TruncF64U_I32,
		0xAC: ExtendI32S_I64,
		0xAD: ExtendI32U_I64,
		0xAE: TruncF32S_I64,
		0xAF: TruncF32U_I64,
		0xB0: TruncF64S_I64,
		0xB1: TruncF64U_I64,
		0xB2: ConvertI32S_F32,
		0xB3: ConvertI32U_F32,
		0xB4: ConvertI64S_F32,
		0xB5: ConvertI64U_F32,
		0xB6: DemoteF64_F32,
		0xB7: ConvertI32S_F64,
		0xB8: ConvertI32U_F64,
		0xB9: ConvertI64S_F64,
		0xBA: ConvertI64U_F64,
		0xBB: PromoteF32_F64,
		0xBC: ReinterpretF32_I32,
		0xBD: ReinterpretF64_I64,
		0xBE: ReinterpretI32_F32,
		0xBF: ReinterpretI64_F64,
		0xC0: Extend8S_I32,
		0xC1: Extend16S_I32,
		0xC2: Extend8S_I64,
		0xC3: Extend16S_I64,
		0xC4: Extend32S_I64,
	},
	0xFC: {
		0x00: TruncSatF32S_I32,
		0x01: TruncSatF32U_I32,
		0x02: TruncSatF64S_I32,
		0x03: TruncSatF64U_I32,
		0x04: TruncSatF32S_I64,
		0x05: TruncSatF32U_I64,
		0x06: TruncSatF64S_I64,
		0x07: TruncSatF64U_I64,
		0x08: MemoryInit,
		0x09: DataDrop,
		0x0A: MemoryCopy,
		0x0B: MemoryFill,
		0x0C: TableInit,
		0x0D: ElemDrop,
		0x0E: TableCopy,
		0x0F: TableGrow,
		0x10: TableSize,
		0x11: TableFill,
	},
	0xFD: {
		0x00: Load_V128,
		0x01: Load8x8S_V128,
		0x02: Load8x8U_V128,
		0x03: Load16x4S_V128,
		0x04: Load16x4U_V128,
		0x05: Load32x2S_V128,
		0x06: Load32x2U_V128,
		0x07: Load8Splat_V128,
		0x08: Load16Splat_V128,
		0x09: Load32Splat_V128,
		0x0A: Load64Splat_V128,
		0x0B: Store_V128,
		0x0C: Const_V128,
		0x0D: Shuffle_I8x16,
		0x0E: Swizzle_I8x16,
		0x0F: Splat_I8x16,
		0x10: Splat_I16x8,
		0x11: Splat_I32x4,
		0x12: Splat_I64x2,
		0x13: Splat_F32x4,
		0x14: Splat_F64x2,
		0x15: ExtractLaneS_I8x16,
		0x16: ExtractLaneU_I8x16,
		0x17: ReplaceLane_I8x16,
		0x18: ExtractLaneS_I16x8,
		0x19: ExtractLaneU_I16x8,
		0x1A: ReplaceLane_I16x8,
		0x1B: ExtractLane_I32x4,
		0x1C: ReplaceLane_I32x4,
		0x1D: ExtractLane_I64x2,
		0x1E: ReplaceLane_I64x2,
		0x1F: ExtractLane_F32x4,
		0x20: ReplaceLane_F32x4,
		0x21: ExtractLane_F64x2,
		0x22: ReplaceLane_F64x2,
		0x23: Eq_I8x16,
		0x24: Ne_I8x16,
		0x25: LtS_I8x16,
		0x26: LtU_I8x16,
		0x27: GtS_I8x16,
		0x28: GtU_I8x16,
		0x29: LeS_I8x16,
		0x2A: LeU_I8x16,
		0x2B: GeS_I8x16,
		0x2C: GeU_I8x16,
		0x2D: Eq_I16x8,
		0x2E: Ne_I16x8,
		0x2F: LtS_I16x8,
		0x30: LtU_I16x8,
		0x31: GtS_I16x8,
		0x32: GtU_I16x8,
		0x33: LeS_I16x8,
		0x34: LeU_I16x8,
		0x35: GeS_I16x8,
		0x36: GeU_I16x8,
		0x37: Eq_I32x4,
		0x38: Ne_I32x4,
		0x39: LtS_I32x4,
		0x3A: LtU_I32x4,
		0x3B: GtS_I32x4,
		0x3C: GtU_I32x4,
		0x3D: LeS_I32x4,
		0x3E: LeU_I32x4,
		0x3F: GeS_I32x4,
		0x40: GeU_I32x4,
		0x41: Eq_F32x4,
		0x42: Ne_F32x4,
		0x43: Lt_F32x4,
		0x44: Gt_F32x4,
		0x45: Le_F32x4,
		0x46: Ge_F32x4,
		0x47: Eq_F64x2,
		0x48: Ne_F64x2,
		0x49: Lt_F64x2,
		0x4A: Gt_F64x2,
		0x4B: Le_F64x2,
		0x4C: Ge_F64x2,
		0x4D: Not_V128,
		0x4E: And_V128,
		0x4F: Andnot_V128,
		0x50: Or_V128,
		0x51: Xor_V128,
		0x52: Bitselect_V128,
		0x53: AnyTrue_V128,
		0x54: Load8Lane_V128,
		0x55: Load16Lane_V128,
		0x56: Load32Lane_V128,
		0x57: Load64Lane_V128,
		0x58: Store8Lane_V128,
		0x59: Store16Lane_V128,
		0x5A: Store32Lane_V128,
		0x5B: Store64Lane_V128,
		0x5C: Load32Zero_V128,
		0x5D: Load64Zero_V128,
		0x5E: DemoteF64x2Zero_F32x4,
		0x5F: PromoteLowF32x4_F64x2,
		0x60: Abs_I8x16,
		0x61: Neg_I8x16,
		0x62: Popcnt_I8x16,
		0x63: AllTrue_I8x16,
		0x64: Bitmask_I8x16,
		0x65: NarrowI16x8S_I8x16,
		0x66: NarrowI16x8U_I8x16,
		0x67: Ceil_F32x4,
		0x68: Floor_F32x4,
		0x69: Trunc_F32x4,
		0x6A: Nearest_F32x4,
		0x6B: Shl_I8x16,
		0x6C: ShrS_I8x16,
		0x6D: ShrU_I8x16,
		0x6E: Add_I8x16,
		0x6F: AddSatS_I8x16,
		0x70: AddSatU_I8x16,
		0x71: Sub_I8x16,
		0x72: SubSatS_I8x16,
		0x73: SubSatU_I8x16,
		0x74: Ceil_F64x2,
		0x75: Floor_F64x2,
		0x76: MinS_I8x16,
		0x77: MinU_I8x16,
		0x78: MaxS_I8x16,
		0x79: MaxU_I8x16,
		0x7A: Trunc_F64x2,
		0x7B: AvgrU_I8x16,
		0x7C: ExtaddPairwiseI8x16S_I16x8,
		0x7D: ExtaddPairwiseI8x16U_I16x8,
		0x7E: ExtaddPairwiseI16x8S_I32x4,
		0x7F: ExtaddPairwiseI16x8U_I32x4,
		0x80: Abs_I16x8,
		0x81: Neg_I16x8,
		0x82: Q15mulrSatS_I16x8,
		0x83: Alltrue_I16x8,
		0x84: Bitmask_I16x8,
		0x85: NarrowI32x4S_I16x8,
		0x86: NarrowI32x4U_I16x8,
		0x87: ExtendLowI8x16S_I16x8,
		0x88: ExtendHighI8x16S_I16x8,
		0x89: ExtendLowI8x16U_I16x8,
		0x8A: ExtendHighI8x16U_I16x8,
		0x8B: Shl_I16x8,
		0x8C: ShrS_I16x8,
		0x8D: ShrU_I16x8,
		0x8E: Add_I16x8,
		0x8F: AddSatS_I16x8,
		0x90: AddSatU_I16x8,
		0x91: Sub_I16x8,
		0x92: SubSatS_I16x8,
		0x93: SubSatU_I16x8,
		0x94: Nearest_F64x2,
		0x95: Mul_I16x8,
		0x96: MinS_I16x8,
		0x97: MinU_I16x8,
		0x98: MaxS_I16x8,
		0x99: MaxU_I16x8,
		0x9B: AvgrU_I16x8,
		0x9C: ExtmulLowI8x16S_I16x8,
		0x9D: ExtmulHighI8x16S_I16x8,
		0x9E: ExtmulLowI8x16U_I16x8,
		0x9F: ExtmulHighI8x16U_I16x8,
		0xA0: Abs_I32x4,
		0xA1: Neg_I32x4,
		0xA3: Alltrue_I32x4,
		0xA4: Bitmask_I32x4,
		0xA7: ExtendLowI16x8S_I32x4,
		0xA8: ExtendHighI16x8S_I32x4,
		0xA9: ExtendLowI16x8U_I32x4,
		0xAA: ExtendHighI16x8U_I32x4,
		0xAB: Shl_I32x4,
		0xAC: ShrS_I32x4,
		0xAD: ShrU_I32x4,
		0xAE: Add_I32x4,
		0xB1: Sub_I32x4,
		0xB5: Mul_I32x4,
		0xB6: MinS_I32x4,
		0xB7: MinU_I32x4,
		0xB8: MaxS_I32x4,
		0xB9: MaxU_I32x4,
		0xBA: DotI16x8S_I32x4,
		0xBC: ExtmulLowI16x8S_I32x4,
		0xBD: ExtmulHighI16x8S_I32x4,
		0xBE: ExtmulLowI16x8U_I32x4,
		0xBF: ExtmulHighI16x8U_I32x4,
		0xC0: Abs_I64x2,
		0xC1: Neg_I64x2,
		0xC3: Alltrue_I64x2,
		0xC4: Bitmask_I64x2,
		0xC7: ExtendLowI32x4S_I64x2,
		0xC8: ExtendHighI32x4S_I64x2,
		0xC9: ExtendLowI32x4U_I64x2,
		0xCA: ExtendHighI32x4U_I64x2,
		0xCB: Shl_I64x2,
		0xCC: ShrS_I64x2,
		0xCD: ShrU_I64x2,
		0xCE: Add_I64x2,
		0xD1: Sub_I64x2,
		0xD5: Mul_I64x2,
		0xD6: Eq_I64x2,
		0xD7: Ne_I64x2,
		0xD8: LtS_I64x2,
		0xD9: GtS_I64x2,
		0xDA: LeS_I64x2,
		0xDB: GeS_I64x2,
		0xDC: ExtmulLowI32x4S_I64x2,
		0xDD: ExtmulHighI32x4S_I64x2,
		0xDE: ExtmulLowI32x4U_I64x2,
		0xDF: ExtmulHighI32x4U_I64x2,
		0xE0: Abs_F32x4,
		0xE1: Neg_F32x4,
		0xE3: Sqrt_F32x4,
		0xE4: Add_F32x4,
		0xE5: Sub_F32x4,
		0xE6: Mul_F32x4,
		0xE7: Div_F32x4,
		0xE8: Min_F32x4,
		0xE9: Max_F32x4,
		0xEA: Pmin_F32x4,
		0xEB: Pmax_F32x4,
		0xEC: Abs_F64x2,
		0xED: Neg_F64x2,
		0xEF: Sqrt_F64x2,
		0xF0: Add_F64x2,
		0xF1: Sub_F64x2,
		0xF2: Mul_F64x2,
		0xF3: Div_F64x2,
		0xF4: Min_F64x2,
		0xF5: Max_F64x2,
		0xF6: Pmin_F64x2,
		0xF7: Pmax_F64x2,
		0xF8: TruncSatF32x4S_I32x4,
		0xF9: TruncSatF32x4U_I32x4,
		0xFA: ConvertI32x4S_F32x4,
		0xFB: ConvertI32x4U_F32x4,
		0xFC: TruncSatF64x2SZero_I32x4,
		0xFD: TruncSatF64x2UZero_I32x4,
		0xFE: ConvertLowI32x4S_F64x2,
		0xFF: ConvertLowI32x4U_F64x2,
	},
}

var nameMap = map[string]Instr{
	"unreachable":                   Unreachable,
	"nop":                           Nop,
	"block":                         Block,
	"loop":                          Loop,
	"if":                            If,
	"else":                          Else,
	"end":                           End,
	"br":                            Br,
	"br_if":                         BrIf,
	"br_table":                      BrTable,
	"return":                        Return,
	"call":                          Call,
	"call_indirect":                 CallIndirect,
	"drop":                          Drop,
	"select":                        Select,
	"select_type":                   SelectType,
	"local_get":                     LocalGet,
	"local_set":                     LocalSet,
	"local_tee":                     LocalTee,
	"global_get":                    GlobalGet,
	"global_set":                    GlobalSet,
	"table_get":                     TableGet,
	"table_set":                     TableSet,
	"table_init":                    TableInit,
	"elem_drop":                     ElemDrop,
	"table_copy":                    TableCopy,
	"table_grow":                    TableGrow,
	"table_size":                    TableSize,
	"table_fill":                    TableFill,
	"load.i32":                      Load_I32,
	"load.i64":                      Load_I64,
	"load.f32":                      Load_F32,
	"load.f64":                      Load_F64,
	"load8_s.i32":                   Load8S_I32,
	"load8_u.i32":                   Load8U_I32,
	"load16_s.i32":                  Load16S_I32,
	"load16_u.i32":                  Load16U_I32,
	"load8_s.i64":                   Load8S_I64,
	"load8_u.i64":                   Load8U_I64,
	"load16_s.i64":                  Load16S_I64,
	"load16_u.i64":                  Load16U_I64,
	"load32_s.i64":                  Load32S_I64,
	"load32_u.i64":                  Load32U_I64,
	"store.i32":                     Store_I32,
	"store.i64":                     Store_I64,
	"store.f32":                     Store_F32,
	"store.f64":                     Store_F64,
	"store8.i32":                    Store8_I32,
	"store16.i32":                   Store16_I32,
	"store8.i64":                    Store8_I64,
	"store16.i64":                   Store16_I64,
	"store32.i64":                   Store32_I64,
	"memory_size":                   MemorySize,
	"memory_grow":                   MemoryGrow,
	"memory_init":                   MemoryInit,
	"data_drop":                     DataDrop,
	"memory_copy":                   MemoryCopy,
	"memory_fill":                   MemoryFill,
	"const.i32":                     Const_I32,
	"const.i64":                     Const_I64,
	"const.f32":                     Const_F32,
	"const.f64":                     Const_F64,
	"eqz.i32":                       Eqz_I32,
	"eq.i32":                        Eq_I32,
	"ne.i32":                        Ne_I32,
	"lt_s.i32":                      LtS_I32,
	"lt_u.i32":                      LtU_I32,
	"gt_s.i32":                      GtS_I32,
	"gt_u.i32":                      GtU_I32,
	"le_s.i32":                      LeS_I32,
	"le_u.i32":                      LeU_I32,
	"ge_s.i32":                      GeS_I32,
	"ge_u.i32":                      GeU_I32,
	"eqz.i64":                       Eqz_I64,
	"eq.i64":                        Eq_I64,
	"ne.i64":                        Ne_I64,
	"lt_s.i64":                      LtS_I64,
	"lt_u.i64":                      LtU_I64,
	"gt_s.i64":                      GtS_I64,
	"gt_u.i64":                      GtU_I64,
	"le_s.i64":                      LeS_I64,
	"le_u.i64":                      LeU_I64,
	"ge_s.i64":                      GeS_I64,
	"ge_u.i64":                      GeU_I64,
	"eq.f32":                        Eq_F32,
	"ne.f32":                        Ne_F32,
	"lt.f32":                        Lt_F32,
	"gt.f32":                        Gt_F32,
	"le.f32":                        Le_F32,
	"ge.f32":                        Ge_F32,
	"eq.f64":                        Eq_F64,
	"ne.f64":                        Ne_F64,
	"lt.f64":                        Lt_F64,
	"gt.f64":                        Gt_F64,
	"le.f64":                        Le_F64,
	"ge.f64":                        Ge_F64,
	"clz.i32":                       Clz_I32,
	"ctz.i32":                       Ctz_I32,
	"popcnt.i32":                    Popcnt_I32,
	"add.i32":                       Add_I32,
	"sub.i32":                       Sub_I32,
	"mul.i32":                       Mul_I32,
	"div_s.i32":                     DivS_I32,
	"div_u.i32":                     DivU_I32,
	"rem_s.i32":                     RemS_I32,
	"rem_u.i32":                     RemU_I32,
	"and.i32":                       And_I32,
	"or.i32":                        Or_I32,
	"xor.i32":                       Xor_I32,
	"shl.i32":                       Shl_I32,
	"shr_s.i32":                     ShrS_I32,
	"shr_u.i32":                     ShrU_I32,
	"rot_l.i32":                     RotL_I32,
	"rot_r.i32":                     RotR_I32,
	"clz.i64":                       Clz_I64,
	"ctz.i64":                       Ctz_I64,
	"popcnt.i64":                    Popcnt_I64,
	"add.i64":                       Add_I64,
	"sub.i64":                       Sub_I64,
	"mul.i64":                       Mul_I64,
	"div_s.i64":                     DivS_I64,
	"div_u.i64":                     DivU_I64,
	"rem_s.i64":                     RemS_I64,
	"rem_u.i64":                     RemU_I64,
	"and.i64":                       And_I64,
	"or.i64":                        Or_I64,
	"xor.i64":                       Xor_I64,
	"shl.i64":                       Shl_I64,
	"shr_s.i64":                     ShrS_I64,
	"shr_u.i64":                     ShrU_I64,
	"rot_l.i64":                     RotL_I64,
	"rot_r.i64":                     RotR_I64,
	"abs.f32":                       Abs_F32,
	"neg.f32":                       Neg_F32,
	"ceil.f32":                      Ceil_F32,
	"floor.f32":                     Floor_F32,
	"trunc.f32":                     Trunc_F32,
	"nearest.f32":                   Nearest_F32,
	"sqrt.f32":                      Sqrt_F32,
	"add.f32":                       Add_F32,
	"sub.f32":                       Sub_F32,
	"mul.f32":                       Mul_F32,
	"div.f32":                       Div_F32,
	"min.f32":                       Min_F32,
	"max.f32":                       Max_F32,
	"copysign.f32":                  Copysign_F32,
	"abs.f64":                       Abs_F64,
	"neg.f64":                       Neg_F64,
	"ceil.f64":                      Ceil_F64,
	"floor.f64":                     Floor_F64,
	"trunc.f64":                     Trunc_F64,
	"nearest.f64":                   Nearest_F64,
	"sqrt.f64":                      Sqrt_F64,
	"add.f64":                       Add_F64,
	"sub.f64":                       Sub_F64,
	"mul.f64":                       Mul_F64,
	"div.f64":                       Div_F64,
	"min.f64":                       Min_F64,
	"max.f64":                       Max_F64,
	"copysign.f64":                  Copysign_F64,
	"wrap_i64.i32":                  WrapI64_I32,
	"trunc_f32_s.i32":               TruncF32S_I32,
	"trunc_f32_u.i32":               TruncF32U_I32,
	"trunc_f64_s.i32":               TruncF64S_I32,
	"trunc_f64_u.i32":               TruncF64U_I32,
	"extend_i32_s.i64":              ExtendI32S_I64,
	"extend_i32_u.i64":              ExtendI32U_I64,
	"trunc_f32_s.i64":               TruncF32S_I64,
	"trunc_f32_u.i64":               TruncF32U_I64,
	"trunc_f64_s.i64":               TruncF64S_I64,
	"trunc_f64_u.i64":               TruncF64U_I64,
	"convert_i32_s.f32":             ConvertI32S_F32,
	"convert_i32_u.f32":             ConvertI32U_F32,
	"convert_i64_s.f32":             ConvertI64S_F32,
	"convert_i64_u.f32":             ConvertI64U_F32,
	"demote_f64.f32":                DemoteF64_F32,
	"convert_i32_s.f64":             ConvertI32S_F64,
	"convert_i32_u.f64":             ConvertI32U_F64,
	"convert_i64_s.f64":             ConvertI64S_F64,
	"convert_i64_u.f64":             ConvertI64U_F64,
	"promote_f32.f64":               PromoteF32_F64,
	"reinterpret_f32.i32":           ReinterpretF32_I32,
	"reinterpret_f64.i64":           ReinterpretF64_I64,
	"reinterpret_i32.f32":           ReinterpretI32_F32,
	"reinterpret_i64.f64":           ReinterpretI64_F64,
	"extend8_s.i32":                 Extend8S_I32,
	"extend16_s.i32":                Extend16S_I32,
	"extend8_s.i64":                 Extend8S_I64,
	"extend16_s.i64":                Extend16S_I64,
	"extend32_s.i64":                Extend32S_I64,
	"trunc_sat_f32_s.i32":           TruncSatF32S_I32,
	"trunc_sat_f32_u.i32":           TruncSatF32U_I32,
	"trunc_sat_f64_s.i32":           TruncSatF64S_I32,
	"trunc_sat_f64_u.i32":           TruncSatF64U_I32,
	"trunc_sat_f32_s.i64":           TruncSatF32S_I64,
	"trunc_sat_f32_u.i64":           TruncSatF32U_I64,
	"trunc_sat_f64_s.i64":           TruncSatF64S_I64,
	"trunc_sat_f64_u.i64":           TruncSatF64U_I64,
	"load.v128":                     Load_V128,
	"load8x8_s.v128":                Load8x8S_V128,
	"load8x8_u.v128":                Load8x8U_V128,
	"load16x4_s.v128":               Load16x4S_V128,
	"load16x4_u.v128":               Load16x4U_V128,
	"load32x2_s.v128":               Load32x2S_V128,
	"load32x2_u.v128":               Load32x2U_V128,
	"load8_splat.v128":              Load8Splat_V128,
	"load16_splat.v128":             Load16Splat_V128,
	"load32_splat.v128":             Load32Splat_V128,
	"load64_splat.v128":             Load64Splat_V128,
	"load32_zero.v128":              Load32Zero_V128,
	"load64_zero.v128":              Load64Zero_V128,
	"store.v128":                    Store_V128,
	"load8_lane.v128":               Load8Lane_V128,
	"load16_lane.v128":              Load16Lane_V128,
	"load32_lane.v128":              Load32Lane_V128,
	"load64_lane.v128":              Load64Lane_V128,
	"store8_lane.v128":              Store8Lane_V128,
	"store16_lane.v128":             Store16Lane_V128,
	"store32_lane.v128":             Store32Lane_V128,
	"store64_lane.v128":             Store64Lane_V128,
	"const.v128":                    Const_V128,
	"shuffle.i8x16":                 Shuffle_I8x16,
	"extract_lane_s.i8x16":          ExtractLaneS_I8x16,
	"extract_lane_u.i8x16":          ExtractLaneU_I8x16,
	"replace_lane.i8x16":            ReplaceLane_I8x16,
	"extract_lane_s.i16x8":          ExtractLaneS_I16x8,
	"extract_lane_u.i16x8":          ExtractLaneU_I16x8,
	"replace_lane.i16x8":            ReplaceLane_I16x8,
	"extract_lane.i32x4":            ExtractLane_I32x4,
	"replace_lane.i32x4":            ReplaceLane_I32x4,
	"extract_lane.i64x2":            ExtractLane_I64x2,
	"replace_lane.i64x2":            ReplaceLane_I64x2,
	"extract_lane.f32x4":            ExtractLane_F32x4,
	"replace_lane.f32x4":            ReplaceLane_F32x4,
	"extract_lane.f64x2":            ExtractLane_F64x2,
	"replace_lane.f64x2":            ReplaceLane_F64x2,
	"swizzle.i8x16":                 Swizzle_I8x16,
	"splat.i8x16":                   Splat_I8x16,
	"splat.i16x8":                   Splat_I16x8,
	"splat.i32x4":                   Splat_I32x4,
	"splat.i64x2":                   Splat_I64x2,
	"splat.f32x4":                   Splat_F32x4,
	"splat.f64x2":                   Splat_F64x2,
	"eq.i8x16":                      Eq_I8x16,
	"ne.i8x16":                      Ne_I8x16,
	"lt_s.i8x16":                    LtS_I8x16,
	"lt_u.i8x16":                    LtU_I8x16,
	"gt_s.i8x16":                    GtS_I8x16,
	"gt_u.i8x16":                    GtU_I8x16,
	"le_s.i8x16":                    LeS_I8x16,
	"le_u.i8x16":                    LeU_I8x16,
	"ge_s.i8x16":                    GeS_I8x16,
	"ge_u.i8x16":                    GeU_I8x16,
	"eq.i16x8":                      Eq_I16x8,
	"ne.i16x8":                      Ne_I16x8,
	"lt_s.i16x8":                    LtS_I16x8,
	"lt_u.i16x8":                    LtU_I16x8,
	"gt_s.i16x8":                    GtS_I16x8,
	"gt_u.i16x8":                    GtU_I16x8,
	"le_s.i16x8":                    LeS_I16x8,
	"le_u.i16x8":                    LeU_I16x8,
	"ge_s.i16x8":                    GeS_I16x8,
	"ge_u.i16x8":                    GeU_I16x8,
	"eq.i32x4":                      Eq_I32x4,
	"ne.i32x4":                      Ne_I32x4,
	"lt_s.i32x4":                    LtS_I32x4,
	"lt_u.i32x4":                    LtU_I32x4,
	"gt_s.i32x4":                    GtS_I32x4,
	"gt_u.i32x4":                    GtU_I32x4,
	"le_s.i32x4":                    LeS_I32x4,
	"le_u.i32x4":                    LeU_I32x4,
	"ge_s.i32x4":                    GeS_I32x4,
	"ge_u.i32x4":                    GeU_I32x4,
	"eq.i64x2":                      Eq_I64x2,
	"ne.i64x2":                      Ne_I64x2,
	"lt_s.i64x2":                    LtS_I64x2,
	"gt_s.i64x2":                    GtS_I64x2,
	"le_s.i64x2":                    LeS_I64x2,
	"ge_s.i64x2":                    GeS_I64x2,
	"eq.f32x4":                      Eq_F32x4,
	"ne.f32x4":                      Ne_F32x4,
	"lt.f32x4":                      Lt_F32x4,
	"gt.f32x4":                      Gt_F32x4,
	"le.f32x4":                      Le_F32x4,
	"ge.f32x4":                      Ge_F32x4,
	"eq.f64x2":                      Eq_F64x2,
	"ne.f64x2":                      Ne_F64x2,
	"lt.f64x2":                      Lt_F64x2,
	"gt.f64x2":                      Gt_F64x2,
	"le.f64x2":                      Le_F64x2,
	"ge.f64x2":                      Ge_F64x2,
	"not.v128":                      Not_V128,
	"and.v128":                      And_V128,
	"andnot.v128":                   Andnot_V128,
	"or.v128":                       Or_V128,
	"xor.v128":                      Xor_V128,
	"bitselect.v128":                Bitselect_V128,
	"any_true.v128":                 AnyTrue_V128,
	"abs.i8x16":                     Abs_I8x16,
	"neg.i8x16":                     Neg_I8x16,
	"popcnt.i8x16":                  Popcnt_I8x16,
	"all_true.i8x16":                AllTrue_I8x16,
	"bitmask.i8x16":                 Bitmask_I8x16,
	"narrow_i16x8_s.i8x16":          NarrowI16x8S_I8x16,
	"narrow_i16x8_u.i8x16":          NarrowI16x8U_I8x16,
	"shl.i8x16":                     Shl_I8x16,
	"shr_s.i8x16":                   ShrS_I8x16,
	"shr_u.i8x16":                   ShrU_I8x16,
	"add.i8x16":                     Add_I8x16,
	"add_sat_s.i8x16":               AddSatS_I8x16,
	"add_sat_u.i8x16":               AddSatU_I8x16,
	"sub.i8x16":                     Sub_I8x16,
	"sub_sat_s.i8x16":               SubSatS_I8x16,
	"sub_sat_u.i8x16":               SubSatU_I8x16,
	"min_s.i8x16":                   MinS_I8x16,
	"min_u.i8x16":                   MinU_I8x16,
	"max_s.i8x16":                   MaxS_I8x16,
	"max_u.i8x16":                   MaxU_I8x16,
	"avgr_u.i8x16":                  AvgrU_I8x16,
	"extadd_pairwise_i8x16_s.i16x8": ExtaddPairwiseI8x16S_I16x8,
	"extadd_pairwise_i8x16_u.i16x8": ExtaddPairwiseI8x16U_I16x8,
	"abs.i16x8":                     Abs_I16x8,
	"neg.i16x8":                     Neg_I16x8,
	"q15mulr_sat_s.i16x8":           Q15mulrSatS_I16x8,
	"alltrue.i16x8":                 Alltrue_I16x8,
	"bitmask.i16x8":                 Bitmask_I16x8,
	"narrow_i32x4_s.i16x8":          NarrowI32x4S_I16x8,
	"narrow_i32x4_u.i16x8":          NarrowI32x4U_I16x8,
	"extend_low_i8x16_s.i16x8":      ExtendLowI8x16S_I16x8,
	"extend_high_i8x16_s.i16x8":     ExtendHighI8x16S_I16x8,
	"extend_low_i8x16_u.i16x8":      ExtendLowI8x16U_I16x8,
	"extend_high_i8x16_u.i16x8":     ExtendHighI8x16U_I16x8,
	"shl.i16x8":                     Shl_I16x8,
	"shr_s.i16x8":                   ShrS_I16x8,
	"shr_u.i16x8":                   ShrU_I16x8,
	"add.i16x8":                     Add_I16x8,
	"add_sat_s.i16x8":               AddSatS_I16x8,
	"add_sat_u.i16x8":               AddSatU_I16x8,
	"sub.i16x8":                     Sub_I16x8,
	"sub_sat_s.i16x8":               SubSatS_I16x8,
	"sub_sat_u.i16x8":               SubSatU_I16x8,
	"mul.i16x8":                     Mul_I16x8,
	"min_s.i16x8":                   MinS_I16x8,
	"min_u.i16x8":                   MinU_I16x8,
	"max_s.i16x8":                   MaxS_I16x8,
	"max_u.i16x8":                   MaxU_I16x8,
	"avgr_u.i16x8":                  AvgrU_I16x8,
	"extmul_low_i8x16_s.i16x8":      ExtmulLowI8x16S_I16x8,
	"extmul_high_i8x16_s.i16x8":     ExtmulHighI8x16S_I16x8,
	"extmul_low_i8x16_u.i16x8":      ExtmulLowI8x16U_I16x8,
	"extmul_high_i8x16_u.i16x8":     ExtmulHighI8x16U_I16x8,
	"extadd_pairwise_i16x8_s.i32x4": ExtaddPairwiseI16x8S_I32x4,
	"extadd_pairwise_i16x8_u.i32x4": ExtaddPairwiseI16x8U_I32x4,
	"abs.i32x4":                     Abs_I32x4,
	"neg.i32x4":                     Neg_I32x4,
	"alltrue.i32x4":                 Alltrue_I32x4,
	"bitmask.i32x4":                 Bitmask_I32x4,
	"extend_low_i16x8_s.i32x4":      ExtendLowI16x8S_I32x4,
	"extend_high_i16x8_s.i32x4":     ExtendHighI16x8S_I32x4,
	"extend_low_i16x8_u.i32x4":      ExtendLowI16x8U_I32x4,
	"extend_high_i16x8_u.i32x4":     ExtendHighI16x8U_I32x4,
	"shl.i32x4":                     Shl_I32x4,
	"shr_s.i32x4":                   ShrS_I32x4,
	"shr_u.i32x4":                   ShrU_I32x4,
	"add.i32x4":                     Add_I32x4,
	"sub.i32x4":                     Sub_I32x4,
	"mul.i32x4":                     Mul_I32x4,
	"min_s.i32x4":                   MinS_I32x4,
	"min_u.i32x4":                   MinU_I32x4,
	"max_s.i32x4":                   MaxS_I32x4,
	"max_u.i32x4":                   MaxU_I32x4,
	"dot_i16x8_s.i32x4":             DotI16x8S_I32x4,
	"extmul_low_i16x8_s.i32x4":      ExtmulLowI16x8S_I32x4,
	"extmul_high_i16x8_s.i32x4":     ExtmulHighI16x8S_I32x4,
	"extmul_low_i16x8_u.i32x4":      ExtmulLowI16x8U_I32x4,
	"extmul_high_i16x8_u.i32x4":     ExtmulHighI16x8U_I32x4,
	"abs.i64x2":                     Abs_I64x2,
	"neg.i64x2":                     Neg_I64x2,
	"alltrue.i64x2":                 Alltrue_I64x2,
	"bitmask.i64x2":                 Bitmask_I64x2,
	"extend_low_i32x4_s.i64x2":      ExtendLowI32x4S_I64x2,
	"extend_high_i32x4_s.i64x2":     ExtendHighI32x4S_I64x2,
	"extend_low_i32x4_u.i64x2":      ExtendLowI32x4U_I64x2,
	"extend_high_i32x4_u.i64x2":     ExtendHighI32x4U_I64x2,
	"shl.i64x2":                     Shl_I64x2,
	"shr_s.i64x2":                   ShrS_I64x2,
	"shr_u.i64x2":                   ShrU_I64x2,
	"add.i64x2":                     Add_I64x2,
	"sub.i64x2":                     Sub_I64x2,
	"mul.i64x2":                     Mul_I64x2,
	"extmul_low_i32x4_s.i64x2":      ExtmulLowI32x4S_I64x2,
	"extmul_high_i32x4_s.i64x2":     ExtmulHighI32x4S_I64x2,
	"extmul_low_i32x4_u.i64x2":      ExtmulLowI32x4U_I64x2,
	"extmul_high_i32x4_u.i64x2":     ExtmulHighI32x4U_I64x2,
	"ceil.f32x4":                    Ceil_F32x4,
	"floor.f32x4":                   Floor_F32x4,
	"trunc.f32x4":                   Trunc_F32x4,
	"nearest.f32x4":                 Nearest_F32x4,
	"abs.f32x4":                     Abs_F32x4,
	"neg.f32x4":                     Neg_F32x4,
	"sqrt.f32x4":                    Sqrt_F32x4,
	"add.f32x4":                     Add_F32x4,
	"sub.f32x4":                     Sub_F32x4,
	"mul.f32x4":                     Mul_F32x4,
	"div.f32x4":                     Div_F32x4,
	"min.f32x4":                     Min_F32x4,
	"max.f32x4":                     Max_F32x4,
	"pmin.f32x4":                    Pmin_F32x4,
	"pmax.f32x4":                    Pmax_F32x4,
	"ceil.f64x2":                    Ceil_F64x2,
	"floor.f64x2":                   Floor_F64x2,
	"trunc.f64x2":                   Trunc_F64x2,
	"nearest.f64x2":                 Nearest_F64x2,
	"abs.f64x2":                     Abs_F64x2,
	"neg.f64x2":                     Neg_F64x2,
	"sqrt.f64x2":                    Sqrt_F64x2,
	"add.f64x2":                     Add_F64x2,
	"sub.f64x2":                     Sub_F64x2,
	"mul.f64x2":                     Mul_F64x2,
	"div.f64x2":                     Div_F64x2,
	"min.f64x2":                     Min_F64x2,
	"max.f64x2":                     Max_F64x2,
	"pmin.f64x2":                    Pmin_F64x2,
	"pmax.f64x2":                    Pmax_F64x2,
	"trunc_sat_f32x4_s.i32x4":       TruncSatF32x4S_I32x4,
	"trunc_sat_f32x4_u.i32x4":       TruncSatF32x4U_I32x4,
	"convert_i32x4_s.f32x4":         ConvertI32x4S_F32x4,
	"convert_i32x4_u.f32x4":         ConvertI32x4U_F32x4,
	"trunc_sat_f64x2_s_zero.i32x4":  TruncSatF64x2SZero_I32x4,
	"trunc_sat_f64x2_u_zero.i32x4":  TruncSatF64x2UZero_I32x4,
	"convert_low_i32x4_s.f64x2":     ConvertLowI32x4S_F64x2,
	"convert_low_i32x4_u.f64x2":     ConvertLowI32x4U_F64x2,
	"demote_f64x2_zero.f32x4":       DemoteF64x2Zero_F32x4,
	"promote_low_f32x4.f64x2":       PromoteLowF32x4_F64x2,
}

var name2Map = map[string]Instr{
	"Unreachable":                Unreachable,
	"Nop":                        Nop,
	"Block":                      Block,
	"Loop":                       Loop,
	"If":                         If,
	"Else":                       Else,
	"End":                        End,
	"Br":                         Br,
	"BrIf":                       BrIf,
	"BrTable":                    BrTable,
	"Return":                     Return,
	"Call":                       Call,
	"CallIndirect":               CallIndirect,
	"Drop":                       Drop,
	"Select":                     Select,
	"SelectType":                 SelectType,
	"LocalGet":                   LocalGet,
	"LocalSet":                   LocalSet,
	"LocalTee":                   LocalTee,
	"GlobalGet":                  GlobalGet,
	"GlobalSet":                  GlobalSet,
	"TableGet":                   TableGet,
	"TableSet":                   TableSet,
	"TableInit":                  TableInit,
	"ElemDrop":                   ElemDrop,
	"TableCopy":                  TableCopy,
	"TableGrow":                  TableGrow,
	"TableSize":                  TableSize,
	"TableFill":                  TableFill,
	"Load_I32":                   Load_I32,
	"Load_I64":                   Load_I64,
	"Load_F32":                   Load_F32,
	"Load_F64":                   Load_F64,
	"Load8S_I32":                 Load8S_I32,
	"Load8U_I32":                 Load8U_I32,
	"Load16S_I32":                Load16S_I32,
	"Load16U_I32":                Load16U_I32,
	"Load8S_I64":                 Load8S_I64,
	"Load8U_I64":                 Load8U_I64,
	"Load16S_I64":                Load16S_I64,
	"Load16U_I64":                Load16U_I64,
	"Load32S_I64":                Load32S_I64,
	"Load32U_I64":                Load32U_I64,
	"Store_I32":                  Store_I32,
	"Store_I64":                  Store_I64,
	"Store_F32":                  Store_F32,
	"Store_F64":                  Store_F64,
	"Store8_I32":                 Store8_I32,
	"Store16_I32":                Store16_I32,
	"Store8_I64":                 Store8_I64,
	"Store16_I64":                Store16_I64,
	"Store32_I64":                Store32_I64,
	"MemorySize":                 MemorySize,
	"MemoryGrow":                 MemoryGrow,
	"MemoryInit":                 MemoryInit,
	"DataDrop":                   DataDrop,
	"MemoryCopy":                 MemoryCopy,
	"MemoryFill":                 MemoryFill,
	"Const_I32":                  Const_I32,
	"Const_I64":                  Const_I64,
	"Const_F32":                  Const_F32,
	"Const_F64":                  Const_F64,
	"Eqz_I32":                    Eqz_I32,
	"Eq_I32":                     Eq_I32,
	"Ne_I32":                     Ne_I32,
	"LtS_I32":                    LtS_I32,
	"LtU_I32":                    LtU_I32,
	"GtS_I32":                    GtS_I32,
	"GtU_I32":                    GtU_I32,
	"LeS_I32":                    LeS_I32,
	"LeU_I32":                    LeU_I32,
	"GeS_I32":                    GeS_I32,
	"GeU_I32":                    GeU_I32,
	"Eqz_I64":                    Eqz_I64,
	"Eq_I64":                     Eq_I64,
	"Ne_I64":                     Ne_I64,
	"LtS_I64":                    LtS_I64,
	"LtU_I64":                    LtU_I64,
	"GtS_I64":                    GtS_I64,
	"GtU_I64":                    GtU_I64,
	"LeS_I64":                    LeS_I64,
	"LeU_I64":                    LeU_I64,
	"GeS_I64":                    GeS_I64,
	"GeU_I64":                    GeU_I64,
	"Eq_F32":                     Eq_F32,
	"Ne_F32":                     Ne_F32,
	"Lt_F32":                     Lt_F32,
	"Gt_F32":                     Gt_F32,
	"Le_F32":                     Le_F32,
	"Ge_F32":                     Ge_F32,
	"Eq_F64":                     Eq_F64,
	"Ne_F64":                     Ne_F64,
	"Lt_F64":                     Lt_F64,
	"Gt_F64":                     Gt_F64,
	"Le_F64":                     Le_F64,
	"Ge_F64":                     Ge_F64,
	"Clz_I32":                    Clz_I32,
	"Ctz_I32":                    Ctz_I32,
	"Popcnt_I32":                 Popcnt_I32,
	"Add_I32":                    Add_I32,
	"Sub_I32":                    Sub_I32,
	"Mul_I32":                    Mul_I32,
	"DivS_I32":                   DivS_I32,
	"DivU_I32":                   DivU_I32,
	"RemS_I32":                   RemS_I32,
	"RemU_I32":                   RemU_I32,
	"And_I32":                    And_I32,
	"Or_I32":                     Or_I32,
	"Xor_I32":                    Xor_I32,
	"Shl_I32":                    Shl_I32,
	"ShrS_I32":                   ShrS_I32,
	"ShrU_I32":                   ShrU_I32,
	"RotL_I32":                   RotL_I32,
	"RotR_I32":                   RotR_I32,
	"Clz_I64":                    Clz_I64,
	"Ctz_I64":                    Ctz_I64,
	"Popcnt_I64":                 Popcnt_I64,
	"Add_I64":                    Add_I64,
	"Sub_I64":                    Sub_I64,
	"Mul_I64":                    Mul_I64,
	"DivS_I64":                   DivS_I64,
	"DivU_I64":                   DivU_I64,
	"RemS_I64":                   RemS_I64,
	"RemU_I64":                   RemU_I64,
	"And_I64":                    And_I64,
	"Or_I64":                     Or_I64,
	"Xor_I64":                    Xor_I64,
	"Shl_I64":                    Shl_I64,
	"ShrS_I64":                   ShrS_I64,
	"ShrU_I64":                   ShrU_I64,
	"RotL_I64":                   RotL_I64,
	"RotR_I64":                   RotR_I64,
	"Abs_F32":                    Abs_F32,
	"Neg_F32":                    Neg_F32,
	"Ceil_F32":                   Ceil_F32,
	"Floor_F32":                  Floor_F32,
	"Trunc_F32":                  Trunc_F32,
	"Nearest_F32":                Nearest_F32,
	"Sqrt_F32":                   Sqrt_F32,
	"Add_F32":                    Add_F32,
	"Sub_F32":                    Sub_F32,
	"Mul_F32":                    Mul_F32,
	"Div_F32":                    Div_F32,
	"Min_F32":                    Min_F32,
	"Max_F32":                    Max_F32,
	"Copysign_F32":               Copysign_F32,
	"Abs_F64":                    Abs_F64,
	"Neg_F64":                    Neg_F64,
	"Ceil_F64":                   Ceil_F64,
	"Floor_F64":                  Floor_F64,
	"Trunc_F64":                  Trunc_F64,
	"Nearest_F64":                Nearest_F64,
	"Sqrt_F64":                   Sqrt_F64,
	"Add_F64":                    Add_F64,
	"Sub_F64":                    Sub_F64,
	"Mul_F64":                    Mul_F64,
	"Div_F64":                    Div_F64,
	"Min_F64":                    Min_F64,
	"Max_F64":                    Max_F64,
	"Copysign_F64":               Copysign_F64,
	"WrapI64_I32":                WrapI64_I32,
	"TruncF32S_I32":              TruncF32S_I32,
	"TruncF32U_I32":              TruncF32U_I32,
	"TruncF64S_I32":              TruncF64S_I32,
	"TruncF64U_I32":              TruncF64U_I32,
	"ExtendI32S_I64":             ExtendI32S_I64,
	"ExtendI32U_I64":             ExtendI32U_I64,
	"TruncF32S_I64":              TruncF32S_I64,
	"TruncF32U_I64":              TruncF32U_I64,
	"TruncF64S_I64":              TruncF64S_I64,
	"TruncF64U_I64":              TruncF64U_I64,
	"ConvertI32S_F32":            ConvertI32S_F32,
	"ConvertI32U_F32":            ConvertI32U_F32,
	"ConvertI64S_F32":            ConvertI64S_F32,
	"ConvertI64U_F32":            ConvertI64U_F32,
	"DemoteF64_F32":              DemoteF64_F32,
	"ConvertI32S_F64":            ConvertI32S_F64,
	"ConvertI32U_F64":            ConvertI32U_F64,
	"ConvertI64S_F64":            ConvertI64S_F64,
	"ConvertI64U_F64":            ConvertI64U_F64,
	"PromoteF32_F64":             PromoteF32_F64,
	"ReinterpretF32_I32":         ReinterpretF32_I32,
	"ReinterpretF64_I64":         ReinterpretF64_I64,
	"ReinterpretI32_F32":         ReinterpretI32_F32,
	"ReinterpretI64_F64":         ReinterpretI64_F64,
	"Extend8S_I32":               Extend8S_I32,
	"Extend16S_I32":              Extend16S_I32,
	"Extend8S_I64":               Extend8S_I64,
	"Extend16S_I64":              Extend16S_I64,
	"Extend32S_I64":              Extend32S_I64,
	"TruncSatF32S_I32":           TruncSatF32S_I32,
	"TruncSatF32U_I32":           TruncSatF32U_I32,
	"TruncSatF64S_I32":           TruncSatF64S_I32,
	"TruncSatF64U_I32":           TruncSatF64U_I32,
	"TruncSatF32S_I64":           TruncSatF32S_I64,
	"TruncSatF32U_I64":           TruncSatF32U_I64,
	"TruncSatF64S_I64":           TruncSatF64S_I64,
	"TruncSatF64U_I64":           TruncSatF64U_I64,
	"Load_V128":                  Load_V128,
	"Load8x8S_V128":              Load8x8S_V128,
	"Load8x8U_V128":              Load8x8U_V128,
	"Load16x4S_V128":             Load16x4S_V128,
	"Load16x4U_V128":             Load16x4U_V128,
	"Load32x2S_V128":             Load32x2S_V128,
	"Load32x2U_V128":             Load32x2U_V128,
	"Load8Splat_V128":            Load8Splat_V128,
	"Load16Splat_V128":           Load16Splat_V128,
	"Load32Splat_V128":           Load32Splat_V128,
	"Load64Splat_V128":           Load64Splat_V128,
	"Load32Zero_V128":            Load32Zero_V128,
	"Load64Zero_V128":            Load64Zero_V128,
	"Store_V128":                 Store_V128,
	"Load8Lane_V128":             Load8Lane_V128,
	"Load16Lane_V128":            Load16Lane_V128,
	"Load32Lane_V128":            Load32Lane_V128,
	"Load64Lane_V128":            Load64Lane_V128,
	"Store8Lane_V128":            Store8Lane_V128,
	"Store16Lane_V128":           Store16Lane_V128,
	"Store32Lane_V128":           Store32Lane_V128,
	"Store64Lane_V128":           Store64Lane_V128,
	"Const_V128":                 Const_V128,
	"Shuffle_I8x16":              Shuffle_I8x16,
	"ExtractLaneS_I8x16":         ExtractLaneS_I8x16,
	"ExtractLaneU_I8x16":         ExtractLaneU_I8x16,
	"ReplaceLane_I8x16":          ReplaceLane_I8x16,
	"ExtractLaneS_I16x8":         ExtractLaneS_I16x8,
	"ExtractLaneU_I16x8":         ExtractLaneU_I16x8,
	"ReplaceLane_I16x8":          ReplaceLane_I16x8,
	"ExtractLane_I32x4":          ExtractLane_I32x4,
	"ReplaceLane_I32x4":          ReplaceLane_I32x4,
	"ExtractLane_I64x2":          ExtractLane_I64x2,
	"ReplaceLane_I64x2":          ReplaceLane_I64x2,
	"ExtractLane_F32x4":          ExtractLane_F32x4,
	"ReplaceLane_F32x4":          ReplaceLane_F32x4,
	"ExtractLane_F64x2":          ExtractLane_F64x2,
	"ReplaceLane_F64x2":          ReplaceLane_F64x2,
	"Swizzle_I8x16":              Swizzle_I8x16,
	"Splat_I8x16":                Splat_I8x16,
	"Splat_I16x8":                Splat_I16x8,
	"Splat_I32x4":                Splat_I32x4,
	"Splat_I64x2":                Splat_I64x2,
	"Splat_F32x4":                Splat_F32x4,
	"Splat_F64x2":                Splat_F64x2,
	"Eq_I8x16":                   Eq_I8x16,
	"Ne_I8x16":                   Ne_I8x16,
	"LtS_I8x16":                  LtS_I8x16,
	"LtU_I8x16":                  LtU_I8x16,
	"GtS_I8x16":                  GtS_I8x16,
	"GtU_I8x16":                  GtU_I8x16,
	"LeS_I8x16":                  LeS_I8x16,
	"LeU_I8x16":                  LeU_I8x16,
	"GeS_I8x16":                  GeS_I8x16,
	"GeU_I8x16":                  GeU_I8x16,
	"Eq_I16x8":                   Eq_I16x8,
	"Ne_I16x8":                   Ne_I16x8,
	"LtS_I16x8":                  LtS_I16x8,
	"LtU_I16x8":                  LtU_I16x8,
	"GtS_I16x8":                  GtS_I16x8,
	"GtU_I16x8":                  GtU_I16x8,
	"LeS_I16x8":                  LeS_I16x8,
	"LeU_I16x8":                  LeU_I16x8,
	"GeS_I16x8":                  GeS_I16x8,
	"GeU_I16x8":                  GeU_I16x8,
	"Eq_I32x4":                   Eq_I32x4,
	"Ne_I32x4":                   Ne_I32x4,
	"LtS_I32x4":                  LtS_I32x4,
	"LtU_I32x4":                  LtU_I32x4,
	"GtS_I32x4":                  GtS_I32x4,
	"GtU_I32x4":                  GtU_I32x4,
	"LeS_I32x4":                  LeS_I32x4,
	"LeU_I32x4":                  LeU_I32x4,
	"GeS_I32x4":                  GeS_I32x4,
	"GeU_I32x4":                  GeU_I32x4,
	"Eq_I64x2":                   Eq_I64x2,
	"Ne_I64x2":                   Ne_I64x2,
	"LtS_I64x2":                  LtS_I64x2,
	"GtS_I64x2":                  GtS_I64x2,
	"LeS_I64x2":                  LeS_I64x2,
	"GeS_I64x2":                  GeS_I64x2,
	"Eq_F32x4":                   Eq_F32x4,
	"Ne_F32x4":                   Ne_F32x4,
	"Lt_F32x4":                   Lt_F32x4,
	"Gt_F32x4":                   Gt_F32x4,
	"Le_F32x4":                   Le_F32x4,
	"Ge_F32x4":                   Ge_F32x4,
	"Eq_F64x2":                   Eq_F64x2,
	"Ne_F64x2":                   Ne_F64x2,
	"Lt_F64x2":                   Lt_F64x2,
	"Gt_F64x2":                   Gt_F64x2,
	"Le_F64x2":                   Le_F64x2,
	"Ge_F64x2":                   Ge_F64x2,
	"Not_V128":                   Not_V128,
	"And_V128":                   And_V128,
	"Andnot_V128":                Andnot_V128,
	"Or_V128":                    Or_V128,
	"Xor_V128":                   Xor_V128,
	"Bitselect_V128":             Bitselect_V128,
	"AnyTrue_V128":               AnyTrue_V128,
	"Abs_I8x16":                  Abs_I8x16,
	"Neg_I8x16":                  Neg_I8x16,
	"Popcnt_I8x16":               Popcnt_I8x16,
	"AllTrue_I8x16":              AllTrue_I8x16,
	"Bitmask_I8x16":              Bitmask_I8x16,
	"NarrowI16x8S_I8x16":         NarrowI16x8S_I8x16,
	"NarrowI16x8U_I8x16":         NarrowI16x8U_I8x16,
	"Shl_I8x16":                  Shl_I8x16,
	"ShrS_I8x16":                 ShrS_I8x16,
	"ShrU_I8x16":                 ShrU_I8x16,
	"Add_I8x16":                  Add_I8x16,
	"AddSatS_I8x16":              AddSatS_I8x16,
	"AddSatU_I8x16":              AddSatU_I8x16,
	"Sub_I8x16":                  Sub_I8x16,
	"SubSatS_I8x16":              SubSatS_I8x16,
	"SubSatU_I8x16":              SubSatU_I8x16,
	"MinS_I8x16":                 MinS_I8x16,
	"MinU_I8x16":                 MinU_I8x16,
	"MaxS_I8x16":                 MaxS_I8x16,
	"MaxU_I8x16":                 MaxU_I8x16,
	"AvgrU_I8x16":                AvgrU_I8x16,
	"ExtaddPairwiseI8x16S_I16x8": ExtaddPairwiseI8x16S_I16x8,
	"ExtaddPairwiseI8x16U_I16x8": ExtaddPairwiseI8x16U_I16x8,
	"Abs_I16x8":                  Abs_I16x8,
	"Neg_I16x8":                  Neg_I16x8,
	"Q15mulrSatS_I16x8":          Q15mulrSatS_I16x8,
	"Alltrue_I16x8":              Alltrue_I16x8,
	"Bitmask_I16x8":              Bitmask_I16x8,
	"NarrowI32x4S_I16x8":         NarrowI32x4S_I16x8,
	"NarrowI32x4U_I16x8":         NarrowI32x4U_I16x8,
	"ExtendLowI8x16S_I16x8":      ExtendLowI8x16S_I16x8,
	"ExtendHighI8x16S_I16x8":     ExtendHighI8x16S_I16x8,
	"ExtendLowI8x16U_I16x8":      ExtendLowI8x16U_I16x8,
	"ExtendHighI8x16U_I16x8":     ExtendHighI8x16U_I16x8,
	"Shl_I16x8":                  Shl_I16x8,
	"ShrS_I16x8":                 ShrS_I16x8,
	"ShrU_I16x8":                 ShrU_I16x8,
	"Add_I16x8":                  Add_I16x8,
	"AddSatS_I16x8":              AddSatS_I16x8,
	"AddSatU_I16x8":              AddSatU_I16x8,
	"Sub_I16x8":                  Sub_I16x8,
	"SubSatS_I16x8":              SubSatS_I16x8,
	"SubSatU_I16x8":              SubSatU_I16x8,
	"Mul_I16x8":                  Mul_I16x8,
	"MinS_I16x8":                 MinS_I16x8,
	"MinU_I16x8":                 MinU_I16x8,
	"MaxS_I16x8":                 MaxS_I16x8,
	"MaxU_I16x8":                 MaxU_I16x8,
	"AvgrU_I16x8":                AvgrU_I16x8,
	"ExtmulLowI8x16S_I16x8":      ExtmulLowI8x16S_I16x8,
	"ExtmulHighI8x16S_I16x8":     ExtmulHighI8x16S_I16x8,
	"ExtmulLowI8x16U_I16x8":      ExtmulLowI8x16U_I16x8,
	"ExtmulHighI8x16U_I16x8":     ExtmulHighI8x16U_I16x8,
	"ExtaddPairwiseI16x8S_I32x4": ExtaddPairwiseI16x8S_I32x4,
	"ExtaddPairwiseI16x8U_I32x4": ExtaddPairwiseI16x8U_I32x4,
	"Abs_I32x4":                  Abs_I32x4,
	"Neg_I32x4":                  Neg_I32x4,
	"Alltrue_I32x4":              Alltrue_I32x4,
	"Bitmask_I32x4":              Bitmask_I32x4,
	"ExtendLowI16x8S_I32x4":      ExtendLowI16x8S_I32x4,
	"ExtendHighI16x8S_I32x4":     ExtendHighI16x8S_I32x4,
	"ExtendLowI16x8U_I32x4":      ExtendLowI16x8U_I32x4,
	"ExtendHighI16x8U_I32x4":     ExtendHighI16x8U_I32x4,
	"Shl_I32x4":                  Shl_I32x4,
	"ShrS_I32x4":                 ShrS_I32x4,
	"ShrU_I32x4":                 ShrU_I32x4,
	"Add_I32x4":                  Add_I32x4,
	"Sub_I32x4":                  Sub_I32x4,
	"Mul_I32x4":                  Mul_I32x4,
	"MinS_I32x4":                 MinS_I32x4,
	"MinU_I32x4":                 MinU_I32x4,
	"MaxS_I32x4":                 MaxS_I32x4,
	"MaxU_I32x4":                 MaxU_I32x4,
	"DotI16x8S_I32x4":            DotI16x8S_I32x4,
	"ExtmulLowI16x8S_I32x4":      ExtmulLowI16x8S_I32x4,
	"ExtmulHighI16x8S_I32x4":     ExtmulHighI16x8S_I32x4,
	"ExtmulLowI16x8U_I32x4":      ExtmulLowI16x8U_I32x4,
	"ExtmulHighI16x8U_I32x4":     ExtmulHighI16x8U_I32x4,
	"Abs_I64x2":                  Abs_I64x2,
	"Neg_I64x2":                  Neg_I64x2,
	"Alltrue_I64x2":              Alltrue_I64x2,
	"Bitmask_I64x2":              Bitmask_I64x2,
	"ExtendLowI32x4S_I64x2":      ExtendLowI32x4S_I64x2,
	"ExtendHighI32x4S_I64x2":     ExtendHighI32x4S_I64x2,
	"ExtendLowI32x4U_I64x2":      ExtendLowI32x4U_I64x2,
	"ExtendHighI32x4U_I64x2":     ExtendHighI32x4U_I64x2,
	"Shl_I64x2":                  Shl_I64x2,
	"ShrS_I64x2":                 ShrS_I64x2,
	"ShrU_I64x2":                 ShrU_I64x2,
	"Add_I64x2":                  Add_I64x2,
	"Sub_I64x2":                  Sub_I64x2,
	"Mul_I64x2":                  Mul_I64x2,
	"ExtmulLowI32x4S_I64x2":      ExtmulLowI32x4S_I64x2,
	"ExtmulHighI32x4S_I64x2":     ExtmulHighI32x4S_I64x2,
	"ExtmulLowI32x4U_I64x2":      ExtmulLowI32x4U_I64x2,
	"ExtmulHighI32x4U_I64x2":     ExtmulHighI32x4U_I64x2,
	"Ceil_F32x4":                 Ceil_F32x4,
	"Floor_F32x4":                Floor_F32x4,
	"Trunc_F32x4":                Trunc_F32x4,
	"Nearest_F32x4":              Nearest_F32x4,
	"Abs_F32x4":                  Abs_F32x4,
	"Neg_F32x4":                  Neg_F32x4,
	"Sqrt_F32x4":                 Sqrt_F32x4,
	"Add_F32x4":                  Add_F32x4,
	"Sub_F32x4":                  Sub_F32x4,
	"Mul_F32x4":                  Mul_F32x4,
	"Div_F32x4":                  Div_F32x4,
	"Min_F32x4":                  Min_F32x4,
	"Max_F32x4":                  Max_F32x4,
	"Pmin_F32x4":                 Pmin_F32x4,
	"Pmax_F32x4":                 Pmax_F32x4,
	"Ceil_F64x2":                 Ceil_F64x2,
	"Floor_F64x2":                Floor_F64x2,
	"Trunc_F64x2":                Trunc_F64x2,
	"Nearest_F64x2":              Nearest_F64x2,
	"Abs_F64x2":                  Abs_F64x2,
	"Neg_F64x2":                  Neg_F64x2,
	"Sqrt_F64x2":                 Sqrt_F64x2,
	"Add_F64x2":                  Add_F64x2,
	"Sub_F64x2":                  Sub_F64x2,
	"Mul_F64x2":                  Mul_F64x2,
	"Div_F64x2":                  Div_F64x2,
	"Min_F64x2":                  Min_F64x2,
	"Max_F64x2":                  Max_F64x2,
	"Pmin_F64x2":                 Pmin_F64x2,
	"Pmax_F64x2":                 Pmax_F64x2,
	"TruncSatF32x4S_I32x4":       TruncSatF32x4S_I32x4,
	"TruncSatF32x4U_I32x4":       TruncSatF32x4U_I32x4,
	"ConvertI32x4S_F32x4":        ConvertI32x4S_F32x4,
	"ConvertI32x4U_F32x4":        ConvertI32x4U_F32x4,
	"TruncSatF64x2SZero_I32x4":   TruncSatF64x2SZero_I32x4,
	"TruncSatF64x2UZero_I32x4":   TruncSatF64x2UZero_I32x4,
	"ConvertLowI32x4S_F64x2":     ConvertLowI32x4S_F64x2,
	"ConvertLowI32x4U_F64x2":     ConvertLowI32x4U_F64x2,
	"DemoteF64x2Zero_F32x4":      DemoteF64x2Zero_F32x4,
	"PromoteLowF32x4_F64x2":      PromoteLowF32x4_F64x2,
}
