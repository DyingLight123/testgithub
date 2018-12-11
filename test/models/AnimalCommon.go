package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type AnimalCommon struct {
	Id 			int
	Userid 		int
	Animalid 	int
	Created 	time.Time
	Describe 	string
	Likecount 	int
}

func AddAnimalCommon(common *AnimalCommon) (int64, error) {
	o := orm.NewOrm()
	common1 := new(AnimalCommon)

	common1.Userid = common.Userid
	common1.Animalid = common.Animalid
	common1.Describe = common.Describe
	common1.Created = time.Now()
	common1.Likecount =0

	id, err := o.Insert(common1)
	return id, err
}

/*orderid :
	1:created;
	2:-created;
	3:likecount;
	4:-likecount;*/
func GetAnimalCommonListByAnimalId(page int, animalid int, orderid int) ([]*AnimalCommon, error) {
	o := orm.NewOrm()
	animalcommon := make([]*AnimalCommon, 0)
	qs := o.QueryTable("animalcommon")
	offset := (page - 1) * 10
	var err error
	switch orderid {
	case 1:
		_, err = qs.Filter("animalid", animalid).OrderBy("created").Limit(10, offset).All(&animalcommon)
	case 2:
		_, err = qs.Filter("animalid", animalid).OrderBy("-created").Limit(10, offset).All(&animalcommon)
	case 3:
		_, err = qs.Filter("animalid", animalid).OrderBy("likecount").Limit(10, offset).All(&animalcommon)
	case 4:
		_, err = qs.Filter("animalid", animalid).OrderBy("-likecount").Limit(10, offset).All(&animalcommon)
	default:
		_, err = qs.Filter("animalid", animalid).OrderBy("created").Limit(10, offset).All(&animalcommon)
	}
	//_, err := qs.Filter("animalid", animalid).OrderBy("created").Limit(10, offset).All(&animalcommon)
	if err != nil {
		return nil, err
	} else {
		return animalcommon, nil
	}
}

func UpdateLikeCountById(id int) error {
	o := orm.NewOrm()
	animalcommon := new(AnimalCommon)

	qs := o.QueryTable("animalcommon")
	err := qs.Filter("id", id).One(animalcommon)
	if err != nil {
		return err
	}
	animalcommon.Likecount += 1
	_, err1 := o.Update(animalcommon)
	if err1 != nil {
		return err1
	}
	return nil
}

func DeleteAnimalCommonAll(animalid int) error {
	o := orm.NewOrm()

	qs := o.QueryTable("animalcommon")
	_, err := qs.Filter("animalid", animalid).Delete()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func GetAnimalCommonIdByAnimalId(animalid int) ([]int, error) {
	o := orm.NewOrm()
	common := make([]*AnimalCommon, 0)
	qs := o.QueryTable("animalcommon")
	_, err := qs.Filter("animalid", animalid).All(&common)
	if err != nil {
		return nil, err
	}
	id := make([]int, 0)
	for i := 0; i < len(common); i++ {
		id = append(id, common[i].Id)
	}
	return id, nil
}

func GetAnimalCommonListByUserId(userid int, page int, orderid int) ([]*AnimalCommon, error) {
	o := orm.NewOrm()
	common := make([]*AnimalCommon, 0)

	qs := o.QueryTable("animalcommon")
	var err error
	offset := (page - 1) * 10
	switch orderid {
	case 1:
		_, err = qs.Filter("userid", userid).OrderBy("created").Limit(10, offset).All(&common)
	case 2:
		_, err = qs.Filter("userid", userid).OrderBy("-created").Limit(10, offset).All(&common)
	case 3:
		_, err = qs.Filter("userid", userid).OrderBy("likecount").Limit(10, offset).All(&common)
	case 4:
		_, err = qs.Filter("userid", userid).OrderBy("-likecount").Limit(10, offset).All(&common)
	default:
		_, err = qs.Filter("userid", userid).OrderBy("created").Limit(10, offset).All(&common)
	}
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err == nil {
		return common, nil
	} else {
		return nil, err
	}
}

func GetAnimalCommonUserIdById(id int) (int, error) {
	o := orm.NewOrm()
	common := new(AnimalCommon)
	qs := o.QueryTable("animalcommon")
	err := qs.Filter("id", id).One(common)
	if err == orm.ErrNoRows {
		return -1, nil
	} else if err == nil {
		return common.Userid, nil
	} else {
		return -1, err
	}
}

func DeleteAnimalCommonOneById(id int) error {
	o := orm.NewOrm()
	qs := o.QueryTable("animalcommon")
	_, err := qs.Filter("id", id).Delete()
	if err != nil {
		return err
	}
	return nil
}



