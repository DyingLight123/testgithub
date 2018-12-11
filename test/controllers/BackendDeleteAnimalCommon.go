package controllers

import (
	"encoding/json"
	"log"
	"testgithub/test/models"
)

type BackendDeleteAnimalCommonController struct {
	BaseController
}

type BackendDeleteAnimalCommonParam struct {
	Animalcommonid	int	`json:"animalcommonid"`
}

type APIRequestBackendDeleteAnimalCommon struct {
	models.APIRequestCommon
	BackendDeleteAnimalCommonParam `json:"params"`
}

func (delete BackendDeleteAnimalCommonController) Post() {
	var request APIRequestBackendDeleteAnimalCommon
	json.Unmarshal(delete.Ctx.Input.RequestBody, &request)
	if models.GetUserOneName(request.User) == false {
		delete.Data["json"] = map[string]interface{} {"id": request.Id, "code": 0, "desc": "",
			"data": map[string]interface{}{"message": "找不到该用户名对应的id！"}}
	} else {
		userid, err := models.GetUserIdByName(request.User)
		if err != nil {
			log.Fatal(err)
		}
		userid2, err := models.GetAnimalCommonUserIdById(request.Animalcommonid)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(userid, userid2)
		if userid2 != userid {
			delete.Data["json"] = map[string]interface{} {"id": request.Id, "code": 0, "desc": "",
				"data": map[string]interface{}{"message": "您无权删除该评论！"}}
		} else {
			err := models.DeleteAnimalLikeByAnimalCommonId(request.Animalcommonid)
			if err != nil {
				log.Fatal(err)
			}
			err = models.DeleteAnimalCommonOneById(request.Animalcommonid)
			if err != nil {
				log.Fatal(err)
			}
			delete.Data["json"] = map[string]interface{} {"id": request.Id, "code": 0, "desc": "",
				"data": map[string]interface{}{"message": "删除成功！"}}
		}
	}
	delete.ServeJSON()
}
