package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"testgithub/test/models"
)

type LoginController struct {
	BaseController
}

type LoginParam struct {
	UserName           	string 	`json:"user_name"`
	PassWord           	string 	`json:"password"`
	SuccessRedirectURL 	string 	`json:"success_redirect_url"`
	//CaptchaID 			string	`json:"captcha_id"`
	//CaptchaVal 			string	`json:"captcha_val"`
}
type APIRequestLogin struct {
	models.APIRequestCommon
	LoginParam `json:"params"`
}

func (login LoginController) GetLogin() {
	returnURL := login.GetString("url")
	login.Data["json"] = map[string]interface{} {"id": "",
		"code": 0, "desc": "", "data": map[string]interface{}{"returnURL": returnURL}}
	login.ServeJSON()
	/*captchaid := captcha.New()
	login.Data["json"] = map[string]interface{} {"id": "",
		"code": 0, "desc": "", "data": map[string]interface{} {"captchaId": captchaid}}
	login.ServeJSON()*/
}

func (login LoginController) PostLogin() {
	var requestlogin APIRequestLogin
	json.Unmarshal(login.Ctx.Input.RequestBody, &requestlogin)
	if models.GetUserOneName(requestlogin.UserName) == false {
		login.Data["json"] = map[string]interface{} {"id": requestlogin.Id, "code": 0, "desc": "",
			"data": map[string]interface{} {"message": "用户名不存在！"}}
	} else {
		password, err := models.GetUserOneByNama(requestlogin.UserName)
		if err != nil {
			log.Fatal("用户密码查询失败：", err)
			login.Data["json"] = map[string]interface{} {"id": requestlogin.Id, "code": 1, "desc": err,
				"data": map[string]interface{} {"message": "密码查询失败！"}}
		} else {
			if password == requestlogin.PassWord {
				login.SetSession("username", requestlogin.UserName)
				fmt.Println(login.CruSession)
				login.Data["json"] = map[string]interface{} {"id": requestlogin.Id, "code": 0, "desc": "",
					"data": map[string]interface{} {"message": "登录成功！"}}
				login.Redirect(requestlogin.SuccessRedirectURL, 302)
			} else {
				login.Data["json"] = map[string]interface{} {"id": requestlogin.Id, "code": 0, "desc": "",
					"data": map[string]interface{} {"message": "密码错误！"}}
			}
		}
	}
	login.ServeJSON()
}
