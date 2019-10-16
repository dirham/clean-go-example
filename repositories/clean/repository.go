package clean

import "clean-example/models"

// Repository interface hold communication logic between app and outside sevices. i.g: DB communication etc..
type Repository interface {
	FindByID(id int) (*models.Clean, error)
}
