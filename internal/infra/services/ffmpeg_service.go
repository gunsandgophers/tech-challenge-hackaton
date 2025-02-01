package services

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"tech-challenge-hackaton/internal/infra/clients"
)

type FFMPEGService struct {
	client *clients.FFMPEGClient
}

func NewFFMPEGService(client *clients.FFMPEGClient) *FFMPEGService {
	return &FFMPEGService{
		client: client,
	}
}

func (f *FFMPEGService) Snapshot(videoID string, localVideoDir string, filename string, interval int) (string, error) {
	duration, err := f.client.VideoDirationInSeconds(filename)
	if err != nil {
		return "", err
	}

	framesPath := filepath.Join(localVideoDir, "frames")
	if err := os.MkdirAll(filepath.Dir(framesPath), 0775); err != nil {
		return "", err
	}

	videoFilenameComplete := fmt.Sprintf("%s/%s", localVideoDir, filename)
	for curr := 0; curr < int(duration); curr += interval {
		if err := f.client.Snapshot(videoFilenameComplete, framesPath, curr); err != nil {
			log.Println(err.Error())
		}
	}

	zipFilenameComplete := fmt.Sprintf("%s/%s.zip", localVideoDir, videoID)
	zipDirectory(framesPath, zipFilenameComplete)
	return zipFilenameComplete, nil
}

func zipDirectory(sourceDir, outputZip string) error {
	zipFile, err := os.Create(outputZip)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	return filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		zipEntry, err := zipWriter.Create(relPath)
		if err != nil {
			return err
		}

		_, err = io.Copy(zipEntry, file)
		return err
	})
}
