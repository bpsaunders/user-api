package transformers

import (
	"github.com/bpsaunders/user-api/models"
)

// UserTransform provides an interface by with to transform a user resource
type UserTransform interface {
	ToRest(entity *models.UserDao) *models.User
	ToRestArray(entities *[]*models.UserDao) *[]*models.User
	ToEntity(rest *models.User) *models.UserDao
}

// UserTransformer is a concrete implementation of the UserTransform interface
type UserTransformer struct{}

// NewUserTransformer returns a new implementation of the UserTransform interface
func NewUserTransformer() UserTransform {
	return &UserTransformer{}
}

// ToRest converts a database entity to a REST resource
func (*UserTransformer) ToRest(entity *models.UserDao) *models.User {

	return &models.User{
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		Email:     entity.Email,
		Country:   entity.Country,
	}
}

// ToEntity converts a REST resource to a database entity
func (*UserTransformer) ToEntity(rest *models.User) *models.UserDao {

	return &models.UserDao{
		ID:        rest.ID,
		FirstName: rest.FirstName,
		LastName:  rest.LastName,
		Email:     rest.Email,
		Country:   rest.Country,
	}
}

// ToRestArray converts an array of database entities to an array of REST resources
func (t *UserTransformer) ToRestArray(entities *[]*models.UserDao) *[]*models.User {

	arr := make([]*models.User, 0)

	if len(*entities) > 0 {
		for _, entity := range *entities {
			rest := t.ToRest(entity)
			rest.ID = entity.ID
			arr = append(arr, rest)
		}
	}

	return &arr
}
