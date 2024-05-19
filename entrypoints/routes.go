package entrypoints

import (
	"log"

	"time"

	"github.com/gin-gonic/gin"
	character_routes "github.com/juanmiguelar/character_generator/entrypoints/handlers/character"
	"github.com/juanmiguelar/character_generator/internal/core/character"
)

func CreateRoutes() *gin.Engine {
	r := gin.New()
	r.Use(Logger())
	character_routes.CharacterRouter(r, character_routes.HandlerDependencies{
		CharacterBuilder: character.NewDirector(),
	})

	return r
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		// after request
		latency := time.Since(t)
		log.Printf("Latency: %s", latency.String())

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
