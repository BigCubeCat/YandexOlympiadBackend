package api

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"search/db"
)

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
	recipes, er := db.FindRecipes(request.GoodIngredients, request.BadIngredients)
	if er != nil {
		c.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message":  "Not found",
			"response": recipes,
		})
		return er
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message":  "OK",
		"response": recipes,
	})
	return nil
}
