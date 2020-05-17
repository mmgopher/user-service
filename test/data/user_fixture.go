package data

import (
	"time"

	"github.com/mmgopher/user-service/app/api/response"
)

// NotExistingUserID represents ID of user not existing in DB
var NotExistingUserID = 50001

var (
	// GetUserSuccessResponse is a success response for GET /v1/users/1 endpoint.
	GetUserSuccessResponse = response.User{
		ID:        1,
		Name:      "Sonny",
		Surname:   "Watts",
		Gender:    "male",
		Age:       30,
		Address:   "1754  Arron Smith Drive",
		CreatedAt: time.Date(2020, 04, 25, 16, 47, 59, 0, time.UTC),
	}
)
