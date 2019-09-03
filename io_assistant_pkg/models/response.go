package models

type FulfillmentResponse struct {
	Message      map[string]interface{} `json:"message,omitempty"`
	MessageGroup *MessageGroup          `json:"messageGroup,omitempty"`
	UserContext  map[string]interface{} `json:"userContext,omitempty"`
	Context      map[string]interface{} `json:"context,omitempty"`
	Complete     bool                   `json:"complete"`
}

func NewMessageGroupResponse(messages []map[string]interface{}, workflowUrl string, userContext map[string]interface{}, context map[string]interface{}) FulfillmentResponse {
	return FulfillmentResponse{
		MessageGroup: &MessageGroup{
			Messages:      messages,
			OnCompleteUrl: workflowUrl,
		},
		UserContext: userContext,
		Context:     context,
	}
}

func NewMessageResponse(message map[string]interface{}, userContext map[string]interface{}, context map[string]interface{}, complete bool) FulfillmentResponse {
	return FulfillmentResponse{
		Message:     message,
		UserContext: userContext,
		Context:     context,
		Complete:    complete,
	}
}
