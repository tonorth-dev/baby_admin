package baby

import (
	"time"
	"baby_admin/server/global"
)

// ParentingCategory 育儿指导分类表
type ParentingCategory struct {
	global.GVA_MODEL
	Name        string `json:"name" gorm:"size:50;not null;comment:分类名称"`
	Description string `json:"description" gorm:"size:255;comment:分类描述"`
	Icon        string `json:"icon" gorm:"size:255;comment:分类图标"`
	AgeRange    string `json:"age_range" gorm:"size:20;comment:适用年龄范围"`
	SortOrder   int    `json:"sort_order" gorm:"default:0;comment:排序权重"`
	IsActive    bool   `json:"is_active" gorm:"default:true;comment:是否启用"`
}

// TableName 指定表名
func (ParentingCategory) TableName() string {
	return "parenting_categories"
}

// ParentingArticle 育儿指导文章表
type ParentingArticle struct {
	global.GVA_MODEL
	CategoryID   uint      `json:"category_id" gorm:"not null;comment:分类ID"`
	Title        string    `json:"title" gorm:"size:200;not null;comment:文章标题"`
	Summary      string    `json:"summary" gorm:"size:500;comment:文章摘要"`
	Content      string    `json:"content" gorm:"type:longtext;comment:文章内容"`
	CoverURL     string    `json:"cover_url" gorm:"size:500;comment:封面图片URL"`
	Author       string    `json:"author" gorm:"size:50;comment:作者"`
	AuthorAvatar string    `json:"author_avatar" gorm:"size:500;comment:作者头像"`
	AuthorTitle  string    `json:"author_title" gorm:"size:100;comment:作者职称"`
	AgeRange     string    `json:"age_range" gorm:"size:20;comment:适用年龄范围"`
	Tags         string    `json:"tags" gorm:"size:255;comment:标签,JSON格式"`
	ViewCount    int64     `json:"view_count" gorm:"default:0;comment:浏览次数"`
	LikeCount    int64     `json:"like_count" gorm:"default:0;comment:点赞次数"`
	ShareCount   int64     `json:"share_count" gorm:"default:0;comment:分享次数"`
	ReadTime     int       `json:"read_time" gorm:"comment:预计阅读时间(分钟)"`
	Difficulty   int       `json:"difficulty" gorm:"default:1;comment:难度等级:1简单,2中等,3困难"`
	IsRecommend  bool      `json:"is_recommend" gorm:"default:false;comment:是否推荐"`
	IsVIP        bool      `json:"is_vip" gorm:"default:false;comment:是否VIP专享"`
	IsActive     bool      `json:"is_active" gorm:"default:true;comment:是否启用"`
	PublishedAt  time.Time `json:"published_at" gorm:"comment:发布时间"`
	SortOrder    int       `json:"sort_order" gorm:"default:0;comment:排序权重"`
}

// TableName 指定表名
func (ParentingArticle) TableName() string {
	return "parenting_articles"
}

// ParentingVideo 育儿指导视频表
type ParentingVideo struct {
	global.GVA_MODEL
	CategoryID   uint   `json:"category_id" gorm:"not null;comment:分类ID"`
	Title        string `json:"title" gorm:"size:200;not null;comment:视频标题"`
	Description  string `json:"description" gorm:"type:text;comment:视频描述"`
	VideoURL     string `json:"video_url" gorm:"size:500;not null;comment:视频文件URL"`
	CoverURL     string `json:"cover_url" gorm:"size:500;comment:封面图片URL"`
	Duration     int    `json:"duration" gorm:"comment:视频时长(秒)"`
	FileSize     int64  `json:"file_size" gorm:"comment:文件大小(字节)"`
	Resolution   string `json:"resolution" gorm:"size:20;comment:分辨率"`
	Author       string `json:"author" gorm:"size:50;comment:作者"`
	AuthorAvatar string `json:"author_avatar" gorm:"size:500;comment:作者头像"`
	AuthorTitle  string `json:"author_title" gorm:"size:100;comment:作者职称"`
	AgeRange     string `json:"age_range" gorm:"size:20;comment:适用年龄范围"`
	Tags         string `json:"tags" gorm:"size:255;comment:标签,JSON格式"`
	ViewCount    int64  `json:"view_count" gorm:"default:0;comment:观看次数"`
	LikeCount    int64  `json:"like_count" gorm:"default:0;comment:点赞次数"`
	ShareCount   int64  `json:"share_count" gorm:"default:0;comment:分享次数"`
	IsRecommend  bool   `json:"is_recommend" gorm:"default:false;comment:是否推荐"`
	IsVIP        bool   `json:"is_vip" gorm:"default:false;comment:是否VIP专享"`
	IsActive     bool   `json:"is_active" gorm:"default:true;comment:是否启用"`
	PublishedAt  time.Time `json:"published_at" gorm:"comment:发布时间"`
	SortOrder    int    `json:"sort_order" gorm:"default:0;comment:排序权重"`
}

