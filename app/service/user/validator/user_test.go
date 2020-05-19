// +build unit

package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mmgopher/user-service/app/api/request"
	"github.com/mmgopher/user-service/app/httperrors"
)

func TestValidateCreateUserRequestOK(t *testing.T) {

	err := ValidateCreateUserRequest(&request.CreateUser{
		Name:    "name",
		Gender:  "male",
		Surname: "surname",
		Age:     10,
		Address: "address",
	})
	assert.Nil(t, err)
}
func TestValidateCreateUserRequestError(t *testing.T) {

	var testData = []struct {
		name          string
		request       *request.CreateUser
		expectedError error
	}{
		{
			"EmptyName",
			&request.CreateUser{
				Surname: "surname",
				Gender:  "male",
				Age:     30,
				Address: "address",
			},
			httperrors.UserNameEmpty,
		},
		{
			"EmptySurname",
			&request.CreateUser{
				Name:    "name",
				Gender:  "male",
				Age:     30,
				Address: "address",
			},
			httperrors.UserSurnameEmpty,
		},
		{
			"EmptyGender",
			&request.CreateUser{
				Name:    "name",
				Surname: "surname",
				Gender:  "",
				Age:     30,
				Address: "address",
			},
			httperrors.UserGenderEmpty,
		},
		{
			"NotSupportedGender",
			&request.CreateUser{
				Name:    "name",
				Surname: "surname",
				Gender:  "gender",
				Age:     30,
				Address: "address",
			},
			httperrors.UserGenderNotSupported("gender"),
		},
		{
			"IncorrectAge",
			&request.CreateUser{
				Name:    "name",
				Gender:  "male",
				Surname: "surname",
				Age:     0,
				Address: "address",
			},
			httperrors.UserAgeIncorrect,
		},
		{
			"EmptyAddress",
			&request.CreateUser{
				Name:    "name",
				Gender:  "male",
				Surname: "surname",
				Age:     10,
			},
			httperrors.UserAddressEmpty,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCreateUserRequest(tt.request)
			require.NotNil(t, err)
			assert.EqualError(t, tt.expectedError, err.Error())
		})
	}
}

func TestValidateUpdateUserRequestOK(t *testing.T) {

	err := ValidateUpdateUserRequest(&request.UpdateUser{
		Name:    "name",
		Gender:  "male",
		Surname: "surname",
		Age:     10,
		Address: "address",
	})
	assert.Nil(t, err)
}

func TestValidateUpdateUserRequestError(t *testing.T) {

	var testData = []struct {
		name          string
		request       *request.UpdateUser
		expectedError error
	}{
		{
			"EmptyName",
			&request.UpdateUser{
				Surname: "surname",
				Gender:  "male",
				Age:     30,
				Address: "address",
			},
			httperrors.UserNameEmpty,
		},
		{
			"EmptySurname",
			&request.UpdateUser{
				Name:    "name",
				Gender:  "male",
				Age:     30,
				Address: "address",
			},
			httperrors.UserSurnameEmpty,
		},
		{
			"EmptyGender",
			&request.UpdateUser{
				Name:    "name",
				Surname: "surname",
				Gender:  "",
				Age:     30,
				Address: "address",
			},
			httperrors.UserGenderEmpty,
		},
		{
			"NotSupportedGender",
			&request.UpdateUser{
				Name:    "name",
				Surname: "surname",
				Gender:  "gender",
				Age:     30,
				Address: "address",
			},
			httperrors.UserGenderNotSupported("gender"),
		},
		{
			"IncorrectAge",
			&request.UpdateUser{
				Name:    "name",
				Gender:  "male",
				Surname: "surname",
				Age:     0,
				Address: "address",
			},
			httperrors.UserAgeIncorrect,
		},
		{
			"EmptyAddress",
			&request.UpdateUser{
				Name:    "name",
				Gender:  "male",
				Surname: "surname",
				Age:     10,
			},
			httperrors.UserAddressEmpty,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUpdateUserRequest(tt.request)
			require.NotNil(t, err)
			assert.EqualError(t, tt.expectedError, err.Error())
		})
	}
}

func TestValidateFindeUserRequestError(t *testing.T) {

	var testData = []struct {
		name          string
		request       *request.FindUsers
		expectedError error
	}{
		{
			"AfterIDAndBeforeIDeclared",
			&request.FindUsers{
				AfterID:  4,
				BeforeID: 5,
			},
			httperrors.PaginationAfterIDAndBeforeIDDeclared,
		},
		{
			"AfterIDNegative",
			&request.FindUsers{
				AfterID:  -1,
				BeforeID: 0,
			},
			httperrors.PaginationAfterIDNegative,
		},
		{
			"BeforeIDNegative",
			&request.FindUsers{
				AfterID:  0,
				BeforeID: -5,
			},
			httperrors.PaginationBeforeIDNegative,
		},
		{
			"LimitNegative",
			&request.FindUsers{
				AfterID: 0,
				Limit:   -1,
			},
			httperrors.PaginationLimitNegative,
		},
		{
			"SortColumnMissingOrder",
			&request.FindUsers{
				AfterID: 0,
				Limit:   10,
				Sort:    "id",
			},
			httperrors.PaginationSortIncorrectFormat,
		},
		{
			"SortColumnWrongOrder",
			&request.FindUsers{
				AfterID: 0,
				Limit:   10,
				Sort:    "id:down",
			},
			httperrors.PaginationSortIncorrectFormat,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateFindUsersRequest(tt.request)
			require.NotNil(t, err)
			assert.EqualError(t, tt.expectedError, err.Error())
		})
	}
}
