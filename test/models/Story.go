package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Story struct {
	Id 			int
	Userid		int
	Title 		string
	Describe 	string
	Created 	time.Time
}

func AddStory(story *Story) (int64, error) {
	o := orm.NewOrm()
	story1 := new(Story)

	story1.Userid = story.Userid
	story1.Title = story.Title
	story1.Describe = story.Describe
	story1.Created = time.Now()

	id, err := o.Insert(story1)
	return id, err
}

func GetStoryAllByUserId(page int, userid int) ([]*Story, error) {
	o := orm.NewOrm()
	story := make([]*Story, 0)
	qs := o.QueryTable("story")
	offset := (page - 1) * 10
	_, err := qs.Filter("userid", userid).OrderBy("-created").Limit(10, offset).All(&story)
	if err != nil {
		return nil, err
	} else {
		return story, nil
	}
}

func GetStoryOneById(id int) (*Story, error) {
	o := orm.NewOrm()
	story := new(Story)
	qs := o.QueryTable("story")
	err := qs.Filter("id", id).One(story)
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err == nil {
		return story, nil
	} else {
		return nil, err
	}
}

func UpdateStory(id int, title string, desctibe string) error {
	o := orm.NewOrm()
	story := new(Story)
	qs := o.QueryTable("story")
	err := qs.Filter("id", id).One(story)
	if err != nil {
		return err
	}
	story.Title = title
	story.Describe = desctibe
	_, err1 := o.Update(story)
	if err1 != nil {
		return err1
	}
	return nil
}

func GetStoryOneUserIdById(id int) (int, error) {
	o := orm.NewOrm()
	story := new(Story)

	qs := o.QueryTable("story")
	err := qs.Filter("id", id).One(story)
	if err != nil {
		return 0, err
	} else {
		return story.Userid, nil
	}
}

func DeleteStory(id int) error {
	o := orm.NewOrm()
	qs := o.QueryTable("story")
	_, err := qs.Filter("id", id).Delete()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func GetListStory() ([]*Story, error) {
	o := orm.NewOrm()
	story := make([]*Story, 0)
	qs := o.QueryTable("story")
	_, err := qs.OrderBy("-created").Limit(8).All(&story)
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err == nil {
		return story, nil
	} else {
		return nil, err
	}
}

func GetListStoryRecent(page int, orderid int) ([]*Story, error) {
	o := orm.NewOrm()
	story := make([]*Story, 0)
	qs := o.QueryTable("story")
	offset := (page - 1) * 50
	var err error
	switch orderid {
	case 1:
		_, err = qs.OrderBy("-created").Limit(50, offset).All(&story)
	case 2:
		_, err = qs.OrderBy("created").Limit(50, offset).All(&story)
	}
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err == nil {
		return story, nil
	} else {
		return nil, err
	}
}