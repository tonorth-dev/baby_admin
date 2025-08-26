package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MiniprogramUserRouter struct{}

// InitMiniprogramUserRouter 初始化小程序用户路由
func (m *MiniprogramUserRouter) InitMiniprogramUserRouter(Router *gin.RouterGroup) {
	miniprogramRouter := Router.Group("miniprogram")
	miniprogramRouterWithoutRecord := Router.Group("miniprogram")
	
	// 不记录操作日志的路由组
	miniprogramRouterWithoutRecord.Use(middleware.OperationRecord())

	{
		// 无需鉴权的接口
		miniprogramRouter.POST("login", miniprogramUserApi.Login) // 小程序登录
	}
	
	// 需要小程序鉴权的路由组
	miniprogramAuthRouter := miniprogramRouter.Group("")
	miniprogramAuthRouter.Use(middleware.MiniprogramJWTAuth())
	{
		miniprogramAuthRouter.GET("getUserInfo", miniprogramUserApi.GetUserInfo)       // 获取用户信息
		miniprogramAuthRouter.PUT("updateProfile", miniprogramUserApi.UpdateProfile) // 更新用户资料
	}
	
	// 可选鉴权的路由组（用于一些既可以匿名访问也可以登录访问的接口）
	miniprogramOptionalRouter := miniprogramRouter.Group("")
	miniprogramOptionalRouter.Use(middleware.MiniprogramOptionalAuth())
	{
		// 这里可以添加一些可选登录的接口
		// 例如：商品列表、公告列表等
	}
}

// InitAppUserRouter 初始化APP用户路由
func (m *MiniprogramUserRouter) InitAppUserRouter(Router *gin.RouterGroup) {
	// APP注册和登录（无需鉴权）
	appRouter := Router.Group("app")
	{
		appRouter.POST("register", miniprogramUserApi.AppRegister) // APP注册
		appRouter.POST("login", miniprogramUserApi.AppLogin)       // APP登录
	}
	
	// 需要鉴权的用户相关接口
	userRouter := Router.Group("user")
	userRouter.Use(middleware.MiniprogramJWTAuth()) // 使用同一套JWT中间件
	{
		userRouter.GET("info", miniprogramUserApi.GetUserInfo)              // 获取用户信息（复用）
		userRouter.PUT("profile", miniprogramUserApi.UpdateUserProfile)     // 更新用户资料
		userRouter.PUT("changePassword", miniprogramUserApi.ChangePassword) // 修改密码
	}
}