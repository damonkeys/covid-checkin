package main

// This server uses environment variables. If they are not set the server won't start. It expects the following environment variables:
//   * SERVER_PORT       - the server is listening on this portnumber

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	biz "github.com/damonkeys/ch3ck1n/biz/business"
	"github.com/damonkeys/ch3ck1n/checkins/checkin"
	"github.com/damonkeys/ch3ck1n/monkeys/database"
	"github.com/damonkeys/ch3ck1n/monkeys/tracing"
	"github.com/qor/admin"
	"github.com/qor/qor"
)

type (
	// ServerConfigStruct holds the server-config
	ServerConfigStruct struct {
		Port             string `env:"SERVER_PORT"`
		DatabaseBiz      database.ConfigStruct
		DatabaseCheckins database.ConfigStruct
	}
)

var serverConfig ServerConfigStruct

func main() {
	// Initalize
	// tracer init
	closer, _, ctx := tracing.InitJaeger("admin")
	defer closer.Close()

	// read config from environment variables to struct
	readEnvVars(ctx)

	// Register route
	mux := http.NewServeMux()

	// Main-Admin-Page
	Admin := admin.New(&admin.AdminConfig{
		SiteName: "chckr",
	})
	// amount to /admin, so visit `/admin` to view the admin interface
	Admin.MountTo("/admin", mux)

	// Biz
	dbBiz := addBizAdmin(ctx, mux)
	defer dbBiz.Close()
	Admin.AddMenu(&admin.Menu{Name: "DBBiz", RelativePath: "/biz"})

	// Checkins
	dbCheckins := addCheckinsAdmin(ctx, mux)
	defer dbCheckins.Close()
	Admin.AddMenu(&admin.Menu{Name: "DBCheckins", RelativePath: "/checkins"})

	fmt.Printf("Listening on: %s\n", serverConfig.Port)
	http.ListenAndServe(":"+serverConfig.Port, mux)
}

func addBizAdmin(c context.Context, mux *http.ServeMux) *gorm.DB {
	span := tracing.EnterWithContext(c)
	defer span.Finish()
	// open database connection
	db, err := gorm.Open("mysql", serverConfig.DatabaseBiz.User+":"+serverConfig.DatabaseBiz.Password+"@"+serverConfig.DatabaseBiz.Server+"/"+serverConfig.DatabaseBiz.Name+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
		db.Close()
	}

	Admin := admin.New(&admin.AdminConfig{DB: db})
	Admin.MountTo("/admin/biz", mux)

	// Create resources from GORM-backend model
	businesses := Admin.AddResource(&biz.Business{})
	businesses.NewAttrs("-UUID")
	businessInfoMeta := businesses.Meta(&admin.Meta{Name: "BusinessInfos"})
	businessInfoResource := businessInfoMeta.Resource
	businessInfoResource.NewAttrs("-UUID")
	businessInfoResource.Meta(&admin.Meta{Name: "Description", Type: "rich_editor"})
	addDefaultScopes(c, businesses)

	businessInfos := Admin.AddResource(&biz.BusinessInfo{})
	businessInfos.NewAttrs("-UUID")
	addDefaultScopes(c, businessInfos)
	return db
}

func addCheckinsAdmin(c context.Context, mux *http.ServeMux) *gorm.DB {
	span := tracing.EnterWithContext(c)
	defer span.Finish()
	// open database connection
	db, err := gorm.Open("mysql", serverConfig.DatabaseCheckins.User+":"+serverConfig.DatabaseCheckins.Password+"@"+serverConfig.DatabaseCheckins.Server+"/"+serverConfig.DatabaseCheckins.Name+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
		db.Close()
	}

	Admin := admin.New(&admin.AdminConfig{DB: db})
	Admin.MountTo("/admin/checkins", mux)

	// Create resources from GORM-backend model
	checkins := Admin.AddResource(&checkin.Checkin{})
	checkins.NewAttrs("-UUID")
	addDefaultScopes(c, checkins)
	return db
}

func addDefaultScopes(c context.Context, resource *admin.Resource) {
	span := tracing.EnterWithContext(c)
	defer span.Finish()

	resource.Scope(&admin.Scope{Name: "All", Handler: func(db *gorm.DB, context *qor.Context) *gorm.DB {
		return db.Unscoped().Find(resource.NewStruct())
	}})
}

func readEnvVars(c context.Context) {
	span := tracing.EnterWithContext(c)
	defer span.Finish()

	serverConfig = ServerConfigStruct{}
	serverConfig.Port = os.Getenv("SERVER_PORT")

	serverConfig.DatabaseBiz.Name = os.Getenv("DB_BIZ_NAME")
	serverConfig.DatabaseBiz.Password = os.Getenv("DB_BIZ_PASSWORD")
	serverConfig.DatabaseBiz.Server = os.Getenv("DB_BIZ_HOST")
	serverConfig.DatabaseBiz.User = os.Getenv("DB_BIZ_USER")

	serverConfig.DatabaseCheckins.Name = os.Getenv("DB_CHECKINS_NAME")
	serverConfig.DatabaseCheckins.Password = os.Getenv("DB_CHECKINS_PASSWORD")
	serverConfig.DatabaseCheckins.Server = os.Getenv("DB_CHECKINS_HOST")
	serverConfig.DatabaseCheckins.User = os.Getenv("DB_CHECKINS_USER")
}
