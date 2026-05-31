package database

import (
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"
)

func Connect(
    host string,
    port string,
    user string,
    pass string,
    db string,
) (*sql.DB,error) {

    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        host,
        port,
        user,
        pass,
        db,
    )

    return sql.Open("postgres", dsn)
}
