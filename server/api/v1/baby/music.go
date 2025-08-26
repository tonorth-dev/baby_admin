package baby

import (
	"baby_admin/server/global"
	"baby_admin/server/model/baby/request"
	commonReq "baby_admin/server/model/common/request"
	"baby_admin/server/model/common/response"
	systemReq "baby_admin/server/model/system/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type MusicApi struct{}

// GetMusicCategories 获取音乐分类列表
// @Tags Music
// @Summary 获取音乐分类列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=[]response.MusicCategoryResponse,msg=string} "获取成功"
// @Router /music/categories [get]
func (m *MusicApi) GetMusicCategories(c *gin.Context) {
	categories, err := musicService.GetMusicCategories()
	if err != nil {
		global.GVA_LOG.Error("获取音乐分类失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(categories, "获取成功", c)
}

// GetMusicList 获取音乐列表
// @Tags Music
// @Summary 获取音乐列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.MusicSearch true "搜索参数"
// @Success 200 {object} response.Response{data=response.MusicListResponse,msg=string} "获取成功"
// @Router /music/list [get]
func (m *MusicApi) GetMusicList(c *gin.Context) {
	var req request.MusicSearch
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
		req.PageSize = 20
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

	list, err := musicService.GetMusicList(customClaims.BaseClaims.ID, &req)
	if err != nil {
		global.GVA_LOG.Error("获取音乐列表失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(list, "获取成功", c)
}

// GetMusicDetail 获取音乐详情
// @Tags Music
// @Summary 获取音乐详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id path int true "音乐ID"
// @Success 200 {object} response.Response{data=response.MusicResponse,msg=string} "获取成功"
// @Router /music/{id} [get]
func (m *MusicApi) GetMusicDetail(c *gin.Context) {
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

	music, err := musicService.GetMusicDetail(customClaims.BaseClaims.ID, uint(id))
	if err != nil {
		global.GVA_LOG.Error("获取音乐详情失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(music, "获取成功", c)
}

// PlayMusic 播放音乐
// @Tags Music
// @Summary 播放音乐
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PlayMusicRequest true "播放信息"
// @Success 200 {object} response.Response{msg=string} "播放成功"
// @Router /music/play [post]
func (m *MusicApi) PlayMusic(c *gin.Context) {
	var req request.PlayMusicRequest
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

	err = musicService.PlayMusic(customClaims.BaseClaims.ID, &req)
	if err != nil {
		global.GVA_LOG.Error("播放音乐失败!", zap.Error(err))
		response.FailWithMessage("播放失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("播放成功", c)
}

// ToggleFavorite 切换收藏状态
// @Tags Music
// @Summary 切换音乐收藏状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id path int true "音乐ID"
// @Success 200 {object} response.Response{data=bool,msg=string} "操作成功"
// @Router /music/{id}/favorite [post]
func (m *MusicApi) ToggleFavorite(c *gin.Context) {
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

	isFavorited, err := musicService.ToggleFavorite(customClaims.BaseClaims.ID, uint(id))
	if err != nil {
		global.GVA_LOG.Error("切换收藏状态失败!", zap.Error(err))
		response.FailWithMessage("操作失败: "+err.Error(), c)
		return
	}

	message := "已取消收藏"
	if isFavorited {
		message = "已添加收藏"
	}

	response.OkWithDetailed(isFavorited, message, c)
}

// GetUserFavorites 获取用户收藏的音乐
// @Tags Music
// @Summary 获取用户收藏的音乐
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query commonReq.PageInfo true "分页参数"
// @Success 200 {object} response.Response{data=response.MusicListResponse,msg=string} "获取成功"
// @Router /music/favorites [get]
func (m *MusicApi) GetUserFavorites(c *gin.Context) {
	var req commonReq.PageInfo
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
		req.PageSize = 20
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

	list, err := musicService.GetUserFavorites(customClaims.BaseClaims.ID, &req)
	if err != nil {
		global.GVA_LOG.Error("获取收藏音乐失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(list, "获取成功", c)
}

// GetPlayHistory 获取播放历史
// @Tags Music
// @Summary 获取播放历史
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query commonReq.PageInfo true "分页参数"
// @Success 200 {object} response.Response{data=[]response.MusicHistoryResponse,msg=string} "获取成功"
// @Router /music/history [get]
func (m *MusicApi) GetPlayHistory(c *gin.Context) {
	var req commonReq.PageInfo
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
		req.PageSize = 20
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

	history, err := musicService.GetPlayHistory(customClaims.BaseClaims.ID, &req)
	if err != nil {
		global.GVA_LOG.Error("获取播放历史失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(history, "获取成功", c)
}

// GetRecommendations 获取音乐推荐
// @Tags Music
// @Summary 获取音乐推荐
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param baby_id query int false "宝宝ID"
// @Success 200 {object} response.Response{data=[]response.RecommendationResponse,msg=string} "获取成功"
// @Router /music/recommendations [get]
func (m *MusicApi) GetRecommendations(c *gin.Context) {
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

	// 获取宝宝ID参数
	var babyID uint
	if babyIDStr := c.Query("baby_id"); babyIDStr != "" {
		if id, err := strconv.ParseUint(babyIDStr, 10, 32); err == nil {
			babyID = uint(id)
		}
	}

	recommendations, err := musicService.GetRecommendations(customClaims.BaseClaims.ID, babyID)
	if err != nil {
		global.GVA_LOG.Error("获取音乐推荐失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(recommendations, "获取成功", c)
}
