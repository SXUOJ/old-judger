package compile

const (
	BasePath   = "/sxu-judger"
	TmpPath    = BasePath + "/run"
	OutputPath = BasePath + "/output"
	SamplePath = BasePath + "/sample"
)

const (
	SUCCEED = iota
	LOOKUP_FAILED
	COMPILE_FAILED
)
