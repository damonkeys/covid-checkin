package main

// This server uses environment variables. If they are not set the server won't start. It expects the following environment variables:
//   * SERVER_PORT       - the server is listening on this portnumber

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	biz "github.com/damonkeys/ch3ck1n/biz/models"
	"github.com/damonkeys/ch3ck1n/monkeys/config"
	"github.com/damonkeys/ch3ck1n/monkeys/database"
	"github.com/damonkeys/ch3ck1n/monkeys/tracing"
	"github.com/qor/admin"
	"github.com/qor/qor"
)

var (
	// Admin defines a static site for links
	Admin *admin.Admin
	// AdminBiz reprents the biz-DB connection for businesses
	AdminBiz *admin.Admin
)

type (
	// ServerConfigStruct holds the server-config
	ServerConfigStruct struct {
		Port     string                `env:"SERVER_PORT"`
		Database database.ConfigStruct `json:"database"`
	}
)

var serverConfig ServerConfigStruct

func main() {
	// Initalize
	// tracer init
	closer, _, ctx := tracing.InitJaeger("admin")
	defer closer.Close()

	// read config from environment variables to struct
	configInterface, err := config.ReadEnvVars(ctx, ServerConfigStruct{})
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	serverConfig = configInterface.(ServerConfigStruct)

	Admin = admin.New(&admin.AdminConfig{
		SiteName: "QOR-Links",
	})
	Admin.AddMenu(&admin.Menu{Name: "DBBiz", RelativePath: "/biz"})
	dbBiz := addBizAdmin()
	defer dbBiz.Close()
	// Register route
	mux := http.NewServeMux()

	// amount to /admin, so visit `/admin` to view the admin interface
	Admin.MountTo("/admin", mux)
	AdminBiz.MountTo("/admin/biz", mux)
	fmt.Printf("Listening on: %s\n", serverConfig.Port)
	http.ListenAndServe(":"+serverConfig.Port, mux)
}

func addBizAdmin() (db *gorm.DB) {
	// open database connection
	DBBiz, err := gorm.Open("mysql", serverConfig.Database.User+":"+serverConfig.Database.Password+"@"+serverConfig.Database.Server+"/"+serverConfig.Database.Name+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Println(err)
		DBBiz.Close()
	}

	AdminBiz = admin.New(&admin.AdminConfig{DB: DBBiz})

	// Create resources from GORM-backend model
	businesses := AdminBiz.AddResource(&biz.Business{})
	businesses.NewAttrs("-UUID")
	businessInfoMeta := businesses.Meta(&admin.Meta{Name: "BusinessInfos"})
	businessInfoResource := businessInfoMeta.Resource
	businessInfoResource.NewAttrs("-UUID")
	addDefaultScopes(businesses)

	businessInfos := AdminBiz.AddResource(&biz.BusinessInfo{})
	businessInfos.NewAttrs("-UUID")
	addDefaultScopes(businessInfos)

	return DBBiz
}

func addDefaultScopes(resource *admin.Resource) {
	resource.Scope(&admin.Scope{Name: "All", Handler: func(db *gorm.DB, context *qor.Context) *gorm.DB {
		return db.Unscoped().Find(resource.NewStruct())
	}})
}
