package seeder

import (
	"log"
	"time"

	"github.com/morelmiles/go-events/internals/models"
	"gorm.io/gorm"
)

var events = []models.Event{
	{
		ID:          1,
		Name:        "Reggae Fest",
		CoverImage:  "https://images.quicket.co.za/0437147_247_247.jpeg",
		Venue:       "Kira Town",
		EventTime:   time.Now(),
		Date:        time.Now(),
		Description: "This is the description",
		Organizer:   "Africa",
		TicketPrice: 2500,
		Category:    "fundraiser",
		Slug:        "party-after-arty",
	},
	{
		ID:          2,
		Name:        "Nyam Nya,",
		CoverImage:  "https://images.quicket.co.za/0437147_247_247.jpeg",
		Venue:       "Kampala, Uganda",
		EventTime:   time.Now(),
		Date:        time.Now(),
		Description: "Fun filled event ",
		Organizer:   "Talent Africa",
		TicketPrice: 50000,
		Category:    "fundraiser",
		Slug:        "party-after-arty",
	},
}

var users = []models.User{
	{
		ID:          1,
		FirstName:   "Luigi",
		LastName:    "Morel",
		PhoneNumber: "256739303233",
		Email:       "hi@luigimorel.co",
		Password:    "dsdsdsd",
	},
	{
		ID:          2,
		FirstName:   "Karo",
		LastName:    "Musegge",
		PhoneNumber: "256739323233",
		Email:       "hi@luigimsorel.co",
		Password:    "dsdsdsd",
	},
}

func InsertData(db *gorm.DB) {
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
