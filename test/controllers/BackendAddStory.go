package controllers

import (
	"encoding/json"
	"log"
	"testgithub/test/models"
)

type BackendAddStoryController struct {
	BaseController
}

type BankendAddStoryParam struct {
	Title    string `json:"title"`
	Describe string `json:"describe"`
}
type APIRequestBankendAddStory struct {
	models.APIRequestCommon
	BankendAddStoryParam `json:"params"`
}

func (backendstory *BackendAddStoryController) AddStory() {
	var request APIRequestBankendAddStory
	json.Unmarshal(backendstory.Ctx.Input.RequestBody, &request)
	if models.GetUserOneName(request.User) == false {
		backendstory.Data["json"] = map[string]interface{} {"id": request.Id, "code": 0, "desc": "",
			"data": map[string]interface{}{"message": "找不到该用户名对应的id！"}}
	} else {
		userid, err := models.GetUserIdByName(request.User)
		if err != nil {
			log.Fatal(err)
			backendstory.Data["json"] = map[string]interface{} {"id": request.Id, "code": 1, "desc": err,
				"data": map[string]interface{}{"message": "获取用户名出错！"}}
		} else {
			story := new(models.Story)
			story.Userid = userid
			story.Title = request.Title
			story.Describe = request.Describe
			id, err1 := models.AddStory(story)
			if err1 != nil {
				log.Fatal(err1)
				backendstory.Data["json"] = map[string]interface{} {"id": request.Id, "code": 1, "desc": err1,
					"data": map[string]interface{}{"message": "插入story表出错！"}}
			} else {
				backendstory.Data["json"] = map[string]interface{} {"id": request.Id, "code": 0, "desc": "",
					"data": map[string]interface{}{"message": "插入成功！", "id": id}}
			}
		}
	}
	backendstory.ServeJSON()
}
