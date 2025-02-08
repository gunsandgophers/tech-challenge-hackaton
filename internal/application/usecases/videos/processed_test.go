package videos

import (
	"errors"
	"tech-challenge-hackaton/internal/application/entities"
	"tech-challenge-hackaton/internal/application/vo"
	"tech-challenge-hackaton/internal/tests/mocks"
	"testing"

	"github.com/google/uuid"
)

func TestUpdateProcessedVideoUseCase_Execute(t *testing.T) {

	videoRepository := mocks.NewVideoRepositoryInterfaceMock(t)

	videoInput := UpdateProcessedVideoInput{
		VideoID:  uuid.NewString(),
		Filename: "/tmp/algumVideo.mp4",
	}

	video, _ := entities.RestoreVideo(
		videoInput.VideoID, uuid.NewString(),
		vo.VideoStatusAwaiting, videoInput.Filename,
		vo.MIMETypeMP4,
	)

	type args struct {
		input UpdateProcessedVideoInput
	}
	tests := []struct {
		name    string
		u       *UpdateProcessedVideoUseCase
		args    args
		wantErr bool
	}{
		{
			name: "Success: processed video",
			u:    NewUpdateProcessedVideoUseCase(videoRepository),
			args: args{
				input: videoInput,
			},
			wantErr: false,
		},
		{
			name: "Failure: get video",
			u:    NewUpdateProcessedVideoUseCase(videoRepository),
			args: args{
				input: videoInput,
			},
			wantErr: true,
		},
		{
			name: "Failure: update video",
			u:    NewUpdateProcessedVideoUseCase(videoRepository),
			args: args{
				input: videoInput,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Success: processed video" {
				videoRepository.On("Get", videoInput.VideoID).Return(video, nil).Once()
				videoRepository.On("Update", video).Return(nil).Once()
			}
			if tt.name == "Failure: get video" {
				videoRepository.On("Get", videoInput.VideoID).Return(nil, errors.New("error")).Once()

			}
			if tt.name == "Failure: update video" {
				videoRepository.On("Get", videoInput.VideoID).Return(video, nil).Once()
				videoRepository.On("Update", video).Return(errors.New("error")).Once()
			}
			if err := tt.u.Execute(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("UpdateProcessedVideoUseCase.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
