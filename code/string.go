package main

import "strings"

var escapeCharacterMap = map[string]rune{
	"\\\\": '\\',
	"\\'":  '\'',
	"\\t":  '\t',
	"\\n":  '\n',
	"\\r":  '\r',
	"\\b":  '\b',
	"\\f":  '\f',
	"\\\"": '"',
	"\\?":  '?',
}

var charEscapeMap = map[rune]string{}

func init() {
	charEscapeMap = make(map[rune]string, len(escapeCharacterMap))
	for k, v := range escapeCharacterMap {
		charEscapeMap[v] = k
	}
}

type StringEncoderEngine struct {
}

func (s *StringEncoderEngine) Name() string {
	return "string encode"
}

func (s *StringEncoderEngine) Encode(bytes []byte) ([]byte, error) {
	input := string(bytes)
	res := StringEncode(input)
	return []byte(res), nil
}

func (s *StringEncoderEngine) Decode(bytes []byte) ([]byte, error) {
	input := string(bytes)
	res := StringDecode(input)
	return []byte(res), nil
}

func StringDecode(input string) string {
	var sb strings.Builder
	var tmpChar rune
	var meetLine bool
	for i := 0; i < len(input); i++ {
		ch := rune(input[i])
		switch meetLine {
		case true:
			meetLine = false
			cc := string([]rune{tmpChar, ch})
			if val, ok := escapeCharacterMap[cc]; ok {
				sb.WriteRune(val)
			} else {
				sb.WriteString(cc)
			}
			continue
		case false:
			if ch == '\\' {
				meetLine = true
				tmpChar = ch
				continue
			}
			sb.WriteRune(ch)
		}
	}
	return sb.String()
}

func StringEncode(input string) string {
	var sb strings.Builder
	for i := 0; i < len(input); i++ {
		ch := rune(input[i])
		val, ok := charEscapeMap[ch]
		if ok {
			sb.WriteString(val)
			continue
		}
		sb.WriteRune(ch)
	}
	return sb.String()
}
