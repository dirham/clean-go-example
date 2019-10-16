package clean

import (
	"clean-example/models"
	"time"

	"github.com/globalsign/mgo"
)

// defined provider here
// you can sparate driver into different stuct like this :
type mgoRepository struct {
	MgoDatabase *mgo.Database
}

// this will throw an error(s) if you not implement all function in repository interface
func NewCleanRepository(DB *mgo.Database) Repository {
	return &mgoRepository{MgoDatabase: DB}
}

// implementation of interface
func (mr *mgoRepository) FindByID(id int) (*models.Clean, error) {
	// do logic here
	return &models.Clean{ID: id, Optional: time.Now()}, nil
}
