package repository

import (
	"database/sql"

	"github.com/train-do/Golang-Generics/model"
)

type RepoLocation struct {
	Db *sql.DB
}

func NewRepoLocation(db *sql.DB) *RepoLocation {
	return &RepoLocation{db}
}
func (r *RepoLocation) FindByDestinationId(Location *[]model.Location, id int) error {
	query := `select * from "Location" l where destination_id = $1;`
	rows, _ := r.Db.Query(query, id)
	for rows.Next() {
		var location model.Location
		rows.Scan(&location.Id, &location.DestinationId, &location.Lat, &location.Lng)
		*Location = append(*Location, location)
	}
	return nil
}
