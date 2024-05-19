package character

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/juanmiguelar/character_generator/internal/core/character/character_random_builder"
	"github.com/juanmiguelar/character_generator/internal/core/character/providers"
)

type CharacterBuilderType string

var (
	NORMAL CharacterBuilderType = "NORMAL"
	RANDOM CharacterBuilderType = "RANDOM"
)

func GetCharacterBuilder(builder_type CharacterBuilderType) (providers.CharacterBuilderProvider, error) {
	switch builder_type {
	case RANDOM:
		return character_random_builder.NewRandomCharacterBuilder(rand.New(rand.NewSource(time.Now().UnixNano()))), nil
	default:
		return nil, fmt.Errorf("[%w]: builder type is not supported: [%s]", errors.ErrUnsupported, builder_type)
	}
}
