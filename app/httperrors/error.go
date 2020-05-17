package httperrors

import (
	"fmt"
)

// Common Application errors.
// Format AAXXXBB
// AA - endpoint
// XXX - http error code
// BB - custom error code
var (
	InternalServerError = NewHTTPInternalServerError(
		1050000, "internal server error",
	)
	RequestBodyParsingError = NewBadRequest(
		1040000, "could not parse the request body",
	)
	QueryParametersParsingError = NewBadRequest(
		1040001, "could not parse the query parameters",
	)

	PathParametersParsingError = NewBadRequest(
		1040002, "could not parse the path parameters",
	)

	EntityNotFoundError = func(entity string) *HTTPError {
		return NewNotFound(1040400, fmt.Sprintf(
			"`%s` entity not found", entity),
		)
	}
)
