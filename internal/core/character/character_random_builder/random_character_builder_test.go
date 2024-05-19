package character_random_builder

import (
	"math/rand"
	"testing"
	"time"

	"github.com/juanmiguelar/character_generator/internal/core/entities"
	"github.com/stretchr/testify/assert"
)

func newRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func TestNewRandomCharacterBuilder(t *testing.T) {
	r := newRand()
	builder := NewRandomCharacterBuilder(r)
	assert.NotNil(t, builder)
	assert.Equal(t, entities.Character{}, builder.character)
}

func TestSetName(t *testing.T) {
	r := newRand()
	builder := NewRandomCharacterBuilder(r)
	err := builder.SetName()
	assert.NoError(t, err)
	assert.Contains(t, []string{"Alice", "Bob", "Charlie", "Diana", "Emma"}, builder.character.Name)
}

func TestSetAge(t *testing.T) {
	r := newRand()
	builder := NewRandomCharacterBuilder(r)
	err := builder.SetAge()
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, builder.character.Age, 18)
	assert.LessOrEqual(t, builder.character.Age, 60)
}

func TestSetSex(t *testing.T) {
	r := newRand()
	builder := NewRandomCharacterBuilder(r)
	err := builder.SetSex()
	assert.NoError(t, err)
	assert.Contains(t, []string{"Male", "Female"}, builder.character.Sex)
}

func TestSetBackdropStory(t *testing.T) {
	r := newRand()
	builder := NewRandomCharacterBuilder(r)
	err := builder.SetBackdropStory()
	assert.NoError(t, err)
	assert.Contains(t, []string{"Once upon a time...", "In a land far away...", "In a world of magic..."}, builder.character.BackdropStory)
}

func TestSetClothesStyle(t *testing.T) {
	r := newRand()
	builder := NewRandomCharacterBuilder(r)
	err := builder.SetClothesStyle()
	assert.NoError(t, err)
	assert.Contains(t, []string{"Casual", "Formal", "Sporty"}, builder.character.ClothesStyle)
}

func TestBuild(t *testing.T) {
	r := newRand()
	builder := NewRandomCharacterBuilder(r)
	err := builder.SetName()
	assert.NoError(t, err)
	err = builder.SetAge()
	assert.NoError(t, err)
	err = builder.SetSex()
	assert.NoError(t, err)
	err = builder.SetBackdropStory()
	assert.NoError(t, err)
	err = builder.SetClothesStyle()
	assert.NoError(t, err)
	character := builder.Build()
	assert.NotNil(t, character)
	assert.NotEmpty(t, character.Name)
	assert.GreaterOrEqual(t, character.Age, 18)
	assert.LessOrEqual(t, character.Age, 60)
	assert.Contains(t, []string{"Male", "Female"}, character.Sex)
	assert.Contains(t, []string{"Once upon a time...", "In a land far away...", "In a world of magic..."}, character.BackdropStory)
	assert.Contains(t, []string{"Casual", "Formal", "Sporty"}, character.ClothesStyle)
}
