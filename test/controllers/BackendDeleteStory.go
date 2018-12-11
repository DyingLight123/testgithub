package controllers

import (
	"encoding/json"
	"log"
	"testgithub/test/models"
)

type BackendDeleteStoryController struct {
	BaseController
}

type BackendDeleteStoryParam struct {
	Id 			int 		`json:"id"`
}

type APIRequestBackendDeleteStory struct {
	models.APIRequestCommon
	BackendDeleteStoryParam `json:"params"`
}

func (backendstory *BackendDeleteStoryController) Post() {
	var request APIRequestBackendDeleteStory
	json.Unmarshal(backendstory.Ctx.Input.RequestBody, &request)
	userid1, err1 := models.GetStoryOneUserIdById(request.BackendDeleteStoryParam.Id)
	if err1 != nil {
		log.Fatal(err1)
	} else {
		userid2, err2 := models.GetUserIdByName(request.User)
		if err2 != nil {
			log.Fatal(err2)
		} else {
			//fmt.Println(userid1, userid2)
			if userid1 != userid2 {
				backendstory.Data["json"] = map[string]interface{} {"id": request.APIRequestCommon.Id, "code": 0, "desc": "",
					"data": map[string]interface{}{"message": "您无权删除该文章！"}}
			} else {
				err := models.DeleteStoryLikeAll(request.BackendDeleteStoryParam.Id)
				if err != nil {
					log.Fatal(err)
				}
				err = models.DeleteStoryCommonAll(request.BackendDeleteStoryParam.Id)
				if err != nil {
					log.Fatal(err)
				}
				err = models.DeleteStory(request.BackendDeleteStoryParam.Id)
				if err != nil {
					log.Fatal(err)
				}
				backendstory.Data["json"] = map[string]interface{} {"id": request.APIRequestCommon.Id, "code": 0, "desc": "",
					"data": map[string]interface{}{"message": "删除成功! "}}
			}

		}
	}
	backendstory.ServeJSON()
}
