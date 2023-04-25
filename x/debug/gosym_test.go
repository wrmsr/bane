package debug

import (
	"bufio"
	"debug/elf"
	"debug/gosym"
	"debug/macho"
	"fmt"
	"os"
	"testing"

	rtu "github.com/wrmsr/bane/core/runtime"
	log "github.com/wrmsr/bane/core/slog"
)

func crackElf(file string, t *testing.T) (*elf.File, *gosym.Table) {
	f, err := elf.Open(file)
	if err != nil {
		t.Fatal(err)
	}
	s := f.Section(".gosymtab")
	if s == nil {
		t.Skip("no .gosymtab section")
	}
	symdat, err := s.Data()
	if err != nil {
		f.Close()
		t.Fatalf("reading %s gosymtab: %v", file, err)
	}
	pclndat, err := f.Section(".gopclntab").Data()
	if err != nil {
		f.Close()
		t.Fatalf("reading %s gopclntab: %v", file, err)
	}
	pcln := gosym.NewLineTable(pclndat, f.Section(".text").Addr)
	tab, err := gosym.NewTable(symdat, pcln)
	if err != nil {
		f.Close()
		t.Fatalf("parsing %s gosymtab: %v", file, err)
	}
	return f, tab
}

func crackMacho(file string, t *testing.T) (*macho.File, *gosym.Table) {
	f, err := macho.Open(file)
	if err != nil {
		t.Fatal(err)
	}
	s := f.Section("__gosymtab")
	if s == nil {
		t.Skip("no __gosymtab section")
	}
	symdat, err := s.Data()
	if err != nil {
		f.Close()
		t.Fatalf("reading %s gosymtab: %v", file, err)
	}
	pclndat, err := f.Section("__gopclntab").Data()
	if err != nil {
		f.Close()
		t.Fatalf("reading %s gopclntab: %v", file, err)
	}
	pcln := gosym.NewLineTable(pclndat, f.Section("__text").Addr)
	tab, err := gosym.NewTable(symdat, pcln)
	if err != nil {
		f.Close()
		t.Fatalf("parsing %s gosymtab: %v", file, err)
	}
	return f, tab
}

func TestCrack(t *testing.T) {
	f, tab := crackMacho(os.Args[0], t)
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

func TestLogStd(t *testing.T) {
	log.Printf("hi")

	f, err := macho.Open(os.Args[0])
	if err != nil {
		t.Fatal(err)
	}
	f.Close()

	//	for _, f := range tab.Funcs {
	//		fmt.Println(f.Func.Name)
	//	}

	for _, s := range f.Symtab.Syms {
		if s.Name == "log.std" {
			fmt.Println(s)
			fmt.Println("!!")
		}
	}
}
