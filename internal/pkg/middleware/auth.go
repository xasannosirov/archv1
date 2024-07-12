package middleware

import (
	"archv1/internal/pkg/errors"
	"log"
	"net/http"

	"archv1/internal/pkg/config"
	"archv1/internal/pkg/tokens"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/cast"
)

// JWTRoleAuth ...
type JWTRoleAuth struct {
	enforcer   *casbin.Enforcer
	cfg        config.Config
	jwtHandler tokens.JWTHandler
}

// NewAuthorizer ...
func NewAuthorizer(e *casbin.Enforcer, jwtHandler tokens.JWTHandler, cfg config.Config) gin.HandlerFunc {
	a := &JWTRoleAuth{
		enforcer:   e,
		cfg:        cfg,
		jwtHandler: jwtHandler,
	}

	return func(c *gin.Context) {
		allow, err := a.CheckPermission(c.Request)
		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			if v.Errors == jwt.ValidationErrorExpired {
				a.RequireRefresh(c)
			} else {
				a.RequirePermission(c)
			}
		} else if !allow {
			a.RequirePermission(c)
		}
	}
}

// CheckPermission check permissions as role, path and method
func (a *JWTRoleAuth) CheckPermission(r *http.Request) (bool, error) {
	role, err := a.GetRole(r)
	if err != nil {
		return false, err
	}

	method := r.Method
	path := r.URL.Path

	allowed, err := a.enforcer.Enforce(role, path, method)
	if err != nil {
		log.Println("failed to check permission: ", err)
		return false, err
	}

	return allowed, nil
}

// GetRole gets role from http request
func (a *JWTRoleAuth) GetRole(r *http.Request) (string, error) {
	var (
		role   string
		claims jwt.MapClaims
		err    error
	)

	jwtToken := r.Header.Get("Authorization")
	if jwtToken == "" {
		return "unauthorized", nil
	}

	a.jwtHandler.Token = jwtToken
	claims, err = a.jwtHandler.ExtractClaims()
	if err != nil {
		return "", err
	}

	if cast.ToString(claims["role"]) == "sudo" {
		role = "sudo"
	} else if cast.ToString(claims["role"]) == "admin" {
		role = "admin"
	} else if cast.ToString(claims["role"]) == "user" {
		role = "user"
	} else {
		role = "unknown"
	}

	return role, nil
}

// RequireRefresh response with 401
func (a *JWTRoleAuth) RequireRefresh(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, errors.Error{
		Message: "Require refresh",
	})
	c.AbortWithStatus(401)
}

// RequirePermission response with 403
func (a *JWTRoleAuth) RequirePermission(c *gin.Context) {
	c.JSON(http.StatusForbidden, errors.Error{
		Message: "You have no access this page",
	})
	c.AbortWithStatus(403)
}
