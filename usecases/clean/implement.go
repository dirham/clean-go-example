package clean

import (
	"clean-example/models"
	"clean-example/repositories/clean"
)

type MgoUsecase struct {
	mongoRepo clean.Repository
}

// NewMgoUsecase is constructor, recieve implementation of repository
func NewMgoUsecase(mr clean.Repository) *MgoUsecase {
	return &MgoUsecase{mongoRepo: mr}
}

// this method call implementation of Repository
func (mu *MgoUsecase) FindByID(id int) (*models.Clean, error) {
	return mu.mongoRepo.FindByID(id)
}
