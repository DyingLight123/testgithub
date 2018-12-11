package controllers

import (
	"encoding/json"
	"testgithub/test/models"
)

type BackendHomeController struct {
	BaseController
}

type HomeParam struct {

}

type APIRequestHome struct {
	models.APIRequestCommon
	HomeParam `json:"params"`
}

func (home *BackendHomeController) Home() {
	var request APIRequestHome
	json.Unmarshal(home.Ctx.Input.RequestBody, &request)
	if models.GetUserOneName(request.User) == false {
		returnURL := home.Ctx.Request.URL.Path
		home.Data["json"] = map[string]interface{} {"id": "", "code": 0, "desc": "",
			"data": map[string]interface{} {"message": "请先登录！", "returnURL": returnURL}}
	} else {
		home.Data["json"] = map[string]interface{} {"id": "", "code": 0, "desc": "",
			"data": map[string]interface{} {"username": request.User}}
	}
	home.ServeJSON()
}
