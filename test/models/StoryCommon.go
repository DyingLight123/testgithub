package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type StoryCommon struct {
	Id  		int
	Userid 		int
	Storyid 	int
	Describe 	string
	Created 	time.Time
	Likecount 	int
}

func AddStoryCommon(common *StoryCommon) (int64, error) {
	o := orm.NewOrm()
	common1 := new(StoryCommon)

	common1.Userid = common.Userid
	common1.Storyid= common.Storyid
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
func GetStoryCommonListByStoryId(page int, storyid int, orderid int) ([]*StoryCommon, error) {
	o := orm.NewOrm()
	storycommon := make([]*StoryCommon, 0)
	qs := o.QueryTable("storycommon")
	offset := (page - 1) * 10
	var err error
	switch orderid {
	case 1:
		_, err = qs.Filter("storyid", storyid).OrderBy("created").Limit(10, offset).All(&storycommon)
	case 2:
		_, err = qs.Filter("storyid", storyid).OrderBy("-created").Limit(10, offset).All(&storycommon)
	case 3:
		_, err = qs.Filter("storyid", storyid).OrderBy("likecount").Limit(10, offset).All(&storycommon)
	case 4:
		_, err = qs.Filter("storyid", storyid).OrderBy("-likecount").Limit(10, offset).All(&storycommon)
	default:
		_, err = qs.Filter("storyid", storyid).OrderBy("created").Limit(10, offset).All(&storycommon)
	}
	if err != nil {
		return nil, err
	} else {
		return storycommon, nil
	}
}

func UpdateStoryLikeCountById(id int) error {
	o := orm.NewOrm()
	storycommon := new(StoryCommon)

	qs := o.QueryTable("storycommon")
	err := qs.Filter("id", id).One(storycommon)
	if err != nil {
		return err
	}
	storycommon.Likecount += 1
	_, err1 := o.Update(storycommon)
	if err1 != nil {
		return err1
	}
	return nil
}

func DeleteStoryCommonAll(storyid int) error {
	o := orm.NewOrm()

	qs := o.QueryTable("storycommon")
	_, err := qs.Filter("storyid", storyid).Delete()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func GetStoryCommonIdByStoryId(storyid int) ([]int, error) {
	o := orm.NewOrm()
	common := make([]*StoryCommon, 0)
	qs := o.QueryTable("storycommon")
	_, err := qs.Filter("storyid", storyid).All(&common)
	if err != nil {
		return nil, err
	}
	id := make([]int, 0)
	for i := 0; i < len(common); i++ {
		id = append(id, common[i].Id)
	}
	return id, nil
}

func GetStoryCommonListByUserId(userid int, page int, orderid int) ([]*StoryCommon, error) {
	o := orm.NewOrm()
	common := make([]*StoryCommon, 0)

	qs := o.QueryTable("storycommon")
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

func GetStoryCommonUserIdById(id int) (int, error) {
	o := orm.NewOrm()
	common := new(StoryCommon)
	qs := o.QueryTable("storycommon")
	err := qs.Filter("id", id).One(common)
	if err == orm.ErrNoRows {
		return -1, nil
	} else if err == nil {
		return common.Userid, nil
	} else {
		return -1, err
	}
}

func DeleteStoryCommonOneById(id int) error {
	o := orm.NewOrm()
	qs := o.QueryTable("storycommon")
	_, err := qs.Filter("id", id).Delete()
	if err != nil {
		return err
	}
	return nil
}
