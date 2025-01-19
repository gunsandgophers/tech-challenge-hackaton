package app

import (
	"tech-challenge-hackaton/internal/core/repositories"
	"tech-challenge-hackaton/internal/core/services"
	httpserver "tech-challenge-hackaton/internal/infra/http"
)

type APIApp struct {
	httpServer      httpserver.HTTPServer
	storageService  services.StorageServiceInterface
	videoRepository repositories.VideoRepositoryInterface
}

func NewAPIApp(
	httpServer httpserver.HTTPServer,
	storageService services.StorageServiceInterface,
	videoRepository repositories.VideoRepositoryInterface,
) *APIApp {
	app := &APIApp{}

	// // HTTP SERVER
	app.httpServer = httpServer

	// SERVICES
	app.storageService = storageService

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
