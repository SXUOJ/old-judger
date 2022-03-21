package lang

type Python3 struct {
	lang
}

func newPython3(sourcePath, binaryPath string) *Python3 {
	return &Python3{}
}

func (p *Python3) NeedCompile() bool {
	return false
}
