// Package controller loads the routes for each of the controllers.
package controller

import (
	"github.com/TheTwoOfUs/gotest/controller/about"
	"github.com/TheTwoOfUs/gotest/controller/debug"
	"github.com/TheTwoOfUs/gotest/controller/home"
	"github.com/TheTwoOfUs/gotest/controller/login"
	"github.com/TheTwoOfUs/gotest/controller/notepad"
	"github.com/TheTwoOfUs/gotest/controller/register"
	"github.com/TheTwoOfUs/gotest/controller/static"
	"github.com/TheTwoOfUs/gotest/controller/status"
)

// LoadRoutes loads the routes for each of the controllers.
func LoadRoutes() {
	about.Load()
	debug.Load()
	register.Load()
	login.Load()
	home.Load()
	static.Load()
	status.Load()
	notepad.Load()
}
