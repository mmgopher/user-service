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

// Application errors for `POST /v1/users` and `PUT /v1/users/:user_id
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

// Application errors for GET /v1/users endpoint
var (
	PaginationAfterIDNegative = NewBadRequest(
		2140001, "`afterID` can not be negative",
	)

	PaginationBeforeIDNegative = NewBadRequest(
		2140002, "`beforeID` can not be negative",
	)

	PaginationAfterIDAndBeforeIDDeclared = NewBadRequest(
		2140003, "`afterID` and beforeID can not both been declared",
	)

	PaginationLimitNegative = NewBadRequest(
		2140004, "`limit` can not be negative",
	)

	PaginationSortIncorrectFormat = NewBadRequest(
		2140005, "`sort` parameter does not match sort pattern",
	)
)
