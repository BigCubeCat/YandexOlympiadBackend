package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"search/db"
)

func CreateRecipe(c *fiber.Ctx) error {
	c.Accepts("application/json") // "application/json"
	var err error
	request := new(RecipeJSON)
	if err = c.BodyParser(request); err != nil {
		c.Status(http.StatusBadRequest)
		return err
	}
	fmt.Println(request)
	recipe := db.Recipe{
		Title:       request.Title,
		Description: request.Description,
		Link:        request.Link,
		Ingredients: []*db.Ingredient{},
	}
	for _, ingr := range request.Ingredients {
		recipe.Ingredients = append(recipe.Ingredients, &db.Ingredient{Title: ingr})
	}
	fmt.Println(recipe)
	if err = db.DB.Create(&recipe).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return err
	}
	c.Status(http.StatusOK)
	return nil
}
