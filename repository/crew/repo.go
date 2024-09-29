package crew

import (
	"cruiseapp/model"
	"cruiseapp/repository"
	"database/sql"
)

type CrewRankRepository interface {
	FindById(id int64) (*model.CrewRank, error)
	Save(cr *model.CrewRank) error
	Update(cr *model.CrewRank) error
	Delete(id int64) error
}
type CrewMemberRepository interface {
	FindById(id int64) (*model.CrewMember, error)
	Save(cm *model.CrewMember) error
	Update(cm *model.CrewMember) error
	Delete(id int64) error
}

type PgCrewRankRepository struct {
	conn *sql.DB
}

type PgCrewMemberRepository struct {
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

func NewPgCrewMemberRepository(conn *sql.DB) PgCrewMemberRepository {
	return PgCrewMemberRepository{
		conn: conn,
	}
}

func (repo PgCrewMemberRepository) FindById(id int64) (*model.CrewMember, error) {
	var cm model.CrewMember
	err := repo.conn.QueryRow("SELECT id, crew_rank, person_id FROM crew_member WHERE id = $1", id).Scan(&cm.Id, &cm.CrewRankId, &cm.PersonId)
	if err != nil {
		return nil, &repository.NotFoundError{}
	}

	return &cm, nil
}

func (repo PgCrewMemberRepository) Save(cm *model.CrewMember) error {
	var id int64
	err := repo.conn.QueryRow("INSERT INTO crew_member (crew_rank, person_id) VALUES ($1, $2) RETURNING id", cm.CrewRankId, cm.PersonId).Scan(&id)
	cm.Id = id

	return err

}

func (repo PgCrewMemberRepository) Update(cm *model.CrewMember) error {
	stmt := "UPDATE crew_member SET crew_rank = $1 WHERE id = $2"
	res, err := repo.conn.Exec(stmt, cm.CrewRankId, cm.Id)
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

func (repo PgCrewMemberRepository) Delete(id int64) error {
	res, err := repo.conn.Exec("DELETE FROM crew_member WHERE id = $1", id)
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
