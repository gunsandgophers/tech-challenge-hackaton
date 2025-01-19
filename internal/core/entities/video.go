package entities

import (
	"tech-challenge-hackaton/internal/core/errors"

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

func (s MIMEType) IsValid() bool {

	for _, t := range MIME_TYPES {
		if t == s {
			return true
		}
	}

	return false
}

const (
	VIDEO_STATUS_AWAITING VideoStatus = "AWAITING"
	VIDEO_STATUS_FINISHED VideoStatus = "FINISHED"
	VIDEO_STATUS_CANCELED VideoStatus = "CANCELED"
)

const (
	MIME_TYPE_MPEG MIMEType = "video/mpeg"
	MIME_TYPE_MP4  MIMEType = "video/mp4"
)

var MIME_TYPES = []MIMEType{MIME_TYPE_MPEG, MIME_TYPE_MP4}

type Video struct {
	id       string
	status   VideoStatus
	filename string
	mimeType MIMEType
}

func CreateVideo(filename string, mimeType MIMEType) (*Video, error) {
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
) (*Video, error) {
	if mimeType.IsValid() {
		return &Video{
			id:       id,
			status:   status,
			filename: filename,
			mimeType: mimeType,
		}, nil
	}

	return nil, errors.ErrMimeTypeInvalid
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
