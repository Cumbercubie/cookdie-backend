package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"cookdie/restaurants"
	// "cookdie/vendors"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// var env = os.Getenv("ENV")

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

	group := r.Group("/api/v1")

	restaurantRouteHandler := restaurants.NewRouteHandler()
	restaurantService.RegisterApiRoutes(group)
	r.Run(":3000")

}

func main() {
	// ctx := context.Background()
	initGinServer()
}
