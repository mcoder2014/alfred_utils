package service

import (
	"strconv"
	"strings"

	"alfred_utils/range/model"
	"alfred_utils/utils"
)

func GenInteger(req *model.IntRequest) error {
	if req == nil {
		return nil
	}

	array := getIntArray(req.Start, req.Interval, req.End)

	var res strings.Builder
	for i, v := range array {
		if i != 0 {
			res.WriteString(", ")
		}
		res.WriteString(strconv.Itoa(v))
	}

	utils.NewCopyableItem(req.Wf, "生成序列", res.String())

	return nil
}

func getIntArray(start int, interval int, end int) []int {
	var res []int
	for i := start; i <= end; i += interval {
		res = append(res, i)
	}
	return res
}
