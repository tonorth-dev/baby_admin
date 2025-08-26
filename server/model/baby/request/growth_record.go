package request

import (
	"baby_admin/server/model/common/request"
	"time"
)

// GrowthRecordSearch 成长记录搜索条件
type GrowthRecordSearch struct {
	request.PageInfo
	UserID     uint   `json:"user_id" form:"user_id"`
	BabyID     uint   `json:"baby_id" form:"baby_id"`
	RecordType int    `json:"record_type" form:"record_type"`
	StartDate  string `json:"start_date" form:"start_date"`
	EndDate    string `json:"end_date" form:"end_date"`
	Keyword    string `json:"keyword" form:"keyword"`
	Tags       string `json:"tags" form:"tags"`
}

// CreateGrowthRecordRequest 创建成长记录请求
type CreateGrowthRecordRequest struct {
	BabyID     uint               `json:"baby_id" binding:"required"`
	Title      string             `json:"title" binding:"required" validate:"min=1,max=100"`
	Content    string             `json:"content"`
	RecordType int                `json:"record_type" binding:"required,oneof=1 2 3 4"`
	RecordDate time.Time          `json:"record_date" binding:"required"`
	Tags       string             `json:"tags"`
	Weight     float64            `json:"weight"`
	Height     float64            `json:"height"`
	MediaFiles []MediaFileRequest `json:"media_files"`
	Milestone  string             `json:"milestone"`
	IsPrivate  bool               `json:"is_private"`
}

// UpdateGrowthRecordRequest 更新成长记录请求
type UpdateGrowthRecordRequest struct {
	ID         uint               `json:"id" binding:"required"`
	BabyID     uint               `json:"baby_id" binding:"required"`
	Title      string             `json:"title" binding:"required" validate:"min=1,max=100"`
	Content    string             `json:"content"`
	RecordType int                `json:"record_type" binding:"required,oneof=1 2 3 4"`
	RecordDate time.Time          `json:"record_date" binding:"required"`
	Tags       string             `json:"tags"`
	Weight     float64            `json:"weight"`
	Height     float64            `json:"height"`
	MediaFiles []MediaFileRequest `json:"media_files"`
	Milestone  string             `json:"milestone"`
	IsPrivate  bool               `json:"is_private"`
}

// MediaFileRequest 媒体文件请求
type MediaFileRequest struct {
	URL      string `json:"url" binding:"required"`
	Type     string `json:"type" binding:"required,oneof=image video"`
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
}
