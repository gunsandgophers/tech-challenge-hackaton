package app

import (
	"tech-challenge-hackaton/internal/application/repositories"
	"tech-challenge-hackaton/internal/application/services"
	httpserver "tech-challenge-hackaton/internal/infra/http"
)

type APIApp struct {
	httpServer      httpserver.HTTPServer
	storageService  services.StorageServiceInterface
	videoRepository repositories.VideoRepositoryInterface
	queueService    services.QueueServiceInterface
}

func NewAPIApp(
	httpServer httpserver.HTTPServer,
	storageService services.StorageServiceInterface,
	videoRepository repositories.VideoRepositoryInterface,
	queueService services.QueueServiceInterface,
) *APIApp {
	app := &APIApp{}

	// // HTTP SERVER
	app.httpServer = httpServer

	// SERVICES
	app.storageService = storageService
	app.queueService = queueService

	//REPOSITORIES
	app.videoRepository = videoRepository

	// ROUTES
	app.configRoutes()
	return app
}

func (app *APIApp) configRoutes() {
	registerRouters(app)
}

func (app *APIApp) HTTPServer() httpserver.HTTPServer {
	return app.httpServer
}

func (app *APIApp) Run() {
	app.httpServer.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
