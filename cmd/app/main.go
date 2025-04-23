package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Quszlet/subscription-service/internal/api/handler"
	"github.com/Quszlet/subscription-service/internal/repository"
	"github.com/Quszlet/subscription-service/internal/services"
	"github.com/Quszlet/subscription-service/internal/utils"
	"github.com/Quszlet/subscription-service/pkg/logger"
)

func main() {
	logFile, err := utils.SetupLogsFiles()
	if err != nil {
		panic(err)
	}

	slogLogger := utils.CreateSlogLogger(logFile)

	slog := logger.NewSlogAdapter(slogLogger)

	config := repository.NewConfigPostgres(
		os.Getenv("HOST_DB"),
		os.Getenv("PORT_DB"),
		os.Getenv("USER_DB"),
		os.Getenv("PASSWORD_DB"),
		os.Getenv("NAME_DB"),
		os.Getenv("SSL_MODE"),
	)

	db, err := repository.NewPostgresDB(config)
	if err != nil {
		slog.Error(err.Error())
	}

	repos := repository.NewRepository(db, slog)
	service := service.NewService(repos, slog)
	handler := handler.NewHandler(service, slog)

	routes := handler.InitRoutes()


	slog.Info(fmt.Sprintf("Listening on %s", os.Getenv("PORT_SRV")))

	srv := &http.Server{
		Addr:           ":" + os.Getenv("PORT_SRV"),
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	slog.Error(srv.ListenAndServe().Error())
}
