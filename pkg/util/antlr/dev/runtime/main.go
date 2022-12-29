package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"

	"github.com/wrmsr/bane/pkg/util/check"
)

const AntlrPackage = "github.com/antlr/antlr4/runtime/Go/antlr"
const AntlrVersion = "v1.4.10"

type DownloadedGoMod struct {
	Path     string `json:"Path"`
	Version  string `json:"Version"`
	Info     string `json:"Info"`
	GoMod    string `json:"GoMod"`
	Zip      string `json:"Zip"`
	Dir      string `json:"Dir"`
	Sum      string `json:"Sum"`
	GoModSum string `json:"GoModSum"`
	Origin   struct {
		VCS    string `json:"VCS"`
		URL    string `json:"URL"`
		Subdir string `json:"Subdir"`
		Ref    string `json:"Ref"`
		Hash   string `json:"Hash"`
	} `json:"Origin"`

	X map[string]any `json:"-"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	cmd := exec.CommandContext(
		ctx,
		"go",
		"mod",
		"download",
		"-json",
		fmt.Sprintf("%s@%s", AntlrPackage, AntlrVersion),
	)

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	check.Must(cmd.Run())

	var mod DownloadedGoMod
	check.Must(json.Unmarshal(outb.Bytes(), &mod))
	fmt.Println(mod)
}
