package baby

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
	"baby_admin/server/global"
	"baby_admin/server/model/baby"
	"baby_admin/server/model/baby/request"
	"baby_admin/server/model/baby/response"
	"gorm.io/gorm"
)

type GrowthRecordService struct{}

// CreateGrowthRecord 创建成长记录
func (s *GrowthRecordService) CreateGrowthRecord(userID uint, req *request.CreateGrowthRecordRequest) error {
	// 验证宝宝是否属于当前用户
	var babyProfile baby.BabyProfile
	err := global.GVA_DB.Where("id = ? AND user_id = ?", req.BabyID, userID).First(&babyProfile).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("宝宝档案不存在")
		}
		return err
	}

	// 序列化媒体文件
	var mediaFilesJSON string
	if len(req.MediaFiles) > 0 {
		mediaFiles := make([]baby.MediaFile, len(req.MediaFiles))
		for i, mf := range req.MediaFiles {
			mediaFiles[i] = baby.MediaFile{
				URL:      mf.URL,
				Type:     mf.Type,
				Filename: mf.Filename,
				Size:     mf.Size,
			}
		}
		mediaFilesBytes, _ := json.Marshal(mediaFiles)
		mediaFilesJSON = string(mediaFilesBytes)
	}

	// 处理标签
	var tagsJSON string
	if req.Tags != "" {
		tags := strings.Split(req.Tags, ",")
		for i, tag := range tags {
			tags[i] = strings.TrimSpace(tag)
		}
		tagsBytes, _ := json.Marshal(tags)
		tagsJSON = string(tagsBytes)
	}

	record := &baby.GrowthRecord{
		UserID:     userID,
		BabyID:     req.BabyID,
		Title:      req.Title,
		Content:    req.Content,
		RecordType: req.RecordType,
		RecordDate: req.RecordDate,
		Tags:       tagsJSON,
		Weight:     req.Weight,
		Height:     req.Height,
		MediaFiles: mediaFilesJSON,
		Milestone:  req.Milestone,
		IsPrivate:  req.IsPrivate,
	}

	return global.GVA_DB.Create(record).Error
}

// GetGrowthRecord 获取成长记录详情
func (s *GrowthRecordService) GetGrowthRecord(id uint, userID uint) (*response.GrowthRecordResponse, error) {
	var record baby.GrowthRecord
	err := global.GVA_DB.Where("id = ? AND user_id = ?", id, userID).First(&record).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("成长记录不存在")
		}
		return nil, err
	}

	// 获取宝宝姓名
	var babyProfile baby.BabyProfile
	global.GVA_DB.Where("id = ?", record.BabyID).First(&babyProfile)

	var resp response.GrowthRecordResponse
	resp.FromGrowthRecord(&record, babyProfile.Name)
	return &resp, nil
}

// GetGrowthRecordList 获取成长记录列表
func (s *GrowthRecordService) GetGrowthRecordList(userID uint, req *request.GrowthRecordSearch) (*response.GrowthRecordListResponse, error) {
	db := global.GVA_DB.Model(&baby.GrowthRecord{}).Where("user_id = ?", userID)

	// 搜索条件
	if req.BabyID > 0 {
		db = db.Where("baby_id = ?", req.BabyID)
	}
	if req.RecordType > 0 {
		db = db.Where("record_type = ?", req.RecordType)
	}
	if req.StartDate != "" {
		db = db.Where("record_date >= ?", req.StartDate)
	}
	if req.EndDate != "" {
		db = db.Where("record_date <= ?", req.EndDate)
	}
	if req.Keyword != "" {
		db = db.Where("title LIKE ? OR content LIKE ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}
	if req.Tags != "" {
		db = db.Where("tags LIKE ?", "%"+req.Tags+"%")
	}

	// 获取总数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}

	// 分页查询
	var records []baby.GrowthRecord
	offset := (req.Page - 1) * req.PageSize
	err := db.Offset(offset).Limit(req.PageSize).Order("record_date DESC, created_at DESC").Find(&records).Error
	if err != nil {
		return nil, err
	}

	// 获取所有相关的宝宝信息
	var babyIDs []uint
	for _, record := range records {
		babyIDs = append(babyIDs, record.BabyID)
	}

	var babies []baby.BabyProfile
	if len(babyIDs) > 0 {
		global.GVA_DB.Where("id IN ?", babyIDs).Find(&babies)
	}

	// 创建宝宝ID到姓名的映射
	babyNameMap := make(map[uint]string)
	for _, baby := range babies {
		babyNameMap[baby.ID] = baby.Name
	}

	// 转换响应
	var list []response.GrowthRecordResponse
	for _, record := range records {
		var resp response.GrowthRecordResponse
		babyName := babyNameMap[record.BabyID]
		resp.FromGrowthRecord(&record, babyName)
		list = append(list, resp)
	}

	return &response.GrowthRecordListResponse{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

// UpdateGrowthRecord 更新成长记录
func (s *GrowthRecordService) UpdateGrowthRecord(userID uint, req *request.UpdateGrowthRecordRequest) error {
	// 检查记录是否存在
	var record baby.GrowthRecord
	err := global.GVA_DB.Where("id = ? AND user_id = ?", req.ID, userID).First(&record).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("成长记录不存在")
		}
		return err
	}

	// 验证宝宝是否属于当前用户
	var babyProfile baby.BabyProfile
	err = global.GVA_DB.Where("id = ? AND user_id = ?", req.BabyID, userID).First(&babyProfile).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("宝宝档案不存在")
		}
		return err
	}

	// 序列化媒体文件
	var mediaFilesJSON string
	if len(req.MediaFiles) > 0 {
		mediaFiles := make([]baby.MediaFile, len(req.MediaFiles))
		for i, mf := range req.MediaFiles {
			mediaFiles[i] = baby.MediaFile{
				URL:      mf.URL,
				Type:     mf.Type,
				Filename: mf.Filename,
				Size:     mf.Size,
			}
		}
		mediaFilesBytes, _ := json.Marshal(mediaFiles)
		mediaFilesJSON = string(mediaFilesBytes)
	}

	// 处理标签
	var tagsJSON string
	if req.Tags != "" {
		tags := strings.Split(req.Tags, ",")
		for i, tag := range tags {
			tags[i] = strings.TrimSpace(tag)
		}
		tagsBytes, _ := json.Marshal(tags)
		tagsJSON = string(tagsBytes)
	}

	// 更新记录
	updates := map[string]interface{}{
		"baby_id":     req.BabyID,
		"title":       req.Title,
		"content":     req.Content,
		"record_type": req.RecordType,
		"record_date": req.RecordDate,
		"tags":        tagsJSON,
		"weight":      req.Weight,
		"height":      req.Height,
		"media_files": mediaFilesJSON,
		"milestone":   req.Milestone,
		"is_private":  req.IsPrivate,
	}

	return global.GVA_DB.Model(&record).Updates(updates).Error
}

