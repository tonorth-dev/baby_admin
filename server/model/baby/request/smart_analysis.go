package request

import (
	"baby_admin/server/model/common/request"
	"time"
)

// SleepRecordSearch 睡眠记录搜索条件
type SleepRecordSearch struct {
	request.PageInfo
	UserID    uint   `json:"user_id" form:"user_id"`
	BabyID    uint   `json:"baby_id" form:"baby_id"`
	StartDate string `json:"start_date" form:"start_date"`
	EndDate   string `json:"end_date" form:"end_date"`
	Status    int    `json:"status" form:"status"`
}

// CreateSleepRecordRequest 创建睡眠记录请求
type CreateSleepRecordRequest struct {
	BabyID      uint       `json:"baby_id" binding:"required"`
	DeviceID    uint       `json:"device_id"`
	StartTime   time.Time  `json:"start_time" binding:"required"`
	EndTime     *time.Time `json:"end_time"`
	Quality     int        `json:"quality"`
	Temperature float64    `json:"temperature"`
	Humidity    float64    `json:"humidity"`
	NoiseLevel  float64    `json:"noise_level"`
	Notes       string     `json:"notes"`
}

// UpdateSleepRecordRequest 更新睡眠记录请求
type UpdateSleepRecordRequest struct {
	ID          uint       `json:"id" binding:"required"`
	EndTime     *time.Time `json:"end_time"`
	Duration    int        `json:"duration"`
	Quality     int        `json:"quality"`
	DeepSleep   int        `json:"deep_sleep"`
	LightSleep  int        `json:"light_sleep"`
	AwakeCount  int        `json:"awake_count"`
	Temperature float64    `json:"temperature"`
	Humidity    float64    `json:"humidity"`
	NoiseLevel  float64    `json:"noise_level"`
	Status      int        `json:"status"`
	Notes       string     `json:"notes"`
}

// CryDetectionSearch 哭声检测搜索条件
type CryDetectionSearch struct {
	request.PageInfo
	UserID      uint   `json:"user_id" form:"user_id"`
	BabyID      uint   `json:"baby_id" form:"baby_id"`
	CryType     int    `json:"cry_type" form:"cry_type"`
	StartDate   string `json:"start_date" form:"start_date"`
	EndDate     string `json:"end_date" form:"end_date"`
	IsProcessed *bool  `json:"is_processed" form:"is_processed"`
}

// CreateCryDetectionRequest 创建哭声检测请求
type CreateCryDetectionRequest struct {
	BabyID     uint      `json:"baby_id" binding:"required"`
	DeviceID   uint      `json:"device_id"`
	DetectedAt time.Time `json:"detected_at" binding:"required"`
	Duration   int       `json:"duration"`
	Intensity  int       `json:"intensity"`
	CryType    int       `json:"cry_type"`
	Confidence float64   `json:"confidence"`
	AudioURL   string    `json:"audio_url"`
	VideoURL   string    `json:"video_url"`
}

// MovementDetectionSearch 动作检测搜索条件
type MovementDetectionSearch struct {
	request.PageInfo
	UserID       uint   `json:"user_id" form:"user_id"`
	BabyID       uint   `json:"baby_id" form:"baby_id"`
	MovementType int    `json:"movement_type" form:"movement_type"`
	StartDate    string `json:"start_date" form:"start_date"`
	EndDate      string `json:"end_date" form:"end_date"`
	IsAlert      *bool  `json:"is_alert" form:"is_alert"`
}

// CreateMovementDetectionRequest 创建动作检测请求
type CreateMovementDetectionRequest struct {
	BabyID       uint      `json:"baby_id" binding:"required"`
	DeviceID     uint      `json:"device_id"`
	DetectedAt   time.Time `json:"detected_at" binding:"required"`
	MovementType int       `json:"movement_type"`
	Intensity    int       `json:"intensity"`
	Duration     int       `json:"duration"`
	VideoURL     string    `json:"video_url"`
	ThumbnailURL string    `json:"thumbnail_url"`
	IsAlert      bool      `json:"is_alert"`
}

// SmartAlertSearch 智能警报搜索条件
type SmartAlertSearch struct {
	request.PageInfo
	UserID      uint  `json:"user_id" form:"user_id"`
	BabyID      uint  `json:"baby_id" form:"baby_id"`
	AlertType   int   `json:"alert_type" form:"alert_type"`
	AlertLevel  int   `json:"alert_level" form:"alert_level"`
	IsRead      *bool `json:"is_read" form:"is_read"`
	IsProcessed *bool `json:"is_processed" form:"is_processed"`
}

// CreateSmartAlertRequest 创建智能警报请求
type CreateSmartAlertRequest struct {
	BabyID      uint   `json:"baby_id"`
	DeviceID    uint   `json:"device_id"`
	AlertType   int    `json:"alert_type" binding:"required"`
	AlertLevel  int    `json:"alert_level" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Message     string `json:"message"`
	TriggerData string `json:"trigger_data"`
	MediaURL    string `json:"media_url"`
}

// ProcessAlertRequest 处理警报请求
type ProcessAlertRequest struct {
	ID         uint   `json:"id" binding:"required"`
	UserAction string `json:"user_action" binding:"required"`
}

// EnvironmentDataSearch 环境数据搜索条件
type EnvironmentDataSearch struct {
	request.PageInfo
	UserID    uint   `json:"user_id" form:"user_id"`
	DeviceID  uint   `json:"device_id" form:"device_id"`
	StartDate string `json:"start_date" form:"start_date"`
	EndDate   string `json:"end_date" form:"end_date"`
}

// CreateEnvironmentDataRequest 创建环境数据请求
type CreateEnvironmentDataRequest struct {
	DeviceID    uint    `json:"device_id" binding:"required"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	AirQuality  int     `json:"air_quality"`
	NoiseLevel  float64 `json:"noise_level"`
	Brightness  int     `json:"brightness"`
	CO2Level    int     `json:"co2_level"`
}

// AnalysisReportRequest 分析报告请求
type AnalysisReportRequest struct {
	BabyID     uint      `json:"baby_id" binding:"required"`
	ReportType int       `json:"report_type" binding:"required,oneof=1 2 3"`
	StartDate  time.Time `json:"start_date" binding:"required"`
	EndDate    time.Time `json:"end_date" binding:"required"`
}
