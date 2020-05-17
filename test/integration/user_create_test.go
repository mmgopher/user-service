// build integration

package integration

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mmgopher/user-service/app"
	"github.com/mmgopher/user-service/app/api/request"
	"github.com/mmgopher/user-service/app/api/response"
	"github.com/mmgopher/user-service/app/httperrors"
	"github.com/mmgopher/user-service/test/data"
	"github.com/mmgopher/user-service/test/helpers"
)

// TestCreateUserOK makes test of POST /v1/users
func TestCreateUserOK(t *testing.T) {
	httpService := helpers.NewHTTPService(http.DefaultClient)
	request, err := json.Marshal(data.CreateUserRequest)
	assert.Nil(t, err)
	statusCode, respBody, err := httpService.DoRequest(
		http.MethodPost,
		os.Getenv("APP_BASE_URL")+app.RootPath+app.CreateUserRoute,
		nil,
		nil,
		request,
	)
	require.Nil(t, err)

	var target response.CreateUser
	assert.Nil(t, json.Unmarshal(respBody, &target))
	assert.Equal(t, http.StatusCreated, statusCode)
	assert.True(t, target.ID > 0)

	// Get new user to check if exist
	statusCode, respBody, err = httpService.DoRequest(
		http.MethodGet,
		helpers.StrReplace(
			os.Getenv("APP_BASE_URL")+app.RootPath+app.GetUserRoute,
			":user_id",
			target.ID,
		),
		nil,
		nil,
		nil,
	)
	require.Nil(t, err)
	var user response.User
	assert.Nil(t, json.Unmarshal(respBody, &user))
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, data.CreateUserRequest.Name, user.Name)
	assert.Equal(t, data.CreateUserRequest.Surname, user.Surname)
	assert.Equal(t, data.CreateUserRequest.Age, user.Age)
	assert.Equal(t, data.CreateUserRequest.Gender, user.Gender)
}

func TestCreateUserError(t *testing.T) {

	var testData = []struct {
		testName      string
		request       request.CreateUser
		expectedError *httperrors.HTTPError
	}{
		{
			testName: "EmptyName",
			request: request.CreateUser{
				Surname: "test",
				Age:     23,
				Gender:  "male",
			},
			expectedError: httperrors.UserNameEmpty,
		},
		{
			testName: "EmptySurname",
			request: request.CreateUser{
				Name:   "test",
				Age:    23,
				Gender: "male",
			},
			expectedError: httperrors.UserSurnameEmpty,
		},
		{
			testName: "IncorrectAge",
			request: request.CreateUser{
				Name:    "test",
				Surname: "test",
				Age:     -1,
				Gender:  "male",
			},
			expectedError: httperrors.UserAgeIncorrect,
		},
		{
			testName: "NotSupportedGender",
			request: request.CreateUser{
				Name:    "test",
				Surname: "test",
				Age:     23,
				Gender:  "other",
			},
			expectedError: httperrors.UserGenderNotSupported("other"),
		},
		{
			testName: "AlreadyRegistered",
			request: request.CreateUser{
				Name:    "Sonny",
				Surname: "Watts",
				Age:     23,
				Gender:  "male",
				Address: "address",
			},
			expectedError: httperrors.UserAlreadyRegistered,
		},
	}

	httpService := helpers.NewHTTPService(http.DefaultClient)
	for _, tt := range testData {
		t.Run(tt.testName, func(t *testing.T) {
			request, err := json.Marshal(tt.request)
			assert.Nil(t, err)
			statusCode, respBody, err := httpService.DoRequest(
				http.MethodPost,
				os.Getenv("APP_BASE_URL")+app.RootPath+app.CreateUserRoute,
				nil,
				nil,
				request,
			)
			require.Nil(t, err)
			assert.Equal(t, tt.expectedError.HTTPCode, statusCode)
			assert.JSONEq(t, tt.expectedError.Error(), string(respBody))
		})
	}
}
