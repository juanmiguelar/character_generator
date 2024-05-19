package character

import "github.com/juanmiguelar/character_generator/entrypoints/handlers/character/providers"

type HandlerDependencies struct {
	CharacterBuilder providers.CharacterDirectorProvider
}

type handler struct {
	characterBuilder providers.CharacterDirectorProvider
}

func newHandler(params HandlerDependencies) handler {
	return handler{
		characterBuilder: params.CharacterBuilder,
	}
}
