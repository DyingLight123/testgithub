package controllers

import (
	"encoding/json"
	"log"
	"testgithub/test/models"
)

type StoryController struct {
	BaseController
}

type StoryParam struct {
	Page int `json:"page"`
	Type int `json:"type"`
}

type APIRequestStory struct {
	models.APIRequestCommon
	StoryParam `json:"params"`
}

func (story *StoryController) Get() {
	page, err := story.GetInt("page")
	if err != nil {
		page = 1
	}
	orderid, err := story.GetInt("orderid")
	if err != nil {
		orderid = 2
	}
	liststory, err := models.GetListStoryRecent(page, orderid)
	if err != nil {
		log.Fatal(err)
	} else if liststory == nil {
		data := make(map[string]interface{})
		data["message1"] = "没有storoy数据！"
		story.Data["json"] = map[string]interface{}{"id": "", "code": 0, "desc": "",
			"data": data}
	} else {
		liststorydata := make([]map[string]interface{}, 0)
		for i := 0; i < len(liststory); i++ {
			m := make(map[string]interface{})
			j, _ := json.Marshal(&liststory[i])
			json.Unmarshal(j, &m)
			liststorydata = append(liststorydata, m)
		}
		story.Data["json"] = map[string]interface{}{"id": "", "code": 0, "desc": "",
			"data": liststorydata}
	}
	story.ServeJSON()
}