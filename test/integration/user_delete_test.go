// +build integration

package integration

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mmgopher/user-service/app"
	"github.com/mmgopher/user-service/app/httperrors"
	"github.com/mmgopher/user-service/test/data"
	"github.com/mmgopher/user-service/test/helpers"
)

// TestGetUserOK makes test of DELETE /v1/users/:user_id.
func TestDeleteUserOK(t *testing.T) {

	userID := 2
	httpService := helpers.NewHTTPService(http.DefaultClient)
	// User should exists
	statusCode, _, err := httpService.DoRequest(
		http.MethodGet,
		helpers.StrReplace(
			os.Getenv("APP_BASE_URL")+app.RootPath+app.GetUserRoute,
			":user_id",
			userID,
		),
		nil,
		nil,
		nil,
	)
	require.Nil(t, err)
	assert.Equal(t, http.StatusOK, statusCode)

	// Remove user
	statusCode, _, err = httpService.DoRequest(
		http.MethodDelete,
		helpers.StrReplace(
			os.Getenv("APP_BASE_URL")+app.RootPath+app.DeleteUserRoute,
			":user_id",
			userID,
		),
		nil,
		nil,
		nil,
	)
	require.Nil(t, err)
	assert.Equal(t, http.StatusOK, statusCode)

	// User should not exists
	statusCode, _, err = httpService.DoRequest(
		http.MethodGet,
		helpers.StrReplace(
			os.Getenv("APP_BASE_URL")+app.RootPath+app.GetUserRoute,
			":user_id",
			userID,
		),
		nil,
		nil,
		nil,
	)
	require.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, statusCode)
}

// TestDeleteUserError makes test of DELETE /v1/users/:user_id.
func TestDeleteUserError(t *testing.T) {
	httpService := helpers.NewHTTPService(http.DefaultClient)
	statusCode, respBody, err := httpService.DoRequest(
		http.MethodDelete,
		helpers.StrReplace(
			os.Getenv("APP_BASE_URL")+app.RootPath+app.DeleteUserRoute,
			":user_id",
			data.NotExistingUserID,
		),
		nil,
		nil,
		nil,
	)
	expectedError := httperrors.EntityNotFoundError("user")
	require.Nil(t, err)
	assert.Equal(t, expectedError.HTTPCode, statusCode)
	assert.JSONEq(t, expectedError.Error(), string(respBody))
}
