package controller

import (
	"fmt"
	"integral-mall/user/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	UserController struct {
		userLogic *logic.UserLogic
	}
)

func NewUserController(userLogic *logic.UserLogic) *UserController {
	return &UserController{
		userLogic: userLogic,
	}
}

//register
func (c *UserController) Register(ctx *gin.Context) {
	r := new(logic.RegisterRequest)
	if err := ctx.ShouldBindJSON(r); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		ctx.Abort()
		return
	}
	fmt.Println(r)
	res, err := c.userLogic.Register(r)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
