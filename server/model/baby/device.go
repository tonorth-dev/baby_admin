package baby

import (
	"time"
	"baby_admin/server/global"
)

// Device 设备表
type Device struct {
	global.GVA_MODEL
	UserID       uint      `json:"user_id" gorm:"not null;comment:用户ID"`
	DeviceName   string    `json:"device_name" gorm:"size:100;not null;comment:设备名称"`
	DeviceType   int       `json:"device_type" gorm:"not null;comment:设备类型:1摄像头,2传感器,3音响,4其他"`
	ProductID    string    `json:"product_id" gorm:"size:50;not null;comment:腾讯云产品ID"`
	DeviceSecret string    `json:"device_secret" gorm:"size:100;not null;comment:设备密钥"`
	SerialNumber string    `json:"serial_number" gorm:"size:50;comment:设备序列号"`
	MacAddress   string    `json:"mac_address" gorm:"size:20;comment:MAC地址"`
	FirmwareVersion string `json:"firmware_version" gorm:"size:20;comment:固件版本"`
	HardwareVersion string `json:"hardware_version" gorm:"size:20;comment:硬件版本"`
	Location     string    `json:"location" gorm:"size:100;comment:设备位置"`
	Description  string    `json:"description" gorm:"size:255;comment:设备描述"`
	Status       int       `json:"status" gorm:"default:1;comment:设备状态:1在线,2离线,3故障,4维护"`
	IsActive     bool      `json:"is_active" gorm:"default:true;comment:是否启用"`
	LastOnlineAt *time.Time `json:"last_online_at" gorm:"comment:最后在线时间"`
	ActivatedAt  time.Time  `json:"activated_at" gorm:"comment:激活时间"`
}

// TableName 指定表名
func (Device) TableName() string {
	return "devices"
}

// DeviceConfig 设备配置表
type DeviceConfig struct {
	global.GVA_MODEL
	DeviceID     uint   `json:"device_id" gorm:"not null;comment:设备ID"`
	ConfigKey    string `json:"config_key" gorm:"size:50;not null;comment:配置键"`
	ConfigValue  string `json:"config_value" gorm:"type:text;comment:配置值"`
	ValueType    string `json:"value_type" gorm:"size:20;default:'string';comment:值类型:string,int,float,bool,json"`
	Description  string `json:"description" gorm:"size:255;comment:配置描述"`
	IsEditable   bool   `json:"is_editable" gorm:"default:true;comment:是否可编辑"`
	DefaultValue string `json:"default_value" gorm:"type:text;comment:默认值"`
}

// TableName 指定表名
func (DeviceConfig) TableName() string {
	return "device_configs"
}

// DeviceCommand 设备命令记录表
type DeviceCommand struct {
	global.GVA_MODEL
	UserID      uint      `json:"user_id" gorm:"not null;comment:用户ID"`
	DeviceID    uint      `json:"device_id" gorm:"not null;comment:设备ID"`
	CommandType int       `json:"command_type" gorm:"not null;comment:命令类型:1云台控制,2夜视切换,3音量调节,4重启,5其他"`
	Command     string    `json:"command" gorm:"size:100;not null;comment:命令名称"`
	Parameters  string    `json:"parameters" gorm:"type:text;comment:命令参数(JSON格式)"`
	Status      int       `json:"status" gorm:"default:1;comment:执行状态:1待执行,2执行中,3成功,4失败"`
	Response    string    `json:"response" gorm:"type:text;comment:设备响应"`
	ExecutedAt  *time.Time `json:"executed_at" gorm:"comment:执行时间"`
	CompletedAt *time.Time `json:"completed_at" gorm:"comment:完成时间"`
	ErrorMsg    string    `json:"error_msg" gorm:"size:255;comment:错误信息"`
}

// TableName 指定表名
func (DeviceCommand) TableName() string {
	return "device_commands"
}

// DeviceStatus 设备状态记录表
type DeviceStatus struct {
	global.GVA_MODEL
	DeviceID      uint      `json:"device_id" gorm:"not null;comment:设备ID"`
	StatusType    int       `json:"status_type" gorm:"not null;comment:状态类型:1在线状态,2电量状态,3信号强度,4其他"`
	StatusValue   string    `json:"status_value" gorm:"size:100;comment:状态值"`
	StatusData    string    `json:"status_data" gorm:"type:text;comment:状态数据(JSON格式)"`
	RecordedAt    time.Time `json:"recorded_at" gorm:"not null;comment:记录时间"`
	IsAlert       bool      `json:"is_alert" gorm:"default:false;comment:是否需要警报"`
}

