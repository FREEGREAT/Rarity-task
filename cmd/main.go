package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"test_task/internal/handler"
	"test_task/internal/handler/middleware"
	"test_task/internal/service/api"
	"test_task/pkg/logger"
	"test_task/pkg/server"
)

func main() {

	log := logger.New("debug")

	if err := godotenv.Load(); err != nil {
		log.Warn("No .env file found — using environment variables from Kubernetes")
	}

	raribleClient := api.NewRaribleApiClient(os.Getenv("BASE_URL"), os.Getenv("MAINNET"))

	handlers := handler.NewHandler(log, raribleClient)
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.RequestLogger(log))
	r.Use(middleware.Recovery(log))

	handlers.InitRoutes(r)

	srv := new(server.Server)
	port := "8080"

	log.Info("Server is running,", "port", port)

	if err := srv.Run(port, r); err != nil {
		log.Fatal("Error occurred while running HTTP server")
	}

}
