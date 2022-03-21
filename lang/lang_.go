package lang

import (
	"errors"
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

type lang struct {
	bin  string
	args string

	real_time_limit string
	cpu_time_limit  string
	memory_limit    string
}

type Lang interface {
	NeedCompile() bool

	Bin() string
	Args() string
	RealTimeLimit() string
	CpuTimeLimit() string
	MemoryLimit() string
}

func NewLang(langType, sourcePath, binaryPath string) (Lang, error) {
	switch langType {
	case "C":
		return newC(sourcePath, binaryPath), nil
	case "Cpp":
		return newCpp(sourcePath, binaryPath), nil
	default:
		return nil, ERROR_NOT_SUPPORT_LANG
	}
}