// TableName 指定表名
func (DeviceStatus) TableName() string {
	return "device_statuses"
}

// DeviceShare 设备分享表
type DeviceShare struct {
	global.GVA_MODEL
	DeviceID     uint       `json:"device_id" gorm:"not null;comment:设备ID"`
	OwnerID      uint       `json:"owner_id" gorm:"not null;comment:设备所有者ID"`
	SharedUserID uint       `json:"shared_user_id" gorm:"not null;comment:被分享用户ID"`
	ShareType    int        `json:"share_type" gorm:"not null;comment:分享类型:1只读,2控制,3管理"`
	Permissions  string     `json:"permissions" gorm:"type:text;comment:权限配置(JSON格式)"`
	ExpiredAt    *time.Time `json:"expired_at" gorm:"comment:过期时间"`
	IsActive     bool       `json:"is_active" gorm:"default:true;comment:是否有效"`
	AcceptedAt   *time.Time `json:"accepted_at" gorm:"comment:接受时间"`
}

// TableName 指定表名
func (DeviceShare) TableName() string {
	return "device_shares"
}

// DeviceLog 设备日志表
type DeviceLog struct {
	global.GVA_MODEL
	DeviceID    uint      `json:"device_id" gorm:"not null;comment:设备ID"`
	LogLevel    int       `json:"log_level" gorm:"not null;comment:日志级别:1调试,2信息,3警告,4错误"`
	LogType     int       `json:"log_type" gorm:"not null;comment:日志类型:1系统,2业务,3安全,4其他"`
	Message     string    `json:"message" gorm:"type:text;not null;comment:日志消息"`
	Details     string    `json:"details" gorm:"type:text;comment:详细信息(JSON格式)"`
	Source      string    `json:"source" gorm:"size:50;comment:日志来源"`
	RecordedAt  time.Time `json:"recorded_at" gorm:"not null;comment:记录时间"`
}

// TableName 指定表名
func (DeviceLog) TableName() string {
	return "device_logs"
}

// 方法定义

// GetDeviceTypeText 获取设备类型文本
func (d *Device) GetDeviceTypeText() string {
	switch d.DeviceType {
	case 1:
		return "摄像头"
	case 2:
		return "传感器"
	case 3:
		return "音响"
	case 4:
		return "其他"
	default:
		return "未知设备"
	}
}

// GetStatusText 获取设备状态文本
func (d *Device) GetStatusText() string {
	switch d.Status {
	case 1:
		return "在线"
	case 2:
		return "离线"
	case 3:
		return "故障"
	case 4:
		return "维护"
	default:
		return "未知状态"
	}
}

// IsOnline 判断设备是否在线
func (d *Device) IsOnline() bool {
	return d.Status == 1
}

// GetCommandTypeText 获取命令类型文本
func (dc *DeviceCommand) GetCommandTypeText() string {
	switch dc.CommandType {
	case 1:
		return "云台控制"
	case 2:
		return "夜视切换"
	case 3:
		return "音量调节"
	case 4:
		return "设备重启"
	case 5:
		return "其他命令"
	default:
		return "未知命令"
	}
}

// GetStatusText 获取命令状态文本
func (dc *DeviceCommand) GetStatusText() string {
	switch dc.Status {
	case 1:
		return "待执行"
	case 2:
		return "执行中"
	case 3:
		return "成功"
	case 4:
		return "失败"
	default:
		return "未知状态"
	}
}

// GetShareTypeText 获取分享类型文本
func (ds *DeviceShare) GetShareTypeText() string {
	switch ds.ShareType {
	case 1:
		return "只读权限"
	case 2:
		return "控制权限"
	case 3:
		return "管理权限"
	default:
		return "未知权限"
	}
}

// IsExpired 判断分享是否已过期
func (ds *DeviceShare) IsExpired() bool {
	if ds.ExpiredAt == nil {
		return false
	}
	return time.Now().After(*ds.ExpiredAt)
}

// GetLogLevelText 获取日志级别文本
func (dl *DeviceLog) GetLogLevelText() string {
	switch dl.LogLevel {
	case 1:
		return "调试"
	case 2:
		return "信息"
	case 3:
		return "警告"
	case 4:
		return "错误"
	default:
		return "未知级别"
	}
}

// GetLogTypeText 获取日志类型文本
func (dl *DeviceLog) GetLogTypeText() string {
	switch dl.LogType {
	case 1:
		return "系统日志"
	case 2:
		return "业务日志"
	case 3:
		return "安全日志"
	case 4:
		return "其他日志"
	default:
		return "未知类型"
	}
}