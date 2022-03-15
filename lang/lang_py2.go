package lang

import (
	"os/exec"
)

type Python2 struct {
	SourcePath string
}

func newPython2(sourcePath string) *Python2 {
	return &Python2{
		SourcePath: sourcePath,
	}
}

func (p *Python2) NeedCompile() bool {
	return false
}

func (p *Python2) Compile() *exec.Cmd {
	return nil
}

func (p *Python2) Run() *exec.Cmd {
	return exec.Command(
		"/usr/bin/python",
		p.SourcePath,
	)
}
