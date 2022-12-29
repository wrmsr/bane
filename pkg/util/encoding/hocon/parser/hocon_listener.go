// Code generated from Hocon.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Hocon

import antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

// HoconListener is a complete listener for a parse tree produced by HoconParser.
type HoconListener interface {
	antlr.ParseTreeListener

	// EnterHocon is called when entering the hocon production.
	EnterHocon(c *HoconContext)

	// EnterProp is called when entering the prop production.
	EnterProp(c *PropContext)

	// EnterObj is called when entering the obj production.
	EnterObj(c *ObjContext)

	// EnterObjectBegin is called when entering the objectBegin production.
	EnterObjectBegin(c *ObjectBeginContext)

	// EnterObjectEnd is called when entering the objectEnd production.
	EnterObjectEnd(c *ObjectEndContext)

	// EnterObjectData is called when entering the objectData production.
	EnterObjectData(c *ObjectDataContext)

	// EnterArrayData is called when entering the arrayData production.
	EnterArrayData(c *ArrayDataContext)

	// EnterStringData is called when entering the stringData production.
	EnterStringData(c *StringDataContext)

	// EnterReferenceData is called when entering the referenceData production.
	EnterReferenceData(c *ReferenceDataContext)

	// EnterNumberData is called when entering the numberData production.
	EnterNumberData(c *NumberDataContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// EnterPath is called when entering the path production.
	EnterPath(c *PathContext)

	// EnterArrayBegin is called when entering the arrayBegin production.
	EnterArrayBegin(c *ArrayBeginContext)

	// EnterArrayEnd is called when entering the arrayEnd production.
	EnterArrayEnd(c *ArrayEndContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterArrayValue is called when entering the arrayValue production.
	EnterArrayValue(c *ArrayValueContext)

	// EnterArrayString is called when entering the arrayString production.
	EnterArrayString(c *ArrayStringContext)

	// EnterArrayReference is called when entering the arrayReference production.
	EnterArrayReference(c *ArrayReferenceContext)

	// EnterArrayNumber is called when entering the arrayNumber production.
	EnterArrayNumber(c *ArrayNumberContext)

	// EnterArrayObj is called when entering the arrayObj production.
	EnterArrayObj(c *ArrayObjContext)

	// EnterArrayArray is called when entering the arrayArray production.
	EnterArrayArray(c *ArrayArrayContext)

	// EnterV_string is called when entering the v_string production.
	EnterV_string(c *V_stringContext)

	// EnterV_rawstring is called when entering the v_rawstring production.
	EnterV_rawstring(c *V_rawstringContext)

	// EnterV_reference is called when entering the v_reference production.
	EnterV_reference(c *V_referenceContext)

	// EnterRawstring is called when entering the rawstring production.
	EnterRawstring(c *RawstringContext)

	// ExitHocon is called when exiting the hocon production.
	ExitHocon(c *HoconContext)

	// ExitProp is called when exiting the prop production.
	ExitProp(c *PropContext)

	// ExitObj is called when exiting the obj production.
	ExitObj(c *ObjContext)

	// ExitObjectBegin is called when exiting the objectBegin production.
	ExitObjectBegin(c *ObjectBeginContext)

	// ExitObjectEnd is called when exiting the objectEnd production.
	ExitObjectEnd(c *ObjectEndContext)

	// ExitObjectData is called when exiting the objectData production.
	ExitObjectData(c *ObjectDataContext)

	// ExitArrayData is called when exiting the arrayData production.
	ExitArrayData(c *ArrayDataContext)

	// ExitStringData is called when exiting the stringData production.
	ExitStringData(c *StringDataContext)

	// ExitReferenceData is called when exiting the referenceData production.
	ExitReferenceData(c *ReferenceDataContext)

	// ExitNumberData is called when exiting the numberData production.
	ExitNumberData(c *NumberDataContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)

	// ExitPath is called when exiting the path production.
	ExitPath(c *PathContext)

	// ExitArrayBegin is called when exiting the arrayBegin production.
	ExitArrayBegin(c *ArrayBeginContext)

	// ExitArrayEnd is called when exiting the arrayEnd production.
	ExitArrayEnd(c *ArrayEndContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitArrayValue is called when exiting the arrayValue production.
	ExitArrayValue(c *ArrayValueContext)

	// ExitArrayString is called when exiting the arrayString production.
	ExitArrayString(c *ArrayStringContext)

	// ExitArrayReference is called when exiting the arrayReference production.
	ExitArrayReference(c *ArrayReferenceContext)

	// ExitArrayNumber is called when exiting the arrayNumber production.
	ExitArrayNumber(c *ArrayNumberContext)

	// ExitArrayObj is called when exiting the arrayObj production.
	ExitArrayObj(c *ArrayObjContext)

	// ExitArrayArray is called when exiting the arrayArray production.
	ExitArrayArray(c *ArrayArrayContext)

	// ExitV_string is called when exiting the v_string production.
	ExitV_string(c *V_stringContext)

	// ExitV_rawstring is called when exiting the v_rawstring production.
	ExitV_rawstring(c *V_rawstringContext)

	// ExitV_reference is called when exiting the v_reference production.
	ExitV_reference(c *V_referenceContext)

	// ExitRawstring is called when exiting the rawstring production.
	ExitRawstring(c *RawstringContext)
}
