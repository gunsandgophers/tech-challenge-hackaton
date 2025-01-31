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
	config := utils.Must(configs.LoadConfig("../.", ".env"))
	// CLIENTS
	storageClient := clients.NewS3Client(
		config.AWSS3Region,
		config.AWSS3AccessKeyID,
		config.AWSS3SecretAccessKey,
		config.AWSS3BaseEndpoint,
	)
	queueClient := clients.NewSQSClient(
		config.AWSSQSRegion,
		config.AWSSQSAccessKeyID,
		config.AWSSQSSecretAccessKey,
		config.AWSSQSBaseEndpoint,
	)
	// SERVICES
	storageService := services.NewAwsS3Service(storageClient, config.AWSS3BucketName)
	queueService := services.NewAwsSQSService(queueClient, config.QueueProcessVideo)


}

