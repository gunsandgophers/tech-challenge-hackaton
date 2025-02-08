package entities

import (
	"tech-challenge-hackaton/internal/application/vo"
	"testing"

	"github.com/google/uuid"
)

func TestVideo_Validate(t *testing.T) {
	type fields struct {
		id       string
		userID   string
		status   vo.VideoStatus
		filename string
		mimeType vo.MIMEType
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{ 
		{
			name:"Success: validate",
			fields: fields{
				id: uuid.NewString(),
				userID: uuid.NewString(),
				status: vo.VideoStatusAwaiting,
				filename: "algumVideo.mp4",
				mimeType: vo.MIMETypeMP4,
			},
			wantErr: false,
		},
		{
			name:"Failure: video id",
			fields: fields{},
			wantErr: true,
		},
		{
			name:"Failure: user id",
			fields: fields{
				id: uuid.NewString(),
			},
			wantErr: true,
		},
		{
			name:"Failure: filename",
			fields: fields{
				id: uuid.NewString(),
				userID: uuid.NewString(),
			},
			wantErr: true,
		},
		{
			name:"Failure: status",
			fields: fields{
				id: uuid.NewString(),
				userID: uuid.NewString(),
				filename: "algumVideo.mp4",
			},
			wantErr: true,
		},
		{
			name:"Failure: mime type",
			fields: fields{
				id: uuid.NewString(),
				userID: uuid.NewString(),
				status: vo.VideoStatusAwaiting,
				filename: "algumVideo.mp4",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Video{
				id:       tt.fields.id,
				userID:   tt.fields.userID,
				status:   tt.fields.status,
				filename: tt.fields.filename,
				mimeType: tt.fields.mimeType,
			}
			if err := v.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Video.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVideo_IsAvaiableToDownload(t *testing.T) {
	userID := uuid.NewString()

	type fields struct {
		id       string
		userID   string
		status   vo.VideoStatus
		filename string
		mimeType vo.MIMEType
	}
	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Success: is avaiable",
			fields: fields{
				id: uuid.NewString(),
				userID: userID,
				status: vo.VideoStatusFinished,
				filename: "algumVideo.mp4",
				mimeType: vo.MIMETypeMP4,
			},
			args: args{userID: userID},
			wantErr: false,
		},
		{
			name: "Failure: different user id",
			fields: fields{
				id: uuid.NewString(),
				userID: uuid.NewString(),
				status: vo.VideoStatusFinished,
				filename: "algumVideo.mp4",
				mimeType: vo.MIMETypeMP4,
			},
			args: args{userID: uuid.NewString()},
			wantErr: true,
		},
		{
			name: "Failure: status not finished",
			fields: fields{
				id: uuid.NewString(),
				userID: uuid.NewString(),
				status: vo.VideoStatusAwaiting,
				filename: "algumVideo.mp4",
				mimeType: vo.MIMETypeMP4,
			},
			args: args{userID: uuid.NewString()},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Video{
				id:       tt.fields.id,
				userID:   tt.fields.userID,
				status:   tt.fields.status,
				filename: tt.fields.filename,
				mimeType: tt.fields.mimeType,
			}
			if err := v.IsAvaiableToDownload(tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("Video.IsAvaiableToDownload() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
