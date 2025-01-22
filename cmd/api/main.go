package main

import (
	"tech-challenge-hackaton/configs"
	"tech-challenge-hackaton/internal/infra/clients"
	"tech-challenge-hackaton/internal/infra/database"
	httpserver "tech-challenge-hackaton/internal/infra/http"
	"tech-challenge-hackaton/internal/infra/repositories"
	"tech-challenge-hackaton/internal/infra/services"
	"tech-challenge-hackaton/internal/utils"
	"tech-challenge-hackaton/internal/web/app"
)

func main() {
	config := utils.Must(configs.LoadConfig("../.", ".env.dev"))
	// HTTP Server and Database Connection
	httpServer := httpserver.NewGinHTTPServerAdapter()
	connection := database.NewPGXConnectionAdapter(
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPassword,
		config.DBName,
	)
	// CLIENTS
	storageClient := clients.NewS3Client(
		config.AWSRegion,
		config.AWSAccessKeyID,
		config.AWSSecretAccessKey,
		config.AWSBaseEndpoint,
	)
	queueClient := clients.NewSQSClient(
		config.AWSRegion,
		config.AWSAccessKeyID,
		config.AWSSecretAccessKey,
		config.AWSBaseEndpoint,
	)
	cognitoClient := clients.NewCognitoClient(
		config.AWSRegion,
		config.AWSAppClientID,
		config.AWSUserPoolID,
	)
	// SERVICES
	storageService := services.NewAwsS3Service(storageClient, config.AWSBucketName)
	queueService := services.NewAwsSQSService(queueClient, config.QueueProcessVideo)
	userManagerService := services.NewAWSCognitoService(cognitoClient)
	// REPOSITORIES
	videoRepository := repositories.NewVideoRepositoryDB(connection)
	// APP
	app := app.NewAPIApp(
		httpServer,
		storageService,
		videoRepository,
		queueService,
		userManagerService,
	)
	app.Run()
	defer connection.Close()
}
