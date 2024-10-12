package cruise

import (
	"cruiseapp/model"
	"database/sql"
	"log"
)

type CruiseRepository interface {
	FindById(id int64) (*model.Cruise, error)
	Save(c *model.Cruise) error
	// Update(c *model.Cruise) error
	// Delete(id int64) error
}

type PgCruiseRepository struct {
	conn *sql.DB
}

func NewPgCruiseRepository(conn *sql.DB) PgCruiseRepository {
	return PgCruiseRepository{
		conn: conn,
	}
}

func (repo PgCruiseRepository) FindById(id int64) (*model.Cruise, error) {
	var c model.Cruise
	err := repo.conn.QueryRow("SELECT id, start_date, end_date, from_location, to_location FROM cruise WHERE id = $1", id).Scan(&c.Id, &c.StartDate, &c.EndDate, &c.FromLocation.Id, &c.ToLocation.Id)
	if err != nil {
		return nil, err
	}
	portSelect := "SELECT location FROM port WHERE id = $1"
	err = repo.conn.QueryRow(portSelect, c.FromLocation.Id).Scan(&c.FromLocation.Location)
	if err != nil {
		return nil, err
	}
	err = repo.conn.QueryRow(portSelect, c.ToLocation.Id).Scan(&c.ToLocation.Location)
	if err != nil {
		return nil, err
	}
	rows, err := repo.conn.Query("SELECT id, person_id, crew_rank FROM crew_member cm JOIN cruise_crew_member ccm ON cm.id = ccm.crew_member_id WHERE ccm.cruise_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var cm model.CrewMember
		if err := rows.Scan(&cm.Id, &cm.PersonId, &cm.CrewRankId); err != nil {
			return nil, err
		}
		c.Crew = append(c.Crew, &cm)
	}

	return &c, nil
}

// TODO fix error handling
func (repo PgCruiseRepository) Save(c *model.Cruise) error {
	var id int64
	err := repo.conn.QueryRow(`INSERT INTO cruise (start_date, end_date, from_location, to_location) VALUES ($1, $2, $3, $4) RETURNING id`, c.StartDate, c.EndDate, c.FromLocation.Id, c.ToLocation.Id).Scan(&id)
	c.Id = id
	if err != nil {
		log.Println(err)
	}

	for _, crewMember := range c.Crew {
		_, err := repo.conn.Exec(`INSERT INTO cruise_crew_member (cruise_id, crew_member_id) VALUES ($1, $2)`, c.Id, crewMember.Id)
		if err != nil {
			log.Println(err)
		}
	}

	return err
}
