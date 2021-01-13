package router

import (
	"nginx-web/api/v1"
	"nginx-web/middleware"
	"github.com/gin-gonic/gin"
)

func InitDomainRuleRouter(Router *gin.RouterGroup) {
	DomainRuleRouter := Router.Group("Rule").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		DomainRuleRouter.POST("createDomainRule", v1.CreateDomainRule)   // 新建DomainRule
		DomainRuleRouter.DELETE("deleteDomainRule", v1.DeleteDomainRule) // 删除DomainRule
		DomainRuleRouter.DELETE("deleteDomainRuleByIds", v1.DeleteDomainRuleByIds) // 批量删除DomainRule
		DomainRuleRouter.PUT("updateDomainRule", v1.UpdateDomainRule)    // 更新DomainRule
		DomainRuleRouter.GET("findDomainRule", v1.FindDomainRule)        // 根据ID获取DomainRule
		DomainRuleRouter.GET("getDomainRuleList", v1.GetDomainRuleList)  // 获取DomainRule列表
	}
}
