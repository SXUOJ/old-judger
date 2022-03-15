package judge

import (
	"path/filepath"
	"sync"

	"github.com/Sxu-Online-Judge/judger/lang"
	"github.com/Sxu-Online-Judge/judger/model"
)

const (
	samplePath = "sample"
	tmpPath    = "/tmp"
)

type Judger struct {
	problemId      string
	problemType    string
	codeType       string
	codeSourcePath string

	binPath string

	sampleDir string
	outputDir string

	timeLimit   int64
	memoryLimit int64
}

func NewJudger(submit *model.Submit) *Judger {
	return &Judger{
		problemId:      submit.ProblemId,
		problemType:    submit.ProblemType,
		codeType:       submit.CodeType,
		codeSourcePath: submit.CodeSourcePath,

		binPath:   filepath.Join(tmpPath, submit.ProblemId, submit.SubmitId),
		sampleDir: filepath.Join(samplePath, submit.ProblemId, submit.SubmitId),
		outputDir: filepath.Join(tmpPath, submit.ProblemId, submit.SubmitId),

		timeLimit:   submit.TimeLimit,
		memoryLimit: submit.MemoryLimit,
	}
}

func (judger *Judger) Judge() *[]model.Result {
	lang, err := judger.newLang()
	if err != nil {
		//TODO:
	}

	if ok := lang.NeedCompile(); ok {
		cmd := lang.Compile()
		if err := cmd.Run(); err != nil {
			//TODO:
		}
	}

	filepathNames, err := filepath.Glob(judger.sampleDir)
	if err != nil {
		//TODO:
	}

	sampleCount := len(filepathNames) / 2
	result := make([]model.Result, sampleCount)
	var lock sync.Mutex
	for i := 0; i < sampleCount; i++ {
		go func(sampleId int) {
			_result := judger.judgerOneByOne(sampleId)

			lock.Lock()
			result = append(result, *_result)
			lock.Unlock()
		}(i + 1)
	}

	return &result
}

func (judger *Judger) newLang() (lang.Lang, error) {
	return lang.NewLang(judger.codeType,
		judger.codeSourcePath,
		judger.binPath,
	)
}

func (judger *Judger) judgerOneByOne(sampleId int) *model.Result {

	return nil
}
