package services

type SnapshotServiceInterface interface {
	Snapshot(videoID, filename string, interval int) (string, error)
}

