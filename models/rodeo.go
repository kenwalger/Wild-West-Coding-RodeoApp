package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Rodeo struct {
	ID        primitive.ObjectID
	Name      string
	ProRodeo  bool
	StartDate string
	EndDate   string
	Venue     struct {
		VenueId        int
		Name           string
		Seats          [][]int
		SeatsAvailable int
		Coordinates    []float64
	}
	PublishedAt time.Time
	UpdatedAt   time.Time
	Events      struct {
		RoughStock []string
		Timed      []string
		Other      []string
	}
}
