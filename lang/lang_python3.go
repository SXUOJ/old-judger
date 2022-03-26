package lang

import (
	"strings"
)

type Python3 struct {
	bin  string
	args string

	runnerConfig
}

func newPython3(sourcePath, binaryPath string) *Python3 {
	binPath := strings.Join([]string{binaryPath, ".py"}, "")
	return &Python3{
		bin: "/usr/bin/cp",
		args: strings.Join([]string{
			sourcePath,
			binPath,
		}, "&"),
		runnerConfig: runnerConfig{
			bin: "/usr/bin/python3",
			args: strings.Join([]string{
				binPath,
			}, "&"),
		},
	}
}

func (python3 *Python3) NeedCompile() bool {
	return false
}

func (python3 *Python3) CompileBin() string {
	return python3.bin
}

func (python3 *Python3) CompileArgs() string {
	return python3.args
}

func (python *Python3) RealTimeLimit() string {
	return ""
}

func (python *Python3) CpuTimeLimit() string {
	return ""
}

func (python *Python3) MemoryLimit() string {
	return ""
}

func (python *Python3) RunBin() string {
	return python.runnerConfig.bin
}

func (python *Python3) RunArgs() string {
	return python.runnerConfig.args
}
