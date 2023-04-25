package consts

/*
https://en.wikipedia.org/wiki/LEB128
https://en.wikipedia.org/wiki/LEB128#Signed_LEB128
*/

const (
	// numtype

	I32 = 0x7F
	I64 = 0x7E
	F32 = 0x7D
	F64 = 0x7C

	// vectype

	V128 = 0x7B
	I8   = 0x7A
	I16  = 0x79

	// reftype

	FuncRef   = 0x70
	ExternRef = 0x6F

	FuncType = 0x60

	MinLimit = 0x00
	MaxLimit = 0x01

	ConstGlobalType = 0x00
	VarGlobalType   = 0x01

	BlockType = 0x40

	// prefixes

	MathPrefix = 0xFC
	SimdPrefix = 0xFD

	// sections

	CustomSection    = 0
	TypeSection      = 1
	ImportSection    = 2
	FunctionSection  = 3
	TableSection     = 4
	MemorySection    = 5
	GlobalSection    = 6
	ExportSection    = 7
	StartSection     = 8
	ElementSection   = 9
	CodeSection      = 10
	DataSection      = 11
	DataCountSection = 12

	FuncImport   = 0x00
	TableImport  = 0x01
	MemImport    = 0x02
	GlobalImport = 0x03

	FuncExport   = 0x00
	TableExport  = 0x01
	MemExport    = 0x01
	GlobalExport = 0x03

	/*
	   elemsec ::= seg* :section9(vec(elem)) ⇒ seg*
	   elem ::=
	     0:u32 𝑒:expr 𝑦 * :vec(funcidx) ⇒ {type funcref, init ((ref.func 𝑦) end) * , mode active {table 0, offset 𝑒}}
	   | 1:u32 et : elemkind 𝑦 * :vec(funcidx) ⇒ {type et, init ((ref.func 𝑦) end) * , mode passive}
	   | 2:u32 𝑥:tableidx 𝑒:expr et : elemkind 𝑦 * :vec(funcidx) ⇒ {type et, init ((ref.func 𝑦) end) * , mode active {table 𝑥, offset 𝑒}}
	   | 3:u32 et : elemkind 𝑦 * :vec(funcidx) ⇒ {type et, init ((ref.func 𝑦) end) * , mode declarative}
	   | 4:u32 𝑒:expr el * :vec(expr) ⇒ {type funcref, init el * , mode active {table 0, offset 𝑒}}
	   | 5:u32 et : reftype el * :vec(expr) ⇒ {type 𝑒𝑡, init el * , mode passive}
	   | 6:u32 𝑥:tableidx 𝑒:expr et : reftype el * :vec(expr) ⇒ {type 𝑒𝑡, init el * , mode active {table 𝑥, offset 𝑒}}
	   | 7:u32 et : reftype el * :vec(expr) ⇒ {type 𝑒𝑡, init el * , mode declarative}
	   elemkind ::= 0x00 ⇒ funcref

	   datasec ::= seg* :section11(vec(data)) ⇒ seg*
	   data ::=
	     0:u32 𝑒:expr 𝑏 * :vec(byte) ⇒ {init 𝑏 * , mode active {memory 0, offset 𝑒}}
	   | 1:u32 𝑏 * :vec(byte) ⇒ {init 𝑏 * , mode passive}
	   | 2:u32 𝑥:memidx 𝑒:expr 𝑏 * :vec(byte) ⇒ {init 𝑏 * , mode active {memory 𝑥, offset 𝑒}}
	*/

	// special

	Magic   = 0x6D_73_61_00
	Version = 0x00_00_00_01
)
