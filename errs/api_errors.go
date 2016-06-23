package errs

import "fmt"

// APIError defines a standard format for API errors.
type APIError struct {
	// The status code.
	Status int `json:"status"`
	// The description of the API error.
	Description string `json:"description"`
	// The token uniquely identifying the API error.
	ErrorCode string `json:"errorCode"`
	// Additional infos.
	Params map[string]interface{} `json:"params,omitempty"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("%s : %s", e.ErrorCode, e.Description)
}

var (
	APIInternal = APIError{
		Description: "An internal error occured. Please retry later.",
		ErrorCode:   "INTERNAL_ERROR",
	}
	APIJsonRendering = APIError{
		Description: "The JSON rendering failed.",
		ErrorCode:   "JSON_RENDERING_ERROR",
	}
	APIBodyDecoding = APIError{
		Description: "Could not decode the JSON request.",
		ErrorCode:   "BODY_DECODING_ERROR",
	}
	APIUnauthorized = APIError{
		Description: "Authorization Required.",
		ErrorCode:   "AUTHORIZATION_REQUIRED",
	}
	APIForbidden = APIError{
		Description: "The specified resource was not found or you don't have sufficient permissions.",
		ErrorCode:   "FORBIDDEN",
	}
	APIFilterDecoding = APIError{
		Description: "Could not decode the given filter.",
		ErrorCode:   "FILTER_DECODING_ERROR",
	}
	APIValidation = APIError{
		Description: "The model validation failed.",
		ErrorCode:   "VALIDATION_ERROR",
	}
)
