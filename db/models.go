package db

type Recipe struct {
	RecipeId    uint `gorm:"primaryKey"`
	Title       string
	Description string
	Rating      uint
	Link        string
	Ingredients []*Ingredient `gorm:"many2many:ingredients"`
}
type Ingredient struct {
	Title   string    `gorm:"primaryKey"`
	Recipes []*Recipe `gorm:"many2many:ingredients"`
}
