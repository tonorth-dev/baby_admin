package response

import (
	"time"
	"baby_admin/server/model/baby"
)

// BabyProfileResponse 宝宝档案响应
type BabyProfileResponse struct {
	ID         uint      `json:"id"`
	UserID     uint      `json:"user_id"`
	Name       string    `json:"name"`
	Gender     int       `json:"gender"`
	GenderText string    `json:"gender_text"`
	Birthday   time.Time `json:"birthday"`
	Avatar     string    `json:"avatar"`
	Weight     float64   `json:"weight"`
	Height     float64   `json:"height"`
	BloodType  string    `json:"blood_type"`
	Remark     string    `json:"remark"`
	IsActive   bool      `json:"is_active"`
	Age        int       `json:"age"`      // 年龄(月)
	AgeText    string    `json:"age_text"` // 年龄描述
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// FromBabyProfile 从BabyProfile模型转换
func (r *BabyProfileResponse) FromBabyProfile(profile *baby.BabyProfile) {
	r.ID = profile.ID
	r.UserID = profile.UserID
	r.Name = profile.Name
	r.Gender = profile.Gender
	r.GenderText = getGenderText(profile.Gender)
	r.Birthday = profile.Birthday
	r.Avatar = profile.Avatar
	r.Weight = profile.Weight
	r.Height = profile.Height
	r.BloodType = profile.BloodType
	r.Remark = profile.Remark
	r.IsActive = profile.IsActive
	r.Age = profile.GetAge()
	r.AgeText = profile.GetAgeText()
	r.CreatedAt = profile.CreatedAt
	r.UpdatedAt = profile.UpdatedAt
}

// getGenderText 获取性别文本
func getGenderText(gender int) string {
	switch gender {
	case 1:
		return "男"
	case 2:
		return "女"
	default:
		return "未知"
	}
}

// BabyProfileListResponse 宝宝档案列表响应
type BabyProfileListResponse struct {
	List     []BabyProfileResponse `json:"list"`
	Total    int64                 `json:"total"`
	Page     int                   `json:"page"`
	PageSize int                   `json:"page_size"`
}
