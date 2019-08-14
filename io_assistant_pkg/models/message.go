package models

import "errors"

type MessageBase struct {
	Id        string           `json:"id,omitempty"`
	Type      MessageType      `json:"type,omitempty"`
	Icon      string           `json:"icon,omitempty"`
	Text      string           `json:"text,omitempty"`
	Options   []*MessageOption `json:"options,omitempty"`
	Label     string           `json:"label,omitempty"`
	YesAnswer string           `json:"yesAnswer,omitempty"`
	NoAnswer  string           `json:"noAnswer,omitempty"`
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
	if m.Type == PromptBool {
		if m.YesAnswer == "" || m.NoAnswer == "" {
			return errors.New("yes/no response values are required for prompt bool types")
		}
	}
	return nil
}
