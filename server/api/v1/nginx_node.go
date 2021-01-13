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

// @Tags NginxNode
// @Summary 创建NginxNode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxNode true "创建NginxNode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Node/createNginxNode [post]
func CreateNginxNode(c *gin.Context) {
	var Node model.NginxNode
	_ = c.ShouldBindJSON(&Node)
	if err := service.CreateNginxNode(Node); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags NginxNode
// @Summary 删除NginxNode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxNode true "删除NginxNode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Node/deleteNginxNode [delete]
func DeleteNginxNode(c *gin.Context) {
	var Node model.NginxNode
	_ = c.ShouldBindJSON(&Node)
	if err := service.DeleteNginxNode(Node); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags NginxNode
// @Summary 批量删除NginxNode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除NginxNode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Node/deleteNginxNodeByIds [delete]
func DeleteNginxNodeByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteNginxNodeByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags NginxNode
// @Summary 更新NginxNode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxNode true "更新NginxNode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Node/updateNginxNode [put]
func UpdateNginxNode(c *gin.Context) {
	var Node model.NginxNode
	_ = c.ShouldBindJSON(&Node)
	if err := service.UpdateNginxNode(&Node); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags NginxNode
// @Summary 用id查询NginxNode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxNode true "用id查询NginxNode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Node/findNginxNode [get]
func FindNginxNode(c *gin.Context) {
	var Node model.NginxNode
	_ = c.ShouldBindQuery(&Node)
	if err, reNode := service.GetNginxNode(Node.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reNode": reNode}, c)
	}
}

// @Tags NginxNode
// @Summary 获取NginxNode列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.NginxNodeSearch true "获取NginxNode列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Node/getNginxNodeList [get]
func GetNginxNodeByClusterId(c *gin.Context) {
	var Node model.NginxNode
	_ = c.ShouldBindQuery(&Node)
	if err, Nodes := service.GetNginxNodeByClusterId(Node.ClusterId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"Nodes": Nodes}, c)
	}
}

// @Tags NginxNode
// @Summary 分页获取NginxNode列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.NginxNodeSearch true "分页获取NginxNode列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Node/getNginxNodeList [get]
//func GetNginxNodeList(c *gin.Context) {
//	var pageInfo request.NginxNodeSearch
//	_ = c.ShouldBindQuery(&pageInfo)
//	if err, list, total := service.GetNginxNodeInfoList(pageInfo); err != nil {
//	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
//        response.FailWithMessage("获取失败", c)
//    } else {
//        response.OkWithDetailed(response.PageResult{
//            List:     list,
//            Total:    total,
//            Page:     pageInfo.Page,
//            PageSize: pageInfo.PageSize,
//        }, "获取成功", c)
//    }
//}



