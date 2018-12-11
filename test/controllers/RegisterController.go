package controllers

import (
	"encoding/json"
	"log"
	"testgithub/test/models"
)

type RegisterController struct {
	BaseController
}

type APIRequestRegister struct {
	models.APIResponseCommon
	LoginParam `json:"params"`
}

func (register *RegisterController) Get() {
	register.Data["json"] = map[string]interface{} {"message": "yes"}
	register.ServeJSON()
}

func (register *RegisterController) Post() {
	var requestRegister APIRequestRegister
	json.Unmarshal(register.Ctx.Input.RequestBody, &requestRegister)
	if models.GetUserOneName(requestRegister.UserName) == true {
		register.Data["json"] = map[string]interface{} {"id": requestRegister.Id, "code": 0, "desc": "",
			"data": map[string]interface{} {"message": "用户名被占用！"}}
	} else {
		user := new(models.User)
		user.Name = requestRegister.UserName
		user.Pwd = requestRegister.PassWord
		id, err := models.AddUser(user)
		if err != nil {
			log.Fatal("注册用户失败：", err)
			register.Data["json"] = map[string]interface{} {"id": requestRegister.Id, "code": 1, "desc": err,
				"data": map[string]interface{} {"message": "注册用户失败！"}}
		}
		register.Data["json"] = map[string]interface{} {"id": requestRegister.Id, "code": 0, "desc": "",
			"data": map[string]interface{} {"id": id, "message": "注册成功！"}}
		if requestRegister.SuccessRedirectURL != "" {
			register.Redirect(requestRegister.SuccessRedirectURL, 302)
		}
	}
	register.ServeJSON()
}
