package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ID          uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name        string    `gorm:"size:255;not null;" json:"name"`
	CoverImage  string    `gorm:"size:255;not null;" json:"coverImage"`
	Venue       string    `gorm:"size:255;not null;" json:"venue"`
	EventTime   time.Time `gorm:"type:text;not null;" json:"eventTime"`
	Date        time.Time `gorm:"type:text;not null;" json:"date"`
	Description string    `gorm:"type:text;not null;" json:"description"`
	Organizer   string    `gorm:"size:255;not null;" json:"organizer"`
	TicketPrice uint64    `gorm:"size:255;not null;" json:"ticketPrice"`
}
