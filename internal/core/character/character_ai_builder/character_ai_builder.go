package character_ai_builder

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/juanmiguelar/character_generator/internal/core/entities"
)

// characterAIBuilder implements the CharacterBuilder interface.
type characterAIBuilder struct {
	character entities.Character
	apiKey    string
	apiURL    string
	model     string
}

// NewCharacterAIBuilder creates a new characterAIBuilder with the provided OpenAI API key and model.
func NewCharacterAIBuilder(apiKey string) *characterAIBuilder {
	return &characterAIBuilder{
		character: entities.Character{},
		apiKey:    apiKey,
		apiURL:    "https://api.openai.com/v1/chat/completions",
		model:     "gpt-3.5-turbo",
	}
}

// callChatGPTAPI makes a request to the OpenAI API with the provided prompt.
func (cab *characterAIBuilder) callChatGPTAPI(prompt string) (string, error) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"model": cab.model,
		"messages": []map[string]string{
			{"role": "system", "content": "You are a helpful assistant that generates creative content."},
			{"role": "user", "content": prompt},
		},
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", cab.apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+cab.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to call OpenAI API: %s - %s", resp.Status, string(body))
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	if choices, ok := response["choices"].([]interface{}); ok && len(choices) > 0 {
		if choice, ok := choices[0].(map[string]interface{}); ok {
			if message, ok := choice["message"].(map[string]interface{}); ok {
				if content, ok := message["content"].(string); ok {
					return content, nil
				}
			}
		}
	}

	return "", errors.New("invalid response format")
}

// SetName sets a random name for the character using ChatGPT API.
func (cab *characterAIBuilder) SetName() error {
	prompt := "Generate a creative name for a fantasy character."
	name, err := cab.callChatGPTAPI(prompt)
	if err != nil {
		return err
	}
	cab.character.Name = name
	return nil
}

// SetAge sets a random age for the character locally.
func (cab *characterAIBuilder) SetAge() error {
	rand.Seed(time.Now().UnixNano())
	cab.character.Age = rand.Intn(43) + 18 // Random age between 18 and 60
	return nil
}

// SetSex sets a random sex for the character using ChatGPT API.
func (cab *characterAIBuilder) SetSex() error {
	prompt := "Generate a random sex (Male or Female) for a fantasy character."
	sex, err := cab.callChatGPTAPI(prompt)
	if err != nil {
		return err
	}
	cab.character.Sex = sex
	return nil
}

// SetBackdropStory sets a random backdrop story for the character using ChatGPT API.
func (cab *characterAIBuilder) SetBackdropStory() error {
	prompt := fmt.Sprintf(`Generate a creative backdrop story for a fantasy character. 
	Which name is [%s], 
	age is [%d], 
	gender is [%s] 
	and clothes_style is [%s]`, 
	cab.character.Name, cab.character.Age, cab.character.Sex, cab.character.ClothesStyle)
	story, err := cab.callChatGPTAPI(prompt)
	if err != nil {
		return err
	}
	cab.character.BackdropStory = story
	return nil
}

// SetClothesStyle sets a random clothes style for the character using ChatGPT API.
func (cab *characterAIBuilder) SetClothesStyle() error {
	prompt := "Generate a random clothes style for a fantasy character."
	style, err := cab.callChatGPTAPI(prompt)
	if err != nil {
		return err
	}
	cab.character.ClothesStyle = style
	return nil
}

// Build returns the built character.
func (cab *characterAIBuilder) Build() entities.Character {
	return cab.character
}


