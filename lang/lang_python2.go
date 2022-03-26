package lang

import (
	"strings"
)

type Python2 struct {
	bin  string
	args string

	runnerConfig
}

func newPython2(sourcePath, binaryPath string) *Python2 {
	binPath := strings.Join([]string{binaryPath, ".py"}, "")
	return &Python2{
		bin: "/usr/bin/cp",
		args: strings.Join([]string{
			sourcePath,
			binPath,
		}, "&"),

		runnerConfig: runnerConfig{
			bin: "/usr/bin/python",
			args: strings.Join([]string{
				binPath,
				"",
			}, "&"),
		},
	}
}

func (python2 *Python2) NeedCompile() bool {
	return false
}

func (python2 *Python2) CompileBin() string {
	return python2.bin
}

func (python2 *Python2) CompileArgs() string {
	return python2.args
}

func (python *Python2) RealTimeLimit() string {
	return ""
}

func (python *Python2) CpuTimeLimit() string {
	return ""
}

func (python *Python2) MemoryLimit() string {
	return ""
}

func (python *Python2) RunBin() string {
	return python.runnerConfig.bin
}

func (python *Python2) RunArgs() string {
	return python.runnerConfig.args
}
