package judge

import (
	"github.com/isther/judger/model"
)

type Judger struct {
	Compiler
	Runner
}

func NewJudger(submit *model.Submit) *Judger {
	return &Judger{
		Compiler: *newCompiler(submit),
		Runner:   *newRunner(submit),
	}
}
