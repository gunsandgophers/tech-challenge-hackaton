package main

import (
	"tech-challenge-hackaton/configs"
	"tech-challenge-hackaton/internal/consumers"
	"tech-challenge-hackaton/internal/infra/clients"
	"tech-challenge-hackaton/internal/infra/services"
	"tech-challenge-hackaton/internal/utils"
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
	ffmpegClient := clients.NewFFMPEGClient()
	// SERVICES
	storageService := services.NewAwsS3Service(storageClient, config.AWSS3BucketName)
	queueService := services.NewAwsSQSService(
		queueClient,
		config.QueueProcessVideo,
		config.QueueResultVideo,
	)
	snapshotService := services.NewFFMPEGService(ffmpegClient)

	consumer := consumers.NewVideoUploadedConsumer(
		10,
		queueService,
		storageService,
		snapshotService,
	)
	consumer.Run()
}
