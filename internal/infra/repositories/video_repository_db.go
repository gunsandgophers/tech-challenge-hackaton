package repositories

import (
	"tech-challenge-hackaton/internal/application/entities"
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
	INSERT INTO videos(id, user_id, filename, status, mimeType)
	VALUES ($1, $2, $3, $4, $5);
	`
	return r.conn.Exec(
		sql,
		video.GetID(),
		video.GetUserID(),
		video.GetFilename(),
		video.GetStatus().String(),
		video.GetMimeType().String(),
	)
}
