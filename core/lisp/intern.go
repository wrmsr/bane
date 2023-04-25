package lisp

import stru "github.com/wrmsr/bane/core/strings"

var internedStrings = stru.Intern(
	"(",
	")",
	"*",
	"+",
	"-",
	"/",
	"begin",
	"define",
	"if",
	"let",
)
