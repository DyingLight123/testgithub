package controllers

import (
	"encoding/json"
	"log"
	"testgithub/test/models"
)

type BackendAddAnimalController struct {
	BaseController
}

type BankendAddAnimalParam struct {
	Title 		string 		`json:"title"`
	Place 		string		`json:"place"`
	Phone 		string		`json:"phone"`
	Email 		string		`json:"email"`
	Describe	string		`json:"describe"`
}
type APIRequestBankendAddAnimal struct {
	models.APIRequestCommon
	BankendAddAnimalParam `json:"params"`
}

func (addanimal *BackendAddAnimalController) AddAnimal() {
	var request APIRequestBankendAddAnimal
	json.Unmarshal(addanimal.Ctx.Input.RequestBody, &request)
	if models.GetUserOneName(request.User) == false {
		addanimal.Data["json"] = map[string]interface{} {"id": request.Id, "code": 0, "desc": "",
			"data": map[string]interface{}{"message": "找不到该用户名对应的id！"}}
	} else {
		userid, err := models.GetUserIdByName(request.User)
		if err != nil {
			log.Fatal(err)
			addanimal.Data["json"] = map[string]interface{} {"id": request.Id, "code": 1, "desc": err,
				"data": map[string]interface{}{"message": "获取用户id失败！"}}
		}
		//fmt.Println(userid)
		animal := new(models.StrayAnimal)
		animal.Userid = userid
		animal.Title = request.Title
		animal.Place = request.Place
		animal.Phone = request.Phone
		animal.Email = request.Email
		animal.Describe = request.Describe

		id, err1 := models.AddAnimal(animal)
		if err1 != nil {
			log.Fatal(err1)
			addanimal.Data["json"] = map[string]interface{} {"id": request.Id, "code": 1, "desc": err1,
				"data": map[string]interface{}{"message": "添加信息失败！"}}
		}
		addanimal.Data["json"] = map[string]interface{} {"id": request.Id, "code": 0, "desc": "",
			"data": map[string]interface{}{"message": "信息添加成功！", "id": id}}
	}
	addanimal.ServeJSON()
}
