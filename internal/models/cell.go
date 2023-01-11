package models

type Cell struct {
	Row     int        `json:"row,omitempty"`
	Column  int        `json:"column,omitempty"`
	Stage   int        `json:"stage,omitempty"`
	StrName string     `json:"str_name,omitempty"`
	Content []*Product `json:"content,omitempty"`
}
