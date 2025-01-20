package main

import (
	"log"
	"tech-challenge-hackaton/internal/infra/app"
	"tech-challenge-hackaton/internal/infra/clients"
	"tech-challenge-hackaton/internal/infra/database"
	httpserver "tech-challenge-hackaton/internal/infra/http"
	"tech-challenge-hackaton/internal/infra/repositories"
	"tech-challenge-hackaton/internal/infra/services"
)

func main() {
	httpServer := httpserver.NewGinHTTPServerAdapter()
	connection := database.NewPGXConnectionAdapter()

	// CLIENTS
	storageClient := clients.NewS3Client()
	queueClient := clients.NewSQSClient()

	// SERVICES
	storageService := services.NewAwsS3Service(storageClient)
	queueService := services.NewAwsSQSService(queueClient)

	// REPOSITORIES
	videoRepository := repositories.NewVideoRepositoryDB(connection)

	app := app.NewAPIApp(
		httpServer,
		storageService,
		videoRepository,
		queueService,
	)

	app.Run()
	defer connection.Close()
}
