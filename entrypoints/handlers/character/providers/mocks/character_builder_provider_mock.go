package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/juanmiguelar/character_generator/internal/core/entities"
	"github.com/juanmiguelar/character_generator/internal/core/character/providers"
)

// MockCharacterDirector is a mock implementation of the CharacterDirectorProvider interface
type MockCharacterDirector struct {
	mock.Mock
}

func (m *MockCharacterDirector) Build() (entities.Character, error) {
	args := m.Called()
	return args.Get(0).(entities.Character), args.Error(1)
}

func (m *MockCharacterDirector) SetBuilder(b providers.CharacterBuilderProvider) {
	m.Called(b)
}
