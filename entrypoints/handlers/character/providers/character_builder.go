package providers

import (
	"github.com/juanmiguelar/character_generator/internal/core/character/providers"
	"github.com/juanmiguelar/character_generator/internal/core/entities"
)

// CharacterBuilderDirector defines the interface for building characters.
type CharacterBuilderDirector interface {
	Build() (entities.Character, error)
	SetBuilder(b providers.CharacterBuilderProvider)
}
