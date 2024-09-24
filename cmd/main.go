package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	config "go-fiber-hex-arch/internal/adapter/config"
	"go-fiber-hex-arch/internal/adapter/http"
	myLogger "go-fiber-hex-arch/internal/adapter/logger"
	"go-fiber-hex-arch/internal/adapter/middleware"
	"go-fiber-hex-arch/internal/core/service"
	"go-fiber-hex-arch/internal/storage/mongodb"
	"go-fiber-hex-arch/internal/storage/mysql"
	"go-fiber-hex-arch/internal/storage/mysql/repository"
	"go-fiber-hex-arch/internal/util"
	"log/slog"
	"os"
	"time"
)

func main() {
	if err := util.InitTimeZone(); err != nil {
		panic(err.Error())
	}

	time.Local = util.Loc

	cfg, err := config.New()
	if err != nil {
		slog.Error("Error loading environment", "error", err)
		os.Exit(1)
	}

	myLogger.Set()

	slog.Info("Starting the app", "app", cfg.App.Name, "env", cfg.App.Env)

	//_, err = mongodb.NewDB(cfg.DB)
	//if err != nil {
	//	slog.Error("Error init MongoDB", "error", err)
	//	os.Exit(1)
	//}

	app := fiber.New()

	mySqlDB, err := mysql.ConnectMySQL(cfg.DB)
	if err != nil {
		panic(err)
	}

	mongoDB, err := mongodb.NewDB(cfg.MONGO)
	if err != nil {
		panic(err)
	}

	productRepo := repository.NewProductRepositoryDB(mySqlDB)
	productService := service.NewProductService(productRepo)

	app.Use(middleware.MonitoringFuncPerformance(mongoDB, cfg))

	http.SetupRoutes(app, productService, mongoDB, cfg.HTTP)

	app.Use(logger.New(logger.Config{
		Format:     "${time} [${ip}]:${port} | ${status} | ${method} ${path} | ${latency}\n",
		TimeFormat: "2006/01/02 15:04:05",
	}))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON("Hello Word")
	})

	listenPort := fmt.Sprintf(":%s", cfg.HTTP.Port)
	err = app.Listen(listenPort)
	if err != nil {
		slog.Error(fmt.Sprintf("Error listen to port %s", cfg.HTTP.Port), "error", err)
		os.Exit(1)
	}
}
