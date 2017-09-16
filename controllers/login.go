package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"net/http"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.tpl"
}

func (c *LoginController) Post() {
	c.Ctx.Request.ParseForm()
	username := c.Ctx.Request.FormValue("username")
	password := c.Ctx.Request.FormValue("password")
	logs.Info("Username: ", username)
	logs.Info("Password: ", password)
	c.SetSession(sessionName, username+" "+password)
	c.Redirect("/admin", http.StatusFound)
}
