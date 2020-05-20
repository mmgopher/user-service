package main

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/mmgopher/user-service/app"
	"github.com/mmgopher/user-service/app/config"
	"github.com/mmgopher/user-service/app/controller"
	"github.com/mmgopher/user-service/app/dao"
	"github.com/mmgopher/user-service/app/db"
	"github.com/mmgopher/user-service/app/service/user"
)

func main() {

	cfg, err := config.New()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatal("impossible to initialize config")
	}

	logLevel, err := log.ParseLevel(strings.ToLower(cfg.LogLevel))
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatal("unsupported log level")
	}

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logLevel)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)

	postgresConnection, err := db.GetConnection(
		cfg.DBType,
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBName,
		cfg.DBHost,
	)

	if err != nil {
		log.Fatalf("impossible to establish connection to %s database: %+v", cfg.DBType, err)
	}

	userService := user.NewService(
		dao.NewUserRepository(postgresConnection),
	)

	router := app.NewRouter(cfg, controller.New(
		userService,
	))
	router.Run()
}
