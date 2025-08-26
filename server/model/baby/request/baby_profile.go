package request

import (
	"baby_admin/server/model/baby"
	"baby_admin/server/model/common/request"
	"time"
)

// BabyProfileSearch 宝宝档案搜索条件
type BabyProfileSearch struct {
	request.PageInfo
	UserID uint   `json:"user_id" form:"user_id"`
	Name   string `json:"name" form:"name"`
	Gender int    `json:"gender" form:"gender"`
}

// CreateBabyProfileRequest 创建宝宝档案请求
type CreateBabyProfileRequest struct {
	Name      string    `json:"name" binding:"required" validate:"min=1,max=50"`
	Gender    int       `json:"gender" binding:"required,oneof=1 2"`
	Birthday  time.Time `json:"birthday" binding:"required"`
	Avatar    string    `json:"avatar"`
	Weight    float64   `json:"weight"`
	Height    float64   `json:"height"`
	BloodType string    `json:"blood_type"`
	Remark    string    `json:"remark"`
	IsActive  bool      `json:"is_active"`
}

// UpdateBabyProfileRequest 更新宝宝档案请求
type UpdateBabyProfileRequest struct {
	ID        uint      `json:"id" binding:"required"`
	Name      string    `json:"name" binding:"required" validate:"min=1,max=50"`
	Gender    int       `json:"gender" binding:"required,oneof=1 2"`
	Birthday  time.Time `json:"birthday" binding:"required"`
	Avatar    string    `json:"avatar"`
	Weight    float64   `json:"weight"`
	Height    float64   `json:"height"`
	BloodType string    `json:"blood_type"`
	Remark    string    `json:"remark"`
	IsActive  bool      `json:"is_active"`
}

// ToBabyProfile 转换为BabyProfile模型
func (req *CreateBabyProfileRequest) ToBabyProfile(userID uint) *baby.BabyProfile {
	return &baby.BabyProfile{
		UserID:    userID,
		Name:      req.Name,
		Gender:    req.Gender,
		Birthday:  req.Birthday,
		Avatar:    req.Avatar,
		Weight:    req.Weight,
		Height:    req.Height,
		BloodType: req.BloodType,
		Remark:    req.Remark,
		IsActive:  req.IsActive,
	}
}
