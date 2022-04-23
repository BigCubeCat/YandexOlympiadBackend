package db

type Recipe struct {
	RecipeId    uint `gorm:"primary_key:true"`
	Title       string
	Description string
	Link        string
	Energy      string
	Steps       string
	CountRates  int
	Rating      float32
	IngSet      IngredientsSet `gorm:"foreignKey:set_id"`
}
type IngredientsSet struct {
	SetId       uint         `gorm:"primary_key:true"`
	Ingredients []Ingredient `gorm:"many2many:recipe_ingredients"`
	Counts      string       // just JSON with data. like {"молоко": "3мл"}
}
type Ingredient struct {
	Title string `gorm:"primary_key:true"`
}
