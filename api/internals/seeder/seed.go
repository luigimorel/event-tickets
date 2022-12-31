package seeder

import (
	"log"
	"time"

	"github.com/morelmiles/go-events/internals/models"
	"gorm.io/gorm"
)

var events = []models.Event{
	{
		Name:        "Reggae Fest",
		CoverImage:  "https://images.quicket.co.za/0437147_247_247.jpeg",
		Venue:       "Online",
		EventTime:   time.Now(),
		Date:        time.Now(),
		Description: "This is the description",
		Organizer:   "Africa",
		TicketPrice: 2500,
		Category:    "fundraiser",
		Slug:        "/party-after-arty",
	},
	{
		Name:        "NyaNya",
		CoverImage:  "https://images.quicket.co.za/0351438_247_247.jpeg",
		Venue:       "Kampala, Uganda",
		EventTime:   time.Now(),
		Date:        time.Now(),
		Description: "Fun filled event",
		Organizer:   "Talent Africa",
		TicketPrice: 50000,
		Category:    "fundraiser",
		Slug:        "/nyanya",
	},
	{
		Name:        "Lira Festival",
		CoverImage:  "https://images.quicket.co.za/0437147_247_247.jpeg",
		Venue:       "Kampala, Uganda",
		EventTime:   time.Now(),
		Date:        time.Now(),
		Description: "Fun filled event",
		Organizer:   "Talent Africa",
		TicketPrice: 50000,
		Category:    "fundraiser",
		Slug:        "/party-after-arty",
	},
	{
		Name:        "Hotels Uganda",
		CoverImage:  "https://images.quicket.co.za/0437121_247_247.jpeg",
		Venue:       "Kampala, Uganda",
		EventTime:   time.Now(),
		Date:        time.Now(),
		Description: "Fun filled event",
		Organizer:   "Talent Africa",
		TicketPrice: 50000,
		Category:    "fundraiser",
		Slug:        "/party-after-arty",
	},
}

var users = []models.User{
	{
		Name:        "Luigi Morel",
		PhoneNumber: "256739303233",
		Email:       "hi@luigimorel.co",
		Password:    "dsdsdsd",
	},
	{
		Name:        "Luigi Morel",
		PhoneNumber: "256739323233",
		Email:       "hi@luigimsorel.co",
		Password:    "dsdsdsd",
	},
}

func InsertSeedData(db *gorm.DB) {
	err := db.Migrator().DropTable(&models.Event{})
	if err != nil {
		log.Fatalf("cannot drop events table")
	}

	err = db.Migrator().DropTable(&models.User{})
	if err != nil {
		log.Fatalf("cannot drop users table")
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Event{})

	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

		err = db.Debug().Model(&models.Event{}).Create(&events[i]).Error
		if err != nil {
			log.Fatalf("cannot seed events table: %v", err)
		}
	}
}
