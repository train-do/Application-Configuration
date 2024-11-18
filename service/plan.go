package service

import (
	"github.com/train-do/Golang-Generics/model"
	"github.com/train-do/Golang-Generics/repository"
)

type ServiceTourPlan struct {
	Repo *repository.RepoTourPlan
}

func NewServiceTourPlan(repo *repository.RepoTourPlan) *ServiceTourPlan {
	return &ServiceTourPlan{repo}
}
func (s *ServiceTourPlan) GetByDestinationId(TourPlan *[]model.TourPlan, id int) error {
	err := s.Repo.FindByDestinationId(TourPlan, id)
	if err != nil {
		return err
	}
	return nil
}
