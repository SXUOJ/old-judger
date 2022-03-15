package judge

import "github.com/Sxu-Online-Judge/judger/model"

type Judger struct {
	ProblemId      string
	ProblemType    string
	CodeType       string
	CodeSourcePath string

	TimeLimit   int64
	MemoryLimit int64
}

func NewJudger(submit *model.Submit) *Judger {
	return &Judger{
		ProblemId:      submit.ProblemId,
		ProblemType:    submit.ProblemType,
		CodeType:       submit.CodeType,
		CodeSourcePath: submit.CodeSourcePath,
		TimeLimit:      submit.TimeLimit,
		MemoryLimit:    submit.MemoryLimit,
	}
}

func (judger *Judger) Judge() []*model.Result {
	return nil
}

func (judger *Judger) judgerOneByOne(id int, input_sample_path, output_sample_path, output_path string) *model.Result {
	return nil
}
