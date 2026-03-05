package server

import (
	eventsHttp "cookdie/events/http"
	"cookdie/events/internal/aws"
	"cookdie/events/internal/producer"
	eventService "cookdie/events/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

type EventServiceConfig struct {
	QueueURL string
}

type App struct {
	eventService eventsHttp.EventService
}

func (a *App) RegisterApiRoutes(r *gin.Engine) error {
	logger.Info("Registering Api Routes for event...")

	routeHandler := eventsHttp.NewRouteHandler(
		&eventsHttp.RouteHandlerConfig{
			EventService: a.eventService,
		},
	)

	routeHandler.RegisterRoutes(r)

	return nil

}
func NewApp(cfgLogger *zap.SugaredLogger, config *EventServiceConfig) *App {
	logger = cfgLogger

	if len(config.QueueURL) == 0 {
		logger.Fatal("Queue's URL is not defined")
	}
	awsConfig, err := aws.LoadAWSConfig()

	if err != nil {
		logger.Error("awsConfig failed", err.Error())
	}
	SQSQueue := producer.NewSQSProducer(awsConfig, config.QueueURL)

	service := eventService.NewEventService(
		&eventService.EventServiceConfig{
			Queue: SQSQueue,
		})

	return &App{
		eventService: service,
	}
}

func StartEventServer(logger *zap.SugaredLogger, db *pgxpool.Pool, r *gin.Engine) {
	eventApp := NewApp(logger, &EventServiceConfig{
		QueueURL: "",
	})

	eventApp.RegisterApiRoutes(r)
}
