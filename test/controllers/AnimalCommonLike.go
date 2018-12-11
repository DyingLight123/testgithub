package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"testgithub/test/models"
)

type AnimalCommonLikeController struct {
	BaseController
}

type AnimalCommonLikeParam struct {
	AnimalCommonId int `json:"animalcommonid"`
}

type APIRequestAnimalCommonLike struct {
	models.APIRequestCommon
	AnimalCommonLikeParam `json:"params"`
}

func (like *AnimalCommonLikeController) Post() {
	animalid, err := like.GetInt("animalid")
	if err != nil {
		like.Redirect("/home/animal", 302)
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
	var request APIRequestAnimalCommonLike
	json.Unmarshal(like.Ctx.Input.RequestBody, &request)
	if models.GetUserOneName(request.User) == false {
		returnURL := "/home/animal/common"
		returnURL = returnURL + fmt.Sprintf("?animalid=%d&page=%d&orderid=%d", animalid, page, orderid)
		like.Data["json"] = map[string]interface{}{"id": request.Id, "code": 0, "desc": "",
			"data": map[string]interface{}{"message": "请先登录！", "returnURL": returnURL}}
	} else {
		userid, err := models.GetUserIdByName(request.User)
		if err != nil {
			log.Fatal(err)
		}
		if models.IsEmptyAnimalCommonLike(userid, request.AnimalCommonId) == true {
			id, err := models.AddAnimalCommonLike(userid, request.AnimalCommonId)
			if err != nil {
				log.Fatal(err)
			}
			err = models.UpdateLikeCountById(request.AnimalCommonId)
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
	strURL := "/home/animal/common"
	strURL = strURL + fmt.Sprintf("?animalid=%d&page=%d&orderid=%d", animalid, page, orderid)
	like.Redirect(strURL, 302)
	like.ServeJSON()
}
