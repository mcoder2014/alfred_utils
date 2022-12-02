package model

type Method = string

const (
	Gen   Method = "gen"
	Trans Method = "trans"
)

type Request struct {
	Method Method `json:"method"`
	Param  string `json:"param"`
}
