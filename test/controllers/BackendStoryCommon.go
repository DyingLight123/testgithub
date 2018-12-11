package controllers

import (
	"encoding/json"
	"log"
	"testgithub/test/models"
)

type BackendStoryCommonController struct {
	BaseController
}

type BackendStoryCommonParam struct {

}

type APIRequestBackendStoryCommon struct {
	models.APIRequestCommon
	BackendStoryCommonParam `json:"params"`
}

func (common *BackendStoryCommonController) Post() {
	page, err := common.GetInt("page")
	if err != nil {
		page = 1
	}
	orderid, err := common.GetInt("orderid")
	if err != nil {
		orderid = 1
	}
	var request APIRequestBackendStoryCommon
	json.Unmarshal(common.Ctx.Input.RequestBody, &request)
	if models.GetUserOneName(request.User) == false {
		common.Data["json"] = map[string]interface{} {"id": request.Id, "code": 0, "desc": "",
			"data": map[string]interface{}{"message": "找不到该用户名对应的id！"}}
	} else {
		userid, err := models.GetUserIdByName(request.User)
		if err != nil {
			log.Fatal(err)
		}
		commonlist, err := models.GetStoryCommonListByUserId(userid, page, orderid)
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
