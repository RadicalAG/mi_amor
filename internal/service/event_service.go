package service

import (
	"context"
	"time"

	"radical/red_letter/internal/internal_error"
	"radical/red_letter/internal/model"
	"radical/red_letter/internal/repository"
)

type eventService struct {
	repo     repository.EventRepository
	userRepo repository.UserRepository
}

func NewEventService(repo repository.EventRepository, userRepo repository.UserRepository) *eventService {
	return &eventService{
		repo:     repo,
		userRepo: userRepo,
	}
}

type EventService interface {
	CreateEvent(ctx context.Context, name, description, userID string, startDate, endDate int64) (createdID string, err error)
	GetEventByID(ctx context.Context, eventID string) (event *model.Event, err error)
}

func (e *eventService) CreateEvent(ctx context.Context, name, description, userID string, startDate, endDate int64) (createdID string, err error) {
	err = e.validateRequestCreateEvent(name, startDate, endDate)
	if err != nil {
		return "", err
	}

	exist := e.userRepo.CheckIfUserExistByID(ctx, userID)
	if !exist {
		return "", internal_error.Unauthorized("")
	}

	createdID, err = e.repo.CreateEvent(ctx, &model.Event{
		Name:        name,
		Description: description,
		StartDate:   startDate,
		EndDate:     endDate,
	}, userID)

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

func (e *eventService) validateRequestCreateEvent(name string, startDate, endDate int64) error {
	if name == "" {
		return internal_error.CannotBeEmptyError("name")
	}

	current := time.Now()
	start := time.Unix(startDate, 0)
	end := time.Unix(endDate, 0)
	if start.Before(current) {
		return internal_error.InvalidError("start date")
	}
	if end.Before(current) || end.Before(start) {
		return internal_error.InvalidError("end date")
	}

	return nil
}
