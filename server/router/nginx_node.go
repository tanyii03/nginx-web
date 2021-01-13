package router

import (
	"nginx-web/api/v1"
	"nginx-web/middleware"
	"github.com/gin-gonic/gin"
)

func InitNginxNodeRouter(Router *gin.RouterGroup) {
	NginxNodeRouter := Router.Group("Node").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		NginxNodeRouter.POST("createNginxNode", v1.CreateNginxNode)   // 新建NginxNode
		NginxNodeRouter.DELETE("deleteNginxNode", v1.DeleteNginxNode) // 删除NginxNode
		NginxNodeRouter.DELETE("deleteNginxNodeByIds", v1.DeleteNginxNodeByIds) // 批量删除NginxNode
		NginxNodeRouter.PUT("updateNginxNode", v1.UpdateNginxNode)    // 更新NginxNode
		NginxNodeRouter.GET("findNginxNode", v1.FindNginxNode)        // 根据ID获取NginxNode
		NginxNodeRouter.GET("getNginxNodeList", v1.GetNginxNodeByClusterId)  // 获取NginxNode列表
	}
}
