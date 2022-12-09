package model

import aw "github.com/deanishe/awgo"

type NumType = string

const (
	Integer NumType = "int"
)

type Request struct {
	Type     NumType `json:"type"`
	Start    string  `json:"start"`
	Interval string  `json:"interval"`
	End      string  `json:"end"`
}

type IntRequest struct {
	Start    int `json:"start"`
	Interval int `json:"interval"`
	End      int `json:"end"`

	Wf *aw.Workflow
}
