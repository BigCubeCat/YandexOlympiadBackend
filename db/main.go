package db

import (
	"errors"
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
		log.Fatalln("Cant open DB")
		return
	}
	err = DB.AutoMigrate(&Recipe{}, &Ingredient{}, &IngredientsSet{})
	if err != nil {
		log.Fatalln("Cant Auto-migrate")
		return
	}
}

func GetSession() *gorm.DB {
	return DB.Session(&gorm.Session{})
}

func FindRecipes(goodIngredients []string, badIngredients []string) (Recipe, error) {
	var (
		recipes []Recipe
		sets    []IngredientsSet
		err     error
	)
	dict := make(map[uint]int)
	err = GetSession().Preload("Ingredients").Find(&sets).Error
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
	maximumSuggest := 0
	for _, value := range dict {
		if maximumSuggest < value {
			maximumSuggest = value
		}
	}
	for key, count := range dict {
		if count == maximumSuggest {
			possibleRecipes = append(possibleRecipes, key)
		}
	}
	// В теории, ID сета и рецепта равны. TODO: fix this shit
	err = GetSession().Preload("IngSet").Preload("IngSet.Ingredients").
		Where("recipe_id IN ?", possibleRecipes).Order("rating").
		Find(&recipes).Error
	if len(recipes) == 0 {
		return Recipe{}, errors.New("not found")
	}
	if rand.Intn(2) == 1 {
		return recipes[0], err
	}
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(recipes))
	return recipes[index], err
}

func FindByTitle(title string, bad []string) (Recipe, error) {
	var recipes []Recipe
	var goodRecipes []Recipe
	var calm bool
	err := GetSession().Preload("IngSet").Preload("IngSet.Ingredients").
		Where("Title LIKE lower(?)", title+"%").Order("rating").Find(&recipes).Error
	if err != nil {
		return Recipe{}, err
	}
	for _, recipe := range recipes {
		// Check, that recipe not contains "bad" ingredients
		calm = true
		for _, ingr := range recipe.IngSet.Ingredients {
			if slices.Contains(bad, ingr.Title) {
				calm = false
				break
			}
		}
		if calm {
			goodRecipes = append(goodRecipes, recipe)
		}
	}
	if len(recipes) == 0 {
		return Recipe{}, errors.New("not found")
	}
	if rand.Intn(2) == 1 {
		return recipes[0], err
	}
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(recipes))
	return recipes[index], err
}

func FindById(id string) (Recipe, error) {
	var recipes []Recipe
	err := GetSession().Preload("IngSet").Preload("IngSet.Ingredients").
		Where("recipe_id LIKE ?", id).Order("rating").Find(&recipes).Error
	if err != nil {
		return Recipe{}, err
	}
	if len(recipes) == 0 {
		return Recipe{}, errors.New("not found")
	}
	return recipes[0], err
}
