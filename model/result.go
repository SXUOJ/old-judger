package model

// Result result from sandbox
type Result struct {
	Status string `json:"status,omitempty"`

	CpuTime  string `json:"cpu_time,omitempty"`
	RealTime string `json:"real_time,omitempty"`
	Memory   string `json:"memory,omitempty"`

	Signal   string `json:"signal,omitempty"`
	ErrorInf string `json:"msg,omitempty"`
}
