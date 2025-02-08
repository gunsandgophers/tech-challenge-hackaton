package videos

import (
	"errors"
	"reflect"
	"tech-challenge-hackaton/internal/application/entities"
	"tech-challenge-hackaton/internal/application/vo"
	"tech-challenge-hackaton/internal/tests/mocks"
	"testing"

	"github.com/google/uuid"
)

func TestListVideosUseCase_Execute(t *testing.T) {

	videoRepository := mocks.NewVideoRepositoryInterfaceMock(t)
	userID := uuid.NewString()

	video, _ := entities.CreateVideo(userID, "/tmp/algumVideo.mp4", vo.MIMETypeMP4)

	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		lv      *ListVideosUseCase
		args    args
		want    *ListVideosOutput
		wantErr bool
	}{
		{
			name: "Success: list videos",
			lv:   NewListVideosUseCase(videoRepository),
			args: args{userID: userID},
			want: &ListVideosOutput{Videos: []VideoOutput{
				{
					ID:       video.GetID(),
					UserID:   video.GetUserID(),
					Status:   string(video.GetStatus()),
					Filename: video.GetFilename(),
					MIMEType: video.GetMimeType().String(),
				},
			}},
			wantErr: false,
		},
		{
			name:    "Failure: list videos",
			lv:      NewListVideosUseCase(videoRepository),
			args:    args{userID: userID},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Success: list videos" {
				videoRepository.On("ListByUserID", tt.args.userID).Return([]*entities.Video{video}, nil).Once()
			}
			if tt.name == "Failure: list videos" {
				videoRepository.On("ListByUserID", tt.args.userID).Return(nil, errors.New("error")).Once()
			}
			got, err := tt.lv.Execute(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListVideosUseCase.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListVideosUseCase.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
