package baby

import (
	"baby_admin/server/global"
	"baby_admin/server/model/baby/request"
	"baby_admin/server/model/common/response"
	systemReq "baby_admin/server/model/system/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type BabyProfileApi struct{}

// CreateBabyProfile 创建宝宝档案
// @Tags BabyProfile
// @Summary 创建宝宝档案
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CreateBabyProfileRequest true "宝宝档案信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /baby/profile [post]
func (b *BabyProfileApi) CreateBabyProfile(c *gin.Context) {
	var req request.CreateBabyProfileRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取当前用户ID
	claims, exist := c.Get("claims")
	if !exist {
		response.FailWithMessage("未登录或非法访问", c)
		return
	}

	customClaims, ok := claims.(*systemReq.CustomClaims)
	if !ok {
		response.FailWithMessage("token解析失败", c)
		return
	}

	err = babyProfileService.CreateBabyProfile(customClaims.BaseClaims.ID, &req)
	if err != nil {
		global.GVA_LOG.Error("创建宝宝档案失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

// GetBabyProfile 获取宝宝档案详情
// @Tags BabyProfile
// @Summary 获取宝宝档案详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id path int true "宝宝档案ID"
// @Success 200 {object} response.Response{data=response.BabyProfileResponse,msg=string} "获取成功"
// @Router /baby/profile/{id} [get]
func (b *BabyProfileApi) GetBabyProfile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("ID格式错误", c)
		return
	}

	// 获取当前用户ID
	claims, exist := c.Get("claims")
	if !exist {
		response.FailWithMessage("未登录或非法访问", c)
		return
	}

	customClaims, ok := claims.(*systemReq.CustomClaims)
	if !ok {
		response.FailWithMessage("token解析失败", c)
		return
	}

	profile, err := babyProfileService.GetBabyProfile(uint(id), customClaims.BaseClaims.ID)
	if err != nil {
		global.GVA_LOG.Error("获取宝宝档案失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(profile, "获取成功", c)
}

// GetActiveBabyProfile 获取当前活跃的宝宝档案
// @Tags BabyProfile
// @Summary 获取当前活跃的宝宝档案
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=response.BabyProfileResponse,msg=string} "获取成功"
// @Router /baby/profile/active [get]
func (b *BabyProfileApi) GetActiveBabyProfile(c *gin.Context) {
	// 获取当前用户ID
	claims, exist := c.Get("claims")
	if !exist {
		response.FailWithMessage("未登录或非法访问", c)
		return
	}

	customClaims, ok := claims.(*systemReq.CustomClaims)
	if !ok {
		response.FailWithMessage("token解析失败", c)
		return
	}

	profile, err := babyProfileService.GetActiveBabyProfile(customClaims.BaseClaims.ID)
	if err != nil {
		global.GVA_LOG.Error("获取活跃宝宝档案失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(profile, "获取成功", c)
}

// GetBabyProfileList 获取宝宝档案列表
// @Tags BabyProfile
// @Summary 获取宝宝档案列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.BabyProfileSearch true "分页参数"
// @Success 200 {object} response.Response{data=response.BabyProfileListResponse,msg=string} "获取成功"
// @Router /baby/profile/list [get]
func (b *BabyProfileApi) GetBabyProfileList(c *gin.Context) {
	var req request.BabyProfileSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 设置默认分页参数
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 获取当前用户ID
	claims, exist := c.Get("claims")
	if !exist {
		response.FailWithMessage("未登录或非法访问", c)
		return
	}

	customClaims, ok := claims.(*systemReq.CustomClaims)
	if !ok {
		response.FailWithMessage("token解析失败", c)
		return
	}

	list, err := babyProfileService.GetBabyProfileList(customClaims.BaseClaims.ID, &req)
	if err != nil {
		global.GVA_LOG.Error("获取宝宝档案列表失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(list, "获取成功", c)
}

// UpdateBabyProfile 更新宝宝档案
// @Tags BabyProfile
// @Summary 更新宝宝档案
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UpdateBabyProfileRequest true "宝宝档案信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /baby/profile [put]
func (b *BabyProfileApi) UpdateBabyProfile(c *gin.Context) {
	var req request.UpdateBabyProfileRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取当前用户ID
	claims, exist := c.Get("claims")
	if !exist {
		response.FailWithMessage("未登录或非法访问", c)
		return
	}

	customClaims, ok := claims.(*systemReq.CustomClaims)
	if !ok {
		response.FailWithMessage("token解析失败", c)
		return
	}

	err = babyProfileService.UpdateBabyProfile(customClaims.BaseClaims.ID, &req)
	if err != nil {
		global.GVA_LOG.Error("更新宝宝档案失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

// DeleteBabyProfile 删除宝宝档案
// @Tags BabyProfile
// @Summary 删除宝宝档案
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id path int true "宝宝档案ID"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /baby/profile/{id} [delete]
func (b *BabyProfileApi) DeleteBabyProfile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("ID格式错误", c)
		return
	}

	// 获取当前用户ID
	claims, exist := c.Get("claims")
	if !exist {
		response.FailWithMessage("未登录或非法访问", c)
		return
	}

	customClaims, ok := claims.(*systemReq.CustomClaims)
	if !ok {
		response.FailWithMessage("token解析失败", c)
		return
	}

	err = babyProfileService.DeleteBabyProfile(uint(id), customClaims.BaseClaims.ID)
	if err != nil {
		global.GVA_LOG.Error("删除宝宝档案失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// SetActiveBaby 设置活跃宝宝
// @Tags BabyProfile
// @Summary 设置活跃宝宝
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id path int true "宝宝档案ID"
// @Success 200 {object} response.Response{msg=string} "设置成功"
// @Router /baby/profile/{id}/setActive [put]
func (b *BabyProfileApi) SetActiveBaby(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("ID格式错误", c)
		return
	}

	// 获取当前用户ID
	claims, exist := c.Get("claims")
	if !exist {
		response.FailWithMessage("未登录或非法访问", c)
		return
	}

	customClaims, ok := claims.(*systemReq.CustomClaims)
	if !ok {
		response.FailWithMessage("token解析失败", c)
		return
	}

	err = babyProfileService.SetActiveBaby(uint(id), customClaims.BaseClaims.ID)
	if err != nil {
		global.GVA_LOG.Error("设置活跃宝宝失败!", zap.Error(err))
		response.FailWithMessage("设置失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("设置成功", c)
}
