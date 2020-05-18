// +build unit

package middleware

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mmgopher/user-service/app/httperrors"
	"github.com/stretchr/testify/assert"
)

func TestValidateCreatorID_OK(t *testing.T) {

	expectedUserID := 1
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)

	param := gin.Param{
		Key:   UserIDParamKey,
		Value: fmt.Sprintf("%d", expectedUserID),
	}

	context.Params = append(context.Params, param)
	ValidateUserID(context)

	respBody, err := ioutil.ReadAll(recorder.Result().Body)
	assert.Nil(t, err)

	assert.Equal(t, expectedUserID, context.GetInt(UserIDParamKey))
	assert.Empty(t, string(respBody))
}

func TestValidateCreatorID_Error(t *testing.T) {

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)

	param := gin.Param{
		Key:   UserIDParamKey,
		Value: "user",
	}
	context.Params = append(context.Params, param)
	ValidateUserID(context)

	respBody, err := ioutil.ReadAll(recorder.Result().Body)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.JSONEq(t, httperrors.PathParametersParsingError.Error(), string(respBody))
}
