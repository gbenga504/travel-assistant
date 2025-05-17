package lib

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gbenga504/travel-assistant/lib/airport"
	airportcontroller "github.com/gbenga504/travel-assistant/lib/airport/controller"
	airportrepository "github.com/gbenga504/travel-assistant/lib/airport/repository"
	airportservice "github.com/gbenga504/travel-assistant/lib/airport/service"
	"github.com/gbenga504/travel-assistant/lib/health"
	healthcontroller "github.com/gbenga504/travel-assistant/lib/health/controller"
	"github.com/gbenga504/travel-assistant/lib/middlewares"
	"github.com/gbenga504/travel-assistant/lib/thread"
	threadcontroller "github.com/gbenga504/travel-assistant/lib/thread/controller"
	threadrepository "github.com/gbenga504/travel-assistant/lib/thread/repository"
	threadservice "github.com/gbenga504/travel-assistant/lib/thread/service"
	"github.com/gbenga504/travel-assistant/utils"
	"github.com/gbenga504/travel-assistant/utils/agent/llms/gemini"
	"github.com/gbenga504/travel-assistant/utils/db"
	"github.com/gbenga504/travel-assistant/utils/db/mongodb"
	"github.com/gbenga504/travel-assistant/utils/errors"
	llmcontext "github.com/gbenga504/travel-assistant/utils/llm_context"
	"github.com/gbenga504/travel-assistant/utils/logger"
	"github.com/gin-gonic/gin"
)

type Server struct {
	addr         string
	geminiClient *gemini.GeminiClient
	db           db.Db
	httpServer   *http.Server
}

func NewServer(addr string) *Server {
	GEMINI_API_KEY := utils.LookupEnv("GEMINI_API_KEY")
	geminiClient := gemini.NewClient(context.Background(), GEMINI_API_KEY)

	httpHandler := gin.New()

	DATABASE_NAME := utils.LookupEnv("DATABASE_NAME")
	db := mongodb.Connect(DATABASE_NAME)

	// Apply global middlewares
	httpHandler.Use(middlewares.CORSMiddleware())

	v1 := httpHandler.Group("/api/v1")

	// Health
	healthController := healthcontroller.NewHealthController()
	health.ConnectRoutes(v1, healthController)

	// Airport
	airportRepository := airportrepository.NewAirportRepository(db)
	airportService := airportservice.NewAirportService(airportRepository)
	airportController := airportcontroller.NewAirportController(airportService)
	airport.ConnectRoutes(v1, airportController)

	// Thread
	threadRepository := threadrepository.NewThreadRepository(db)
	threadService := threadservice.NewThreadService(threadRepository, geminiClient, &llmcontext.LLMContext{
		AirportRepository: airportRepository,
	})
	threadController := threadcontroller.NewThreadController(threadService)
	thread.ConnectRoutes(v1, threadController)

	return &Server{
		addr:         addr,
		geminiClient: geminiClient,
		db:           db,
		httpServer:   &http.Server{Addr: addr, Handler: httpHandler},
	}
}

func (s *Server) Run() {
	// We run to run this in goroutine because ListenAndServe blocks and we need to
	// run some other code e.g taking care of graceful shutdowns
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Error occurred causing server shutdown", logger.ErrorOpt{
				Name:          errors.Name(errors.ErrServerClosed),
				Message:       errors.Message(errors.ErrServerClosed),
				OriginalError: err.Error(),
			})
		}
	}()
}

func (s *Server) ShutdownGracefully() {
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
	defer s.db.Close()
	defer s.geminiClient.Close()

	// When shutdown is initiated, our server stops receiving connections,
	// try to finish up with ongoing connections and we gracefully shutdown the server after the timeout set in ctx
	// it also blocks until all is done
	if err := s.httpServer.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown", logger.ErrorOpt{
			Name:          errors.Name(errors.ErrServerClosed),
			Message:       errors.Message(errors.ErrServerClosed),
			OriginalError: err.Error(),
		})
	}

	logger.Info("Server exiting")
}
