package dto

type CreateEventRequest struct {
	Name        string `json:"name" binding:""`
	Description string `json:"description" binding:""`
	StartDate   int64  `json:"start_date" binding:"required"`
	EndDate     int64  `json:"end_date" binding:"required"`
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
