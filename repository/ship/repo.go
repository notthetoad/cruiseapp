package ship

import (
	"cruiseapp/model"
	"database/sql"
)

type ShipModelRepository interface {
	Save(shipModel *model.ShipModel) error
}

type ShipRepository interface {
	Save(ship *model.Ship) error
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

func (repo PgShipRepository) Save(ship *model.Ship) error {
	var id int64
	err := repo.conn.QueryRow("INSERT INTO ship (name, serial_number, ship_model_id) VALUES ($1, $2, $3) RETURNING id", ship.Name, ship.SerialNumber, ship.ShipModelId).Scan(&id)
	ship.Id = id

	return err
}
