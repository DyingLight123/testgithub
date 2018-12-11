package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"testgithub/test/models"
)

type StoryCommonLikeController struct {
	BaseController
}

type StoryCommonLikeParam struct {
	StoryCommonId int `json:"storycommonid"`
}

type APIRequestStoryCommonLike struct {
	models.APIRequestCommon
	StoryCommonLikeParam `json:"params"`
}

func (like *StoryCommonLikeController) Post() {
	storyid, err := like.GetInt("storyid")
	if err != nil {
		like.Redirect("/home/story", 302)
		return
	}
	page, err := like.GetInt("page")
	if err != nil {
		page = 1
	}
	orderid, err := like.GetInt("orderid")
	if err != nil {
		orderid = 1
	}
	var request APIRequestStoryCommonLike
	json.Unmarshal(like.Ctx.Input.RequestBody, &request)
	if models.GetUserOneName(request.User) == false {
		returnURL := "/home/story/common"
		returnURL = returnURL + fmt.Sprintf("?storyid=%d&page=%d&orderid=%d", storyid, page, orderid)
		like.Data["json"] = map[string]interface{}{"id": request.Id, "code": 0, "desc": "",
			"data": map[string]interface{}{"message": "请先登录！", "returnURL": returnURL}}
	} else {
		userid, err := models.GetUserIdByName(request.User)
		if err != nil {
			log.Fatal(err)
		}
		if models.IsEmptyStoryCommonLike(userid, request.StoryCommonId) == true {
			id, err := models.AddStoryCommonLike(userid, request.StoryCommonId)
			if err != nil {
				log.Fatal(err)
			}
			err = models.UpdateStoryLikeCountById(request.StoryCommonId)
			if err != nil {
				log.Fatal(err)
			}
			like.Data["json"] = map[string]interface{}{"id": request.Id, "code": 0, "desc": "",
				"data": map[string]interface{}{"message": "点赞成功！", "id": id}}
		} else {
			like.Data["json"] = map[string]interface{}{"id": request.Id, "code": 0, "desc": "",
				"data": map[string]interface{}{"message": "你已经点过了！"}}
		}

	}
	strURL := "/home/story/common"
	strURL = strURL + fmt.Sprintf("?storyid=%d&page=%d&orderid=%d", storyid, page, orderid)
	like.Redirect(strURL, 302)
	like.ServeJSON()
}
