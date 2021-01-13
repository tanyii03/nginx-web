package router

import (
	"nginx-web/api/v1"
	"nginx-web/middleware"
	"github.com/gin-gonic/gin"
)

func InitPublishHistoryRouter(Router *gin.RouterGroup) {
	PublishHistoryRouter := Router.Group("History").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		PublishHistoryRouter.POST("createPublishHistory", v1.CreatePublishHistory)   // 新建PublishHistory
		PublishHistoryRouter.DELETE("deletePublishHistory", v1.DeletePublishHistory) // 删除PublishHistory
		PublishHistoryRouter.DELETE("deletePublishHistoryByIds", v1.DeletePublishHistoryByIds) // 批量删除PublishHistory
		PublishHistoryRouter.POST("viewPublishHistory", v1.ViewPublishHistory)    // 查看PublishHistory
		PublishHistoryRouter.GET("findPublishHistory", v1.FindPublishHistory)        // 根据ID获取PublishHistory
		PublishHistoryRouter.GET("getPublishHistoryList", v1.GetPublishHistoryList)  // 获取PublishHistory列表
	}
}
