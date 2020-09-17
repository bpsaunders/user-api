package transformers

import (
	"github.com/bpsaunders/user-api/models"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const id = "id"
const firstName = "firstName"
const lastName = "lastName"
const email = "email"
const country = "country"

func TestUnitToRest(t *testing.T) {

	transformer := NewUserTransformer()

	Convey("Given I have a fully populated user db entity", t, func() {

		entity := &models.UserDao{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			Country:   country,
		}

		Convey("When I transform the entity to a REST resource", func() {

			rest := transformer.ToRest(entity)

			Convey("Then I expect all fields besides id to be mapped to the REST resource", func() {

				So(rest.FirstName, ShouldEqual, firstName)
				So(rest.LastName, ShouldEqual, lastName)
				So(rest.Email, ShouldEqual, email)
				So(rest.Country, ShouldEqual, country)
				So(rest.ID, ShouldEqual, "")
			})
		})
	})
}

func TestUnitToEntity(t *testing.T) {

	transformer := NewUserTransformer()

	Convey("Given I have a fully populated user REST resource", t, func() {

		rest := &models.User{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			Country:   country,
		}

		Convey("When I transform the REST resource to a database entity", func() {

			entity := transformer.ToEntity(rest)

			Convey("Then I expect all fields including id to be mapped to the database entity", func() {

				So(entity.FirstName, ShouldEqual, firstName)
				So(entity.LastName, ShouldEqual, lastName)
				So(entity.Email, ShouldEqual, email)
				So(entity.Country, ShouldEqual, country)
				So(entity.ID, ShouldEqual, id)
			})
		})
	})
}

func TestUnitToRestArray(t *testing.T) {

	transformer := NewUserTransformer()

	Convey("Given I have an array containing a single fully populated user db entity", t, func() {

		entity := &models.UserDao{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			Country:   country,
		}

		entityArray := []*models.UserDao{entity}

		Convey("When I transform the entity array to a REST array", func() {

			restArray := transformer.ToRestArray(&entityArray)

			Convey("Then I expect all fields including id to be mapped to the REST resource in the array", func() {

				So(len(*restArray), ShouldEqual, 1)

				rest := (*restArray)[0]
				So(rest.FirstName, ShouldEqual, firstName)
				So(rest.LastName, ShouldEqual, lastName)
				So(rest.Email, ShouldEqual, email)
				So(rest.Country, ShouldEqual, country)
				So(rest.ID, ShouldEqual, id)
			})
		})
	})
}
