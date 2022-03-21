package judge

import "strconv"

type CompileResult struct {
	Status  string `json:"status,omitempty"`
	CpuTime string `json:"cpu_time,omitempty"`
	Memory  string `json:"real_time,omitempty"`
	Signal  string `json:"signal,omitempty"`

	ErrorInf string `json:"msg,omitempty"`
}

type JudgeResult struct {
	SampleId string `json:"sample_id,omitempty"`

	Result string `json:"status,omitempty"`

	CpuTime  string `json:"cpu_time,omitempty"`
	RealTime string `json:"real_time,omitempty"`
	Memory   string `json:"memory,omitempty"`

	Signal string `json:"signal,omitempty"`

	ErrorInf string `json:"msg,omitempty"`
}

const (
	_ = iota
	StatusAC
	StatusWA
	StatusCE
	StatusRE
	StatusTLE
	StatusMLE
	StatusOLE
	StatusPE
	StatusSE
)

var judgeStatus = map[string]string{
	"1": "Accepted",
	"2": "Wrong Answer",
	"3": "Compile Error",
	"4": "Runtime Error",
	"5": "Time Limit Exceed",
	"6": "Memory Limit Exceed",
	"7": "Output Limit Exceed",
	"8": "Presentation Error",
	"9": "System Error",
}

func GetJudgeStatus(result int) string {
	if v, ok := judgeStatus[strconv.FormatInt(int64(result), 10)]; ok {
		return v
	}
	return "no"
}
