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
