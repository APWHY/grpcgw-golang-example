package pet

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/APWHY/grpcgw-golang-example/app"
)

// go:generate moq -out mock_ecomm_repository.test.go . EcommRepository

// Repository -- the interface that we use to exchange data with the database
type Repository interface {
	GetPets(context.Context) ([]Pet, error)
	InsertPet(context.Context, Pet) (Pet, error)
}

// NewPetRepository -- creates a new implementation of the Repository interface
func NewPetRepository(config *app.Config) Repository {
	retVal := &Queries{DB: app.NewDB(config.DbConnectionString)}
	err := retVal.DB.Sync2(new(Pet))
	if err != nil {
		log.Error(err, "error syncing pet")
	}
	return retVal
}
