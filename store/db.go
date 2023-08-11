package store

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

var s *sqliteConnection

type sqliteConnection struct {
	dbfile string
	conn   *sql.DB
}

func (s *sqliteConnection) InitializeTables() error {
	_, err := s.conn.Exec(`
		CREATE TABLE IF NOT EXISTS OpnsOrder (
			id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			date DATE NOT NULL,
			total INTEGER NOT NULL,
			paid BOOL NOT NULL DEFAULT 0
		);

		CREATE TABLE IF NOT EXISTS OrderItem (
			id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			order_id INTEGER NOT NULL,
			details TEXT NOT NULL,
			price INTEGER NOT NULL,
			persons_nb INTEGER NOT NULL DEFAULT 0,
			kids_pets_nb INTEGER NOT NULL DEFAULT 0,
			type INTEGER NOT NULL,
			background INTEGER NOT NULL,
			done BOOL NOT NULL DEFAULT 0,
			FOREIGN KEY (order_id) REFERENCES OpnsOrder(id)
		);`,
	)
	return err
}

func (s *sqliteConnection) connect() (*sqliteConnection, error) {
	if s.conn != nil {
		return s, s.conn.Ping()
	}
	conn, err := sql.Open("sqlite3", s.dbfile)
	if err != nil {
		return nil, err
	} else if err := conn.Ping(); err != nil {
		return nil, err
	}
	s.conn = conn
	return s, nil
}

func GetConnection(dbfile string) (*sqliteConnection, error) {
	if dbfile == "" {
		return nil, errors.New("dbfile can't be an empty string!")
	}
	if s != nil && s.dbfile != dbfile {
		s.dbfile = dbfile
		return s.connect()
	}
	s = &sqliteConnection{dbfile: dbfile}
	return s.connect()
}
