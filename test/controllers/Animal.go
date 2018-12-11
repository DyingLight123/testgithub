package controllers

import (
	"encoding/json"
	"log"
	"testgithub/test/models"
)

type AnimalController struct {
	BaseController
}

type AnimalParam struct {
	Page int `json:"page"`
	Type int `json:"type"`
}

type APIRequestAnimal struct {
	models.APIRequestCommon
	AnimalParam `json:"params"`
}

func (animal *AnimalController) Get() {
	page, err := animal.GetInt("page")
	if err != nil {
		page = 1
	}
	orderid, err := animal.GetInt("orderid")
	if err != nil {
		orderid = 1
	}
	listanimal, err := models.GetListAnimalRecent(page, orderid)
	if err != nil {
		log.Fatal(err)
	} else if listanimal == nil {
		data := make(map[string]interface{})
		data["message1"] = "没有animal数据！"
		animal.Data["json"] = map[string]interface{}{"id": "", "code": 0, "desc": "",
			"data": data}
	} else {
		listanimaldata := make([]map[string]interface{}, 0)
		for i := 0; i < len(listanimal); i++ {
			m := make(map[string]interface{})
			j, _ := json.Marshal(&listanimal[i])
			json.Unmarshal(j, &m)
			listanimaldata = append(listanimaldata, m)
		}
		animal.Data["json"] = map[string]interface{}{"id": "", "code": 0, "desc": "",
			"data": listanimaldata}
	}
	animal.ServeJSON()
}

/*func (animal *AnimalController) Post() {
	var request APIRequestAnimal
	json.Unmarshal(animal.Ctx.Input.RequestBody, &request)
	switch request.Type {
	case 1:

	default:

	}
}*/
