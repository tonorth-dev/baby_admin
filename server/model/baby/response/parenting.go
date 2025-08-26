package response

import (
	"baby_admin/server/model/baby"
	"fmt"
	"time"
)

// ParentingCategoryResponse 育儿分类响应
type ParentingCategoryResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Icon         string `json:"icon"`
	AgeRange     string `json:"age_range"`
	SortOrder    int    `json:"sort_order"`
	ArticleCount int64  `json:"article_count"` // 该分类下的文章数量
	VideoCount   int64  `json:"video_count"`   // 该分类下的视频数量
	IsActive     bool   `json:"is_active"`
}

// ParentingArticleResponse 育儿文章响应
type ParentingArticleResponse struct {
	ID             uint      `json:"id"`
	CategoryID     uint      `json:"category_id"`
	CategoryName   string    `json:"category_name"`
	Title          string    `json:"title"`
	Summary        string    `json:"summary"`
	Content        string    `json:"content,omitempty"` // 列表时不返回内容
	CoverURL       string    `json:"cover_url"`
	Author         string    `json:"author"`
	AuthorAvatar   string    `json:"author_avatar"`
	AuthorTitle    string    `json:"author_title"`
	AgeRange       string    `json:"age_range"`
	Tags           []string  `json:"tags"`
	ViewCount      int64     `json:"view_count"`
	LikeCount      int64     `json:"like_count"`
	ShareCount     int64     `json:"share_count"`
	ReadTime       int       `json:"read_time"`
	ReadTimeText   string    `json:"read_time_text"`
	Difficulty     int       `json:"difficulty"`
	DifficultyText string    `json:"difficulty_text"`
	IsRecommend    bool      `json:"is_recommend"`
	IsVIP          bool      `json:"is_vip"`
	IsFavorited    bool      `json:"is_favorited"`  // 用户是否收藏
	ReadProgress   int       `json:"read_progress"` // 用户阅读进度
	IsRead         bool      `json:"is_read"`       // 用户是否已读
	PublishedAt    time.Time `json:"published_at"`
	CreatedAt      time.Time `json:"created_at"`
}

// FromParentingArticle 从ParentingArticle模型转换
func (p *ParentingArticleResponse) FromParentingArticle(article *baby.ParentingArticle, categoryName string, isFavorited bool, readProgress int, isRead bool, includeContent bool) {
	p.ID = article.ID
	p.CategoryID = article.CategoryID
	p.CategoryName = categoryName
	p.Title = article.Title
	p.Summary = article.Summary
	if includeContent {
		p.Content = article.Content
	}
	p.CoverURL = article.CoverURL
	p.Author = article.Author
	p.AuthorAvatar = article.AuthorAvatar
	p.AuthorTitle = article.AuthorTitle
	p.AgeRange = article.AgeRange
	p.Tags = parseTags(article.Tags)
	p.ViewCount = article.ViewCount
	p.LikeCount = article.LikeCount
	p.ShareCount = article.ShareCount
	p.ReadTime = article.ReadTime
	p.ReadTimeText = formatReadTime(article.ReadTime)
	p.Difficulty = article.Difficulty
	p.DifficultyText = article.GetDifficultyText()
	p.IsRecommend = article.IsRecommend
	p.IsVIP = article.IsVIP
	p.IsFavorited = isFavorited
	p.ReadProgress = readProgress
	p.IsRead = isRead
	p.PublishedAt = article.PublishedAt
	p.CreatedAt = article.CreatedAt
}

