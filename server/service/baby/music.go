package baby

import (
	"errors"
	"fmt"
	"baby_admin/server/global"
	"baby_admin/server/model/baby"
	"baby_admin/server/model/baby/request"
	"baby_admin/server/model/baby/response"
	commonRequest "baby_admin/server/model/common/request"
	"gorm.io/gorm"
)

type MusicService struct{}

// GetMusicCategories 获取音乐分类列表
func (s *MusicService) GetMusicCategories() ([]response.MusicCategoryResponse, error) {
	var categories []baby.MusicCategory
	err := global.GVA_DB.Where("is_active = ?", true).Order("sort_order ASC, id ASC").Find(&categories).Error
	if err != nil {
		return nil, err
	}

	var result []response.MusicCategoryResponse
	for _, category := range categories {
		// 统计该分类下的音乐数量
		var count int64
		global.GVA_DB.Model(&baby.Music{}).Where("category_id = ? AND is_active = ?", category.ID, true).Count(&count)

		result = append(result, response.MusicCategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			Icon:        category.Icon,
			AgeRange:    category.AgeRange,
			SortOrder:   category.SortOrder,
			MusicCount:  count,
			IsActive:    category.IsActive,
		})
	}

	return result, nil
}

// GetMusicList 获取音乐列表
func (s *MusicService) GetMusicList(userID uint, req *request.MusicSearch) (*response.MusicListResponse, error) {
	db := global.GVA_DB.Model(&baby.Music{}).Where("is_active = ?", true)

	// 搜索条件
	if req.CategoryID > 0 {
		db = db.Where("category_id = ?", req.CategoryID)
	}
	if req.AgeRange != "" {
		db = db.Where("age_range = ? OR age_range = ''", req.AgeRange)
	}
	if req.Keyword != "" {
		db = db.Where("title LIKE ? OR artist LIKE ? OR description LIKE ?", 
			"%"+req.Keyword+"%", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}
	if req.Tags != "" {
		db = db.Where("tags LIKE ?", "%"+req.Tags+"%")
	}
	if req.IsVIP != nil {
		db = db.Where("is_vip = ?", *req.IsVIP)
	}

	// 获取总数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}

	// 分页查询
	var musics []baby.Music
	offset := (req.Page - 1) * req.PageSize
	err := db.Offset(offset).Limit(req.PageSize).Order("sort_order ASC, play_count DESC, id DESC").Find(&musics).Error
	if err != nil {
		return nil, err
	}

	// 获取分类信息
	var categoryIDs []uint
	for _, music := range musics {
		categoryIDs = append(categoryIDs, music.CategoryID)
	}

	var categories []baby.MusicCategory
	if len(categoryIDs) > 0 {
		global.GVA_DB.Where("id IN ?", categoryIDs).Find(&categories)
	}

	categoryMap := make(map[uint]string)
	for _, category := range categories {
		categoryMap[category.ID] = category.Name
	}

	// 获取用户收藏信息
	var musicIDs []uint
	for _, music := range musics {
		musicIDs = append(musicIDs, music.ID)
	}

	var favorites []baby.UserMusicFavorite
	if len(musicIDs) > 0 {
		global.GVA_DB.Where("user_id = ? AND music_id IN ?", userID, musicIDs).Find(&favorites)
	}

	favoriteMap := make(map[uint]bool)
	for _, favorite := range favorites {
		favoriteMap[favorite.MusicID] = true
	}

	// 转换响应
	var list []response.MusicResponse
	for _, music := range musics {
		var resp response.MusicResponse
		categoryName := categoryMap[music.CategoryID]
		isFavorited := favoriteMap[music.ID]
		resp.FromMusic(&music, categoryName, isFavorited)
		list = append(list, resp)
	}

	return &response.MusicListResponse{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

// GetMusicDetail 获取音乐详情
func (s *MusicService) GetMusicDetail(userID uint, musicID uint) (*response.MusicResponse, error) {
	var music baby.Music
	err := global.GVA_DB.Where("id = ? AND is_active = ?", musicID, true).First(&music).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("音乐不存在")
		}
		return nil, err
	}

	// 获取分类名称
	var category baby.MusicCategory
	global.GVA_DB.Where("id = ?", music.CategoryID).First(&category)

	// 检查是否收藏
	var favorite baby.UserMusicFavorite
	isFavorited := false
	err = global.GVA_DB.Where("user_id = ? AND music_id = ?", userID, musicID).First(&favorite).Error
	if err == nil {
		isFavorited = true
	}

	var resp response.MusicResponse
	resp.FromMusic(&music, category.Name, isFavorited)
	return &resp, nil
}

// PlayMusic 播放音乐（记录播放历史）
func (s *MusicService) PlayMusic(userID uint, req *request.PlayMusicRequest) error {
	// 检查音乐是否存在
	var music baby.Music
	err := global.GVA_DB.Where("id = ? AND is_active = ?", req.MusicID, true).First(&music).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("音乐不存在")
		}
		return err
	}

	// 如果指定了宝宝ID，验证宝宝是否属于当前用户
	if req.BabyID > 0 {
		var babyProfile baby.BabyProfile
		err = global.GVA_DB.Where("id = ? AND user_id = ?", req.BabyID, userID).First(&babyProfile).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("宝宝档案不存在")
			}
			return err
		}
	}

	// 更新音乐播放次数
	global.GVA_DB.Model(&music).Update("play_count", gorm.Expr("play_count + ?", 1))

	// 记录播放历史
	history := &baby.UserMusicHistory{
		UserID:     userID,
		MusicID:    req.MusicID,
		BabyID:     req.BabyID,
		PlayTime:   req.PlayTime,
		IsFinished: req.PlayTime >= music.Duration,
	}

	return global.GVA_DB.Create(history).Error
}

