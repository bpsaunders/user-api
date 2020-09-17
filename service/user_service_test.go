package service

import (
	"errors"
	"github.com/bpsaunders/user-api/db"
	"github.com/bpsaunders/user-api/models"
	"github.com/bpsaunders/user-api/transformers"
	"github.com/bpsaunders/user-api/validators"
	"github.com/golang/mock/gomock"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const email = "email"
const id = "id"

func TestUnitCreateUser(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	transformer := transformers.NewMockUserTransform(mockCtrl)
	validator := validators.NewMockUserValidate(mockCtrl)
	client := db.NewMockClient(mockCtrl)

	svc := &UserServiceImpl{
		transformer: transformer,
		validator:   validator,
		db:          client,
	}

	rest := models.User{
		Email: email,
	}

	Convey("Given I attempt to create a rest resource with validation errors", t, func() {

		validationErrors := []validators.ValidationError{{}}

		validator.EXPECT().Validate(&rest).Return(validationErrors)

		responseType, validationErrs, err := svc.CreateUser(&rest)

		Convey("Then I expect an 'invalid-data' response type", func() {

			So(responseType, ShouldEqual, InvalidData)

			Convey("And validation errors should be returned", func() {

				So(validationErrs, ShouldResemble, validationErrors)

				Convey("And no errors should be present", func() {

					So(err, ShouldBeNil)
				})
			})
		})
	})

	Convey("Given I attempt to create a rest resource without validation errors", t, func() {

		var validationErrors []validators.ValidationError

		validator.EXPECT().Validate(&rest).Return(validationErrors)

		Convey("But the user already exists", func() {

			client.EXPECT().UserExistsWithEmail(email).Return(true, nil)

			responseType, validationErrs, err := svc.CreateUser(&rest)

			Convey("Then I expect a 'conflict' response type", func() {

				So(responseType, ShouldEqual, Conflict)

				Convey("But validation errors should be empty", func() {

					So(len(validationErrs), ShouldEqual, 0)

					Convey("And no errors should be present", func() {

						So(err, ShouldBeNil)
					})
				})
			})
		})
	})

	Convey("Given I attempt to create a rest resource without validation errors", t, func() {

		var validationErrors []validators.ValidationError

		validator.EXPECT().Validate(&rest).Return(validationErrors)

		Convey("But I am returned an error when checking if the user exists", func() {

			dbErr := errors.New("error when checking if a user exists")

			client.EXPECT().UserExistsWithEmail(email).Return(false, dbErr)

			responseType, validationErrs, err := svc.CreateUser(&rest)

			Convey("Then I expect an 'error' response type", func() {

				So(responseType, ShouldEqual, Error)

				Convey("But validation errors should be empty", func() {

					So(len(validationErrs), ShouldEqual, 0)

					Convey("And errors should be returned", func() {

						So(err, ShouldEqual, dbErr)
					})
				})
			})
		})
	})

	Convey("Given I attempt to create a rest resource without validation errors", t, func() {

		var validationErrors []validators.ValidationError

		validator.EXPECT().Validate(&rest).Return(validationErrors)

		Convey("And the user doesn't exist", func() {

			client.EXPECT().UserExistsWithEmail(email).Return(false, nil)

			Convey("Then the REST resource is transformed to a db entity", func() {

				entity := models.UserDao{}

				transformer.EXPECT().ToEntity(&rest).Return(&entity)

				Convey("But if there's an error when saving the user to the db", func() {

					dbErr := errors.New("error saving the user to the db")

					client.EXPECT().CreateUser(&entity).Return(dbErr)

					responseType, validationErrs, err := svc.CreateUser(&rest)

					Convey("Then I expect an 'error' response type", func() {

						So(responseType, ShouldEqual, Error)

						Convey("But validation errors should be empty", func() {

							So(len(validationErrs), ShouldEqual, 0)

							Convey("And errors should be returned", func() {

								So(err, ShouldEqual, dbErr)
							})
						})
					})
				})
			})
		})
	})

	Convey("Given I attempt to create a rest resource without validation errors", t, func() {

		var validationErrors []validators.ValidationError

		validator.EXPECT().Validate(&rest).Return(validationErrors)

		Convey("And the user doesn't exist", func() {

			client.EXPECT().UserExistsWithEmail(email).Return(false, nil)

			Convey("Then the REST resource is transformed to a db entity", func() {

				entity := models.UserDao{}

				transformer.EXPECT().ToEntity(&rest).Return(&entity)

				Convey("And if there's an error when saving the user to the db", func() {

					client.EXPECT().CreateUser(&entity).Return(nil)

					responseType, validationErrs, err := svc.CreateUser(&rest)

					Convey("Then I expect a 'success' response type", func() {

						So(responseType, ShouldEqual, Success)

						Convey("And validation errors should be empty", func() {

							So(len(validationErrs), ShouldEqual, 0)

							Convey("And no errors should be returned", func() {

								So(err, ShouldBeNil)

								Convey("And an id should be returned in the REST resource", func() {

									So(rest.ID, ShouldNotBeBlank)
								})
							})
						})
					})
				})
			})
		})
	})
}

func TestUnitGetUser(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	transformer := transformers.NewMockUserTransform(mockCtrl)
	validator := validators.NewMockUserValidate(mockCtrl)
	client := db.NewMockClient(mockCtrl)

	svc := &UserServiceImpl{
		transformer: transformer,
		validator:   validator,
		db:          client,
	}

	Convey("Given I encounter errors when fetching a user", t, func() {

		dbErr := errors.New("error when fetching a user")

		client.EXPECT().GetUser(id).Return(nil, dbErr)

		responseType, user, err := svc.GetUser(id)

		Convey("Then I expect an 'error' response type", func() {

			So(responseType, ShouldEqual, Error)

			Convey("And user should be nil", func() {

				So(user, ShouldBeNil)

				Convey("And errors should be returned", func() {

					So(err, ShouldEqual, dbErr)
				})
			})
		})
	})

	Convey("Given I don't find the user I'm fetching", t, func() {

		client.EXPECT().GetUser(id).Return(nil, nil)

		responseType, user, err := svc.GetUser(id)

		Convey("Then I expect a 'not-found' response type", func() {

			So(responseType, ShouldEqual, NotFound)

			Convey("And user should be nil", func() {

				So(user, ShouldBeNil)

				Convey("And errors should be nil", func() {

					So(err, ShouldBeNil)
				})
			})
		})
	})

	Convey("Given I find the user I'm fetching", t, func() {

		entity := models.UserDao{}

		client.EXPECT().GetUser(id).Return(&entity, nil)

		rest := models.User{}

		transformer.EXPECT().ToRest(&entity).Return(&rest)

		responseType, user, err := svc.GetUser(id)

		Convey("Then I expect a 'success' response type", func() {

			So(responseType, ShouldEqual, Success)

			Convey("And user should be returned", func() {

				So(user, ShouldEqual, &rest)

				Convey("And errors should be nil", func() {

					So(err, ShouldBeNil)
				})
			})
		})
	})
}

func TestUnitGetAllUsers(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	transformer := transformers.NewMockUserTransform(mockCtrl)
	validator := validators.NewMockUserValidate(mockCtrl)
	client := db.NewMockClient(mockCtrl)

	svc := &UserServiceImpl{
		transformer: transformer,
		validator:   validator,
		db:          client,
	}

	Convey("Given I encounter errors when fetching all users", t, func() {

		dbErr := errors.New("error when fetching all users")

		client.EXPECT().GetAllUsers().Return(nil, dbErr)

		responseType, users, err := svc.GetAllUsers()

		Convey("Then I expect an 'error' response type", func() {

			So(responseType, ShouldEqual, Error)

			Convey("And users should be nil", func() {

				So(users, ShouldBeNil)

				Convey("And errors should be returned", func() {

					So(err, ShouldEqual, dbErr)
				})
			})
		})
	})

	Convey("Given I successfully fetch all users", t, func() {

		entities := make([]*models.UserDao, 0)

		client.EXPECT().GetAllUsers().Return(&entities, nil)

		restResources := make([]*models.User, 0)

		transformer.EXPECT().ToRestArray(&entities).Return(&restResources)

		responseType, users, err := svc.GetAllUsers()

		Convey("Then I expect a 'success' response type", func() {

			So(responseType, ShouldEqual, Success)

			Convey("And users should be returned", func() {

				So(users, ShouldEqual, &restResources)

				Convey("And errors should not be returned", func() {

					So(err, ShouldBeNil)
				})
			})
		})
	})
}

func TestUnitShutdown(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	transformer := transformers.NewMockUserTransform(mockCtrl)
	validator := validators.NewMockUserValidate(mockCtrl)
	client := db.NewMockClient(mockCtrl)

	svc := &UserServiceImpl{
		transformer: transformer,
		validator:   validator,
		db:          client,
	}

	Convey("Verify db is shutdown on application shutdown", t, func() {

		client.EXPECT().Shutdown().Times(1)

		svc.Shutdown()
	})
}
