package lang

import (
	"strings"
)

type Python3 struct {
	bin  string
	args string
}

func newPython3(sourcePath, binaryPath string) *C {
	return &C{
		bin: "/usr/bin/python3",
		args: strings.Join([]string{
			sourcePath,
		}, "&"),
	}
}

func (python3 *Python3) NeedCompile() bool {
	return false
}

func (python3 *Python3) Bin() string {
	return python3.bin
}

func (python3 *Python3) Args() string {
	return python3.args
}
