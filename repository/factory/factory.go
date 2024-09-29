package factory

import (
	"context"
	"cruiseapp/database"
	"cruiseapp/repository/crew"
	"cruiseapp/repository/person"
	"cruiseapp/repository/port"
	"cruiseapp/repository/ship"
	"database/sql"
	"net/http"
)

const MIDDLEWARE_CTX_KEY = "repo_middleware_ctx_key"

type RepoFactory interface {
	CreatePortRepo() port.PortRepository
	CreateShipModelRepo() ship.ShipModelRepository
	CreateShipRepo() ship.ShipRepository
	CreateCrewRankRepo() crew.CrewRankRepository
	CreateCrewMemberRepo() crew.CrewMemberRepository
	CreatePersonRepo() person.PersonRepository
}

type PgRepoFactory struct {
	Conn *sql.DB
}

func (factory PgRepoFactory) CreatePortRepo() port.PortRepository {
	repo := port.NewPgPortRepository(factory.Conn)

	return repo
}

func (factory PgRepoFactory) CreateShipModelRepo() ship.ShipModelRepository {
	repo := ship.NewPgShipModelRepository(factory.Conn)

	return repo
}

func (factory PgRepoFactory) CreateShipRepo() ship.ShipRepository {
	repo := ship.NewPgShipRepository(factory.Conn)

	return repo
}

func (factory PgRepoFactory) CreateCrewRankRepo() crew.CrewRankRepository {
	repo := crew.NewPgCrewRankRepository(factory.Conn)

	return repo
}

func (factory PgRepoFactory) CreateCrewMemberRepo() crew.CrewMemberRepository {
	repo := crew.NewPgCrewMemberRepository(factory.Conn)

	return repo
}

func (factory PgRepoFactory) CreatePersonRepo() person.PersonRepository {
	repo := person.NewPgPersonRepository(factory.Conn)

	return repo
}

// TODO make it generic
func PgRepoFactoryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := database.GetDb(r)
		var factory RepoFactory = PgRepoFactory{Conn: db}

		ctx := context.WithValue(r.Context(), MIDDLEWARE_CTX_KEY, factory)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetRepoFactory(r *http.Request) RepoFactory {
	return r.Context().Value(MIDDLEWARE_CTX_KEY).(RepoFactory)
}
