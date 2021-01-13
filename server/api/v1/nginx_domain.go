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

// @Tags NginxDomain
// @Summary 创建NginxDomain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxDomain true "创建NginxDomain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Domain/createNginxDomain [post]
func CreateNginxDomain(c *gin.Context) {
	var Domain model.NginxDomain
	_ = c.ShouldBindJSON(&Domain)
	if err := service.CreateNginxDomain(Domain); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		var Config model.AutoConfigStruct
		Config.Domain = Domain
		service.CreateConfigTemp(Config)
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags NginxDomain
// @Summary 删除NginxDomain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxDomain true "删除NginxDomain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Domain/deleteNginxDomain [delete]
func DeleteNginxDomain(c *gin.Context) {
	var Domain model.NginxDomain
	_ = c.ShouldBindJSON(&Domain)
	User := getUserName(c)
	if err := service.DeleteNginxDomain(Domain, User); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags NginxDomain
// @Summary 批量删除NginxDomain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除NginxDomain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Domain/deleteNginxDomainByIds [delete]
func DeleteNginxDomainByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteNginxDomainByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags NginxDomain
// @Summary 更新NginxDomain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxDomain true "更新NginxDomain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Domain/updateNginxDomain [put]
func UpdateNginxDomain(c *gin.Context) {
	var Domain model.NginxDomain
	_ = c.ShouldBindJSON(&Domain)
	if err := service.UpdateNginxDomain(&Domain); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		var Config model.AutoConfigStruct
		Config.Domain = Domain
		service.CreateConfigTemp(Config)
		response.OkWithMessage("更新成功", c)
	}
}


func PreviewNginxDomain(c *gin.Context) {
	var Domain model.NginxDomain
	_ = c.ShouldBindJSON(&Domain)
	data, err := ioutil.ReadFile("autoCode/domain." + Domain.DomainName + ".conf")
	if err != nil {
		global.GVA_LOG.Error("预览失败!", zap.Any("err", err))
		response.FailWithMessage("预览失败", c)
	} else {
		response.OkWithData(string(data), c)
	}
}


// @Tags NginxDomain
// @Summary 用id查询NginxDomain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxDomain true "用id查询NginxDomain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Domain/findNginxDomain [get]
func FindNginxDomain(c *gin.Context) {
	var Domain model.NginxDomain
	_ = c.ShouldBindQuery(&Domain)
	if err, reDomain := service.GetNginxDomain(Domain.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reDomain": reDomain}, c)
	}
}


func GetNginxDomainByClusterId(c *gin.Context) {
	var Info model.NginxCluster
	_ = c.ShouldBindQuery(&Info)
	if err, reDomains := service.GetNginxDomainInfoByClusterId(Info); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.NginxDomainListResponse{NginxDomains: reDomains}, "获取成功", c)
	}
}



// @Tags NginxDomain
// @Summary 分页获取NginxDomain列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.NginxDomainSearch true "分页获取NginxDomain列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Domain/getNginxDomainList [get]
//func GetNginxDomainList(c *gin.Context) {
//	var pageInfo request.NginxDomainSearch
//	_ = c.ShouldBindQuery(&pageInfo)
//	if err, list, total := service.GetNginxDomainInfoList(pageInfo); err != nil {
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
