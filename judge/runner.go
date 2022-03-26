package judge

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/isther/judger/lang"
	"github.com/isther/judger/model"
	"github.com/isther/judger/util"
)

type Runner struct {
	codeType   string
	binaryPath string

	sampleDir string
	outputDir string

	timeLimit   string
	memoryLimit string
}

type RunResult struct {
	SampleId string `json:"sample_id,omitempty"`
	model.Result
}

func newRunner(submit *model.Submit) *Runner {
	return &Runner{
		codeType:   submit.CodeType,
		binaryPath: filepath.Join(TmpPath, submit.SubmitId),
		sampleDir:  filepath.Join(SamplePath, submit.ProblemId),
		outputDir:  filepath.Join(OutputPath, submit.SubmitId),

		timeLimit:   submit.TimeLimit,
		memoryLimit: submit.MemoryLimit,
	}
}

func (runner *Runner) Run() *[]RunResult {
	sampleCount := 0
	files, _ := ioutil.ReadDir(runner.sampleDir)
	for _, fi := range files {
		if fi.IsDir() {
			continue
		} else {
			sampleCount++
		}
	}

	// log.Println("sample count: ", sampleCount)
	if ok, err := util.PathExists(runner.outputDir); err != nil {
		log.Println("Check if path exists failed")
	} else {
		if ok {
			log.Println("Output dir exists: ", runner.outputDir)
		} else {
			if err := os.MkdirAll(runner.outputDir, 0755); err != nil {
				log.Println("Make dir failed")
				return nil
			}
		}
	}

	result := make([]RunResult, 0, sampleCount/2)
	var lock sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < sampleCount/2; i++ {
		wg.Add(1)
		go func(id int) {
			sampleId := strconv.FormatInt(int64(id), 10)
			_result := runner.judgerOneByOne(sampleId)

			lock.Lock()
			result = append(result, *_result)
			lock.Unlock()
			defer wg.Done()
		}(i + 1)
	}
	wg.Wait()
	return &result
}

func (runner *Runner) judgerOneByOne(sampleId string) (_result *RunResult) {
	lang, err := lang.NewLang(runner.codeType, "", runner.binaryPath)
	if err != nil {
		log.Println("New lang failed")
		return &RunResult{
			SampleId: sampleId,
			Result: model.Result{
				Status: GetJudgeStatus(strconv.FormatInt(StatusSE, 10)),
			},
		}
	}

	runnerByOne := exec.Command("sandbox",
		"--bin_path", lang.RunBin(),
		"--args", lang.RunArgs(),
		"--seccomp_rule_name", "general",
		"--input_path", filepath.Join(runner.sampleDir, strings.Join([]string{sampleId, ".in"}, "")),
		"--output_path", filepath.Join(runner.outputDir, strings.Join([]string{sampleId, ".out"}, "")),
		// "--error_path ",
		// "--log_path",
		"--real_time_limit", runner.timeLimit,
		"--cpu_time_limit", runner.timeLimit,
		"--max_memory", runner.memoryLimit,
		// "--max_stack", "",
		// "--max_output_size", "",
		// "--max_process_number", "",
	)

	var o bytes.Buffer
	runnerByOne.Stdin = os.Stdin
	runnerByOne.Stdout = &o
	runnerByOne.Stderr = os.Stderr

	// log.Println(runner.Args)

	if err := runnerByOne.Run(); err != nil {
		log.Fatal("Error: ", err)
		_result.Status = strconv.FormatInt(StatusSE, 10)
		return
	}

	json.Unmarshal(o.Bytes(), &_result)
	_result.SampleId = sampleId
	if _result.Status != strconv.FormatInt(SUCCEED, 10) {
		_result.Status = GetJudgeStatus(_result.Status)
		return
	}

	if ok := runner.Compare(sampleId); ok {
		_result.Status = GetJudgeStatus(strconv.FormatInt(StatusAC, 10))
	} else {
		_result.Status = GetJudgeStatus(strconv.FormatInt(StatusWA, 10))
	}
	// log.Println(_result)

	return _result
}

func (runner *Runner) Compare(sampleId string) bool {
	//TODO: presentation judge
	outPath := filepath.Join(runner.outputDir, strings.Join([]string{sampleId, ".out"}, ""))
	ansPath := filepath.Join(runner.sampleDir, strings.Join([]string{sampleId, ".out"}, ""))

	b, err := ioutil.ReadFile(ansPath)
	if err != nil {
		b = []byte{}
	}

	o, err := ioutil.ReadFile(outPath)
	if err != nil {
		o = []byte{}
	}

	ans := plain(b)
	out := plain(o)
	// log.Printf("ans:= %s", ans)
	// log.Printf("out:= %s", out)

	if out == ans {
		return true
	}
	return false
}

func plain(raw []byte) string {
	buf := bufio.NewScanner(bytes.NewReader(raw))
	var b bytes.Buffer
	newline := []byte{'\n'}
	for buf.Scan() {
		b.Write(bytes.TrimSpace(buf.Bytes()))
		b.Write(newline)
	}
	return b.String()
}
