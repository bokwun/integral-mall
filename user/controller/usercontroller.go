package controller

//控制层处理请求操作后台
import (
	"integral-mall/common/baseresponse"
	"integral-mall/common/i18n"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"message": i18n.ErrBind})
		ctx.Abort()
		return
	}
	res, err := c.userLogic.Register(r)
	baseresponse.HttpResponse(ctx, res, err)
}

//login
func (c *UserController) Login(ctx *gin.Context) {
	r := new(logic.LoginRequest)
	if err := ctx.ShouldBindJSON(r); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": i18n.ErrBind})
		ctx.Abort()
		return
	}
	res, err := c.userLogic.Login(r)
	baseresponse.HttpResponse(ctx, res, err)
}
