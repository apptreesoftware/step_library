package models

type MessageOption struct {
	Id   string      `json:"id"`
	Text string      `json:"text"`
	Type MessageType `json:"type"`
}
