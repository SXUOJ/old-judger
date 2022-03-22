package compile

const (
	RunPath    = "./test"
	TmpPath    = RunPath + "/tmp"
	OutputPath = RunPath + "/output"
	SamplePath = RunPath + "/sample"
)

const (
	SUCCEED = iota
	LOOKUP_FAILED
	COMPILE_FAILED
)
