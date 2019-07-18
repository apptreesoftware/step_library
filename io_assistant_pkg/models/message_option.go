package models

import "errors"

type MessageOption struct {
	Id   string      `json:"id"`
	Text string      `json:"text"`
	Type MessageType `json:"type,omitempty"`
}

func (m MessageOption) Validate() error {
	if m.Type.IsPrompt() && m.Id == "" {
		return errors.New("message ID is required for prompt messages")
	}
	if m.Text == "" {
		return errors.New("message text is required")
	}
	return nil
}
