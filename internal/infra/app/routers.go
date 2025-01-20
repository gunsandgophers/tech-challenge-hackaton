package app

import (
	"tech-challenge-hackaton/internal/infra/controllers"
	httpserver "tech-challenge-hackaton/internal/infra/http"
)

func registerRouters(app *APIApp) {

	videoController := controllers.
		NewUploadVideoController(app.storageService, app.videoRepository, app.queueService)

	baseUrl := "/api/v1"
	app.httpServer.(httpserver.HTTPRoutes).SetBasePath(baseUrl)

	// videos
	app.httpServer.(httpserver.HTTPRoutes).POST("/upload-video", videoController.UploadVideos)

}
