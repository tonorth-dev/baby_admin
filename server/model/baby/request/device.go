package request

import (
	"baby_admin/server/model/common/request"
	"time"
)

// DeviceSearch 设备搜索条件
type DeviceSearch struct {
	request.PageInfo
	UserID     uint   `json:"user_id" form:"user_id"`
	DeviceType int    `json:"device_type" form:"device_type"`
	Status     int    `json:"status" form:"status"`
	Location   string `json:"location" form:"location"`
	Keyword    string `json:"keyword" form:"keyword"`
}

// AddDeviceRequest 添加设备请求
type AddDeviceRequest struct {
	DeviceName   string `json:"device_name" binding:"required,min=1,max=100"`
	DeviceType   int    `json:"device_type" binding:"required,oneof=1 2 3 4"`
	ProductID    string `json:"product_id" binding:"required"`
	DeviceSecret string `json:"device_secret" binding:"required"`
	SerialNumber string `json:"serial_number"`
	MacAddress   string `json:"mac_address"`
	Location     string `json:"location"`
	Description  string `json:"description"`
}

// UpdateDeviceRequest 更新设备请求
type UpdateDeviceRequest struct {
	ID          uint   `json:"id" binding:"required"`
	DeviceName  string `json:"device_name" binding:"required,min=1,max=100"`
	Location    string `json:"location"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

// DeviceCommandRequest 设备命令请求
type DeviceCommandRequest struct {
	DeviceID    uint   `json:"device_id" binding:"required"`
	CommandType int    `json:"command_type" binding:"required,oneof=1 2 3 4 5"`
	Command     string `json:"command" binding:"required"`
	Parameters  string `json:"parameters"` // JSON格式的参数
}

// UpdateDeviceConfigRequest 更新设备配置请求
type UpdateDeviceConfigRequest struct {
	DeviceID uint                      `json:"device_id" binding:"required"`
	Configs  []DeviceConfigItemRequest `json:"configs" binding:"required,min=1"`
}

// DeviceConfigItemRequest 设备配置项请求
type DeviceConfigItemRequest struct {
	ConfigKey   string `json:"config_key" binding:"required"`
	ConfigValue string `json:"config_value" binding:"required"`
}

// ShareDeviceRequest 分享设备请求
type ShareDeviceRequest struct {
	DeviceID     uint       `json:"device_id" binding:"required"`
	SharedUserID uint       `json:"shared_user_id" binding:"required"`
	ShareType    int        `json:"share_type" binding:"required,oneof=1 2 3"`
	Permissions  string     `json:"permissions"` // JSON格式的权限配置
	ExpiredAt    *time.Time `json:"expired_at"`
}

// AcceptShareRequest 接受分享请求
type AcceptShareRequest struct {
	ShareID uint `json:"share_id" binding:"required"`
}

// DeviceLogSearch 设备日志搜索条件
type DeviceLogSearch struct {
	request.PageInfo
	DeviceID  uint   `json:"device_id" form:"device_id"`
	LogLevel  int    `json:"log_level" form:"log_level"`
	LogType   int    `json:"log_type" form:"log_type"`
	StartDate string `json:"start_date" form:"start_date"`
	EndDate   string `json:"end_date" form:"end_date"`
	Keyword   string `json:"keyword" form:"keyword"`
}

// DeviceStatusSearch 设备状态搜索条件
type DeviceStatusSearch struct {
	request.PageInfo
	DeviceID   uint   `json:"device_id" form:"device_id"`
	StatusType int    `json:"status_type" form:"status_type"`
	StartDate  string `json:"start_date" form:"start_date"`
	EndDate    string `json:"end_date" form:"end_date"`
	IsAlert    *bool  `json:"is_alert" form:"is_alert"`
}

// CreateDeviceLogRequest 创建设备日志请求
type CreateDeviceLogRequest struct {
	DeviceID uint   `json:"device_id" binding:"required"`
	LogLevel int    `json:"log_level" binding:"required,oneof=1 2 3 4"`
	LogType  int    `json:"log_type" binding:"required,oneof=1 2 3 4"`
	Message  string `json:"message" binding:"required"`
	Details  string `json:"details"`
	Source   string `json:"source"`
}

// CreateDeviceStatusRequest 创建设备状态请求
type CreateDeviceStatusRequest struct {
	DeviceID    uint   `json:"device_id" binding:"required"`
	StatusType  int    `json:"status_type" binding:"required"`
	StatusValue string `json:"status_value"`
	StatusData  string `json:"status_data"`
	IsAlert     bool   `json:"is_alert"`
}
