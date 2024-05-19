package character

import (
	"github.com/juanmiguelar/character_generator/internal/core/character/providers"
	"github.com/juanmiguelar/character_generator/internal/core/entities"
	"golang.org/x/sync/errgroup"
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
    g := new(errgroup.Group)
    g.Go(d.character_builder.SetAge)
    g.Go(d.character_builder.SetClothesStyle)
    g.Go(d.character_builder.SetName)
    g.Go(d.character_builder.SetSex)
    if err := g.Wait(); err != nil {
        return character, err
    }
	if err := d.character_builder.SetBackdropStory(); err != nil {
		return character, err
	}
	return d.character_builder.Build(), nil
}
