package main

import (
	"fmt"
	"net/http"

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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hell12")
	})

	slog.Info("Listening on :8030")

	http.ListenAndServe(":8030", nil)
}
