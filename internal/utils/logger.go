package utils

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Quszlet/subscription-service/pkg/logger"
)

func CreateFolderForLogs() error {
	nameDir := "logs"
	_, err := os.Stat(nameDir)
	if os.IsNotExist(err) {
		if err := os.Mkdir("logs", os.ModePerm); err != nil {
			return err
		}

	}

	return nil
}

func CreateFileForLog() (*os.File, error) {
	logFile, err := os.OpenFile("logs/subscription-service.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return logFile, nil
}

func SetupLogsFiles() (*os.File, error) {
	if err := CreateFolderForLogs(); err != nil {
		return nil, err
	}

	logfile, err := CreateFileForLog()
	if err != nil {
		return nil, err
	}

	fmt.Println("Logs file create with directory!")

	return logfile, nil
}

func CreateSlogLogger(file *os.File) *slog.Logger {
	consoleHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	fileHandler := slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	multiHandler := logger.NewMultiHandler(consoleHandler, fileHandler)
	log := slog.New(multiHandler)

	return log
}
