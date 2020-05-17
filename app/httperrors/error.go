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

// Application error for `POST /userss` and `PUT /users/:user_id
var (
	UserNameEmpty = NewBadRequest(
		2040001, "`name` can't be empty",
	)

	UserSurnameEmpty = NewBadRequest(
		2040002, "`surnname` can't be empty",
	)

	UserAgeIncorrect = NewBadRequest(
		2040003, "`age` should be `>0` and `<120`",
	)

	UserGenderEmpty = NewBadRequest(
		2040004, "`gender` can't be empty",
	)

	UserGenderNotSupported = func(gender string) *HTTPError {
		return NewBadRequest(2040005, fmt.Sprintf(
			"`gender` %s is not supported", gender),
		)
	}

	UserAddressEmpty = NewBadRequest(
		2040006, "`address` can't be empty",
	)

	UserAlreadyRegistered = NewBadRequest(
		2040007, "user with provided name and surname already exists",
	)
)
