package character

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func (h handler) CreateCharacter(c *gin.Context) {
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