// +build integration

package integration

import (
	"encoding/json"
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

// TestGetUserOK makes test of GET /v1/users/:user_id
func TestGetUserOK(t *testing.T) {
	httpService := helpers.NewHTTPService(http.DefaultClient)
	statusCode, respBody, err := httpService.DoRequest(
		http.MethodGet,
		helpers.StrReplace(
			os.Getenv("APP_BASE_URL")+app.RootPath+app.GetUserRoute,
			":user_id",
			1,
		),
		nil,
		nil,
		nil,
	)
	require.Nil(t, err)
	expectedJSONResponse, err := json.Marshal(data.GetUserSuccessResponse)
	require.Nil(t, err)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.JSONEq(t, string(expectedJSONResponse), string(respBody))
}

// TestGetUserError makes test of GET /v1/users/:user_id.
func TestGetUserError(t *testing.T) {
	httpService := helpers.NewHTTPService(http.DefaultClient)
	statusCode, respBody, err := httpService.DoRequest(
		http.MethodGet,
		helpers.StrReplace(
			os.Getenv("APP_BASE_URL")+app.RootPath+app.GetUserRoute,
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
