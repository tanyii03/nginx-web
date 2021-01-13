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

// @Tags NginxDomainCert
// @Summary 创建NginxDomainCert
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxDomainCert true "创建NginxDomainCert"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Cert/createNginxDomainCert [post]
func CreateNginxDomainCert(c *gin.Context) {
	var Cert model.NginxDomainCertFile
	_ = c.ShouldBindJSON(&Cert)
	if err := service.CreateNginxDomainCertByFile(Cert); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags NginxDomainCert
// @Summary 删除NginxDomainCert
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxDomainCert true "删除NginxDomainCert"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Cert/deleteNginxDomainCert [delete]
func DeleteNginxDomainCert(c *gin.Context) {
	var Cert model.NginxDomainCert
	_ = c.ShouldBindJSON(&Cert)
	if err := service.DeleteNginxDomainCert(Cert); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags NginxDomainCert
// @Summary 批量删除NginxDomainCert
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除NginxDomainCert"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Cert/deleteNginxDomainCertByIds [delete]
func DeleteNginxDomainCertByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteNginxDomainCertByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags NginxDomainCert
// @Summary 更新NginxDomainCert
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxDomainCert true "更新NginxDomainCert"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Cert/updateNginxDomainCert [put]
func UpdateNginxDomainCert(c *gin.Context) {
	var Cert model.NginxDomainCert
	_ = c.ShouldBindJSON(&Cert)
	if err := service.UpdateNginxDomainCert(&Cert); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags NginxDomainCert
// @Summary 用id查询NginxDomainCert
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxDomainCert true "用id查询NginxDomainCert"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Cert/findNginxDomainCert [get]
func FindNginxDomainCert(c *gin.Context) {
	var Cert model.NginxDomainCert
	_ = c.ShouldBindQuery(&Cert)
	if err, reCert := service.GetNginxDomainCert(Cert.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reCert": reCert}, c)
	}
}

// @Tags NginxDomainCert
// @Summary 分页获取NginxDomainCert列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.NginxDomainCertSearch true "分页获取NginxDomainCert列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Cert/getNginxDomainCertList [get]
func GetNginxDomainCertList(c *gin.Context) {
	var pageInfo request.NginxDomainCertSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetNginxDomainCertInfoList(pageInfo); err != nil {
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
