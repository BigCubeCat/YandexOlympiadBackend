package db

func FindRecipes(ingredients []string, stopList []string) ([]Recipe, error) {
	var recipes []Recipe
	err := db.Order("rating").Not(map[string]interface{}{"title": stopList}).Find(&recipes, ingredients).Error
	return recipes, err
}

func CreateRecipesIngredients(ingredients []string) ([]Ingredient, error) {
	var (
		ingredientObject Ingredient
		result           []Ingredient
		err              error
	)
	for _, ingredient := range ingredients {
		if err = db.Where("title = ?", ingredient).First(&ingredientObject).Error; err != nil {
		} else {
			result = append(result, ingredientObject)
		}
	}
	return result, nil
}
