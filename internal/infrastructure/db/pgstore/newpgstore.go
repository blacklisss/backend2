package pgstore

import "database/sql"

func NewDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
