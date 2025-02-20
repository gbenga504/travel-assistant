package main

import (
	"log/slog"

	"github.com/gbenga504/travel-assistant/lib"
	"github.com/gbenga504/travel-assistant/utils/logger"
)

func main() {
	slog.SetDefault(logger.NewLogger())

	lib.RunApp()
}
