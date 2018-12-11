package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"testgithub/test/models"
)

type AnimalCommonController struct {
	BaseController
}

type AnimalCommonParam struct {
	Animalid int    `json:"animalid"`
	Describe string `json:"describe"`
}

type APIRequestAnimalCommon struct {
	models.APIRequestCommon
	AnimalCommonParam `json:"params"`
}

func (animalcommon *AnimalCommonController) Get() {
	data := make(map[string]interface{})
	var err error
	animalid, err := animalcommon.GetInt("animalid")
	if err != nil {
		animalcommon.Redirect("/home/animal", 302)
		return
	}
	page, err := animalcommon.GetInt("page")
	if err != nil {
		page = 1
	}
	orderid, err := animalcommon.GetInt("orderid")
	if err != nil {
		orderid = 1
	}
	//fmt.Println(animalcommon.animalid, animalcommon.page, animalcommon.orderid)
	animal, err := models.GetAnimalOneById(animalid)
	if err != nil {
		log.Fatal(err)
	} else if animal == nil {
		data["message"] = "给的animalid不存在！"
		animalcommon.Data["json"] = map[string]interface{}{"id": "", "code": 1, "desc": "animalid不存在！",
			"data": data}
	} else {
		m := make(map[string]interface{})
		j, _ := json.Marshal(&animal)
		json.Unmarshal(j, &m)
		data["animal"] = m
		listanimalcommon, err := models.GetAnimalCommonListByAnimalId(page, animalid, orderid)
		if err != nil {
			log.Fatal(err)
		}
		list := make([]map[string]interface{}, 0)
		for i := 0; i < len(listanimalcommon); i++ {
			n := make(map[string]interface{})
			j, _ := json.Marshal(&listanimalcommon[i])
			json.Unmarshal(j, &n)
			list = append(list, n)
		}
		data["animalcommon"] = list
		animalcommon.Data["json"] = map[string]interface{}{"id": "", "code": 0, "desc": "",
			"data": data}
	}
	//animalcommon.Data["json"] = map[string]interface{}{"id": "", "code": 0, "desc": "",
	//	"data": data}
	animalcommon.ServeJSON()
}

func (animalcommon *AnimalCommonController) Post() {
	animalid, err := animalcommon.GetInt("animalid")
	if err != nil {
		animalcommon.Redirect("/home/animal", 302)
		return
	}
	page, err := animalcommon.GetInt("page")
	if err != nil {
		page = 1
	}
	orderid, err := animalcommon.GetInt("orderid")
	if err != nil {
		orderid = 1
	}
	var request APIRequestAnimalCommon
	json.Unmarshal(animalcommon.Ctx.Input.RequestBody, &request)
	//fmt.Println(animalcommon.animalid, animalcommon.page, animalcommon.orderid)
	if models.GetUserOneName(request.User) == false {
		returnURL := animalcommon.Ctx.Request.URL.Path
		returnURL = returnURL + fmt.Sprintf("?animalid=%d&page=%d&orderid=%d", animalid,
			page, orderid)
		animalcommon.Data["json"] = map[string]interface{}{"id": request.Id, "code": 0, "desc": "",
			"data": map[string]interface{}{"message": "请先登录！", "returnURL": returnURL}}
	} else {
		userid, err := models.GetUserIdByName(request.User)
		if err != nil {
			log.Fatal(err)
		}
		common := new(models.AnimalCommon)
		common.Userid = userid
		common.Animalid = request.Animalid
		common.Describe = request.Describe
		id, err := models.AddAnimalCommon(common)
		if err != nil {
			log.Fatal(err)
		}
		animalcommon.Data["json"] = map[string]interface{}{"id": request.Id, "code": 0, "desc": "",
			"data": map[string]interface{}{"message": "添加成功！", "id": id}}
	}
	animalcommon.ServeJSON()
}

