package api

import (
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
	recipe := db.Recipe{
		Title:       request.Title,
		Description: request.Description,
		Link:        request.Link,
		Energy:      request.Energy,
		Steps:       request.Steps,
		IngSet:      db.IngredientsSet{},
	}
	for _, ingr := range request.Ingredients {
		recipe.IngSet.Ingredients = append(recipe.IngSet.Ingredients, db.Ingredient{Title: ingr})
	}
	recipe.IngSet.Counts = request.IngredientsSetCounts
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

func GetRecipes(c *fiber.Ctx) error {
	c.Accepts("application/json") // "application/json"
	request := new(RequestJSON)
	if err := c.BodyParser(request); err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message":  "error",
			"response": []RecipeJSON{},
		})
		return err
	}
	recipe, er := db.FindRecipes(request.GoodIngredients, request.BadIngredients)
	if er != nil {
		c.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message":  "Not found",
			"response": recipe,
		})
		return er
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message":  "OK",
		"response": recipe,
	})
	return nil
}

func GetByTitle(c *fiber.Ctx) error {
	c.Accepts("application/json") // "application/json"
	request := new(RecipeTitleJSON)
	if err := c.BodyParser(request); err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message":  "error",
			"response": []RecipeJSON{},
		})
		return err
	}
	recipe, er := db.FindByTitle(request.Title, request.BadIngredients)
	if er != nil {
		c.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message":  "Not found",
			"response": recipe,
		})
		return er
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message":  "OK",
		"response": recipe,
	})
	return nil
}
