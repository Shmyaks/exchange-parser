// Package handlers handlers of some functions or methods
package handlers

import (
	customErrors "github.com/Shmyaks/exchange-parser-server/app/internal/models/custom_errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// ValidateBody validate body from request
func ValidateBody(scheme interface{}, c *fiber.Ctx) error {
	if err := c.BodyParser(scheme); err != nil {
		return customErrors.ValidationError

	}

	err := validate.Struct(scheme)
	if err != nil {
		// for _, err := range err.(validator.ValidationErrors) {
		// 	var element customErrors.ValidationError
		// 	element.FailedField = err.StructNamespace()
		// 	element.Tag = err.Tag()
		// 	element.Value = err.Param()
		// 	errors = append(errors, &element)
		// }

		return customErrors.ValidationError
	}

	return nil
}

// ValidateQuery validate query params from request
// Example: url/api/path?{name}=SOMETEXT
func ValidateQuery(scheme interface{}, c *fiber.Ctx) error {
	if err := c.QueryParser(scheme); err != nil {
		return customErrors.ValidationError

	}

	err := validate.Struct(scheme)
	if err != nil {
		// for _, err := range err.(validator.ValidationErrors) {
		// 	var element customErrors.ValidationError
		// 	element.FailedField = err.StructNamespace()
		// 	element.Tag = err.Tag()
		// 	element.Value = err.Param()
		// 	errors = append(errors, &element)
		// }

		return customErrors.ValidationError
	}

	return nil
}

// ValidateParam validate params from request
// Example: url/api/path/{name}
func ValidateParam(scheme interface{}, c *fiber.Ctx) error {
	if err := c.ParamsParser(scheme); err != nil {
		return customErrors.ValidationError

	}
	err := validate.Struct(scheme)
	if err != nil {
		// for _, err := range err.(validator.ValidationErrors) {
		// 	var element customErrors.ValidationError
		// 	element.FailedField = err.StructNamespace()
		// 	element.Tag = err.Tag()
		// 	element.Value = err.Param()
		// 	errors = append(errors, &element)
		// }

		return customErrors.ValidationError
	}
	return nil
}
