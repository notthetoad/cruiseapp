package port

import (
	"cruiseapp/model"
	"cruiseapp/repository"
	"database/sql"
)

type PortRepository interface {
	FindById(id int64) (*model.Port, error)
	Save(port *model.Port) error
}

type PgPortRepository struct {
	conn *sql.DB
}

func NewPgPortRepository(conn *sql.DB) PgPortRepository {
	return PgPortRepository{
		conn: conn,
	}
}

func (pr PgPortRepository) FindById(id int64) (*model.Port, error) {
	var p model.Port
	row := pr.conn.QueryRow(`SELECT id, location FROM port WHERE id = $1`, id)
	err := row.Scan(&p.Id, &p.Location)
	if err != nil {
		return nil, &repository.NotFoundError{}
	}
	return &p, nil
}

func (pr PgPortRepository) Save(port *model.Port) error {
	var id int64
	err := pr.conn.QueryRow("INSERT INTO port (location) VALUES ($1) RETURNING id", port.Location).Scan(&id)
	port.Id = id

	return err
}
