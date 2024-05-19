package character

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juanmiguelar/character_generator/internal/core/character"
)

// createCharacterParams defines the structure for the query parameters
type createCharacterParams struct {
	GeneratorType character.CharacterBuilderType `form:"generatorType"`
}

func (h handler) CreateCharacter(c *gin.Context) {
	var params createCharacterParams

	// Bind the query parameters to the struct
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(params.GeneratorType) > 0 {
		builder, err := character.GetCharacterBuilder(params.GeneratorType)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		h.characterBuilder.SetBuilder(builder)
	}

	// Build the character using the builder pattern
	character, err := h.characterBuilder.Build()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, character)
}
