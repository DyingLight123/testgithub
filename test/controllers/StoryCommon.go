package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"testgithub/test/models"
)

type StoryCommonController struct {
	BaseController
}

type StoryCommonParam struct {
	Storyid  int    `json:"storyid"`
	Describe string `json:"describe"`
}

type APIRequestStoryCommon struct {
	models.APIRequestCommon
	StoryCommonParam `json:"params"`
}

func (storycommon *StoryCommonController) Get() {
	data := make(map[string]interface{})
	var err error
	storyid, err := storycommon.GetInt("storyid")
	if err != nil {
		storycommon.Redirect("/home/story", 302)
		return
	}
	page, err := storycommon.GetInt("page")
	if err != nil {
		page = 1
	}
	orderid, err := storycommon.GetInt("orderid")
	if err != nil {
		orderid = 1
	}
	story, err := models.GetStoryOneById(storyid)
	if err != nil {
		log.Fatal(err)
	} else if story == nil {
		data["message"] = "给的storyid不存在！"
		storycommon.Data["json"] = map[string]interface{}{"id": "", "code": 1, "desc": "storyid不存在！",
			"data": data}
	} else {
		m := make(map[string]interface{})
		j, _ := json.Marshal(&story)
		json.Unmarshal(j, &m)
		data["story"] = m
		liststorycommon, err := models.GetStoryCommonListByStoryId(page, storyid, orderid)
		if err != nil {
			log.Fatal(err)
		}
		list := make([]map[string]interface{}, 0)
		for i := 0; i < len(liststorycommon); i++ {
			n := make(map[string]interface{})
			j, _ := json.Marshal(&liststorycommon[i])
			json.Unmarshal(j, &n)
			list = append(list, n)
		}
		data["storycommon"] = list
		storycommon.Data["json"] = map[string]interface{}{"id": "", "code": 0, "desc": "",
			"data": data}
	}
	storycommon.ServeJSON()
}

func (storycommon *StoryCommonController) Post() {
	storyid, err := storycommon.GetInt("storyid")
	if err != nil {
		storycommon.Redirect("/home/story", 302)
		return
	}
	page, err := storycommon.GetInt("page")
	if err != nil {
		page = 1
	}
	orderid, err := storycommon.GetInt("orderid")
	if err != nil {
		orderid = 1
	}
	var request APIRequestStoryCommon
	json.Unmarshal(storycommon.Ctx.Input.RequestBody, &request)
	if models.GetUserOneName(request.User) == false {
		returnURL := storycommon.Ctx.Request.URL.Path
		returnURL = returnURL + fmt.Sprintf("?storyid=%d&page=%d&orderid=%d", storyid,
			page, orderid)
		storycommon.Data["json"] = map[string]interface{}{"id": request.Id, "code": 0, "desc": "",
			"data": map[string]interface{}{"message": "请先登录！", "returnURL": returnURL}}
	} else {
		userid, err := models.GetUserIdByName(request.User)
		if err != nil {
			log.Fatal(err)
		}
		common := new(models.StoryCommon)
		common.Userid = userid
		common.Storyid = request.Storyid
		common.Describe = request.Describe
		id, err := models.AddStoryCommon(common)
		if err != nil {
			log.Fatal(err)
		}
		storycommon.Data["json"] = map[string]interface{}{"id": request.Id, "code": 0, "desc": "",
			"data": map[string]interface{}{"message": "添加成功！", "id": id}}
	}
	storycommon.ServeJSON()
}
