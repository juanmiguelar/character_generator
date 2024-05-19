package character

import (
	"math/rand"
	"github.com/juanmiguelar/character_generator/internal/core/entities"
)

// randomCharacterBuilder implements the CharacterBuilder interface.
type randomCharacterBuilder struct {
	character entities.Character
}

// NewrandomCharacterBuilder creates a new randomCharacterBuilder.
func newRandomCharacterBuilder() *randomCharacterBuilder {
	return &randomCharacterBuilder{
		character: entities.Character{},
	}
}

// SetName sets a random name for the character.
func (rcb *randomCharacterBuilder) SetName() error {
	// Generate random name (for demonstration purposes, just using a fixed set of names)
	names := []string{"Alice", "Bob", "Charlie", "Diana", "Emma"}
	rcb.character.Name = names[rand.Intn(len(names))]
	return nil
}

// SetAge sets a random age for the character.
func (rcb *randomCharacterBuilder) SetAge() error {
	// Generate random age between 18 and 60
	rcb.character.Age = rand.Intn(43) + 18
	return nil
}

// SetSex sets a random sex for the character.
func (rcb *randomCharacterBuilder) SetSex() error {
	// Generate random sex (for demonstration purposes, just using Male or Female)
	sexes := []string{"Male", "Female"}
	rcb.character.Sex = sexes[rand.Intn(len(sexes))]
	return nil
}

// SetBackdropStory sets a random backdrop story for the character.
func (rcb *randomCharacterBuilder) SetBackdropStory() error {
	// Generate random backdrop story (for demonstration purposes, just using a fixed set of stories)
	stories := []string{"Once upon a time...", "In a land far away...", "In a world of magic..."}
	rcb.character.BackdropStory = stories[rand.Intn(len(stories))]
	return nil
}

// SetClothesStyle sets a random clothes style for the character.
func (rcb *randomCharacterBuilder) SetClothesStyle() error {
	// Generate random clothes style (for demonstration purposes, just using a fixed set of styles)
	styles := []string{"Casual", "Formal", "Sporty"}
	rcb.character.ClothesStyle = styles[rand.Intn(len(styles))]
	return nil
}

// Build returns the built character.
func (rcb *randomCharacterBuilder) Build() entities.Character {
	return rcb.character
}