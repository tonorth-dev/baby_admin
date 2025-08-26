package response

import (
	"time"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/google/uuid"
)

// MiniprogramLoginResponse 小程序登录响应
type MiniprogramLoginResponse struct {
	User      system.MiniprogramUser `json:"user"`      // 用户信息
	Token     string                 `json:"token"`     // token
	ExpiresAt int64                  `json:"expiresAt"` // 过期时间
}

// MiniprogramUserResponse 小程序用户信息响应
type MiniprogramUserResponse struct {
	User system.MiniprogramUser `json:"user"` // 用户信息
}

// WechatSession 微信登录凭证校验响应
type WechatSession struct {
	OpenID     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionID    string `json:"unionid"`     // 用户在开放平台的唯一标识符（当且仅当在微信开放平台下的应用中）
	ErrCode    int    `json:"errcode"`     // 错误码
	ErrMsg     string `json:"errmsg"`      // 错误信息
}

// AppLoginResponse APP登录响应
type AppLoginResponse struct {
	User      UserInfo `json:"user"`
	Token     string   `json:"token"`
	ExpiresAt int64    `json:"expiresAt"`
}

// AppRegisterResponse APP注册响应
type AppRegisterResponse struct {
	User      UserInfo `json:"user"`
	Token     string   `json:"token"`
	ExpiresAt int64    `json:"expiresAt"`
}

// UserInfo 用户信息（脱敏后）
type UserInfo struct {
	ID        uint      `json:"id"`
	UUID      uuid.UUID `json:"uuid"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	NickName  string    `json:"nickName"`
	Avatar    string    `json:"avatar"`
	Gender    int       `json:"gender"`
	Phone     string    `json:"phone"`
	City      string    `json:"city"`
	Province  string    `json:"province"`
	Country   string    `json:"country"`
	LoginType int       `json:"loginType"`
	Enable    int       `json:"enable"`
	LastLogin *time.Time `json:"lastLogin"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// FromMiniprogramUser 从MiniprogramUser转换
func (u *UserInfo) FromMiniprogramUser(user *system.MiniprogramUser) {
	u.ID = user.ID
	u.UUID = user.UUID
	u.Username = user.Username
	u.Email = user.Email
	u.NickName = user.NickName
	u.Avatar = user.Avatar
	u.Gender = user.Gender
	u.Phone = user.Phone
	u.City = user.City
	u.Province = user.Province
	u.Country = user.Country
	u.LoginType = user.LoginType
	u.Enable = user.Enable
	u.LastLogin = user.LastLogin
	u.CreatedAt = user.CreatedAt
	u.UpdatedAt = user.UpdatedAt
}