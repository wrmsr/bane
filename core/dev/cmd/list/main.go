//go:build !nodev

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/wrmsr/bane/core/check"
	ju "github.com/wrmsr/bane/core/json"
	osu "github.com/wrmsr/bane/core/os"
	"github.com/wrmsr/bane/core/slices"
)

type Origin struct {
	Vcs    string `json:",omitempty"`
	Url    string `json:",omitempty"`
	Subdir string `json:",omitempty"`

	TagPrefix string `json:",omitempty"`
	TagSum    string `json:",omitempty"`

	Ref  string `json:",omitempty"`
	Hash string `json:",omitempty"`

	RepoSum string `json:",omitempty"`
}

type ModuleError struct {
	Err string
}

type Module struct {
	Path       string       `json:",omitempty"`
	Version    string       `json:",omitempty"`
	Query      string       `json:",omitempty"`
	Versions   []string     `json:",omitempty"`
	Replace    *Module      `json:",omitempty"`
	Time       *time.Time   `json:",omitempty"`
	Update     *Module      `json:",omitempty"`
	Main       bool         `json:",omitempty"`
	Indirect   bool         `json:",omitempty"`
	Dir        string       `json:",omitempty"`
	GoMod      string       `json:",omitempty"`
	GoVersion  string       `json:",omitempty"`
	Retracted  []string     `json:",omitempty"`
	Deprecated string       `json:",omitempty"`
	Error      *ModuleError `json:",omitempty"`

	Origin *Origin `json:",omitempty"`
	Reuse  bool    `json:",omitempty"`
}

type PackageError struct {
	ImportStack      []string
	Pos              string
	Err              error
	IsImportCycle    bool
	Hard             bool
	alwaysPrintStack bool
}

type Package struct {
	Dir           string   `json:",omitempty"`
	ImportPath    string   `json:",omitempty"`
	ImportComment string   `json:",omitempty"`
	Name          string   `json:",omitempty"`
	Doc           string   `json:",omitempty"`
	Target        string   `json:",omitempty"`
	Shlib         string   `json:",omitempty"`
	Root          string   `json:",omitempty"`
	ConflictDir   string   `json:",omitempty"`
	ForTest       string   `json:",omitempty"`
	Export        string   `json:",omitempty"`
	BuildID       string   `json:",omitempty"`
	Module        *Module  `json:",omitempty"`
	Match         []string `json:",omitempty"`
	Goroot        bool     `json:",omitempty"`
	Standard      bool     `json:",omitempty"`
	DepOnly       bool     `json:",omitempty"`
	BinaryOnly    bool     `json:",omitempty"`
	Incomplete    bool     `json:",omitempty"`

	Stale       bool   `json:",omitempty"`
	StaleReason string `json:",omitempty"`

	GoFiles           []string `json:",omitempty"`
	CgoFiles          []string `json:",omitempty"`
	CompiledGoFiles   []string `json:",omitempty"`
	IgnoredGoFiles    []string `json:",omitempty"`
	InvalidGoFiles    []string `json:",omitempty"`
	IgnoredOtherFiles []string `json:",omitempty"`
	CFiles            []string `json:",omitempty"`
	CXXFiles          []string `json:",omitempty"`
	MFiles            []string `json:",omitempty"`
	HFiles            []string `json:",omitempty"`
	FFiles            []string `json:",omitempty"`
	SFiles            []string `json:",omitempty"`
	SwigFiles         []string `json:",omitempty"`
	SwigCXXFiles      []string `json:",omitempty"`
	SysoFiles         []string `json:",omitempty"`

	EmbedPatterns []string `json:",omitempty"`
	EmbedFiles    []string `json:",omitempty"`

	CgoCFLAGS    []string `json:",omitempty"`
	CgoCPPFLAGS  []string `json:",omitempty"`
	CgoCXXFLAGS  []string `json:",omitempty"`
	CgoFFLAGS    []string `json:",omitempty"`
	CgoLDFLAGS   []string `json:",omitempty"`
	CgoPkgConfig []string `json:",omitempty"`

	Imports   []string          `json:",omitempty"`
	ImportMap map[string]string `json:",omitempty"`
	Deps      []string          `json:",omitempty"`

	Error      *PackageError   `json:",omitempty"`
	DepsErrors []*PackageError `json:",omitempty"`

	TestGoFiles        []string `json:",omitempty"`
	TestImports        []string `json:",omitempty"`
	TestEmbedPatterns  []string `json:",omitempty"`
	TestEmbedFiles     []string `json:",omitempty"`
	XTestGoFiles       []string `json:",omitempty"`
	XTestImports       []string `json:",omitempty"`
	XTestEmbedPatterns []string `json:",omitempty"`
	XTestEmbedFiles    []string `json:",omitempty"`
}

func main() {
	flags := flag.NewFlagSet("list", flag.ExitOnError)

	flagDeps := flags.Bool("deps", false, "")
	flagFind := flags.Bool("find", false, "")

	flagJson := flags.Bool("json", false, "")
	flagIgnoreFiles := flags.Bool("ignorefiles", false, "")

	check.Must(flags.Parse(os.Args[1:]))

	cmdArgs := []string{
		"list",
		"-json",
	}
	if *flagDeps {
		cmdArgs = append(cmdArgs, "-deps")
	}
	if *flagFind {
		cmdArgs = append(cmdArgs, "-find")
	}
	cmdArgs = append(cmdArgs, flags.Args()...)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	cmd := exec.CommandContext(
		ctx,
		"go",
		cmdArgs...,
	)

	var outb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = os.Stderr
	check.Must(cmd.Run())

	var pkgs []*Package
	jd := json.NewDecoder(bytes.NewReader(outb.Bytes()))
	for jd.More() {
		var pkg *Package
		check.Must(jd.Decode(&pkg))
		pkgs = append(pkgs, pkg)
	}

	if *flagIgnoreFiles {
		pkgs = slices.Filter(func(p *Package) bool {
			return p.Dir == "" || !check.Must1(osu.Exists(filepath.Join(p.Dir, ".goignore")))
		}, pkgs)
	}

	if *flagJson {
		fmt.Println(check.Must1(ju.MarshalPretty(pkgs)))
	} else {
		for _, p := range pkgs {
			fmt.Println(p.ImportPath)
		}
	}
}
