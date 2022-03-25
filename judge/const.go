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
