package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/morelmiles/go-events/config"

	"github.com/morelmiles/go-events/internals/models"
)

func checkIfEventExists(eventId string) bool {
	var event models.Event
	config.DB.First(&event, eventId)

	return event.ID != 0
}

// GetEvents - Returns a list of all events.
func GetEvents(w http.ResponseWriter, r *http.Request) {
	var events []models.Event

	config.DB.Find(&events)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&events)
}

// GetHomePageEvents - Returns only 4 results
func GetHomePageEvents(w http.ResponseWriter, r *http.Request) {
	var events []models.Event

	config.DB.Limit(4).Find(&events)
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
	config.DB.First(&event, eventId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

// CreateEvent - Creates a new event
func CreateEvent(w http.ResponseWriter, r *http.Request) {

	var event models.Event
	var err error

	json.NewDecoder(r.Body).Decode(&event)

	newEvent := config.DB.Create(&event)

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

	config.DB.First(&event, eventId)
	json.NewDecoder(r.Body).Decode(&event)
	config.DB.Save(&event)
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
	config.DB.Delete(&event, eventId)
	json.NewEncoder(w).Encode(event)
}
