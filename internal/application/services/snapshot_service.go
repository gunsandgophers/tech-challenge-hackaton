package services

type SnapshotServiceInterface interface {
	Snapshot(videoID string, localVideoDir string, filename string, interval int) (string, error)
}

