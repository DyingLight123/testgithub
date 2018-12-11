package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id 		int
	Name 	string
	Pwd 	string
}

func AddUser(user *User) (int64, error) {
	o := orm.NewOrm()
	user1 := new(User)

	user1.Name = user.Name
	user1.Pwd = user.Pwd

	id, err := o.Insert(user1)
	return id, err
}

func GetUserOneName(name string) bool {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable("user")
	err := qs.Filter("name", name).One(user)
	//fmt.Println(user)
	if err == orm.ErrNoRows {
		return false
	}
	return true
}

func GetUserOneByNama(name string) (string, error) {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable("user")
	err := qs.Filter("name", name).One(user)
	if err != nil {
		return "", err
	} else {
		return user.Pwd, nil
	}
}

func GetUserOne(name string) (*User, error) {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable("user")
	err := qs.Filter("name", name).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(name1 string, name string, pwd string) error {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable("user")
	err := qs.Filter("name", name1).One(user)
	if err != nil {
		return err
	}
	user.Name = name
	user.Pwd = pwd
	_, err1 := o.Update(user)
	if err1 != nil {
		return err1
	}
	return nil
}

func GetUserIdByName(name string) (int, error) {
	o := orm.NewOrm()
	user := new(User)

	qs := o.QueryTable("user")
	err := qs.Filter("name", name).One(user)
	if err != nil {
		return 0, err
	} else {
		return user.Id, nil
	}
}