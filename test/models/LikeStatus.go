package models

import (
	"github.com/astaxie/beego/orm"
	"log"
)

type LikeStatus struct {
	Id 				int
	Userid 			int
	Animalcommonid 	int
	Status 			int
}

func AddAnimalCommonLike(userid int, animalcommonid int) (int64, error) {
	o := orm.NewOrm()
	like := new(LikeStatus)

	like.Userid = userid
	like.Animalcommonid = animalcommonid
	like.Status = 1

	id, err := o.Insert(like)
	return id, err
}

func IsEmptyAnimalCommonLike(userid int, animalcommonid int) bool {
	o := orm.NewOrm()
	like := new(LikeStatus)

	qs := o.QueryTable("likestatus")
	err := qs.Filter("userid", userid).Filter("animalcommonid", animalcommonid).One(like)
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

func DeleteAnimalLikeAll(animalid int) error {
	animalcommonid, err := GetAnimalCommonIdByAnimalId(animalid)
	if err != nil {
		log.Fatal(err)
	}
	o := orm.NewOrm()
	qs := o.QueryTable("likestatus")
	for i := 0; i < len(animalcommonid); i++ {
		_, err := qs.Filter("animalcommonid", animalcommonid[i]).Delete()
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteAnimalLikeByAnimalCommonId(animalcommonid int) error {
	o := orm.NewOrm()
	qs := o.QueryTable("likestatus")
	_, err := qs.Filter("animalcommonid", animalcommonid).Delete()
	if err != nil {
		return err
	}
	return nil
}
