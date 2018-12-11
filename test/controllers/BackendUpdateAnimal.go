package controllers

import (
	"encoding/json"
	"log"
	"testgithub/test/models"
)

type BackendUpdateAnimalController struct {
	BaseController
}

type BankendUpdateAnimalParam struct {
	Id 			int 		`json:"id"`
	Title 		string 		`json:"title"`
	Place 		string		`json:"place"`
	Phone 		string		`json:"phone"`
	Email 		string		`json:"email"`
	Describe	string		`json:"describe"`
	Status 		int 		`json:"status"`
}

type APIRequestBankendUpdateAnimal struct {
	models.APIRequestCommon
	BankendUpdateAnimalParam `json:"params"`
}

func (updateanimal *BackendUpdateAnimalController) Get() {
	id, err := updateanimal.GetInt("id")
	if err != nil {
		log.Fatal(err)
	}
	animal, err1 := models.GetAnimalOneById(id)
	if err1 != nil{
		log.Fatal(err1)
	} else if animal == nil {
		updateanimal.Data["json"] = map[string]interface{} {"id": "", "code": 0, "desc": "",
			"data": map[string]interface{}{"message": "没找到您给定id的animal"}}
	} else {
		m := make(map[string]interface{})
		j, _ := json.Marshal(&animal)
		json.Unmarshal(j, &m)
		updateanimal.Data["json"] = map[string]interface{} {"id": "", "code": 0, "desc": "",
			"data": m}
	}
	updateanimal.ServeJSON()
}

func (updateanimal *BackendUpdateAnimalController) Post() {
	var request APIRequestBankendUpdateAnimal
	json.Unmarshal(updateanimal.Ctx.Input.RequestBody, &request)
	//fmt.Println(request)
	err := models.UpdateAnimal(request.BankendUpdateAnimalParam.Id, request.Title,
		request.Place, request.Phone, request.Email, request.Describe, request.Status)
	if err != nil {
		log.Fatal(err)
		updateanimal.Data["json"] = map[string]interface{} {"id": request.APIRequestCommon.Id, "code": 1, "desc": err,
			"data": map[string]interface{} {"message": "更新信息出错！"}}
	}
	updateanimal.Data["json"] = map[string]interface{} {"id": request.APIRequestCommon.Id, "code": 0, "desc": "",
		"data": map[string]interface{}{"message": "更新成功！"}}
	updateanimal.ServeJSON()
}



