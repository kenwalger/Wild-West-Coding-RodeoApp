package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// API rodeo data
// Rodeo data structure for the application
//
// swagger:model rodeo
type Rodeo struct {
	// ID for a specific rodeo
	//
	// required:true
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	// Name of the rodeo event
	//
	// required:true
	Name string `json:"name"`
	// ProRodeo is a True/False value of the Pro Rodeo status of the event.
	//
	// required:true
	ProRodeo bool `json:"pro_rodeo"`
	// StartDate is the start date of the rodeo
	//
	// required:false
	StartDate string `json:"start_date"`
	// EndDate is the end date of the rodeo
	//
	// required:false
	EndDate string `json:"end_date"`
	// Venue is a data structure about the rodeo venue
	//
	// required:false
	Venue struct {
		VenueId        int       `json:"venue_id"`
		Name           string    `json:"name"`
		Seats          [][]int   `json:"seats"`
		SeatsAvailable int       `json:"seats_available"`
		Coordinates    []float64 `json:"coordinates"`
	} `json:"venue"`
	// PublishedAt is the date the rodeo was added to the database and is auto-generated.
	//
	// required:true
	PublishedAt time.Time `json:"published_at"`
	// UpdatedAt is the date the rodeo was lasted updated in the database and is auto-generated.
	//
	// required:true
	UpdatedAt time.Time `json:"updated_at"`
	// Events is a data structure of the events offered at a rodeo
	//
	// required:false
	Events struct {
		RoughStock []string `json:"rough_stock"`
		Timed      []string `json:"timed"`
		Other      []string `json:"other"`
	} `json:"events"`
}
