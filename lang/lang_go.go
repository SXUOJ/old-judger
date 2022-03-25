package lang

import (
	"strings"
)

type Golang struct {
	bin  string
	args string
}

func newGolang(sourcePath, binaryPath string) *C {
	return &C{
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
