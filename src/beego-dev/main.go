package main

import (
	_ "beego-api/routers"
	"fmt"

	// "github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
	// "github.com/beego/beego"
	// "github.com/beego/beego/orm"
	// "github.com/beego/beego/v2/server/web"
	// "github.com/beego/beego/v2"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"

	// don't forget this
	_ "github.com/go-sql-driver/mysql"
	// _ "github.com/mattn/go-sqlite3"
)

// User -
type User struct {
	ID   int    `orm:"column(id)"`
	Name string `orm:"column(name)"`
}

func init() {
	// runmode := beego.AppConfig.String("runmode")
	sqlconn, _ := beego.AppConfig.String("sqlconn")

	// need to register models in init
	orm.RegisterModel(new(User))

	err := orm.RegisterDataBase("default", "mysql", sqlconn)
	if err != nil {
		fmt.Println("Core database connect error!")
		return
	}
}

func main() {

	// automatically build table
	orm.RunSyncdb("default", false, true)

	// create orm object
	o := orm.NewOrm()

	// data
	user := new(User)
	user.Name = "mike"

	// insert data
	o.Insert(user)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
