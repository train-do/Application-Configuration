package model

import "github.com/lib/pq"

type TourPlan struct {
	Id            int
	DestinationId int
	Day           int
	Description   string
	Facilities    pq.StringArray
}
