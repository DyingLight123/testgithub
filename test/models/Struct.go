package models

type APIRequestCommon struct {
	Id      string `json:"id"`
	User    string `json:"user"`
	Version string `json:"version"`
}

type APIResponseCommon struct {
	Id 		string 	`json:"id"`
	Code 	int 	`json:"code"`
	Desc 	string	`json:"desc"`
}
