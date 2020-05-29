package middleware

import (
	"net/http"
	"time"

	jwt "jwt-demo/internal/middleware/jwtauth"

	"jwt-demo/internal/user"

	"github.com/gin-gonic/gin"
)

// AuthInit ..
func AuthInit() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "jwt auth",
		Key:             []byte("secret key"),
		Timeout:         time.Duration(1) * time.Hour,
		MaxRefresh:      time.Duration(1) * time.Hour,
		IdentityKey:     "secret-key", // you should change to more secureity char
		PayloadFunc:     PayloadFunc,
		IdentityHandler: IdentityHandler,
		Authenticator:   Authenticator,
		Unauthorized:    Unauthorized,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})
}

// PayloadFunc ..
func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(*user.Model)
		return jwt.MapClaims{
			jwt.IdentityKey: u.ID,
			jwt.UserName:    u.UserName,
			jwt.Email:       u.Email,
		}
	}
	return jwt.MapClaims{}
}

// IdentityHandler ..
func IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return map[string]interface{}{
		"IdentityKey": claims["identity"],
		"UserName":    claims["username"],
		"Email":       claims["email"],
	}
}

// Authenticator ..
func Authenticator(c *gin.Context) (interface{}, error) {
	var loginReq user.Login

	if err := c.ShouldBind(&loginReq); err != nil {
		return nil, jwt.ErrMissingLoginValues
	}

	userSvc := user.NewService()
	userModel, err := userSvc.Authenticate(&loginReq)
	if err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	return map[string]interface{}{"user": userModel}, nil
}

// Unauthorized ..
func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  message,
	})
}
