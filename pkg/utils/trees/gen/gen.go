package main

import (
	"fmt"
	"os"

	"golang.org/x/tools/go/packages"

	"github.com/wrmsr/bane/pkg/utils/log"
)

func main() {
	fmt.Printf("Running %s go on %s\n", os.Args[0], os.Getenv("GOFILE"))

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  cwd = %s\n", cwd)
	fmt.Printf("  os.Args = %#v\n", os.Args)

	for _, ev := range []string{"GOARCH", "GOOS", "GOFILE", "GOLINE", "GOPACKAGE", "DOLLAR"} {
		fmt.Println("  ", ev, "=", os.Getenv(ev))
	}

	log.Println(os.Args)
	cfg := &packages.Config{
		Mode: packages.NeedSyntax | packages.NeedTypes,
	}
	pkgs, err := packages.Load(cfg, "github.com/wrmsr/bane/pkg/utils/inject")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(pkgs)
}
