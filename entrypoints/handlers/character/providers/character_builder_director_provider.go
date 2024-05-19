package providers

import (
	"github.com/juanmiguelar/character_generator/internal/core/character/providers"
	"github.com/juanmiguelar/character_generator/internal/core/entities"
)

// CharacterDirectorProvider defines the interface for building characters.
type CharacterDirectorProvider interface {
	Build() (entities.Character, error)
	SetBuilder(b providers.CharacterBuilderProvider)
}
