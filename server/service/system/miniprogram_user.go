package system

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/captcha"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type MiniprogramUserService struct{}

// Login 小程序用户登录
func (m *MiniprogramUserService) Login(loginReq request.MiniprogramLogin) (userInter system.Login, token string, expiresAt int64, err error) {
	// 1. 调用微信接口获取session_key和openid
	session, err := m.getWechatSession(loginReq.Code)
	if err != nil {
		return nil, "", 0, err
	}

	if session.ErrCode != 0 {
		return nil, "", 0, fmt.Errorf("微信登录失败: %s", session.ErrMsg)
	}

	// 2. 查找或创建用户
	var user system.MiniprogramUser
	err = global.GVA_DB.Where("open_id = ?", session.OpenID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 首次登录，创建新用户
			now := time.Now()
			user = system.MiniprogramUser{
				UUID:       uuid.New(),
				OpenID:     session.OpenID,
				UnionID:    session.UnionID,
				SessionKey: session.SessionKey,
				Enable:     1,
				LastLogin:  &now,
			}

			// 如果有加密的用户信息，尝试解密
			if loginReq.EncryptedData != "" && loginReq.IV != "" {
				userInfo, decryptErr := m.decryptUserInfo(loginReq.EncryptedData, loginReq.IV, session.SessionKey)
				if decryptErr == nil {
					user.NickName = userInfo.NickName
					user.Avatar = userInfo.Avatar
					user.Gender = userInfo.Gender
					user.City = userInfo.City
					user.Province = userInfo.Province
					user.Country = userInfo.Country
					user.Language = userInfo.Language
				} else {
					global.GVA_LOG.Warn("解密用户信息失败", zap.Error(decryptErr))
				}
			}

			err = global.GVA_DB.Create(&user).Error
			if err != nil {
				return nil, "", 0, err
			}
		} else {
			return nil, "", 0, err
		}
	} else {
		// 用户已存在，更新session_key和最后登录时间
		now := time.Now()
		user.SessionKey = session.SessionKey
		user.LastLogin = &now
		if session.UnionID != "" {
			user.UnionID = session.UnionID
		}
		
		err = global.GVA_DB.Save(&user).Error
		if err != nil {
			return nil, "", 0, err
		}
	}

	// 3. 生成JWT token
	j := utils.NewJWT()
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		Username:    user.OpenID,
		NickName:    user.NickName,
		AuthorityId: user.GetAuthorityId(),
	})
	token, err = j.CreateToken(claims)
	if err != nil {
		return nil, "", 0, err
	}

	// 4. 设置Redis缓存（如果启用）
	if global.GVA_CONFIG.System.UseMultipoint {
		err = utils.SetRedisJWT(token, user.OpenID)
		if err != nil {
			return nil, "", 0, err
		}
	}

	return &user, token, claims.ExpiresAt.Unix(), nil
}

// GetUserInfo 获取小程序用户信息
func (m *MiniprogramUserService) GetUserInfo(uuid string) (user system.MiniprogramUser, err error) {
	err = global.GVA_DB.Where("uuid = ?", uuid).First(&user).Error
	return
}

// UpdateUserProfile 更新用户资料
func (m *MiniprogramUserService) UpdateUserProfile(uuid string, updateReq request.MiniprogramUpdateProfile) (err error) {
	updates := make(map[string]interface{})
	
	if updateReq.NickName != "" {
		updates["nick_name"] = updateReq.NickName
	}
	if updateReq.Avatar != "" {
		updates["avatar"] = updateReq.Avatar
	}
	if updateReq.Phone != "" {
		updates["phone"] = updateReq.Phone
	}
	
	if len(updates) > 0 {
		err = global.GVA_DB.Model(&system.MiniprogramUser{}).Where("uuid = ?", uuid).Updates(updates).Error
	}
	return
}

// getWechatSession 调用微信接口获取session信息
func (m *MiniprogramUserService) getWechatSession(code string) (*response.WechatSession, error) {
	// 从配置文件读取小程序配置
	appID := global.GVA_CONFIG.Miniprogram.AppID
	appSecret := global.GVA_CONFIG.Miniprogram.AppSecret
	
	if appID == "" || appSecret == "" {
		return nil, errors.New("小程序配置不完整，请检查config.yaml中的miniprogram配置")
	}
	
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", 
		appID, appSecret, code)
	
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	var session response.WechatSession
	err = json.Unmarshal(body, &session)
	if err != nil {
		return nil, err
	}
	
	return &session, nil
}

// decryptUserInfo 解密微信用户信息
func (m *MiniprogramUserService) decryptUserInfo(encryptedData, iv, sessionKey string) (*request.MiniprogramUserInfo, error) {
	// Base64解码
	aesKey, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, err
	}
	
	cipherText, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}
	
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}
	
	// AES-128-CBC解密
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}
	
	if len(cipherText) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	
	mode := cipher.NewCBCDecrypter(block, ivBytes)
	mode.CryptBlocks(cipherText, cipherText)
	
	// 去除PKCS7填充
	cipherText = pkcs7UnPadding(cipherText)
	
	// 解析JSON
	var userInfo request.MiniprogramUserInfo
	err = json.Unmarshal(cipherText, &userInfo)
	if err != nil {
		return nil, err
	}
	
	return &userInfo, nil
}

