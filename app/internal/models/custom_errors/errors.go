// Package customErrors ...
package customErrors

import "github.com/gofiber/fiber/v2"

// EmptyError, ValidationError Custom errors
var (
	EmptyError      = fiber.NewError(404, "Empty error")
	ValidationError = fiber.NewError(422, "Validation error")
)
