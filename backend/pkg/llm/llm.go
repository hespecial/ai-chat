package llm

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	roleUser  = "user"
	chatApi   = "/v1/chat/completions"
	chatModel = "doubao-seed-1.6-thinking"
	//voiceApi   = "/v1/audio/speech"
	//voiceModel = "fnlp/MOSS-TTSD-v0.5"

	voiceApi = "/v1/voice/tts"
)

type LLM struct {
	BaseUrl string
	ApiKey  string
	wssUrl  url.URL
}

func NewLLM(baseUrl string, apiKey string) *LLM {
	return &LLM{
		BaseUrl: baseUrl,
		ApiKey:  apiKey,
		wssUrl:  url.URL{Scheme: "wss", Host: baseUrl[len("https://"):], Path: voiceApi},
	}
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

type TtsPayload struct {
	Audio   *Audio   `json:"audio"`
	Request *Request `json:"request"`
}
type Audio struct {
	VoiceType  string  `json:"voice_type"`
	Encoding   string  `json:"encoding"`
	SpeedRatio float64 `json:"speed_ratio"`
}
type Request struct {
	Text string `json:"text"`
}
type RelayTTSResponse struct {
	Reqid     string `json:"reqid"`
	Operation string `json:"operation"`
	Sequence  int    `json:"sequence"`
	Data      string `json:"data"`
	Addition  struct {
		Duration string `json:"duration"`
	} `json:"addition"`
}

func newDefaultTtsPayload(input string, voice string) *TtsPayload {
	return &TtsPayload{
		Audio: &Audio{
			VoiceType:  voice,
			Encoding:   "mp3",
			SpeedRatio: 1.0,
		},
		Request: &Request{
			Text: input,
		},
	}
}

func (llm *LLM) TransferTextToVoice(text string, voice string) ([]byte, error) {
	return llm.wssStreamToBytes(text, voice)
}

func (llm *LLM) wssStreamToBytes(text, voice string) ([]byte, error) {
	payload := newDefaultTtsPayload(text, voice)

	input, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	var header = http.Header{
		"Authorization": []string{fmt.Sprintf("Bearer %s", llm.ApiKey)},
		"VoiceType":     []string{voice},
	}
	c, _, err := websocket.DefaultDialer.Dial(llm.wssUrl.String(), header)
	if err != nil {
		return nil, fmt.Errorf("dial err: %v", err)
	}
	defer c.Close()

	err = c.WriteMessage(websocket.BinaryMessage, input)
	if err != nil {
		return nil, fmt.Errorf("write message fail: %v", err)
	}

	var audio []byte
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			return nil, fmt.Errorf("read message fail: %v", err)
		}

		var resp RelayTTSResponse
		err = json.Unmarshal(message, &resp)
		if err != nil {
			fmt.Println("unmarshal fail, err:", err.Error())
			continue
		}

		d, err := base64.StdEncoding.DecodeString(resp.Data)
		if err != nil {
			fmt.Println("decode fail, err:", err.Error())
			continue
		}

		audio = append(audio, d...)

		// 根据响应判断是否结束（假设Sequence < 0表示结束）
		if resp.Sequence < 0 {
			break
		}
	}

	return audio, nil
}
