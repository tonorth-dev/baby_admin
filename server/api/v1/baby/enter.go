package baby

import "baby_admin/server/service"

type ApiGroup struct {
	BabyProfileApi
	MusicApi
}

var (
	babyProfileService = service.ServiceGroupApp.BabyServiceGroup.BabyProfileService
	musicService       = service.ServiceGroupApp.BabyServiceGroup.MusicService
)
