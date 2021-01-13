package router

import (
	"nginx-web/api/v1"
	"nginx-web/middleware"
	"github.com/gin-gonic/gin"
)

func InitNginxDomainCertRouter(Router *gin.RouterGroup) {
	NginxDomainCertRouter := Router.Group("Cert").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		NginxDomainCertRouter.POST("createNginxDomainCert", v1.CreateNginxDomainCert)   // 新建NginxDomainCert
		NginxDomainCertRouter.DELETE("deleteNginxDomainCert", v1.DeleteNginxDomainCert) // 删除NginxDomainCert
		NginxDomainCertRouter.DELETE("deleteNginxDomainCertByIds", v1.DeleteNginxDomainCertByIds) // 批量删除NginxDomainCert
		NginxDomainCertRouter.PUT("updateNginxDomainCert", v1.UpdateNginxDomainCert)    // 更新NginxDomainCert
		NginxDomainCertRouter.GET("findNginxDomainCert", v1.FindNginxDomainCert)        // 根据ID获取NginxDomainCert
		NginxDomainCertRouter.GET("getNginxDomainCertList", v1.GetNginxDomainCertList)  // 获取NginxDomainCert列表
	}
}
