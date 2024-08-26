// internal/storage/db.go

package storage

import (
    "database/sql"
    "fmt"
    "log"

 //   "github.com/lib/pq"
    "github.com/marktsarkov/notesapp/internal/config"
)

type DB struct {
    Conn *sql.DB
}

func NewDB(cfg config.DatabaseConfig) (*DB, error) {
    dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
        cfg.Host, 
		cfg.Port, 
		cfg.User, 
		cfg.Password, 
		cfg.DBName, 
		cfg.SSLMode)

    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }

    if err = db.Ping(); err != nil {
        return nil, err
    }

    log.Println("Успешное подключение к postgres")
    return &DB{Conn: db}, nil
}
