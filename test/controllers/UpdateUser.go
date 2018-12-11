package controllers

import (
	"encoding/json"
	"log"
	"testgithub/test/models"
)

type UpdateUserController struct {
	BaseController
}

type UPdateUserParam struct {
	UserName           	string 	`json:"user_name"`
	PassWord           	string 	`json:"password"`
}
type APIRequestUpdateUser struct {
	models.APIRequestCommon
	UPdateUserParam `json:"params"`
}

func (updateuser *UpdateUserController) Get() {
	username := updateuser.GetString("username")
	user, err := models.GetUserOne(username)
	if err != nil {
		log.Fatal("数据库查询出错：", err)
		updateuser.Data["json"] = map[string]interface{} {"id": "", "code": 1, "desc": err,
			"data": map[string]interface{} {"message": "用户详细信息查询失败！"}}
	}
	updateuser.Data["json"] = map[string]interface{} {"id": "", "code": 0, "desc": "",
		"data": map[string]interface{} {"name": user.Name, "password": user.Pwd}}
	updateuser.ServeJSON()
}

func (updateuser *UpdateUserController) Post() {
	var request APIRequestUpdateUser
	json.Unmarshal(updateuser.Ctx.Input.RequestBody, &request)
	if models.GetUserOneName(request.UserName) == true {
		updateuser.Data["json"] = map[string]interface{} {"id": request.Id, "code": 0, "desc": "",
			"data": map[string]interface{} {"message": "用户名被占用"}}
	} else {
		err := models.UpdateUser(request.User, request.UserName, request.PassWord)
		if err != nil {
			updateuser.Data["json"] = map[string]interface{} {"id": request.Id, "code": 1, "desc": err,
				"data": map[string]interface{} {"message": "用户信息更新失败！"}}
		}
		updateuser.Data["json"] = map[string]interface{} {"id": request.Id, "code": 0, "desc": "",
			"data": map[string]interface{} {"message": "用户信息更新成功！"}}
	}
	updateuser.ServeJSON()
}
