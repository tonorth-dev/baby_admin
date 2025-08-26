package baby

import "baby_admin/server/api/v1"

type RouterGroup struct {
	BabyProfileRouter
	MusicRouter
}

var (
	babyProfileApi = v1.ApiGroupApp.BabyApiGroup.BabyProfileApi
	musicApi       = v1.ApiGroupApp.BabyApiGroup.MusicApi
)