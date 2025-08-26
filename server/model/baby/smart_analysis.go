package baby

import (
	"time"
	"baby_admin/server/global"
)

// SleepRecord 睡眠记录表
type SleepRecord struct {
	global.GVA_MODEL
	UserID      uint      `json:"user_id" gorm:"not null;comment:用户ID"`
	BabyID      uint      `json:"baby_id" gorm:"not null;comment:宝宝ID"`
	DeviceID    uint      `json:"device_id" gorm:"comment:设备ID"`
	StartTime   time.Time `json:"start_time" gorm:"not null;comment:睡眠开始时间"`
	EndTime     *time.Time `json:"end_time" gorm:"comment:睡眠结束时间"`
	Duration    int       `json:"duration" gorm:"comment:睡眠时长(分钟)"`
	Quality     int       `json:"quality" gorm:"comment:睡眠质量评分(1-10)"`
	DeepSleep   int       `json:"deep_sleep" gorm:"comment:深度睡眠时长(分钟)"`
	LightSleep  int       `json:"light_sleep" gorm:"comment:浅度睡眠时长(分钟)"`
	AwakeCount  int       `json:"awake_count" gorm:"default:0;comment:醒来次数"`
	Temperature float64   `json:"temperature" gorm:"type:decimal(4,2);comment:环境温度"`
	Humidity    float64   `json:"humidity" gorm:"type:decimal(5,2);comment:环境湿度"`
	NoiseLevel  float64   `json:"noise_level" gorm:"type:decimal(5,2);comment:噪音水平(dB)"`
	Status      int       `json:"status" gorm:"default:1;comment:状态:1进行中,2已完成,3异常"`
	Notes       string    `json:"notes" gorm:"type:text;comment:备注"`
}

// TableName 指定表名
func (SleepRecord) TableName() string {
	return "sleep_records"
}

// CryDetection 哭声检测记录表
type CryDetection struct {
	global.GVA_MODEL
	UserID       uint      `json:"user_id" gorm:"not null;comment:用户ID"`
	BabyID       uint      `json:"baby_id" gorm:"not null;comment:宝宝ID"`
	DeviceID     uint      `json:"device_id" gorm:"comment:设备ID"`
	DetectedAt   time.Time `json:"detected_at" gorm:"not null;comment:检测时间"`
	Duration     int       `json:"duration" gorm:"comment:哭声持续时长(秒)"`
	Intensity    int       `json:"intensity" gorm:"comment:哭声强度(1-10)"`
	CryType      int       `json:"cry_type" gorm:"comment:哭声类型:1饥饿,2困倦,3不适,4疼痛,5其他"`
	Confidence   float64   `json:"confidence" gorm:"type:decimal(5,2);comment:识别置信度"`
	AudioURL     string    `json:"audio_url" gorm:"size:500;comment:音频文件URL"`
	VideoURL     string    `json:"video_url" gorm:"size:500;comment:视频文件URL"`
	IsProcessed  bool      `json:"is_processed" gorm:"default:false;comment:是否已处理"`
	UserFeedback int       `json:"user_feedback" gorm:"comment:用户反馈:1准确,2不准确"`
}

// TableName 指定表名
func (CryDetection) TableName() string {
	return "cry_detections"
}

// MovementDetection 动作检测记录表
type MovementDetection struct {
	global.GVA_MODEL
	UserID      uint      `json:"user_id" gorm:"not null;comment:用户ID"`
	BabyID      uint      `json:"baby_id" gorm:"not null;comment:宝宝ID"`
	DeviceID    uint      `json:"device_id" gorm:"comment:设备ID"`
	DetectedAt  time.Time `json:"detected_at" gorm:"not null;comment:检测时间"`
	MovementType int      `json:"movement_type" gorm:"comment:动作类型:1翻身,2踢腿,3手部动作,4头部动作,5异常动作"`
	Intensity   int       `json:"intensity" gorm:"comment:动作强度(1-10)"`
	Duration    int       `json:"duration" gorm:"comment:动作持续时长(秒)"`
	VideoURL    string    `json:"video_url" gorm:"size:500;comment:视频文件URL"`
	ThumbnailURL string   `json:"thumbnail_url" gorm:"size:500;comment:缩略图URL"`
	IsAlert     bool      `json:"is_alert" gorm:"default:false;comment:是否需要警报"`
}

// TableName 指定表名
func (MovementDetection) TableName() string {
	return "movement_detections"
}

