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

// @Tags DomainRuleIns
// @Summary 创建DomainRuleIns
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DomainRuleIns true "创建DomainRuleIns"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Ins/createDomainRuleIns [post]
func CreateDomainRuleIns(c *gin.Context) {
	var Ins model.DomainRuleIns
	_ = c.ShouldBindJSON(&Ins)
	if err := service.CreateDomainRuleIns(Ins); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags DomainRuleIns
// @Summary 删除DomainRuleIns
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DomainRuleIns true "删除DomainRuleIns"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Ins/deleteDomainRuleIns [delete]
func DeleteDomainRuleIns(c *gin.Context) {
	var Ins model.DomainRuleIns
	_ = c.ShouldBindJSON(&Ins)
	if err := service.DeleteDomainRuleIns(Ins); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags DomainRuleIns
// @Summary 批量删除DomainRuleIns
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DomainRuleIns"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Ins/deleteDomainRuleInsByIds [delete]
func DeleteDomainRuleInsByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteDomainRuleInsByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags DomainRuleIns
// @Summary 更新DomainRuleIns
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DomainRuleIns true "更新DomainRuleIns"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Ins/updateDomainRuleIns [put]
func UpdateDomainRuleIns(c *gin.Context) {
	var Ins model.DomainRuleIns
	_ = c.ShouldBindJSON(&Ins)
	if err := service.UpdateDomainRuleIns(&Ins); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags DomainRuleIns
// @Summary 用id查询DomainRuleIns
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DomainRuleIns true "用id查询DomainRuleIns"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Ins/findDomainRuleIns [get]
func FindDomainRuleIns(c *gin.Context) {
	var Ins model.DomainRuleIns
	_ = c.ShouldBindQuery(&Ins)
	if err, reIns := service.GetDomainRuleIns(Ins.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reIns": reIns}, c)
	}
}

// @Tags DomainRuleIns
// @Summary 分页获取DomainRuleIns列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DomainRuleInsSearch true "分页获取DomainRuleIns列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Ins/getDomainRuleInsList [get]
func GetDomainRuleInsList(c *gin.Context) {
	var pageInfo request.DomainRuleInsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDomainRuleInsInfoList(pageInfo); err != nil {
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
