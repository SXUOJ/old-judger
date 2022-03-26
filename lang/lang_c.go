package lang

import (
	"strconv"
	"strings"
)

type C struct {
	compilerConfig
	runnerConfig
}

func newC(sourcePath, binaryPath string) *C {
	return &C{
		compilerConfig: compilerConfig{
			bin: "/usr/bin/gcc",
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

func (c *C) NeedCompile() bool {
	return true
}

func (c *C) CompileBin() string {
	return c.compilerConfig.bin
}

func (c *C) CompileArgs() string {
	return c.compilerConfig.args
}

func (c *C) RealTimeLimit() string {
	return c.real_time_limit
}

func (c *C) CpuTimeLimit() string {
	return c.cpu_time_limit
}

func (c *C) MemoryLimit() string {
	return c.memory_limit
}

func (c *C) RunBin() string {
	return c.runnerConfig.bin
}

func (c *C) RunArgs() string {
	return c.runnerConfig.args
}
