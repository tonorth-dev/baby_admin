package baby

import (
	"baby_admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MusicRouter struct{}

// InitMusicRouter 初始化音乐路由
func (m *MusicRouter) InitMusicRouter(Router *gin.RouterGroup) {
	musicRouter := Router.Group("music")
	musicRouter.Use(middleware.MiniprogramJWTAuth()) // 需要登录验证
	{
		musicRouter.GET("categories", musicApi.GetMusicCategories)     // 获取音乐分类
		musicRouter.GET("list", musicApi.GetMusicList)                 // 获取音乐列表
		musicRouter.GET("recommendations", musicApi.GetRecommendations) // 获取推荐音乐
		musicRouter.GET("favorites", musicApi.GetUserFavorites)        // 获取收藏音乐
		musicRouter.GET("history", musicApi.GetPlayHistory)            // 获取播放历史
		musicRouter.GET(":id", musicApi.GetMusicDetail)                // 获取音乐详情
		musicRouter.POST("play", musicApi.PlayMusic)                   // 播放音乐
		musicRouter.POST(":id/favorite", musicApi.ToggleFavorite)      // 切换收藏状态
	}
}