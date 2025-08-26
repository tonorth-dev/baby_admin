package request

import "baby_admin/server/model/common/request"

// MusicSearch 音乐搜索条件
type MusicSearch struct {
	request.PageInfo
	CategoryID uint   `json:"category_id" form:"category_id"`
	AgeRange   string `json:"age_range" form:"age_range"`
	Keyword    string `json:"keyword" form:"keyword"`
	Tags       string `json:"tags" form:"tags"`
	IsVIP      *bool  `json:"is_vip" form:"is_vip"`
}

// PlayMusicRequest 播放音乐请求
type PlayMusicRequest struct {
	MusicID  uint `json:"music_id" binding:"required"`
	BabyID   uint `json:"baby_id"`
	PlayTime int  `json:"play_time"` // 播放时长(秒)
}

// CreatePlaylistRequest 创建播放列表请求
type CreatePlaylistRequest struct {
	Name        string `json:"name" binding:"required,min=1,max=50"`
	Description string `json:"description"`
	CoverURL    string `json:"cover_url"`
	IsPublic    bool   `json:"is_public"`
}

// UpdatePlaylistRequest 更新播放列表请求
type UpdatePlaylistRequest struct {
	ID          uint   `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required,min=1,max=50"`
	Description string `json:"description"`
	CoverURL    string `json:"cover_url"`
	IsPublic    bool   `json:"is_public"`
}

// AddMusicToPlaylistRequest 添加音乐到播放列表请求
type AddMusicToPlaylistRequest struct {
	PlaylistID uint   `json:"playlist_id" binding:"required"`
	MusicIDs   []uint `json:"music_ids" binding:"required,min=1"`
}

// RemoveMusicFromPlaylistRequest 从播放列表移除音乐请求
type RemoveMusicFromPlaylistRequest struct {
	PlaylistID uint   `json:"playlist_id" binding:"required"`
	MusicIDs   []uint `json:"music_ids" binding:"required,min=1"`
}
