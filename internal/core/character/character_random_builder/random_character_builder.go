package character_random_builder

import (
	"errors"
	"math/rand"

	"github.com/juanmiguelar/character_generator/internal/core/entities"
)

// randomCharacterBuilder implements the CharacterBuilder interface.
type randomCharacterBuilder struct {
	character entities.Character
	rand      *rand.Rand
}

// NewRandomCharacterBuilder creates a new randomCharacterBuilder with a custom random generator.
func NewRandomCharacterBuilder(r *rand.Rand) *randomCharacterBuilder {
	return &randomCharacterBuilder{
		character: entities.Character{},
		rand:      r,
	}
}

// SetName sets a random name for the character.
func (rcb *randomCharacterBuilder) SetName() error {
	names := []string{"Alice", "Bob", "Charlie", "Diana", "Emma"}
	if len(names) == 0 {
		return errors.New("no names available")
	}
	rcb.character.Name = names[rcb.rand.Intn(len(names))]
	return nil
}

// SetAge sets a random age for the character.
func (rcb *randomCharacterBuilder) SetAge() error {
	ageRange := 43
	if ageRange <= 0 {
		return errors.New("invalid age range")
	}
	rcb.character.Age = rcb.rand.Intn(ageRange) + 18
	return nil
}

// SetSex sets a random sex for the character.
func (rcb *randomCharacterBuilder) SetSex() error {
	sexes := []string{"Male", "Female"}
	if len(sexes) == 0 {
		return errors.New("no sexes available")
	}
	rcb.character.Sex = sexes[rcb.rand.Intn(len(sexes))]
	return nil
}

// SetBackdropStory sets a random backdrop story for the character.
func (rcb *randomCharacterBuilder) SetBackdropStory() error {
	stories := []string{"Once upon a time...", "In a land far away...", "In a world of magic..."}
	if len(stories) == 0 {
		return errors.New("no backdrop stories available")
	}
	rcb.character.BackdropStory = stories[rcb.rand.Intn(len(stories))]
	return nil
}

// SetClothesStyle sets a random clothes style for the character.
func (rcb *randomCharacterBuilder) SetClothesStyle() error {
	styles := []string{"Casual", "Formal", "Sporty"}
	if len(styles) == 0 {
		return errors.New("no clothes styles available")
	}
	rcb.character.ClothesStyle = styles[rcb.rand.Intn(len(styles))]
	return nil
}

// Build returns the built character.
func (rcb *randomCharacterBuilder) Build() entities.Character {
	return rcb.character
}
