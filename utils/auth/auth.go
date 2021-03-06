package auth

import (
	//	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/juusechec/jwt-beego"

	"strings"
)

func GetClaims(ctx *context.Context) string {
	//	tokenString := c.Ctx.Input.Query("tokenString")
	auth := ctx.Input.Header("Authorization")
	if len(auth) > 0 {
		header := strings.Split(auth, " ")
		if len(header) != 2 || header[0] != "Bearer" {
			return ""
		}

		et := jwtbeego.EasyToken{}
		valid, claims, _ := et.ValidateToken(header[1])
		if !valid {
			return ""
		}
		return claims
	}
	return ""
}

func GetUsername(ctx *context.Context) (username string) {
	claims := GetClaims(ctx)
	return claims

}
