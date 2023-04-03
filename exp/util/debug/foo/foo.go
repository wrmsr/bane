//
/*

golang.org/x/text/cases.nlTitle: 1000b6ac0 - 1000b6ce0
[  358]    358     Code            0x00000001000b6ac0 0x0000000100366ac0 0x0000000000000220 0x000e0000 golang.org/x/text/cases.nlTitle
0x2B0000 = 2818048
*/
package main

import (
	"bufio"
	"debug/gosym"
	"debug/macho"
	"fmt"
	"log"
	"os"

	rtu "github.com/wrmsr/bane/pkg/util/runtime"
)

func crackMacho(file string) (*macho.File, *gosym.Table) {
	f, err := macho.Open(file)
	if err != nil {
		panic(err)
	}
	s := f.Section("__gosymtab")
	if s == nil {
		panic("no __gosymtab section")
	}
	symdat, err := s.Data()
	if err != nil {
		f.Close()
		panic(err)
	}
	pclndat, err := f.Section("__gopclntab").Data()
	if err != nil {
		f.Close()
		panic(err)
	}
	pcln := gosym.NewLineTable(pclndat, f.Section("__text").Addr)
	tab, err := gosym.NewTable(symdat, pcln)
	if err != nil {
		f.Close()
		panic(err)
	}
	return f, tab
}

func main() {
	log.Println("hi")

	f, tab := crackMacho(os.Args[0])
	f.Close()

	pc := uint64(rtu.GetPc())
	fmt.Printf("pc: %x\n", pc)
	for _, f := range tab.Funcs {
		fmt.Printf("%s: %x - %x\n", f.Name, f.Entry, f.End)
		if pc >= f.Entry && pc < f.End {
			fmt.Printf("!!!!\n")
		}
	}

	fmt.Println(os.Getpid())
	bufio.NewScanner(os.Stdin).Scan()
}
