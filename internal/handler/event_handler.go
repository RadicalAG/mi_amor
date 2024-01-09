package handler

import (
	"net/http"
	"radical/red_letter/internal/api_error"
	"radical/red_letter/internal/dto"
	"radical/red_letter/internal/middleware"
	"radical/red_letter/internal/utils"

	"radical/red_letter/internal/service"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	service        service.EventService
	authmiddleware middleware.AuthMiddleware
}

func NewEventHandler(service service.EventService, authmiddleware middleware.AuthMiddleware) *EventHandler {
	return &EventHandler{
		service:        service,
		authmiddleware: authmiddleware,
	}
}

func (t *EventHandler) RegisterHandler(r *gin.Engine) *gin.Engine {
	r.POST("/event", t.authmiddleware.TokenAuthorization(), t.CreateEvent)
	r.GET("/event/:id", t.GetEventByID)
	return r
}

func (t *EventHandler) CreateEvent(c *gin.Context) {
	var req dto.CreateEventRequest
	claim, err := utils.GetClaimsFromToken(c)
	if err != nil {
		c.Error(err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(api_error.NewApiError(http.StatusBadRequest, "invalid body"))
		return
	}

	createdID, err := t.service.CreateEvent(c, req.Name, req.Description, claim.ID)
	if err != nil {
		c.Error(api_error.FromError(err))
		return
	}
	res := dto.CreateEventResponse{
		ID: createdID,
	}

	c.JSON(http.StatusCreated, JsonSuccessFormater("Event Created", res))
}

func (t *EventHandler) GetEventByID(c *gin.Context) {
	id := c.Param("id")
	event, err := t.service.GetEventByID(c, id)
	if err != nil {
		c.Error(api_error.FromError(err))
		return
	}
	res := dto.GetEventResponse{
		Event: dto.EventResponse{
			Name:        event.Name,
			Description: event.Description,
		},
	}

	c.JSON(http.StatusOK, JsonSuccessFormater("Get Event Successfully", res))

}
