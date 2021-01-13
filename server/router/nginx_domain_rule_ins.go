package router

import (
	"nginx-web/api/v1"
	"nginx-web/middleware"
	"github.com/gin-gonic/gin"
)

func InitDomainRuleInsRouter(Router *gin.RouterGroup) {
	DomainRuleInsRouter := Router.Group("Ins").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		DomainRuleInsRouter.POST("createDomainRuleIns", v1.CreateDomainRuleIns)   // 新建DomainRuleIns
		DomainRuleInsRouter.DELETE("deleteDomainRuleIns", v1.DeleteDomainRuleIns) // 删除DomainRuleIns
		DomainRuleInsRouter.DELETE("deleteDomainRuleInsByIds", v1.DeleteDomainRuleInsByIds) // 批量删除DomainRuleIns
		DomainRuleInsRouter.PUT("updateDomainRuleIns", v1.UpdateDomainRuleIns)    // 更新DomainRuleIns
		DomainRuleInsRouter.GET("findDomainRuleIns", v1.FindDomainRuleIns)        // 根据ID获取DomainRuleIns
		DomainRuleInsRouter.GET("getDomainRuleInsList", v1.GetDomainRuleInsList)  // 获取DomainRuleIns列表
	}
}
