package http

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ilyakaznacheev/cleanenv"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"

	"github.com/bariasabda/kumparan/config"
	"github.com/bariasabda/kumparan/domain/repository"
	"github.com/bariasabda/kumparan/domain/service"
)

func Execute() {
	// Configuration
	var cfg config.Config

	err := cleanenv.ReadConfig("./config/config.yml", &cfg)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// connect postgresql
	db, err := gorm.Open(postgres.Open(cfg.Postgresql.Url), &gorm.Config{})
	if err != nil {
		log.Fatalf("Tidak Konek DB Errornya : %s", err)
	}

	repository := repository.NewRepository(db)
	service := service.NewService(cfg, repository)

	// Run
	handler := NewHandler(service)
	r := NewRouter(*handler)
	r.routes()
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = cfg.Base.Port
	}
	r.router.Run(port)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Fatalf("app - Run - signal: %s", s.String())
	default:
	}
}
