package handler

import (
	"cruiseapp/model"
	"cruiseapp/repository/crew"
	"cruiseapp/repository/cruise"
	"cruiseapp/repository/person"
	"cruiseapp/repository/port"
	"cruiseapp/repository/ship"
)

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

	return nil
}

func (factory MockRepoFactory) CreatePersonRepo() person.PersonRepository {

	return nil
}

func (factory MockRepoFactory) CreateCruiseRepo() cruise.CruiseRepository {

	return nil
}

type MockPortRepository struct {
}

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
