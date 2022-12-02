package main

import (
	"flag"

	"alfred_utils/timex/model"

	aw "github.com/deanishe/awgo"
	"github.com/sirupsen/logrus"
)

// Workflow is the main API
var wf *aw.Workflow

func Init() {
	// Create a new Workflow using default settings.
	// Critical settings are provided by Alfred via environment variables,
	// so this *will* die in flames if not run in an Alfred-like environment.
	wf = aw.New()
}

func Serve(req *model.Request) {
	var err error
	switch req.Method {
	case model.Gen:
		err = Gen(wf)
	case model.Trans:
		err = Trans(wf, req.Param)
	default:
		wf.NewItem("Not support current method").Subtitle("method: " + req.Method).Valid(false)
	}
	if err != nil {
		wf.FatalError(err)
	}
}

func main() {
	req := getRequest()

	// Wrap your entry point with Run() to catch and log panics and
	// show an error in Alfred instead of silently dying
	Init()
	wf.Run(func() {
		Serve(&req)

		// Send results to Alfred
		wf.SendFeedback()
	})
}

func getRequest() model.Request {
	var res model.Request

	flag.StringVar(&res.Method, "method", "gen", "方法，支持 gen, trans")
	flag.StringVar(&res.Param, "param", "123456789", "时间戳、2022/12/01等")

	// 从arguments中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。未注册却使用flag -help时，会返回ErrHelp。
	flag.Parse()
	logrus.Infof("Receive request: %+v", res)
	return res
}
