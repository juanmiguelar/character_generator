package character

import (
	"errors"
	"testing"

	"github.com/juanmiguelar/character_generator/internal/core/character/providers/mocks"
	"github.com/juanmiguelar/character_generator/internal/core/entities"
	"github.com/stretchr/testify/assert"
)

func TestNewDirector(t *testing.T) {
	director := NewDirector()
	assert.NotNil(t, director)
	assert.NotNil(t, director.character_builder)
}

func TestSetBuilder(t *testing.T) {
	director := NewDirector()
	mockBuilder := new(mocks.MockCharacterBuilder)

	director.SetBuilder(mockBuilder)
	assert.Equal(t, mockBuilder, director.character_builder)
}

func TestBuild(t *testing.T) {
	mockBuilder := new(mocks.MockCharacterBuilder)
	director := &characterDirector{
		character_builder: mockBuilder,
	}

	mockBuilder.On("SetAge").Return(nil)
	mockBuilder.On("SetBackdropStory").Return(nil)
	mockBuilder.On("SetClothesStyle").Return(nil)
	mockBuilder.On("SetName").Return(nil)
	mockBuilder.On("SetSex").Return(nil)
	expectedCharacter := entities.Character{
		Name:         "Test",
		Age:          30,
		Sex:          "Male",
		BackdropStory: "Test Story",
		ClothesStyle: "Casual",
	}
	mockBuilder.On("Build").Return(expectedCharacter)

	character, err := director.Build()
	assert.NoError(t, err)
	assert.Equal(t, expectedCharacter, character)

	mockBuilder.AssertExpectations(t)
}

func TestBuildWithError(t *testing.T) {
	mockBuilder := new(mocks.MockCharacterBuilder)
	director := &characterDirector{
		character_builder: mockBuilder,
	}

	mockBuilder.On("SetAge").Return(errors.New("error setting age"))

	character, err := director.Build()
	assert.Error(t, err)
	assert.Equal(t, entities.Character{}, character)
}