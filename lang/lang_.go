package lang

import (
	"errors"
	"os/exec"
)

var (
	ERROR_NOT_SUPPORT_LANG = errors.New("This language is not supported")
)

const (
	_ = iota
	LangC
	LangCpp
	LangGo
	LangJava
	LangPython2
	LangPython3
)

type Lang interface {
	NeedCompile() bool

	Compile() *exec.Cmd
	Run() *exec.Cmd
}

func NewLang(langType, sourcePath, binaryPath string) (Lang, error) {
	switch langType {
	case "C":
		return newC(sourcePath, binaryPath), nil
	case "Cpp":
		return newCpp(sourcePath, binaryPath), nil
	case "Go":
		return newGo(sourcePath, binaryPath), nil
	case "Java":
		return nil, nil
	case "Python2":
		return newPython2(sourcePath), nil
	case "Python3":
		return newPython3(sourcePath), nil
	default:
		return nil, ERROR_NOT_SUPPORT_LANG
	}
}
