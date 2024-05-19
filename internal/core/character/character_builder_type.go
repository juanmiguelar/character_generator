package character

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/juanmiguelar/character_generator/internal/core/character/character_ai_builder"
	"github.com/juanmiguelar/character_generator/internal/core/character/character_random_builder"
	"github.com/juanmiguelar/character_generator/internal/core/character/providers"
)

type CharacterBuilderType string

var (
	AI CharacterBuilderType = "AI"
	RANDOM CharacterBuilderType = "RANDOM"
)

func GetCharacterBuilder(builder_type CharacterBuilderType) (providers.CharacterBuilderProvider, error) {
	switch builder_type {
	case RANDOM:
		return character_random_builder.NewRandomCharacterBuilder(rand.New(rand.NewSource(time.Now().UnixNano()))), nil
	case AI:
		return character_ai_builder.NewCharacterAIBuilder(os.Getenv("OPENAI_API_KEY")), nil
	default:
		return nil, fmt.Errorf("[%w]: builder type is not supported: [%s]", errors.ErrUnsupported, builder_type)
	}
}
