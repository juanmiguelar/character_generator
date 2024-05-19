package character

import "github.com/juanmiguelar/character_generator/entrypoints/handlers/character/providers"

type HandlerDependencies struct {
	CharacterBuilder providers.CharacterBuilderDirector
}

type handler struct {
	characterBuilder providers.CharacterBuilderDirector
}

func newHandler(params HandlerDependencies) handler {
	return handler{
		characterBuilder: params.CharacterBuilder,
	}
}
