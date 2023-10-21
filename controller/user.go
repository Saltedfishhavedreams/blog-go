package controller

import (
	"blog/config"
	"blog/models"
	"blog/pkg/jwt"
	"blog/pkg/logger"
	"blog/response"
	"blog/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RegisterHandler 处理注册请求的函数
// @Summary 注册用户接口
// @Description 用户注册接口
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Register body models.ParamsRegister true "注册参数"
// @Security ApiKeyAuth
// @Success 200
// @Router /register [post]
func RegisterHandler(c *gin.Context) {
	params := new(models.ParamsRegister)

	if err := c.ShouldBindJSON(params); err != nil {
		logger.Error("Register with invalid params", zap.Error(err))
		response.ResponseError(c, response.CodeInvalidParam)
		return
	}

	if err := service.Register(params); err != nil {
		logger.Error("service.Register failed", zap.Error(err))
		if resCode, ok := err.(response.ResCode); ok {
			response.ResponseError(c, resCode)
		} else {
			response.ResponseErrorWithMsg(c, response.CodeServerError, err.Error())
		}

		return
	}

	response.ResponseSuccess(c, nil)
}

// LoginHandler 处理登录请求的函数
// @Summary 用户登录接口
// @Description 用户登录接口
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Login body models.ParamsLogin true "注册参数"
// @Security ApiKeyAuth
// @Success 200
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	params := new(models.ParamsLogin)

	if err := c.ShouldBindJSON(params); err != nil {
		logger.Error("Login with invalid params", zap.Error(err))
		response.ResponseError(c, response.CodeInvalidParam)
		return
	}

	data, err := service.Login(params)
	if err != nil {
		if resCode, ok := err.(response.ResCode); ok {
			response.ResponseError(c, resCode)
		} else {
			response.ResponseErrorWithMsg(c, response.CodeServerError, err.Error())
		}

		return
	}

	response.ResponseSuccess(c, data)
}

// LoginHandler 处理刷新登录状态的函数
// @Summary 用户刷新登录状态接口
// @Description 用户刷新登录状态接口
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Security ApiKeyAuth
// @Success 200
// @Router /refresh_token [get]
func RefreshTokenHandler(c *gin.Context) {
	unknowClaim, exist := c.Get(config.UserClaimKey)
	if claim, ok := unknowClaim.(*jwt.UserClaim); exist && ok {
		data, err := service.RefreshToken(claim)
		if err != nil {
			if resCode, ok := err.(response.ResCode); ok {
				response.ResponseError(c, resCode)
			} else {
				response.ResponseErrorWithMsg(c, response.CodeServerError, err.Error())
			}

			return
		}

		response.ResponseSuccess(c, data)
		return
	}

	response.ResponseError(c, response.CodeInvalidToken)
}
