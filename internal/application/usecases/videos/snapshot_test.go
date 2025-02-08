package videos

import (
	"errors"
	"fmt"
	"os"
	"tech-challenge-hackaton/internal/application/services"
	"tech-challenge-hackaton/internal/tests/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func TestSnapshotUseCase_Execute(t *testing.T) {

	queueService := mocks.NewQueueServiceInterfaceMock(t)
	storageService := mocks.NewStorageServiceInterfaceMock(t)
	snapshotService := mocks.NewSnapshotServiceInterfaceMock(t)

	input := SnapshotInput{
		VideoID:  uuid.NewString(),
		Filename: "algumVideo.mp4",
	}

	dir := fmt.Sprintf("/tmp/video/%s", input.VideoID)
	filenameZip := "algumVideo.zip"
	filenameCompleteZip := fmt.Sprintf("%s/%s", dir, filenameZip)

	type args struct {
		input SnapshotInput
	}
	tests := []struct {
		name    string
		s       *SnapshotUseCase
		args    args
		wantErr bool
	}{
		{
			name: "Success: snapshot",
			s: NewSnapshotUseCase(
				queueService,
				storageService,
				snapshotService,
			),
			args:    args{input: input},
			wantErr: false,
		},
		{
			name: "Failure: Download video",
			s: NewSnapshotUseCase(
				queueService,
				storageService,
				snapshotService,
			),
			args:    args{input: input},
			wantErr: true,
		},
		{
			name: "Failure: snapshot",
			s: NewSnapshotUseCase(
				queueService,
				storageService,
				snapshotService,
			),
			args:    args{input: input},
			wantErr: true,
		},
		{
			name: "Failure: upload video frame zip",
			s: NewSnapshotUseCase(
				queueService,
				storageService,
				snapshotService,
			),
			args:    args{input: input},
			wantErr: true,
		},
		{
			name: "Failure: send video processed message",
			s: NewSnapshotUseCase(
				queueService,
				storageService,
				snapshotService,
			),
			args:    args{input: input},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				t.Errorf("SnapshotUseCase.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			err = os.WriteFile(filenameCompleteZip, []byte("Hello"), 0755)
			if err != nil {
				t.Errorf("SnapshotUseCase.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.name == "Success: snapshot" {
				storageService.On("DownloadVideo", tt.args.input.VideoID,
					tt.args.input.Filename).Return("download", nil).Once()

				storageService.On("GetLocalVideoDir", tt.args.input.VideoID).
					Return(dir).Twice()

				snapshotService.On("Snapshot", tt.args.input.VideoID,
					dir, tt.args.input.Filename, 20).
					Return(filenameCompleteZip, filenameZip, nil).Once()

				storageService.On("UploadZipFrames", filenameZip, mock.Anything).Return("", nil).Once()

				queueService.On("SendVideoProcessedMessage", services.VideoProcessedMessage{
					VideoID:  tt.args.input.VideoID,
					Filename: tt.args.input.Filename,
				}).Return(nil).Once()
			}
			if tt.name == "Failure: Download video" {
				storageService.On("DownloadVideo", tt.args.input.VideoID,
					tt.args.input.Filename).Return("", errors.New("error")).Once()
			}
			if tt.name == "Failure: snapshot" {
				storageService.On("DownloadVideo", tt.args.input.VideoID,
					tt.args.input.Filename).Return("download", nil).Once()

				storageService.On("GetLocalVideoDir", tt.args.input.VideoID).
					Return(dir).Once()

				snapshotService.On("Snapshot", tt.args.input.VideoID,
					dir, tt.args.input.Filename, 20).
					Return("", "", errors.New("error")).Once()
			}
			if tt.name == "Failure: upload video frame zip" {
				storageService.On("DownloadVideo", tt.args.input.VideoID,
					tt.args.input.Filename).Return("download", nil).Once()

				storageService.On("GetLocalVideoDir", tt.args.input.VideoID).
					Return(dir).Twice()

				snapshotService.On("Snapshot", tt.args.input.VideoID,
					dir, tt.args.input.Filename, 20).
					Return(filenameCompleteZip, filenameZip, nil).Once()

				storageService.On("UploadZipFrames", filenameZip, mock.Anything).Return("", errors.New("error")).Once()
			}
			if tt.name == "Failure: send video processed message" {
				storageService.On("DownloadVideo", tt.args.input.VideoID,
					tt.args.input.Filename).Return("download", nil).Once()

				storageService.On("GetLocalVideoDir", tt.args.input.VideoID).
					Return(dir).Twice()

				snapshotService.On("Snapshot", tt.args.input.VideoID,
					dir, tt.args.input.Filename, 20).
					Return(filenameCompleteZip, filenameZip, nil).Once()

				storageService.On("UploadZipFrames", filenameZip, mock.Anything).Return("", nil).Once()

				queueService.On("SendVideoProcessedMessage", services.VideoProcessedMessage{
					VideoID:  tt.args.input.VideoID,
					Filename: tt.args.input.Filename,
				}).Return(errors.New("error")).Once()
			}
			if err := tt.s.Execute(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("SnapshotUseCase.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
