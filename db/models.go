package db

type Recipe struct {
	RecipeId    uint `gorm:"primaryKey"`
	Title       string
	Description string
	Rating      uint
	Link        string
	Ingredients []Ingredient `gorm:"many2many:recipe_ingredients;foreignKey:IngredientId"`
}
type Ingredient struct {
	IngredientID uint `gorm:"primaryKey"`
	Title        string
}
