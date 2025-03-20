package lib

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gbenga504/travel-assistant/lib/routes"
	util "github.com/gbenga504/travel-assistant/utils"
	"github.com/gbenga504/travel-assistant/utils/agent/llms/gemini"
	"github.com/gbenga504/travel-assistant/utils/db"
	"github.com/gbenga504/travel-assistant/utils/db/mongodb"
	"github.com/gbenga504/travel-assistant/utils/errors"
	"github.com/gbenga504/travel-assistant/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		logger.Fatal("Cannot load .env files", logger.ErrorOpt{
			Name:          errors.Name(errors.ErrEnvNotLoaded),
			Message:       errors.Message(errors.ErrEnvNotLoaded),
			OriginalError: err.Error(),
		})
	}
}

func shutdownSrvGracefully(server *http.Server) {
	// Create a channel to listen for signals
	quit := make(chan os.Signal, 1)

	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Wait till we receive a signal
	<-quit
	logger.Info("Shutting down server")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// When shutdown is initiated, our server stops receiving connections,
	// try to finish up with ongoing connections and we gracefully shutdown the server after the timeout set in ctx
	// it also blocks until all is done
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown", logger.ErrorOpt{
			Name:          errors.Name(errors.ErrServerClosed),
			Message:       errors.Message(errors.ErrServerClosed),
			OriginalError: err.Error(),
		})
	}

	logger.Info("Server exiting")
}

func createDatabase() db.Db {
	db := mongodb.NewMongoDB()

	return db
}

func startServer() {
	GEMINI_API_KEY := util.LookupEnv("GEMINI_API_KEY")
	geminiClient := gemini.NewClient(context.Background(), GEMINI_API_KEY)
	defer geminiClient.Close()

	db := createDatabase()
	defer db.Close()

	httpHandler := gin.New()
	routes.Routes(httpHandler, geminiClient)

	PORT := fmt.Sprintf(":%s", util.LookupEnv("PORT"))
	logger.Info(fmt.Sprintf("Listening on Port %s", PORT))

	server := &http.Server{
		Addr:    PORT,
		Handler: httpHandler,
	}

	// We run to run this in goroutine because ListenAndServe blocks and we need to
	// run some other code e.g taking care of graceful shutdowns
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Error occurred causing server shutdown", logger.ErrorOpt{
				Name:          errors.Name(errors.ErrServerClosed),
				Message:       errors.Message(errors.ErrServerClosed),
				OriginalError: err.Error(),
			})
		}
	}()

	shutdownSrvGracefully(server)
}

func RunApp() {
	loadEnv()
	startServer()
}
