//go:build !nodev

package debug

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/wrmsr/bane/pkg/util/check"
)

type psItem struct {
	Pid  int
	Ppid int
	Exe  string
}

func getPsItem(ctx context.Context, pid int) psItem {
	cmd := exec.CommandContext(ctx, "ps", "-o", "pid=,ppid=,command=", strconv.Itoa(pid))
	var outb bytes.Buffer
	cmd.Stdout = &outb
	check.NoErr(cmd.Run())

	outs := outb.String()
	s := outs
	spid, s, ok := strings.Cut(strings.TrimSpace(s), " ")
	if !ok {
		panic(outs)
	}
	sppid, scmd, ok := strings.Cut(strings.TrimSpace(s), " ")
	if !ok {
		panic(outs)
	}

	pid2 := check.Must(strconv.Atoi(strings.TrimSpace(spid)))
	if pid2 != pid {
		panic(fmt.Errorf("pid %d != %d", pid2, pid))
	}

	ppid := check.Must(strconv.Atoi(strings.TrimSpace(sppid)))

	var exe string
	p := 0
	for {
		var c string
		n := strings.Index(scmd[p:], " ")
		if n == -1 {
			c = scmd
		} else {
			c = scmd[:p+n]
		}

		if _, err := os.Stat(c); err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				panic(err)
			}
		} else {
			exe = c
			break
		}

		if n < 0 {
			break
		}
		p += n + 1
	}

	return psItem{
		Pid:  pid,
		Ppid: ppid,
		Exe:  exe,
	}
}

func getPsLineage(ctx context.Context, pid int) []psItem {
	var ret []psItem
	for {
		item := getPsItem(ctx, pid)
		if item.Ppid < 1 {
			break
		}
		ret = append(ret, item)
		pid = item.Ppid
	}
	return ret
}

func IsDebuggerAttached(ctx context.Context) bool {
	lg := getPsLineage(ctx, os.Getpid())
	for _, i := range lg {
		if filepath.Base(i.Exe) == "dlv" {
			return true
		}
	}
	return false
}
