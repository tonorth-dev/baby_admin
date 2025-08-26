package baby

import (
	"baby_admin/server/global"
)

// MusicCategory 音乐分类表
type MusicCategory struct {
	global.GVA_MODEL
	Name        string `json:"name" gorm:"size:50;not null;comment:分类名称"`
	Description string `json:"description" gorm:"size:255;comment:分类描述"`
	Icon        string `json:"icon" gorm:"size:255;comment:分类图标"`
	AgeRange    string `json:"age_range" gorm:"size:20;comment:适用年龄范围"`
	SortOrder   int    `json:"sort_order" gorm:"default:0;comment:排序权重"`
	IsActive    bool   `json:"is_active" gorm:"default:true;comment:是否启用"`
}

// TableName 指定表名
func (MusicCategory) TableName() string {
	return "music_categories"
}

// Music 音乐表
type Music struct {
	global.GVA_MODEL
	CategoryID   uint   `json:"category_id" gorm:"not null;comment:分类ID"`
	Title        string `json:"title" gorm:"size:100;not null;comment:音乐标题"`
	Artist       string `json:"artist" gorm:"size:50;comment:艺术家"`
	Description  string `json:"description" gorm:"type:text;comment:音乐描述"`
	AudioURL     string `json:"audio_url" gorm:"size:500;not null;comment:音频文件URL"`
	CoverURL     string `json:"cover_url" gorm:"size:500;comment:封面图片URL"`
	Duration     int    `json:"duration" gorm:"comment:时长(秒)"`
	AgeRange     string `json:"age_range" gorm:"size:20;comment:适用年龄范围"`
	Tags         string `json:"tags" gorm:"size:255;comment:标签,逗号分隔"`
	PlayCount    int64  `json:"play_count" gorm:"default:0;comment:播放次数"`
	LikeCount    int64  `json:"like_count" gorm:"default:0;comment:点赞次数"`
	FileSize     int64  `json:"file_size" gorm:"comment:文件大小(字节)"`
	Format       string `json:"format" gorm:"size:10;comment:音频格式"`
	IsVIP        bool   `json:"is_vip" gorm:"default:false;comment:是否VIP专享"`
	IsActive     bool   `json:"is_active" gorm:"default:true;comment:是否启用"`
	SortOrder    int    `json:"sort_order" gorm:"default:0;comment:排序权重"`
}

// TableName 指定表名
func (Music) TableName() string {
	return "musics"
}

// UserMusicHistory 用户音乐播放历史
type UserMusicHistory struct {
	global.GVA_MODEL
	UserID     uint `json:"user_id" gorm:"not null;comment:用户ID"`
	MusicID    uint `json:"music_id" gorm:"not null;comment:音乐ID"`
	BabyID     uint `json:"baby_id" gorm:"comment:宝宝ID"`
	PlayTime   int  `json:"play_time" gorm:"default:0;comment:播放时长(秒)"`
	IsFinished bool `json:"is_finished" gorm:"default:false;comment:是否播放完成"`
}

// TableName 指定表名
func (UserMusicHistory) TableName() string {
	return "user_music_histories"
}

// UserMusicFavorite 用户音乐收藏
type UserMusicFavorite struct {
	global.GVA_MODEL
	UserID  uint `json:"user_id" gorm:"not null;comment:用户ID"`
	MusicID uint `json:"music_id" gorm:"not null;comment:音乐ID"`
}

// TableName 指定表名
func (UserMusicFavorite) TableName() string {
	return "user_music_favorites"
}

// Playlist 播放列表
type Playlist struct {
	global.GVA_MODEL
	UserID      uint   `json:"user_id" gorm:"not null;comment:用户ID"`
	Name        string `json:"name" gorm:"size:50;not null;comment:播放列表名称"`
	Description string `json:"description" gorm:"size:255;comment:描述"`
	CoverURL    string `json:"cover_url" gorm:"size:500;comment:封面图片"`
	IsPublic    bool   `json:"is_public" gorm:"default:false;comment:是否公开"`
	MusicCount  int    `json:"music_count" gorm:"default:0;comment:音乐数量"`
}

// TableName 指定表名
func (Playlist) TableName() string {
	return "playlists"
}

// PlaylistMusic 播放列表音乐关联
type PlaylistMusic struct {
	global.GVA_MODEL
	PlaylistID uint `json:"playlist_id" gorm:"not null;comment:播放列表ID"`
	MusicID    uint `json:"music_id" gorm:"not null;comment:音乐ID"`
	SortOrder  int  `json:"sort_order" gorm:"default:0;comment:排序权重"`
}

// TableName 指定表名
func (PlaylistMusic) TableName() string {
	return "playlist_musics"
}