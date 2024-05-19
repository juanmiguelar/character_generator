package entities

// Character represents a random character.
type Character struct {
	Name          string `json:"name"`
	Age           int    `json:"age"`
	Sex           string `json:"sex"`
	BackdropStory string `json:"backdrop_story"`
	ClothesStyle  string `json:"clothes_style"`
}