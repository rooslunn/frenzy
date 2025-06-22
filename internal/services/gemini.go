package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

const (
	AI_MODEL     = "gemini-2.0-flash"
	TOKENS_COUNT = 200

	ErrPromptKeyMissing = "key does not exist: "
)

var AI_PROMTPS = map[string]string{
	"en": "Please, write a short story (about %d characters) using phrase \"%s\"",
	"ru": "Напиши, пожалуйста, короткий рассказ (около %d символов), используя фразу \"%s\"",
}

func GenerateFenzyText(ctx context.Context, aiKey, lang, phrase string) (string, error) {

	aiPrompt, ok := AI_PROMTPS[lang]
	if !ok {
		return "", errors.New(ErrPromptKeyMissing + lang)
	}

	aiPrompt = fmt.Sprintf(aiPrompt, TOKENS_COUNT, phrase)

	client, err := genai.NewClient(ctx, option.WithAPIKey(aiKey))
	if err != nil {
		return "", err
	}

	model := client.GenerativeModel(AI_MODEL)
	resp, err := model.GenerateContent(ctx, genai.Text(aiPrompt))
	if err != nil {
		return "", err
	}

	var s string

	for _, candidate := range resp.Candidates {
		if candidate != nil {
			if candidate.Content.Parts != nil {
				s = string(candidate.Content.Parts[0].(genai.Text))
			}
		}
	}

	return s, nil
}

func GenerateFenzyPicture(ctx context.Context, aiKey, prompt string) (string, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(aiKey))
	if err != nil {
		return "", err
	}

	model := client.GenerativeModel(AI_MODEL)
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	var s string

	for _, candidate := range resp.Candidates {
		if candidate != nil {
			if candidate.Content.Parts != nil {
				s = fmt.Sprintf("Output: %s", string(candidate.Content.Parts[0].(genai.Text)))
			}
		}
	}

	return s, nil
}
