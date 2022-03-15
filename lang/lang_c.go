package lang

import (
	"os/exec"
)

type C struct {
	SourcePath string
	BinaryPath string
	Suffix     string
}

func newC(sourcePath, binaryPath string) *C {
	return &C{
		SourcePath: sourcePath,
		BinaryPath: binaryPath,
		Suffix:     "c",
	}
}

func (c *C) NeedCompile() bool {
	return true
}

func (c *C) Compile() *exec.Cmd {
	return exec.Command(
		"/usr/bin/gcc",
		"-o",
		c.BinaryPath,
		c.SourcePath,
		"-fmax-errors=3",
		"-std=c11",
		"-lm",
		"-w",
		"-O2",
		"-DONLINE_JUDGE",
	)
}

func (c *C) Run() *exec.Cmd {
	return nil
}
