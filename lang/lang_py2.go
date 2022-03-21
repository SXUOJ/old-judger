package lang

type Python2 struct {
	lang
}

func newPython2(sourcePath, binaryPath string) *Python2 {
	return &Python2{}
}

func (p *Python2) NeedCompile() bool {
	return false
}
