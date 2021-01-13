package v1

import (
	"nginx-web/global"
    "nginx-web/model"
    "nginx-web/model/request"
    "nginx-web/model/response"
    "nginx-web/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags NginxPool
// @Summary 创建NginxPool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxPool true "创建NginxPool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Pool/createNginxPool [post]
func CreateNginxPool(c *gin.Context) {
	var Pool model.NginxPool
	_ = c.ShouldBindJSON(&Pool)
	Pool.AuthorityId =  getUserAuthorityId(c)
	if err := service.CreateNginxPool(Pool); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags NginxPool
// @Summary 删除NginxPool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxPool true "删除NginxPool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Pool/deleteNginxPool [delete]
func DeleteNginxPool(c *gin.Context) {
	var Pool model.NginxPool
	_ = c.ShouldBindJSON(&Pool)
	if err := service.DeleteNginxPool(Pool); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败" + err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags NginxPool
// @Summary 批量删除NginxPool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除NginxPool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Pool/deleteNginxPoolByIds [delete]
func DeleteNginxPoolByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteNginxPoolByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags NginxPool
// @Summary 更新NginxPool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxPool true "更新NginxPool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Pool/updateNginxPool [put]
func UpdateNginxPool(c *gin.Context) {
	var Pool model.NginxPool
	_ = c.ShouldBindJSON(&Pool)
	if err := service.UpdateNginxPool(&Pool); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags NginxPool
// @Summary 用id查询NginxPool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxPool true "用id查询NginxPool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Pool/findNginxPool [get]
func FindNginxPool(c *gin.Context) {
	var Pool model.NginxPool
	_ = c.ShouldBindQuery(&Pool)
	if err, rePool := service.GetNginxPool(Pool.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rePool": rePool}, c)
	}
}

// @Tags NginxPool
// @Summary 分页获取NginxPool列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.NginxPoolSearch true "分页获取NginxPool列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Pool/getNginxPoolList [get]
func GetNginxPoolList(c *gin.Context) {
	var pageInfo request.NginxPoolSearch
	_ = c.ShouldBindQuery(&pageInfo)
	authId := getUserAuthorityId(c)
	if err, list, total := service.GetNginxPoolInfoList(pageInfo, authId); err != nil {
	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}




func GetAllNginxPoolList(c *gin.Context) {
	authId := getUserAuthorityId(c)
	if err, list := service.GetAllNginxPoolInfoList(authId); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.NginxPoolListResponse{NginxPools: list}, "获取成功", c)
	}
}