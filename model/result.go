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
