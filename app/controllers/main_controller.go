package app

import (
	"github.com/1garo/daedalus/types"
	"github.com/gofiber/fiber/v2"
)

type CreateRecipeInput struct {
	Ingredients []RecipeIngredient `json:"ingredients" validate:"required,dive"`
	Steps       []Step             `json:"steps"       validate:"required,dive"`
}

type RecipeIngredient struct {
	Amount int `json:"amount" validate:"required"`
	// e.g: cup, tea scoop, oz
	Measure string `json:"measure" validate:"required"` // TODO: this should be a enum, let's just validate for now,
	Name    string `json:"name" validate:"required"`
}

type Step struct {
	Description string `json:"description"`
	Order       int    `json:"order"`
}

// CreateRecipe Add a recipe into the database
func CreateRecipe(c *fiber.Ctx) error {
	input := new(CreateRecipeInput)

	if err := c.BodyParser(input); err != nil {
		errs := []types.ErrorResponse{
			{
				Tag: err.Error(),
			},
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errs})
	}

	myValidator := c.Locals("validator").(*types.XValidator)

	if errs := myValidator.Validate(input); len(errs) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errs})
	}

	return c.JSON(fiber.Map{"data": input})
}
