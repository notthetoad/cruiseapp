package person

import (
	"cruiseapp/model"
	"cruiseapp/repository"
	"database/sql"
	"fmt"
)

type PersonRepository interface {
	FindById(id int64) (*model.Person, error)
	FindAllByIds(ids []int64) ([]*model.Person, error)
	Save(p *model.Person) error
	Update(p *model.Person) error
	Delete(id int64) error
}

type PgPersonRepository struct {
	conn *sql.DB
}

func NewPgPersonRepository(conn *sql.DB) PgPersonRepository {
	return PgPersonRepository{
		conn: conn,
	}
}

func (repo PgPersonRepository) FindById(id int64) (*model.Person, error) {
	var p model.Person
	err := repo.conn.QueryRow("SELECT id, first_name, last_name, email, phone FROM person WHERE id = $1", id).Scan(&p.Id, &p.FirstName, &p.LastName, &p.Email, &p.Phone)
	if err != nil {
		return nil, repository.NewNotFoundError(id)
	}

	return &p, nil
}

func (repo PgPersonRepository) FindAllByIds(ids []int64) ([]*model.Person, error) {
	var people []*model.Person
	stmt, err := repo.conn.Prepare("SELECT id, first_name, last_name, email, phone FROM person WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	for _, id := range ids {
		fmt.Printf("looking up person %d\n", id)
		var p model.Person
		err := stmt.QueryRow(id).Scan(&p.Id, &p.FirstName, &p.LastName, &p.Email, &p.Phone)
		if err != nil {
			return nil, repository.NewNotFoundError(id)
		}
		people = append(people, &p)
	}

	return people, nil
}

func (repo PgPersonRepository) Save(p *model.Person) error {
	var id int64
	err := repo.conn.QueryRow("INSERT INTO person (first_name, last_name, email, phone) VALUES ($1, $2, $3, $4) RETURNING id", p.FirstName, p.LastName, p.Email, p.Phone).Scan(&id)
	p.Id = id

	return err
}

func (repo PgPersonRepository) Update(p *model.Person) error {
	stmt := `
	UPDATE person SET first_name = $1,
	last_name = $2,
	email = $3,
	phone = $4
	WHERE id = $5`

	res, err := repo.conn.Exec(stmt, p.FirstName, p.LastName, p.Email, p.Phone, p.Id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return repository.NewNotFoundError(p.Id)
	}

	return nil
}

func (repo PgPersonRepository) Delete(id int64) error {
	res, err := repo.conn.Exec("DELETE FROM person WHERE id = $1", id)
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
