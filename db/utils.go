package db

import "fmt"

func FindRecipes(ingredients []string, stopList []string) ([]Recipe, error) {
	var recipes []Recipe
	fmt.Println(ingredients, stopList)
	err := DB.Order("rating").Find(&recipes).Error
	//err := DB.Order("rating").Not(map[string]interface{}{"title": stopList}).Find(&recipes, ingredients).Error
	fmt.Println(recipes)
	return recipes, err
}

func CreateRecipesIngredients(ingredients []string) ([]Ingredient, error) {
	var (
		ingredientObject Ingredient
		result           []Ingredient
		err              error
	)
	for _, ingredient := range ingredients {
		if err = DB.Where("title = ?", ingredient).First(&ingredientObject).Error; err != nil {
		} else {
			result = append(result, ingredientObject)
		}
	}
	return result, nil
}
