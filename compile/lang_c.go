package compile

import (
	"strconv"
	"strings"
)

type C struct {
	bin  string
	args string

	real_time_limit string
	cpu_time_limit  string
	memory_limit    string
}

func newC(sourcePath, binaryPath string) *C {
	return &C{
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
	}
}

func (c *C) NeedCompile() bool {
	return true
}

func (c *C) Bin() string {
	return c.bin
}

func (c *C) Args() string {
	return c.args
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
