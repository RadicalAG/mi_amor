package dto

type CreateEventRequest struct {
	Name        string `json:"name" binding:""`
	Description string `json:"description" binding:""`
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
