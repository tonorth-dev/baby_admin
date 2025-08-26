package baby

import (
	"fmt"
	"time"
	"baby_admin/server/global"
)

// BabyProfile 宝宝档案表
type BabyProfile struct {
	global.GVA_MODEL
	UserID      uint      `json:"user_id" gorm:"not null;comment:用户ID"`
	Name        string    `json:"name" gorm:"size:50;not null;comment:宝宝姓名"`
	Gender      int       `json:"gender" gorm:"not null;default:0;comment:性别:0未知,1男,2女"`
	Birthday    time.Time `json:"birthday" gorm:"not null;comment:出生日期"`
	Avatar      string    `json:"avatar" gorm:"size:255;comment:头像URL"`
	Weight      float64   `json:"weight" gorm:"type:decimal(5,2);comment:体重(kg)"`
	Height      float64   `json:"height" gorm:"type:decimal(5,2);comment:身高(cm)"`
	BloodType   string    `json:"blood_type" gorm:"size:10;comment:血型"`
	Remark      string    `json:"remark" gorm:"type:text;comment:备注"`
	IsActive    bool      `json:"is_active" gorm:"default:true;comment:是否为当前活跃宝宝"`
}

// TableName 指定表名
func (BabyProfile) TableName() string {
	return "baby_profiles"
}

// GetAge 计算年龄(月)
func (b *BabyProfile) GetAge() int {
	now := time.Now()
	months := int(now.Month()) - int(b.Birthday.Month())
	years := now.Year() - b.Birthday.Year()
	
	if months < 0 {
		years--
		months += 12
	}
	
	return years*12 + months
}

// GetAgeText 获取年龄文本描述
func (b *BabyProfile) GetAgeText() string {
	months := b.GetAge()
	years := months / 12
	remainingMonths := months % 12
	
	if years > 0 {
		if remainingMonths > 0 {
			return fmt.Sprintf("%d岁%d个月", years, remainingMonths)
		}
		return fmt.Sprintf("%d岁", years)
	}
	return fmt.Sprintf("%d个月", remainingMonths)
}