package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type StrayAnimal struct {
	Id 			int
	Userid 		int
	Title 		string
	Place 		string
	Phone 		string
	Email 		string
	Describe	string
	Created 	time.Time
	Status 		int
}

func AddAnimal(animal *StrayAnimal) (int64, error) {
	o := orm.NewOrm()
	animal1 := new(StrayAnimal)

	animal1.Userid = animal.Userid
	animal1.Title = animal.Title
	animal1.Place = animal.Place
	animal1.Phone = animal.Phone
	animal1.Email = animal.Email
	animal1.Describe = animal.Describe
	animal1.Created = time.Now()
	animal1.Status = 0
	//fmt.Println(animal1)
	id, err := o.Insert(animal1)
	return id, err
}

func GetListAnimal() ([]*StrayAnimal, error) {
	o := orm.NewOrm()
	animal := make([]*StrayAnimal, 0)
	qs := o.QueryTable("animal")
	_, err := qs.OrderBy("-created").Limit(8).All(&animal)
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err == nil {
		return animal, nil
	} else {
		return nil, err
	}
}

func GetListAnimalRecent(page int, orderid int) ([]*StrayAnimal, error) {
	o := orm.NewOrm()
	animal := make([]*StrayAnimal, 0)
	qs := o.QueryTable("animal")
	offset := (page - 1) * 50
	var err error
	switch orderid {
	case 1:
		_, err = qs.OrderBy("-created").Limit(50, offset).All(&animal)
	case 2:
		_, err = qs.OrderBy("created").Limit(50, offset).All(&animal)
	}
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err == nil {
		return animal, nil
	} else {
		return nil, err
	}
}

func GetAnimalAllByUserId(page int, userid int) ([]*StrayAnimal, error) {
	o := orm.NewOrm()
	animal := make([]*StrayAnimal, 0)
	qs := o.QueryTable("animal")
	offset := (page - 1) * 10
	_, err := qs.Filter("userid", userid).OrderBy("-created").Limit(10, offset).All(&animal)
	if err != nil {
		return nil, err
	}
	return animal, nil
}

func GetAnimalOneById(id int) (*StrayAnimal, error) {
	o := orm.NewOrm()
	animal := new(StrayAnimal)
	qs := o.QueryTable("animal")
	err := qs.Filter("id", id).One(animal)
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err == nil {
		return animal, nil
	} else {
		return nil, err
	}
}

func UpdateAnimal(id int, title string, place string, phone string, email string, desctibe string, status int) error {
	o := orm.NewOrm()
	animal := new(StrayAnimal)
	qs := o.QueryTable("animal")
	err := qs.Filter("id", id).One(animal)
	if err != nil {
		return err
	}
	animal.Title = title
	animal.Place = place
	animal.Phone = phone
	animal.Email = email
	animal.Describe = desctibe
	animal.Status = status
	_, err1 := o.Update(animal)
	if err1 != nil {
		return err1
	}
	return nil
}

func DeleteAnimal(id int) error {
	o := orm.NewOrm()
	qs := o.QueryTable("animal")
	_, err := qs.Filter("id", id).Delete()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func GetAnimalOneUserIdById(id int) (int, error) {
	o := orm.NewOrm()
	animal := new(StrayAnimal)

	qs := o.QueryTable("animal")
	err := qs.Filter("id", id).One(animal)
	if err != nil {
		return 0, err
	} else {
		return animal.Userid, nil
	}
}

/*func GetAnimalOneId(id int) bool {
	o := orm.NewOrm()
	animal := new(StrayAnimal)

	qs := o.QueryTable("animal")
	err := qs.Filter("id", id).One(animal)
	if err == orm.ErrNoRows {
		return false
	} else {
		return true
	}
}*/

