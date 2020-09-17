package service

import (
	"github.com/bpsaunders/user-api/config"
	"github.com/bpsaunders/user-api/db"
	"github.com/bpsaunders/user-api/models"
	"github.com/bpsaunders/user-api/transformers"
	"github.com/bpsaunders/user-api/validators"
	"github.com/hashicorp/go-uuid"
)

// UserService provides an interface by which to interact with a User resource
type UserService interface {
	CreateUser(rest *models.User) (ResponseType, []validators.ValidationError, error)
	GetUser(id string) (ResponseType, *models.User, error)
	GetAllUsers() (ResponseType, *[]*models.User, error)
	Shutdown()
}

// UserServiceImpl provides a concrete implementation of the UserService interface
type UserServiceImpl struct {
	transformer transformers.UserTransform
	validator   validators.UserValidate
	db          db.Client
}

// NewUserService returns a new concrete implementation of the UserService interface
func NewUserService(cfg *config.Config) UserService {
	return &UserServiceImpl{
		transformer: transformers.NewUserTransformer(),
		validator:   validators.NewUserValidator(),
		db:          db.NewDatabaseClient(cfg),
	}
}

// CreateUser validates and creates a user resource
func (service *UserServiceImpl) CreateUser(rest *models.User) (ResponseType, []validators.ValidationError, error) {

	// validate the resource first
	validationErrors := service.validator.Validate(rest)
	if len(validationErrors) > 0 {
		return InvalidData, validationErrors, nil
	}

	userExists, err := service.db.UserExistsWithEmail(rest.Email)
	if err != nil {
		return Error, validationErrors, err
	}
	if userExists {
		return Conflict, validationErrors, err
	}

	// no validation errors; generate a unique id and stamp it on the rest resource
	id, err := uuid.GenerateUUID()
	if err != nil {
		return Error, validationErrors, err
	}
	rest.ID = id

	// transformer the rest resource to a DAO entity
	entity := service.transformer.ToEntity(rest)

	// save entity to the db
	err = service.db.CreateUser(entity)
	if err != nil {
		return Error, validationErrors, err
	}

	return Success, validationErrors, err
}

// GetUser fetches an individual user according to an id
func (service *UserServiceImpl) GetUser(id string) (ResponseType, *models.User, error) {

	// fetch the db entity
	entity, err := service.db.GetUser(id)

	if err != nil {
		return Error, nil, err
	}

	// if nil entity, no results were found so cascade that up to the handler
	if entity == nil {
		return NotFound, nil, nil
	}

	return Success, service.transformer.ToRest(entity), err
}

// GetAllUsers returns an array of all users
func (service *UserServiceImpl) GetAllUsers() (ResponseType, *[]*models.User, error) {

	// fetch the db entities
	entities, err := service.db.GetAllUsers()

	if err != nil {
		return Error, nil, err
	}

	return Success, service.transformer.ToRestArray(entities), err
}

// Shutdown provides functionality to clean up resources on application shutdown
func (service *UserServiceImpl) Shutdown() {

	service.db.Shutdown()
}
