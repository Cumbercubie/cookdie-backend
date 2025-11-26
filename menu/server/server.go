package server

import (
	"cookdie/menu"
	menuHttp "cookdie/menu/http"
	sql "cookdie/menu/sql/db/query_gen"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type AppConfig struct {
	DbPool *pgxpool.Pool
}

var logger *zap.SugaredLogger

type App struct {
	menuService *menu.MenuService
}

func (a *App) RegisterApiRoutes(r *gin.Engine) error {
	logger.Info("Registering Api Routes for catalog app...")

	routeHandler := menuHttp.NewRouteHandler(
		&menuHttp.RouteHandlerConfig{
			MenuService: a.menuService,
		},
	)

	routeHandler.RegisterRoutes(r)

	return nil

}
func NewApp(cfgLogger *zap.SugaredLogger, config *AppConfig) *App {
	logger = cfgLogger

	dbHandler := sql.New(config.DbPool)

	menuService := menu.NewMenuService(
		&menu.MenuServiceConfig{
			Store: *dbHandler,
		})

	return &App{
		menuService: menuService,
	}
}

func StartMenuServer(logger *zap.SugaredLogger, db *pgxpool.Pool, r *gin.Engine) {
	menuApp := NewApp(logger, &AppConfig{
		DbPool: db,
	})

	menuApp.RegisterApiRoutes(r)
}
