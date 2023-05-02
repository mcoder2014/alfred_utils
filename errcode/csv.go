package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"

	"github.com/jszwec/csvutil"
)

// LoadErrorCodeFromCSV 从 csv 文件中解析内容
func LoadErrorCodeFromCSV(path string) ([]*ErrCodeInfo, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	content = removeBOM(content)

	var tmp []*ErrCodeInfo
	if err = csvutil.Unmarshal(content, &tmp); err != nil {
		return nil, err
	}
	var res []*ErrCodeInfo
	for _, errcode := range tmp {
		if errcode.Code == "" {
			continue
		}
		res = append(res, errcode)
	}
	return res, nil
}

func removeBOM(content []byte) []byte {
	scanner := bufio.NewScanner(bytes.NewReader(content))
	var res = make([]byte, 0, len(content))

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, "\uFEFF", "", -1)
		res = append(res, []byte(line)...)
		res = append(res, []byte("\n")...)
	}
	return res
}