// DeleteGrowthRecord 删除成长记录
func (s *GrowthRecordService) DeleteGrowthRecord(id uint, userID uint) error {
	var record baby.GrowthRecord
	err := global.GVA_DB.Where("id = ? AND user_id = ?", id, userID).First(&record).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("成长记录不存在")
		}
		return err
	}

	return global.GVA_DB.Delete(&record).Error
}

// GetGrowthStatistics 获取成长统计
func (s *GrowthRecordService) GetGrowthStatistics(userID uint, babyID uint) (*response.GrowthStatistics, error) {
	// 验证宝宝是否属于当前用户
	var babyProfile baby.BabyProfile
	err := global.GVA_DB.Where("id = ? AND user_id = ?", babyID, userID).First(&babyProfile).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("宝宝档案不存在")
		}
		return nil, err
	}

	stats := &response.GrowthStatistics{}

	// 总记录数
	global.GVA_DB.Model(&baby.GrowthRecord{}).Where("user_id = ? AND baby_id = ?", userID, babyID).Count(&stats.TotalRecords)

	// 照片数量
	global.GVA_DB.Model(&baby.GrowthRecord{}).Where("user_id = ? AND baby_id = ? AND record_type = ?", userID, babyID, 2).Count(&stats.PhotoCount)

	// 视频数量
	global.GVA_DB.Model(&baby.GrowthRecord{}).Where("user_id = ? AND baby_id = ? AND record_type = ?", userID, babyID, 3).Count(&stats.VideoCount)

	// 里程碑数量
	global.GVA_DB.Model(&baby.GrowthRecord{}).Where("user_id = ? AND baby_id = ? AND record_type = ?", userID, babyID, 4).Count(&stats.MilestoneCount)

	// 最近记录（5条）
	var recentRecords []baby.GrowthRecord
	err = global.GVA_DB.Where("user_id = ? AND baby_id = ?", userID, babyID).Order("record_date DESC, created_at DESC").Limit(5).Find(&recentRecords).Error
	if err == nil {
		for _, record := range recentRecords {
			var resp response.GrowthRecordResponse
			resp.FromGrowthRecord(&record, babyProfile.Name)
			stats.RecentRecords = append(stats.RecentRecords, resp)
		}
	}

	// 体重记录（最近30天）
	thirtyDaysAgo := time.Now().AddDate(0, 0, -30)
	var weightRecords []baby.GrowthRecord
	err = global.GVA_DB.Where("user_id = ? AND baby_id = ? AND weight > 0 AND record_date >= ?", userID, babyID, thirtyDaysAgo).
		Order("record_date ASC").Find(&weightRecords).Error
	if err == nil {
		for _, record := range weightRecords {
			stats.WeightRecords = append(stats.WeightRecords, response.WeightRecord{
				Date:   record.RecordDate,
				Weight: record.Weight,
			})
		}
	}

	// 身高记录（最近30天）
	var heightRecords []baby.GrowthRecord
	err = global.GVA_DB.Where("user_id = ? AND baby_id = ? AND height > 0 AND record_date >= ?", userID, babyID, thirtyDaysAgo).
		Order("record_date ASC").Find(&heightRecords).Error
	if err == nil {
		for _, record := range heightRecords {
			stats.HeightRecords = append(stats.HeightRecords, response.HeightRecord{
				Date:   record.RecordDate,
				Height: record.Height,
			})
		}
	}

	return stats, nil
}