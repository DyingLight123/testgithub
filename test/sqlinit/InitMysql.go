package sqlinit

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func InitMysql() {
	user := beego.AppConfig.String("mysqluser")
	password := beego.AppConfig.String("mysqlpass")
	host := beego.AppConfig.String("mysqlhost")
	post, err := beego.AppConfig.Int("mysqlpost")
	if err != nil {
		post = 3306
	}
	dbname := beego.AppConfig.String("mysqldb")

	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Local",
			user, password, host, post, dbname))
	//orm.RunSyncdb("default", false, true)
}
