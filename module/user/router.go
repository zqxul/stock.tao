package user

import "stock.tao/module/core"

var group = core.Router.Group("/user")

// init request path to handler
func init() {
	group.POST("/register", Register)
	group.POST("/login", Login)
}
