package model

type Submit struct {
	SubmitId string `json:"submit_id"`

	ProblemId   string `json:"problem_id"`
	ProblemType int64  `json:"problem_type"`

	CodeType   string `json:"code_type"`
	CodeSource string `json:"code_source"`

	Limit
}

type Limit struct {
	TimeLimit   int64 `json:"time_limit"`
	MemoryLimit int64 `json:"memory_limit"`
}
