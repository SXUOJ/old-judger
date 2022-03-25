package lang

import (
	"strings"
)

type Golang struct {
	bin  string
	args string
}

func newGolang(sourcePath, binaryPath string) *Golang {
	return &Golang{
		bin: "/usr/bin/go",
		args: strings.Join([]string{
			"run",
			sourcePath,
		}, "&"),
	}
}

func (golang *Golang) NeedCompile() bool {
	return false
}

func (golang *Golang) Bin() string {
	return golang.bin
}

func (golang *Golang) Args() string {
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
