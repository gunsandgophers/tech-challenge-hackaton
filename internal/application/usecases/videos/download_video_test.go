package videos

import (
	"errors"
	"fmt"
	"reflect"
	"tech-challenge-hackaton/internal/application/entities"
	"tech-challenge-hackaton/internal/application/vo"
	"tech-challenge-hackaton/internal/tests/mocks"
	"testing"

	"github.com/google/uuid"
)

func TestDownloadVideoFramesUseCase_Execute(t *testing.T) {

	storageService := mocks.NewStorageServiceInterfaceMock(t)
	videoRepository := mocks.NewVideoRepositoryInterfaceMock(t)

	video, _ := entities.RestoreVideo(
		uuid.NewString(),
		uuid.NewString(),
		vo.VideoStatusFinished,
		"/tmp/algumvideo.mp4",
		vo.MIMETypeMP4,
	)

	type args struct {
		videoID string
		userID  string
	}
	tests := []struct {
		name    string
		uv      *DownloadVideoFramesUseCase
		args    args
		want    *DownloadVideoFramesUseCaseOutput
		wantErr bool
	}{
		{
			name: "Success: download video frames",
			uv:   NewDownloadVideoFramesUseCase(storageService, videoRepository),
			args: args{videoID: video.GetID(), userID: video.GetUserID()},
			want: &DownloadVideoFramesUseCaseOutput{
				Filename: fmt.Sprintf("%s.zip", video.GetID()),
				MIMEType: "application/zip",
				Content:  []byte{},
			},
			wantErr: false,
		},
		{
			name:    "Failure: get video error",
			uv:      NewDownloadVideoFramesUseCase(storageService, videoRepository),
			args:    args{videoID: video.GetID(), userID: video.GetUserID()},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Failure: download zip",
			uv:      NewDownloadVideoFramesUseCase(storageService, videoRepository),
			args:    args{videoID: video.GetID(), userID: video.GetUserID()},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.name == "Success: download video frames" {
				videoRepository.On("Get", tt.args.videoID).Return(video, nil).Once()
				storageService.On("DownloadZipFrames", tt.args.videoID).Return([]byte{}, nil).Once()
			}
			if tt.name == "Failure: get video error" {
				videoRepository.On("Get", tt.args.videoID).Return(nil, errors.New("error")).Once()
			}
			if tt.name == "Failure: download zip" {
				videoRepository.On("Get", tt.args.videoID).Return(video, nil).Once()
				storageService.On("DownloadZipFrames", tt.args.videoID).Return(nil, errors.New("error")).Once()
			}
			got, err := tt.uv.Execute(tt.args.videoID, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DownloadVideoFramesUseCase.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DownloadVideoFramesUseCase.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
