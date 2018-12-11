package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(User), new(StrayAnimal), new(AnimalCommon), new(LikeStatus), new(Story), new(StoryCommon),
		new(StoryLikeStatus))
}

func (user *User) TableName() string {
	return "user"
}

func (animal *StrayAnimal) TableName() string {
	return "animal"
}

func (animalcommon *AnimalCommon) TableName() string {
	return "animalcommon"
}

func (likestatus *LikeStatus) TableName() string {
	return "likestatus"
}

func (story *Story) TableName() string {
	return "story"
}

func (storycommon *StoryCommon) TableName() string {
	return "storycommon"
}

func (storylikestatus *StoryLikeStatus) TableName() string {
	return "storylikestatus"
}

