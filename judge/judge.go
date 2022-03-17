package judge

import (
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

	"github.com/Sxu-Online-Judge/judger/lang"
	"github.com/Sxu-Online-Judge/judger/model"
	"github.com/Sxu-Online-Judge/judger/util"
)

const (
	dataPath   = "./test"
	samplePath = dataPath + "/sample"
	tmpPath    = dataPath + "/tmp"
)

type Judger struct {
	codeType       string
	codeSourcePath string

	binPath string

	sampleDir string
	outputDir string

	timeLimit   string
	memoryLimit string
}

func NewJudger(submit *model.Submit) *Judger {
	return &Judger{
		codeType:       submit.CodeType,
		codeSourcePath: submit.CodeSourcePath,

		binPath:   filepath.Join(tmpPath, submit.SubmitId),
		sampleDir: filepath.Join(samplePath, submit.ProblemId),
		outputDir: filepath.Join(tmpPath, "output", submit.SubmitId),

		timeLimit:   submit.TimeLimit,
		memoryLimit: submit.MemoryLimit,
	}
}

func (judger *Judger) Print() {
	log.Println("code type: ", judger.codeType)
	log.Println("code source path: ", judger.codeSourcePath)
	log.Println("bin path: ", judger.binPath)
	log.Println("sample dir: ", judger.sampleDir)
	log.Println("output dir: ", judger.outputDir)
	log.Println("time limit: ", judger.timeLimit)
	log.Println("memory limit: ", judger.memoryLimit)
}

func (judger *Judger) Judge() *[]model.Result {
	lang, err := judger.newLang()
	if err != nil {
		//TODO:
		log.Println("New Lang failed")
		return nil
	}

	if ok := lang.NeedCompile(); ok {
		compiler := lang.Compile()
		compiler.Stdin = os.Stdin
		compiler.Stdout = os.Stdout
		compiler.Stderr = os.Stderr

		if err := compiler.Run(); err != nil {
			log.Println("Compile failed")
			return nil
		}
	}

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
				//TODO:
				log.Println("Make dir failed")
				return nil
			}
		}
	}

	result := make([]model.Result, sampleCount)
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

func (judger *Judger) judgerOneByOne(sampleId string) (_result *model.Result) {

	runner := exec.Command("./sandbox",
		"--code_type", "C",
		"--bin_path", judger.binPath,
		"--seccomp_rule_name", "c_cpp",
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
	}

	// log.Printf("Bytes: %s\n", o.Bytes())
	json.Unmarshal(o.Bytes(), &_result)
	_result.SampleId = sampleId
	// log.Println(_result)
	return _result
}

func (judger *Judger) newLang() (lang.Lang, error) {
	return lang.NewLang(judger.codeType,
		judger.codeSourcePath,
		judger.binPath,
	)
}
