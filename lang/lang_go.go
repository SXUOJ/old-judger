package lang

import (
	"os/exec"
)

type Go struct {
	SourcePath string
	BinaryPath string
	Suffix     string
}

func newGo(sourcePath, binaryPath string) *Go {
	return &Go{
		SourcePath: sourcePath,
		BinaryPath: binaryPath,
		Suffix:     "go",
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
