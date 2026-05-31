package repository

import (
    "database/sql"

    "github.com/georgigeorgiev/cloudops/auth-service/internal/model"
)

type PostgresUserRepository struct {
    DB *sql.DB
}

func (r *PostgresUserRepository) Create(
    user *model.User,
) error {

    _, err := r.DB.Exec(
        `
        INSERT INTO users
        (email,password,role)
        VALUES($1,$2,$3)
        `,
        user.Email,
        user.Password,
        user.Role,
    )

    return err
}

func (r *PostgresUserRepository) FindByEmail(
    email string,
) (*model.User,error) {

    var user model.User

    err := r.DB.QueryRow(
        `
        SELECT id,email,password,role
        FROM users
        WHERE email=$1
        `,
        email,
    ).Scan(
        &user.ID,
        &user.Email,
        &user.Password,
        &user.Role,
    )

    if err != nil {
        return nil, err
    }

    return &user,nil
}
