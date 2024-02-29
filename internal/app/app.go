package app

import (
	"app/internal/interface/server/app"
	"app/internal/registry"
	"log"
)

func Run(config *Config) error {
	cnt, err := NewContainer(config)
	if err != nil {
		return err
	}

	appServer := InitAppServer(cnt)

	log.Println("port", config.AppServer.Port)
	if err := appServer.Run(config.AppServer.Port); err != nil {
		return err
	}

	return nil
}

func InitAppServer(cnt *registry.Container) *app.Server {
	s := app.NewAppServer()
	s.InitRoutes(cnt)

	return s
}
