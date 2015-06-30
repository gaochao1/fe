package home

import (
	"github.com/astaxie/beego"
	"github.com/gaochao1/fe/g"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["Shortcut"] = g.Config().Shortcut
	this.TplNames = "home/index.html"
}
