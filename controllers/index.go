package controllers

import (
	"github.com/astaxie/beego"
	"time"
)

type IndexController struct {
	beego.Controller
}

type Blog struct {
	Link  string
	Title string
	Time  time.Time
}

func (c *IndexController) Get() {
	c.TplName = "index.html"

	blogs := []Blog{
		{"/blog/a", "Hello a", time.Now()},
		{"/blog/b", "Hello b", time.Now()},
	}

	c.Data["Blogs"] = blogs
}
