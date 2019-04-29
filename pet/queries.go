package pet

import (
	"context"
	"github.com/asaskevich/govalidator"
	"github.com/go-xorm/xorm"

	log "github.com/sirupsen/logrus"
)

// Queries -- holds the instance of the DB that we interface with
type Queries struct {
	DB *xorm.Engine
}

// GetPets -- gets all the pets stored in our db
func (repo *Queries) GetPets(ctx context.Context) ([]Pet, error) {
	var pets []Pet

	//_, err := repo.DB.get(&pets)

	err := repo.DB.Find(&pets)

	if err != nil {
		log.Warn("Could not get all pets")
		log.Error(err)
		return nil, err
	}
	log.Info("got some pets for you: ", pets)
	return pets, nil
}

// InsertPet -- inserts a pet into our db
func (repo *Queries) InsertPet(ctx context.Context, pet Pet) (Pet, error) {
	_ = ctx
	pet.Clean()
	valid, err := govalidator.ValidateStruct(pet)

	if !valid {
		log.Error(err)
		return pet, err
	}
	_, err = repo.DB.Insert(pet)

	if err != nil {
		log.Warn("Could not insert pet")
		log.Error(err)
		return pet, err
	}

	return pet, nil
}
