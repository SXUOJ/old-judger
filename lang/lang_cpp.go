package lang

import (
	"strconv"
	"strings"
)

type Cpp struct {
	compilerConfig
	runnerConfig
}

func newCpp(sourcePath, binaryPath string) *Cpp {
	return &Cpp{
		compilerConfig: compilerConfig{
			bin: "/usr/bin/g++",
			args: strings.Join([]string{
				"-o",
				binaryPath,
				sourcePath,
				"-fmax-errors=3",
				"-std=c11",
				"-lm",
				"-w",
				"-O2",
				"-DONLINE_JUDGE",
				"",
			}, "&"),
			real_time_limit: "5000",
			cpu_time_limit:  "3000",
			memory_limit:    strconv.FormatInt(128*1024*1024, 10),
		},
		runnerConfig: runnerConfig{
			bin:  binaryPath,
			args: "",
		},
	}
}

func (c *Cpp) NeedCompile() bool {
	return true
}

func (c *Cpp) CompileBin() string {
	return c.compilerConfig.bin
}

func (c *Cpp) CompileArgs() string {
	return c.compilerConfig.args
}

func (c *Cpp) RealTimeLimit() string {
	return c.real_time_limit
}

func (c *Cpp) CpuTimeLimit() string {
	return c.cpu_time_limit
}

func (c *Cpp) MemoryLimit() string {
	return c.memory_limit
}

func (c *Cpp) RunBin() string {
	return c.runnerConfig.bin
}

func (c *Cpp) RunArgs() string {
	return c.runnerConfig.args
}
