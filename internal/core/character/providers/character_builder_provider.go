package providers

import "github.com/juanmiguelar/character_generator/internal/core/entities"

type CharacterBuilderProvider interface {
	SetName() error
	SetAge() error
	SetSex() error
	SetBackdropStory() error
	SetClothesStyle() error
	Build() entities.Character
}