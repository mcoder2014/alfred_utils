package main

const (
	Code        = "code"
	Name        = "name"
	Category    = "category"
	Service     = "service"
	Description = "description"
	RefLink     = "ref_link"
)

type ErrCodeInfo struct {
	Code        string `json:"code" csv:"code,omitempty"`
	Name        string `json:"name" csv:"name,omitempty"`
	Category    string `json:"category" csv:"category,omitempty"`
	Service     string `json:"service" csv:"service,omitempty"`
	Description string `json:"description" csv:"description,omitempty"`
	RefLink     string `json:"ref_link" csv:"ref_link,omitempty"`
}

type Request struct {
	Query       string
	ErrCodeFile string
}
