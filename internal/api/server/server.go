package server

import (
	"context"
	"fmt"
	"github.com/8soat-grupo35/fastfood-order/external"
	"github.com/8soat-grupo35/fastfood-order/internal/api/handlers"
	"net/http"

	_ "github.com/8soat-grupo35/fastfood-order/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Start(cfg external.Config) {
	fmt.Println(context.Background(), fmt.Sprintf("Starting a server at http://%s", cfg.ServerHost))
	app := newApp(cfg)
	app.Logger.Fatal(app.Start(cfg.ServerHost))
}

// @title Swagger Fastfood App API
// @version 1.0
// @description This is a sample API from Fastfood App.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /v1
func newApp(cfg external.Config) *echo.Echo {
	external.ConectaDB(cfg.DatabaseConfig.Host, cfg.DatabaseConfig.User, cfg.DatabaseConfig.Password, cfg.DatabaseConfig.DbName, cfg.DatabaseConfig.Port)

	app := echo.New()
	app.GET("/swagger/*", echoSwagger.WrapHandler)
	app.GET("/", func(echo echo.Context) error {
		return echo.JSON(http.StatusOK, "Alive")
	})

	customerHandler := handlers.NewCustomerHandler(external.DB)
	customerGroupV1 := app.Group("/v1/customer")
	customerGroupV1.GET("", customerHandler.GetAll)
	customerGroupV1.GET("/cpf/:cpf", customerHandler.GetByCpf)
	customerGroupV1.POST("", customerHandler.Create)
	customerGroupV1.PUT("/:id", customerHandler.Update)
	customerGroupV1.DELETE("/:id", customerHandler.Delete)

	itemHandler := handlers.NewItemHandler(external.DB)
	itemV1Group := app.Group("/v1/item")
	itemV1Group.GET("", itemHandler.GetAll)
	itemV1Group.POST("", itemHandler.Create)
	itemV1Group.PUT("/:id", itemHandler.Update)
	itemV1Group.DELETE("/:id", itemHandler.Delete)

	orderHandler := handlers.NewOrderHandler(external.DB)
	orderV1Group := app.Group("/v1/orders")
	orderV1Group.GET("", orderHandler.GetAll)
	orderV1Group.POST("/checkout", orderHandler.Checkout)
	orderV1Group.PATCH("/:id", orderHandler.UpdateStatus)

	return app
}
