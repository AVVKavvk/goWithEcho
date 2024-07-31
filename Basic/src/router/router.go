package router

import (
	"fmt"
	"log"

	"github.com/AVVKavvk/Echo/src/config"
	"github.com/AVVKavvk/Echo/src/controller"
	custommiddleware "github.com/AVVKavvk/Echo/src/middleware"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start() {
	e := echo.New()
	err:=cleanenv.ReadEnv(&config.Cfg);
	if err!=nil {
		log.Fatal(err)
	}
	e.Pre(middleware.RemoveTrailingSlash())
	e.GET("/", controller.ServeHome,custommiddleware.CustomMiddleware)
	e.GET("/product/:id",controller.GetProductById)
	e.GET("/products",controller.GetAllProduct)

	e.POST("/product",controller.CreateProduct,middleware.BodyLimit("1K"))
	e.POST("/pro/:vander",controller.QueryParamDemo)

	e.PUT("/product/:id",controller.UpdateProductById)
	e.DELETE("/product/:id",controller.DeleteProductById)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s",config.Cfg.Port)))
}