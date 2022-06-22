package lisp

import stru "github.com/wrmsr/bane/pkg/util/strings"

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
