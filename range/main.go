package main

import (
	"flag"

	"alfred_utils/range/helper"
	"alfred_utils/range/model"
	"alfred_utils/range/service"

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
	switch req.Type {
	case model.Integer:
		r, err := helper.ToIntegerReq(req, wf)
		if err != nil {
			wf.FatalError(err)
		}
		err = service.GenInteger(r)
	default:
		wf.NewItem("Not support current type").Subtitle("type: " + req.Type).Valid(false)
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

	flag.StringVar(&res.Type, "type", "int", "生成的类型，支持 int")
	flag.StringVar(&res.Start, "start", "1", "起始值")
	flag.StringVar(&res.Interval, "interval", "1", "间隔")
	flag.StringVar(&res.End, "end", "10", "结束值")

	// 从arguments中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。未注册却使用flag -help时，会返回ErrHelp。
	flag.Parse()
	logrus.Infof("Receive request: %+v", res)
	return res
}
