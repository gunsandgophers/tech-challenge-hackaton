package repositories

import (
	"tech-challenge-hackaton/internal/core/entities"
	"tech-challenge-hackaton/internal/infra/database"
)

type VideoRepositoryDB struct {
	conn database.ConnectionDB
}

func NewVideoRepositoryDB(conn database.ConnectionDB) *VideoRepositoryDB {
	return &VideoRepositoryDB{conn: conn}
}

func (r *VideoRepositoryDB) Insert(video *entities.Video) error {
	sql := `
	INSERT INTO videos(id, filename, status, mimeType)
	VALUES ($1, $2, $3, $4);
	`
	return r.conn.Exec(
		sql,
		video.GetID(),
		video.GetFilename(),
		video.GetStatus().String(),
		video.GetMimeType().String(),
	)
}
