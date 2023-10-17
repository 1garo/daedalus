package main

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type (
	ErrorResponse struct {
		Tag         string `json:"description"     validate:"required"`
		FailedField string `json:"field,omitempty"`
	}

	XValidator struct {
		validator *validator.Validate
	}
)

// This is the validator instance
// for more information see: https://github.com/go-playground/validator
var validate = validator.New()

func (v XValidator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = strings.ToLower(err.Field()) // Export struct field name

			elem.Tag = err.Tag() // Export struct tag

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

type CreateRecipeInput struct {
	ID          int                `json:"id"          validate:"required"`
	Ingredients []RecipeIngredient `json:"ingredients" validate:"required"`
	Steps       []Step             `json:"steps"       validate:"required"`
}

type RecipeIngredient struct {
	Amount int `json:"amount"`
	// e.g: cup, tea scoop, oz
	Measure string `json:"measure"` // TODO: this should be a enum, let's just validate for now,
	Name    string `json:"name"`
}

type Step struct {
	Description string `json:"description"`
	Order       int    `json:"order"`
}

func main() {
	myValidator := &XValidator{
		validator: validate,
	}

	app := fiber.New()
	// Initialize default config
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"data": "HealthCheck!"})
	})

	app.Post("/recipe", func(c *fiber.Ctx) error {
		input := new(CreateRecipeInput)

		if err := c.BodyParser(input); err != nil {
			errs := []ErrorResponse{
				{
					Tag: err.Error(),
				},
			}

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errs})
		}

		if errs := myValidator.Validate(input); len(errs) > 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errs})
		}

		return c.JSON(fiber.Map{"data": input})
	})

	app.Get("/recipe", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"data": "Alexandre"})
	})

	app.Listen(":3000")
}
