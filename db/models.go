package db

type Recipe struct {
	RecipeId    int `gorm:"primaryKey"`
	Title       string
	Description string
	Rating      int
	Link        string
	Ingredients []*Ingredient `gorm:"many2many:ingredients"`
}
type Ingredient struct {
	Title   string    `gorm:"primaryKey"`
	Recipes []*Recipe `gorm:"many2many:ingredients"`
}
