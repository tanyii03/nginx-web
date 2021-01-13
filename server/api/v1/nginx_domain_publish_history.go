package v1

import (
	"nginx-web/global"
    "nginx-web/model"
    "nginx-web/model/request"
    "nginx-web/model/response"
    "nginx-web/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
	"io/ioutil"
)

// @Tags PublishHistory
// @Summary 创建PublishHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PublishHistory true "创建PublishHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /History/createPublishHistory [post]
func CreatePublishHistory(c *gin.Context) {
	var History model.PublishHistory
	_ = c.ShouldBindJSON(&History)
	History.User = getUserName(c)
	if err := service.CreatePublishHistory(History); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags PublishHistory
// @Summary 删除PublishHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PublishHistory true "删除PublishHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /History/deletePublishHistory [delete]
func DeletePublishHistory(c *gin.Context) {
	var History model.PublishHistory
	_ = c.ShouldBindJSON(&History)
	if err := service.DeletePublishHistory(History); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags PublishHistory
// @Summary 批量删除PublishHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PublishHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /History/deletePublishHistoryByIds [delete]
func DeletePublishHistoryByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeletePublishHistoryByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags PublishHistory
// @Summary 更新PublishHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PublishHistory true "更新PublishHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /History/updatePublishHistory [put]
func ViewPublishHistory(c *gin.Context) {
	var History model.PublishHistory
	_ = c.ShouldBindJSON(&History)
	data, err := ioutil.ReadFile("log/publish/" + History.Uuid + ".log")
	if  err != nil {
        global.GVA_LOG.Error("查看失败!", zap.Any("err", err))
		response.FailWithMessage("查看失败", c)
	} else {
		response.OkWithData(string(data), c)
	}
}

// @Tags PublishHistory
// @Summary 用id查询PublishHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PublishHistory true "用id查询PublishHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /History/findPublishHistory [get]
func FindPublishHistory(c *gin.Context) {
	var History model.PublishHistory
	_ = c.ShouldBindQuery(&History)
	if err, reHistory := service.GetPublishHistory(History.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reHistory": reHistory}, c)
	}
}

// @Tags PublishHistory
// @Summary 分页获取PublishHistory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PublishHistorySearch true "分页获取PublishHistory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /History/getPublishHistoryList [get]
func GetPublishHistoryList(c *gin.Context) {
	var pageInfo request.PublishHistorySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetPublishHistoryInfoList(pageInfo); err != nil {
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
