package db

type Recipe struct {
	RecipeId    uint `gorm:"primary_key:true"`
	Title       string
	Description string
	CountRates  int
	Rating      float32
	Link        string
	IngSet      IngredientsSet `gorm:"foreignKey:set_id"`
}
type IngredientsSet struct {
	SetId       uint         `gorm:"primary_key:true"`
	Ingredients []Ingredient `gorm:"many2many:recipe_ingredients"`
}
type Ingredient struct {
	Title string `gorm:"primary_key:true"`
}
