package models

type MessageGroup struct {
	OnCompleteUrl string          `json:"onCompleteUrl"`
	Messages      []MessageOption `json:"messages"`
}
