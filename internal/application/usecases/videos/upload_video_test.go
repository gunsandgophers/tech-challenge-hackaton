package videos

import (
	"errors"
	"mime/multipart"
	"tech-challenge-hackaton/internal/application/vo"
	"tech-challenge-hackaton/internal/tests/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func TestUploadVideoUseCase_Execute(t *testing.T) {

	storageService := mocks.NewStorageServiceInterfaceMock(t)
	videoRepository := mocks.NewVideoRepositoryInterfaceMock(t)
	queueService := mocks.NewQueueServiceInterfaceMock(t)

	filename := "filename"

	type args struct {
		filename string
		file     multipart.File
		mimeType string
		userID   string
	}
	tests := []struct {
		name    string
		uv      *UploadVideoUseCase
		args    args
		wantErr bool
	}{
		{
			name: "Success: upload video",
			uv: NewUploadVideoUseCase(
				storageService,
				videoRepository,
				queueService,
			),
			args: args{
				filename: "algumVideo.mp4",
				file:     nil,
				mimeType: vo.MIMETypeMP4.String(),
				userID:   uuid.NewString(),
			},
			wantErr: false,
		},
		{
			name: "Failure: upload video",
			uv: NewUploadVideoUseCase(
				storageService,
				videoRepository,
				queueService,
			),
			args: args{
				filename: "algumVideo.mp4",
				file:     nil,
				mimeType: vo.MIMETypeMP4.String(),
				userID:   uuid.NewString(),
			},
			wantErr: true,
		},
		{
			name: "Failure: insert video",
			uv: NewUploadVideoUseCase(
				storageService,
				videoRepository,
				queueService,
			),
			args: args{
				filename: "algumVideo.mp4",
				file:     nil,
				mimeType: vo.MIMETypeMP4.String(),
				userID:   uuid.NewString(),
			},
			wantErr: true,
		},
		{
			name: "Failure: send video uploaded message",
			uv: NewUploadVideoUseCase(
				storageService,
				videoRepository,
				queueService,
			),
			args: args{
				filename: "algumVideo.mp4",
				file:     nil,
				mimeType: vo.MIMETypeMP4.String(),
				userID:   uuid.NewString(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Success: upload video" {
				storageService.On("UploadVideo", mock.Anything, mock.Anything).
					Return(filename, nil).Once()
				videoRepository.On("Insert", mock.Anything).Return(nil).Once()
				queueService.On("SendVideoUploadedMessage", mock.Anything).
					Return(nil).Once()
			}
			if tt.name == "Failure: upload video" {
				storageService.On("UploadVideo", mock.Anything, mock.Anything).
					Return("", errors.New("error")).Once()
			}
			if tt.name == "Failure: insert video" {
				storageService.On("UploadVideo", mock.Anything, mock.Anything).
					Return(filename, nil).Once()
				videoRepository.On("Insert", mock.Anything).Return(errors.New("error")).Once()
			}
			if tt.name == "Failure: send video uploaded message" {
				storageService.On("UploadVideo", mock.Anything, mock.Anything).
					Return(filename, nil).Once()
				videoRepository.On("Insert", mock.Anything).Return(nil).Once()
				queueService.On("SendVideoUploadedMessage", mock.Anything).
					Return(errors.New("error")).Once()
			}
			_, err := tt.uv.Execute(tt.args.filename, tt.args.file, tt.args.mimeType, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UploadVideoUseCase.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
