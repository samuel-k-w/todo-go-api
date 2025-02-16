package response

import "github.com/gofiber/fiber/v2"

// APIResponse defines a consistent structure for responses.
type APIResponse struct {
	Status  string      `json:"status" `         // "success" or "error"
	Message string      `json:"message" `        // Human-readable message
	Data    interface{} `json:"data,omitempty" ` // Actual data (if any)
}

// SuccessResponse returns a success response with data.
func SuccessResponse(message string, data interface{}) fiber.Map {
	return fiber.Map{
		"status":  "success",
		"message": message,
		"data":    data,
	}
}

// ErrorResponse returns an error response.
func ErrorResponse(message string) fiber.Map {
	return fiber.Map{
		"status":  "error",
		"message": message,
	}
}
