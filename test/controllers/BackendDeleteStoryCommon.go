package controllers

import (
	"encoding/json"
	"log"
	"testgithub/test/models"
)

type BackendDeleteStoryCommonController struct {
	BaseController
}

type BackendDeleteStoryCommonParam struct {
	Storycommonid	int	`json:"storycommonid"`
}

type APIRequestBackendDeleteStoryCommon struct {
	models.APIRequestCommon
	BackendDeleteStoryCommonParam `json:"params"`
}

func (delete BackendDeleteStoryCommonController) Post() {
	var request APIRequestBackendDeleteStoryCommon
	json.Unmarshal(delete.Ctx.Input.RequestBody, &request)
	if models.GetUserOneName(request.User) == false {
		delete.Data["json"] = map[string]interface{} {"id": request.Id, "code": 0, "desc": "",
			"data": map[string]interface{}{"message": "找不到该用户名对应的id！"}}
	} else {
		userid, err := models.GetUserIdByName(request.User)
		if err != nil {
			log.Fatal(err)
		}
		userid2, err := models.GetStoryCommonUserIdById(request.Storycommonid)
		if err != nil {
			log.Fatal(err)
		}
		if userid2 != userid {
			delete.Data["json"] = map[string]interface{} {"id": request.Id, "code": 0, "desc": "",
				"data": map[string]interface{}{"message": "您无权删除该评论！"}}
		} else {
			err := models.DeleteStoryLikeByStoryCommonId(request.Storycommonid)
			if err != nil {
				log.Fatal(err)
			}
			err = models.DeleteStoryCommonOneById(request.Storycommonid)
			if err != nil {
				log.Fatal(err)
			}
			delete.Data["json"] = map[string]interface{} {"id": request.Id, "code": 0, "desc": "",
				"data": map[string]interface{}{"message": "删除成功！"}}
		}
	}
	delete.ServeJSON()
}
