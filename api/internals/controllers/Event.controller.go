package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/morelmiles/go-events/internals/errors"
	"github.com/morelmiles/go-events/internals/helpers"
	"github.com/morelmiles/go-events/internals/models"
	"github.com/morelmiles/go-events/pkg/database"
)

func GetEvents(w http.ResponseWriter, r *http.Request) {
	var events []models.Event
	if err := database.DB.Find(&events).Error; err != nil {
		apiErr := errors.NewDatabaseError("Failed to fetch events")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}
	helpers.JSON(w, http.StatusOK, &events)
}

func GetHomePageEvents(w http.ResponseWriter, r *http.Request) {
	var events []models.Event
	if err := database.DB.Limit(3).Find(&events).Error; err != nil {
		apiErr := errors.NewDatabaseError("Failed to fetch homepage events")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}
	helpers.JSON(w, http.StatusOK, &events)
}

func GetEventById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		apiErr := errors.NewBadRequestError("Invalid event ID format")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}

	var event models.Event
	if err := database.DB.First(&event, id).Error; err != nil {
		apiErr := errors.NewNotFoundError("Event not found")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}
	helpers.JSON(w, http.StatusOK, event)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event

	if err := helpers.DecodeJSONBody(w, r, &event); err != nil {
		apiErr := errors.NewValidationError("Invalid request body")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}

	if err := database.DB.Create(&event).Error; err != nil {
		apiErr := errors.NewDatabaseError("Failed to create event")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}

	helpers.JSON(w, http.StatusCreated, &event)
}

func DeleteEventById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		apiErr := errors.NewBadRequestError("Invalid event ID format")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}

	var event models.Event
	if err := database.DB.First(&event, id).Error; err != nil {
		apiErr := errors.NewNotFoundError("Event not found")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}

	if err := database.DB.Delete(&event).Error; err != nil {
		apiErr := errors.NewDatabaseError("Failed to delete event")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}

	helpers.JSON(w, http.StatusOK, event)
}

func UpdateEventById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		apiErr := errors.NewBadRequestError("Invalid event ID format")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}

	var event models.Event
	if err := database.DB.First(&event, id).Error; err != nil {
		apiErr := errors.NewNotFoundError("Event not found")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}

	if err := helpers.DecodeJSONBody(w, r, &event); err != nil {
		apiErr := errors.NewValidationError("Invalid request body")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}

	if err := database.DB.Save(&event).Error; err != nil {
		apiErr := errors.NewDatabaseError("Failed to update event")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}

	helpers.JSON(w, http.StatusOK, event)
}
