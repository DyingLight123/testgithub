package controllers

import (
	"encoding/json"
	"log"
	"testgithub/test/models"
)

type HomeController struct {
	BaseController
}

func (home HomeController) GetHome() {
	data := make(map[string]interface{})

	animal, err := models.GetListAnimal()
	if err != nil {
		log.Fatal(err)
	} else if animal == nil {
		data["message1"] = "没有animal数据！"
	} else {
		animaldata := make([]map[string]interface{}, 0)
		for i := 0; i < len(animal); i++ {
			m := make(map[string]interface{})
			j, _ := json.Marshal(&animal[i])
			json.Unmarshal(j, &m)
			animaldata = append(animaldata, m)
		}
		data["animaldata"] = animaldata
	}

	story, err1 := models.GetListStory()
	if err1 != nil {
		log.Fatal(err1)
	} else if story == nil {
		data["message2"] = "没有story数据！"
	} else {
		storydata := make([]map[string]interface{}, 0)
		for i := 0; i < len(story); i++ {
			m := make(map[string]interface{})
			j, _ := json.Marshal(&story[i])
			json.Unmarshal(j, &m)
			storydata = append(storydata, m)
		}
		data["storydata"] = storydata
	}

	home.Data["json"]= map[string]interface{}{"id": "", "code": 0, "desc": "",
		"data": data}
	home.ServeJSON()
}
