package port

import (
	"cruiseapp/model"
	"cruiseapp/repository"
	"database/sql"
)

type PortRepository interface {
	FindById(id int64) (*model.Port, error)
	Save(port *model.Port) error
	Update(port *model.Port) error
	Delete(id int64) error
}

type PgPortRepository struct {
	conn *sql.DB
}

func NewPgPortRepository(conn *sql.DB) PgPortRepository {
	return PgPortRepository{
		conn: conn,
	}
}

func (repo PgPortRepository) FindById(id int64) (*model.Port, error) {
	var p model.Port
	row := repo.conn.QueryRow(`SELECT id, location FROM port WHERE id = $1`, id)
	err := row.Scan(&p.Id, &p.Location)
	if err != nil {
		return nil, &repository.NotFoundError{}
	}
	return &p, nil
}

func (repo PgPortRepository) Save(port *model.Port) error {
	var id int64
	err := repo.conn.QueryRow("INSERT INTO port (location) VALUES ($1) RETURNING id", port.Location).Scan(&id)
	port.Id = id

	return err
}

func (repo PgPortRepository) Update(port *model.Port) error {
	stmt := "UPDATE port SET location = $1 WHERE id = $2"
	res, err := repo.conn.Exec(stmt, port.Location, port.Id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return err
	}

	return nil
}

func (repo PgPortRepository) Delete(id int64) error {
	res, err := repo.conn.Exec("DELETE FROM port WHERE id = $1", id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return err
	}

	return nil
}
