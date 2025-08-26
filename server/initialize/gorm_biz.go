package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"baby_admin/server/model/baby"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(
		// 婴儿陪护相关模型
		&baby.BabyProfile{},
		&baby.GrowthRecord{},
		&baby.MusicCategory{},
		&baby.Music{},
		&baby.UserMusicHistory{},
		&baby.UserMusicFavorite{},
		&baby.Playlist{},
		&baby.PlaylistMusic{},
		// 智能分析相关模型
		&baby.SleepRecord{},
		&baby.CryDetection{},
		&baby.MovementDetection{},
		&baby.EnvironmentData{},
		&baby.SmartAlert{},
		&baby.AnalysisReport{},
		// 育儿指导相关模型
		&baby.ParentingCategory{},
		&baby.ParentingArticle{},
		&baby.ParentingVideo{},
		&baby.ParentingMilestone{},
		&baby.UserArticleRead{},
		&baby.UserVideoWatch{},
		&baby.UserContentFavorite{},
		// 设备管理相关模型
		&baby.Device{},
		&baby.DeviceConfig{},
		&baby.DeviceCommand{},
		&baby.DeviceStatus{},
		&baby.DeviceShare{},
		&baby.DeviceLog{},
	)
	if err != nil {
		return err
	}
	return nil
}
