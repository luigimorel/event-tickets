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
	EventTime   time.Time `sql:"type:timestamp without time zone;not null;" json:"eventTime"`
	Date        time.Time `gorm:"type:timestamp without time zone;not null;" json:"date"`
	Description string    `gorm:"type:text;not null;" json:"description"`
	Organizer   string    `gorm:"size:255;not null;" json:"organizer"`
	TicketPrice uint64    `gorm:"size:255;not null;" json:"ticketPrice"`
	Category    string    `gorm:"size:255;not null;" json:"category"`
	Slug        string    `gorm:"size:255;not null;" json:"slug"`
}
