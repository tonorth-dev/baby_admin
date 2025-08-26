package request

import "baby_admin/server/model/common/request"

// ParentingArticleSearch 育儿文章搜索条件
type ParentingArticleSearch struct {
	request.PageInfo
	CategoryID  uint   `json:"category_id" form:"category_id"`
	AgeRange    string `json:"age_range" form:"age_range"`
	Keyword     string `json:"keyword" form:"keyword"`
	Tags        string `json:"tags" form:"tags"`
	Difficulty  int    `json:"difficulty" form:"difficulty"`
	IsRecommend *bool  `json:"is_recommend" form:"is_recommend"`
	IsVIP       *bool  `json:"is_vip" form:"is_vip"`
}

// ParentingVideoSearch 育儿视频搜索条件
type ParentingVideoSearch struct {
	request.PageInfo
	CategoryID  uint   `json:"category_id" form:"category_id"`
	AgeRange    string `json:"age_range" form:"age_range"`
	Keyword     string `json:"keyword" form:"keyword"`
	Tags        string `json:"tags" form:"tags"`
	IsRecommend *bool  `json:"is_recommend" form:"is_recommend"`
	IsVIP       *bool  `json:"is_vip" form:"is_vip"`
}

// ParentingMilestoneSearch 成长里程碑搜索条件
type ParentingMilestoneSearch struct {
	request.PageInfo
	AgeRange    string `json:"age_range" form:"age_range"`
	Category    string `json:"category" form:"category"`
	IsImportant *bool  `json:"is_important" form:"is_important"`
}

// ReadArticleRequest 阅读文章请求
type ReadArticleRequest struct {
	ArticleID  uint `json:"article_id" binding:"required"`
	ReadTime   int  `json:"read_time"`   // 阅读时长(秒)
	Progress   int  `json:"progress"`    // 阅读进度(百分比)
	IsFinished bool `json:"is_finished"` // 是否读完
}

// WatchVideoRequest 观看视频请求
type WatchVideoRequest struct {
	VideoID    uint `json:"video_id" binding:"required"`
	WatchTime  int  `json:"watch_time"`  // 观看时长(秒)
	Progress   int  `json:"progress"`    // 观看进度(百分比)
	IsFinished bool `json:"is_finished"` // 是否看完
}

// ToggleFavoriteRequest 切换收藏请求
type ToggleFavoriteRequest struct {
	ContentType int  `json:"content_type" binding:"required,oneof=1 2"` // 1文章,2视频
	ContentID   uint `json:"content_id" binding:"required"`
}

// GetRecommendationsRequest 获取推荐内容请求
type GetRecommendationsRequest struct {
	BabyID uint `json:"baby_id" form:"baby_id"` // 宝宝ID，用于基于年龄的推荐
	Limit  int  `json:"limit" form:"limit"`     // 推荐数量限制
}