// pkcs7UnPadding PKCS7去填充
func pkcs7UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

// AppRegister APP注册
func (m *MiniprogramUserService) AppRegister(registerReq request.AppRegister) (userInter system.Login, token string, expiresAt int64, err error) {
	// 1. 检查用户名是否已存在
	var existUser system.MiniprogramUser
	err = global.GVA_DB.Where("username = ?", registerReq.Username).First(&existUser).Error
	if err == nil {
		return nil, "", 0, errors.New("用户名已存在")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, "", 0, err
	}

	// 2. 检查邮箱是否已存在
	err = global.GVA_DB.Where("email = ?", registerReq.Email).First(&existUser).Error
	if err == nil {
		return nil, "", 0, errors.New("邮箱已被注册")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, "", 0, err
	}

	// 3. 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", 0, errors.New("密码加密失败")
	}

	// 4. 创建用户
	now := time.Now()
	user := system.MiniprogramUser{
		UUID:      uuid.New(),
		Username:  registerReq.Username,
		Password:  string(hashedPassword),
		Email:     registerReq.Email,
		NickName:  registerReq.NickName,
		Phone:     registerReq.Phone,
		LoginType: 1, // 账号密码登录
		Enable:    1,
		LastLogin: &now,
	}

	err = global.GVA_DB.Create(&user).Error
	if err != nil {
		return nil, "", 0, err
	}

	// 5. 生成JWT token
	j := utils.NewJWT()
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		Username:    user.Username,
		NickName:    user.NickName,
		AuthorityId: user.GetAuthorityId(),
	})
	token, err = j.CreateToken(claims)
	if err != nil {
		return nil, "", 0, err
	}

	// 6. 设置Redis缓存（如果启用）
	if global.GVA_CONFIG.System.UseMultipoint {
		err = utils.SetRedisJWT(token, user.Username)
		if err != nil {
			return nil, "", 0, err
		}
	}

	return &user, token, claims.ExpiresAt.Unix(), nil
}

// AppLogin APP登录
func (m *MiniprogramUserService) AppLogin(loginReq request.AppLogin) (userInter system.Login, token string, expiresAt int64, err error) {
	// 1. 验证验证码
	if !captcha.NewDefaultRedisStore().Verify(loginReq.CaptchaId, loginReq.Captcha, true) {
		return nil, "", 0, errors.New("验证码错误")
	}

	// 2. 查找用户
	var user system.MiniprogramUser
	err = global.GVA_DB.Where("username = ? OR email = ?", loginReq.Username, loginReq.Username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", 0, errors.New("用户不存在")
		}
		return nil, "", 0, err
	}

	// 3. 检查用户状态
	if user.Enable != 1 {
		return nil, "", 0, errors.New("用户已被禁用")
	}

	// 4. 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		return nil, "", 0, errors.New("密码错误")
	}

	// 5. 更新最后登录时间
	now := time.Now()
	user.LastLogin = &now
	err = global.GVA_DB.Save(&user).Error
	if err != nil {
		return nil, "", 0, err
	}

	// 6. 生成JWT token
	j := utils.NewJWT()
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		Username:    user.Username,
		NickName:    user.NickName,
		AuthorityId: user.GetAuthorityId(),
	})
	token, err = j.CreateToken(claims)
	if err != nil {
		return nil, "", 0, err
	}

	// 7. 设置Redis缓存（如果启用）
	if global.GVA_CONFIG.System.UseMultipoint {
		err = utils.SetRedisJWT(token, user.Username)
		if err != nil {
			return nil, "", 0, err
		}
	}

	return &user, token, claims.ExpiresAt.Unix(), nil
}

// UpdateProfile 更新用户资料（通用）
func (m *MiniprogramUserService) UpdateProfile(uuid string, updateReq request.UpdateProfile) error {
	updates := make(map[string]interface{})
	
	updates["nick_name"] = updateReq.NickName
	updates["avatar"] = updateReq.Avatar
	updates["gender"] = updateReq.Gender
	updates["phone"] = updateReq.Phone
	updates["city"] = updateReq.City
	updates["province"] = updateReq.Province
	
	return global.GVA_DB.Model(&system.MiniprogramUser{}).Where("uuid = ?", uuid).Updates(updates).Error
}

// ChangePassword 修改密码
func (m *MiniprogramUserService) ChangePassword(uuid string, changeReq request.ChangePassword) error {
	// 1. 查找用户
	var user system.MiniprogramUser
	err := global.GVA_DB.Where("uuid = ?", uuid).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("用户不存在")
		}
		return err
	}

	// 2. 验证原密码
	if user.Password == "" {
		return errors.New("该用户未设置密码，无法修改密码")
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(changeReq.OldPassword))
	if err != nil {
		return errors.New("原密码错误")
	}

	// 3. 加密新密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(changeReq.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("新密码加密失败")
	}

	// 4. 更新密码
	return global.GVA_DB.Model(&user).Update("password", string(hashedPassword)).Error
}