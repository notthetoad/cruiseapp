package handler

import (
	"context"
	"cruiseapp/model"
	"cruiseapp/repository/crew"
	"cruiseapp/repository/cruise"
	"cruiseapp/repository/factory"
	"cruiseapp/repository/person"
	"cruiseapp/repository/port"
	"cruiseapp/repository/ship"
	"net/http/httptest"
)

func setupRecorderAndCtx() (*httptest.ResponseRecorder, context.Context) {
	rr := httptest.NewRecorder()
	ctx := ctxWithMockRepoFactory(context.Background())
	// hub := MockHub{}
	// nCtx := context.WithValue(ctx, ws.WS_HUB_CTX_KEY, hub)

	return rr, ctx
}

func ctxWithMockRepoFactory(ctx context.Context) context.Context {
	var repoFactory factory.RepoFactory = MockRepoFactory{}
	return factory.CtxWithRepoFactory(ctx, repoFactory)
}

type MockRepoFactory struct {
}

func (factory MockRepoFactory) CreatePortRepo() port.PortRepository {

	return MockPortRepository{}
}

func (factory MockRepoFactory) CreateShipModelRepo() ship.ShipModelRepository {

	return nil
}

func (factory MockRepoFactory) CreateShipRepo() ship.ShipRepository {

	return nil
}

func (factory MockRepoFactory) CreateCrewRankRepo() crew.CrewRankRepository {

	return nil
}

func (factory MockRepoFactory) CreateCrewMemberRepo() crew.CrewMemberRepository {

	return MockCrewMemberRepository{}
}

func (factory MockRepoFactory) CreatePersonRepo() person.PersonRepository {

	return MockPersonRepository{}
}

func (factory MockRepoFactory) CreateCruiseRepo() cruise.CruiseRepository {

	return MockCruiseRepository{}
}

type MockPortRepository struct{}

func (repo MockPortRepository) FindById(id int64) (*model.Port, error) {
	var p model.Port
	return &p, nil
}

func (repo MockPortRepository) Save(port *model.Port) error {
	port.Id = 1
	return nil
}

func (repo MockPortRepository) Update(port *model.Port) error {
	return nil
}

func (repo MockPortRepository) Delete(id int64) error {
	return nil
}

type MockCruiseRepository struct{}

func (repo MockCruiseRepository) FindById(id int64) (*model.Cruise, error) {
	var c model.Cruise
	return &c, nil
}

func (repo MockCruiseRepository) Save(cruise *model.Cruise) error {
	cruise.Id = 1
	return nil
}

func (repo MockCruiseRepository) Update(cruise *model.Cruise) error {
	return nil
}

func (repo MockCruiseRepository) Delete(id int64) error {
	return nil
}

type MockCrewMemberRepository struct{}

func (repo MockCrewMemberRepository) FindById(id int64) (*model.CrewMember, error) {
	var cm model.CrewMember
	return &cm, nil
}

func (repo MockCrewMemberRepository) FindAllByIds(ids []int64) ([]*model.CrewMember, error) {
	var cm []*model.CrewMember
	return cm, nil
}

func (repo MockCrewMemberRepository) Save(cm *model.CrewMember) error {
	cm.Id = 1
	return nil
}

func (repo MockCrewMemberRepository) Update(cruise *model.CrewMember) error {
	return nil
}

func (repo MockCrewMemberRepository) Delete(id int64) error {
	return nil
}

type MockPersonRepository struct{}

func (repo MockPersonRepository) FindById(id int64) (*model.Person, error) {
	var p model.Person
	return &p, nil
}

func (repo MockPersonRepository) FindAllByIds(ids []int64) ([]*model.Person, error) {
	var p []*model.Person
	return p, nil
}

func (repo MockPersonRepository) Save(p *model.Person) error {
	p.Id = 1
	return nil
}

func (repo MockPersonRepository) Update(p *model.Person) error {
	return nil
}

func (repo MockPersonRepository) Delete(id int64) error {
	return nil
}

type MockHub struct{}

func (mh *MockHub) Run() {}
