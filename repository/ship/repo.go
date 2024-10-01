package ship

import (
	"cruiseapp/model"
	"cruiseapp/repository"
	"database/sql"
)

type ShipModelRepository interface {
	Save(shipModel *model.ShipModel) error
	FindById(id int64) (model.ShipModel, error)
	Update(shipModel *model.ShipModel) error
	Delete(id int64) error
}

type ShipRepository interface {
	Save(ship *model.Ship) error
	FindById(id int64) (model.Ship, error)
	Update(ship *model.Ship) error
	Delete(id int64) error
}

type PgShipModelRepository struct {
	conn *sql.DB
}

type PgShipRepository struct {
	conn *sql.DB
}

func NewPgShipModelRepository(conn *sql.DB) PgShipModelRepository {
	return PgShipModelRepository{
		conn: conn,
	}
}

func NewPgShipRepository(conn *sql.DB) PgShipRepository {
	return PgShipRepository{
		conn: conn,
	}
}

func (repo PgShipModelRepository) Save(shipModel *model.ShipModel) error {
	var id int64
	err := repo.conn.QueryRow("INSERT INTO ship_model (name) VALUES ($1) RETURNING id", shipModel.Name).Scan(&id)
	shipModel.Id = id

	return err
}

func (repo PgShipModelRepository) FindById(id int64) (model.ShipModel, error) {
	var sm model.ShipModel
	sm.Id = id
	err := repo.conn.QueryRow("SELECT name FROM ship_model WHERE id = $1", id).Scan(&sm.Name)
	if err != nil {
		return model.ShipModel{}, repository.NewNotFoundError(id)
	}

	return sm, nil
}

func (repo PgShipModelRepository) Update(shipModel *model.ShipModel) error {
	_, err := repo.conn.Exec("UPDATE ship_model SET name = $1 WHERE id = $2", shipModel.Name, shipModel.Id)

	return err
}

func (repo PgShipModelRepository) Delete(id int64) error {
	res, err := repo.conn.Exec("DELETE FROM ship_model WHERE id = $1", id)
	if err != nil {
		return repository.NewForbiddenActionError(id, "delete").WithDetails(err.Error())
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return repository.NewNotFoundError(id)
	}

	return nil
}

func (repo PgShipRepository) Save(ship *model.Ship) error {
	var id int64
	err := repo.conn.QueryRow("INSERT INTO ship (name, serial_number, ship_model_id) VALUES ($1, $2, $3) RETURNING id", ship.Name, ship.SerialNumber, ship.ShipModelId).Scan(&id)
	ship.Id = id

	return err
}

func (repo PgShipRepository) FindById(id int64) (model.Ship, error) {
	var s model.Ship
	s.Id = id
	err := repo.conn.QueryRow("SELECT name, serial_number, ship_model_id FROM ship WHERE id = $1", id).Scan(&s.Name, &s.SerialNumber, &s.ShipModelId)
	if err != nil {
		return model.Ship{}, repository.NewNotFoundError(id)
	}

	return s, nil
}

func (repo PgShipRepository) Update(ship *model.Ship) error {
	stmt := `
	UPDATE ship
	SET name = $1,
	    serial_number = $2,
	    ship_model_id = $3
	WHERE id = $4`
	res, err := repo.conn.Exec(stmt, ship.Name, ship.SerialNumber, ship.ShipModelId, ship.Id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return repository.NewNotFoundError(ship.Id)
	}

	return err
}

func (repo PgShipRepository) Delete(id int64) error {
	res, err := repo.conn.Exec("DELETE FROM ship WHERE id = $1", id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return repository.NewNotFoundError(id)
	}

	return nil
}
