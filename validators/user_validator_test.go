package validators

import (
	"github.com/bpsaunders/user-api/models"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUnitValidate(t *testing.T) {

	validator := NewUserValidator()

	Convey("Given I validate a valid user", t, func() {

		validationErrors := validator.Validate(createValidUser())

		Convey("Then I expect no errors", func() {

			So(len(validationErrors), ShouldEqual, 0)
		})
	})

	Convey("Given I validate a user without a first name", t, func() {

		user := createValidUser()
		user.FirstName = ""
		validationErrors := validator.Validate(user)

		Convey("Then I expect 1 error", func() {

			So(len(validationErrors), ShouldEqual, 1)

			Convey("For first name", func() {

				So(validationErrors[0].Field, ShouldEqual, jsonFieldPrefix+firstNameField)

				Convey("Stating the field is mandatory", func() {

					So(validationErrors[0].Error, ShouldEqual, mandatoryElementMissing)
				})
			})
		})
	})

	Convey("Given I validate a user with a single character first name", t, func() {

		user := createValidUser()
		user.FirstName = "a"
		validationErrors := validator.Validate(user)

		Convey("Then I expect 1 error", func() {

			So(len(validationErrors), ShouldEqual, 1)

			Convey("For first name", func() {

				So(validationErrors[0].Field, ShouldEqual, jsonFieldPrefix+firstNameField)

				Convey("Stating the field is invalid length", func() {

					So(validationErrors[0].Error, ShouldEqual, invalidLength)

					Convey("With params stating the minimum length is 2", func() {

						So(validationErrors[0].Params[minChars], ShouldEqual, 2)
					})
				})
			})
		})
	})

	Convey("Given I validate a user with a 31 character first name", t, func() {

		user := createValidUser()
		user.FirstName = strings.Repeat("a", 31)
		validationErrors := validator.Validate(user)

		Convey("Then I expect 1 error", func() {

			So(len(validationErrors), ShouldEqual, 1)

			Convey("For first name", func() {

				So(validationErrors[0].Field, ShouldEqual, jsonFieldPrefix+firstNameField)

				Convey("Stating the field is invalid length", func() {

					So(validationErrors[0].Error, ShouldEqual, invalidLength)

					Convey("With params stating the maximum length is 30", func() {

						So(validationErrors[0].Params[maxChars], ShouldEqual, 30)
					})
				})
			})
		})
	})

	Convey("Given I validate a user with invalid chars in their first name", t, func() {

		user := createValidUser()
		user.FirstName = "name!"
		validationErrors := validator.Validate(user)

		Convey("Then I expect 1 error", func() {

			So(len(validationErrors), ShouldEqual, 1)

			Convey("For first name", func() {

				So(validationErrors[0].Field, ShouldEqual, jsonFieldPrefix+firstNameField)

				Convey("Stating the field contains invalid characters", func() {

					So(validationErrors[0].Error, ShouldEqual, invalidChars)
				})
			})
		})
	})

	Convey("Given I validate a user without a last name", t, func() {

		user := createValidUser()
		user.LastName = ""
		validationErrors := validator.Validate(user)

		Convey("Then I expect 1 error", func() {

			So(len(validationErrors), ShouldEqual, 1)

			Convey("For last name", func() {

				So(validationErrors[0].Field, ShouldEqual, jsonFieldPrefix+lastNameField)

				Convey("Stating the field is mandatory", func() {

					So(validationErrors[0].Error, ShouldEqual, mandatoryElementMissing)
				})
			})
		})
	})

	Convey("Given I validate a user with a single character last name", t, func() {

		user := createValidUser()
		user.LastName = "a"
		validationErrors := validator.Validate(user)

		Convey("Then I expect 1 error", func() {

			So(len(validationErrors), ShouldEqual, 1)

			Convey("For last name", func() {

				So(validationErrors[0].Field, ShouldEqual, jsonFieldPrefix+lastNameField)

				Convey("Stating the field is invalid length", func() {

					So(validationErrors[0].Error, ShouldEqual, invalidLength)

					Convey("With params stating the minimum length is 2", func() {

						So(validationErrors[0].Params[minChars], ShouldEqual, 2)
					})
				})
			})
		})
	})

	Convey("Given I validate a user with a 31 character last name", t, func() {

		user := createValidUser()
		user.LastName = strings.Repeat("a", 31)
		validationErrors := validator.Validate(user)

		Convey("Then I expect 1 error", func() {

			So(len(validationErrors), ShouldEqual, 1)

			Convey("For last name", func() {

				So(validationErrors[0].Field, ShouldEqual, jsonFieldPrefix+lastNameField)

				Convey("Stating the field is invalid length", func() {

					So(validationErrors[0].Error, ShouldEqual, invalidLength)

					Convey("With params stating the maximum length is 30", func() {

						So(validationErrors[0].Params[maxChars], ShouldEqual, 30)
					})
				})
			})
		})
	})

	Convey("Given I validate a user with invalid chars in their last name", t, func() {

		user := createValidUser()
		user.LastName = "name!"
		validationErrors := validator.Validate(user)

		Convey("Then I expect 1 error", func() {

			So(len(validationErrors), ShouldEqual, 1)

			Convey("For last name", func() {

				So(validationErrors[0].Field, ShouldEqual, jsonFieldPrefix+lastNameField)

				Convey("Stating the field contains invalid characters", func() {

					So(validationErrors[0].Error, ShouldEqual, invalidChars)
				})
			})
		})
	})

	Convey("Given I validate a user without an email", t, func() {

		user := createValidUser()
		user.Email = ""
		validationErrors := validator.Validate(user)

		Convey("Then I expect 1 error", func() {

			So(len(validationErrors), ShouldEqual, 1)

			Convey("For email", func() {

				So(validationErrors[0].Field, ShouldEqual, jsonFieldPrefix+emailField)

				Convey("Stating the field is mandatory", func() {

					So(validationErrors[0].Error, ShouldEqual, mandatoryElementMissing)
				})
			})
		})
	})

	Convey("Given I validate a user with a 121 character email", t, func() {

		user := createValidUser()
		user.Email = strings.Repeat("a", 121)
		validationErrors := validator.Validate(user)

		Convey("Then I expect 1 error", func() {

			So(len(validationErrors), ShouldEqual, 1)

			Convey("For email", func() {

				So(validationErrors[0].Field, ShouldEqual, jsonFieldPrefix+emailField)

				Convey("Stating the field is invalid length", func() {

					So(validationErrors[0].Error, ShouldEqual, invalidLength)

					Convey("With params stating the maximum length is 120", func() {

						So(validationErrors[0].Params[maxChars], ShouldEqual, 120)
					})
				})
			})
		})
	})

	Convey("Given I validate a user with a malformed email", t, func() {

		user := createValidUser()
		user.Email = "invalidEmailAddress"
		validationErrors := validator.Validate(user)

		Convey("Then I expect 1 error", func() {

			So(len(validationErrors), ShouldEqual, 1)

			Convey("For email", func() {

				So(validationErrors[0].Field, ShouldEqual, jsonFieldPrefix+emailField)

				Convey("Stating the field is an invalid format", func() {

					So(validationErrors[0].Error, ShouldEqual, invalidFormat)
				})
			})
		})
	})

	Convey("Given I validate a user without a country", t, func() {

		user := createValidUser()
		user.Country = ""
		validationErrors := validator.Validate(user)

		Convey("Then I expect 1 error", func() {

			So(len(validationErrors), ShouldEqual, 1)

			Convey("For country", func() {

				So(validationErrors[0].Field, ShouldEqual, jsonFieldPrefix+countryField)

				Convey("Stating the field is mandatory", func() {

					So(validationErrors[0].Error, ShouldEqual, mandatoryElementMissing)
				})
			})
		})
	})

	Convey("Given I validate a user with an invalid country code", t, func() {

		user := createValidUser()
		user.Country = "notACountryCode"
		validationErrors := validator.Validate(user)

		Convey("Then I expect 1 error", func() {

			So(len(validationErrors), ShouldEqual, 1)

			Convey("For country", func() {

				So(validationErrors[0].Field, ShouldEqual, jsonFieldPrefix+countryField)

				Convey("Stating the field is an invalid country code", func() {

					So(validationErrors[0].Error, ShouldEqual, invalidCountryCode)
				})
			})
		})
	})
}

func createValidUser() *models.User {

	return &models.User{
		FirstName: "firstName",
		LastName:  "lastName",
		Email:     "user@mail.com",
		Country:   "GB",
	}
}
