package models

type PromptRequest struct {
	Prompt string `json:"prompt" binding:"required"`
}

type PromptResponse struct {
	Prompt   string `json:"prompt"`
	Response string `json:"response"`
}
