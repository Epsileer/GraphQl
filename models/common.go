package models

type ReqQuery struct {
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables"`
	OperationName string                 `json:"operation_name"`
}

type User struct {
	Id     int
	Name   string
	Status string
}
