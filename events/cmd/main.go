package cmd

import (
	"context"
	"cookdie/events/internal/server"

	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func SetupLogger() *zap.SugaredLogger {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating zap logger: %v\n", err)
		os.Exit(1)
	}
	defer zapLogger.Sync()

	return zapLogger.Sugar()

}

func initGinServer() {

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus((http.StatusOK))
			return
		}

		c.Next()
	})

	dbpool, derr := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	if derr != nil {
		log.Fatalf("Unable to create connection pool: %v\n", derr)
		os.Exit(1)
	}

	logger := SetupLogger()

	server.StartEventServer(logger, dbpool, r)
	// restaurantRouteHandler := restauants.NewRouteHandler()
	// restaurantService.RegisterApiRoutes(group)
	r.Run(":3000")

}

func main() {
	// ctx := context.Background()
	initGinServer()
}
