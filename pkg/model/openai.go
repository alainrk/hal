package model

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

// OpenAIModel implements the Model interface for OpenAI
type OpenAIModel struct {
	client    *openai.Client
	modelName string
}

// NewOpenAIModel creates a new OpenAI model
func NewOpenAIModel(apiKey string, modelName string) *OpenAIModel {
	return &OpenAIModel{
		client:    openai.NewClient(apiKey),
		modelName: modelName,
	}
}

// Generate implements Model.Generate
func (m *OpenAIModel) Generate(ctx context.Context, prompt string, options *GenerateOptions) (*Response, error) {
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: prompt,
		},
	}

	if options.SystemPrompt != "" {
		messages = append([]openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: options.SystemPrompt,
			},
		}, messages...)
	}

	req := openai.ChatCompletionRequest{
		Model:       m.modelName,
		Messages:    messages,
		Temperature: options.Temperature,
		MaxTokens:   options.MaxTokens,
		TopP:        options.TopP,
		Stop:        options.StopSequences,
	}

	resp, err := m.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("openai completion failed: %w", err)
	}

	if len(resp.Choices) == 0 {
		return nil, fmt.Errorf("no completion choices returned")
	}

	return &Response{
		Content:      resp.Choices[0].Message.Content,
		TokensUsed:   resp.Usage.TotalTokens,
		FinishReason: string(resp.Choices[0].FinishReason),
		Metadata: map[string]any{
			"model": m.modelName,
			"id":    resp.ID,
		},
	}, nil
}

// GetName returns the model name
func (m *OpenAIModel) GetName() string {
	return m.modelName
}
