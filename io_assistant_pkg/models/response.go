package models

type FulfillmentResponse struct {
	Message      *MessageData           `json:"message,omitempty"`
	MessageGroup *MessageGroup          `json:"messageGroup,omitempty"`
	UserContext  map[string]interface{} `json:"userContext,omitempty"`
	Context      map[string]interface{} `json:"context,omitempty"`
}

func NewMessageGroupResponse(messages []MessageOption, workflowUrl string, userContext map[string]interface{}, context map[string]interface{}) FulfillmentResponse {
	return FulfillmentResponse{
		MessageGroup: &MessageGroup{
			Messages:      messages,
			OnCompleteUrl: workflowUrl,
		},
		UserContext: userContext,
		Context:     context,
	}
}

func NewMessageResponse(input MessageBase, workflowUrl string, userContext map[string]interface{}, context map[string]interface{}) FulfillmentResponse {
	return FulfillmentResponse{
		Message: &MessageData{
			MessageBase:   input,
			OnCompleteUrl: workflowUrl,
		},
		UserContext: userContext,
		Context:     context,
	}
}
