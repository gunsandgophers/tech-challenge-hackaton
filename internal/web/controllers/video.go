package controllers

import (
	"fmt"
	"log"
	"net/http"
	"tech-challenge-hackaton/internal/application/repositories"
	"tech-challenge-hackaton/internal/application/services"
	"tech-challenge-hackaton/internal/application/usecases/videos"
	httpserver "tech-challenge-hackaton/internal/infra/http"
)

type VideoController struct {
	storageService     services.StorageServiceInterface
	videoRepository    repositories.VideoRepositoryInterface
	queueService       services.QueueServiceInterface
	userManagerService services.UserManagerServiceInterface
}

func NewVideoController(
	storageService services.StorageServiceInterface,
	videoRepository repositories.VideoRepositoryInterface,
	queueService services.QueueServiceInterface,
	userManagerService services.UserManagerServiceInterface,
) *VideoController {
	return &VideoController{
		storageService:     storageService,
		videoRepository:    videoRepository,
		queueService:       queueService,
		userManagerService: userManagerService,
	}
}

// Checkout godoc
//
// @Summary		Upload Videos
// @Description	Upload videos to snapshot
// @Tags				videos
// @Accept multipart/form-data
// @Produce		json
// @Param file formData file true "video para upload"
// @Security BearerToken
// @Success		200			{object}	JSONResponse{data=videos.VideoUploadResponseDTO}  "token"
// @Failure		400			{object}	ErrorJSONResponse	"when bad request"
// @Failure		406			{object}	ErrorJSONResponse	"when invalid params or invalid object"
// @Router			/videos/upload [post]
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

// Checkout godoc
//
// @Summary		List Uploaded Videos
// @Description	List uploaded videos by logged user
// @Tags				videos
// @Produce		json
// @Security BearerToken
// @Success		200			{object}	JSONResponse{data=videos.ListVideosOutput}  "token"
// @Failure		400			{object}	ErrorJSONResponse	"when bad request"
// @Failure		406			{object}	ErrorJSONResponse	"when invalid params or invalid object"
// @Router			/videos [get]
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

// Checkout godoc
//
// @Summary		Download Videos Snapshot
// @Description	Downalod videos snapshot
// @Tags				videos
// @Produce application/zip
// @Param id path string true "video ID for download"
// @Security BearerToken
// @Success 200 {file} application/zip "ZIP File with snapshots"
// @Failure		400			{object}	ErrorJSONResponse	"when bad request"
// @Failure		406			{object}	ErrorJSONResponse	"when invalid params or invalid object"
// @Router			/videos/download/{id} [get]
func (cc *VideoController) Download(c httpserver.HTTPContext) {
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

	videoID := c.Param("id")
	downloadVideo := videos.NewDownloadVideoFramesUseCase(cc.storageService, cc.videoRepository)
	output, err := downloadVideo.Execute(videoID, user.ID)
	if err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+output.Filename)
	c.Header("Content-Type", output.MIMEType)
	c.Header("Accept-Length", fmt.Sprintf("%d", len(output.Content)))
	c.Data(http.StatusOK, output.MIMEType, output.Content)
}
