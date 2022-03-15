package lang

import (
	"os/exec"
)

type Python3 struct {
	SourcePath string
}

func newPython3(sourcePath string) *Python3 {
	return &Python3{
		SourcePath: sourcePath,
	}
}

func (p *Python3) NeedCompile() bool {
	return false
}

func (p *Python3) Compile() *exec.Cmd {
	return nil
}

func (p *Python3) Run() *exec.Cmd {
	return exec.Command(
		"/usr/bin/python3",
		p.SourcePath,
	)
}
