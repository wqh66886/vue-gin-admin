package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wqh66886/vue-gin-admin/server/server/model"
)

type Handler struct {
	UserService  model.UserService
	TokenService model.TokenService
}

type Config struct {
	R            *gin.Engine
	UserService  model.UserService
	TokenService model.TokenService
}

// func init() {
// 	if err := godotenv.Load("/Users/wangqihao/vue-gin-admin/local.env"); err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }

func NewHandler(c *Config) {

	h := &Handler{
		UserService:  c.UserService,
		TokenService: c.TokenService,
	}
	g := c.R.Group("/api/account")

	g.GET("/me", h.Me)
	g.POST("/signup", h.SignUp)
	g.POST("/signin", h.SignIn)
	g.POST("/signout", h.SignOut)
	g.POST("/tokens", h.Tokens)
	g.POST("/image", h.Image)
	g.DELETE("/image", h.DeleteImage)
	g.PUT("/details", h.Details)
}

func (h *Handler) SignIn(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "it's sign in",
	})
}

func (h *Handler) SignOut(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "it's sign out",
	})
}

func (h *Handler) Tokens(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "it's tokens",
	})
}

func (h *Handler) Image(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "it's image",
	})
}

func (h *Handler) DeleteImage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "it's delete image",
	})
}

func (h *Handler) Details(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "it's details",
	})
}
