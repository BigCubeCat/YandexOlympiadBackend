package db

import (
	"golang.org/x/exp/slices"
)

func FindRecipes(goodIngredients []string, badIngredients []string) (Recipe, error) {
	var (
		recipe Recipe
		sets   []IngredientsSet
		err    error
	)
	dict := make(map[uint]int)
	err = DB.Preload("Ingredients").Find(&sets).Error
	for _, set := range sets {
		count := 0
		good := true
		for _, ingr := range set.Ingredients {
			// Check, what set not contain ingredient from bad list
			if slices.Contains(badIngredients, ingr.Title) {
				good = false
				break
			}
			if slices.Contains(goodIngredients, ingr.Title) {
				count++
			}
		}
		if count > 0 && good {
			dict[set.SetId] = count
		}
	}
	if err != nil {
		return recipe, err
	}
	maximum := 0
	bestSet := uint(0)
	for key, count := range dict {
		if count > maximum {
			maximum = count
			bestSet = key
		}
	}
	// В теории, ID сета и рецепта равны. TODO: fix this shit
	err = DB.Preload("IngSet").Preload("IngSet.Ingredients").
		Where(bestSet).
		First(&recipe).Error
	return recipe, err
}
