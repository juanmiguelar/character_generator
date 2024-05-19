package mocks

import (
	"github.com/juanmiguelar/character_generator/internal/core/entities"
	"github.com/stretchr/testify/mock"
)

// MockCharacterBuilder is a mock implementation of the CharacterBuilderProvider interface
type MockCharacterBuilder struct {
	mock.Mock
}

func (m *MockCharacterBuilder) SetName() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockCharacterBuilder) SetAge() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockCharacterBuilder) SetSex() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockCharacterBuilder) SetBackdropStory() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockCharacterBuilder) SetClothesStyle() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockCharacterBuilder) Build() entities.Character {
	args := m.Called()
	return args.Get(0).(entities.Character)
}