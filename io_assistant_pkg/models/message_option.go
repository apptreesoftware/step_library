package models

import "errors"

type MessageOption struct {
	Id            string `json:"id"`
	Text          string `json:"text"`
	OnCompleteUrl string `json:"onCompleteUrl,omitempty"`
}

func (m MessageOption) Validate() error {
	if m.Text == "" {
		return errors.New("message text is required")
	} else if m.Id == "" {
		return errors.New("id is required")
	}
	return nil
}
