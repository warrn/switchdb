package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"switchdb/internal/config"
	"switchdb/internal/router"

	"github.com/oklog/run"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	var publicListener, adminListener net.Listener
	if cfg.Production {
		log.Fatal("production unimplemented")
	} else {
		var err error
		publicListener, err = net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.IP, cfg.PublicPort))
		if err != nil {
			log.Fatal(err)
		}

		adminListener, err = net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.IP, cfg.AdminPort))
		if err != nil {
			log.Fatal(err)
		}
	}

	publicRouter := router.NewPublicRouter()
	adminRouter := router.NewAdminRouter()

	// TODO: run group is insufficient, we should hook into interrupt signals and use a
	//       shared context to coordinate graceful shutdown
	rg := run.Group{}

	rg.Add(func() error {
		err := http.Serve(publicListener, publicRouter)
		return err
	}, func(err error) {
		publicListener.Close()
	})

	rg.Add(func() error {
		err := http.Serve(adminListener, adminRouter)
		return err
	}, func(err error) {
		adminListener.Close()
	})

	rg.Run()
}
