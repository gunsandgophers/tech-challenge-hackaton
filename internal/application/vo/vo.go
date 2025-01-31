package vo

import (
	"tech-challenge-hackaton/internal/application/errors"
	"tech-challenge-hackaton/internal/utils"
)

type (
	VideoStatus string
	MIMEType    string
)

const (
	VideoStatusAwaiting VideoStatus = "AWAITING"
	VideoStatusFinished VideoStatus = "FINISHED"
	VideoStatusCanceled VideoStatus = "CANCELED"
)

const (
	MIMETypeMPEG MIMEType = "video/mpeg"
	MIMETypeMP4  MIMEType = "video/mp4"
)

var ValidVideoStatus = []VideoStatus{
	VideoStatusAwaiting,
	VideoStatusCanceled,
	VideoStatusFinished,
}
var ValidMIMETypes = []MIMEType{MIMETypeMPEG, MIMETypeMP4}

func (s VideoStatus) String() string {
	return string(s)
}

func (s VideoStatus) Validate() error {
	if (!utils.AssertIn(s, ValidVideoStatus)) {
		return errors.ErrVideoStatusInvalid
	}
	return nil
}

func (s MIMEType) String() string {
	return string(s)
}

func (s MIMEType) Validate() error {
	if (!utils.AssertIn(s, ValidMIMETypes)) {
		return errors.ErrMimeTypeInvalid
	}
	return nil
}

