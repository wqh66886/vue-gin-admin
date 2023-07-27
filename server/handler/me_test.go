package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wqh66886/vue-gin-admin/server/server/model"
	"github.com/wqh66886/vue-gin-admin/server/server/model/apperrors"
	"github.com/wqh66886/vue-gin-admin/server/server/model/mocks"
)

func TestMe(t *testing.T) {
	// setUp
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockUserResp := &model.User{
			UID:   uid,
			Email: "bob@bob.com",
			Name:  "Bobby Bobson",
		}
		mockUserService := new(mocks.MocUserService)
		mockUserService.On("Get", mock.AnythingOfType("*gin.Context"), uid).Return(mockUserResp, nil)
		rr := httptest.NewRecorder()

		router := gin.Default()
		router.Use(func(ctx *gin.Context) {
			ctx.Set("user", &model.User{
				UID: uid,
			})
		})

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, "/me", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal(gin.H{
			"user": mockUserResp,
		})
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		mockUserService.AssertExpectations(t)
	})
	t.Run("NoContextUser", func(t *testing.T) {
		mockUserService := new(mocks.MocUserService)
		mockUserService.On("Get", mock.Anything, mock.Anything).Return(nil, nil)

		rr := httptest.NewRecorder()

		router := gin.Default()
		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, "/me", nil)

		assert.NoError(t, err)
		router.ServeHTTP(rr, request)
		assert.Equal(t, 500, rr.Code)
		mockUserService.AssertCalled(t, "Get", mock.Anything)
	})

	t.Run("NotFound", func(t *testing.T) {
		uid, _ := uuid.NewRandom()
		mockUserService := new(mocks.MocUserService)
		mockUserService.On("Get", mock.Anything, uid).Return(nil, fmt.Errorf("Some error down call chain"))

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.Use(func(ctx *gin.Context) {
			ctx.Set("user", &model.User{
				UID: uid,
			})
		})

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})
		request, err := http.NewRequest(http.MethodGet, "/me", nil)
		assert.NoError(t, err)
		router.ServeHTTP(rr, request)

		respErr := apperrors.NewNotFound("user", uid.String())

		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})

		assert.NoError(t, err)
		assert.Equal(t, respErr.Status(), rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		mockUserService.AssertExpectations(t)
	})
}
