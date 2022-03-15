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
	TimeLimit   string `json:"time_limit"`
	MemoryLimit string `json:"memory_limit"`
}
