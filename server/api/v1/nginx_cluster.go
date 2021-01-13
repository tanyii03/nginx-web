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

// @Tags NginxCluster
// @Summary 创建NginxCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxCluster true "创建NginxCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Cluster/createNginxCluster [post]
func CreateNginxCluster(c *gin.Context) {
	var Cluster model.NginxCluster
	_ = c.ShouldBindJSON(&Cluster)
	if err := service.CreateNginxCluster(Cluster); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags NginxCluster
// @Summary 删除NginxCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxCluster true "删除NginxCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Cluster/deleteNginxCluster [delete]
func DeleteNginxCluster(c *gin.Context) {
	var Cluster model.NginxCluster
	_ = c.ShouldBindJSON(&Cluster)
	if err := service.DeleteNginxCluster(Cluster); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags NginxCluster
// @Summary 批量删除NginxCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除NginxCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Cluster/deleteNginxClusterByIds [delete]
func DeleteNginxClusterByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteNginxClusterByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags NginxCluster
// @Summary 更新NginxCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxCluster true "更新NginxCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Cluster/updateNginxCluster [put]
func UpdateNginxCluster(c *gin.Context) {
	var Cluster model.NginxCluster
	_ = c.ShouldBindJSON(&Cluster)
	if err := service.UpdateNginxCluster(&Cluster); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags NginxCluster
// @Summary 用id查询NginxCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxCluster true "用id查询NginxCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Cluster/findNginxCluster [get]
func FindNginxCluster(c *gin.Context) {
	var Cluster model.NginxCluster
	_ = c.ShouldBindQuery(&Cluster)
	if err, reCluster := service.GetNginxCluster(Cluster.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reCluster": reCluster}, c)
	}
}

// @Tags NginxCluster
// @Summary 分页获取NginxCluster列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.NginxClusterSearch true "分页获取NginxCluster列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Cluster/getNginxClusterList [get]
func GetNginxClusterList(c *gin.Context) {
	var pageInfo request.NginxClusterSearch
	_ = c.ShouldBindQuery(&pageInfo)
	AuthId := getUserAuthorityId(c)
	if err, list, total := service.GetNginxClusterInfoList(pageInfo, AuthId); err != nil {
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



func GetAllNginxClusterList(c *gin.Context) {
	authId := getUserAuthorityId(c)
	if err, list := service.GetAllNginxClusterInfoList(authId); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.NginxClusterListResponse{NginxClusters: list}, "获取成功", c)
	}
}

