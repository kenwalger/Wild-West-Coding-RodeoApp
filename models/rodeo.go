package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Rodeo struct {
	ID        primitive.ObjectID `json:"_id"`
	Name      string             `json:"name"`
	ProRodeo  bool               `json:"pro_rodeo"`
	StartDate string             `json:"start_date"`
	EndDate   string             `json:"end_date"`
	Venue     struct {
		VenueId        int       `json:"venue_id"`
		Name           string    `json:"name"`
		Seats          [][]int   `json:"seats"`
		SeatsAvailable int       `json:"seats_available"`
		Coordinates    []float64 `json:"coordinates"`
	} `json:"venue"`
	PublishedAt time.Time `json:"published_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Events      struct {
		RoughStock []string `json:"rough_stock"`
		Timed      []string `json:"timed"`
		Other      []string `json:"other"`
	} `json:"events"`
}
