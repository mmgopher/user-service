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
	"github.com/mmgopher/user-service/app/api/response"
	"github.com/mmgopher/user-service/test/data"
	"github.com/mmgopher/user-service/test/helpers"
)

// TestUpdateUserOK makes test of PUT /v1/users/:user_id
func TestUpdateUserOK(t *testing.T) {
	httpService := helpers.NewHTTPService(http.DefaultClient)
	request, err := json.Marshal(data.UpdateUserRequest)
	assert.Nil(t, err)
	statusCode, respBody, err := httpService.DoRequest(
		http.MethodPut,
		helpers.StrReplace(
			os.Getenv("APP_BASE_URL")+app.RootPath+app.UpdateUserRoute,
			":user_id",
			3,
		),
		nil,
		nil,
		request,
	)
	require.Nil(t, err)

	assert.Equal(t, http.StatusOK, statusCode)

	// Get update user to check if updates were applied
	statusCode, respBody, err = httpService.DoRequest(
		http.MethodGet,
		helpers.StrReplace(
			os.Getenv("APP_BASE_URL")+app.RootPath+app.GetUserRoute,
			":user_id",
			3,
		),
		nil,
		nil,
		nil,
	)
	require.Nil(t, err)
	var user response.User
	assert.Nil(t, json.Unmarshal(respBody, &user))
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, data.UpdateUserRequest.Name, user.Name)
	assert.Equal(t, data.UpdateUserRequest.Surname, user.Surname)
	assert.Equal(t, data.UpdateUserRequest.Age, user.Age)
	assert.Equal(t, data.UpdateUserRequest.Gender, user.Gender)
}
