package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wqh66886/vue-gin-admin/server/server/model"
	"github.com/wqh66886/vue-gin-admin/server/server/model/apperrors"
)

func (h *Handler) Me(ctx *gin.Context) {
	user, exits := ctx.Get("user")
	if !exits {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", ctx)
		err := apperrors.NewInternal()
		ctx.JSON(err.Status(), gin.H{
			"err": err,
		})
		return
	}
	uid := user.(*model.User).UID
	u, err := h.UserService.Get(ctx, uid)
	if err != nil {
		log.Printf("Unable to find user: %v\n%v", uid, err)
		e := apperrors.NewNotFound("user", uid.String())
		ctx.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}
