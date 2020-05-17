package app

import (
	"github.com/gin-gonic/gin"

	"github.com/mmgopher/user-service/app/config"
	"github.com/mmgopher/user-service/app/controller"
	"github.com/mmgopher/user-service/app/middleware"
)

// RootPath - current api version
const RootPath = "/v1"

// supported route creators.
const (
	GetUserRoute    = "/users/:user_id"
	UpdateUserRoute = "/users/:user_id"
	DeleteUserRoute = "/users/:user_id"
	CreateUserRoute = "/users"
)

// NewRouter initializes the gin router and routes.
// config is currently not used but can be used to configure CORS
func NewRouter(config *config.Config, controller *controller.Controller) *gin.Engine {

	g := gin.Default()
	v1 := g.Group(RootPath)
	{
		v1.GET(GetUserRoute, middleware.ValidateUserID, controller.GetUser)
		v1.POST(CreateUserRoute, controller.CreateUser)
		v1.PUT(UpdateUserRoute, middleware.ValidateUserID, controller.UpdateUser)
		v1.DELETE(DeleteUserRoute, middleware.ValidateUserID, controller.DeleteUser)
	}
	return g
}
