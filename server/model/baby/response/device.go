package response

import (
	"baby_admin/server/model/baby"
	"time"
)

// DeviceResponse 设备响应
type DeviceResponse struct {
	ID              uint       `json:"id"`
	UserID          uint       `json:"user_id"`
	DeviceName      string     `json:"device_name"`
	DeviceType      int        `json:"device_type"`
	DeviceTypeText  string     `json:"device_type_text"`
	ProductID       string     `json:"product_id"`
	SerialNumber    string     `json:"serial_number"`
	MacAddress      string     `json:"mac_address"`
	FirmwareVersion string     `json:"firmware_version"`
	HardwareVersion string     `json:"hardware_version"`
	Location        string     `json:"location"`
	Description     string     `json:"description"`
	Status          int        `json:"status"`
	StatusText      string     `json:"status_text"`
	IsActive        bool       `json:"is_active"`
	IsOnline        bool       `json:"is_online"`
	LastOnlineAt    *time.Time `json:"last_online_at"`
	ActivatedAt     time.Time  `json:"activated_at"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	// 关联信息
	ShareCount     int    `json:"share_count,omitempty"`     // 分享数量
	CommandCount   int    `json:"command_count,omitempty"`   // 命令数量
	OnlineDuration string `json:"online_duration,omitempty"` // 在线时长
}

// FromDevice 从Device模型转换
func (d *DeviceResponse) FromDevice(device *baby.Device) {
	d.ID = device.ID
	d.UserID = device.UserID
	d.DeviceName = device.DeviceName
	d.DeviceType = device.DeviceType
	d.DeviceTypeText = device.GetDeviceTypeText()
	d.ProductID = device.ProductID
	d.SerialNumber = device.SerialNumber
	d.MacAddress = device.MacAddress
	d.FirmwareVersion = device.FirmwareVersion
	d.HardwareVersion = device.HardwareVersion
	d.Location = device.Location
	d.Description = device.Description
	d.Status = device.Status
	d.StatusText = device.GetStatusText()
	d.IsActive = device.IsActive
	d.IsOnline = device.IsOnline()
	d.LastOnlineAt = device.LastOnlineAt
	d.ActivatedAt = device.ActivatedAt
	d.CreatedAt = device.CreatedAt
	d.UpdatedAt = device.UpdatedAt
}

// DeviceConfigResponse 设备配置响应
type DeviceConfigResponse struct {
	ID           uint      `json:"id"`
	DeviceID     uint      `json:"device_id"`
	ConfigKey    string    `json:"config_key"`
	ConfigValue  string    `json:"config_value"`
	ValueType    string    `json:"value_type"`
	Description  string    `json:"description"`
	IsEditable   bool      `json:"is_editable"`
	DefaultValue string    `json:"default_value"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// DeviceCommandResponse 设备命令响应
type DeviceCommandResponse struct {
	ID              uint       `json:"id"`
	UserID          uint       `json:"user_id"`
	DeviceID        uint       `json:"device_id"`
	DeviceName      string     `json:"device_name"`
	CommandType     int        `json:"command_type"`
	CommandTypeText string     `json:"command_type_text"`
	Command         string     `json:"command"`
	Parameters      string     `json:"parameters"`
	Status          int        `json:"status"`
	StatusText      string     `json:"status_text"`
	Response        string     `json:"response"`
	ExecutedAt      *time.Time `json:"executed_at"`
	CompletedAt     *time.Time `json:"completed_at"`
	ErrorMsg        string     `json:"error_msg"`
	Duration        int        `json:"duration,omitempty"` // 执行耗时(秒)
	CreatedAt       time.Time  `json:"created_at"`
}

// FromDeviceCommand 从DeviceCommand模型转换
func (d *DeviceCommandResponse) FromDeviceCommand(command *baby.DeviceCommand, deviceName string) {
	d.ID = command.ID
	d.UserID = command.UserID
	d.DeviceID = command.DeviceID
	d.DeviceName = deviceName
	d.CommandType = command.CommandType
	d.CommandTypeText = command.GetCommandTypeText()
	d.Command = command.Command
	d.Parameters = command.Parameters
	d.Status = command.Status
	d.StatusText = command.GetStatusText()
	d.Response = command.Response
	d.ExecutedAt = command.ExecutedAt
	d.CompletedAt = command.CompletedAt
	d.ErrorMsg = command.ErrorMsg
	d.CreatedAt = command.CreatedAt

	// 计算执行耗时
	if command.ExecutedAt != nil && command.CompletedAt != nil {
		d.Duration = int(command.CompletedAt.Sub(*command.ExecutedAt).Seconds())
	}
}

// DeviceShareResponse 设备分享响应
type DeviceShareResponse struct {
	ID             uint       `json:"id"`
	DeviceID       uint       `json:"device_id"`
	DeviceName     string     `json:"device_name"`
	OwnerID        uint       `json:"owner_id"`
	OwnerName      string     `json:"owner_name"`
	SharedUserID   uint       `json:"shared_user_id"`
	SharedUserName string     `json:"shared_user_name"`
	ShareType      int        `json:"share_type"`
	ShareTypeText  string     `json:"share_type_text"`
	Permissions    string     `json:"permissions"`
	ExpiredAt      *time.Time `json:"expired_at"`
	IsActive       bool       `json:"is_active"`
	IsExpired      bool       `json:"is_expired"`
	AcceptedAt     *time.Time `json:"accepted_at"`
	CreatedAt      time.Time  `json:"created_at"`
}

// FromDeviceShare 从DeviceShare模型转换
func (d *DeviceShareResponse) FromDeviceShare(share *baby.DeviceShare, deviceName, ownerName, sharedUserName string) {
	d.ID = share.ID
	d.DeviceID = share.DeviceID
	d.DeviceName = deviceName
	d.OwnerID = share.OwnerID
	d.OwnerName = ownerName
	d.SharedUserID = share.SharedUserID
	d.SharedUserName = sharedUserName
	d.ShareType = share.ShareType
	d.ShareTypeText = share.GetShareTypeText()
	d.Permissions = share.Permissions
	d.ExpiredAt = share.ExpiredAt
	d.IsActive = share.IsActive
	d.IsExpired = share.IsExpired()
	d.AcceptedAt = share.AcceptedAt
	d.CreatedAt = share.CreatedAt
}

// DeviceStatusResponse 设备状态响应
type DeviceStatusResponse struct {
	ID          uint      `json:"id"`
	DeviceID    uint      `json:"device_id"`
	StatusType  int       `json:"status_type"`
	StatusValue string    `json:"status_value"`
	StatusData  string    `json:"status_data"`
	RecordedAt  time.Time `json:"recorded_at"`
	IsAlert     bool      `json:"is_alert"`
	CreatedAt   time.Time `json:"created_at"`
}

// DeviceLogResponse 设备日志响应
type DeviceLogResponse struct {
	ID           uint      `json:"id"`
	DeviceID     uint      `json:"device_id"`
	LogLevel     int       `json:"log_level"`
	LogLevelText string    `json:"log_level_text"`
	LogType      int       `json:"log_type"`
	LogTypeText  string    `json:"log_type_text"`
	Message      string    `json:"message"`
	Details      string    `json:"details"`
	Source       string    `json:"source"`
	RecordedAt   time.Time `json:"recorded_at"`
	CreatedAt    time.Time `json:"created_at"`
}

// FromDeviceLog 从DeviceLog模型转换
func (d *DeviceLogResponse) FromDeviceLog(log *baby.DeviceLog) {
	d.ID = log.ID
	d.DeviceID = log.DeviceID
	d.LogLevel = log.LogLevel
	d.LogLevelText = log.GetLogLevelText()
	d.LogType = log.LogType
	d.LogTypeText = log.GetLogTypeText()
	d.Message = log.Message
	d.Details = log.Details
	d.Source = log.Source
	d.RecordedAt = log.RecordedAt
	d.CreatedAt = log.CreatedAt
}

// DeviceListResponse 设备列表响应
type DeviceListResponse struct {
	List     []DeviceResponse `json:"list"`
	Total    int64            `json:"total"`
	Page     int              `json:"page"`
	PageSize int              `json:"page_size"`
}

// DeviceStatistics 设备统计
type DeviceStatistics struct {
	TotalDevices   int64                   `json:"total_devices"`   // 总设备数
	OnlineDevices  int64                   `json:"online_devices"`  // 在线设备数
	OfflineDevices int64                   `json:"offline_devices"` // 离线设备数
	FaultDevices   int64                   `json:"fault_devices"`   // 故障设备数
	DeviceTypes    map[string]int64        `json:"device_types"`    // 设备类型分布
	StatusTrend    []DeviceStatusTrendItem `json:"status_trend"`    // 状态趋势
	RecentCommands []DeviceCommandResponse `json:"recent_commands"` // 最近命令
	RecentLogs     []DeviceLogResponse     `json:"recent_logs"`     // 最近日志
}

// DeviceStatusTrendItem 设备状态趋势项
type DeviceStatusTrendItem struct {
	Date    string `json:"date"`
	Online  int64  `json:"online"`
	Offline int64  `json:"offline"`
	Fault   int64  `json:"fault"`
}

// DeviceDashboard 设备仪表盘
type DeviceDashboard struct {
	Statistics    DeviceStatistics     `json:"statistics"`
	QuickActions  []QuickAction        `json:"quick_actions"`
	Notifications []DeviceNotification `json:"notifications"`
	HealthStatus  DeviceHealthStatus   `json:"health_status"`
}

// QuickAction 快捷操作
type QuickAction struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Action      string `json:"action"`
	Params      string `json:"params"`
}

// DeviceNotification 设备通知
type DeviceNotification struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	Level     string    `json:"level"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}

// DeviceHealthStatus 设备健康状态
type DeviceHealthStatus struct {
	OverallScore int                `json:"overall_score"` // 总体评分
	HealthItems  []DeviceHealthItem `json:"health_items"`  // 健康项目
	Suggestions  []string           `json:"suggestions"`   // 改进建议
}

// DeviceHealthItem 设备健康项目
type DeviceHealthItem struct {
	Name        string `json:"name"`
	Score       int    `json:"score"`
	Status      string `json:"status"`
	Description string `json:"description"`
}
