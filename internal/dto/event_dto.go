package dto

import (
	"time"
)

type CreateEventRequest struct {
	Name        string    `json:"name" binding:""`
	Description string    `json:"description" binding:""`
	StartDate   time.Time `json:"startDate" binding:"required"`
	EndDate     time.Time `json:"endDate" binding:"required"`
}

type CreateEventResponse struct {
	ID string `json:"id" binding:""`
}

type GetEventResponse struct {
	Event EventResponse `json:"event"`
}

type EventResponse struct {
	Name        string `json:"name" binding:""`
	Description string `json:"description" binding:""`
}
