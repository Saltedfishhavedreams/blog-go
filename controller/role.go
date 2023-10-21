package controller

import (
	"blog/response"
	"blog/service"

	"github.com/gin-gonic/gin"
)

// RegisterHandler 处理获取角色列表请求的函数
// @Summary 获取角色列表
// @Description 获取角色列表接口
// @Tags 角色相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Security ApiKeyAuth
// @Success 200
// @Router /role [get]
func GetRoleListHandler(c *gin.Context) {
	data, err := service.GetRoleList()
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

// RegisterHandler 处理获取角色请求的函数
// @Summary 获取角色
// @Description 获取角色接口
// @Tags 角色相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param role_id path string  true "角色 role_id 字段"
// @Security ApiKeyAuth
// @Success 200
// @Router /role/{role_id} [get]
func GetRoleHandler(c *gin.Context) {
	data, err := service.GetRole(c.Param("role_id"))
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
