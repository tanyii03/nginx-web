package router

import (
	"nginx-web/api/v1"
	"nginx-web/middleware"
	"github.com/gin-gonic/gin"
)

func InitNginxClusterRouter(Router *gin.RouterGroup) {
	NginxClusterRouter := Router.Group("Cluster").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		NginxClusterRouter.POST("createNginxCluster", v1.CreateNginxCluster)   // 新建NginxCluster
		NginxClusterRouter.DELETE("deleteNginxCluster", v1.DeleteNginxCluster) // 删除NginxCluster
		NginxClusterRouter.DELETE("deleteNginxClusterByIds", v1.DeleteNginxClusterByIds) // 批量删除NginxCluster
		NginxClusterRouter.PUT("updateNginxCluster", v1.UpdateNginxCluster)    // 更新NginxCluster
		NginxClusterRouter.GET("findNginxCluster", v1.FindNginxCluster)        // 根据ID获取NginxCluster
		NginxClusterRouter.GET("getNginxClusterList", v1.GetNginxClusterList)  // 获取NginxCluster列表
		NginxClusterRouter.GET("getAllNginxClusterList", v1.GetAllNginxClusterList)  // 获取NginxCluster列表
	}
}
