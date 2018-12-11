package models

import (
	"github.com/astaxie/beego/orm"
	"log"
)

type StoryLikeStatus struct {
	Id 				int
	Userid 			int
	Storycommonid	int
	Status 			int
}

func AddStoryCommonLike(userid int, storycommonid int) (int64, error) {
	o := orm.NewOrm()
	like := new(StoryLikeStatus)

	like.Userid = userid
	like.Storycommonid= storycommonid
	like.Status = 1

	id, err := o.Insert(like)
	return id, err
}

func IsEmptyStoryCommonLike(userid int, storycommonid int) bool {
	o := orm.NewOrm()
	like := new(StoryLikeStatus)
	qs := o.QueryTable("storylikestatus")
	err := qs.Filter("userid", userid).Filter("storycommonid", storycommonid).One(like)
	if err == orm.ErrNoRows {
		return true
	} else if err == nil {
		if like.Status == 0 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func DeleteStoryLikeAll(storyid int) error {
	storycommonid, err := GetStoryCommonIdByStoryId(storyid)
	if err != nil {
		log.Fatal(err)
	}
	o := orm.NewOrm()
	qs := o.QueryTable("storylikestatus")
	for i := 0; i < len(storycommonid); i++ {
		_, err := qs.Filter("storycommonid", storycommonid[i]).Delete()
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteStoryLikeByStoryCommonId(storycommonid int) error {
	o := orm.NewOrm()
	qs := o.QueryTable("storylikestatus")
	_, err := qs.Filter("storycommonid", storycommonid).Delete()
	if err != nil {
		return err
	}
	return nil
}
