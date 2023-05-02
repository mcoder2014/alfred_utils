package main

import (
	"fmt"
	"strings"

	"github.com/mcoder2014/alfred_utils/utils"
)

const RelateLimit = 10

func Serve(req *Request) error {
	errcodeList, err := LoadErrorCodeFromCSV(req.ErrCodeFile)
	if err != nil {
		return err
	}
	res := searchRelateCodeFromList(req.Query, errcodeList)
	return InsertToWF(res)
}

func InsertToWF(res []*ErrCodeInfo) error {
	for _, errcode := range res {
		title := fmt.Sprintf("Code: %v Name:%v Category:%v ", errcode.Code, errcode.Name, errcode.Category)
		description := fmt.Sprintf("Description: %v ", errcode.Description)
		copyContent := fmt.Sprintf("code:%v\tcategory:%v\tservice:%v\tdescription:%v",
			errcode.Code, errcode.Category, errcode.Service, errcode.Description)
		utils.NewCopyableItemWithDescription(wf, title, description, copyContent)
	}
	return nil
}

func searchRelateCodeFromList(query string, errcodeList []*ErrCodeInfo) []*ErrCodeInfo {
	var res []*ErrCodeInfo
	var relateMap = make(map[int][]*ErrCodeInfo, RelateLimit)
	for _, errcode := range errcodeList {
		if errcode.Code == query {
			res = append(res, errcode)
			continue
		}
		if strings.Contains(errcode.Code, query) || strings.Contains(query, errcode.Code) {
			delta := len(errcode.Code) - len(query)
			if delta < 0 {
				delta = -delta
			}
			if delta > RelateLimit {
				continue
			}
			relateMap[delta] = append(relateMap[delta], errcode)
		}
	}
	for i := 1; i <= RelateLimit; i++ {
		res = append(res, relateMap[i]...)
	}

	return res
}
