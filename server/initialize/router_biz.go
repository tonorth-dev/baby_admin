package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

// 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}

func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]

	// 注册婴儿陪护相关路由
	babyRouter := router.RouterGroupApp.Baby
	{
		// 宝宝档案路由 - 需要鉴权
		babyRouter.InitBabyProfileRouter(publicGroup) // 公开路由组，但内部会使用鉴权中间件
		// 音乐相关路由 - 需要鉴权
		babyRouter.InitMusicRouter(publicGroup)
	}

	holder(publicGroup, privateGroup)
}
