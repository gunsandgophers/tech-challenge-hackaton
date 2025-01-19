package entities

import (
	"github.com/google/uuid"
)

type (
	VideoStatus string
	MIMEType    string
)

func (s VideoStatus) String() string {
	return string(s)
}

func (s MIMEType) String() string {
	return string(s)
}

const (
	VIDEO_STATUS_AWAITING VideoStatus = "AWAITING"
	VIDEO_STATUS_FINISHED VideoStatus = "FINISHED"
	VIDEO_STATUS_CANCELED VideoStatus = "CANCELED"
)

const (
	MIME_TYPE_MPEG VideoStatus = "video/mpeg"
	MIME_TYPE_MP4  VideoStatus = "video/mp4"
	MIME_TYPE_OGG  VideoStatus = "video/ogg"
	MIME_TYPE_WEBM VideoStatus = "video/webm"
)

type Video struct {
	id       string
	status   VideoStatus
	filename string
	mimeType MIMEType
}

func CreateVideo(filename string, mimeType MIMEType) *Video {
	return RestoreVideo(
		uuid.NewString(),
		VIDEO_STATUS_AWAITING,
		filename,
		mimeType,
	)
}

func RestoreVideo(
	id string,
	status VideoStatus,
	filename string,
	mimeType MIMEType,
) *Video {
	return &Video{
		id:       id,
		status:   status,
		filename: filename,
		mimeType: mimeType,
	}
}

func (v *Video) GetID() string {
	return v.id
}

func (v *Video) GetFilename() string {
	return v.filename
}

func (v *Video) GetStatus() VideoStatus {
	return v.status
}

func (v *Video) GetMimeType() MIMEType {
	return v.mimeType
}
