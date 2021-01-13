package router

import (
	"nginx-web/api/v1"
	"nginx-web/middleware"
	"github.com/gin-gonic/gin"
)

func InitNginxPoolRouter(Router *gin.RouterGroup) {
	NginxPoolRouter := Router.Group("Pool").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		NginxPoolRouter.POST("createNginxPool", v1.CreateNginxPool)   // 新建NginxPool
		NginxPoolRouter.DELETE("deleteNginxPool", v1.DeleteNginxPool) // 删除NginxPool
		NginxPoolRouter.DELETE("deleteNginxPoolByIds", v1.DeleteNginxPoolByIds) // 批量删除NginxPool
		NginxPoolRouter.PUT("updateNginxPool", v1.UpdateNginxPool)    // 更新NginxPool
		NginxPoolRouter.GET("findNginxPool", v1.FindNginxPool)        // 根据ID获取NginxPool
		NginxPoolRouter.GET("getNginxPoolList", v1.GetNginxPoolList)  // 获取NginxPool列表
		NginxPoolRouter.GET("getAllNginxPoolList", v1.GetAllNginxPoolList)  // 获取所有NginxPool列表
	}
}
