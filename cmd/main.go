package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Angstreminus/cinema/config"
	"github.com/Angstreminus/cinema/internal/postgres"
	"github.com/Angstreminus/cinema/logger"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	zaplog := logger.MustInitLogger()
	fmt.Println("Logger initialized")
	db, err := postgres.NewDatabaseHandler(cfg)
	if err != nil {
		zaplog.ZapLogger.Error("Error co connect postgres")
	}

	err = db.PingContext(context.Background())
	if err != nil {
		zaplog.ZapLogger.Fatal("Fail to connect")
		log.Fatal(err)
	}
	zaplog.ZapLogger.Info("Connected and ping success")
}
