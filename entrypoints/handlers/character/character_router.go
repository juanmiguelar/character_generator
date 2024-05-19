package character

import "github.com/gin-gonic/gin"

func CharacterRouter(r *gin.Engine, params HandlerDependencies){
	h := newHandler(params)
	character := r.Group("/character")
	{
		character.POST("create", h.CreateCharacter)
	}
}