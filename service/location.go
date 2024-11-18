package service

import (
	"github.com/train-do/Golang-Generics/model"
	"github.com/train-do/Golang-Generics/repository"
)

type ServiceLocation struct {
	Repo *repository.RepoLocation
}

func NewServiceLocation(repo *repository.RepoLocation) *ServiceLocation {
	return &ServiceLocation{repo}
}
func (s *ServiceLocation) GetByDestinationId(Location *[]model.Location, id int) error {
	err := s.Repo.FindByDestinationId(Location, id)
	if err != nil {
		return err
	}
	return nil
}