// ToggleFavorite 切换收藏状态
func (s *MusicService) ToggleFavorite(userID uint, musicID uint) (bool, error) {
	// 检查音乐是否存在
	var music baby.Music
	err := global.GVA_DB.Where("id = ? AND is_active = ?", musicID, true).First(&music).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.New("音乐不存在")
		}
		return false, err
	}

	// 检查是否已收藏
	var favorite baby.UserMusicFavorite
	err = global.GVA_DB.Where("user_id = ? AND music_id = ?", userID, musicID).First(&favorite).Error
	
	if err == nil {
		// 已收藏，取消收藏
		err = global.GVA_DB.Delete(&favorite).Error
		if err != nil {
			return false, err
		}
		// 更新音乐点赞数
		global.GVA_DB.Model(&music).Update("like_count", gorm.Expr("like_count - ?", 1))
		return false, nil
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		// 未收藏，添加收藏
		newFavorite := &baby.UserMusicFavorite{
			UserID:  userID,
			MusicID: musicID,
		}
		err = global.GVA_DB.Create(newFavorite).Error
		if err != nil {
			return false, err
		}
		// 更新音乐点赞数
		global.GVA_DB.Model(&music).Update("like_count", gorm.Expr("like_count + ?", 1))
		return true, nil
	}

	return false, err
}

