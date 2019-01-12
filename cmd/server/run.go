package main

import (
	"os"

	"github.com/izumin5210/ghrsync/app"
	"go.uber.org/zap"
)

func main() {
	os.Exit(run())
}

func run() int {
	err := app.Run()
	if err != nil {
		zap.L().Error("server was shutdown with errors", zap.Error(err))
		return 1
	}
	return 0
}
