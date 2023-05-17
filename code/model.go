package main

type Request struct {
	Command string `json:"command"`
	Content string `json:"content"`
}

const (
	CommandEncode = "encode"
	CommandDecode = "decode"
)
