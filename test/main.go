package main

import (
	"github.com/astaxie/beego"
	_ "testgithub/test/routers"
	_ "testgithub/test/sqlinit"
)

func main() {
	beego.Run()
}

