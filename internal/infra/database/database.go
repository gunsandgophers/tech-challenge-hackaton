package database

type RowDB interface {
	Scan(dest ...any) error
}

type RowsDB interface {
	Next() bool
	Scan(dest ...any) error
}

type ConnectionDB interface {
	QueryRow(sql string, args ...interface{}) RowDB
	Query(sql string, args ...interface{}) (RowsDB, error)
	Exec(sql string, args ...interface{}) error
	Close()
}
