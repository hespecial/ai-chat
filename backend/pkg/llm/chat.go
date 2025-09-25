package llm

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

const (
	roleUser = "user"
)

type LLM struct {
	Url    string
	ApiKey string
	Model  string
}

type Payload struct {
	Model    string            `json:"model"`
	Messages []*RequestMessage `json:"messages"`
}

type RequestMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func NewLLM(url string, apiKey string, model string) *LLM {
	return &LLM{Url: url, ApiKey: apiKey, Model: model}
}

func (llm *LLM) Call(content string) (*ChatCompletionResponse, error) {
	request, err := http.NewRequest(http.MethodPost, llm.Url, nil)
	if err != nil {
		return nil, err
	}
	auth := strings.Join([]string{"Bearer", llm.ApiKey}, " ")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", auth)

	body, err := json.Marshal(&Payload{
		Model: llm.Model,
		Messages: []*RequestMessage{
			{
				Role:    roleUser,
				Content: content,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	request.Body = io.NopCloser(bytes.NewBuffer(body))
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ChatCompletionResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

type ChatCompletionResponse struct {
	ID      string `json:"id"`
	Choices []struct {
		Message struct {
			Role             string `json:"role"`
			Content          string `json:"content"`
			ReasoningContent string `json:"reasoning_content"`
			ToolCalls        []struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				Function struct {
					Name      string `json:"name"`
					Arguments string `json:"arguments"`
				} `json:"function"`
			} `json:"tool_calls"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Object  string `json:"object"`
}
