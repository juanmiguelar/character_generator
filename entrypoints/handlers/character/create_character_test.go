package character

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/juanmiguelar/character_generator/entrypoints/handlers/character/providers/mocks"
	"github.com/juanmiguelar/character_generator/internal/core/entities"
	"github.com/stretchr/testify/assert"
)

// TestCreateCharacter tests the CreateCharacter handler
func TestCreateCharacter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		// Arrange
		mockBuilder := new(mocks.MockCharacterDirector)
		expectedCharacter := entities.Character{
			Name:         "Test",
			Age:          30,
			Sex:          "Male",
			BackdropStory: "Test Story",
			ClothesStyle: "Casual",
		}
		mockBuilder.On("Build").Return(expectedCharacter, nil)

		h := handler{characterBuilder: mockBuilder}

		// Create a gin router and register the handler
		r := gin.Default()
		r.POST("/characters", h.CreateCharacter)

		// Create a request to pass to the handler
		req, _ := http.NewRequest(http.MethodPost, "/characters", nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"name":"Test","age":30,"sex":"Male","backdrop_story":"Test Story","clothes_style":"Casual"}`, w.Body.String())
		mockBuilder.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		// Arrange
		mockBuilder := new(mocks.MockCharacterDirector)
		mockBuilder.On("Build").Return(entities.Character{}, errors.New("build error"))

		h := handler{characterBuilder: mockBuilder}

		// Create a gin router and register the handler
		r := gin.Default()
		r.POST("/characters", h.CreateCharacter)

		// Create a request to pass to the handler
		req, _ := http.NewRequest(http.MethodPost, "/characters", nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"error":"build error"}`, w.Body.String())
		mockBuilder.AssertExpectations(t)
	})
}
