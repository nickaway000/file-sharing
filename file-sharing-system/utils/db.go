package utils

import (
    "context"
    "github.com/jackc/pgx/v4"
    "log"
)

var Db *pgx.Conn

func ConnectDB() {
    var err error
    connStr := "postgres://postgres:Nikhil%40007@localhost:5432/file_sharing_system?sslmode=disable"

 
    Db, err = pgx.Connect(context.Background(), connStr)
    if err != nil {
        log.Fatal("Unable to connect to the database:", err)
    }
}
