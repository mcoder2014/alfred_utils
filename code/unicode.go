package main

import (
	"fmt"
	"strconv"
	"strings"
)

type UnicodeEngine struct {
}

func (c *UnicodeEngine) Name() string {
	return "unicode"
}

func (c *UnicodeEngine) Encode(content []byte) ([]byte, error) {
	textQuoted := strconv.QuoteToASCII(string(content))
	textUnquoted := textQuoted[1 : len(textQuoted)-1]
	return []byte(textUnquoted), nil
}

func (c *UnicodeEngine) Decode(content []byte) ([]byte, error) {
	str := strings.ToLower(string(content))
	if !c.isUnicode(str) {
		return nil, fmt.Errorf("content is not unicode encoded")
	}

	str, err := strconv.Unquote(strings.Replace(strconv.Quote(str), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

func (c *UnicodeEngine) isUnicode(str string) bool {
	count := strings.Count(str, "\\u")
	return count != 0
}
