package service

import (
	"context"
	"log"
	"time"

	"radical/red_letter/internal/internal_error"
	"radical/red_letter/internal/model"
	"radical/red_letter/internal/repository"
)

type eventService struct {
	repo repository.EventRepository
}

func NewEventService(repo repository.EventRepository) *eventService {
	return &eventService{
		repo: repo,
	}
}

type EventService interface {
	CreateEvent(ctx context.Context, name, description string, startDate, endDate time.Time) (createdID string, err error)
	GetEventByID(ctx context.Context, eventID string) (event *model.Event, err error)
}

func (e *eventService) CreateEvent(ctx context.Context, name, description string, startDate, endDate time.Time) (createdID string, err error) {
	err = e.validateRequestCreateEvent(name, startDate, endDate)
	if err != nil {
		return "", err
	}

	createdID, err = e.repo.CreateEvent(ctx, &model.Event{
		Name:        name,
		Description: description,
		StartDate:   startDate,
		EndDate:     endDate,
	})
	if err != nil {
		return "", err
	}

	return createdID, nil
}

func (e *eventService) GetEventByID(ctx context.Context, eventID string) (event *model.Event, err error) {
	if eventID == "" {
		return nil, internal_error.CannotBeEmptyError("event_id")
	}

	event, err = e.repo.GetEventByID(ctx, eventID)
	if err != nil {
		return nil, err
	}

	return event, nil

}

func (e *eventService) validateRequestCreateEvent(name string, startDate, endDate time.Time) error {
	if name == "" {
		return internal_error.CannotBeEmptyError("name")
	}

	currentTime := time.Now()
	log.Print("the current time is %v", currentTime)
	if startDate.Before(currentTime) {
		return internal_error.InvalidError("start date")
	}
	if endDate.Before(currentTime) || endDate.Before(startDate) {
		return internal_error.InvalidError("end date")
	}

	return nil
}
