package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"search/db"
	"strconv"
)

func CreateRecipe(c *fiber.Ctx) error {
	c.Accepts("application/json") // "application/json"
	var err error
	request := new(RecipeJSON)
	if err = c.BodyParser(&request); err != nil {
		c.Status(http.StatusBadRequest)
		return err
	}
	fmt.Println(request)
	recipe := db.Recipe{
		Title:       request.Title,
		Description: request.Description,
		Link:        request.Link,
		IngSet:      db.IngredientsSet{},
	}
	for _, ingr := range request.Ingredients {
		recipe.IngSet.Ingredients = append(recipe.IngSet.Ingredients, db.Ingredient{Title: ingr})
	}
	if err = db.DB.Create(&recipe).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return err
	}
	c.Status(http.StatusOK)
	return nil
}

func AddRate(c *fiber.Ctx) error {
	c.Accepts("application/json")
	id := c.Params("id")
	rate := c.Params("rate")
	rateNum, err := strconv.Atoi(rate)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return err
	}
	if rateNum < 0 {
		rateNum = 1
	} else if rateNum > 5 {
		rateNum = 5
	}
	var recipe db.Recipe
	if err := db.DB.Where("recipe_id = ?", id).First(&recipe).Error; err != nil {
		return err
	}
	newRating := (recipe.Rating*float32(recipe.CountRates) + float32(rateNum)) / float32(recipe.CountRates+1)
	c.Status(http.StatusOK)
	err = db.DB.Model(&recipe).Updates(
		db.Recipe{
			Rating:     newRating,
			CountRates: recipe.CountRates + 1,
		},
	).Error
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return err
	}
	c.Status(http.StatusOK)
	return nil
}
