package http

import (
	"github.com/astaxie/beego"
	"github.com/gaochao1/fe/g"
	"github.com/gaochao1/fe/http/home"
	"github.com/gaochao1/fe/http/uic"
	uic_model "github.com/gaochao1/fe/model/uic"
)

func Start() {
	if !g.Config().Http.Enabled {
		return
	}

	addr := g.Config().Http.Listen
	if addr == "" {
		return
	}

	home.ConfigRoutes()
	uic.ConfigRoutes()

	beego.AddFuncMap("member", uic_model.MembersByTeamId)
	beego.Run(addr)
}
