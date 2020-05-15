package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/mmgopher/user-service/app"
	"github.com/mmgopher/user-service/app/config"
	"github.com/mmgopher/user-service/app/controller"
)

func main() {

	appCfg, err := config.New()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatal("impossible to initialize config")
	}

	router := app.NewRouter(appCfg, controller.New())
	router.Run()
}
