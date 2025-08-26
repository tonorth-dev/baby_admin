package request

// MiniprogramLogin 小程序登录请求
type MiniprogramLogin struct {
	Code          string `json:"code" binding:"required" example:"微信登录凭证code"`         // 微信登录凭证
	EncryptedData string `json:"encryptedData,omitempty" example:"加密的用户数据"`          // 加密的用户数据（可选）
	IV            string `json:"iv,omitempty" example:"加密算法的初始向量"`                   // 加密算法的初始向量（可选）
	RawData       string `json:"rawData,omitempty" example:"不包含敏感信息的原始数据字符串"`        // 不包含敏感信息的原始数据字符串（可选）
	Signature     string `json:"signature,omitempty" example:"签名"`                   // 签名（可选）
}

// MiniprogramUserInfo 小程序用户信息（用于解密后的用户信息）
type MiniprogramUserInfo struct {
	OpenID    string `json:"openId"`    // 用户openid
	NickName  string `json:"nickName"`  // 用户昵称
	Gender    int    `json:"gender"`    // 用户性别 0：未知、1：男、2：女
	City      string `json:"city"`      // 用户所在城市
	Province  string `json:"province"`  // 用户所在省份
	Country   string `json:"country"`   // 用户所在国家
	Avatar    string `json:"avatarUrl"` // 用户头像
	UnionID   string `json:"unionId"`   // 用户unionid（在满足unionid获取条件的情况下返回）
	Language  string `json:"language"`  // 用户的语言，简体中文为zh_CN
}

// MiniprogramUpdateProfile 更新小程序用户信息
type MiniprogramUpdateProfile struct {
	NickName string `json:"nickName,omitempty"` // 用户昵称
	Avatar   string `json:"avatar,omitempty"`   // 用户头像
	Phone    string `json:"phone,omitempty"`    // 手机号
}

// AppRegister APP注册请求
type AppRegister struct {
	Username string `json:"username" binding:"required,min=3,max=20" validate:"alphanum"`
	Password string `json:"password" binding:"required,min=6,max=20"`
	Email    string `json:"email" binding:"required,email"`
	NickName string `json:"nickName" binding:"required,min=1,max=50"`
	Phone    string `json:"phone" binding:"omitempty,len=11" validate:"numeric"`
}

// AppLogin APP登录请求
type AppLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Captcha  string `json:"captcha" binding:"required"`
	CaptchaId string `json:"captchaId" binding:"required"`
}

// UpdateProfile 更新用户资料请求
type UpdateProfile struct {
	NickName string `json:"nickName" binding:"required,min=1,max=50"`
	Avatar   string `json:"avatar"`
	Gender   int    `json:"gender" binding:"oneof=0 1 2"`
	Phone    string `json:"phone" binding:"omitempty,len=11" validate:"numeric"`
	City     string `json:"city"`
	Province string `json:"province"`
}

// ChangePassword 修改密码请求
type ChangePassword struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required,min=6,max=20"`
}