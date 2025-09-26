package llm

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

const (
	roleUser   = "user"
	chatApi    = "/v1/chat/completions"
	chatModel  = "tencent/Hunyuan-MT-7B"
	voiceApi   = "/v1/audio/speech"
	voiceModel = "fnlp/MOSS-TTSD-v0.5"
)

type LLM struct {
	BaseUrl string
	ApiKey  string
}

func NewLLM(url string, apiKey string) *LLM {
	return &LLM{BaseUrl: url, ApiKey: apiKey}
}

type ChatPayload struct {
	Model    string            `json:"model"`
	Messages []*RequestMessage `json:"messages"`
}
type RequestMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
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

func (llm *LLM) CreateChat(content string) (*ChatCompletionResponse, error) {
	request, err := http.NewRequest(http.MethodPost, llm.BaseUrl+chatApi, nil)
	if err != nil {
		return nil, err
	}
	auth := strings.Join([]string{"Bearer", llm.ApiKey}, " ")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", auth)

	body, err := json.Marshal(&ChatPayload{
		Model: chatModel,
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

type VoicePayload struct {
	MaxTokens      int     `json:"max_tokens"`
	ResponseFormat string  `json:"response_format"`
	SampleRate     int     `json:"sample_rate"`
	Stream         bool    `json:"stream"`
	Speed          float64 `json:"speed"`
	Gain           int     `json:"gain"`
	Model          string  `json:"model"`
	Input          string  `json:"input"`
	Voice          string  `json:"voice"`
}

func newDefaultVoicePayload(input string, voice string) *VoicePayload {
	return &VoicePayload{
		MaxTokens:      2048,
		ResponseFormat: "mp3",
		SampleRate:     32000,
		Stream:         true,
		Speed:          1,
		Gain:           0,
		Model:          voiceModel,
		Input:          input,
		Voice:          voice,
	}
}

func (llm *LLM) TransferTextToVoice(text string, voice string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodPost, llm.BaseUrl+voiceApi, nil)
	if err != nil {
		return nil, err
	}
	auth := strings.Join([]string{"Bearer", llm.ApiKey}, " ")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", auth)

	body, err := json.Marshal(newDefaultVoicePayload(text, voice))
	if err != nil {
		return nil, err
	}
	request.Body = io.NopCloser(bytes.NewBuffer(body))
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bin, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bin, nil
}
