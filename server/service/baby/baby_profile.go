package baby

import (
	"errors"
	"baby_admin/server/global"
	"baby_admin/server/model/baby"
	"baby_admin/server/model/baby/request"
	"baby_admin/server/model/baby/response"
	"gorm.io/gorm"
)

type BabyProfileService struct{}

// CreateBabyProfile 创建宝宝档案
func (s *BabyProfileService) CreateBabyProfile(userID uint, req *request.CreateBabyProfileRequest) error {
	// 如果设置为活跃宝宝，先将其他宝宝设置为非活跃
	if req.IsActive {
		if err := s.deactivateOtherBabies(userID); err != nil {
			return err
		}
	}

	profile := req.ToBabyProfile(userID)
	return global.GVA_DB.Create(profile).Error
}

// GetBabyProfile 获取宝宝档案详情
func (s *BabyProfileService) GetBabyProfile(id uint, userID uint) (*response.BabyProfileResponse, error) {
	var profile baby.BabyProfile
	err := global.GVA_DB.Where("id = ? AND user_id = ?", id, userID).First(&profile).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("宝宝档案不存在")
		}
		return nil, err
	}

	var resp response.BabyProfileResponse
	resp.FromBabyProfile(&profile)
	return &resp, nil
}

// GetActiveBabyProfile 获取当前活跃的宝宝档案
func (s *BabyProfileService) GetActiveBabyProfile(userID uint) (*response.BabyProfileResponse, error) {
	var profile baby.BabyProfile
	err := global.GVA_DB.Where("user_id = ? AND is_active = ?", userID, true).First(&profile).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("未找到活跃的宝宝档案")
		}
		return nil, err
	}

	var resp response.BabyProfileResponse
	resp.FromBabyProfile(&profile)
	return &resp, nil
}

// GetBabyProfileList 获取宝宝档案列表
func (s *BabyProfileService) GetBabyProfileList(userID uint, req *request.BabyProfileSearch) (*response.BabyProfileListResponse, error) {
	db := global.GVA_DB.Model(&baby.BabyProfile{}).Where("user_id = ?", userID)

	// 搜索条件
	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.Gender > 0 {
		db = db.Where("gender = ?", req.Gender)
	}

	// 获取总数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}

	// 分页查询
	var profiles []baby.BabyProfile
	offset := (req.Page - 1) * req.PageSize
	err := db.Offset(offset).Limit(req.PageSize).Order("is_active DESC, updated_at DESC").Find(&profiles).Error
	if err != nil {
		return nil, err
	}

	// 转换响应
	var list []response.BabyProfileResponse
	for _, profile := range profiles {
		var resp response.BabyProfileResponse
		resp.FromBabyProfile(&profile)
		list = append(list, resp)
	}

	return &response.BabyProfileListResponse{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

// UpdateBabyProfile 更新宝宝档案
func (s *BabyProfileService) UpdateBabyProfile(userID uint, req *request.UpdateBabyProfileRequest) error {
	// 检查宝宝档案是否存在
	var profile baby.BabyProfile
	err := global.GVA_DB.Where("id = ? AND user_id = ?", req.ID, userID).First(&profile).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("宝宝档案不存在")
		}
		return err
	}

	// 如果设置为活跃宝宝，先将其他宝宝设置为非活跃
	if req.IsActive && !profile.IsActive {
		if err := s.deactivateOtherBabies(userID); err != nil {
			return err
		}
	}

	// 更新档案信息
	updates := map[string]interface{}{
		"name":       req.Name,
		"gender":     req.Gender,
		"birthday":   req.Birthday,
		"avatar":     req.Avatar,
		"weight":     req.Weight,
		"height":     req.Height,
		"blood_type": req.BloodType,
		"remark":     req.Remark,
		"is_active":  req.IsActive,
	}

	return global.GVA_DB.Model(&profile).Updates(updates).Error
}

// DeleteBabyProfile 删除宝宝档案
func (s *BabyProfileService) DeleteBabyProfile(id uint, userID uint) error {
	// 检查宝宝档案是否存在
	var profile baby.BabyProfile
	err := global.GVA_DB.Where("id = ? AND user_id = ?", id, userID).First(&profile).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("宝宝档案不存在")
		}
		return err
	}

	return global.GVA_DB.Delete(&profile).Error
}

// SetActiveBaby 设置活跃宝宝
func (s *BabyProfileService) SetActiveBaby(id uint, userID uint) error {
	// 检查宝宝档案是否存在
	var profile baby.BabyProfile
	err := global.GVA_DB.Where("id = ? AND user_id = ?", id, userID).First(&profile).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("宝宝档案不存在")
		}
		return err
	}

	// 先将其他宝宝设置为非活跃
	if err := s.deactivateOtherBabies(userID); err != nil {
		return err
	}

	// 设置当前宝宝为活跃
	return global.GVA_DB.Model(&profile).Update("is_active", true).Error
}

// deactivateOtherBabies 将其他宝宝设置为非活跃
func (s *BabyProfileService) deactivateOtherBabies(userID uint) error {
	return global.GVA_DB.Model(&baby.BabyProfile{}).
		Where("user_id = ? AND is_active = ?", userID, true).
		Update("is_active", false).Error
}