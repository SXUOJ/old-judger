package judge

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Sxu-Online-Judge/judger/lang"
)

type Compiler struct {
	codeType       string
	codeSourcePath string
	binPath        string
}

func NewCompiler(submit *Submit) *Compiler {
	return &Compiler{
		codeType:       submit.CodeType,
		codeSourcePath: submit.CodeSourcePath,
		binPath:        filepath.Join(tmpPath, submit.SubmitId),
	}
}

func (c *Compiler) Run() (_result *CompileResult) {
	lang, err := lang.NewLang(c.codeType, c.codeSourcePath, c.binPath)
	if err != nil {
		log.Println("New Lang failed")
		return &CompileResult{}
	}

	info, err := user.Lookup("compiler")
	if err != nil {
		log.Println("Lookup failed")
		return &CompileResult{}
	}

	// env := os.Getenv("PATH")

	var o bytes.Buffer
	var e bytes.Buffer
	if ok := lang.NeedCompile(); ok {
		compiler := exec.Command("./sandbox",
			"--bin_path", lang.Bin(),
			//TODO: seccomp
			// "--seccomp_rule_name", "general",
			// "--input_path", c.codeSourcePath,
			"--real_time_limit", lang.RealTimeLimit(),
			"--cpu_time_limit", lang.CpuTimeLimit(),
			"--max_memory", lang.MemoryLimit(),
			"--max_stack", strconv.FormatInt(128*1024*1024, 10),
			"--max_output_size", strconv.FormatInt(20*1024*1024, 10),
			"--args", lang.Args(),
			// "--env", env,
			"--uid", info.Uid,
			"--gid", info.Gid,
		)

		log.Println(compiler.Args)

		compiler.Stdin = os.Stdin
		compiler.Stdout = &o
		compiler.Stderr = &e

		if err := compiler.Run(); err != nil {
			log.Println("Compile failed")
			log.Printf("%s", o.Bytes())
			log.Printf("%s", e.Bytes())
			return &CompileResult{
				Status:   GetJudgeStatus(StatusCE),
				ErrorInf: strings.Join([]string{o.String(), e.String()}, "\n"),
			}
		}
	}
	json.Unmarshal(o.Bytes(), &_result)
	return _result
}
