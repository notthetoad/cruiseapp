package crew

import (
	"cruiseapp/model"
	"cruiseapp/repository"
	"database/sql"
)

type CrewRankRepository interface {
	FindById(id int64) (*model.CrewRank, error)
	Save(rank *model.CrewRank) error
	Update(rank *model.CrewRank) error
	Delete(id int64) error
}

type PgCrewRankRepository struct {
	conn *sql.DB
}

func NewPgCrewRankRepository(conn *sql.DB) PgCrewRankRepository {
	return PgCrewRankRepository{
		conn: conn,
	}
}

func (repo PgCrewRankRepository) FindById(id int64) (*model.CrewRank, error) {
	var cr model.CrewRank
	err := repo.conn.QueryRow("SELECT id, name FROM crew_rank WHERE id = $1", id).Scan(&cr.Id, &cr.Name)
	if err != nil {
		return nil, &repository.NotFoundError{}
	}

	return &cr, nil
}

func (repo PgCrewRankRepository) Save(cr *model.CrewRank) error {
	var id int64
	err := repo.conn.QueryRow("INSERT INTO crew_rank (name) VALUES ($1) RETURNING id", cr.Name).Scan(&id)
	cr.Id = id

	return err
}

func (repo PgCrewRankRepository) Update(cr *model.CrewRank) error {
	stmt := `UPDATE crew_rank SET name = $1 WHERE id = $2`
	res, err := repo.conn.Exec(stmt, cr.Name, cr.Id)
	if err != nil {
		return &repository.NotFoundError{}
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return &repository.NotFoundError{}
	}
	if rows != 1 {
		return &repository.NotFoundError{}
	}

	return nil

}

func (repo PgCrewRankRepository) Delete(id int64) error {
	res, err := repo.conn.Exec("DELETE FROM crew_rank WHERE id = $1", id)
	if err != nil {
		return &repository.NotFoundError{}
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return &repository.NotFoundError{}
	}
	if rows != 1 {
		return &repository.NotFoundError{}
	}

	return nil
}
