package baby

import (
	"baby_admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BabyProfileRouter struct{}

// InitBabyProfileRouter 初始化宝宝档案路由
func (b *BabyProfileRouter) InitBabyProfileRouter(Router *gin.RouterGroup) {
	babyProfileRouter := Router.Group("baby/profile")
	babyProfileRouter.Use(middleware.MiniprogramJWTAuth()) // 需要登录验证
	{
		babyProfileRouter.POST("", babyProfileApi.CreateBabyProfile)       // 创建宝宝档案
		babyProfileRouter.GET("active", babyProfileApi.GetActiveBabyProfile) // 获取活跃宝宝档案
		babyProfileRouter.GET("list", babyProfileApi.GetBabyProfileList)   // 获取宝宝档案列表
		babyProfileRouter.GET(":id", babyProfileApi.GetBabyProfile)        // 获取宝宝档案详情
		babyProfileRouter.PUT("", babyProfileApi.UpdateBabyProfile)        // 更新宝宝档案
		babyProfileRouter.DELETE(":id", babyProfileApi.DeleteBabyProfile)  // 删除宝宝档案
		babyProfileRouter.PUT(":id/setActive", babyProfileApi.SetActiveBaby) // 设置活跃宝宝
	}
}