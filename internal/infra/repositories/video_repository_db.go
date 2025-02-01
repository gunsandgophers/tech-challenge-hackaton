package repositories

import (
	"log"
	"tech-challenge-hackaton/internal/application/entities"
	"tech-challenge-hackaton/internal/application/vo"
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
	INSERT INTO videos(id, user_id, filename, status, mime_type)
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

func (r *VideoRepositoryDB) ListByUserID(userID string) ([]*entities.Video, error) {
	sql := `
	SELECT 
		id,
		user_id,
		filename,
		status,
		mime_type,
	FROM public.videos
	WHERE
		user_id = $1
	`
	rows, err := r.conn.Query(sql, userID)
	if err != nil {
		return nil, err
	}
	return r.toDomainList(rows), nil
}

func (r *VideoRepositoryDB) toDomainList(rows database.RowsDB) []*entities.Video {
	var videos []*entities.Video
	for rows.Next() {
		if v, err := r.toDomain(rows); err == nil {
			videos = append(videos, v)
		} else {
			log.Println(err)
		}
	}
	return videos
}

func (r *VideoRepositoryDB) toDomain(row database.RowDB) (*entities.Video, error) {
	var (
		id       string
		userID   string
		filename string
		status   string
		mimeType string
	)

	err := row.Scan(
		&id,
		&userID,
		&filename,
		&status,
		&mimeType,
	)
	if err != nil {
		return nil, err
	}

	return entities.RestoreVideo(
		id,
		userID,
		vo.VideoStatus(status),
		filename,
		vo.MIMEType(mimeType),
	)
}
