// Code generated from Minml.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Minml

import antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

// MinmlListener is a complete listener for a parse tree produced by MinmlParser.
type MinmlListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterObj is called when entering the obj production.
	EnterObj(c *ObjContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterStringValue is called when entering the stringValue production.
	EnterStringValue(c *StringValueContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterTrue is called when entering the true production.
	EnterTrue(c *TrueContext)

	// EnterFalse is called when entering the false production.
	EnterFalse(c *FalseContext)

	// EnterNull is called when entering the null production.
	EnterNull(c *NullContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitObj is called when exiting the obj production.
	ExitObj(c *ObjContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitStringValue is called when exiting the stringValue production.
	ExitStringValue(c *StringValueContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitTrue is called when exiting the true production.
	ExitTrue(c *TrueContext)

	// ExitFalse is called when exiting the false production.
	ExitFalse(c *FalseContext)

	// ExitNull is called when exiting the null production.
	ExitNull(c *NullContext)
}
