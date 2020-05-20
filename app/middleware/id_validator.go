package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/mmgopher/user-service/app/httperrors"
)

// Request placeholders keys
const (
	UserIDParamKey = "user_id"
)

// ValidateUserID validates :user_id placeholder from the request.
func ValidateUserID(context *gin.Context) {
	validateURLParamAsNumber(context, UserIDParamKey)
}

func validateURLParamAsNumber(context *gin.Context, paramName string) {

	value, err := strconv.Atoi(context.Param(paramName))
	if err != nil {
		httperrors.Emit(
			context, httperrors.PathParametersParsingError.WithCause(
				errors.Wrapf(err, "wrong format of :%s", paramName),
			),
		)
		context.Abort()
	} else {
		context.Set(paramName, value)
		context.Next()
	}
}
