package controller

import (
	"blog/dao/mysql"
	"blog/models"
	"blog/pkg/logger"
	"blog/response"
	"blog/service"
	"errors"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RegisterHandler 处理注册请求的函数
// @Summary 注册用户接口
// @Description 用户注册接口
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Register body models.ParamRegister true "注册参数"
// @Security ApiKeyAuth
// @Success 200
// @Router /register [post]
func RegisterHandler(c *gin.Context) {
	params := new(models.ParamRegister)

	if err := c.ShouldBindJSON(params); err != nil {
		logger.Error("Register with invalid params", zap.Error(err))
		response.ResponseError(c, response.CodeInvalidParam)
		return
	}

	if err := service.Register(params); err != nil {
		logger.Error("service.Register failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			response.ResponseError(c, response.CodeUserExist)
		} else {
			response.ResponseError(c, response.CodeAddFailed)
		}

		return
	}

	response.ResponseSuccess(c, nil)
}
