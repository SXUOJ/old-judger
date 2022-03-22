package judge

import (
	"github.com/isther/judger/compile"
	"github.com/isther/judger/model"
)

type Judger struct {
	compile.Compiler
	Runner
}

func NewJudger(submit *model.Submit) *Judger {
	return &Judger{
		Compiler: *compile.NewCompiler(submit),
		Runner:   *newRunner(submit),
	}
}
