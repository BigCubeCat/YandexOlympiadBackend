package db

type Recipe struct {
	RecipeId    uint `gorm:"primary_key:true"`
	Title       string
	Description string
	Rating      uint
	Link        string
	Ingredients []Ingredient `gorm:"many2many:recipe_ingredients"`
}
type Ingredient struct {
	Title string `gorm:"primary_key:true"`
}
