package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wqh66886/vue-gin-admin/server/server/model"
	"github.com/wqh66886/vue-gin-admin/server/server/model/apperrors"
)

/*
* 在定义结构体字段时，除字段名称和数据类型外，还可以使用反引号为结构体字段声明元信息，
* 这种元信息称为Tag，用于编译阶段关联到字段当中,如我们将上面例子中的结构体修改为
 */
type signupReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

func (h *Handler) SignUp(ctx *gin.Context) {
	var req signupReq
	if ok := bindData(ctx, &req); !ok {
		return
	}

	u := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	err := h.UserService.SignUp(ctx, u)

	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		ctx.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
}
