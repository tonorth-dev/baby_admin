package response

import (
	"baby_admin/server/model/baby"
	"encoding/json"
	"time"
)

// GrowthRecordResponse 成长记录响应
type GrowthRecordResponse struct {
	ID             uint             `json:"id"`
	UserID         uint             `json:"user_id"`
	BabyID         uint             `json:"baby_id"`
	BabyName       string           `json:"baby_name"` // 关联宝宝姓名
	Title          string           `json:"title"`
	Content        string           `json:"content"`
	RecordType     int              `json:"record_type"`
	RecordTypeText string           `json:"record_type_text"`
	RecordDate     time.Time        `json:"record_date"`
	Tags           []string         `json:"tags"`
	Weight         float64          `json:"weight"`
	Height         float64          `json:"height"`
	MediaFiles     []baby.MediaFile `json:"media_files"`
	Milestone      string           `json:"milestone"`
	IsPrivate      bool             `json:"is_private"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
}

// FromGrowthRecord 从GrowthRecord模型转换
func (r *GrowthRecordResponse) FromGrowthRecord(record *baby.GrowthRecord, babyName string) {
	r.ID = record.ID
	r.UserID = record.UserID
	r.BabyID = record.BabyID
	r.BabyName = babyName
	r.Title = record.Title
	r.Content = record.Content
	r.RecordType = record.RecordType
	r.RecordTypeText = record.GetRecordTypeText()
	r.RecordDate = record.RecordDate
	r.Weight = record.Weight
	r.Height = record.Height
	r.Milestone = record.Milestone
	r.IsPrivate = record.IsPrivate
	r.CreatedAt = record.CreatedAt
	r.UpdatedAt = record.UpdatedAt

	// 解析标签
	if record.Tags != "" {
		r.Tags = parseTags(record.Tags)
	} else {
		r.Tags = []string{}
	}

	// 解析媒体文件
	if record.MediaFiles != "" {
		var mediaFiles []baby.MediaFile
		if err := json.Unmarshal([]byte(record.MediaFiles), &mediaFiles); err == nil {
			r.MediaFiles = mediaFiles
		} else {
			r.MediaFiles = []baby.MediaFile{}
		}
	} else {
		r.MediaFiles = []baby.MediaFile{}
	}
}

// GrowthRecordListResponse 成长记录列表响应
type GrowthRecordListResponse struct {
	List     []GrowthRecordResponse `json:"list"`
	Total    int64                  `json:"total"`
	Page     int                    `json:"page"`
	PageSize int                    `json:"page_size"`
}

// GrowthStatistics 成长统计响应
type GrowthStatistics struct {
	TotalRecords   int64                  `json:"total_records"`   // 总记录数
	PhotoCount     int64                  `json:"photo_count"`     // 照片数量
	VideoCount     int64                  `json:"video_count"`     // 视频数量
	MilestoneCount int64                  `json:"milestone_count"` // 里程碑数量
	RecentRecords  []GrowthRecordResponse `json:"recent_records"`  // 最近记录
	WeightRecords  []WeightRecord         `json:"weight_records"`  // 体重记录
	HeightRecords  []HeightRecord         `json:"height_records"`  // 身高记录
}

// WeightRecord 体重记录
type WeightRecord struct {
	Date   time.Time `json:"date"`
	Weight float64   `json:"weight"`
}

// HeightRecord 身高记录
type HeightRecord struct {
	Date   time.Time `json:"date"`
	Height float64   `json:"height"`
}
