package lang

import (
	"strconv"
	"strings"
)

type Cpp lang

func newCpp(sourcePath, binaryPath string) *Cpp {
	return &Cpp{
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
		runCmd:          binaryPath,
	}
}

func (c *Cpp) NeedCompile() bool {
	return true
}

func (c *Cpp) Bin() string {
	return c.bin
}

func (c *Cpp) Args() string {
	return c.args
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