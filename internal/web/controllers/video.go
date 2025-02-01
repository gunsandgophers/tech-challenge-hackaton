package controllers

import (
	"log"
	"net/http"
	"tech-challenge-hackaton/internal/application/repositories"
	"tech-challenge-hackaton/internal/application/services"
	"tech-challenge-hackaton/internal/application/usecases/videos"
	httpserver "tech-challenge-hackaton/internal/infra/http"
)

type VideoController struct {
	storageService  services.StorageServiceInterface
	videoRepository repositories.VideoRepositoryInterface
	queueService    services.QueueServiceInterface
	userManagerService    services.UserManagerServiceInterface
}

func NewVideoController(
	storageService services.StorageServiceInterface,
	videoRepository repositories.VideoRepositoryInterface,
	queueService services.QueueServiceInterface,
	userManagerService    services.UserManagerServiceInterface,
) *VideoController {
	return &VideoController{
		storageService:  storageService,
		videoRepository: videoRepository,
		queueService:    queueService,
		userManagerService: userManagerService,
	}
}

func (cc *VideoController) Upload(c httpserver.HTTPContext) {
	token, err := cc.userManagerService.ValidateAccessTokenByAuthHeader(c.GetHeader("Authorization"))
	if err != nil {
		sendError(c, http.StatusUnauthorized, err.Error())
		return
	}
	user, err := cc.userManagerService.GetUser(token)
	if err != nil {
		sendError(c, http.StatusUnauthorized, err.Error())
		return
	}

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

		videoUpload, err := usecase.Execute(filename, file, header.Get("Content-Type"), user.ID)
		if err != nil {
			log.Println("Error on upload file - ", err)
		} else {
			videosUpload = append(videosUpload, videoUpload)
		}
	}
	response := videos.VideoUploadResponseDTO{Videos: videosUpload}
	sendSuccess(c, http.StatusCreated, "Upload finished", response)
}

func (cc *VideoController) List(c httpserver.HTTPContext) {
	token, err := cc.userManagerService.ValidateAccessTokenByAuthHeader(c.GetHeader("Authorization"))
	if err != nil {
		sendError(c, http.StatusUnauthorized, err.Error())
		return
	}
	user, err := cc.userManagerService.GetUser(token)
	if err != nil {
		sendError(c, http.StatusUnauthorized, err.Error())
		return
	}

	listVideos := videos.NewListVideosUseCase(cc.videoRepository)
	output, err := listVideos.Execute(user.ID)
	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(c, http.StatusOK, "list videos", output)
}
