// declare api routes
package api

import (
	"github.com/casbin/casbin/v2"
	"github.com/ezzddinne/api/app"
	"github.com/ezzddinne/api/squad"
	"github.com/ezzddinne/api/user"
	"github.com/ezzddinne/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutesApis(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	// auth routes
	user.RoutesAuth(router.Group("/user"), db, enforcer)

	// reset password routes
	user.RoutesUserPassword(router.Group("/user/reset"), db, enforcer)

	// user route
	user.RoutesUsersJWT(router.Group("/user/jwt", middleware.AuthorizeJWT()), db, enforcer)

	// paiment status route
	user.RouteAdmin(router.Group("/user/paiment", middleware.AuthorizeJWT()), db, enforcer)

	// auth jwt routes
	squad.RoutesAuthJWT(router.Group("/auth/jwt", middleware.AuthorizeJWT()), db, enforcer)

	// app routes
	app.RoutesApps(router.Group("/app", middleware.AuthorizeJWT()), db, enforcer)

}
