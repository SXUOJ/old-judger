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

	"github.com/isther/judger/compile"
	"github.com/isther/judger/model"
	"github.com/isther/judger/util"
)

type Runner struct {
	codeType       string
	codeSourcePath string

	binPath string

	sampleDir string
	outputDir string

	timeLimit   string
	memoryLimit string
}

type JudgeResult struct {
	SampleId string `json:"sample_id,omitempty"`

	Status string `json:"status,omitempty"`

	CpuTime  string `json:"cpu_time,omitempty"`
	RealTime string `json:"real_time,omitempty"`
	Memory   string `json:"memory,omitempty"`

	Signal string `json:"signal,omitempty"`

	ErrorInf string `json:"msg,omitempty"`
}

func newRunner(submit *model.Submit) *Runner {
	return &Runner{
		codeType:       submit.CodeType,
		codeSourcePath: submit.CodeSourcePath,

		binPath:   filepath.Join(compile.TmpPath, submit.SubmitId),
		sampleDir: filepath.Join(compile.SamplePath, submit.ProblemId),
		outputDir: filepath.Join(compile.OutputPath, submit.SubmitId),

		timeLimit:   submit.TimeLimit,
		memoryLimit: submit.MemoryLimit,
	}
}

func (judger *Runner) Run() *[]JudgeResult {
	sampleCount := 0
	files, _ := ioutil.ReadDir(judger.sampleDir)
	for _, fi := range files {
		if fi.IsDir() {
			continue
		} else {
			sampleCount++
		}
	}

	// log.Println("sample count: ", sampleCount)
	if ok, err := util.PathExists(judger.outputDir); err != nil {
		log.Println("Check if path exists failed")
	} else {
		if ok {
			log.Println("Output dir exists: ", judger.outputDir)
		} else {
			if err := os.MkdirAll(judger.outputDir, 0755); err != nil {
				log.Println("Make dir failed")
				return nil
			}
		}
	}

	result := make([]JudgeResult, sampleCount/2)
	var lock sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < sampleCount/2; i++ {
		wg.Add(1)
		go func(id int) {
			sampleId := strconv.FormatInt(int64(id), 10)
			_result := judger.judgerOneByOne(sampleId)

			lock.Lock()
			result = append(result, *_result)
			lock.Unlock()
			defer wg.Done()
		}(i + 1)
	}
	wg.Wait()
	return &result
}

func (judger *Runner) judgerOneByOne(sampleId string) (_result *JudgeResult) {

	runner := exec.Command("sandbox",
		"--bin_path", judger.binPath,
		"--seccomp_rule_name", "general",
		"--input_path", filepath.Join(judger.sampleDir, strings.Join([]string{sampleId, ".in"}, "")),
		"--output_path", filepath.Join(judger.outputDir, strings.Join([]string{sampleId, ".out"}, "")),
		// "--error_path ",
		// "--log_path",
		"--real_time_limit", judger.timeLimit,
		"--cpu_time_limit", judger.timeLimit,
		"--max_memory", judger.memoryLimit,
		// "--max_stack", "",
		// "--max_output_size", "",
		// "--max_process_number", "",
	)

	var o bytes.Buffer
	runner.Stdin = os.Stdin
	runner.Stdout = &o
	runner.Stderr = os.Stderr

	// log.Println(runner.Args)

	if err := runner.Run(); err != nil {
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

	if ok := judger.Compare(sampleId); ok {
		_result.Status = GetJudgeStatus(strconv.FormatInt(StatusAC, 10))
	} else {
		_result.Status = GetJudgeStatus(strconv.FormatInt(StatusWA, 10))
	}
	// log.Println(_result)

	return _result
}

func (judger *Runner) Compare(sampleId string) bool {
	//TODO: presentation judge
	outPath := filepath.Join(judger.outputDir, strings.Join([]string{sampleId, ".out"}, ""))
	ansPath := filepath.Join(judger.sampleDir, strings.Join([]string{sampleId, ".out"}, ""))

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
	log.Printf("ans:= %s", ans)
	log.Printf("out:= %s", out)

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
