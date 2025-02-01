package app

import (
	"tech-challenge-hackaton/internal/web/controllers"
	httpserver "tech-challenge-hackaton/internal/infra/http"
)

func registerRouters(app *APIApp) {
	videoController := controllers.NewVideoController(
		app.storageService,
		app.videoRepository,
		app.queueService,
		app.userManagerService,
	)
	authController := controllers.NewAuthControllerController(app.userManagerService)

	baseUrl := "/api/v1"
	app.httpServer.(httpserver.HTTPRoutes).SetBasePath(baseUrl)

	app.httpServer.(httpserver.HTTPRoutes).POST("/auth/login", authController.Login)
	app.httpServer.(httpserver.HTTPRoutes).GET("/auth/protected", authController.EndpointProtectedByAccessToken)
	app.httpServer.(httpserver.HTTPRoutes).POST("/videos/upload", videoController.Upload)
	app.httpServer.(httpserver.HTTPRoutes).GET("/videos/", videoController.List)
	app.httpServer.(httpserver.HTTPRoutes).GET("/videos/download/:id", videoController.Download)
}
