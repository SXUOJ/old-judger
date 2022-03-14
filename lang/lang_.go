package lang

import (
	"os/exec"
)

const (
	_ = iota
	LangC
	LangCpp
	LangJava
	LangPython2
	LangPython3
	LangGo
)

type Lang interface {
	NeedCompile() bool

	Compile() *exec.Cmd
	Run() *exec.Cmd
}
