package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/wqh66886/vue-gin-admin/server/server/model/apperrors"
)

type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

func bindData(ctx *gin.Context, req interface{}) bool {
	if err := ctx.ShouldBind(req); err != nil {
		log.Printf("Error binding data: %+v\n", err)
		if errs, ok := err.(validator.ValidationErrors); ok {
			var invalidArgs []invalidArgument
			for _, err := range errs {
				invalidArgs = append(invalidArgs, invalidArgument{
					Field: err.Field(),
					Value: err.Value().(string),
					Tag:   err.Tag(),
					Param: err.Param(),
				})
			}

			err := apperrors.NewBadRequest("Invalid request parameters, see InvalidArgs")

			ctx.JSON(err.Status(), gin.H{
				"error":       err,
				"invalidArgs": invalidArgs,
			})
			return false
		}
		fallBack := apperrors.NewInternal()
		ctx.JSON(fallBack.Status(), gin.H{
			"error": fallBack,
		})
		return false
	}
	return true
}
