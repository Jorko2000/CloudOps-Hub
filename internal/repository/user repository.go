package repository

import "github.com/georgigeorgiev/cloudops/auth-service/internal/model"

type UserRepository interface {

    Create(
        user *model.User,
    ) error

    FindByEmail(
        email string,
    ) (*model.User,error)
}
