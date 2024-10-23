package factory

import (
	"context"
	"cruiseapp/repository/crew"
	"cruiseapp/repository/cruise"
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
	CreateCruiseRepo() cruise.CruiseRepository
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

func (factory PgRepoFactory) CreateCruiseRepo() cruise.CruiseRepository {
	repo := cruise.NewPgCruiseRepository(factory.Conn)

	return repo
}

func CtxWithRepoFactory(ctx context.Context, factory RepoFactory) context.Context {
	return context.WithValue(ctx, MIDDLEWARE_CTX_KEY, factory)
}

func GetRepoFactory(r *http.Request) RepoFactory {
	return r.Context().Value(MIDDLEWARE_CTX_KEY).(RepoFactory)
}
