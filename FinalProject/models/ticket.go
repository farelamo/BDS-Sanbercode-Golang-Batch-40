package models

import "time"

type Ticket struct {
	Id 			int 		`json:"id,omitempty"`
	Type 		string 		`json:"type,omitempty"`
	UserId 		int 		`json:"user_id,omitempty"`
	StudioId 	int 		`json:"studio_id,omitempty"`
	CreatedAt   time.Time 	`json:"created_at,omitempty"`
	UpdatedAt   time.Time 	`json:"updated_at,omitempty"`
}

type AddTicket struct {
	Type 		string 		`json:"type,omitempty"`
	UserId 		int 		`json:"user_id,omitempty"`
	StudioId 	int 		`json:"studio_id,omitempty"`
}