package lang

import "errors"

const (
	_ = iota
	LangC
	LangCpp
	LangGo
	LangJava
	LangPython2
	LangPython3
)

var (
	ERROR_NOT_SUPPORT_LANG = errors.New("This language is not supported")
)

type compilerConfig struct {
	bin  string
	args string

	real_time_limit string
	cpu_time_limit  string
	memory_limit    string
}

type runnerConfig struct {
	bin  string
	args string
}

// Lang: compile parameters
type Lang interface {
	NeedCompile() bool

	CompileBin() string
	CompileArgs() string
	RealTimeLimit() string
	CpuTimeLimit() string
	MemoryLimit() string

	RunBin() string
	RunArgs() string
}

func NewLang(langType, sourcePath, binaryPath string) (Lang, error) {
	switch langType {
	case "C":
		return newC(sourcePath, binaryPath), nil
	case "Cpp":
		return newCpp(sourcePath, binaryPath), nil
	case "Go":
		return newGolang(sourcePath, binaryPath), nil
	case "Python2":
		return newPython2(sourcePath, binaryPath), nil
	case "Python3":
		return newPython3(sourcePath, binaryPath), nil
	default:
		return nil, ERROR_NOT_SUPPORT_LANG
	}
}
