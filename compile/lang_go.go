package compile

type Golang struct {
	lang
}

func newGolang(sourcePath, binaryPath string) *Golang {
	return &Golang{}
}

func (g *Golang) NeedCompile() bool {
	return true
}
