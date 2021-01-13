package router

import (
	"nginx-web/api/v1"
	"nginx-web/middleware"
	"github.com/gin-gonic/gin"
)

func InitPoolNodeRouter(Router *gin.RouterGroup) {
	PoolNodeRouter := Router.Group("PNode").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		PoolNodeRouter.POST("createPoolNode", v1.CreatePoolNode)   // 新建PoolNode
		PoolNodeRouter.DELETE("deletePoolNode", v1.DeletePoolNode) // 删除PoolNode
		PoolNodeRouter.DELETE("deletePoolNodeByIds", v1.DeletePoolNodeByIds) // 批量删除PoolNode
		PoolNodeRouter.PUT("updatePoolNode", v1.UpdatePoolNode)    // 更新PoolNode
		PoolNodeRouter.GET("findPoolNode", v1.FindPoolNode)        // 根据ID获取PoolNode
		PoolNodeRouter.GET("getPoolNodeList", v1.GetPoolNodeList)  // 获取PoolNode列表
	}
}
