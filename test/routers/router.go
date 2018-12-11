package routers

import (
	"github.com/astaxie/beego"
	"testgithub/test/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/login", &controllers.LoginController{}, "get:GetLogin;post:PostLogin")
	//登录
    beego.Router("/home", &controllers.HomeController{}, "get:GetHome")
	//主页
    beego.Router("/register", &controllers.RegisterController{})
    //注册
	beego.Router("/home/animal", &controllers.AnimalController{})
    //animal展示页面
    beego.Router("/home/animal/common", &controllers.AnimalCommonController{})
    //单个animal展示页面
	beego.Router("/home/animal/common/like", &controllers.AnimalCommonLikeController{})
    //animalcommon点赞
	beego.Router("/home/story", &controllers.StoryController{})
    //故事展示页面
    beego.Router("/home/story/common", &controllers.StoryCommonController{})
    //单个故事展示页面
    beego.Router("/home/story/common/like", &controllers.StoryCommonLikeController{})
    //storycommon点赞
    beego.Router("/backendhome", &controllers.BackendHomeController{}, "post:Home")
	//用户主页面
	beego.Router("/backendhome/updateuser", &controllers.UpdateUserController{})
    //用户修改详细信息
	beego.Router("/backendhome/animal", &controllers.BackendAnimalConntroller{})
    //用户animal
    beego.Router("/backendhome/animal/add", &controllers.BackendAddAnimalController{}, "post:AddAnimal")
	beego.Router("/backendhome/animal/update", &controllers.BackendUpdateAnimalController{})
    beego.Router("/backendhome/animal/delete", &controllers.BackendDeleteAnimalController{})
    //用户添加修改删除animal
    beego.Router("/backendhome/animal/common", &controllers.BackendAnimalCommonController{})
    //用户animal评论
    beego.Router("/backendhome/animal/common/delete", &controllers.BackendDeleteAnimalCommonController{})
    //用户animal评论删除
    beego.Router("/backendhome/story", &controllers.BackendStoryController{})
    //用户故事
    beego.Router("/backendhome/story/add", &controllers.BackendAddStoryController{}, "post:AddStory")
	beego.Router("/backendhome/story/update", &controllers.BackendUpdateStoryController{})
    beego.Router("/backendhome/story/delete", &controllers.BackendDeleteStoryController{})
    //用户添加故事修改删除
	beego.Router("/backendhome/story/common", &controllers.BackendStoryCommonController{})
    //用户故事评论
    beego.Router("/backendhome/story/common/delete", &controllers.BackendDeleteStoryCommonController{})
    //用户故事评论删除
}
