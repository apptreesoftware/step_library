package models

type MessageType string

const (
	Message            MessageType = "message"
	PromptBool         MessageType = "promptBool"
	PromptText         MessageType = "promptText"
	PromptInt          MessageType = "promptInt"
	PromptForm         MessageType = "promptForm"
	PromptSingleSelect MessageType = "promptSingleSelect"
	CaptureImage       MessageType = "captureImage"
	CaptureLocation    MessageType = "captureLocation"
	CaptureBarcode     MessageType = "captureBarcode"
)

func (m MessageType) IsPrompt() bool {
	return m != Message
}
