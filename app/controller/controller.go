package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mmgopher/user-service/app/api/request"
	"github.com/mmgopher/user-service/app/api/response"
	"github.com/mmgopher/user-service/app/httperrors"
	"github.com/mmgopher/user-service/app/middleware"
	"github.com/mmgopher/user-service/app/service/user"
)

// Controller represents Controller layer of application.
type Controller struct {
	userService user.Provider
}

// New creates new instance of Controller.
func New(
	userService user.Provider,
) *Controller {
	return &Controller{
		userService: userService,
	}
}

// GetUser handles GET /v1/users/:user_id endpoint
func (c Controller) GetUser(context *gin.Context) {
	u, err := c.userService.GetUser(context.GetInt(middleware.UserIDParamKey))
	if err != nil {
		httperrors.Emit(context, err)
		return
	}
	context.JSON(http.StatusOK, response.User{
		ID:        u.ID,
		Name:      u.Name,
		Surname:   u.Surname,
		Gender:    u.Gender,
		Age:       u.Age,
		Address:   u.Address,
		CreatedAt: u.CreatedAt,
	})
}

// DeleteUser handles DELETE /v1/users/:user_id endpoint
func (c Controller) DeleteUser(context *gin.Context) {
	if err := c.userService.DeleteUser(context.GetInt(middleware.UserIDParamKey)); err != nil {
		httperrors.Emit(context, err)
		return
	}
	context.JSON(http.StatusOK, gin.H{})
}

// CreateUser handles POST /v1/users endpoint
func (c Controller) CreateUser(context *gin.Context) {

	var req request.CreateUser
	if err := context.ShouldBindJSON(&req); err != nil {
		httperrors.Emit(context, httperrors.RequestBodyParsingError.WithCause(err))
		return
	}

	userID, err := c.userService.CreateUser(&req)

	if err != nil {
		httperrors.Emit(context, err)
		return
	}

	context.JSON(http.StatusCreated, response.CreateUser{
		ID: userID,
	})
}

// UpdateUser handles PUT /v1/users/:user_id endpoint
func (c Controller) UpdateUser(context *gin.Context) {

	var req request.UpdateUser
	if err := context.ShouldBindJSON(&req); err != nil {
		httperrors.Emit(context, httperrors.RequestBodyParsingError.WithCause(err))
		return
	}

	if err := c.userService.UpdateUser(
		context.GetInt(middleware.UserIDParamKey),
		&req,
	); err != nil {
		httperrors.Emit(context, err)
		return
	}
	context.JSON(http.StatusOK, gin.H{})
}

// GetUserList handles GET /v1/users endpoint.
func (c Controller) GetUserList(context *gin.Context) {
	var req request.FindUsers
	if err := context.ShouldBindQuery(&req); err != nil {
		httperrors.Emit(context, httperrors.QueryParametersParsingError.WithCause(err))
		return
	}

	users, beforeID, afterID, err := c.userService.FindUsers(&req)
	if err != nil {
		httperrors.Emit(context, err)
		return
	}

	prevURL, nextURL := getPaginationURLs(context.Request.URL, beforeID, afterID)
	userListResponse := make([]response.User, 0, len(users))

	for _, u := range users {
		userListResponse = append(userListResponse, response.User{
			ID:        u.ID,
			Name:      u.Name,
			Surname:   u.Surname,
			Age:       u.Age,
			Gender:    u.Gender,
			Address:   u.Address,
			CreatedAt: u.CreatedAt,
		})

	}

	context.JSON(http.StatusOK, response.UserListWithPagination{
		Result: userListResponse,
		Pagination: response.Pagination{
			PrevLink: prevURL,
			BeforeID: beforeID,
			NextLink: nextURL,
			AfterID:  afterID,
		},
	})

}
