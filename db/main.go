package db

import (
	"fmt"
	"golang.org/x/exp/slices"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

var DB *gorm.DB

func InitDB(dbName string) {
	var err error
	DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = DB.AutoMigrate(&Recipe{}, &Ingredient{}, &IngredientsSet{})
	if err != nil {
		log.Fatalln("Cant Auto-migrate")
		return
	}
}

func FindRecipes(goodIngredients []string, badIngredients []string) (Recipe, error) {
	var (
		recipes []Recipe
		sets    []IngredientsSet
		err     error
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
		return Recipe{}, err
	}
	var possibleRecipes []uint
	for key, _ := range dict {
		possibleRecipes = append(possibleRecipes, key)
	}
	// В теории, ID сета и рецепта равны. TODO: fix this shit
	err = DB.Preload("IngSet").Preload("IngSet.Ingredients").
		Where("recipe_id IN ?", possibleRecipes).Order("rating").
		Find(&recipes).Error
	if rand.Intn(2) == 1 {
		return recipes[0], err
	}
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(recipes))
	return recipes[index], err
}
