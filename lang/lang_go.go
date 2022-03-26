package lang

import (
	"strings"
)

type Golang struct {
	bin  string
	args string

	runnerConfig
}

func newGolang(sourcePath, binaryPath string) *Golang {
	binPath := strings.Join([]string{binaryPath, ".go"}, "")
	return &Golang{
		bin: "/usr/bin/cp",
		args: strings.Join([]string{
			sourcePath,
			binPath,
			"",
		}, "&"),
		runnerConfig: runnerConfig{
			bin: "/usr/bin/go",
			args: strings.Join([]string{
				"run",
				binPath,
				"",
			}, "&"),
		},
	}
}

func (golang *Golang) NeedCompile() bool {
	return false
}

func (golang *Golang) CompileBin() string {
	return golang.bin
}

func (golang *Golang) CompileArgs() string {
	return golang.args
}

func (golang *Golang) RealTimeLimit() string {
	return ""
}

func (golang *Golang) CpuTimeLimit() string {
	return ""
}

func (golang *Golang) MemoryLimit() string {
	return ""
}

func (golang *Golang) RunBin() string {
	return golang.runnerConfig.bin
}

func (golang *Golang) RunArgs() string {
	return golang.runnerConfig.args
}
