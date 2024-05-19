package character

import (
	"fmt"
	"errors"
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
		return newRandomCharacterBuilder(), nil
	default:
		return nil, fmt.Errorf("[%w]: builder type is not supported: [%s]", errors.ErrUnsupported, builder_type)
	}
}
