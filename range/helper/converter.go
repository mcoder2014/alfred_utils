package helper

import (
	"strconv"

	"alfred_utils/range/model"

	aw "github.com/deanishe/awgo"
)

func ToIntegerReq(req *model.Request, wf *aw.Workflow) (*model.IntRequest, error) {
	if req == nil {
		return nil, nil
	}
	var res model.IntRequest
	var err error
	res.Wf = wf
	res.Start, err = strconv.Atoi(req.Start)
	if err != nil {
		return nil, err
	}
	res.Interval, err = strconv.Atoi(req.Interval)
	if err != nil {
		return nil, err
	}
	res.End, err = strconv.Atoi(req.End)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
