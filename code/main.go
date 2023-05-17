package main

import (
	"flag"
	"fmt"
	"os"

	aw "github.com/deanishe/awgo"
	"github.com/sirupsen/logrus"
)

var wf *aw.Workflow

func Init() {
	// Create a new Workflow using default settings.
	// Critical settings are provided by Alfred via environment variables,
	// so this *will* die in flames if not run in an Alfred-like environment.
	wf = aw.New()
}

func main() {
	req := getRequest()

	// Wrap your entry point with Run() to catch and log panics and
	// show an error in Alfred instead of silently dying
	Init()
	wf.Run(func() {
		err := Serve(&req)
		if err != nil {
			wf.NewItem("Program failed.").Subtitle(fmt.Sprintf("err:%v", err)).Valid(false)
			wf.FatalError(err)
		}

		// Send results to Alfred
		wf.SendFeedback()
	})
}

func getRequest() Request {
	var res Request
	flag.StringVar(&res.Command, "command", "encode", "编码命令，encode, decode")
	flag.StringVar(&res.Content, "content", "", "入参内容")

	// 从arguments中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。未注册却使用flag -help时，会返回ErrHelp。
	flag.Parse()
	pwd, _ := os.Getwd()
	logrus.Infof("PWD:%v", pwd)
	logrus.Infof("Receive request: %+v", res)
	return res
}
