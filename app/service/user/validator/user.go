package validator

import (
	"regexp"
	"strings"

	"github.com/mmgopher/user-service/app/api/request"
	"github.com/mmgopher/user-service/app/httperrors"
)

var sortRegex = regexp.MustCompile("^[a-zA-Z_]*:(asc|desc)$")

var supportedGenderList = map[string]struct{}{
	"male":           {},
	"female":         {},
	"trans":          {},
	"intersex":       {},
	"non-conforming": {},
	"personal":       {},
}

// User validators
var (
	// validateName validates `name` request parameter.
	validateName = func(name string) error {
		if name == "" {
			return httperrors.UserNameEmpty
		}

		return nil
	}

	// validateSurname validates `surname` request parameter.
	validateSurname = func(surname string) error {
		if surname == "" {
			return httperrors.UserSurnameEmpty
		}

		return nil
	}

	// validateAge validates `age` request parameter.
	validateAge = func(age int) error {
		if age < 1 || age > 120 {
			return httperrors.UserAgeIncorrect
		}

		return nil
	}

	// validateGender validates `gender` request parameter.
	validateGender = func(gender string) error {
		if gender == "" {
			return httperrors.UserGenderEmpty
		}

		if _, ok := supportedGenderList[strings.ToLower(gender)]; !ok {
			return httperrors.UserGenderNotSupported(gender)
		}

		return nil
	}

	validateAddress = func(address string) error {
		if address == "" {
			return httperrors.UserAddressEmpty
		}

		return nil
	}
)

// ValidateCreateUserRequest validates POST /v1/users endpoint.
func ValidateCreateUserRequest(request *request.CreateUser) error {

	if err := validateName(request.Name); err != nil {
		return err
	}

	if err := validateSurname(request.Surname); err != nil {
		return err
	}

	if err := validateAge(request.Age); err != nil {
		return err
	}

	if err := validateGender(request.Gender); err != nil {
		return err
	}

	if err := validateAddress(request.Address); err != nil {
		return err
	}

	return nil
}

// ValidateUpdateUserRequest validates PUT /v1/users/:user_id endpoint.
func ValidateUpdateUserRequest(request *request.UpdateUser) error {

	if err := validateName(request.Name); err != nil {
		return err
	}

	if err := validateSurname(request.Surname); err != nil {
		return err
	}

	if err := validateAge(request.Age); err != nil {
		return err
	}

	if err := validateGender(request.Gender); err != nil {
		return err
	}

	if err := validateAddress(request.Address); err != nil {
		return err
	}

	return nil
}

// ValidateFindUsersRequest validates GET /v1/users endpoint.
func ValidateFindUsersRequest(request *request.FindUsers) error {

	if request.AfterID > 0 && request.BeforeID > 0 {
		return httperrors.PaginationAfterIDAndBeforeIDDeclared
	} else if request.AfterID < 0 {
		return httperrors.PaginationAfterIDNegative
	} else if request.BeforeID < 0 {
		return httperrors.PaginationBeforeIDNegative
	}

	if request.Limit < 0 {
		return httperrors.PaginationLimitNegative
	}

	if len(request.Sort) > 0 && !sortRegex.MatchString(request.Sort) {
		return httperrors.PaginationSortIncorrectFormat
	}

	return nil
}
