package pet

import (
	"context"

	log "github.com/sirupsen/logrus"

	"gitlab.com/loveplus/data-ingest/app"
)

// go:generate moq -out mock_ecomm_repository.test.go . EcommRepository
type Repository interface {
	GetPets(ctx context.Context) (int, error)
	InsertPet(ctx context.Context) error
}

func NewPetRepository(config *app.Config) Repository {
	log.Info("Initializing new Pet Repository instance")
	retVal := &Queries{DB: app.NewDB(config.EcommCloudDbConnectionString)}
	log.Info("Instance initialised, trying sync now")
	err := retVal.DB.Sync2(new(Pet))
	if err != nil {
		log.Error(err, "error syncing pet")
	}
	return retVal
}
