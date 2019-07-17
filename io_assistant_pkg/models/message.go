package models

import "errors"

type MessageBase struct {
	Id      string           `json:"id"`
	Type    MessageType      `json:"type"`
	Icon    string           `json:"icon"`
	Text    string           `json:"text"`
	Options []*MessageOption `json:"options"`
	Label   string           `json:"label"`
}

type MessageInput struct {
	MessageBase
	OnCompleteWorkflow string `json:"onCompleteWorkflow"`
}

type MessageData struct {
	MessageBase
	OnCompleteUrl string `json:"onCompleteUrl"`
}

func (m MessageBase) ValidateMessageInput() error {
	if m.Type == "" {
		m.Type = Message
	}
	if m.Id == "" && m.Type.IsPrompt() {
		return errors.New("prompt message must contain an ID")
	}
	if m.Text == "" {
		return errors.New("message text is required")
	}
	if m.Type == PromptSingleSelect {
		if m.Options == nil || len(m.Options) == 0 {
			return errors.New("options are required for a promptSingleSelect message type")
		}
		for _, option := range m.Options {
			if option.Id == "" {
				return errors.New("options must have an ID")
			}
			if option.Text == "" {
				return errors.New("options must have text for display")
			}
		}
	}
	return nil
}
