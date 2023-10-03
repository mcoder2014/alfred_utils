package main

import (
	"fmt"

	"github.com/mcoder2014/alfred_utils/utils"
)

var defaultEngines = []CodeEngine{
	utils.GPtr(UnicodeEngine{}),
	utils.GPtr(StringEncoderEngine{}),
}

func Serve(req *Request) error {
	if req.Content == "" {
		return nil
	}

	for _, engine := range defaultEngines {
		switch req.Command {
		case CommandDecode:
			res, err := engine.Decode([]byte(req.Content))
			if err != nil {
				InsertErrorToWF(engine, err)
				continue
			}
			InsertSuccessToWF(engine, res)
		case CommandEncode:
			res, err := engine.Encode([]byte(req.Content))
			if err != nil {
				InsertErrorToWF(engine, err)
				continue
			}
			InsertSuccessToWF(engine, res)
		default:
			InsertErrorToWF(engine, fmt.Errorf("not support command: %s", req.Command))
		}
	}
	return nil
}

func InsertErrorToWF(c CodeEngine, err error) {
	title := fmt.Sprintf("[%v failed] meet error: %s", c.Name(), err.Error())
	utils.NewCopyableItemWithDescription(wf, title, title, title)
}

func InsertSuccessToWF(c CodeEngine, res []byte) {
	title := fmt.Sprintf("[%v ] result:%v", c.Name(), string(res))
	copyContent := string(res)
	utils.NewCopyableItem(wf, title, copyContent)
}

type CodeEngine interface {
	Name() string
	Encode([]byte) ([]byte, error)
	Decode([]byte) ([]byte, error)
}
