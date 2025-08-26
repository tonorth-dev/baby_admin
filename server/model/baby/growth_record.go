package baby

import (
	"time"
	"baby_admin/server/global"
)

// GrowthRecord 成长记录表
type GrowthRecord struct {
	global.GVA_MODEL
	UserID     uint      `json:"user_id" gorm:"not null;comment:用户ID"`
	BabyID     uint      `json:"baby_id" gorm:"not null;comment:宝宝ID"`
	Title      string    `json:"title" gorm:"size:100;not null;comment:记录标题"`
	Content    string    `json:"content" gorm:"type:text;comment:记录内容"`
	RecordType int       `json:"record_type" gorm:"not null;default:1;comment:记录类型:1文字,2图片,3视频,4里程碑"`
	RecordDate time.Time `json:"record_date" gorm:"not null;comment:记录日期"`
	Tags       string    `json:"tags" gorm:"size:255;comment:标签,逗号分隔"`
	Weight     float64   `json:"weight" gorm:"type:decimal(5,2);comment:体重记录(kg)"`
	Height     float64   `json:"height" gorm:"type:decimal(5,2);comment:身高记录(cm)"`
	MediaFiles string    `json:"media_files" gorm:"type:text;comment:媒体文件URL,JSON格式"`
	Milestone  string    `json:"milestone" gorm:"size:100;comment:里程碑类型"`
	IsPrivate  bool      `json:"is_private" gorm:"default:false;comment:是否私密"`
}

// TableName 指定表名
func (GrowthRecord) TableName() string {
	return "growth_records"
}

// MediaFile 媒体文件结构
type MediaFile struct {
	URL      string `json:"url"`
	Type     string `json:"type"`     // image, video
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
}

// GetRecordTypeText 获取记录类型文本
func (g *GrowthRecord) GetRecordTypeText() string {
	switch g.RecordType {
	case 1:
		return "文字记录"
	case 2:
		return "图片记录"
	case 3:
		return "视频记录"
	case 4:
		return "里程碑"
	default:
		return "未知类型"
	}
}