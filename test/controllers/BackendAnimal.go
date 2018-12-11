package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"testgithub/test/models"
)

type BackendAnimalConntroller struct {
	BaseController
}

type BackendAnimalParam struct {
	Page 	int 	`json:"page"`
}

type APIRequestBackendAnimal struct {
	models.APIRequestCommon
	BackendAnimalParam `json:"params"`
}

func (backendanimal *BackendAnimalConntroller) Post() {
	var request APIRequestBackendAnimal
	json.Unmarshal(backendanimal.Ctx.Input.RequestBody, &request)
	if models.GetUserOneName(request.User) == false {
		backendanimal.Data["json"] = map[string]interface{} {"id": request.Id, "code": 0, "desc": "",
			"data": map[string]interface{}{"message": "找不到该用户名对应的id！"}}
	} else {
		userid, err := models.GetUserIdByName(request.User)
		if err != nil {
			log.Fatal(err)
			backendanimal.Data["json"] = map[string]interface{} {"id": request.Id, "code": 1, "desc": err,
				"data": map[string]interface{}{"message": "获取用户id失败！"}}
		}
		animal, err := models.GetAnimalAllByUserId(request.Page, userid)
		if err != nil {
			log.Fatal(err)
			backendanimal.Data["json"] = map[string]interface{} {"id": request.Id, "code": 1, "desc": err,
				"data": map[string]interface{}{"message": "获取动物列表信息失败！"}}
		}
		//fmt.Println(animal)
		animaldata := make([]map[string]interface{}, 0)
		fmt.Println(len(animal))
		for i := 0; i < len(animal); i++ {
			m := make(map[string]interface{})
			j, _ := json.Marshal(&animal[i])
			json.Unmarshal(j, &m)
			animaldata = append(animaldata, m)
		}
		backendanimal.Data["json"] = map[string]interface{} {"id": request.Id, "code": 0, "desc": "",
			"data": animaldata}
	}
	backendanimal.ServeJSON()
}
