package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MiniprogramUserApi struct{}

// Login 小程序用户登录
// @Tags MiniprogramUser
// @Summary 小程序用户登录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.MiniprogramLogin true "小程序登录信息"
// @Success 200 {object} response.Response{data=systemRes.MiniprogramLoginResponse,msg=string} "登录成功返回用户信息和token"
// @Router /miniprogram/login [post]
func (m *MiniprogramUserApi) Login(c *gin.Context) {
	var req request.MiniprogramLogin
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 验证必要字段
	if req.Code == "" {
		response.FailWithMessage("登录凭证code不能为空", c)
		return
	}

	user, token, expiresAt, err := miniprogramUserService.Login(req)
	if err != nil {
		global.GVA_LOG.Error("小程序登录失败!", zap.Error(err))
		response.FailWithMessage("登录失败: "+err.Error(), c)
		return
	}

	// 设置token到cookie
	utils.SetToken(c, token, int(expiresAt))

	response.OkWithDetailed(systemRes.MiniprogramLoginResponse{
		User:      *user.(*system.MiniprogramUser),
		Token:     token,
		ExpiresAt: expiresAt,
	}, "登录成功", c)
}

// GetUserInfo 获取小程序用户信息
// @Tags MiniprogramUser
// @Summary 获取小程序用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=systemRes.MiniprogramUserResponse,msg=string} "获取用户信息成功"
// @Router /miniprogram/getUserInfo [get]
func (m *MiniprogramUserApi) GetUserInfo(c *gin.Context) {
	claims, exist := c.Get("claims")
	if !exist {
		response.FailWithMessage("未登录或非法访问", c)
		return
	}

	customClaims, ok := claims.(*request.CustomClaims)
	if !ok {
		response.FailWithMessage("token解析失败", c)
		return
	}

	user, err := miniprogramUserService.GetUserInfo(customClaims.UUID.String())
	if err != nil {
		global.GVA_LOG.Error("获取用户信息失败!", zap.Error(err))
		response.FailWithMessage("获取用户信息失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(systemRes.MiniprogramUserResponse{
		User: user,
	}, "获取用户信息成功", c)
}

// UpdateProfile 更新小程序用户资料
// @Tags MiniprogramUser
// @Summary 更新小程序用户资料
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.MiniprogramUpdateProfile true "用户资料信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /miniprogram/updateProfile [put]
func (m *MiniprogramUserApi) UpdateProfile(c *gin.Context) {
	var req request.MiniprogramUpdateProfile
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	claims, exist := c.Get("claims")
	if !exist {
		response.FailWithMessage("未登录或非法访问", c)
		return
	}

	customClaims, ok := claims.(*request.CustomClaims)
	if !ok {
		response.FailWithMessage("token解析失败", c)
		return
	}

	err = miniprogramUserService.UpdateUserProfile(customClaims.UUID.String(), req)
	if err != nil {
		global.GVA_LOG.Error("更新用户资料失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

// AppRegister APP注册
// @Tags MiniprogramUser
// @Summary APP用户注册
// @accept application/json
// @Produce application/json
// @Param data body request.AppRegister true "注册信息"
// @Success 200 {object} response.Response{data=systemRes.AppRegisterResponse,msg=string} "注册成功"
// @Router /app/register [post]
func (m *MiniprogramUserApi) AppRegister(c *gin.Context) {
	var req request.AppRegister
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 参数验证
	if len(req.Username) < 3 || len(req.Username) > 20 {
		response.FailWithMessage("用户名长度必须在3-20位之间", c)
		return
	}
	if len(req.Password) < 6 || len(req.Password) > 20 {
		response.FailWithMessage("密码长度必须在6-20位之间", c)
		return
	}
	if req.Email == "" {
		response.FailWithMessage("邮箱不能为空", c)
		return
	}

	user, token, expiresAt, err := miniprogramUserService.AppRegister(req)
	if err != nil {
		global.GVA_LOG.Error("APP注册失败!", zap.Error(err))
		response.FailWithMessage("注册失败: "+err.Error(), c)
		return
	}

	// 转换为脱敏的用户信息
	var userInfo systemRes.UserInfo
	userInfo.FromMiniprogramUser(user.(*system.MiniprogramUser))

	// 设置token到cookie
	utils.SetToken(c, token, int(expiresAt))

	response.OkWithDetailed(systemRes.AppRegisterResponse{
		User:      userInfo,
		Token:     token,
		ExpiresAt: expiresAt,
	}, "注册成功", c)
}

// AppLogin APP登录
// @Tags MiniprogramUser
// @Summary APP用户登录
// @accept application/json
// @Produce application/json
// @Param data body request.AppLogin true "登录信息"
// @Success 200 {object} response.Response{data=systemRes.AppLoginResponse,msg=string} "登录成功"
// @Router /app/login [post]
func (m *MiniprogramUserApi) AppLogin(c *gin.Context) {
	var req request.AppLogin
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	user, token, expiresAt, err := miniprogramUserService.AppLogin(req)
	if err != nil {
		global.GVA_LOG.Error("APP登录失败!", zap.Error(err))
		response.FailWithMessage("登录失败: "+err.Error(), c)
		return
	}

	// 转换为脱敏的用户信息
	var userInfo systemRes.UserInfo
	userInfo.FromMiniprogramUser(user.(*system.MiniprogramUser))

	// 设置token到cookie
	utils.SetToken(c, token, int(expiresAt))

	response.OkWithDetailed(systemRes.AppLoginResponse{
		User:      userInfo,
		Token:     token,
		ExpiresAt: expiresAt,
	}, "登录成功", c)
}

// UpdateUserProfile 更新用户资料（通用）
// @Tags MiniprogramUser
// @Summary 更新用户资料
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UpdateProfile true "用户资料信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /user/profile [put]
func (m *MiniprogramUserApi) UpdateUserProfile(c *gin.Context) {
	var req request.UpdateProfile
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	claims, exist := c.Get("claims")
	if !exist {
		response.FailWithMessage("未登录或非法访问", c)
		return
	}

	customClaims, ok := claims.(*request.CustomClaims)
	if !ok {
		response.FailWithMessage("token解析失败", c)
		return
	}

	err = miniprogramUserService.UpdateProfile(customClaims.UUID.String(), req)
	if err != nil {
		global.GVA_LOG.Error("更新用户资料失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

// ChangePassword 修改密码
// @Tags MiniprogramUser
// @Summary 修改密码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ChangePassword true "修改密码信息"
// @Success 200 {object} response.Response{msg=string} "修改成功"
// @Router /user/changePassword [put]
func (m *MiniprogramUserApi) ChangePassword(c *gin.Context) {
	var req request.ChangePassword
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	claims, exist := c.Get("claims")
	if !exist {
		response.FailWithMessage("未登录或非法访问", c)
		return
	}

	customClaims, ok := claims.(*request.CustomClaims)
	if !ok {
		response.FailWithMessage("token解析失败", c)
		return
	}

	err = miniprogramUserService.ChangePassword(customClaims.UUID.String(), req)
	if err != nil {
		global.GVA_LOG.Error("修改密码失败!", zap.Error(err))
		response.FailWithMessage("修改失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("修改成功", c)
}