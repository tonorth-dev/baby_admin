package response

import (
	"baby_admin/server/model/baby"
	"time"
)

// MusicCategoryResponse 音乐分类响应
type MusicCategoryResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	AgeRange    string `json:"age_range"`
	SortOrder   int    `json:"sort_order"`
	MusicCount  int64  `json:"music_count"` // 该分类下的音乐数量
	IsActive    bool   `json:"is_active"`
}

// MusicResponse 音乐响应
type MusicResponse struct {
	ID           uint      `json:"id"`
	CategoryID   uint      `json:"category_id"`
	CategoryName string    `json:"category_name"`
	Title        string    `json:"title"`
	Artist       string    `json:"artist"`
	Description  string    `json:"description"`
	AudioURL     string    `json:"audio_url"`
	CoverURL     string    `json:"cover_url"`
	Duration     int       `json:"duration"`
	DurationText string    `json:"duration_text"` // 格式化的时长显示
	AgeRange     string    `json:"age_range"`
	Tags         []string  `json:"tags"`
	PlayCount    int64     `json:"play_count"`
	LikeCount    int64     `json:"like_count"`
	FileSize     int64     `json:"file_size"`
	Format       string    `json:"format"`
	IsVIP        bool      `json:"is_vip"`
	IsActive     bool      `json:"is_active"`
	IsFavorited  bool      `json:"is_favorited"` // 用户是否收藏
	CreatedAt    time.Time `json:"created_at"`
}

// FromMusic 从Music模型转换
func (m *MusicResponse) FromMusic(music *baby.Music, categoryName string, isFavorited bool) {
	m.ID = music.ID
	m.CategoryID = music.CategoryID
	m.CategoryName = categoryName
	m.Title = music.Title
	m.Artist = music.Artist
	m.Description = music.Description
	m.AudioURL = music.AudioURL
	m.CoverURL = music.CoverURL
	m.Duration = music.Duration
	m.DurationText = formatDuration(music.Duration)
	m.AgeRange = music.AgeRange
	m.Tags = parseTags(music.Tags)
	m.PlayCount = music.PlayCount
	m.LikeCount = music.LikeCount
	m.FileSize = music.FileSize
	m.Format = music.Format
	m.IsVIP = music.IsVIP
	m.IsActive = music.IsActive
	m.IsFavorited = isFavorited
	m.CreatedAt = music.CreatedAt
}

// MusicListResponse 音乐列表响应
type MusicListResponse struct {
	List     []MusicResponse `json:"list"`
	Total    int64           `json:"total"`
	Page     int             `json:"page"`
	PageSize int             `json:"page_size"`
}

// PlaylistResponse 播放列表响应
type PlaylistResponse struct {
	ID          uint            `json:"id"`
	UserID      uint            `json:"user_id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	CoverURL    string          `json:"cover_url"`
	IsPublic    bool            `json:"is_public"`
	MusicCount  int             `json:"music_count"`
	Musics      []MusicResponse `json:"musics,omitempty"` // 播放列表中的音乐
	Duration    int             `json:"duration"`         // 总时长
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

// PlaylistListResponse 播放列表列表响应
type PlaylistListResponse struct {
	List     []PlaylistResponse `json:"list"`
	Total    int64              `json:"total"`
	Page     int                `json:"page"`
	PageSize int                `json:"page_size"`
}

// MusicHistoryResponse 音乐播放历史响应
type MusicHistoryResponse struct {
	ID         uint          `json:"id"`
	Music      MusicResponse `json:"music"`
	BabyID     uint          `json:"baby_id"`
	BabyName   string        `json:"baby_name"`
	PlayTime   int           `json:"play_time"`
	IsFinished bool          `json:"is_finished"`
	PlayedAt   time.Time     `json:"played_at"`
}

// RecommendationResponse 音乐推荐响应
type RecommendationResponse struct {
	Title       string          `json:"title"`       // 推荐标题
	Description string          `json:"description"` // 推荐描述
	Musics      []MusicResponse `json:"musics"`      // 推荐音乐列表
	Type        string          `json:"type"`        // 推荐类型：age_based, sleep, popular等
}
