package main

import (
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

	// SERVICES
	storageService := services.NewAwsS3Service(storageClient)

	// REPOSITORIES
	videoRepository := repositories.NewVideoRepositoryDB(connection)

	app := app.NewAPIApp(
		httpServer,
		storageService,
		videoRepository,
	)

	app.Run()
	defer connection.Close()
}
