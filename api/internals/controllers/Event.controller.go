package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/morelmiles/go-events/internals/models"
	"github.com/morelmiles/go-events/pkg/database"
)

func checkIfEventExists(eventId string) bool {
	var event models.Event
	database.DB.First(&event, eventId)

	return event.ID != 0
}

// GetEvents - Returns a list of all events.
func GetEvents(w http.ResponseWriter, r *http.Request) {
	var events []models.Event

	pageStr := mux.Vars(r)["page"]
	pageLimit := mux.Vars(r)["limit"]

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageLimit)
	if err != nil || pageSize < 1 {
		pageSize = 20
	}

	database.DB.Limit(pageSize).Offset((page - 1) * pageSize).Find(&events)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&events)
}

// GetHomePageEvents - Returns only 4 results
func GetHomePageEvents(w http.ResponseWriter, r *http.Request) {
	var events []models.Event

	database.DB.Limit(4).Find(&events)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&events)
}

// GetEventById - Returns a single event with the value of the ID specified.
func GetEventById(w http.ResponseWriter, r *http.Request) {
	eventId := mux.Vars(r)["id"]
	if !checkIfEventExists(eventId) {
		json.NewEncoder(w).Encode("event not found!")
		return
	}

	var event models.Event
	database.DB.First(&event, eventId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

// CreateEvent - Creates a new event
func CreateEvent(w http.ResponseWriter, r *http.Request) {

	var event models.Event
	var err error

	json.NewDecoder(r.Body).Decode(&event)

	newEvent := database.DB.Create(&event)

	err = newEvent.Error

	if err != nil {
		log.Panic(err)
	} else {
		json.NewEncoder(w).Encode(&event)
	}
}

// UpdateEventById -  Updates a single event by the ID specified
func UpdateEventById(w http.ResponseWriter, r *http.Request) {
	eventId := mux.Vars(r)["id"]
	if !checkIfEventExists(eventId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("event not found!")
		return
	}

	var event models.Event

	database.DB.First(&event, eventId)
	json.NewDecoder(r.Body).Decode(&event)
	database.DB.Save(&event)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

// DeleteEventById - Updates a single event by the ID specified.
func DeleteEventById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	eventId := mux.Vars(r)["id"]
	if !checkIfEventExists(eventId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("event not found!")
		return
	}

	var event models.Event
	database.DB.Delete(&event, eventId)
	json.NewEncoder(w).Encode(event)
}
