package controllers

import (
	"encoding/json"
	"log"
	"testgithub/test/models"
)

type BackendStoryController struct {
	BaseController
}

type BankendStoryParam struct {
	Page int `json:"page"`
}

type APIRequestBankendStory struct {
	models.APIRequestCommon
	BankendStoryParam `json:"params"`
}

func (backendstory *BackendStoryController) Post() {
	var request APIRequestBankendStory
	json.Unmarshal(backendstory.Ctx.Input.RequestBody, &request)
	if models.GetUserOneName(request.User) == false {
		backendstory.Data["json"] = map[string]interface{}{"id": request.Id, "code": 0, "desc": "",
			"data": map[string]interface{}{"message": "找不到该用户名对应的id！"}}
	} else {
		userid, err := models.GetUserIdByName(request.User)
		if err != nil {
			log.Fatal(err)
			backendstory.Data["json"] = map[string]interface{}{"id": request.Id, "code": 1, "desc": err,
				"data": map[string]interface{}{"message": "通过username获取userid出错！"}}
		} else {
			story, err1 := models.GetStoryAllByUserId(request.Page, userid)
			if err1 != nil {
				log.Fatal(err1)
				backendstory.Data["json"] = map[string]interface{}{"id": request.Id, "code": 1, "desc": err1,
					"data": map[string]interface{}{"message": "通过userid获取story列表出错！"}}
			} else {
				storydata := make([]map[string]interface{}, 0)
				for i := 0; i < len(story); i++ {
					m := make(map[string]interface{})
					j, _ := json.Marshal(&story[i])
					json.Unmarshal(j, &m)
					storydata = append(storydata, m)
				}
				backendstory.Data["json"] = map[string]interface{}{"id": request.Id, "code": 0, "desc": "",
					"data": storydata}
			}
		}
	}
	backendstory.ServeJSON()
}
