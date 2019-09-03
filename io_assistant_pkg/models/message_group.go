package models

type MessageGroup struct {
	OnCompleteUrl string                   `json:"onCompleteUrl"`
	Messages      []map[string]interface{} `json:"messages"`
}
