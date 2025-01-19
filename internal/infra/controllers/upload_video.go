package controllers

import (
	"log"
	"net/http"
	"tech-challenge-hackaton/internal/core/dtos"
	"tech-challenge-hackaton/internal/core/repositories"
	"tech-challenge-hackaton/internal/core/services"
	"tech-challenge-hackaton/internal/core/use_cases/videos"
	httpserver "tech-challenge-hackaton/internal/infra/http"
)

type UploadVideoController struct {
	storageService  services.StorageServiceInterface
	videoRepository repositories.VideoRepositoryInterface
}

func NewUploadVideoController(
	storageService services.StorageServiceInterface,
	videoRepository repositories.VideoRepositoryInterface,

) *UploadVideoController {
	return &UploadVideoController{
		storageService:  storageService,
		videoRepository: videoRepository,
	}
}

func (cc *UploadVideoController) UploadVideos(c httpserver.HTTPContext) {
	form, err := c.MultipartForm()
	if err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	usecase := videos.NewUploadVideoUseCase(cc.storageService, cc.videoRepository)

	videosUpload := []*dtos.VideoUploadDTO{}

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

	response := dtos.VideoUploadResponseDTO{Videos: videosUpload}

	sendSuccess(c, http.StatusCreated, "Upload finished", response)
}
