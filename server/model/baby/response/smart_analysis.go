package response

import (
	"baby_admin/server/model/baby"
	"fmt"
	"time"
)

// SleepRecordResponse 睡眠记录响应
type SleepRecordResponse struct {
	ID           uint       `json:"id"`
	UserID       uint       `json:"user_id"`
	BabyID       uint       `json:"baby_id"`
	BabyName     string     `json:"baby_name"`
	DeviceID     uint       `json:"device_id"`
	StartTime    time.Time  `json:"start_time"`
	EndTime      *time.Time `json:"end_time"`
	Duration     int        `json:"duration"`
	DurationText string     `json:"duration_text"` // 格式化的时长显示
	Quality      int        `json:"quality"`
	QualityText  string     `json:"quality_text"` // 质量描述
	DeepSleep    int        `json:"deep_sleep"`
	LightSleep   int        `json:"light_sleep"`
	AwakeCount   int        `json:"awake_count"`
	Temperature  float64    `json:"temperature"`
	Humidity     float64    `json:"humidity"`
	NoiseLevel   float64    `json:"noise_level"`
	Status       int        `json:"status"`
	StatusText   string     `json:"status_text"`
	Notes        string     `json:"notes"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

// FromSleepRecord 从SleepRecord模型转换
func (s *SleepRecordResponse) FromSleepRecord(record *baby.SleepRecord, babyName string) {
	s.ID = record.ID
	s.UserID = record.UserID
	s.BabyID = record.BabyID
	s.BabyName = babyName
	s.DeviceID = record.DeviceID
	s.StartTime = record.StartTime
	s.EndTime = record.EndTime
	s.Duration = record.Duration
	s.DurationText = formatMinutesToDuration(record.Duration)
	s.Quality = record.Quality
	s.QualityText = getSleepQualityText(record.Quality)
	s.DeepSleep = record.DeepSleep
	s.LightSleep = record.LightSleep
	s.AwakeCount = record.AwakeCount
	s.Temperature = record.Temperature
	s.Humidity = record.Humidity
	s.NoiseLevel = record.NoiseLevel
	s.Status = record.Status
	s.StatusText = getSleepStatusText(record.Status)
	s.Notes = record.Notes
	s.CreatedAt = record.CreatedAt
	s.UpdatedAt = record.UpdatedAt
}

// CryDetectionResponse 哭声检测响应
type CryDetectionResponse struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	BabyID       uint      `json:"baby_id"`
	BabyName     string    `json:"baby_name"`
	DeviceID     uint      `json:"device_id"`
	DetectedAt   time.Time `json:"detected_at"`
	Duration     int       `json:"duration"`
	Intensity    int       `json:"intensity"`
	CryType      int       `json:"cry_type"`
	CryTypeText  string    `json:"cry_type_text"`
	Confidence   float64   `json:"confidence"`
	AudioURL     string    `json:"audio_url"`
	VideoURL     string    `json:"video_url"`
	IsProcessed  bool      `json:"is_processed"`
	UserFeedback int       `json:"user_feedback"`
	CreatedAt    time.Time `json:"created_at"`
}

// FromCryDetection 从CryDetection模型转换
func (c *CryDetectionResponse) FromCryDetection(detection *baby.CryDetection, babyName string) {
	c.ID = detection.ID
	c.UserID = detection.UserID
	c.BabyID = detection.BabyID
	c.BabyName = babyName
	c.DeviceID = detection.DeviceID
	c.DetectedAt = detection.DetectedAt
	c.Duration = detection.Duration
	c.Intensity = detection.Intensity
	c.CryType = detection.CryType
	c.CryTypeText = detection.GetCryTypeText()
	c.Confidence = detection.Confidence
	c.AudioURL = detection.AudioURL
	c.VideoURL = detection.VideoURL
	c.IsProcessed = detection.IsProcessed
	c.UserFeedback = detection.UserFeedback
	c.CreatedAt = detection.CreatedAt
}

// MovementDetectionResponse 动作检测响应
type MovementDetectionResponse struct {
	ID               uint      `json:"id"`
	UserID           uint      `json:"user_id"`
	BabyID           uint      `json:"baby_id"`
	BabyName         string    `json:"baby_name"`
	DeviceID         uint      `json:"device_id"`
	DetectedAt       time.Time `json:"detected_at"`
	MovementType     int       `json:"movement_type"`
	MovementTypeText string    `json:"movement_type_text"`
	Intensity        int       `json:"intensity"`
	Duration         int       `json:"duration"`
	VideoURL         string    `json:"video_url"`
	ThumbnailURL     string    `json:"thumbnail_url"`
	IsAlert          bool      `json:"is_alert"`
	CreatedAt        time.Time `json:"created_at"`
}

// FromMovementDetection 从MovementDetection模型转换
func (m *MovementDetectionResponse) FromMovementDetection(detection *baby.MovementDetection, babyName string) {
	m.ID = detection.ID
	m.UserID = detection.UserID
	m.BabyID = detection.BabyID
	m.BabyName = babyName
	m.DeviceID = detection.DeviceID
	m.DetectedAt = detection.DetectedAt
	m.MovementType = detection.MovementType
	m.MovementTypeText = detection.GetMovementTypeText()
	m.Intensity = detection.Intensity
	m.Duration = detection.Duration
	m.VideoURL = detection.VideoURL
	m.ThumbnailURL = detection.ThumbnailURL
	m.IsAlert = detection.IsAlert
	m.CreatedAt = detection.CreatedAt
}

// SmartAlertResponse 智能警报响应
type SmartAlertResponse struct {
	ID             uint       `json:"id"`
	UserID         uint       `json:"user_id"`
	BabyID         uint       `json:"baby_id"`
	BabyName       string     `json:"baby_name"`
	DeviceID       uint       `json:"device_id"`
	AlertType      int        `json:"alert_type"`
	AlertTypeText  string     `json:"alert_type_text"`
	AlertLevel     int        `json:"alert_level"`
	AlertLevelText string     `json:"alert_level_text"`
	Title          string     `json:"title"`
	Message        string     `json:"message"`
	TriggerData    string     `json:"trigger_data"`
	MediaURL       string     `json:"media_url"`
	IsRead         bool       `json:"is_read"`
	IsProcessed    bool       `json:"is_processed"`
	ProcessedAt    *time.Time `json:"processed_at"`
	UserAction     string     `json:"user_action"`
	CreatedAt      time.Time  `json:"created_at"`
}

// FromSmartAlert 从SmartAlert模型转换
func (s *SmartAlertResponse) FromSmartAlert(alert *baby.SmartAlert, babyName string) {
	s.ID = alert.ID
	s.UserID = alert.UserID
	s.BabyID = alert.BabyID
	s.BabyName = babyName
	s.DeviceID = alert.DeviceID
	s.AlertType = alert.AlertType
	s.AlertTypeText = alert.GetAlertTypeText()
	s.AlertLevel = alert.AlertLevel
	s.AlertLevelText = alert.GetAlertLevelText()
	s.Title = alert.Title
	s.Message = alert.Message
	s.TriggerData = alert.TriggerData
	s.MediaURL = alert.MediaURL
	s.IsRead = alert.IsRead
	s.IsProcessed = alert.IsProcessed
	s.ProcessedAt = alert.ProcessedAt
	s.UserAction = alert.UserAction
	s.CreatedAt = alert.CreatedAt
}

// EnvironmentDataResponse 环境数据响应
type EnvironmentDataResponse struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	DeviceID    uint      `json:"device_id"`
	RecordedAt  time.Time `json:"recorded_at"`
	Temperature float64   `json:"temperature"`
	Humidity    float64   `json:"humidity"`
	AirQuality  int       `json:"air_quality"`
	NoiseLevel  float64   `json:"noise_level"`
	Brightness  int       `json:"brightness"`
	CO2Level    int       `json:"co2_level"`
	CreatedAt   time.Time `json:"created_at"`
}

// SleepStatistics 睡眠统计
type SleepStatistics struct {
	TotalSleepTime   int     `json:"total_sleep_time"`   // 总睡眠时间(分钟)
	AverageSleepTime int     `json:"average_sleep_time"` // 平均睡眠时间(分钟)
	SleepQuality     float64 `json:"sleep_quality"`      // 平均睡眠质量
	DeepSleepRatio   float64 `json:"deep_sleep_ratio"`   // 深度睡眠比例
	AwakeFrequency   float64 `json:"awake_frequency"`    // 平均醒来频率
	SleepTrend       string  `json:"sleep_trend"`        // 睡眠趋势
}

// CryStatistics 哭声统计
type CryStatistics struct {
	TotalCryCount    int            `json:"total_cry_count"`    // 总哭声次数
	CryFrequency     float64        `json:"cry_frequency"`      // 哭声频率(次/天)
	AverageIntensity float64        `json:"average_intensity"`  // 平均强度
	CryTypeBreakdown map[string]int `json:"cry_type_breakdown"` // 哭声类型分布
	CryTrend         string         `json:"cry_trend"`          // 哭声趋势
}

// EnvironmentStatistics 环境统计
type EnvironmentStatistics struct {
	AverageTemperature float64 `json:"average_temperature"` // 平均温度
	AverageHumidity    float64 `json:"average_humidity"`    // 平均湿度
	AverageNoiseLevel  float64 `json:"average_noise_level"` // 平均噪音水平
	ComfortIndex       float64 `json:"comfort_index"`       // 舒适度指数
	OptimalHours       int     `json:"optimal_hours"`       // 最佳环境小时数
}

// AnalysisDashboard 分析仪表盘
type AnalysisDashboard struct {
	SleepStats       SleepStatistics       `json:"sleep_stats"`
	CryStats         CryStatistics         `json:"cry_stats"`
	EnvironmentStats EnvironmentStatistics `json:"environment_stats"`
	RecentAlerts     []SmartAlertResponse  `json:"recent_alerts"`
	TodaySummary     TodaySummary          `json:"today_summary"`
}

// TodaySummary 今日摘要
type TodaySummary struct {
	SleepTime    int    `json:"sleep_time"`    // 今日睡眠时间(分钟)
	CryCount     int    `json:"cry_count"`     // 今日哭声次数
	SleepQuality int    `json:"sleep_quality"` // 今日睡眠质量
	ComfortLevel string `json:"comfort_level"` // 舒适度等级
}

// 工具函数
func getSleepQualityText(quality int) string {
	switch {
	case quality >= 8:
		return "优秀"
	case quality >= 6:
		return "良好"
	case quality >= 4:
		return "一般"
	case quality >= 2:
		return "较差"
	default:
		return "很差"
	}
}

func getSleepStatusText(status int) string {
	switch status {
	case 1:
		return "进行中"
	case 2:
		return "已完成"
	case 3:
		return "异常"
	default:
		return "未知"
	}
}

func formatMinutesToDuration(minutes int) string {
	if minutes <= 0 {
		return "0分钟"
	}
	hours := minutes / 60
	remainingMinutes := minutes % 60
	if hours > 0 {
		return fmt.Sprintf("%d小时%d分钟", hours, remainingMinutes)
	}
	return fmt.Sprintf("%d分钟", remainingMinutes)
}
