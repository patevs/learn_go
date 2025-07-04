package main_pos

import (
	"fmt"
	"pos_api/businesses/pos/scripts/service"
	_ "pos_api/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/astaxie/beego/session/mysql"
	"github.com/astaxie/beego/toolbox"
	_ "github.com/go-sql-driver/mysql"
)

func main_pos() {
	runmode := beego.AppConfig.String("runmode")

	sqlconn := beego.AppConfig.String("core_dns")
	err := orm.RegisterDataBase("default", "mysql", sqlconn)
	if err != nil {
		fmt.Println("Core database connect error!")
		return
	}
	orm.SetMaxIdleConns("default", 50)
	if runmode == "product" {
		orm.SetMaxOpenConns("default", 600)
	} else {
		orm.SetMaxOpenConns("default", 100)
	}

	sqlconn = beego.AppConfig.String("pos_dns")
	err = orm.RegisterDataBase("pos", "mysql", sqlconn)
	if err != nil {
		fmt.Println("POS database connect error!")
		return
	}
	orm.SetMaxIdleConns("pos", 50)
	if runmode == "product" {
		orm.SetMaxOpenConns("pos", 600)
	} else {
		orm.SetMaxOpenConns("pos", 100)
	}

	// tk := toolbox.NewTask("EmailTask", "0 0/01 * * * *", func() error {
	// 	return EmailScripts.EmailTask()
	// })
	// toolbox.AddTask("EmailTask", tk)

	// tk = toolbox.NewTask("PullAllOtherData", "0 0/16 * * * *", func() error {
	// 	return service.AdvancedPullStaticData()
	// })
	// toolbox.AddTask("PullAllOtherData", tk)

	// temporary stop in test environment
	// only runs in production environment
	// if runmode == "product" {
	// tk = toolbox.NewTask("PullStockData", "0 01/05 * * * *", func() error {
	// 	return service.AdvancedStock()
	// })
	// toolbox.AddTask("PullStockData", tk)
	// // }

	// tk = toolbox.NewTask("PullNonStockData", "0 05/05 * * * *", func() error {
	// 	return service.AdvancedNonStock()
	// })
	// toolbox.AddTask("PullNonStockData", tk)

	// tk = toolbox.NewTask("PullDiscountCodeData", "0 02/05 * * * *", func() error {
	// 	return service.AdvancedDiscountCode()
	// })
	// toolbox.AddTask("PullDiscountCodeData", tk)

	// tk = toolbox.NewTask("PullCustomerData", "0 03/05 * * * *", func() error {
	// 	return service.AdvancedCustomer()
	// })
	// toolbox.AddTask("PullCustomerData", tk)

	// tk = toolbox.NewTask("PullNightData22", "0 01,33 22 * * *", func() error {
	// 	return service.AdvancedNight22()
	// })
	// toolbox.AddTask("PullNightData22", tk)

	// tk = toolbox.NewTask("PullNightData23", "0 01,33 23 * * *", func() error {
	// 	return service.AdvancedNight23()
	// })
	// toolbox.AddTask("PullNightData23", tk)

	// tk = toolbox.NewTask("PullDiscountData", "0 04/05 * * * *", func() error {
	// 	return service.AdvancedDiscount()
	// })
	// toolbox.AddTask("PullDiscountData", tk)

	// tk = toolbox.NewTask("PullSalesPriceData", "0 0/05 * * * *", func() error {
	// 	return service.AdvancedSalesPrice()
	// })
	// toolbox.AddTask("PullSalesPriceData", tk)

	// tk = toolbox.NewTask("InvoiceHold", "0 04/30 * * * *", func() error {
	// 	return service.WriteHoldInvoiceToAdvanced()
	// })
	// toolbox.AddTask("InvoiceHold", tk)

	// tk = toolbox.NewTask("LaybyHold", "0 03/30 * * * *", func() error {
	// 	return service.WriteHoldLaybyToAdvanced()
	// })
	// toolbox.AddTask("LaybyHold", tk)

	// tk = toolbox.NewTask("AccountPaymentHold", "0 06/30 * * * *", func() error {
	// 	return service.WriteHoldAccountPaymentToAdvanced()
	// })
	// toolbox.AddTask("AccountPaymentHold", tk)

	// tk = toolbox.NewTask("DeleteUserLogs", "0 01/30 22 * * *", func() error {
	// 	return service.DeleteUserLogs()
	// })
	// toolbox.AddTask("DeleteUserLogs", tk)

	// tk = toolbox.NewTask("DeleteFELogs", "0 03/30 22 * * *", func() error {
	// 	return service.DeleteFELogs()
	// })
	// toolbox.AddTask("DeleteFELogs", tk)

	// force logout
	tk := toolbox.NewTask("ForceLogOut", "0 02 01 * * *", func() error {
		return service.LogoutSystem()
	})
	toolbox.AddTask("ForceLogOut", tk)

	// runs in all modes
	toolbox.StartTask()
	defer toolbox.StopTask()

	UploadDir := beego.AppConfig.String("UploadDir")
	beego.SetStaticPath("/Uploads", UploadDir)
	cdnUrls := beego.AppConfig.Strings("cdn_urls")
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowCredentials: true,
		AllowOrigins:     cdnUrls,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "X-Requested-With", "Authorization", "Access-Control-Request-Method", "Access-Control-Request-Headers", "Host", "Content-Type", "Accept", "if-modified-since", "soapaction"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Credentials", "Access-Control-Expose-Headers", "Content-Type", "Authorization"},
	}))
	beego.BConfig.Log.AccessLogs = true

	beego.Run()
}
