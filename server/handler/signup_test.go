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

	t.Run("Successful Token Creation", func(t *testing.T) {
		u := &model.User{
			Email:    "bob@bob.com",
			Password: "123456789",
		}

		mockTokenResp := &model.TokenPair{
			IDToken:      "idToken",
			RefreshToken: "refreshToken",
		}

		mockTokenService := new(mocks.MockTokenService)
		mockUserService := new(mocks.MocUserService)

		mockUserService.On("SignUp", mock.AnythingOfType("*gin.Context"), u).Return(nil)
		mockTokenService.On("NewPairFromUser", mock.AnythingOfType("*gin.Context"), u, "").Return(mockTokenResp, nil)

		rr := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			R:            router,
			UserService:  mockUserService,
			TokenService: mockTokenService,
		})

		reqBody, err := json.Marshal(gin.H{
			"email":    u.Email,
			"password": u.Password,
		})

		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))

		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal(gin.H{
			"tokens": mockTokenResp,
		})

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		mockUserService.AssertExpectations(t)
		mockTokenService.AssertExpectations(t)
	})

	t.Run("Failed Token Creation", func(t *testing.T) {
		u := &model.User{
			Email:    "bb@bb.com",
			Password: "123456789",
		}

		mockErrorResponse := apperrors.NewInternal()

		mockUserService := new(mocks.MocUserService)
		mockTokenService := new(mocks.MockTokenService)

		mockUserService.On("SignUp", mock.AnythingOfType("*gin.Context"), u).Return(nil)
		mockTokenService.On("NewPairFromUser", mock.AnythingOfType("*gin.Context"), u, "").Return(nil, mockErrorResponse)

		rr := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			R:            router,
			UserService:  mockUserService,
			TokenService: mockTokenService,
		})

		reqBody, err := json.Marshal(gin.H{
			"email":    u.Email,
			"password": u.Password,
		})

		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))

		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal(gin.H{
			"error": mockErrorResponse,
		})

		assert.NoError(t, err)
		assert.Equal(t, mockErrorResponse.Status(), rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		mockUserService.AssertExpectations(t)
		mockTokenService.AssertExpectations(t)
	})
}