// GetUserFavorites 获取用户收藏的音乐
func (s *MusicService) GetUserFavorites(userID uint, req *commonRequest.PageInfo) (*response.MusicListResponse, error) {
	// 获取用户收藏的音乐ID列表
	var favorites []baby.UserMusicFavorite
	db := global.GVA_DB.Where("user_id = ?", userID)
	
	// 获取总数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}

	// 分页查询
	offset := (req.Page - 1) * req.PageSize
	err := db.Offset(offset).Limit(req.PageSize).Order("created_at DESC").Find(&favorites).Error
	if err != nil {
		return nil, err
	}

	if len(favorites) == 0 {
		return &response.MusicListResponse{
			List:     []response.MusicResponse{},
			Total:    0,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, nil
	}

	// 获取音乐详情
	var musicIDs []uint
	for _, favorite := range favorites {
		musicIDs = append(musicIDs, favorite.MusicID)
	}

	var musics []baby.Music
	err = global.GVA_DB.Where("id IN ? AND is_active = ?", musicIDs, true).Find(&musics).Error
	if err != nil {
		return nil, err
	}

	// 获取分类信息
	var categoryIDs []uint
	for _, music := range musics {
		categoryIDs = append(categoryIDs, music.CategoryID)
	}

	var categories []baby.MusicCategory
	if len(categoryIDs) > 0 {
		global.GVA_DB.Where("id IN ?", categoryIDs).Find(&categories)
	}

	categoryMap := make(map[uint]string)
	for _, category := range categories {
		categoryMap[category.ID] = category.Name
	}

	// 转换响应
	var list []response.MusicResponse
	for _, music := range musics {
		var resp response.MusicResponse
		categoryName := categoryMap[music.CategoryID]
		resp.FromMusic(&music, categoryName, true) // 收藏列表中的都是已收藏的
		list = append(list, resp)
	}

	return &response.MusicListResponse{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

// GetPlayHistory 获取播放历史
func (s *MusicService) GetPlayHistory(userID uint, req *commonRequest.PageInfo) ([]response.MusicHistoryResponse, error) {
	var histories []baby.UserMusicHistory
	db := global.GVA_DB.Where("user_id = ?", userID)
	
	offset := (req.Page - 1) * req.PageSize
	err := db.Offset(offset).Limit(req.PageSize).Order("created_at DESC").Find(&histories).Error
	if err != nil {
		return nil, err
	}

	if len(histories) == 0 {
		return []response.MusicHistoryResponse{}, nil
	}

	// 获取相关的音乐和宝宝信息
	var musicIDs []uint
	var babyIDs []uint
	for _, history := range histories {
		musicIDs = append(musicIDs, history.MusicID)
		if history.BabyID > 0 {
			babyIDs = append(babyIDs, history.BabyID)
		}
	}

	// 获取音乐信息
	var musics []baby.Music
	if len(musicIDs) > 0 {
		global.GVA_DB.Where("id IN ?", musicIDs).Find(&musics)
	}

	musicMap := make(map[uint]baby.Music)
	for _, music := range musics {
		musicMap[music.ID] = music
	}

	// 获取宝宝信息
	var babies []baby.BabyProfile
	if len(babyIDs) > 0 {
		global.GVA_DB.Where("id IN ?", babyIDs).Find(&babies)
	}

	babyMap := make(map[uint]string)
	for _, baby := range babies {
		babyMap[baby.ID] = baby.Name
	}

	// 获取分类信息
	var categoryIDs []uint
	for _, music := range musics {
		categoryIDs = append(categoryIDs, music.CategoryID)
	}

	var categories []baby.MusicCategory
	if len(categoryIDs) > 0 {
		global.GVA_DB.Where("id IN ?", categoryIDs).Find(&categories)
	}

	categoryMap := make(map[uint]string)
	for _, category := range categories {
		categoryMap[category.ID] = category.Name
	}

	// 转换响应
	var result []response.MusicHistoryResponse
	for _, history := range histories {
		music, exists := musicMap[history.MusicID]
		if !exists {
			continue
		}

		var musicResp response.MusicResponse
		categoryName := categoryMap[music.CategoryID]
		musicResp.FromMusic(&music, categoryName, false) // 这里不需要检查收藏状态

		babyName := ""
		if history.BabyID > 0 {
			babyName = babyMap[history.BabyID]
		}

		result = append(result, response.MusicHistoryResponse{
			ID:         history.ID,
			Music:      musicResp,
			BabyID:     history.BabyID,
			BabyName:   babyName,
			PlayTime:   history.PlayTime,
			IsFinished: history.IsFinished,
			PlayedAt:   history.CreatedAt,
		})
	}

	return result, nil
}

// GetRecommendations 获取音乐推荐
func (s *MusicService) GetRecommendations(userID uint, babyID uint) ([]response.RecommendationResponse, error) {
	var recommendations []response.RecommendationResponse

	// 如果指定了宝宝ID，获取宝宝信息进行基于年龄的推荐
	if babyID > 0 {
		var babyProfile baby.BabyProfile
		err := global.GVA_DB.Where("id = ? AND user_id = ?", babyID, userID).First(&babyProfile).Error
		if err == nil {
			ageInMonths := babyProfile.GetAge()
			ageRange := getAgeRange(ageInMonths)
			
			var ageBasedMusics []baby.Music
			global.GVA_DB.Where("age_range = ? AND is_active = ?", ageRange, true).
				Order("play_count DESC").Limit(10).Find(&ageBasedMusics)

			if len(ageBasedMusics) > 0 {
				var musicResponses []response.MusicResponse
				for _, music := range ageBasedMusics {
					var resp response.MusicResponse
					resp.FromMusic(&music, "", false)
					musicResponses = append(musicResponses, resp)
				}

				recommendations = append(recommendations, response.RecommendationResponse{
					Title:       fmt.Sprintf("适合%s的音乐", babyProfile.GetAgeText()),
					Description: "根据宝宝年龄推荐的音乐",
					Musics:      musicResponses,
					Type:        "age_based",
				})
			}
		}
	}

	// 安抚催眠音乐推荐
	var sleepMusics []baby.Music
	global.GVA_DB.Where("tags LIKE ? AND is_active = ?", "%催眠%", true).
		Order("play_count DESC").Limit(8).Find(&sleepMusics)

	if len(sleepMusics) > 0 {
		var musicResponses []response.MusicResponse
		for _, music := range sleepMusics {
			var resp response.MusicResponse
			resp.FromMusic(&music, "", false)
			musicResponses = append(musicResponses, resp)
		}

		recommendations = append(recommendations, response.RecommendationResponse{
			Title:       "安抚催眠",
			Description: "让宝宝安然入睡的音乐",
			Musics:      musicResponses,
			Type:        "sleep",
		})
	}

	// 热门音乐推荐
	var popularMusics []baby.Music
	global.GVA_DB.Where("is_active = ?", true).
		Order("play_count DESC").Limit(10).Find(&popularMusics)

	if len(popularMusics) > 0 {
		var musicResponses []response.MusicResponse
		for _, music := range popularMusics {
			var resp response.MusicResponse
			resp.FromMusic(&music, "", false)
			musicResponses = append(musicResponses, resp)
		}

		recommendations = append(recommendations, response.RecommendationResponse{
			Title:       "热门音乐",
			Description: "最受欢迎的婴幼儿音乐",
			Musics:      musicResponses,
			Type:        "popular",
		})
	}

	return recommendations, nil
}

// getAgeRange 根据月龄获取年龄段
func getAgeRange(ageInMonths int) string {
	if ageInMonths <= 6 {
		return "0-6月"
	} else if ageInMonths <= 12 {
		return "6-12月"
	} else if ageInMonths <= 24 {
		return "1-2岁"
	} else if ageInMonths <= 36 {
		return "2-3岁"
	} else {
		return "3岁+"
	}
}