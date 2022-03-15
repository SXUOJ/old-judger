package lang

import (
	"os/exec"
)

type Go struct {
	SourcePath string
	BinaryPath string
}

func newGo(sourcePath, binaryPath string) *Go {
	return &Go{
		SourcePath: sourcePath,
		BinaryPath: binaryPath,
	}
}

func (g *Go) NeedCompile() bool {
	return false
}

func (g *Go) Compile() *exec.Cmd {
	return nil
}

func (g *Go) Run() *exec.Cmd {
	return nil
}