// TableName 指定表名
func (ParentingVideo) TableName() string {
	return "parenting_videos"
}

// ParentingMilestone 成长里程碑表
type ParentingMilestone struct {
	global.GVA_MODEL
	AgeRange    string `json:"age_range" gorm:"size:20;not null;comment:年龄范围"`
	Category    string `json:"category" gorm:"size:50;not null;comment:发育类别"`
	Title       string `json:"title" gorm:"size:100;not null;comment:里程碑标题"`
	Description string `json:"description" gorm:"type:text;comment:描述"`
	Tips        string `json:"tips" gorm:"type:text;comment:小贴士"`
	IsImportant bool   `json:"is_important" gorm:"default:false;comment:是否重要里程碑"`
	SortOrder   int    `json:"sort_order" gorm:"default:0;comment:排序权重"`
	IsActive    bool   `json:"is_active" gorm:"default:true;comment:是否启用"`
}

// TableName 指定表名
func (ParentingMilestone) TableName() string {
	return "parenting_milestones"
}

// UserArticleRead 用户文章阅读记录表
type UserArticleRead struct {
	global.GVA_MODEL
	UserID    uint      `json:"user_id" gorm:"not null;comment:用户ID"`
	ArticleID uint      `json:"article_id" gorm:"not null;comment:文章ID"`
	ReadTime  int       `json:"read_time" gorm:"default:0;comment:阅读时长(秒)"`
	Progress  int       `json:"progress" gorm:"default:0;comment:阅读进度(百分比)"`
	IsFinished bool     `json:"is_finished" gorm:"default:false;comment:是否读完"`
	LastReadAt time.Time `json:"last_read_at" gorm:"comment:最后阅读时间"`
}

// TableName 指定表名
func (UserArticleRead) TableName() string {
	return "user_article_reads"
}

// UserVideoWatch 用户视频观看记录表
type UserVideoWatch struct {
	global.GVA_MODEL
	UserID      uint      `json:"user_id" gorm:"not null;comment:用户ID"`
	VideoID     uint      `json:"video_id" gorm:"not null;comment:视频ID"`
	WatchTime   int       `json:"watch_time" gorm:"default:0;comment:观看时长(秒)"`
	Progress    int       `json:"progress" gorm:"default:0;comment:观看进度(百分比)"`
	IsFinished  bool      `json:"is_finished" gorm:"default:false;comment:是否看完"`
	LastWatchAt time.Time `json:"last_watch_at" gorm:"comment:最后观看时间"`
}

// TableName 指定表名
func (UserVideoWatch) TableName() string {
	return "user_video_watches"
}

// UserContentFavorite 用户内容收藏表
type UserContentFavorite struct {
	global.GVA_MODEL
	UserID      uint `json:"user_id" gorm:"not null;comment:用户ID"`
	ContentType int  `json:"content_type" gorm:"not null;comment:内容类型:1文章,2视频"`
	ContentID   uint `json:"content_id" gorm:"not null;comment:内容ID"`
}

// TableName 指定表名
func (UserContentFavorite) TableName() string {
	return "user_content_favorites"
}

// GetDifficultyText 获取难度等级文本
func (p *ParentingArticle) GetDifficultyText() string {
	switch p.Difficulty {
	case 1:
		return "简单"
	case 2:
		return "中等"
	case 3:
		return "困难"
	default:
		return "未知"
	}
}