package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/bpsaunders/user-api/models"
	"github.com/bpsaunders/user-api/service"
	"github.com/bpsaunders/user-api/validators"
	"github.com/golang/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUnitCreateUser(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	svc := service.NewMockUserService(mockCtrl)

	handler := NewCreateUserHandler(svc)

	Convey("Given I create a user and encounter errors", t, func() {

		user := models.User{}

		var body io.Reader
		b, err := json.Marshal(user)
		if err != nil {
			t.Fatal("failed to marshal request body")
		}
		body = bytes.NewReader(b)

		req := httptest.NewRequest(http.MethodPost, "/users", body).WithContext(context.Background())
		res := httptest.NewRecorder()

		svc.EXPECT().CreateUser(&user).Return(service.Error, nil, errors.New("error when creating user"))

		handler.ServeHTTP(res, req)

		Convey("Then I expect a 500 response", func() {

			So(res.Code, ShouldEqual, http.StatusInternalServerError)
		})
	})

	Convey("Given I create a user that already exists", t, func() {

		user := models.User{}

		var body io.Reader
		b, err := json.Marshal(user)
		if err != nil {
			t.Fatal("failed to marshal request body")
		}
		body = bytes.NewReader(b)

		req := httptest.NewRequest(http.MethodPost, "/users", body).WithContext(context.Background())
		res := httptest.NewRecorder()

		svc.EXPECT().CreateUser(&user).Return(service.Conflict, []validators.ValidationError{}, nil)

		handler.ServeHTTP(res, req)

		Convey("Then I expect a 409 response", func() {

			So(res.Code, ShouldEqual, http.StatusConflict)
		})
	})

	Convey("Given I create a user with validation errors", t, func() {

		user := models.User{}

		var body io.Reader
		b, err := json.Marshal(user)
		if err != nil {
			t.Fatal("failed to marshal request body")
		}
		body = bytes.NewReader(b)

		req := httptest.NewRequest(http.MethodPost, "/users", body).WithContext(context.Background())
		res := httptest.NewRecorder()

		svc.EXPECT().CreateUser(&user).Return(service.InvalidData, []validators.ValidationError{{}}, nil)

		handler.ServeHTTP(res, req)

		Convey("Then I expect a 400 response with a body", func() {

			So(res.Code, ShouldEqual, http.StatusBadRequest)
			So(res.Body, ShouldNotBeNil)
		})
	})

	Convey("Given I create a user without errors", t, func() {

		user := models.User{}

		var body io.Reader
		b, err := json.Marshal(user)
		if err != nil {
			t.Fatal("failed to marshal request body")
		}
		body = bytes.NewReader(b)

		req := httptest.NewRequest(http.MethodPost, "/users", body).WithContext(context.Background())
		res := httptest.NewRecorder()

		svc.EXPECT().CreateUser(&user).Return(service.Success, []validators.ValidationError{}, nil)

		handler.ServeHTTP(res, req)

		Convey("Then I expect a 201 response with a body", func() {

			So(res.Code, ShouldEqual, http.StatusCreated)
			So(res.Body, ShouldNotBeNil)
		})
	})
}

func TestUnitGetAllUsers(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	svc := service.NewMockUserService(mockCtrl)

	handler := NewGetAllUsersHandler(svc)

	Convey("Given I fetch all users and encounter errors", t, func() {

		req := httptest.NewRequest(http.MethodGet, "/users", nil).WithContext(context.Background())
		res := httptest.NewRecorder()

		svc.EXPECT().GetAllUsers().Return(service.Error, nil, errors.New("error when fetching all users"))

		handler.ServeHTTP(res, req)

		Convey("Then I expect a 500 response", func() {

			So(res.Code, ShouldEqual, http.StatusInternalServerError)
		})
	})

	Convey("Given I successfully fetch all users", t, func() {

		req := httptest.NewRequest(http.MethodGet, "/users", nil).WithContext(context.Background())
		res := httptest.NewRecorder()

		svc.EXPECT().GetAllUsers().Return(service.Success, &[]*models.User{}, nil)

		handler.ServeHTTP(res, req)

		Convey("Then I expect a 200 response with a body", func() {

			So(res.Code, ShouldEqual, http.StatusOK)
			So(res.Body, ShouldNotBeNil)
		})
	})
}
