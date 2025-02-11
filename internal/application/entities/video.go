package entities

import (
	"errors"
	"fmt"
	"tech-challenge-hackaton/internal/application/vo"
	"tech-challenge-hackaton/internal/utils"

	"github.com/google/uuid"
)

type Video struct {
	id       string
	userID   string
	status   vo.VideoStatus
	filename string
	mimeType vo.MIMEType
}

func CreateVideo(userID, filename string, mimeType vo.MIMEType) (*Video, error) {
	return RestoreVideo(
		uuid.NewString(),
		userID,
		vo.VideoStatusAwaiting,
		filename,
		mimeType,
	)
}

func RestoreVideo(id, userID string, status vo.VideoStatus, filename string, mimeType vo.MIMEType) (*Video, error) {
	video := &Video{
		id:       id,
		userID:   userID,
		status:   status,
		filename: filename,
		mimeType: mimeType,
	}
	err := video.Validate()
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (v *Video) Validate() error {
	if !utils.AssertNotEmpty(v.id) {
		return errors.New("video id invalid")
	}
	if !utils.AssertNotEmpty(v.userID) {
		return errors.New("user id invalid")
	}
	if !utils.AssertNotEmpty(v.filename) {
		return errors.New("filename invalid")
	}
	if err := v.status.Validate(); err != nil {
		return err
	}
	if err := v.mimeType.Validate(); err != nil {
		return err
	}
	return nil
}

func (v *Video) GetID() string {
	return v.id
}

func (v *Video) GetUserID() string {
	return v.userID
}

func (v *Video) GetFilename() string {
	return v.filename
}

func (v *Video) GetFullFilename() string {
	return fmt.Sprintf("%s-%s", v.GetID(), v.GetFilename())
}

func (v *Video) GetStatus() vo.VideoStatus {
	return v.status
}

func (v *Video) GetMimeType() vo.MIMEType {
	return v.mimeType
}

func (v *Video) Finished() {
	v.status = vo.VideoStatusFinished
}

func (v *Video) IsAvaiableToDownload(userID string) error {
	if v.GetUserID() != userID {
		return fmt.Errorf("User is not the video owner")
	}
	if v.GetStatus() != vo.VideoStatusFinished {
		return fmt.Errorf("frames not avaiable yet")
	}
	return nil
}
