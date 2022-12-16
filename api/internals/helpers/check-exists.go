package helpers

import (
	"github.com/morelmiles/go-events/config"
)

type Model struct {
	Id int
}

func CheckIfIdExists(model Model, id string) bool {
	config.DB.First(&model, id)

	return model.Id != 0
}
