package middleware

import (
	"errors"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// MiniprogramJWTAuth 小程序JWT鉴权中间件
// 这个中间件专门用于小程序用户的鉴权，不影响原有的admin鉴权逻辑
func MiniprogramJWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token，小程序可以从header或者query parameter获取
		token := utils.GetToken(c)
		if token == "" {
			response.NoAuth("小程序用户未登录或非法访问，请登录", c)
			c.Abort()
			return
		}

		// 检查token是否在黑名单
		if isBlacklist(token) {
			response.NoAuth("您的帐户异地登陆或令牌失效", c)
			utils.ClearToken(c)
			c.Abort()
			return
		}

		j := utils.NewJWT()
		// 解析token
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.TokenExpired) {
				response.NoAuth("登录已过期，请重新登录", c)
				utils.ClearToken(c)
				c.Abort()
				return
			}
			response.NoAuth(err.Error(), c)
			utils.ClearToken(c)
			c.Abort()
			return
		}

		// 验证是否是小程序用户（通过AuthorityId判断，小程序用户的AuthorityId为999）
		if claims.AuthorityId != 999 {
			response.NoAuth("非小程序用户，无权访问", c)
			c.Abort()
			return
		}

		// 可选：验证小程序用户是否被禁用
		// 这里可以根据需要添加用户状态检查逻辑
		// if user, err := miniprogramUserService.GetUserInfo(claims.UUID.String()); err != nil || user.Enable == 2 {
		//     response.NoAuth("用户已被禁用", c)
		//     c.Abort()
		//     return
		// }

		c.Set("claims", claims)
		
		// token刷新逻辑
		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			dr, _ := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
			utils.SetToken(c, newToken, int(dr.Seconds()))
			if global.GVA_CONFIG.System.UseMultipoint {
				// 记录新的活跃jwt（使用openid作为key）
				_ = utils.SetRedisJWT(newToken, newClaims.Username)
			}
		}
		c.Next()

		// 处理响应头中的新token
		if newToken, exists := c.Get("new-token"); exists {
			c.Header("new-token", newToken.(string))
		}
		if newExpiresAt, exists := c.Get("new-expires-at"); exists {
			c.Header("new-expires-at", newExpiresAt.(string))
		}
	}
}

// MiniprogramOptionalAuth 小程序可选鉴权中间件
// 如果有token则验证，没有token也可以通过（用于一些可选登录的接口）
func MiniprogramOptionalAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := utils.GetToken(c)
		if token == "" {
			// 没有token，直接继续
			c.Next()
			return
		}

		// 有token则验证
		if isBlacklist(token) {
			c.Next()
			return
		}

		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			// token无效，但不阻止请求
			c.Next()
			return
		}

		// 验证是否是小程序用户
		if claims.AuthorityId == 999 {
			c.Set("claims", claims)
		}

		c.Next()
	}
}