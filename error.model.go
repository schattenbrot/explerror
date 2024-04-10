package explerror

// Error represents the structure for the error response.
// swagger:model
type Error struct {
	// StatusCode represents the HTTP status code of the error.
	// example: 400
	StatusCode int `json:"statusCode"`

	// Message represents the error message.
	// example: "descriptive error message"
	Message string `json:"message"`
}
