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

// @Tags PoolNode
// @Summary 创建PoolNode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PoolNode true "创建PoolNode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PNode/createPoolNode [post]
func CreatePoolNode(c *gin.Context) {
	var PNode model.PoolNode
	_ = c.ShouldBindJSON(&PNode)
	if err := service.CreatePoolNode(PNode); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags PoolNode
// @Summary 删除PoolNode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PoolNode true "删除PoolNode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PNode/deletePoolNode [delete]
func DeletePoolNode(c *gin.Context) {
	var PNode model.PoolNode
	_ = c.ShouldBindJSON(&PNode)
	if err := service.DeletePoolNode(PNode); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags PoolNode
// @Summary 批量删除PoolNode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PoolNode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /PNode/deletePoolNodeByIds [delete]
func DeletePoolNodeByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeletePoolNodeByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags PoolNode
// @Summary 更新PoolNode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PoolNode true "更新PoolNode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PNode/updatePoolNode [put]
func UpdatePoolNode(c *gin.Context) {
	var PNode model.PoolNode
	_ = c.ShouldBindJSON(&PNode)
	if err := service.UpdatePoolNode(&PNode); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags PoolNode
// @Summary 用id查询PoolNode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PoolNode true "用id查询PoolNode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PNode/findPoolNode [get]
func FindPoolNode(c *gin.Context) {
	var PNode model.PoolNode
	_ = c.ShouldBindQuery(&PNode)
	if err, rePNode := service.GetPoolNode(PNode.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rePNode": rePNode}, c)
	}
}

// @Tags PoolNode
// @Summary 分页获取PoolNode列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PoolNodeSearch true "分页获取PoolNode列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PNode/getPoolNodeList [get]
func GetPoolNodeList(c *gin.Context) {
	var pageInfo request.PoolNodeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetPoolNodeInfoList(pageInfo); err != nil {
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
