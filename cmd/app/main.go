package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Quszlet/subscription-service/internal/repository"
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
	);

	_, err = repository.NewPostgresDB(config)
	if err != nil {
		slog.Error(err.Error())
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hell12")
	})

	slog.Info("Listening on :8030")

	http.ListenAndServe(":8030", nil)
}
