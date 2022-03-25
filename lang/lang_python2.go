package lang

import (
	"strings"
)

type Python2 struct {
	bin  string
	args string
}

func newPython2(sourcePath, binaryPath string) *Python2 {
	return &Python2{
		bin: "/usr/bin/python",
		args: strings.Join([]string{
			sourcePath,
		}, "&"),
	}
}

func (python2 *Python2) NeedCompile() bool {
	return false
}

func (python2 *Python2) Bin() string {
	return python2.bin
}

func (python2 *Python2) Args() string {
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
