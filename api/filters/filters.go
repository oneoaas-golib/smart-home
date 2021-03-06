package filters

import (
	"github.com/e154/smart-home/api/log"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func RegisterFilters() {
	log.Info("Filters initialize...")

	// CORS
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "UPDATE"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type", "access_token"},
	}))

	// register access filter
	AccessFilter()
}
