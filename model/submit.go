package model

type Submit struct {
	SubmitId string `json:"submit_id"`

	ProblemId   string `json:"problem_id"`
	ProblemType string `json:"problem_type"`

	CodeType       string `json:"code_type"`
	CodeSourcePath string `json:"code_source_path"`

	Limit
}

type Limit struct {
	TimeLimit   int64 `json:"time_limit,string"`
	MemoryLimit int64 `json:"memory_limit,string"`
}
