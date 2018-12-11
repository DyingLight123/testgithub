package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	//username 	string
}

func (base *BaseController) CheckLogin() {
	base.SetSession("username", "admin1")
	a := base.GetSession("username")
	fmt.Println(base.CruSession)
	if a == nil {
		urlstr := base.URLFor("LoginController.GetLogin") + "?url="
		returnURL := base.Ctx.Request.URL.Path
		base.Redirect(urlstr + returnURL, 302)
	}
}
