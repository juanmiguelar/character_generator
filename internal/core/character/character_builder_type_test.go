package character

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetCharacterBuilder(t *testing.T) {
	// Test with RANDOM builder type
	t.Run("RandomBuilder", func(t *testing.T) {
		builder, err := GetCharacterBuilder(RANDOM)
		require.NoError(t, err)
		assert.NotNil(t, builder)
	})

	// Test with unsupported builder type
	t.Run("UnsupportedBuilder", func(t *testing.T) {
		builder, err := GetCharacterBuilder(CharacterBuilderType("UNSUPPORTED"))
		require.Error(t, err)
		assert.Nil(t, builder)
		assert.Contains(t, err.Error(), "builder type is not supported")
	})
}
