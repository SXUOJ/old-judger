package lang

import (
	"os/exec"
)

type Cpp struct {
	SourcePath string
	BinaryPath string
}

func newCpp(sourcePath, binaryPath string) *Cpp {
	return &Cpp{
		SourcePath: sourcePath,
		BinaryPath: binaryPath,
	}
}

func (c *Cpp) NeedCompile() bool {
	return true
}

func (c *Cpp) Compile() *exec.Cmd {
	return exec.Command(
		"/usr/bin/g++",
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

func (c *Cpp) Run() *exec.Cmd {
	return nil
}