// EnvironmentData 环境数据记录表
type EnvironmentData struct {
	global.GVA_MODEL
	UserID      uint      `json:"user_id" gorm:"not null;comment:用户ID"`
	DeviceID    uint      `json:"device_id" gorm:"not null;comment:设备ID"`
	RecordedAt  time.Time `json:"recorded_at" gorm:"not null;comment:记录时间"`
	Temperature float64   `json:"temperature" gorm:"type:decimal(4,2);comment:温度(摄氏度)"`
	Humidity    float64   `json:"humidity" gorm:"type:decimal(5,2);comment:湿度(%)"`
	AirQuality  int       `json:"air_quality" gorm:"comment:空气质量指数"`
	NoiseLevel  float64   `json:"noise_level" gorm:"type:decimal(5,2);comment:噪音水平(dB)"`
	Brightness  int       `json:"brightness" gorm:"comment:亮度(lux)"`
	CO2Level    int       `json:"co2_level" gorm:"comment:二氧化碳浓度(ppm)"`
}

// TableName 指定表名
func (EnvironmentData) TableName() string {
	return "environment_data"
}

// SmartAlert 智能警报记录表
type SmartAlert struct {
	global.GVA_MODEL
	UserID       uint      `json:"user_id" gorm:"not null;comment:用户ID"`
	BabyID       uint      `json:"baby_id" gorm:"comment:宝宝ID"`
	DeviceID     uint      `json:"device_id" gorm:"comment:设备ID"`
	AlertType    int       `json:"alert_type" gorm:"not null;comment:警报类型:1哭声,2异常动作,3环境异常,4离开检测,5其他"`
	AlertLevel   int       `json:"alert_level" gorm:"not null;comment:警报等级:1低,2中,3高,4紧急"`
	Title        string    `json:"title" gorm:"size:100;not null;comment:警报标题"`
	Message      string    `json:"message" gorm:"type:text;comment:警报消息"`
	TriggerData  string    `json:"trigger_data" gorm:"type:text;comment:触发数据(JSON格式)"`
	MediaURL     string    `json:"media_url" gorm:"size:500;comment:相关媒体文件URL"`
	IsRead       bool      `json:"is_read" gorm:"default:false;comment:是否已读"`
	IsProcessed  bool      `json:"is_processed" gorm:"default:false;comment:是否已处理"`
	ProcessedAt  *time.Time `json:"processed_at" gorm:"comment:处理时间"`
	UserAction   string    `json:"user_action" gorm:"size:50;comment:用户处理动作"`
}

// TableName 指定表名
func (SmartAlert) TableName() string {
	return "smart_alerts"
}

// AnalysisReport 分析报告表
type AnalysisReport struct {
	global.GVA_MODEL
	UserID      uint      `json:"user_id" gorm:"not null;comment:用户ID"`
	BabyID      uint      `json:"baby_id" gorm:"not null;comment:宝宝ID"`
	ReportType  int       `json:"report_type" gorm:"not null;comment:报告类型:1日报,2周报,3月报"`
	ReportDate  time.Time `json:"report_date" gorm:"not null;comment:报告日期"`
	Title       string    `json:"title" gorm:"size:100;not null;comment:报告标题"`
	Summary     string    `json:"summary" gorm:"type:text;comment:摘要"`
	Content     string    `json:"content" gorm:"type:longtext;comment:报告内容(JSON格式)"`
	SleepData   string    `json:"sleep_data" gorm:"type:text;comment:睡眠数据统计"`
	CryData     string    `json:"cry_data" gorm:"type:text;comment:哭声数据统计"`
	EnvironData string    `json:"environ_data" gorm:"type:text;comment:环境数据统计"`
	Suggestions string    `json:"suggestions" gorm:"type:text;comment:建议和改进"`
	IsGenerated bool      `json:"is_generated" gorm:"default:false;comment:是否已生成"`
}

// TableName 指定表名
func (AnalysisReport) TableName() string {
	return "analysis_reports"
}

// GetCryTypeText 获取哭声类型文本
func (c *CryDetection) GetCryTypeText() string {
	switch c.CryType {
	case 1:
		return "饥饿"
	case 2:
		return "困倦"
	case 3:
		return "不适"
	case 4:
		return "疼痛"
	case 5:
		return "其他"
	default:
		return "未知"
	}
}

// GetMovementTypeText 获取动作类型文本
func (m *MovementDetection) GetMovementTypeText() string {
	switch m.MovementType {
	case 1:
		return "翻身"
	case 2:
		return "踢腿"
	case 3:
		return "手部动作"
	case 4:
		return "头部动作"
	case 5:
		return "异常动作"
	default:
		return "未知动作"
	}
}

// GetAlertTypeText 获取警报类型文本
func (s *SmartAlert) GetAlertTypeText() string {
	switch s.AlertType {
	case 1:
		return "哭声警报"
	case 2:
		return "异常动作"
	case 3:
		return "环境异常"
	case 4:
		return "离开检测"
	case 5:
		return "其他警报"
	default:
		return "未知警报"
	}
}

// GetAlertLevelText 获取警报等级文本
func (s *SmartAlert) GetAlertLevelText() string {
	switch s.AlertLevel {
	case 1:
		return "低"
	case 2:
		return "中"
	case 3:
		return "高"
	case 4:
		return "紧急"
	default:
		return "未知"
	}
}