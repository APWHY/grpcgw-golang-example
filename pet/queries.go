package pet

import (
	"context"
	"github.com/go-xorm/xorm"

	log "github.com/sirupsen/logrus"
)

const (
	LoveSchema      = "sbjpa"
	SalesOrderTable = "sales_order"
	PaidOrderStatus = 1
)

type Queries struct {
	DB *xorm.Engine
}

func (repo *Queries) GetPets(ctx context.Context) (int, error) {
	var pets []Pet

	//_, err := repo.DB.get(&pets)

	err := repo.DB.Find(&pets)

	if err != nil {
		log.Warn("Could not get all pets")
		log.Error(err)
		return 0, err
	}
	log.Info("got some pets for you: ", pets)
	return len(pets), nil
}

func (repo *Queries) InsertPet(ctx context.Context) error {
	_ = ctx

	newPet := &Pet{
		Name:    "fozz",
		Owner:   "Carlton",
		Species: "salmon",
		Sex:     "d",
		Birth:   "2001-10-11",
		//Death:   nil,
	}

	_, err := repo.DB.Insert(newPet)

	if err != nil {
		log.Warn("Could not insert pet")
		log.Error(err)
		return err
	}

	return nil
}
