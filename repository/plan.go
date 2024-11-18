package repository

import (
	"database/sql"

	"github.com/train-do/Golang-Generics/model"
)

type RepoTourPlan struct {
	Db *sql.DB
}

func NewRepoTourPlan(db *sql.DB) *RepoTourPlan {
	return &RepoTourPlan{db}
}
func (r *RepoTourPlan) FindByDestinationId(TourPlan *[]model.TourPlan, id int) error {
	query := `select * from "TourPlan" where destination_id = $1;`
	rows, _ := r.Db.Query(query, id)
	for rows.Next() {
		var tourPlan model.TourPlan
		rows.Scan(&tourPlan.Id, &tourPlan.DestinationId, &tourPlan.Day, &tourPlan.Description, &tourPlan.Facilities)
		*TourPlan = append(*TourPlan, tourPlan)
	}
	return nil
}
