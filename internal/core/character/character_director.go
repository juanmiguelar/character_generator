package character

import (
	"github.com/juanmiguelar/character_generator/internal/core/character/providers"
	"github.com/juanmiguelar/character_generator/internal/core/entities"
)

type characterDirector struct {
	character_builder providers.CharacterBuilderProvider
}

func NewDirector() (*characterDirector) {
	default_builder, err := GetCharacterBuilder(RANDOM)
	if err != nil {
		panic(err)
	}
	return &characterDirector{
		character_builder: default_builder,
	}
}

func (d *characterDirector) SetBuilder(b providers.CharacterBuilderProvider) {
	d.character_builder = b
}

func (d *characterDirector) Build() (character entities.Character, err error) {
	if err := d.character_builder.SetAge(); err != nil {
        return character, err
    }
    if err := d.character_builder.SetBackdropStory(); err != nil {
        return character, err
    }
    if err := d.character_builder.SetClothesStyle(); err != nil {
        return character, err
    }
    if err := d.character_builder.SetName(); err != nil {
        return character, err
    }
    if err := d.character_builder.SetSex(); err != nil {
        return character, err
    }
	return d.character_builder.Build(), nil
}
