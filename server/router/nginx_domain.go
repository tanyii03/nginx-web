package router

import (
	"nginx-web/api/v1"
	"nginx-web/middleware"
	"github.com/gin-gonic/gin"
)

func InitNginxDomainRouter(Router *gin.RouterGroup) {
	NginxDomainRouter := Router.Group("Domain").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		NginxDomainRouter.POST("createNginxDomain", v1.CreateNginxDomain)   // 新建NginxDomain
		NginxDomainRouter.DELETE("deleteNginxDomain", v1.DeleteNginxDomain) // 删除NginxDomain
		NginxDomainRouter.DELETE("deleteNginxDomainByIds", v1.DeleteNginxDomainByIds) // 批量删除NginxDomain
		NginxDomainRouter.PUT("updateNginxDomain", v1.UpdateNginxDomain)    // 更新NginxDomain
		NginxDomainRouter.GET("findNginxDomain", v1.FindNginxDomain)        // 根据ID获取NginxDomain
		NginxDomainRouter.GET("getNginxDomainList", v1.GetNginxDomainByClusterId)  // 获取NginxDomain列表
		NginxDomainRouter.PUT("previewNginxDomain", v1.PreviewNginxDomain)    // 预览NginxDomain
	}
}
