package system

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/google/uuid"
)

var _ Login = new(MiniprogramUser)

// MiniprogramUser 微信小程序用户表（现支持APP常规注册登录）
type MiniprogramUser struct {
	global.GVA_MODEL
	UUID       uuid.UUID      `json:"uuid" gorm:"index;comment:用户UUID"`                    // 用户UUID
	OpenID     string         `json:"openId" gorm:"unique;comment:微信OpenID"`                // 微信OpenID（可为空，支持非微信注册）
	UnionID    string         `json:"unionId" gorm:"comment:微信UnionID"`                     // 微信UnionID
	Username   string         `json:"username" gorm:"unique;comment:用户名"`                   // 用户名（APP注册使用）
	Password   string         `json:"-" gorm:"comment:密码哈希"`                                // 密码哈希，不返回给前端
	Email      string         `json:"email" gorm:"unique;comment:邮箱"`                       // 邮箱
	NickName   string         `json:"nickName" gorm:"comment:用户昵称"`                        // 用户昵称
	Avatar     string         `json:"avatar" gorm:"comment:用户头像"`                          // 用户头像
	Gender     int            `json:"gender" gorm:"default:0;comment:性别 0未知 1男 2女"`        // 性别 0未知 1男 2女
	Language   string         `json:"language" gorm:"comment:语言"`                          // 语言
	City       string         `json:"city" gorm:"comment:城市"`                              // 城市
	Province   string         `json:"province" gorm:"comment:省份"`                          // 省份
	Country    string         `json:"country" gorm:"comment:国家"`                           // 国家
	Phone      string         `json:"phone" gorm:"comment:手机号"`                            // 手机号
	Enable     int            `json:"enable" gorm:"default:1;comment:用户状态 1正常 2冻结"`        // 用户状态 1正常 2冻结
	LoginType  int            `json:"loginType" gorm:"default:0;comment:登录类型 0微信 1账号密码"`   // 登录类型 0微信 1账号密码
	LastLogin  *time.Time     `json:"lastLogin" gorm:"comment:最后登录时间"`                     // 最后登录时间
	SessionKey string         `json:"-" gorm:"comment:微信会话密钥"`                             // 微信会话密钥，不返回给前端
}

func (MiniprogramUser) TableName() string {
	return "miniprogram_users"
}

// Login 接口实现
func (m *MiniprogramUser) GetUsername() string {
	if m.Username != "" {
		return m.Username
	}
	return m.OpenID
}

func (m *MiniprogramUser) GetNickname() string {
	return m.NickName
}

func (m *MiniprogramUser) GetUUID() uuid.UUID {
	return m.UUID
}

func (m *MiniprogramUser) GetUserId() uint {
	return m.ID
}

func (m *MiniprogramUser) GetAuthorityId() uint {
	// 小程序用户使用固定的权限角色ID，比如999，需要在权限表中创建对应角色
	return 999
}

func (m *MiniprogramUser) GetUserInfo() any {
	return *m
}