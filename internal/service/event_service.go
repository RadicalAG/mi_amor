package service

import (
	"context"

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
	CreateEvent(ctx context.Context, name, description, userID string) (createdID string, err error)
	GetEventByID(ctx context.Context, eventID string) (event *model.Event, err error)
}

func (e *eventService) CreateEvent(ctx context.Context, name, description, userID string) (createdID string, err error) {
	if name == "" {
		return "", internal_error.CannotBeEmptyError("name")
	}
	exist := e.userRepo.CheckIfUserExistByID(ctx, userID)
	if !exist {
		return "", internal_error.Unauthorized("")
	}
	createdID, err = e.repo.CreateEvent(ctx, &model.Event{
		Name:        name,
		Description: description,
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
