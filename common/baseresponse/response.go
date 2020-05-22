package baseresponse

import (
	"integral-mall/common/baseerror"
	"integral-mall/common/i18n"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func ParamError(ctx *gin.Context, err validator.ValidationErrors) {
// }

func HttpResponse(ctx *gin.Context, res, err interface{}) {
	baseError, ok := err.(*baseerror.BaseError)
	if ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": baseError.Error()})
		ctx.Abort()
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": i18n.ErrServer})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": res})
}
