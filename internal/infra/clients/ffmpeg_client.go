package clients

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type FFMPEGClient struct {
}

func NewFFMPEGClient() *FFMPEGClient {
	return &FFMPEGClient{}
}

func (f *FFMPEGClient) VideoDirationInSeconds(videoPath string) (float64, error) {
	durationBytes, err := exec.Command(
		"ffprobe",
		"-v",
		"error",
		"-show_entries",
		"format=duration",
		"-of",
		"default=noprint_wrappers=1:nokey=1",
		videoPath,
	).Output()
	if err != nil {
		return 0.0, err
	}
	return strconv.ParseFloat(strings.TrimSpace(string(durationBytes)), 64)
}

func (f *FFMPEGClient) Snapshot(videoPath string, momentInSeconds float64) error {
	t := time.Unix(int64(momentInSeconds), 0).UTC()
	timeFormat := t.Format(time.TimeOnly)
	frameFileName := fmt.Sprintf("%sframe_at_%s.jpg", videoPath, timeFormat)
	_, err := exec.Command("ffmpeg", "-ss", timeFormat, "-i", videoPath, "-frames:v", "1", "-q:v", "2", frameFileName).Output()
	return err
}
