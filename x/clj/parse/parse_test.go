/*
Copyright (c) 2014 Caleb Spare

# MIT License

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
package parse

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/wrmsr/bane/core/check"
	ju "github.com/wrmsr/bane/core/json"
	msh "github.com/wrmsr/bane/core/marshal"
)

var testCases = []struct {
	s    string
	want string
}{
	// all types
	{"true", "true"},
	{`\s`, "char('s')"},
	{"; comment!", `comment("; comment!")`},
	{"@foo", "deref"},
	{"#(+ % 3)", "lambda(length=3)"},
	{"#_(a b c)", "discard"},
	{":foobar", "keyword(:foobar)"},
	{"(foo bar baz)", "list(length=3)"},
	{"{:a b :c d}", "map(length=2)"},
	{"#:foo{:a 1}", "map(ns=:foo, length=1)"},
	{"#::{:b 1234}", "map(ns=::, length=1)"},
	{`^String`, "metadata"},
	{"nil", "nil"},
	{"123.456", "num(123.456)"},
	{"foo", "sym(foo)"},
	{"'(foobar)", "quote"},
	{`#"^asdf"`, `regex("^asdf")`},
	{"#{1 2 3}", "set(length=3)"},
	{"#?(:clj 1)", "reader-cond(length=2)"},
	{"#?@(:clj :a :default :b)", "reader-cond-splice(length=4)"},
	{`"foo"`, `string("foo")`},
	{"`(1 2 3)", "syntax quote"},
	{"#foo", "tag(foo)"},
	{"~foo", "unquote"},
	{"~@foo", "unquote splice"},
	{"#'asdf", "varquote(asdf)"},
	{"[a b c]", "vector(length=3)"},

	// issue 13
	{"#_foobar", "discard"},

	// issue 32
	{"#=foo", "eval"},
	{"#^foo", "metadata"},
	{"#! hello!", `comment("#! hello!")`},

	// issue 35
	{"a%b%", "sym(a%b%)"},
	{":100%>50%", "keyword(:100%>50%)"},
}

func TestAll(t *testing.T) {
	for _, tc := range testCases {
		tree, err := Reader(strings.NewReader(tc.s), "temp", IncludeNonSemantic)
		if err != nil {
			t.Fatalf("error parsing %q: %s", tc.s, err)
		}
		if len(tree.Roots) != 1 {
			t.Errorf("got %d roots for %q; want 1", len(tree.Roots), tc.s)
			continue
		}
		root := tree.Roots[0]
		got := root.String()
		if got != tc.want {
			t.Errorf("for %q: got %s; want %s", tc.s, got, tc.want)
			continue
		}
		fmt.Println(check.Must1(ju.MarshalPretty(check.Must1(msh.Marshal(&root)))))
	}
}

func TestIgnoreReaderDiscard(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want string
	}{
		{"#_ignore", ""},
		{"a #_ignore b", "sym(a) sym(b)"},
		{"[a b #_ignore]", "vector(length=2) sym(a) sym(b)"},
	} {
		tree, err := Reader(strings.NewReader(tc.s), "temp", IgnoreReaderDiscard)
		if err != nil {
			t.Fatalf("error parsing %q: %s", tc.s, err)
		}
		got := strings.Join(tree.flatStrings(), " ")
		if got != tc.want {
			t.Errorf("for %q: got %s; want %s", tc.s, got, tc.want)
			continue
		}
	}
}

func TestIgnoreCommentForm(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want string
	}{
		{"(comment 1 2 3)", ""},
		{"'(comment 1)", "quote list(length=2) sym(comment) num(1)"},
		{"a (comment 1) b", "sym(a) sym(b)"},
		{"[a b (comment 1)]", "vector(length=2) sym(a) sym(b)"},
	} {
		tree, err := Reader(strings.NewReader(tc.s), "temp", IgnoreCommentForm)
		if err != nil {
			t.Fatalf("error parsing %q: %s", tc.s, err)
		}
		got := strings.Join(tree.flatStrings(), " ")
		if got != tc.want {
			t.Errorf("for %q: got %s; want %s", tc.s, got, tc.want)
			continue
		}
	}
}

// Issue 32.
func TestUnreadable(t *testing.T) {
	_, err := Reader(strings.NewReader("#<X Y Z>"), "temp", IncludeNonSemantic)
	if err == nil {
		t.Fatal("got nil error on unreadable dispatch macro")
	}
	if !strings.Contains(err.Error(), "unreadable") {
		t.Fatalf("for unreadable dispatch macro, got wrong error %s", err)
	}
}

// Issue 33.
func TestCommentCarriageReturn(t *testing.T) {
	const input = "3;a\r4"
	tree, err := Reader(strings.NewReader(input), "temp", IncludeNonSemantic)
	if err != nil {
		t.Fatalf("error parsing %q: %s", input, err)
	}
	got := tree.flatStrings()
	want := []string{"num(3)", `comment(";a")`, "num(4)"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("for %q, got %v; want %v", input, got, want)
	}
}

// Issue 37.
func TestInternalNewlines(t *testing.T) {
	const input = "[3\n4]"
	tree, err := Reader(strings.NewReader(input), "temp", IncludeNonSemantic)
	if err != nil {
		t.Fatalf("error parsing %q: %s", input, err)
	}
	got := tree.flatStrings()
	want := []string{"vector(length=2)", "num(3)", "newline", "num(4)"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("for %q, got %v; want %v", input, got, want)
	}
}

// Issue 48.
func TestUnterminatedQuotes(t *testing.T) {
	for _, input := range []string{"@", "'", "`", "~", "~@"} {
		_, err := Reader(strings.NewReader(input), "temp", IncludeNonSemantic)
		if !strings.HasSuffix(err.Error(), "unexpected EOF") {
			t.Errorf("for %q, got err=%v; want unexpected EOF", input, err)
		}
	}

	const input = "';hello\na"
	tree, err := Reader(strings.NewReader(input), "temp", IncludeNonSemantic)
	if err != nil {
		t.Fatalf("error parsing %q: %s", input, err)
	}
	got := tree.flatStrings()
	want := []string{"quote", "sym(a)"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("for %q: got %v; want %v", input, got, want)
	}
}

func TestDeref(t *testing.T) {
	input := `(.foo bar)`
	tree, err := Reader(strings.NewReader(input), "temp", IncludeNonSemantic)
	if err != nil {
		t.Fatalf("error parsing %q: %s", input, err)
	}
	got := tree.flatStrings()
	fmt.Println(got)
}

// flatStrings gives a flattened string representation of t by calling String on each node in the tree in a depth-first
// traversal.
func (t *Tree) flatStrings() []string {
	var nodes []string
	var visit func(n Node)
	visit = func(n Node) {
		nodes = append(nodes, n.String())
		for _, child := range n.Children() {
			visit(child)
		}
	}
	for _, root := range t.Roots {
		visit(root)
	}
	return nodes
}
