package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PGXConnectionAdapter struct {
	pool *pgxpool.Pool
}

func NewPGXConnectionAdapter(DBHost, DBPort, DBUser, DBPassword, DBName string) *PGXConnectionAdapter {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s",
		DBHost, DBPort, DBUser, DBPassword, DBName)

	pool, err := pgxpool.New(context.Background(), psqlInfo)
	if err != nil {
		log.Println("Unable to create connection pool", err)
		panic(1)
	}

	return &PGXConnectionAdapter{pool: pool}
}

func (p *PGXConnectionAdapter) QueryRow(sql string, args ...interface{}) RowDB {
	return p.pool.QueryRow(context.Background(), sql, args...)
}

func (p *PGXConnectionAdapter) Query(sql string, args ...interface{}) (RowsDB, error) {
	return p.pool.Query(context.Background(), sql, args...)
}

func (p *PGXConnectionAdapter) Exec(sql string, args ...interface{}) error {
	_, err := p.pool.Exec(context.Background(), sql, args...)
	return err
}

func (p *PGXConnectionAdapter) Close() {
	p.pool.Close()
}
