/*
Copyright (c) 2013 John Barton

MIT License

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
package dotenv

import (
	"bytes"
	"strings"
	"testing"
)

func parseAndCompare(t *testing.T, rawEnvLine string, expectedKey string, expectedValue string) {
	key, value, _ := NewParser().parseLine(rawEnvLine)
	if key != expectedKey || value != expectedValue {
		t.Errorf("Expected '%v' to parse as '%v' => '%v', got '%v' => '%v' instead", rawEnvLine, expectedKey, expectedValue, key, value)
	}
}

func TestParse(t *testing.T) {
	envMap, err := Parse(bytes.NewReader([]byte("ONE=1\nTWO='2'\nTHREE = \"3\"")), nil)
	expectedValues := map[string]string{
		"ONE":   "1",
		"TWO":   "2",
		"THREE": "3",
	}
	if err != nil {
		t.Fatalf("error parsing env: %v", err)
	}
	for key, value := range expectedValues {
		if envMap[key] != value {
			t.Errorf("expected %s to be %s, got %s", key, value, envMap[key])
		}
	}
}

func TestExpanding(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]string
	}{
		{
			"expands variables found in values",
			"FOO=test\nBAR=$FOO",
			map[string]string{"FOO": "test", "BAR": "test"},
		},
		{
			"parses variables wrapped in brackets",
			"FOO=test\nBAR=${FOO}bar",
			map[string]string{"FOO": "test", "BAR": "testbar"},
		},
		{
			"expands undefined variables to an empty string",
			"BAR=$FOO",
			map[string]string{"BAR": ""},
		},
		{
			"expands variables in double quoted strings",
			"FOO=test\nBAR=\"quote $FOO\"",
			map[string]string{"FOO": "test", "BAR": "quote test"},
		},
		{
			"does not expand variables in single quoted strings",
			"BAR='quote $FOO'",
			map[string]string{"BAR": "quote $FOO"},
		},
		{
			"does not expand escaped variables",
			`FOO="foo\$BAR"`,
			map[string]string{"FOO": "foo$BAR"},
		},
		{
			"does not expand escaped variables",
			`FOO="foo\${BAR}"`,
			map[string]string{"FOO": "foo${BAR}"},
		},
		{
			"does not expand escaped variables",
			"FOO=test\nBAR=\"foo\\${FOO} ${FOO}\"",
			map[string]string{"FOO": "test", "BAR": "foo${FOO} test"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			env, err := Parse(strings.NewReader(tt.input), nil)
			if err != nil {
				t.Errorf("Error: %s", err.Error())
			}
			for k, v := range tt.expected {
				if strings.Compare(env[k], v) != 0 {
					t.Errorf("Expected: %s, Actual: %s", v, env[k])
				}
			}
		})
	}

}

func TestParsing(t *testing.T) {
	// unquoted values
	parseAndCompare(t, "FOO=bar", "FOO", "bar")

	// parses values with spaces around equal sign
	parseAndCompare(t, "FOO =bar", "FOO", "bar")
	parseAndCompare(t, "FOO= bar", "FOO", "bar")

	// parses double quoted values
	parseAndCompare(t, `FOO="bar"`, "FOO", "bar")

	// parses single quoted values
	parseAndCompare(t, "FOO='bar'", "FOO", "bar")

	// parses escaped double quotes
	parseAndCompare(t, `FOO="escaped\"bar"`, "FOO", `escaped"bar`)

	// parses single quotes inside double quotes
	parseAndCompare(t, `FOO="'d'"`, "FOO", `'d'`)

	// parses yaml style options
	parseAndCompare(t, "OPTION_A: 1", "OPTION_A", "1")

	//parses yaml values with equal signs
	parseAndCompare(t, "OPTION_A: Foo=bar", "OPTION_A", "Foo=bar")

	// parses non-yaml options with colons
	parseAndCompare(t, "OPTION_A=1:B", "OPTION_A", "1:B")

	// parses export keyword
	parseAndCompare(t, "export OPTION_A=2", "OPTION_A", "2")
	parseAndCompare(t, `export OPTION_B='\n'`, "OPTION_B", "\\n")
	parseAndCompare(t, "export exportFoo=2", "exportFoo", "2")
	parseAndCompare(t, "exportFOO=2", "exportFOO", "2")
	parseAndCompare(t, "export_FOO =2", "export_FOO", "2")
	parseAndCompare(t, "export.FOO= 2", "export.FOO", "2")
	parseAndCompare(t, "export\tOPTION_A=2", "OPTION_A", "2")
	parseAndCompare(t, "  export OPTION_A=2", "OPTION_A", "2")
	parseAndCompare(t, "\texport OPTION_A=2", "OPTION_A", "2")

	// it 'expands newlines in quoted strings' do
	// expect(env('FOO="bar\nbaz"')).to eql('FOO' => "bar\nbaz")
	parseAndCompare(t, `FOO="bar\nbaz"`, "FOO", "bar\nbaz")

	// it 'parses variables with "." in the name' do
	// expect(env('FOO.BAR=foobar')).to eql('FOO.BAR' => 'foobar')
	parseAndCompare(t, "FOO.BAR=foobar", "FOO.BAR", "foobar")

	// it 'parses variables with several "=" in the value' do
	// expect(env('FOO=foobar=')).to eql('FOO' => 'foobar=')
	parseAndCompare(t, "FOO=foobar=", "FOO", "foobar=")

	// it 'strips unquoted values' do
	// expect(env('foo=bar ')).to eql('foo' => 'bar') # not 'bar '
	parseAndCompare(t, "FOO=bar ", "FOO", "bar")

	// it 'ignores inline comments' do
	// expect(env("foo=bar # this is foo")).to eql('foo' => 'bar')
	parseAndCompare(t, "FOO=bar # this is foo", "FOO", "bar")

	// it 'allows # in quoted value' do
	// expect(env('foo="bar#baz" # comment')).to eql('foo' => 'bar#baz')
	parseAndCompare(t, `FOO="bar#baz" # comment`, "FOO", "bar#baz")
	parseAndCompare(t, "FOO='bar#baz' # comment", "FOO", "bar#baz")
	parseAndCompare(t, `FOO="bar#baz#bang" # comment`, "FOO", "bar#baz#bang")

	// it 'parses # in quoted values' do
	// expect(env('foo="ba#r"')).to eql('foo' => 'ba#r')
	// expect(env("foo='ba#r'")).to eql('foo' => 'ba#r')
	parseAndCompare(t, `FOO="ba#r"`, "FOO", "ba#r")
	parseAndCompare(t, "FOO='ba#r'", "FOO", "ba#r")

	//newlines and backslashes should be escaped
	parseAndCompare(t, `FOO="bar\n\ b\az"`, "FOO", "bar\n baz")
	parseAndCompare(t, `FOO="bar\\\n\ b\az"`, "FOO", "bar\\\n baz")
	parseAndCompare(t, `FOO="bar\\r\ b\az"`, "FOO", "bar\\r baz")

	parseAndCompare(t, `="value"`, "", "value")
	parseAndCompare(t, `KEY="`, "KEY", "\"")
	parseAndCompare(t, `KEY="value`, "KEY", "\"value")

	// leading whitespace should be ignored
	parseAndCompare(t, " KEY =value", "KEY", "value")
	parseAndCompare(t, "   KEY=value", "KEY", "value")
	parseAndCompare(t, "\tKEY=value", "KEY", "value")

	// it 'throws an error if line format is incorrect' do
	// expect{env('lol$wut')}.to raise_error(Dotenv::FormatError)
	badlyFormattedLine := "lol$wut"
	_, _, err := NewParser().parseLine(badlyFormattedLine)
	if err == nil {
		t.Errorf("Expected \"%v\" to return error, but it didn't", badlyFormattedLine)
	}
}

func TestLinesToIgnore(t *testing.T) {
	// it 'ignores empty lines' do
	// expect(env("\n \t  \nfoo=bar\n \nfizz=buzz")).to eql('foo' => 'bar', 'fizz' => 'buzz')
	if !isIgnoredLine("\n") {
		t.Error("Line with nothing but line break wasn't ignored")
	}

	if !isIgnoredLine("\r\n") {
		t.Error("Line with nothing but windows-style line break wasn't ignored")
	}

	if !isIgnoredLine("\t\t ") {
		t.Error("Line full of whitespace wasn't ignored")
	}

	// it 'ignores comment lines' do
	// expect(env("\n\n\n # HERE GOES FOO \nfoo=bar")).to eql('foo' => 'bar')
	if !isIgnoredLine("# comment") {
		t.Error("Comment wasn't ignored")
	}

	if !isIgnoredLine("\t#comment") {
		t.Error("Indented comment wasn't ignored")
	}

	// make sure we're not getting false positives
	if isIgnoredLine(`export OPTION_B='\n'`) {
		t.Error("ignoring a perfectly valid line to parse")
	}
}
