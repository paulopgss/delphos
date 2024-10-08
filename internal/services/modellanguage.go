package services

import (
	"bufio"
	"bytes"
	"encoding/json"
	"net/http"
)

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

func SendPromptToModelStream(prompt string) (<-chan string, error) {
	payload := OllamaRequest{
		Model:  "llama3.1",
		Prompt: prompt,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("http://127.0.0.1:11434/api/generate", "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	responseStream := make(chan string)

	go func() {
		defer resp.Body.Close()
		defer close(responseStream)

		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			var part map[string]interface{}
			if err := json.Unmarshal(scanner.Bytes(), &part); err == nil {
				if text, ok := part["response"].(string); ok {
					responseStream <- text
				}
				if done, ok := part["done"].(bool); ok && done {
					break
				}
			}
		}
	}()

	return responseStream, nil
}
