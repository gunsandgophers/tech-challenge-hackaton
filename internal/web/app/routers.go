package app

import (
	"tech-challenge-hackaton/internal/web/controllers"
	httpserver "tech-challenge-hackaton/internal/infra/http"
)

func registerRouters(app *APIApp) {
	videoController := controllers.
		NewUploadVideoController(app.storageService, app.videoRepository, app.queueService, app.userManagerService)
	authController := controllers.NewAuthControllerController(app.userManagerService)

	baseUrl := "/api/v1"
	app.httpServer.(httpserver.HTTPRoutes).SetBasePath(baseUrl)

	// videos
	app.httpServer.(httpserver.HTTPRoutes).POST("/upload-video", videoController.UploadVideos)
	app.httpServer.(httpserver.HTTPRoutes).POST("/auth/login", authController.Login)
	app.httpServer.(httpserver.HTTPRoutes).GET("/auth/protected", authController.EndpointProtectedByAccessToken)
}
