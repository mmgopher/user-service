package app

import (
	"github.com/gin-gonic/gin"

	"github.com/mmgopher/user-service/app/config"
	"github.com/mmgopher/user-service/app/controller"
)

// NewRouter initializes the gin router and routes.
func NewRouter(config *config.Config, controller *controller.Controller) *gin.Engine {

	g := gin.Default()

	return g
}