// ParentingVideoResponse 育儿视频响应
type ParentingVideoResponse struct {
	ID            uint      `json:"id"`
	CategoryID    uint      `json:"category_id"`
	CategoryName  string    `json:"category_name"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	VideoURL      string    `json:"video_url"`
	CoverURL      string    `json:"cover_url"`
	Duration      int       `json:"duration"`
	DurationText  string    `json:"duration_text"`
	FileSize      int64     `json:"file_size"`
	Resolution    string    `json:"resolution"`
	Author        string    `json:"author"`
	AuthorAvatar  string    `json:"author_avatar"`
	AuthorTitle   string    `json:"author_title"`
	AgeRange      string    `json:"age_range"`
	Tags          []string  `json:"tags"`
	ViewCount     int64     `json:"view_count"`
	LikeCount     int64     `json:"like_count"`
	ShareCount    int64     `json:"share_count"`
	IsRecommend   bool      `json:"is_recommend"`
	IsVIP         bool      `json:"is_vip"`
	IsFavorited   bool      `json:"is_favorited"`   // 用户是否收藏
	WatchProgress int       `json:"watch_progress"` // 用户观看进度
	IsWatched     bool      `json:"is_watched"`     // 用户是否已观看
	PublishedAt   time.Time `json:"published_at"`
	CreatedAt     time.Time `json:"created_at"`
}

// FromParentingVideo 从ParentingVideo模型转换
func (p *ParentingVideoResponse) FromParentingVideo(video *baby.ParentingVideo, categoryName string, isFavorited bool, watchProgress int, isWatched bool) {
	p.ID = video.ID
	p.CategoryID = video.CategoryID
	p.CategoryName = categoryName
	p.Title = video.Title
	p.Description = video.Description
	p.VideoURL = video.VideoURL
	p.CoverURL = video.CoverURL
	p.Duration = video.Duration
	p.DurationText = formatDuration(video.Duration)
	p.FileSize = video.FileSize
	p.Resolution = video.Resolution
	p.Author = video.Author
	p.AuthorAvatar = video.AuthorAvatar
	p.AuthorTitle = video.AuthorTitle
	p.AgeRange = video.AgeRange
	p.Tags = parseTags(video.Tags)
	p.ViewCount = video.ViewCount
	p.LikeCount = video.LikeCount
	p.ShareCount = video.ShareCount
	p.IsRecommend = video.IsRecommend
	p.IsVIP = video.IsVIP
	p.IsFavorited = isFavorited
	p.WatchProgress = watchProgress
	p.IsWatched = isWatched
	p.PublishedAt = video.PublishedAt
	p.CreatedAt = video.CreatedAt
}

// ParentingMilestoneResponse 成长里程碑响应
type ParentingMilestoneResponse struct {
	ID          uint      `json:"id"`
	AgeRange    string    `json:"age_range"`
	Category    string    `json:"category"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Tips        string    `json:"tips"`
	IsImportant bool      `json:"is_important"`
	SortOrder   int       `json:"sort_order"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
}

// ParentingArticleListResponse 育儿文章列表响应
type ParentingArticleListResponse struct {
	List     []ParentingArticleResponse `json:"list"`
	Total    int64                      `json:"total"`
	Page     int                        `json:"page"`
	PageSize int                        `json:"page_size"`
}

// ParentingVideoListResponse 育儿视频列表响应
type ParentingVideoListResponse struct {
	List     []ParentingVideoResponse `json:"list"`
	Total    int64                    `json:"total"`
	Page     int                      `json:"page"`
	PageSize int                      `json:"page_size"`
}

// ParentingMilestoneListResponse 成长里程碑列表响应
type ParentingMilestoneListResponse struct {
	List     []ParentingMilestoneResponse `json:"list"`
	Total    int64                        `json:"total"`
	Page     int                          `json:"page"`
	PageSize int                          `json:"page_size"`
}

// ParentingRecommendationResponse 育儿推荐响应
type ParentingRecommendationResponse struct {
	Title       string                     `json:"title"`       // 推荐标题
	Description string                     `json:"description"` // 推荐描述
	Articles    []ParentingArticleResponse `json:"articles"`    // 推荐文章
	Videos      []ParentingVideoResponse   `json:"videos"`      // 推荐视频
	Type        string                     `json:"type"`        // 推荐类型
}

// UserLearningProgress 用户学习进度统计
type UserLearningProgress struct {
	TotalArticles   int64   `json:"total_articles"`   // 总文章数
	ReadArticles    int64   `json:"read_articles"`    // 已读文章数
	TotalVideos     int64   `json:"total_videos"`     // 总视频数
	WatchedVideos   int64   `json:"watched_videos"`   // 已观看视频数
	TotalReadTime   int     `json:"total_read_time"`  // 总阅读时间(分钟)
	TotalWatchTime  int     `json:"total_watch_time"` // 总观看时间(分钟)
	ArticleProgress float64 `json:"article_progress"` // 文章学习进度
	VideoProgress   float64 `json:"video_progress"`   // 视频学习进度
	OverallProgress float64 `json:"overall_progress"` // 总体学习进度
	WeeklyGoal      int     `json:"weekly_goal"`      // 周学习目标(分钟)
	WeeklyCompleted int     `json:"weekly_completed"` // 本周已完成(分钟)
	ConsecutiveDays int     `json:"consecutive_days"` // 连续学习天数
}

// formatReadTime 格式化阅读时长
func formatReadTime(minutes int) string {
	if minutes <= 0 {
		return "< 1分钟"
	} else if minutes < 60 {
		return fmt.Sprintf("%d分钟", minutes)
	} else {
		hours := minutes / 60
		remainingMinutes := minutes % 60
		if remainingMinutes == 0 {
			return fmt.Sprintf("%d小时", hours)
		}
		return fmt.Sprintf("%d小时%d分钟", hours, remainingMinutes)
	}
}
