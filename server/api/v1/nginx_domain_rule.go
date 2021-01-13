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

// @Tags DomainRule
// @Summary 创建DomainRule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DomainRule true "创建DomainRule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Rule/createDomainRule [post]
func CreateDomainRule(c *gin.Context) {
	var Rule model.DomainRule
	_ = c.ShouldBindJSON(&Rule)
	if err := service.CreateDomainRule(Rule); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags DomainRule
// @Summary 删除DomainRule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DomainRule true "删除DomainRule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Rule/deleteDomainRule [delete]
func DeleteDomainRule(c *gin.Context) {
	var Rule model.DomainRule
	_ = c.ShouldBindJSON(&Rule)
	if err := service.DeleteDomainRule(Rule); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags DomainRule
// @Summary 批量删除DomainRule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DomainRule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Rule/deleteDomainRuleByIds [delete]
func DeleteDomainRuleByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteDomainRuleByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags DomainRule
// @Summary 更新DomainRule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DomainRule true "更新DomainRule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Rule/updateDomainRule [put]
func UpdateDomainRule(c *gin.Context) {
	var Rule model.DomainRule
	_ = c.ShouldBindJSON(&Rule)
	if err := service.UpdateDomainRule(&Rule); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags DomainRule
// @Summary 用id查询DomainRule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DomainRule true "用id查询DomainRule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Rule/findDomainRule [get]
func FindDomainRule(c *gin.Context) {
	var Rule model.DomainRule
	_ = c.ShouldBindQuery(&Rule)
	if err, reRule := service.GetDomainRule(Rule.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reRule": reRule}, c)
	}
}

// @Tags DomainRule
// @Summary 分页获取DomainRule列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DomainRuleSearch true "分页获取DomainRule列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Rule/getDomainRuleList [get]
func GetDomainRuleList(c *gin.Context) {
	var pageInfo request.DomainRuleSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDomainRuleInfoList(pageInfo); err != nil {
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
