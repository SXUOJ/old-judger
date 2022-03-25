package judge

const (
	BasePath   = "/sxu-judger"
	CodePath   = BasePath + "/code"
	TmpPath    = BasePath + "/run"
	OutputPath = BasePath + "/output"
	SamplePath = BasePath + "/sample"
)

const (
	SUCCEED = iota
	LOOKUP_FAILED
	COMPILE_FAILED
)

const (
	StatusAC  = 1
	StatusWA  = 2
	StatusCE  = 3
	StatusRE  = 4
	StatusTLE = 5
	StatusMLE = 6
	StatusOLE = 7
	StatusPE  = 8
	StatusSE  = 9
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

func GetJudgeStatus(result string) string {
	if v, ok := judgeStatus[result]; ok {
		return v
	}
	return "no"
}
