package controllers

import (
	"encoding/json"
	"log"
	"testgithub/test/models"
)

type BackendUpdateStoryController struct {
	BaseController
}

type BankendUpdateStoryParam struct {
	Id 			int 		`json:"id"`
	Title 		string 		`json:"title"`
	Describe	string		`json:"describe"`
}

type APIRequestBankendUpdateStory struct {
	models.APIRequestCommon
	BankendUpdateStoryParam `json:"params"`
}

func (backendstory *BackendUpdateStoryController) Get() {
	id, err := backendstory.GetInt("id")
	if err != nil {
		log.Fatal(err)
	}
	story, err1 := models.GetStoryOneById(id)
	if err1 != nil{
		log.Fatal(err1)
	} else if story == nil {
		backendstory.Data["json"] = map[string]interface{} {"id": "", "code": 0, "desc": "",
			"data": map[string]interface{}{"message": "没找到您给定id的story"}}
	} else {
		m := make(map[string]interface{})
		j, _ := json.Marshal(&story)
		json.Unmarshal(j, &m)
		backendstory.Data["json"] = map[string]interface{} {"id": "", "code": 0, "desc": "",
			"data": m}
	}
	backendstory.ServeJSON()
}

func (backendstory *BackendUpdateStoryController) Post() {
	var request APIRequestBankendUpdateStory
	json.Unmarshal(backendstory.Ctx.Input.RequestBody, &request)
	//fmt.Println(request)
	err := models.UpdateStory(request.BankendUpdateStoryParam.Id, request.Title, request.Describe)
	if err != nil {
		log.Fatal(err)
		backendstory.Data["json"] = map[string]interface{} {"id": request.APIRequestCommon.Id, "code": 1, "desc": err,
			"data": map[string]interface{} {"message": "更新信息出错！"}}
	}
	backendstory.Data["json"] = map[string]interface{} {"id": request.APIRequestCommon.Id, "code": 0, "desc": "",
		"data": map[string]interface{}{"message": "更新成功！"}}
	backendstory.ServeJSON()
}