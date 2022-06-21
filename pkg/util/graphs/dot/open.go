package dot

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func Open(ctx context.Context, gv string) error {
	tf, err := ioutil.TempFile("", "dot.pdf")
	if err != nil {
		return err
	}

	cmd := exec.CommandContext(ctx, "dot", "-Tpdf")
	cmd.Stdin = strings.NewReader(gv)
	cmd.Stdout = tf

	var errb bytes.Buffer
	cmd.Stderr = &errb
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("dot error: %s", errb.String())
	}

	fn := tf.Name() + ".pdf"
	if err := os.Rename(tf.Name(), fn); err != nil {
		return err
	}

	cmd = exec.CommandContext(ctx, "open", fn)
	errb = bytes.Buffer{}
	cmd.Stderr = &errb
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("dot error: %s", errb.String())
	}

	return nil
}
