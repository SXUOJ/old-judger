package model

type JudgeStatus int

type Result struct {
	Status JudgeStatus `json:"status"`

	TimeUsed   int64 `json:"time_used,omitempty"`
	MemoryUsed int64 `json:"memory_used,omitempty"`

	StdOutput string `json:"std_output,omitempty"`

	FileName map[string]int64

	ErrorInf string `json:"msg"`
}

const (
	Waiting JudgeStatus = 0

	StatusAC JudgeStatus = 1

	StatusWA JudgeStatus = 2

	StatusCE JudgeStatus = 3

	StatusRE  JudgeStatus = 4
	StatusTLE JudgeStatus = 5
	StatusMLE JudgeStatus = 6
	StatusOLE JudgeStatus = 7

	StatusPE JudgeStatus = 8

	StatusSE JudgeStatus = 9
)

func NewResult() *Result {
	return &Result{}
}

func (r *Result) Accepted(time, memory int64) {
	r.Status = StatusAC
	r.TimeUsed = time
	r.MemoryUsed = memory
}

func (r *Result) CompileError(msg string) {
	r.Status = StatusCE
	r.ErrorInf = "Compile Error"
}

func (r *Result) RuntimeError() {
	r.Status = StatusRE
	r.ErrorInf = "Runtime Error"
}

func (r *Result) TimeLimitExceed(time, memory int64) {
	r.Status = StatusTLE
	r.TimeUsed = time
	r.MemoryUsed = memory
	r.ErrorInf = "Time Limit Exceeded Error"
}
func (r *Result) MemoryLimitExceed(time, memory int64) {
	r.Status = StatusTLE
	r.TimeUsed = time
	r.MemoryUsed = memory
	r.ErrorInf = "Memory Limit Exceeded Error"
}

func (r *Result) OutputLimitExceed(time, memory int64) {
	r.Status = StatusTLE
	r.TimeUsed = time
	r.MemoryUsed = memory
	r.ErrorInf = "Status Limit Exceeded Error"
}

func (r *Result) PresentationError(time, memory int64) {
	r.Status = StatusPE
	r.TimeUsed = time
	r.MemoryUsed = memory
	r.ErrorInf = "Presentation Error"
}

func (r *Result) WrongAnswer(time, memory int64, input, output, stdOutput string) {
	r.Status = StatusWA
	r.TimeUsed = time
	r.MemoryUsed = memory
	r.StdOutput = stdOutput
}
