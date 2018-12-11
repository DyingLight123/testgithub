package controllers

import (
	"encoding/json"
	"log"
	"testgithub/test/models"
)

type BackendDeleteAnimalController struct {
	BaseController
}

type BackendDeleteAnimalParam struct {
	Id 			int 		`json:"id"`
}

type APIRequestBackendDeleteAnimal struct {
	models.APIRequestCommon
	BackendDeleteAnimalParam `json:"params"`
}

func (backenddel * BackendDeleteAnimalController) Post() {
	var request APIRequestBackendDeleteAnimal
	json.Unmarshal(backenddel.Ctx.Input.RequestBody, &request)
	userid1, err := models.GetAnimalOneUserIdById(request.BackendDeleteAnimalParam.Id)
	if err != nil {
		log.Fatal(err)
	} else {
		userid2, err := models.GetUserIdByName(request.User)
		if err != nil {
			log.Fatal(err)
		} else {
			//fmt.Println(userid1, userid2)
			if userid1 != userid2 {
				backenddel.Data["json"] = map[string]interface{} {"id": request.APIRequestCommon.Id, "code": 0, "desc": "",
					"data": map[string]interface{}{"message": "您无权删除该文章！"}}
			} else {
				err := models.DeleteAnimalLikeAll(request.BackendDeleteAnimalParam.Id)
				if err != nil {
					log.Fatal(err)
				}
				err = models.DeleteAnimalCommonAll(request.BackendDeleteAnimalParam.Id)
				if err != nil {
					log.Fatal(err)
				}
				err = models.DeleteAnimal(request.BackendDeleteAnimalParam.Id)
				if err != nil {
					log.Fatal(err)
				}
				backenddel.Data["json"] = map[string]interface{} {"id": "", "code": 0, "desc": "",
					"data": map[string]interface{}{"message": "删除成功! "}}
			}

		}
	}
	backenddel.ServeJSON()
}

