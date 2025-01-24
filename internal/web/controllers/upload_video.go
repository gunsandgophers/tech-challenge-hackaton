package controllers

import (
	"log"
	"net/http"
	"tech-challenge-hackaton/internal/application/repositories"
	"tech-challenge-hackaton/internal/application/services"
	"tech-challenge-hackaton/internal/application/usecases/videos"
	httpserver "tech-challenge-hackaton/internal/infra/http"
)

type UploadVideoController struct {
	storageService  services.StorageServiceInterface
	videoRepository repositories.VideoRepositoryInterface
	queueService    services.QueueServiceInterface
	userManagerService    services.UserManagerServiceInterface
}

func NewUploadVideoController(
	storageService services.StorageServiceInterface,
	videoRepository repositories.VideoRepositoryInterface,
	queueService services.QueueServiceInterface,
	userManagerService    services.UserManagerServiceInterface,
) *UploadVideoController {
	return &UploadVideoController{
		storageService:  storageService,
		videoRepository: videoRepository,
		queueService:    queueService,
		userManagerService: userManagerService,
	}
}

func (cc *UploadVideoController) UploadVideos(c httpserver.HTTPContext) {
	form, err := c.MultipartForm()
	if err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	usecase := videos.NewUploadVideoUseCase(
		cc.storageService,
		cc.videoRepository,
		cc.queueService,
	)

	videosUpload := []*videos.VideoUploadDTO{}

	for _, video := range form.File {

		filename := video[0].Filename
		header := video[0].Header

		file, err := video[0].Open()
		if err != nil {
			log.Println("Error on open file - ", err)
		}

		videoUpload, err := usecase.Execute(filename, file, header.Get("Content-Type"))
		if err != nil {
			log.Println("Error on upload file - ", err)
		} else {
			videosUpload = append(videosUpload, videoUpload)
		}

	}

	response := videos.VideoUploadResponseDTO{Videos: videosUpload}

	sendSuccess(c, http.StatusCreated, "Upload finished", response)
}

