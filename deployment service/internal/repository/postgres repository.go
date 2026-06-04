package repository

import (
	"database/sql"

	"deployment-service/internal/model"
)

type PostgresRepository struct {
	DB *sql.DB
}

func (r *PostgresRepository) Save(
	d *model.Deployment,
) error {

	_, err := r.DB.Exec(
		`
		INSERT INTO deployments
		(app_name,image,replicas,status)
		VALUES($1,$2,$3,$4)
		`,
		d.AppName,
		d.Image,
		d.Replicas,
		d.Status,
	)

	return err
}
