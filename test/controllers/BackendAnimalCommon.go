package controllers

import (
	"encoding/json"
	"log"
	"testgithub/test/models"
)

type BackendAnimalCommonController struct {
	BaseController
}

type BackendAnimalCommonParam struct {

}

type APIRequestBackendAnimalCommon struct {
	models.APIRequestCommon
	BackendAnimalCommonParam `json:"params"`
}

func (common *BackendAnimalCommonController) Post() {
	page, err := common.GetInt("page")
	if err != nil {
		page = 1
	}
	orderid, err := common.GetInt("orderid")
	if err != nil {
		orderid = 1
	}
	var request APIRequestBackendAnimalCommon
	json.Unmarshal(common.Ctx.Input.RequestBody, &request)
	if models.GetUserOneName(request.User) == false {
		common.Data["json"] = map[string]interface{} {"id": request.Id, "code": 0, "desc": "",
			"data": map[string]interface{}{"message": "找不到该用户名对应的id！"}}
	} else {
		userid, err := models.GetUserIdByName(request.User)
		if err != nil {
			log.Fatal(err)
		}
		commonlist, err := models.GetAnimalCommonListByUserId(userid, page, orderid)
		//fmt.Println(commonlist)
		if err != nil {
			log.Fatal(err)
		}
		if len(commonlist) == 0 {
			common.Data["json"] = map[string]interface{} {"id": request.Id, "code": 0, "desc": "",
				"data": map[string]interface{}{"message": "您还没有评论，请赶紧加入评论吧！"}}
		} else {
			list := make([]map[string]interface{}, 0)
			for i := 0; i < len(commonlist); i++ {
				m := make(map[string]interface{})
				j, _ := json.Marshal(&commonlist[i])
				json.Unmarshal(j, &m)
				list = append(list, m)
			}
			common.Data["json"] = map[string]interface{} {"id": request.Id, "code": 0, "desc": "",
				"data": list}
		}
	}
	common.ServeJSON()
}
