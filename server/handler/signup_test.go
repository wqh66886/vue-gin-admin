package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wqh66886/vue-gin-admin/server/server/model"
	"github.com/wqh66886/vue-gin-admin/server/server/model/apperrors"
	"github.com/wqh66886/vue-gin-admin/server/server/model/mocks"
)

func TestSignUp(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Email and Password Required", func(t *testing.T) {
		mockUserService := new(mocks.MocUserService)

		mockUserService.On("SignUp", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*model.User")).Return(nil)

		rr := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		reqBody, err := json.Marshal(gin.H{
			"email": "",
		})
		assert.NoError(t, err)
		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))

		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, request)
		assert.Equal(t, 400, rr.Code)
		mockUserService.AssertCalled(t, "SignUp")
	})

	t.Run("Invalid email", func(t *testing.T) {
		mockUserService := new(mocks.MocUserService)

		mockUserService.On("SignUp", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*model.User")).Return(nil)

		rr := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		reqBody, err := json.Marshal(gin.H{
			"email":    "bob@bob",
			"password": "supersecret1234",
		})
		assert.NoError(t, err)
		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))

		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, request)
		assert.Equal(t, 400, rr.Code)
		mockUserService.AssertCalled(t, "SignUp")
	})

	t.Run("password too short ", func(t *testing.T) {
		mockUserService := new(mocks.MocUserService)

		mockUserService.On("SignUp", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*model.User")).Return(nil)

		rr := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		reqBody, err := json.Marshal(gin.H{
			"email":    "bob@bob.com",
			"password": "super",
		})
		assert.NoError(t, err)
		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))

		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, request)
		assert.Equal(t, 400, rr.Code)
		mockUserService.AssertCalled(t, "SignUp")
	})

	t.Run("password too long ", func(t *testing.T) {
		mockUserService := new(mocks.MocUserService)

		mockUserService.On("SignUp", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*model.User")).Return(nil)

		rr := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		reqBody, err := json.Marshal(gin.H{
			"email":    "bob@bob.com",
			"password": "superadfdfadfadfadfadfadfadfadfadfadfadfadfadfadfadfadfadfadfddfadfdfadfa",
		})
		assert.NoError(t, err)
		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))

		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, request)
		assert.Equal(t, 400, rr.Code)
		mockUserService.AssertCalled(t, "SignUp")
	})

	t.Run("Error returnrf from UserService", func(t *testing.T) {
		u := &model.User{
			Email:    "bob@bob.com",
			Password: "superpassword",
		}

		mockUserService := new(mocks.MocUserService)
		mockUserService.On("SignUp", mock.AnythingOfType("*gin.Context"), u).Return(apperrors.NewConflict("User Already Exits", u.Email))

		rr := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		reqBody, err := json.Marshal(gin.H{
			"emai":     u.Email,
			"password": u.Password,
		})
		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)
		assert.Equal(t, 400, rr.Code)
		mockUserService.AssertExpectations(t)
	})
}
