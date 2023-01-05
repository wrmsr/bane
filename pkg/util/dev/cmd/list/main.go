package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/wrmsr/bane/pkg/util/check"
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

type ModuleError struct {
	Err string // error text
}

func main() {
	flags := flag.NewFlagSet("list", flag.ExitOnError)

	flagDeps := flags.Bool("deps", false, "")
	flagFind := flags.Bool("find", false, "")

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
	check.Must(cmd.Run())

	jd := json.NewDecoder(bytes.NewReader(outb.Bytes()))
	for jd.More() {
		var m map[string]any
		check.Must(jd.Decode(&m))
		fmt.Println(m)
	}
}
