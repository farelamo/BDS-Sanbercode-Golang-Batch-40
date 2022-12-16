package models

import "time"

type Ticket struct {
	Id 			int 		`json:"id,omitempty"`
	Type 		string 		`json:"type,omitempty"`
	ProfileId 	int 		`json:"profile_id,omitempty"`
	StudioId 	int 		`json:"studio_id,omitempty"`
	CreatedAt   time.Time 	`json:"created_at,omitempty"`
	UpdatedAt   time.Time 	`json:"updated_at,omitempty"`
}