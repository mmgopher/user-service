package validator

import (
	"testing"

	"github.com/mmgopher/user-service/app/api/request"
)

func TestValidateCreateUserRequestError(t *testing.T) {

	var testData = []struct {
		testName      string
		request       request.CreateUser
		expectedError error
	}{}

}
