package main

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
	{Name: "invalid"},
	{I: Unreachable, Class: Control, Name: "unreachable", Name2: "Unreachable", Op: uint8(0x0)},
	{I: Nop, Class: Control, Name: "nop", Name2: "Nop", Op: uint8(0x1)},
	{I: Block, Class: Control, Name: "block", Name2: "Block", Op: uint8(0x2)},
	{I: Loop, Class: Control, Name: "loop", Name2: "Loop", Op: uint8(0x3)},
	{I: If, Class: Control, Name: "if", Name2: "If", Op: uint8(0x4)},
	{I: Else, Class: Control, Name: "else", Name2: "Else", Op: uint8(0x5)},
	{I: End, Class: Control, Name: "end", Name2: "End", Op: uint8(0xB)},
	{I: Br, Class: Control, Name: "br", Name2: "Br", Op: uint8(0xC)},
	{I: BrIf, Class: Control, Name: "br_if", Name2: "BrIf", Op: uint8(0xD), A: I32},
	{I: BrTable, Class: Control, Name: "br_table", Name2: "BrTable", Op: uint8(0xE), A: I32},
	{I: Return, Class: Control, Name: "return", Name2: "Return", Op: uint8(0xF)},
	{I: Call, Class: Control, Name: "call", Name2: "Call", Op: uint8(0x10)},
	{I: CallIndirect, Class: Control, Name: "call_indirect", Name2: "CallIndirect", Op: uint8(0x11)},
	{I: Drop, Class: Parametric, Name: "drop", Name2: "Drop", Op: uint8(0x1A)},
	{I: Select, Class: Parametric, Name: "select", Name2: "Select", Op: uint8(0x1B), C: I32},
	{I: SelectType, Class: Parametric, Name: "select_type", Name2: "SelectType", Op: uint8(0x1C), C: I32},
	{I: LocalGet, Class: Variable, Name: "local_get", Name2: "LocalGet", Op: uint8(0x20)},
	{I: LocalSet, Class: Variable, Name: "local_set", Name2: "LocalSet", Op: uint8(0x21)},
	{I: LocalTee, Class: Variable, Name: "local_tee", Name2: "LocalTee", Op: uint8(0x22)},
	{I: GlobalGet, Class: Variable, Name: "global_get", Name2: "GlobalGet", Op: uint8(0x23)},
	{I: GlobalSet, Class: Variable, Name: "global_set", Name2: "GlobalSet", Op: uint8(0x24)},
	{I: TableGet, Class: Table, Name: "table_get", Name2: "TableGet", Op: uint8(0x25), A: I32},
	{I: TableSet, Class: Table, Name: "table_set", Name2: "TableSet", Op: uint8(0x26), A: I32},
	{I: TableInit, Class: Table, Name: "table_init", Name2: "TableInit", OpPfx: uint8(0xFC), Op: uint8(0xC), A: I32, B: I32, C: I32},
	{I: ElemDrop, Class: Table, Name: "elem_drop", Name2: "ElemDrop", OpPfx: uint8(0xFC), Op: uint8(0xD)},
	{I: TableCopy, Class: Table, Name: "table_copy", Name2: "TableCopy", OpPfx: uint8(0xFC), Op: uint8(0xE), A: I32, B: I32, C: I32},
	{I: TableGrow, Class: Table, Name: "table_grow", Name2: "TableGrow", OpPfx: uint8(0xFC), Op: uint8(0xF), B: I32},
	{I: TableSize, Class: Table, Name: "table_size", Name2: "TableSize", OpPfx: uint8(0xFC), Op: uint8(0x10)},
	{I: TableFill, Class: Table, Name: "table_fill", Name2: "TableFill", OpPfx: uint8(0xFC), Op: uint8(0x11), A: I32, C: I32},
	{I: Load_I32, Class: Memory, Name: "load.i32", Name2: "Load_I32", Op: uint8(0x28), T: I32, A: I32, Ma: Load, Mz: 4},
	{I: Load_I64, Class: Memory, Name: "load.i64", Name2: "Load_I64", Op: uint8(0x29), T: I64, A: I32, Ma: Load, Mz: 8},
	{I: Load_F32, Class: Memory, Name: "load.f32", Name2: "Load_F32", Op: uint8(0x2A), T: F32, A: I32, Ma: Load, Mz: 4},
	{I: Load_F64, Class: Memory, Name: "load.f64", Name2: "Load_F64", Op: uint8(0x2B), T: F64, A: I32, Ma: Load, Mz: 8},
	{I: Load8S_I32, Class: Memory, Name: "load8_s.i32", Name2: "Load8S_I32", Op: uint8(0x2C), T: I32, A: I32, Ma: Load, Mz: 1},
	{I: Load8U_I32, Class: Memory, Name: "load8_u.i32", Name2: "Load8U_I32", Op: uint8(0x2D), T: I32, A: I32, Ma: Load, Mz: 1},
	{I: Load16S_I32, Class: Memory, Name: "load16_s.i32", Name2: "Load16S_I32", Op: uint8(0x2E), T: I32, A: I32, Ma: Load, Mz: 2},
	{I: Load16U_I32, Class: Memory, Name: "load16_u.i32", Name2: "Load16U_I32", Op: uint8(0x2F), T: I32, A: I32, Ma: Load, Mz: 2},
	{I: Load8S_I64, Class: Memory, Name: "load8_s.i64", Name2: "Load8S_I64", Op: uint8(0x30), T: I64, A: I32, Ma: Load, Mz: 1},
	{I: Load8U_I64, Class: Memory, Name: "load8_u.i64", Name2: "Load8U_I64", Op: uint8(0x31), T: I64, A: I32, Ma: Load, Mz: 1},
	{I: Load16S_I64, Class: Memory, Name: "load16_s.i64", Name2: "Load16S_I64", Op: uint8(0x32), T: I64, A: I32, Ma: Load, Mz: 2},
	{I: Load16U_I64, Class: Memory, Name: "load16_u.i64", Name2: "Load16U_I64", Op: uint8(0x33), T: I64, A: I32, Ma: Load, Mz: 2},
	{I: Load32S_I64, Class: Memory, Name: "load32_s.i64", Name2: "Load32S_I64", Op: uint8(0x34), T: I64, A: I32, Ma: Load, Mz: 4},
	{I: Load32U_I64, Class: Memory, Name: "load32_u.i64", Name2: "Load32U_I64", Op: uint8(0x35), T: I64, A: I32, Ma: Load, Mz: 4},
	{I: Store_I32, Class: Memory, Name: "store.i32", Name2: "Store_I32", Op: uint8(0x36), A: I32, B: I32, Ma: Store, Mz: 4},
	{I: Store_I64, Class: Memory, Name: "store.i64", Name2: "Store_I64", Op: uint8(0x37), A: I32, B: I64, Ma: Store, Mz: 8},
	{I: Store_F32, Class: Memory, Name: "store.f32", Name2: "Store_F32", Op: uint8(0x38), A: I32, B: F32, Ma: Store, Mz: 4},
	{I: Store_F64, Class: Memory, Name: "store.f64", Name2: "Store_F64", Op: uint8(0x39), A: I32, B: F64, Ma: Store, Mz: 8},
	{I: Store8_I32, Class: Memory, Name: "store8.i32", Name2: "Store8_I32", Op: uint8(0x3A), A: I32, B: I32, Ma: Store, Mz: 1},
	{I: Store16_I32, Class: Memory, Name: "store16.i32", Name2: "Store16_I32", Op: uint8(0x3B), A: I32, B: I32, Ma: Store, Mz: 2},
	{I: Store8_I64, Class: Memory, Name: "store8.i64", Name2: "Store8_I64", Op: uint8(0x3C), A: I32, B: I64, Ma: Store, Mz: 1},
	{I: Store16_I64, Class: Memory, Name: "store16.i64", Name2: "Store16_I64", Op: uint8(0x3D), A: I32, B: I64, Ma: Store, Mz: 2},
	{I: Store32_I64, Class: Memory, Name: "store32.i64", Name2: "Store32_I64", Op: uint8(0x3E), A: I32, B: I64, Ma: Store, Mz: 4},
	{I: MemorySize, Class: Memory, Name: "memory_size", Name2: "MemorySize", Op: uint8(0x3F), T: I32},
	{I: MemoryGrow, Class: Memory, Name: "memory_grow", Name2: "MemoryGrow", Op: uint8(0x40), T: I32, A: I32},
	{I: MemoryInit, Class: Memory, Name: "memory_init", Name2: "MemoryInit", OpPfx: uint8(0xFC), Op: uint8(0x8), A: I32, B: I32, C: I32},
	{I: DataDrop, Class: Memory, Name: "data_drop", Name2: "DataDrop", OpPfx: uint8(0xFC), Op: uint8(0x9)},
	{I: MemoryCopy, Class: Memory, Name: "memory_copy", Name2: "MemoryCopy", OpPfx: uint8(0xFC), Op: uint8(0xA), A: I32, B: I32, C: I32},
	{I: MemoryFill, Class: Memory, Name: "memory_fill", Name2: "MemoryFill", OpPfx: uint8(0xFC), Op: uint8(0xB), A: I32, B: I32, C: I32},
	{I: Const_I32, Class: Numeric, Name: "const.i32", Name2: "Const_I32", Op: uint8(0x41), T: I32},
	{I: Const_I64, Class: Numeric, Name: "const.i64", Name2: "Const_I64", Op: uint8(0x42), T: F32},
	{I: Const_F32, Class: Numeric, Name: "const.f32", Name2: "Const_F32", Op: uint8(0x43), T: I64},
	{I: Const_F64, Class: Numeric, Name: "const.f64", Name2: "Const_F64", Op: uint8(0x44), T: F64},
	{I: Eqz_I32, Class: Numeric, Name: "eqz.i32", Name2: "Eqz_I32", Op: uint8(0x45), T: I32, A: I32},
	{I: Eq_I32, Class: Numeric, Name: "eq.i32", Name2: "Eq_I32", Op: uint8(0x46), T: I32, A: I32, B: I32},
	{I: Ne_I32, Class: Numeric, Name: "ne.i32", Name2: "Ne_I32", Op: uint8(0x47), T: I32, A: I32, B: I32},
	{I: LtS_I32, Class: Numeric, Name: "lt_s.i32", Name2: "LtS_I32", Op: uint8(0x48), T: I32, A: I32, B: I32},
	{I: LtU_I32, Class: Numeric, Name: "lt_u.i32", Name2: "LtU_I32", Op: uint8(0x49), T: I32, A: I32, B: I32},
	{I: GtS_I32, Class: Numeric, Name: "gt_s.i32", Name2: "GtS_I32", Op: uint8(0x4A), T: I32, A: I32, B: I32},
	{I: GtU_I32, Class: Numeric, Name: "gt_u.i32", Name2: "GtU_I32", Op: uint8(0x4B), T: I32, A: I32, B: I32},
	{I: LeS_I32, Class: Numeric, Name: "le_s.i32", Name2: "LeS_I32", Op: uint8(0x4C), T: I32, A: I32, B: I32},
	{I: LeU_I32, Class: Numeric, Name: "le_u.i32", Name2: "LeU_I32", Op: uint8(0x4D), T: I32, A: I32, B: I32},
	{I: GeS_I32, Class: Numeric, Name: "ge_s.i32", Name2: "GeS_I32", Op: uint8(0x4E), T: I32, A: I32, B: I32},
	{I: GeU_I32, Class: Numeric, Name: "ge_u.i32", Name2: "GeU_I32", Op: uint8(0x4F), T: I32, A: I32, B: I32},
	{I: Eqz_I64, Class: Numeric, Name: "eqz.i64", Name2: "Eqz_I64", Op: uint8(0x50), T: I32, A: I64},
	{I: Eq_I64, Class: Numeric, Name: "eq.i64", Name2: "Eq_I64", Op: uint8(0x51), T: I32, A: I64, B: I64},
	{I: Ne_I64, Class: Numeric, Name: "ne.i64", Name2: "Ne_I64", Op: uint8(0x52), T: I32, A: I64, B: I64},
	{I: LtS_I64, Class: Numeric, Name: "lt_s.i64", Name2: "LtS_I64", Op: uint8(0x53), T: I32, A: I64, B: I64},
	{I: LtU_I64, Class: Numeric, Name: "lt_u.i64", Name2: "LtU_I64", Op: uint8(0x54), T: I32, A: I64, B: I64},
	{I: GtS_I64, Class: Numeric, Name: "gt_s.i64", Name2: "GtS_I64", Op: uint8(0x55), T: I32, A: I64, B: I64},
	{I: GtU_I64, Class: Numeric, Name: "gt_u.i64", Name2: "GtU_I64", Op: uint8(0x56), T: I32, A: I64, B: I64},
	{I: LeS_I64, Class: Numeric, Name: "le_s.i64", Name2: "LeS_I64", Op: uint8(0x57), T: I32, A: I64, B: I64},
	{I: LeU_I64, Class: Numeric, Name: "le_u.i64", Name2: "LeU_I64", Op: uint8(0x58), T: I32, A: I64, B: I64},
	{I: GeS_I64, Class: Numeric, Name: "ge_s.i64", Name2: "GeS_I64", Op: uint8(0x59), T: I32, A: I64, B: I64},
	{I: GeU_I64, Class: Numeric, Name: "ge_u.i64", Name2: "GeU_I64", Op: uint8(0x5A), T: I32, A: I64, B: I64},
	{I: Eq_F32, Class: Numeric, Name: "eq.f32", Name2: "Eq_F32", Op: uint8(0x5B), T: I32, A: F32},
	{I: Ne_F32, Class: Numeric, Name: "ne.f32", Name2: "Ne_F32", Op: uint8(0x5C), T: I32, A: F32, B: F32},
	{I: Lt_F32, Class: Numeric, Name: "lt.f32", Name2: "Lt_F32", Op: uint8(0x5D), T: I32, A: F32, B: F32},
	{I: Gt_F32, Class: Numeric, Name: "gt.f32", Name2: "Gt_F32", Op: uint8(0x5E), T: I32, A: F32, B: F32},
	{I: Le_F32, Class: Numeric, Name: "le.f32", Name2: "Le_F32", Op: uint8(0x5F), T: I32, A: F32, B: F32},
	{I: Ge_F32, Class: Numeric, Name: "ge.f32", Name2: "Ge_F32", Op: uint8(0x60), T: I32, A: F32, B: F32},
	{I: Eq_F64, Class: Numeric, Name: "eq.f64", Name2: "Eq_F64", Op: uint8(0x61), T: I64, A: F64},
	{I: Ne_F64, Class: Numeric, Name: "ne.f64", Name2: "Ne_F64", Op: uint8(0x62), T: I64, A: F64, B: F64},
	{I: Lt_F64, Class: Numeric, Name: "lt.f64", Name2: "Lt_F64", Op: uint8(0x63), T: I64, A: F64, B: F64},
	{I: Gt_F64, Class: Numeric, Name: "gt.f64", Name2: "Gt_F64", Op: uint8(0x64), T: I64, A: F64, B: F64},
	{I: Le_F64, Class: Numeric, Name: "le.f64", Name2: "Le_F64", Op: uint8(0x65), T: I64, A: F64, B: F64},
	{I: Ge_F64, Class: Numeric, Name: "ge.f64", Name2: "Ge_F64", Op: uint8(0x66), T: I64, A: F64, B: F64},
	{I: Clz_I32, Class: Numeric, Name: "clz.i32", Name2: "Clz_I32", Op: uint8(0x67), T: I32, A: I32},
	{I: Ctz_I32, Class: Numeric, Name: "ctz.i32", Name2: "Ctz_I32", Op: uint8(0x68), T: I32, A: I32},
	{I: Popcnt_I32, Class: Numeric, Name: "popcnt.i32", Name2: "Popcnt_I32", Op: uint8(0x69), T: I32, A: I32},
	{I: Add_I32, Class: Numeric, Name: "add.i32", Name2: "Add_I32", Op: uint8(0x6A), T: I32, A: I32, B: I32},
	{I: Sub_I32, Class: Numeric, Name: "sub.i32", Name2: "Sub_I32", Op: uint8(0x6B), T: I32, A: I32, B: I32},
	{I: Mul_I32, Class: Numeric, Name: "mul.i32", Name2: "Mul_I32", Op: uint8(0x6C), T: I32, A: I32, B: I32},
	{I: DivS_I32, Class: Numeric, Name: "div_s.i32", Name2: "DivS_I32", Op: uint8(0x6D), T: I32, A: I32, B: I32},
	{I: DivU_I32, Class: Numeric, Name: "div_u.i32", Name2: "DivU_I32", Op: uint8(0x6E), T: I32, A: I32, B: I32},
	{I: RemS_I32, Class: Numeric, Name: "rem_s.i32", Name2: "RemS_I32", Op: uint8(0x6F), T: I32, A: I32, B: I32},
	{I: RemU_I32, Class: Numeric, Name: "rem_u.i32", Name2: "RemU_I32", Op: uint8(0x70), T: I32, A: I32, B: I32},
	{I: And_I32, Class: Numeric, Name: "and.i32", Name2: "And_I32", Op: uint8(0x71), T: I32, A: I32, B: I32},
	{I: Or_I32, Class: Numeric, Name: "or.i32", Name2: "Or_I32", Op: uint8(0x72), T: I32, A: I32, B: I32},
	{I: Xor_I32, Class: Numeric, Name: "xor.i32", Name2: "Xor_I32", Op: uint8(0x73), T: I32, A: I32, B: I32},
	{I: Shl_I32, Class: Numeric, Name: "shl.i32", Name2: "Shl_I32", Op: uint8(0x74), T: I32, A: I32, B: I32},
	{I: ShrS_I32, Class: Numeric, Name: "shr_s.i32", Name2: "ShrS_I32", Op: uint8(0x75), T: I32, A: I32, B: I32},
	{I: ShrU_I32, Class: Numeric, Name: "shr_u.i32", Name2: "ShrU_I32", Op: uint8(0x76), T: I32, A: I32, B: I32},
	{I: RotL_I32, Class: Numeric, Name: "rot_l.i32", Name2: "RotL_I32", Op: uint8(0x77), T: I32, A: I32, B: I32},
	{I: RotR_I32, Class: Numeric, Name: "rot_r.i32", Name2: "RotR_I32", Op: uint8(0x78), T: I32, A: I32, B: I32},
	{I: Clz_I64, Class: Numeric, Name: "clz.i64", Name2: "Clz_I64", Op: uint8(0x79), T: I64, A: I64},
	{I: Ctz_I64, Class: Numeric, Name: "ctz.i64", Name2: "Ctz_I64", Op: uint8(0x7A), T: I64, A: I64},
	{I: Popcnt_I64, Class: Numeric, Name: "popcnt.i64", Name2: "Popcnt_I64", Op: uint8(0x7B), T: I64, A: I64},
	{I: Add_I64, Class: Numeric, Name: "add.i64", Name2: "Add_I64", Op: uint8(0x7C), T: I64, A: I64, B: I64},
	{I: Sub_I64, Class: Numeric, Name: "sub.i64", Name2: "Sub_I64", Op: uint8(0x7D), T: I64, A: I64, B: I64},
	{I: Mul_I64, Class: Numeric, Name: "mul.i64", Name2: "Mul_I64", Op: uint8(0x7E), T: I64, A: I64, B: I64},
	{I: DivS_I64, Class: Numeric, Name: "div_s.i64", Name2: "DivS_I64", Op: uint8(0x7F), T: I64, A: I64, B: I64},
	{I: DivU_I64, Class: Numeric, Name: "div_u.i64", Name2: "DivU_I64", Op: uint8(0x80), T: I64, A: I64, B: I64},
	{I: RemS_I64, Class: Numeric, Name: "rem_s.i64", Name2: "RemS_I64", Op: uint8(0x81), T: I64, A: I64, B: I64},
	{I: RemU_I64, Class: Numeric, Name: "rem_u.i64", Name2: "RemU_I64", Op: uint8(0x82), T: I64, A: I64, B: I64},
	{I: And_I64, Class: Numeric, Name: "and.i64", Name2: "And_I64", Op: uint8(0x83), T: I64, A: I64, B: I64},
	{I: Or_I64, Class: Numeric, Name: "or.i64", Name2: "Or_I64", Op: uint8(0x84), T: I64, A: I64, B: I64},
	{I: Xor_I64, Class: Numeric, Name: "xor.i64", Name2: "Xor_I64", Op: uint8(0x85), T: I64, A: I64, B: I64},
	{I: Shl_I64, Class: Numeric, Name: "shl.i64", Name2: "Shl_I64", Op: uint8(0x86), T: I64, A: I64, B: I64},
	{I: ShrS_I64, Class: Numeric, Name: "shr_s.i64", Name2: "ShrS_I64", Op: uint8(0x87), T: I64, A: I64, B: I64},
	{I: ShrU_I64, Class: Numeric, Name: "shr_u.i64", Name2: "ShrU_I64", Op: uint8(0x88), T: I64, A: I64, B: I64},
	{I: RotL_I64, Class: Numeric, Name: "rot_l.i64", Name2: "RotL_I64", Op: uint8(0x89), T: I64, A: I64, B: I64},
	{I: RotR_I64, Class: Numeric, Name: "rot_r.i64", Name2: "RotR_I64", Op: uint8(0x8A), T: I64, A: I64, B: I64},
	{I: Abs_F32, Class: Numeric, Name: "abs.f32", Name2: "Abs_F32", Op: uint8(0x8B), T: F32, A: F32, B: F32},
	{I: Neg_F32, Class: Numeric, Name: "neg.f32", Name2: "Neg_F32", Op: uint8(0x8C), T: F32, A: F32, B: F32},
	{I: Ceil_F32, Class: Numeric, Name: "ceil.f32", Name2: "Ceil_F32", Op: uint8(0x8D), T: F32, A: F32, B: F32},
	{I: Floor_F32, Class: Numeric, Name: "floor.f32", Name2: "Floor_F32", Op: uint8(0x8E), T: F32, A: F32, B: F32},
	{I: Trunc_F32, Class: Numeric, Name: "trunc.f32", Name2: "Trunc_F32", Op: uint8(0x8F), T: F32, A: F32, B: F32},
	{I: Nearest_F32, Class: Numeric, Name: "nearest.f32", Name2: "Nearest_F32", Op: uint8(0x90), T: F32, A: F32, B: F32},
	{I: Sqrt_F32, Class: Numeric, Name: "sqrt.f32", Name2: "Sqrt_F32", Op: uint8(0x91), T: F32, A: F32, B: F32},
	{I: Add_F32, Class: Numeric, Name: "add.f32", Name2: "Add_F32", Op: uint8(0x92), T: F32, A: F32, B: F32},
	{I: Sub_F32, Class: Numeric, Name: "sub.f32", Name2: "Sub_F32", Op: uint8(0x93), T: F32, A: F32, B: F32},
	{I: Mul_F32, Class: Numeric, Name: "mul.f32", Name2: "Mul_F32", Op: uint8(0x94), T: F32, A: F32, B: F32},
	{I: Div_F32, Class: Numeric, Name: "div.f32", Name2: "Div_F32", Op: uint8(0x95), T: F32, A: F32, B: F32},
	{I: Min_F32, Class: Numeric, Name: "min.f32", Name2: "Min_F32", Op: uint8(0x96), T: F32, A: F32, B: F32},
	{I: Max_F32, Class: Numeric, Name: "max.f32", Name2: "Max_F32", Op: uint8(0x97), T: F32, A: F32, B: F32},
	{I: Copysign_F32, Class: Numeric, Name: "copysign.f32", Name2: "Copysign_F32", Op: uint8(0x98), T: F32, A: F32, B: F32},
	{I: Abs_F64, Class: Numeric, Name: "abs.f64", Name2: "Abs_F64", Op: uint8(0x99), T: F64, A: F64, B: F64},
	{I: Neg_F64, Class: Numeric, Name: "neg.f64", Name2: "Neg_F64", Op: uint8(0x9A), T: F64, A: F64, B: F64},
	{I: Ceil_F64, Class: Numeric, Name: "ceil.f64", Name2: "Ceil_F64", Op: uint8(0x9B), T: F64, A: F64, B: F64},
	{I: Floor_F64, Class: Numeric, Name: "floor.f64", Name2: "Floor_F64", Op: uint8(0x9C), T: F64, A: F64, B: F64},
	{I: Trunc_F64, Class: Numeric, Name: "trunc.f64", Name2: "Trunc_F64", Op: uint8(0x9D), T: F64, A: F64, B: F64},
	{I: Nearest_F64, Class: Numeric, Name: "nearest.f64", Name2: "Nearest_F64", Op: uint8(0x9E), T: F64, A: F64, B: F64},
	{I: Sqrt_F64, Class: Numeric, Name: "sqrt.f64", Name2: "Sqrt_F64", Op: uint8(0x9F), T: F64, A: F64, B: F64},
	{I: Add_F64, Class: Numeric, Name: "add.f64", Name2: "Add_F64", Op: uint8(0xA0), T: F64, A: F64, B: F64},
	{I: Sub_F64, Class: Numeric, Name: "sub.f64", Name2: "Sub_F64", Op: uint8(0xA1), T: F64, A: F64, B: F64},
	{I: Mul_F64, Class: Numeric, Name: "mul.f64", Name2: "Mul_F64", Op: uint8(0xA2), T: F64, A: F64, B: F64},
	{I: Div_F64, Class: Numeric, Name: "div.f64", Name2: "Div_F64", Op: uint8(0xA3), T: F64, A: F64, B: F64},
	{I: Min_F64, Class: Numeric, Name: "min.f64", Name2: "Min_F64", Op: uint8(0xA4), T: F64, A: F64, B: F64},
	{I: Max_F64, Class: Numeric, Name: "max.f64", Name2: "Max_F64", Op: uint8(0xA5), T: F64, A: F64, B: F64},
	{I: Copysign_F64, Class: Numeric, Name: "copysign.f64", Name2: "Copysign_F64", Op: uint8(0xA6), T: F64, A: F64, B: F64},
	{I: WrapI64_I32, Class: Numeric, Name: "wrap_i64.i32", Name2: "WrapI64_I32", Op: uint8(0xA7), T: I32, A: I64},
	{I: TruncF32S_I32, Class: Numeric, Name: "trunc_f32_s.i32", Name2: "TruncF32S_I32", Op: uint8(0xA8), T: I32, A: F32},
	{I: TruncF32U_I32, Class: Numeric, Name: "trunc_f32_u.i32", Name2: "TruncF32U_I32", Op: uint8(0xA9), T: I32, A: F32},
	{I: TruncF64S_I32, Class: Numeric, Name: "trunc_f64_s.i32", Name2: "TruncF64S_I32", Op: uint8(0xAA), T: I32, A: F64},
	{I: TruncF64U_I32, Class: Numeric, Name: "trunc_f64_u.i32", Name2: "TruncF64U_I32", Op: uint8(0xAB), T: I32, A: F64},
	{I: ExtendI32S_I64, Class: Numeric, Name: "extend_i32_s.i64", Name2: "ExtendI32S_I64", Op: uint8(0xAC), T: I64, A: I32},
	{I: ExtendI32U_I64, Class: Numeric, Name: "extend_i32_u.i64", Name2: "ExtendI32U_I64", Op: uint8(0xAD), T: I64, A: I32},
	{I: TruncF32S_I64, Class: Numeric, Name: "trunc_f32_s.i64", Name2: "TruncF32S_I64", Op: uint8(0xAE), T: I64, A: F32},
	{I: TruncF32U_I64, Class: Numeric, Name: "trunc_f32_u.i64", Name2: "TruncF32U_I64", Op: uint8(0xAF), T: I64, A: F32},
	{I: TruncF64S_I64, Class: Numeric, Name: "trunc_f64_s.i64", Name2: "TruncF64S_I64", Op: uint8(0xB0), T: I64, A: F64},
	{I: TruncF64U_I64, Class: Numeric, Name: "trunc_f64_u.i64", Name2: "TruncF64U_I64", Op: uint8(0xB1), T: I64, A: F64},
	{I: ConvertI32S_F32, Class: Numeric, Name: "convert_i32_s.f32", Name2: "ConvertI32S_F32", Op: uint8(0xB2), T: F32, A: I32},
	{I: ConvertI32U_F32, Class: Numeric, Name: "convert_i32_u.f32", Name2: "ConvertI32U_F32", Op: uint8(0xB3), T: F32, A: I32},
	{I: ConvertI64S_F32, Class: Numeric, Name: "convert_i64_s.f32", Name2: "ConvertI64S_F32", Op: uint8(0xB4), T: F32, A: I64},
	{I: ConvertI64U_F32, Class: Numeric, Name: "convert_i64_u.f32", Name2: "ConvertI64U_F32", Op: uint8(0xB5), T: F32, A: I64},
	{I: DemoteF64_F32, Class: Numeric, Name: "demote_f64.f32", Name2: "DemoteF64_F32", Op: uint8(0xB6), T: F32, A: F64},
	{I: ConvertI32S_F64, Class: Numeric, Name: "convert_i32_s.f64", Name2: "ConvertI32S_F64", Op: uint8(0xB7), T: F64, A: I32},
	{I: ConvertI32U_F64, Class: Numeric, Name: "convert_i32_u.f64", Name2: "ConvertI32U_F64", Op: uint8(0xB8), T: F64, A: I32},
	{I: ConvertI64S_F64, Class: Numeric, Name: "convert_i64_s.f64", Name2: "ConvertI64S_F64", Op: uint8(0xB9), T: F64, A: I64},
	{I: ConvertI64U_F64, Class: Numeric, Name: "convert_i64_u.f64", Name2: "ConvertI64U_F64", Op: uint8(0xBA), T: F64, A: I64},
	{I: PromoteF32_F64, Class: Numeric, Name: "promote_f32.f64", Name2: "PromoteF32_F64", Op: uint8(0xBB), T: F64, A: F32},
	{I: ReinterpretF32_I32, Class: Numeric, Name: "reinterpret_f32.i32", Name2: "ReinterpretF32_I32", Op: uint8(0xBC), T: I32, A: F32},
	{I: ReinterpretF64_I64, Class: Numeric, Name: "reinterpret_f64.i64", Name2: "ReinterpretF64_I64", Op: uint8(0xBD), T: I64, A: F64},
	{I: ReinterpretI32_F32, Class: Numeric, Name: "reinterpret_i32.f32", Name2: "ReinterpretI32_F32", Op: uint8(0xBE), T: F32, A: I32},
	{I: ReinterpretI64_F64, Class: Numeric, Name: "reinterpret_i64.f64", Name2: "ReinterpretI64_F64", Op: uint8(0xBF), T: F64, A: I64},
	{I: Extend8S_I32, Class: Numeric, Name: "extend8_s.i32", Name2: "Extend8S_I32", Op: uint8(0xC0), T: I32, A: I32},
	{I: Extend16S_I32, Class: Numeric, Name: "extend16_s.i32", Name2: "Extend16S_I32", Op: uint8(0xC1), T: I32, A: I32},
	{I: Extend8S_I64, Class: Numeric, Name: "extend8_s.i64", Name2: "Extend8S_I64", Op: uint8(0xC2), T: I64, A: I64},
	{I: Extend16S_I64, Class: Numeric, Name: "extend16_s.i64", Name2: "Extend16S_I64", Op: uint8(0xC3), T: I64, A: I64},
	{I: Extend32S_I64, Class: Numeric, Name: "extend32_s.i64", Name2: "Extend32S_I64", Op: uint8(0xC4), T: I64, A: I64},
	{I: TruncSatF32S_I32, Class: Numeric, Name: "trunc_sat_f32_s.i32", Name2: "TruncSatF32S_I32", OpPfx: uint8(0xFC), Op: uint8(0x0), T: I32, A: F32},
	{I: TruncSatF32U_I32, Class: Numeric, Name: "trunc_sat_f32_u.i32", Name2: "TruncSatF32U_I32", OpPfx: uint8(0xFC), Op: uint8(0x1), T: I32, A: F32},
	{I: TruncSatF64S_I32, Class: Numeric, Name: "trunc_sat_f64_s.i32", Name2: "TruncSatF64S_I32", OpPfx: uint8(0xFC), Op: uint8(0x2), T: I32, A: F64},
	{I: TruncSatF64U_I32, Class: Numeric, Name: "trunc_sat_f64_u.i32", Name2: "TruncSatF64U_I32", OpPfx: uint8(0xFC), Op: uint8(0x3), T: I32, A: F64},
	{I: TruncSatF32S_I64, Class: Numeric, Name: "trunc_sat_f32_s.i64", Name2: "TruncSatF32S_I64", OpPfx: uint8(0xFC), Op: uint8(0x4), T: I64, A: F32},
	{I: TruncSatF32U_I64, Class: Numeric, Name: "trunc_sat_f32_u.i64", Name2: "TruncSatF32U_I64", OpPfx: uint8(0xFC), Op: uint8(0x5), T: I64, A: F32},
	{I: TruncSatF64S_I64, Class: Numeric, Name: "trunc_sat_f64_s.i64", Name2: "TruncSatF64S_I64", OpPfx: uint8(0xFC), Op: uint8(0x6), T: I64, A: F64},
	{I: TruncSatF64U_I64, Class: Numeric, Name: "trunc_sat_f64_u.i64", Name2: "TruncSatF64U_I64", OpPfx: uint8(0xFC), Op: uint8(0x7), T: I64, A: F64},
	{I: Load_V128, Class: Vector, Name: "load.v128", Name2: "Load_V128", OpPfx: uint8(0xFD), Op: uint8(0x0), T: V128, A: I32, Ma: Load, Mz: 16},
	{I: Load8x8S_V128, Class: Vector, Name: "load8x8_s.v128", Name2: "Load8x8S_V128", OpPfx: uint8(0xFD), Op: uint8(0x1), T: V128, A: I32, Ma: Load, Mz: 8},
	{I: Load8x8U_V128, Class: Vector, Name: "load8x8_u.v128", Name2: "Load8x8U_V128", OpPfx: uint8(0xFD), Op: uint8(0x2), T: V128, A: I32, Ma: Load, Mz: 8},
	{I: Load16x4S_V128, Class: Vector, Name: "load16x4_s.v128", Name2: "Load16x4S_V128", OpPfx: uint8(0xFD), Op: uint8(0x3), T: V128, A: I32, Ma: Load, Mz: 8},
	{I: Load16x4U_V128, Class: Vector, Name: "load16x4_u.v128", Name2: "Load16x4U_V128", OpPfx: uint8(0xFD), Op: uint8(0x4), T: V128, A: I32, Ma: Load, Mz: 8},
	{I: Load32x2S_V128, Class: Vector, Name: "load32x2_s.v128", Name2: "Load32x2S_V128", OpPfx: uint8(0xFD), Op: uint8(0x5), T: V128, A: I32, Ma: Load, Mz: 8},
	{I: Load32x2U_V128, Class: Vector, Name: "load32x2_u.v128", Name2: "Load32x2U_V128", OpPfx: uint8(0xFD), Op: uint8(0x6), T: V128, A: I32, Ma: Load, Mz: 8},
	{I: Load8Splat_V128, Class: Vector, Name: "load8_splat.v128", Name2: "Load8Splat_V128", OpPfx: uint8(0xFD), Op: uint8(0x7), T: V128, A: I32, Ma: Load, Mz: 1},
	{I: Load16Splat_V128, Class: Vector, Name: "load16_splat.v128", Name2: "Load16Splat_V128", OpPfx: uint8(0xFD), Op: uint8(0x8), T: V128, A: I32, Ma: Load, Mz: 2},
	{I: Load32Splat_V128, Class: Vector, Name: "load32_splat.v128", Name2: "Load32Splat_V128", OpPfx: uint8(0xFD), Op: uint8(0x9), T: V128, A: I32, Ma: Load, Mz: 4},
	{I: Load64Splat_V128, Class: Vector, Name: "load64_splat.v128", Name2: "Load64Splat_V128", OpPfx: uint8(0xFD), Op: uint8(0xA), T: V128, A: I32, Ma: Load, Mz: 8},
	{I: Load32Zero_V128, Class: Vector, Name: "load32_zero.v128", Name2: "Load32Zero_V128", OpPfx: uint8(0xFD), Op: uint8(0x5C), T: V128, A: I32, Ma: Load, Mz: 4},
	{I: Load64Zero_V128, Class: Vector, Name: "load64_zero.v128", Name2: "Load64Zero_V128", OpPfx: uint8(0xFD), Op: uint8(0x5D), T: V128, A: I32, Ma: Load, Mz: 8},
	{I: Store_V128, Class: Vector, Name: "store.v128", Name2: "Store_V128", OpPfx: uint8(0xFD), Op: uint8(0xB), A: I32, B: V128, Ma: Store, Mz: 16},
	{I: Load8Lane_V128, Class: Vector, Name: "load8_lane.v128", Name2: "Load8Lane_V128", OpPfx: uint8(0xFD), Op: uint8(0x54), T: V128, A: I32, B: V128, Ma: Load, Mz: 1},
	{I: Load16Lane_V128, Class: Vector, Name: "load16_lane.v128", Name2: "Load16Lane_V128", OpPfx: uint8(0xFD), Op: uint8(0x55), T: V128, A: I32, B: V128, Ma: Load, Mz: 2},
	{I: Load32Lane_V128, Class: Vector, Name: "load32_lane.v128", Name2: "Load32Lane_V128", OpPfx: uint8(0xFD), Op: uint8(0x56), T: V128, A: I32, B: V128, Ma: Load, Mz: 4},
	{I: Load64Lane_V128, Class: Vector, Name: "load64_lane.v128", Name2: "Load64Lane_V128", OpPfx: uint8(0xFD), Op: uint8(0x57), T: V128, A: I32, B: V128, Ma: Load, Mz: 8},
	{I: Store8Lane_V128, Class: Vector, Name: "store8_lane.v128", Name2: "Store8Lane_V128", OpPfx: uint8(0xFD), Op: uint8(0x58), A: I32, B: V128, Ma: Store, Mz: 1},
	{I: Store16Lane_V128, Class: Vector, Name: "store16_lane.v128", Name2: "Store16Lane_V128", OpPfx: uint8(0xFD), Op: uint8(0x59), A: I32, B: V128, Ma: Store, Mz: 2},
	{I: Store32Lane_V128, Class: Vector, Name: "store32_lane.v128", Name2: "Store32Lane_V128", OpPfx: uint8(0xFD), Op: uint8(0x5A), A: I32, B: V128, Ma: Store, Mz: 4},
	{I: Store64Lane_V128, Class: Vector, Name: "store64_lane.v128", Name2: "Store64Lane_V128", OpPfx: uint8(0xFD), Op: uint8(0x5B), A: I32, B: V128, Ma: Store, Mz: 8},
	{I: Const_V128, Class: Vector, Name: "const.v128", Name2: "Const_V128", OpPfx: uint8(0xFD), Op: uint8(0xC), T: V128},
	{I: Shuffle_I8x16, Class: Vector, Name: "shuffle.i8x16", Name2: "Shuffle_I8x16", OpPfx: uint8(0xFD), Op: uint8(0xD), T: V128, A: V128, B: V128},
	{I: ExtractLaneS_I8x16, Class: Vector, Name: "extract_lane_s.i8x16", Name2: "ExtractLaneS_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x15), T: I32, A: V128},
	{I: ExtractLaneU_I8x16, Class: Vector, Name: "extract_lane_u.i8x16", Name2: "ExtractLaneU_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x16), T: I32, A: V128},
	{I: ReplaceLane_I8x16, Class: Vector, Name: "replace_lane.i8x16", Name2: "ReplaceLane_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x17), T: V128, A: V128, B: I32},
	{I: ExtractLaneS_I16x8, Class: Vector, Name: "extract_lane_s.i16x8", Name2: "ExtractLaneS_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x18), T: I32, A: V128},
	{I: ExtractLaneU_I16x8, Class: Vector, Name: "extract_lane_u.i16x8", Name2: "ExtractLaneU_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x19), T: I32, A: V128},
	{I: ReplaceLane_I16x8, Class: Vector, Name: "replace_lane.i16x8", Name2: "ReplaceLane_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x1A), T: V128, A: V128, B: I32},
	{I: ExtractLane_I32x4, Class: Vector, Name: "extract_lane.i32x4", Name2: "ExtractLane_I32x4", OpPfx: uint8(0xFD), Op: uint8(0x1B), T: I32, A: V128},
	{I: ReplaceLane_I32x4, Class: Vector, Name: "replace_lane.i32x4", Name2: "ReplaceLane_I32x4", OpPfx: uint8(0xFD), Op: uint8(0x1C), T: V128, A: V128, B: I32},
	{I: ExtractLane_I64x2, Class: Vector, Name: "extract_lane.i64x2", Name2: "ExtractLane_I64x2", OpPfx: uint8(0xFD), Op: uint8(0x1D), T: I64, A: V128},
	{I: ReplaceLane_I64x2, Class: Vector, Name: "replace_lane.i64x2", Name2: "ReplaceLane_I64x2", OpPfx: uint8(0xFD), Op: uint8(0x1E), T: V128, A: V128, B: I64},
	{I: ExtractLane_F32x4, Class: Vector, Name: "extract_lane.f32x4", Name2: "ExtractLane_F32x4", OpPfx: uint8(0xFD), Op: uint8(0x1F), T: F32, A: V128},
	{I: ReplaceLane_F32x4, Class: Vector, Name: "replace_lane.f32x4", Name2: "ReplaceLane_F32x4", OpPfx: uint8(0xFD), Op: uint8(0x20), T: V128, A: V128, B: F32},
	{I: ExtractLane_F64x2, Class: Vector, Name: "extract_lane.f64x2", Name2: "ExtractLane_F64x2", OpPfx: uint8(0xFD), Op: uint8(0x21), T: F64, A: V128},
	{I: ReplaceLane_F64x2, Class: Vector, Name: "replace_lane.f64x2", Name2: "ReplaceLane_F64x2", OpPfx: uint8(0xFD), Op: uint8(0x22), T: V128, A: V128, B: F64},
	{I: Swizzle_I8x16, Class: Vector, Name: "swizzle.i8x16", Name2: "Swizzle_I8x16", OpPfx: uint8(0xFD), Op: uint8(0xE), T: V128, A: V128, B: V128},
	{I: Splat_I8x16, Class: Vector, Name: "splat.i8x16", Name2: "Splat_I8x16", OpPfx: uint8(0xFD), Op: uint8(0xF), T: V128, A: I32},
	{I: Splat_I16x8, Class: Vector, Name: "splat.i16x8", Name2: "Splat_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x10), T: V128, A: I32},
	{I: Splat_I32x4, Class: Vector, Name: "splat.i32x4", Name2: "Splat_I32x4", OpPfx: uint8(0xFD), Op: uint8(0x11), T: V128, A: I32},
	{I: Splat_I64x2, Class: Vector, Name: "splat.i64x2", Name2: "Splat_I64x2", OpPfx: uint8(0xFD), Op: uint8(0x12), T: V128, A: I64},
	{I: Splat_F32x4, Class: Vector, Name: "splat.f32x4", Name2: "Splat_F32x4", OpPfx: uint8(0xFD), Op: uint8(0x13), T: V128, A: F32},
	{I: Splat_F64x2, Class: Vector, Name: "splat.f64x2", Name2: "Splat_F64x2", OpPfx: uint8(0xFD), Op: uint8(0x14), T: V128, A: F64},
	{I: Eq_I8x16, Class: Vector, Name: "eq.i8x16", Name2: "Eq_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x23), T: V128, A: V128, B: V128},
	{I: Ne_I8x16, Class: Vector, Name: "ne.i8x16", Name2: "Ne_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x24), T: V128, A: V128, B: V128},
	{I: LtS_I8x16, Class: Vector, Name: "lt_s.i8x16", Name2: "LtS_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x25), T: V128, A: V128, B: V128},
	{I: LtU_I8x16, Class: Vector, Name: "lt_u.i8x16", Name2: "LtU_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x26), T: V128, A: V128, B: V128},
	{I: GtS_I8x16, Class: Vector, Name: "gt_s.i8x16", Name2: "GtS_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x27), T: V128, A: V128, B: V128},
	{I: GtU_I8x16, Class: Vector, Name: "gt_u.i8x16", Name2: "GtU_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x28), T: V128, A: V128, B: V128},
	{I: LeS_I8x16, Class: Vector, Name: "le_s.i8x16", Name2: "LeS_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x29), T: V128, A: V128, B: V128},
	{I: LeU_I8x16, Class: Vector, Name: "le_u.i8x16", Name2: "LeU_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x2A), T: V128, A: V128, B: V128},
	{I: GeS_I8x16, Class: Vector, Name: "ge_s.i8x16", Name2: "GeS_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x2B), T: V128, A: V128, B: V128},
	{I: GeU_I8x16, Class: Vector, Name: "ge_u.i8x16", Name2: "GeU_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x2C), T: V128, A: V128, B: V128},
	{I: Eq_I16x8, Class: Vector, Name: "eq.i16x8", Name2: "Eq_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x2D), T: V128, A: V128, B: V128},
	{I: Ne_I16x8, Class: Vector, Name: "ne.i16x8", Name2: "Ne_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x2E), T: V128, A: V128, B: V128},
	{I: LtS_I16x8, Class: Vector, Name: "lt_s.i16x8", Name2: "LtS_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x2F), T: V128, A: V128, B: V128},
	{I: LtU_I16x8, Class: Vector, Name: "lt_u.i16x8", Name2: "LtU_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x30), T: V128, A: V128, B: V128},
	{I: GtS_I16x8, Class: Vector, Name: "gt_s.i16x8", Name2: "GtS_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x31), T: V128, A: V128, B: V128},
	{I: GtU_I16x8, Class: Vector, Name: "gt_u.i16x8", Name2: "GtU_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x32), T: V128, A: V128, B: V128},
	{I: LeS_I16x8, Class: Vector, Name: "le_s.i16x8", Name2: "LeS_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x33), T: V128, A: V128, B: V128},
	{I: LeU_I16x8, Class: Vector, Name: "le_u.i16x8", Name2: "LeU_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x34), T: V128, A: V128, B: V128},
	{I: GeS_I16x8, Class: Vector, Name: "ge_s.i16x8", Name2: "GeS_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x35), T: V128, A: V128, B: V128},
	{I: GeU_I16x8, Class: Vector, Name: "ge_u.i16x8", Name2: "GeU_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x36), T: V128, A: V128, B: V128},
	{I: Eq_I32x4, Class: Vector, Name: "eq.i32x4", Name2: "Eq_I32x4", OpPfx: uint8(0xFD), Op: uint8(0x37), T: V128, A: V128, B: V128},
	{I: Ne_I32x4, Class: Vector, Name: "ne.i32x4", Name2: "Ne_I32x4", OpPfx: uint8(0xFD), Op: uint8(0x38), T: V128, A: V128, B: V128},
	{I: LtS_I32x4, Class: Vector, Name: "lt_s.i32x4", Name2: "LtS_I32x4", OpPfx: uint8(0xFD), Op: uint8(0x39), T: V128, A: V128, B: V128},
	{I: LtU_I32x4, Class: Vector, Name: "lt_u.i32x4", Name2: "LtU_I32x4", OpPfx: uint8(0xFD), Op: uint8(0x3A), T: V128, A: V128, B: V128},
	{I: GtS_I32x4, Class: Vector, Name: "gt_s.i32x4", Name2: "GtS_I32x4", OpPfx: uint8(0xFD), Op: uint8(0x3B), T: V128, A: V128, B: V128},
	{I: GtU_I32x4, Class: Vector, Name: "gt_u.i32x4", Name2: "GtU_I32x4", OpPfx: uint8(0xFD), Op: uint8(0x3C), T: V128, A: V128, B: V128},
	{I: LeS_I32x4, Class: Vector, Name: "le_s.i32x4", Name2: "LeS_I32x4", OpPfx: uint8(0xFD), Op: uint8(0x3D), T: V128, A: V128, B: V128},
	{I: LeU_I32x4, Class: Vector, Name: "le_u.i32x4", Name2: "LeU_I32x4", OpPfx: uint8(0xFD), Op: uint8(0x3E), T: V128, A: V128, B: V128},
	{I: GeS_I32x4, Class: Vector, Name: "ge_s.i32x4", Name2: "GeS_I32x4", OpPfx: uint8(0xFD), Op: uint8(0x3F), T: V128, A: V128, B: V128},
	{I: GeU_I32x4, Class: Vector, Name: "ge_u.i32x4", Name2: "GeU_I32x4", OpPfx: uint8(0xFD), Op: uint8(0x40), T: V128, A: V128, B: V128},
	{I: Eq_I64x2, Class: Vector, Name: "eq.i64x2", Name2: "Eq_I64x2", OpPfx: uint8(0xFD), Op: uint8(0xD6), T: V128, A: V128, B: V128},
	{I: Ne_I64x2, Class: Vector, Name: "ne.i64x2", Name2: "Ne_I64x2", OpPfx: uint8(0xFD), Op: uint8(0xD7), T: V128, A: V128, B: V128},
	{I: LtS_I64x2, Class: Vector, Name: "lt_s.i64x2", Name2: "LtS_I64x2", OpPfx: uint8(0xFD), Op: uint8(0xD8), T: V128, A: V128, B: V128},
	{I: GtS_I64x2, Class: Vector, Name: "gt_s.i64x2", Name2: "GtS_I64x2", OpPfx: uint8(0xFD), Op: uint8(0xD9), T: V128, A: V128, B: V128},
	{I: LeS_I64x2, Class: Vector, Name: "le_s.i64x2", Name2: "LeS_I64x2", OpPfx: uint8(0xFD), Op: uint8(0xDA), T: V128, A: V128, B: V128},
	{I: GeS_I64x2, Class: Vector, Name: "ge_s.i64x2", Name2: "GeS_I64x2", OpPfx: uint8(0xFD), Op: uint8(0xDB), T: V128, A: V128, B: V128},
	{I: Eq_F32x4, Class: Vector, Name: "eq.f32x4", Name2: "Eq_F32x4", OpPfx: uint8(0xFD), Op: uint8(0x41), T: V128, A: V128, B: V128},
	{I: Ne_F32x4, Class: Vector, Name: "ne.f32x4", Name2: "Ne_F32x4", OpPfx: uint8(0xFD), Op: uint8(0x42), T: V128, A: V128, B: V128},
	{I: Lt_F32x4, Class: Vector, Name: "lt.f32x4", Name2: "Lt_F32x4", OpPfx: uint8(0xFD), Op: uint8(0x43), T: V128, A: V128, B: V128},
	{I: Gt_F32x4, Class: Vector, Name: "gt.f32x4", Name2: "Gt_F32x4", OpPfx: uint8(0xFD), Op: uint8(0x44), T: V128, A: V128, B: V128},
	{I: Le_F32x4, Class: Vector, Name: "le.f32x4", Name2: "Le_F32x4", OpPfx: uint8(0xFD), Op: uint8(0x45), T: V128, A: V128, B: V128},
	{I: Ge_F32x4, Class: Vector, Name: "ge.f32x4", Name2: "Ge_F32x4", OpPfx: uint8(0xFD), Op: uint8(0x46), T: V128, A: V128, B: V128},
	{I: Eq_F64x2, Class: Vector, Name: "eq.f64x2", Name2: "Eq_F64x2", OpPfx: uint8(0xFD), Op: uint8(0x47), T: V128, A: V128, B: V128},
	{I: Ne_F64x2, Class: Vector, Name: "ne.f64x2", Name2: "Ne_F64x2", OpPfx: uint8(0xFD), Op: uint8(0x48), T: V128, A: V128, B: V128},
	{I: Lt_F64x2, Class: Vector, Name: "lt.f64x2", Name2: "Lt_F64x2", OpPfx: uint8(0xFD), Op: uint8(0x49), T: V128, A: V128, B: V128},
	{I: Gt_F64x2, Class: Vector, Name: "gt.f64x2", Name2: "Gt_F64x2", OpPfx: uint8(0xFD), Op: uint8(0x4A), T: V128, A: V128, B: V128},
	{I: Le_F64x2, Class: Vector, Name: "le.f64x2", Name2: "Le_F64x2", OpPfx: uint8(0xFD), Op: uint8(0x4B), T: V128, A: V128, B: V128},
	{I: Ge_F64x2, Class: Vector, Name: "ge.f64x2", Name2: "Ge_F64x2", OpPfx: uint8(0xFD), Op: uint8(0x4C), T: V128, A: V128, B: V128},
	{I: Not_V128, Class: Vector, Name: "not.v128", Name2: "Not_V128", OpPfx: uint8(0xFD), Op: uint8(0x4D), T: V128, A: V128},
	{I: And_V128, Class: Vector, Name: "and.v128", Name2: "And_V128", OpPfx: uint8(0xFD), Op: uint8(0x4E), T: V128, A: V128, B: V128},
	{I: Andnot_V128, Class: Vector, Name: "andnot.v128", Name2: "Andnot_V128", OpPfx: uint8(0xFD), Op: uint8(0x4F), T: V128, A: V128, B: V128},
	{I: Or_V128, Class: Vector, Name: "or.v128", Name2: "Or_V128", OpPfx: uint8(0xFD), Op: uint8(0x50), T: V128, A: V128, B: V128},
	{I: Xor_V128, Class: Vector, Name: "xor.v128", Name2: "Xor_V128", OpPfx: uint8(0xFD), Op: uint8(0x51), T: V128, A: V128, B: V128},
	{I: Bitselect_V128, Class: Vector, Name: "bitselect.v128", Name2: "Bitselect_V128", OpPfx: uint8(0xFD), Op: uint8(0x52), T: V128, A: V128, B: V128, C: V128},
	{I: AnyTrue_V128, Class: Vector, Name: "any_true.v128", Name2: "AnyTrue_V128", OpPfx: uint8(0xFD), Op: uint8(0x53), T: I32, A: V128},
	{I: Abs_I8x16, Class: Vector, Name: "abs.i8x16", Name2: "Abs_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x60), T: V128, A: V128},
	{I: Neg_I8x16, Class: Vector, Name: "neg.i8x16", Name2: "Neg_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x61), T: V128, A: V128},
	{I: Popcnt_I8x16, Class: Vector, Name: "popcnt.i8x16", Name2: "Popcnt_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x62), T: V128, A: V128},
	{I: AllTrue_I8x16, Class: Vector, Name: "all_true.i8x16", Name2: "AllTrue_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x63), T: I32, A: V128},
	{I: Bitmask_I8x16, Class: Vector, Name: "bitmask.i8x16", Name2: "Bitmask_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x64), T: I32, A: V128},
	{I: NarrowI16x8S_I8x16, Class: Vector, Name: "narrow_i16x8_s.i8x16", Name2: "NarrowI16x8S_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x65), T: V128, A: V128},
	{I: NarrowI16x8U_I8x16, Class: Vector, Name: "narrow_i16x8_u.i8x16", Name2: "NarrowI16x8U_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x66), T: V128, A: V128},
	{I: Shl_I8x16, Class: Vector, Name: "shl.i8x16", Name2: "Shl_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x6B), T: V128, A: V128, B: I32},
	{I: ShrS_I8x16, Class: Vector, Name: "shr_s.i8x16", Name2: "ShrS_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x6C), T: V128, A: V128, B: I32},
	{I: ShrU_I8x16, Class: Vector, Name: "shr_u.i8x16", Name2: "ShrU_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x6D), T: V128, A: V128, B: I32},
	{I: Add_I8x16, Class: Vector, Name: "add.i8x16", Name2: "Add_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x6E), T: V128, A: V128, B: V128},
	{I: AddSatS_I8x16, Class: Vector, Name: "add_sat_s.i8x16", Name2: "AddSatS_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x6F), T: V128, A: V128, B: V128},
	{I: AddSatU_I8x16, Class: Vector, Name: "add_sat_u.i8x16", Name2: "AddSatU_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x70), T: V128, A: V128, B: V128},
	{I: Sub_I8x16, Class: Vector, Name: "sub.i8x16", Name2: "Sub_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x71), T: V128, A: V128, B: V128},
	{I: SubSatS_I8x16, Class: Vector, Name: "sub_sat_s.i8x16", Name2: "SubSatS_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x72), T: V128, A: V128, B: V128},
	{I: SubSatU_I8x16, Class: Vector, Name: "sub_sat_u.i8x16", Name2: "SubSatU_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x73), T: V128, A: V128, B: V128},
	{I: MinS_I8x16, Class: Vector, Name: "min_s.i8x16", Name2: "MinS_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x76), T: V128, A: V128, B: V128},
	{I: MinU_I8x16, Class: Vector, Name: "min_u.i8x16", Name2: "MinU_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x77), T: V128, A: V128, B: V128},
	{I: MaxS_I8x16, Class: Vector, Name: "max_s.i8x16", Name2: "MaxS_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x78), T: V128, A: V128, B: V128},
	{I: MaxU_I8x16, Class: Vector, Name: "max_u.i8x16", Name2: "MaxU_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x79), T: V128, A: V128, B: V128},
	{I: AvgrU_I8x16, Class: Vector, Name: "avgr_u.i8x16", Name2: "AvgrU_I8x16", OpPfx: uint8(0xFD), Op: uint8(0x7B), T: V128, A: V128, B: V128},
	{I: ExtaddPairwiseI8x16S_I16x8, Class: Vector, Name: "extadd_pairwise_i8x16_s.i16x8", Name2: "ExtaddPairwiseI8x16S_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x7C), T: V128, A: V128},
	{I: ExtaddPairwiseI8x16U_I16x8, Class: Vector, Name: "extadd_pairwise_i8x16_u.i16x8", Name2: "ExtaddPairwiseI8x16U_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x7D), T: V128, A: V128},
	{I: Abs_I16x8, Class: Vector, Name: "abs.i16x8", Name2: "Abs_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x80), T: V128, A: V128},
	{I: Neg_I16x8, Class: Vector, Name: "neg.i16x8", Name2: "Neg_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x81), T: V128, A: V128},
	{I: Q15mulrSatS_I16x8, Class: Vector, Name: "q15mulr_sat_s.i16x8", Name2: "Q15mulrSatS_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x82), T: V128, A: V128, B: V128},
	{I: Alltrue_I16x8, Class: Vector, Name: "alltrue.i16x8", Name2: "Alltrue_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x83), T: I32, A: V128},
	{I: Bitmask_I16x8, Class: Vector, Name: "bitmask.i16x8", Name2: "Bitmask_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x84), T: I32, A: V128},
	{I: NarrowI32x4S_I16x8, Class: Vector, Name: "narrow_i32x4_s.i16x8", Name2: "NarrowI32x4S_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x85), T: V128, A: V128, B: V128},
	{I: NarrowI32x4U_I16x8, Class: Vector, Name: "narrow_i32x4_u.i16x8", Name2: "NarrowI32x4U_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x86), T: V128, A: V128, B: V128},
	{I: ExtendLowI8x16S_I16x8, Class: Vector, Name: "extend_low_i8x16_s.i16x8", Name2: "ExtendLowI8x16S_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x87), T: V128, A: V128},
	{I: ExtendHighI8x16S_I16x8, Class: Vector, Name: "extend_high_i8x16_s.i16x8", Name2: "ExtendHighI8x16S_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x88), T: V128, A: V128},
	{I: ExtendLowI8x16U_I16x8, Class: Vector, Name: "extend_low_i8x16_u.i16x8", Name2: "ExtendLowI8x16U_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x89), T: V128, A: V128},
	{I: ExtendHighI8x16U_I16x8, Class: Vector, Name: "extend_high_i8x16_u.i16x8", Name2: "ExtendHighI8x16U_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x8A), T: V128, A: V128},
	{I: Shl_I16x8, Class: Vector, Name: "shl.i16x8", Name2: "Shl_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x8B), T: V128, A: V128, B: I32},
	{I: ShrS_I16x8, Class: Vector, Name: "shr_s.i16x8", Name2: "ShrS_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x8C), T: V128, A: V128, B: I32},
	{I: ShrU_I16x8, Class: Vector, Name: "shr_u.i16x8", Name2: "ShrU_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x8D), T: V128, A: V128, B: I32},
	{I: Add_I16x8, Class: Vector, Name: "add.i16x8", Name2: "Add_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x8E), T: I32, A: V128, B: V128},
	{I: AddSatS_I16x8, Class: Vector, Name: "add_sat_s.i16x8", Name2: "AddSatS_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x8F), T: I32, A: V128, B: V128},
	{I: AddSatU_I16x8, Class: Vector, Name: "add_sat_u.i16x8", Name2: "AddSatU_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x90), T: I32, A: V128, B: V128},
	{I: Sub_I16x8, Class: Vector, Name: "sub.i16x8", Name2: "Sub_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x91), T: I32, A: V128, B: V128},
	{I: SubSatS_I16x8, Class: Vector, Name: "sub_sat_s.i16x8", Name2: "SubSatS_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x92), T: I32, A: V128, B: V128},
	{I: SubSatU_I16x8, Class: Vector, Name: "sub_sat_u.i16x8", Name2: "SubSatU_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x93), T: I32, A: V128, B: V128},
	{I: Mul_I16x8, Class: Vector, Name: "mul.i16x8", Name2: "Mul_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x95), T: I32, A: V128, B: V128},
	{I: MinS_I16x8, Class: Vector, Name: "min_s.i16x8", Name2: "MinS_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x96), T: I32, A: V128, B: V128},
	{I: MinU_I16x8, Class: Vector, Name: "min_u.i16x8", Name2: "MinU_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x97), T: I32, A: V128, B: V128},
	{I: MaxS_I16x8, Class: Vector, Name: "max_s.i16x8", Name2: "MaxS_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x98), T: I32, A: V128, B: V128},
	{I: MaxU_I16x8, Class: Vector, Name: "max_u.i16x8", Name2: "MaxU_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x99), T: I32, A: V128, B: V128},
	{I: AvgrU_I16x8, Class: Vector, Name: "avgr_u.i16x8", Name2: "AvgrU_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x9B), T: I32, A: V128, B: V128},
	{I: ExtmulLowI8x16S_I16x8, Class: Vector, Name: "extmul_low_i8x16_s.i16x8", Name2: "ExtmulLowI8x16S_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x9C), T: I32, A: V128, B: V128},
	{I: ExtmulHighI8x16S_I16x8, Class: Vector, Name: "extmul_high_i8x16_s.i16x8", Name2: "ExtmulHighI8x16S_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x9D), T: I32, A: V128, B: V128},
	{I: ExtmulLowI8x16U_I16x8, Class: Vector, Name: "extmul_low_i8x16_u.i16x8", Name2: "ExtmulLowI8x16U_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x9E), T: I32, A: V128, B: V128},
	{I: ExtmulHighI8x16U_I16x8, Class: Vector, Name: "extmul_high_i8x16_u.i16x8", Name2: "ExtmulHighI8x16U_I16x8", OpPfx: uint8(0xFD), Op: uint8(0x9F), T: I32, A: V128, B: V128},
	{I: ExtaddPairwiseI16x8S_I32x4, Class: Vector, Name: "extadd_pairwise_i16x8_s.i32x4", Name2: "ExtaddPairwiseI16x8S_I32x4", OpPfx: uint8(0xFD), Op: uint8(0x7E), T: V128, A: V128},
	{I: ExtaddPairwiseI16x8U_I32x4, Class: Vector, Name: "extadd_pairwise_i16x8_u.i32x4", Name2: "ExtaddPairwiseI16x8U_I32x4", OpPfx: uint8(0xFD), Op: uint8(0x7F), T: V128, A: V128},
	{I: Abs_I32x4, Class: Vector, Name: "abs.i32x4", Name2: "Abs_I32x4", OpPfx: uint8(0xFD), Op: uint8(0xA0), T: I32, A: V128},
	{I: Neg_I32x4, Class: Vector, Name: "neg.i32x4", Name2: "Neg_I32x4", OpPfx: uint8(0xFD), Op: uint8(0xA1), T: I32, A: V128},
	{I: Alltrue_I32x4, Class: Vector, Name: "alltrue.i32x4", Name2: "Alltrue_I32x4", OpPfx: uint8(0xFD), Op: uint8(0xA3), T: I32, A: V128},
	{I: Bitmask_I32x4, Class: Vector, Name: "bitmask.i32x4", Name2: "Bitmask_I32x4", OpPfx: uint8(0xFD), Op: uint8(0xA4), T: I32, A: V128},
	{I: ExtendLowI16x8S_I32x4, Class: Vector, Name: "extend_low_i16x8_s.i32x4", Name2: "ExtendLowI16x8S_I32x4", OpPfx: uint8(0xFD), Op: uint8(0xA7), T: V128, A: V128},
	{I: ExtendHighI16x8S_I32x4, Class: Vector, Name: "extend_high_i16x8_s.i32x4", Name2: "ExtendHighI16x8S_I32x4", OpPfx: uint8(0xFD), Op: uint8(0xA8), T: V128, A: V128},
	{I: ExtendLowI16x8U_I32x4, Class: Vector, Name: "extend_low_i16x8_u.i32x4", Name2: "ExtendLowI16x8U_I32x4", OpPfx: uint8(0xFD), Op: uint8(0xA9), T: V128, A: V128},
	{I: ExtendHighI16x8U_I32x4, Class: Vector, Name: "extend_high_i16x8_u.i32x4", Name2: "ExtendHighI16x8U_I32x4", OpPfx: uint8(0xFD), Op: uint8(0xAA), T: V128, A: V128},
	{I: Shl_I32x4, Class: Vector, Name: "shl.i32x4", Name2: "Shl_I32x4", OpPfx: uint8(0xFD), Op: uint8(0xAB), T: V128, A: V128, B: I32},
	{I: ShrS_I32x4, Class: Vector, Name: "shr_s.i32x4", Name2: "ShrS_I32x4", OpPfx: uint8(0xFD), Op: uint8(0xAC), T: V128, A: V128, B: I32},
	{I: ShrU_I32x4, Class: Vector, Name: "shr_u.i32x4", Name2: "ShrU_I32x4", OpPfx: uint8(0xFD), Op: uint8(0xAD), T: V128, A: V128, B: I32},
	{I: Add_I32x4, Class: Vector, Name: "add.i32x4", Name2: "Add_I32x4", OpPfx: uint8(0xFD), Op: uint8(0xAE), T: V128, A: V128, B: V128},
	{I: Sub_I32x4, Class: Vector, Name: "sub.i32x4", Name2: "Sub_I32x4", OpPfx: uint8(0xFD), Op: uint8(0xB1), T: V128, A: V128, B: V128},
	{I: Mul_I32x4, Class: Vector, Name: "mul.i32x4", Name2: "Mul_I32x4", OpPfx: uint8(0xFD), Op: uint8(0xB5), T: V128, A: V128, B: V128},
	{I: MinS_I32x4, Class: Vector, Name: "min_s.i32x4", Name2: "MinS_I32x4", OpPfx: uint8(0xFD), Op: uint8(0xB6), T: V128, A: V128, B: V128},
	{I: MinU_I32x4, Class: Vector, Name: "min_u.i32x4", Name2: "MinU_I32x4", OpPfx: uint8(0xFD), Op: uint8(0xB7), T: V128, A: V128, B: V128},
	{I: MaxS_I32x4, Class: Vector, Name: "max_s.i32x4", Name2: "MaxS_I32x4", OpPfx: uint8(0xFD), Op: uint8(0xB8), T: V128, A: V128, B: V128},
	{I: MaxU_I32x4, Class: Vector, Name: "max_u.i32x4", Name2: "MaxU_I32x4", OpPfx: uint8(0xFD), Op: uint8(0xB9), T: V128, A: V128, B: V128},
	{I: DotI16x8S_I32x4, Class: Vector, Name: "dot_i16x8_s.i32x4", Name2: "DotI16x8S_I32x4", OpPfx: uint8(0xFD), Op: uint8(0xBA), T: V128, A: V128, B: V128},
	{I: ExtmulLowI16x8S_I32x4, Class: Vector, Name: "extmul_low_i16x8_s.i32x4", Name2: "ExtmulLowI16x8S_I32x4", OpPfx: uint8(0xFD), Op: uint8(0xBC), T: V128, A: V128, B: V128},
	{I: ExtmulHighI16x8S_I32x4, Class: Vector, Name: "extmul_high_i16x8_s.i32x4", Name2: "ExtmulHighI16x8S_I32x4", OpPfx: uint8(0xFD), Op: uint8(0xBD), T: V128, A: V128, B: V128},
	{I: ExtmulLowI16x8U_I32x4, Class: Vector, Name: "extmul_low_i16x8_u.i32x4", Name2: "ExtmulLowI16x8U_I32x4", OpPfx: uint8(0xFD), Op: uint8(0xBE), T: V128, A: V128, B: V128},
	{I: ExtmulHighI16x8U_I32x4, Class: Vector, Name: "extmul_high_i16x8_u.i32x4", Name2: "ExtmulHighI16x8U_I32x4", OpPfx: uint8(0xFD), Op: uint8(0xBF), T: V128, A: V128, B: V128},
	{I: Abs_I64x2, Class: Vector, Name: "abs.i64x2", Name2: "Abs_I64x2", OpPfx: uint8(0xFD), Op: uint8(0xC0), T: V128, A: V128},
	{I: Neg_I64x2, Class: Vector, Name: "neg.i64x2", Name2: "Neg_I64x2", OpPfx: uint8(0xFD), Op: uint8(0xC1), T: V128, A: V128},
	{I: Alltrue_I64x2, Class: Vector, Name: "alltrue.i64x2", Name2: "Alltrue_I64x2", OpPfx: uint8(0xFD), Op: uint8(0xC3), T: I32, A: V128},
	{I: Bitmask_I64x2, Class: Vector, Name: "bitmask.i64x2", Name2: "Bitmask_I64x2", OpPfx: uint8(0xFD), Op: uint8(0xC4), T: I32, A: V128},
	{I: ExtendLowI32x4S_I64x2, Class: Vector, Name: "extend_low_i32x4_s.i64x2", Name2: "ExtendLowI32x4S_I64x2", OpPfx: uint8(0xFD), Op: uint8(0xC7), T: V128, A: V128},
	{I: ExtendHighI32x4S_I64x2, Class: Vector, Name: "extend_high_i32x4_s.i64x2", Name2: "ExtendHighI32x4S_I64x2", OpPfx: uint8(0xFD), Op: uint8(0xC8), T: V128, A: V128},
	{I: ExtendLowI32x4U_I64x2, Class: Vector, Name: "extend_low_i32x4_u.i64x2", Name2: "ExtendLowI32x4U_I64x2", OpPfx: uint8(0xFD), Op: uint8(0xC9), T: V128, A: V128},
	{I: ExtendHighI32x4U_I64x2, Class: Vector, Name: "extend_high_i32x4_u.i64x2", Name2: "ExtendHighI32x4U_I64x2", OpPfx: uint8(0xFD), Op: uint8(0xCA), T: V128, A: V128},
	{I: Shl_I64x2, Class: Vector, Name: "shl.i64x2", Name2: "Shl_I64x2", OpPfx: uint8(0xFD), Op: uint8(0xCB), T: V128, A: V128, B: I32},
	{I: ShrS_I64x2, Class: Vector, Name: "shr_s.i64x2", Name2: "ShrS_I64x2", OpPfx: uint8(0xFD), Op: uint8(0xCC), T: V128, A: V128, B: I32},
	{I: ShrU_I64x2, Class: Vector, Name: "shr_u.i64x2", Name2: "ShrU_I64x2", OpPfx: uint8(0xFD), Op: uint8(0xCD), T: V128, A: V128, B: I32},
	{I: Add_I64x2, Class: Vector, Name: "add.i64x2", Name2: "Add_I64x2", OpPfx: uint8(0xFD), Op: uint8(0xCE), T: V128, A: V128, B: V128},
	{I: Sub_I64x2, Class: Vector, Name: "sub.i64x2", Name2: "Sub_I64x2", OpPfx: uint8(0xFD), Op: uint8(0xD1), T: V128, A: V128, B: V128},
	{I: Mul_I64x2, Class: Vector, Name: "mul.i64x2", Name2: "Mul_I64x2", OpPfx: uint8(0xFD), Op: uint8(0xD5), T: V128, A: V128, B: V128},
	{I: ExtmulLowI32x4S_I64x2, Class: Vector, Name: "extmul_low_i32x4_s.i64x2", Name2: "ExtmulLowI32x4S_I64x2", OpPfx: uint8(0xFD), Op: uint8(0xDC), T: V128, A: V128, B: V128},
	{I: ExtmulHighI32x4S_I64x2, Class: Vector, Name: "extmul_high_i32x4_s.i64x2", Name2: "ExtmulHighI32x4S_I64x2", OpPfx: uint8(0xFD), Op: uint8(0xDD), T: V128, A: V128, B: V128},
	{I: ExtmulLowI32x4U_I64x2, Class: Vector, Name: "extmul_low_i32x4_u.i64x2", Name2: "ExtmulLowI32x4U_I64x2", OpPfx: uint8(0xFD), Op: uint8(0xDE), T: V128, A: V128, B: V128},
	{I: ExtmulHighI32x4U_I64x2, Class: Vector, Name: "extmul_high_i32x4_u.i64x2", Name2: "ExtmulHighI32x4U_I64x2", OpPfx: uint8(0xFD), Op: uint8(0xDF), T: V128, A: V128, B: V128},
	{I: Ceil_F32x4, Class: Vector, Name: "ceil.f32x4", Name2: "Ceil_F32x4", OpPfx: uint8(0xFD), Op: uint8(0x67), T: V128, A: V128},
	{I: Floor_F32x4, Class: Vector, Name: "floor.f32x4", Name2: "Floor_F32x4", OpPfx: uint8(0xFD), Op: uint8(0x68), T: V128, A: V128},
	{I: Trunc_F32x4, Class: Vector, Name: "trunc.f32x4", Name2: "Trunc_F32x4", OpPfx: uint8(0xFD), Op: uint8(0x69), T: V128, A: V128},
	{I: Nearest_F32x4, Class: Vector, Name: "nearest.f32x4", Name2: "Nearest_F32x4", OpPfx: uint8(0xFD), Op: uint8(0x6A), T: V128, A: V128},
	{I: Abs_F32x4, Class: Vector, Name: "abs.f32x4", Name2: "Abs_F32x4", OpPfx: uint8(0xFD), Op: uint8(0xE0), T: V128, A: V128},
	{I: Neg_F32x4, Class: Vector, Name: "neg.f32x4", Name2: "Neg_F32x4", OpPfx: uint8(0xFD), Op: uint8(0xE1), T: V128, A: V128},
	{I: Sqrt_F32x4, Class: Vector, Name: "sqrt.f32x4", Name2: "Sqrt_F32x4", OpPfx: uint8(0xFD), Op: uint8(0xE3), T: V128, A: V128},
	{I: Add_F32x4, Class: Vector, Name: "add.f32x4", Name2: "Add_F32x4", OpPfx: uint8(0xFD), Op: uint8(0xE4), T: V128, A: V128, B: V128},
	{I: Sub_F32x4, Class: Vector, Name: "sub.f32x4", Name2: "Sub_F32x4", OpPfx: uint8(0xFD), Op: uint8(0xE5), T: V128, A: V128, B: V128},
	{I: Mul_F32x4, Class: Vector, Name: "mul.f32x4", Name2: "Mul_F32x4", OpPfx: uint8(0xFD), Op: uint8(0xE6), T: V128, A: V128, B: V128},
	{I: Div_F32x4, Class: Vector, Name: "div.f32x4", Name2: "Div_F32x4", OpPfx: uint8(0xFD), Op: uint8(0xE7), T: V128, A: V128, B: V128},
	{I: Min_F32x4, Class: Vector, Name: "min.f32x4", Name2: "Min_F32x4", OpPfx: uint8(0xFD), Op: uint8(0xE8), T: V128, A: V128, B: V128},
	{I: Max_F32x4, Class: Vector, Name: "max.f32x4", Name2: "Max_F32x4", OpPfx: uint8(0xFD), Op: uint8(0xE9), T: V128, A: V128, B: V128},
	{I: Pmin_F32x4, Class: Vector, Name: "pmin.f32x4", Name2: "Pmin_F32x4", OpPfx: uint8(0xFD), Op: uint8(0xEA), T: V128, A: V128, B: V128},
	{I: Pmax_F32x4, Class: Vector, Name: "pmax.f32x4", Name2: "Pmax_F32x4", OpPfx: uint8(0xFD), Op: uint8(0xEB), T: V128, A: V128, B: V128},
	{I: Ceil_F64x2, Class: Vector, Name: "ceil.f64x2", Name2: "Ceil_F64x2", OpPfx: uint8(0xFD), Op: uint8(0x74), T: V128, A: V128},
	{I: Floor_F64x2, Class: Vector, Name: "floor.f64x2", Name2: "Floor_F64x2", OpPfx: uint8(0xFD), Op: uint8(0x75), T: V128, A: V128},
	{I: Trunc_F64x2, Class: Vector, Name: "trunc.f64x2", Name2: "Trunc_F64x2", OpPfx: uint8(0xFD), Op: uint8(0x7A), T: V128, A: V128},
	{I: Nearest_F64x2, Class: Vector, Name: "nearest.f64x2", Name2: "Nearest_F64x2", OpPfx: uint8(0xFD), Op: uint8(0x94), T: V128, A: V128},
	{I: Abs_F64x2, Class: Vector, Name: "abs.f64x2", Name2: "Abs_F64x2", OpPfx: uint8(0xFD), Op: uint8(0xEC), T: V128, A: V128},
	{I: Neg_F64x2, Class: Vector, Name: "neg.f64x2", Name2: "Neg_F64x2", OpPfx: uint8(0xFD), Op: uint8(0xED), T: V128, A: V128},
	{I: Sqrt_F64x2, Class: Vector, Name: "sqrt.f64x2", Name2: "Sqrt_F64x2", OpPfx: uint8(0xFD), Op: uint8(0xEF), T: V128, A: V128},
	{I: Add_F64x2, Class: Vector, Name: "add.f64x2", Name2: "Add_F64x2", OpPfx: uint8(0xFD), Op: uint8(0xF0), T: V128, A: V128, B: V128},
	{I: Sub_F64x2, Class: Vector, Name: "sub.f64x2", Name2: "Sub_F64x2", OpPfx: uint8(0xFD), Op: uint8(0xF1), T: V128, A: V128, B: V128},
	{I: Mul_F64x2, Class: Vector, Name: "mul.f64x2", Name2: "Mul_F64x2", OpPfx: uint8(0xFD), Op: uint8(0xF2), T: V128, A: V128, B: V128},
	{I: Div_F64x2, Class: Vector, Name: "div.f64x2", Name2: "Div_F64x2", OpPfx: uint8(0xFD), Op: uint8(0xF3), T: V128, A: V128, B: V128},
	{I: Min_F64x2, Class: Vector, Name: "min.f64x2", Name2: "Min_F64x2", OpPfx: uint8(0xFD), Op: uint8(0xF4), T: V128, A: V128, B: V128},
	{I: Max_F64x2, Class: Vector, Name: "max.f64x2", Name2: "Max_F64x2", OpPfx: uint8(0xFD), Op: uint8(0xF5), T: V128, A: V128, B: V128},
	{I: Pmin_F64x2, Class: Vector, Name: "pmin.f64x2", Name2: "Pmin_F64x2", OpPfx: uint8(0xFD), Op: uint8(0xF6), T: V128, A: V128, B: V128},
	{I: Pmax_F64x2, Class: Vector, Name: "pmax.f64x2", Name2: "Pmax_F64x2", OpPfx: uint8(0xFD), Op: uint8(0xF7), T: V128, A: V128, B: V128},
	{I: TruncSatF32x4S_I32x4, Class: Vector, Name: "trunc_sat_f32x4_s.i32x4", Name2: "TruncSatF32x4S_I32x4", OpPfx: uint8(0xFD), Op: uint8(0xF8), T: V128, A: V128},
	{I: TruncSatF32x4U_I32x4, Class: Vector, Name: "trunc_sat_f32x4_u.i32x4", Name2: "TruncSatF32x4U_I32x4", OpPfx: uint8(0xFD), Op: uint8(0xF9), T: V128, A: V128},
	{I: ConvertI32x4S_F32x4, Class: Vector, Name: "convert_i32x4_s.f32x4", Name2: "ConvertI32x4S_F32x4", OpPfx: uint8(0xFD), Op: uint8(0xFA), T: V128, A: V128},
	{I: ConvertI32x4U_F32x4, Class: Vector, Name: "convert_i32x4_u.f32x4", Name2: "ConvertI32x4U_F32x4", OpPfx: uint8(0xFD), Op: uint8(0xFB), T: V128, A: V128},
	{I: TruncSatF64x2SZero_I32x4, Class: Vector, Name: "trunc_sat_f64x2_s_zero.i32x4", Name2: "TruncSatF64x2SZero_I32x4", OpPfx: uint8(0xFD), Op: uint8(0xFC), T: V128, A: V128},
	{I: TruncSatF64x2UZero_I32x4, Class: Vector, Name: "trunc_sat_f64x2_u_zero.i32x4", Name2: "TruncSatF64x2UZero_I32x4", OpPfx: uint8(0xFD), Op: uint8(0xFD), T: V128, A: V128},
	{I: ConvertLowI32x4S_F64x2, Class: Vector, Name: "convert_low_i32x4_s.f64x2", Name2: "ConvertLowI32x4S_F64x2", OpPfx: uint8(0xFD), Op: uint8(0xFE), T: V128, A: V128},
	{I: ConvertLowI32x4U_F64x2, Class: Vector, Name: "convert_low_i32x4_u.f64x2", Name2: "ConvertLowI32x4U_F64x2", OpPfx: uint8(0xFD), Op: uint8(0xFF), T: V128, A: V128},
	{I: DemoteF64x2Zero_F32x4, Class: Vector, Name: "demote_f64x2_zero.f32x4", Name2: "DemoteF64x2Zero_F32x4", OpPfx: uint8(0xFD), Op: uint8(0x5E), T: V128, A: V128},
	{I: PromoteLowF32x4_F64x2, Class: Vector, Name: "promote_low_f32x4.f64x2", Name2: "PromoteLowF32x4_F64x2", OpPfx: uint8(0xFD), Op: uint8(0x5F), T: V128, A: V128},
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
