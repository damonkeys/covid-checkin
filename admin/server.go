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

	authx "github.com/damonkeys/ch3ck1n/authx/models"
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
		DatabaseChckr    database.ConfigStruct
		DatabaseCheckins database.ConfigStruct
	}
)

var serverConfig ServerConfigStruct

func main() {
	// Initalize
	// tracer init
	closer, span, ctx := tracing.InitJaeger("admin")
	defer closer.Close()

	tracing.LogString(span, "Welcome", "Admin-Server started")

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

	// Chckr
	dbChckr := addChckrAdmin(ctx, mux)
	defer dbChckr.Close()
	Admin.AddMenu(&admin.Menu{Name: "DBChckr", RelativePath: "/chckr"})

	// Checkins
	dbCheckins := addCheckinsAdmin(ctx, mux)
	defer dbCheckins.Close()
	Admin.AddMenu(&admin.Menu{Name: "DBCheckins", RelativePath: "/checkins"})

	fmt.Printf("Listening on: %s\n", serverConfig.Port)
	http.ListenAndServe(":"+serverConfig.Port, mux)
}

func addChckrAdmin(c context.Context, mux *http.ServeMux) *gorm.DB {
	span := tracing.EnterWithContext(c)
	defer span.Finish()
	// open database connection
	db, err := gorm.Open("mysql", serverConfig.DatabaseChckr.User+":"+serverConfig.DatabaseChckr.Password+"@("+serverConfig.DatabaseChckr.Server+")/"+serverConfig.DatabaseChckr.Name+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
		db.Close()
	}

	Admin := admin.New(&admin.AdminConfig{DB: db})
	Admin.MountTo("/admin/chckr", mux)
	// Remove CSRF-check for working behind Kong
	Admin.GetRouter().GetMiddleware("csrf_check").Handler = func(context *admin.Context, middleware *admin.Middleware) { middleware.Next(context) }

	// Create resources from GORM-backend model
	businesses := Admin.AddResource(&biz.Business{}, &admin.Config{Menu: []string{"Biz"}})
	businesses.NewAttrs("-UUID")
	businessInfoMeta := businesses.Meta(&admin.Meta{Name: "BusinessInfos"})
	businessInfoResource := businessInfoMeta.Resource
	businessInfoResource.NewAttrs("-UUID")
	businessInfoResource.Meta(&admin.Meta{Name: "Description", Type: "rich_editor"})
	addDefaultScopes(c, businesses)

	businessInfos := Admin.AddResource(&biz.BusinessInfo{}, &admin.Config{Menu: []string{"Biz"}})
	businessInfos.NewAttrs("-UUID")
	addDefaultScopes(c, businessInfos)

	users := Admin.AddResource(&authx.User{}, &admin.Config{Menu: []string{"Authentications"}})
	addDefaultScopes(c, users)

	providers := Admin.AddResource(&authx.Provider{}, &admin.Config{Menu: []string{"Authentications"}})
	addDefaultScopes(c, providers)

	return db
}

func addCheckinsAdmin(c context.Context, mux *http.ServeMux) *gorm.DB {
	span := tracing.EnterWithContext(c)
	defer span.Finish()
	// open database connection
	db, err := gorm.Open("mysql", serverConfig.DatabaseCheckins.User+":"+serverConfig.DatabaseCheckins.Password+"@("+serverConfig.DatabaseCheckins.Server+")/"+serverConfig.DatabaseCheckins.Name+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
		db.Close()
	}

	Admin := admin.New(&admin.AdminConfig{DB: db})
	Admin.MountTo("/admin/checkins", mux)
	// Remove CSRF-check for working behind Kong
	Admin.GetRouter().GetMiddleware("csrf_check").Handler = func(context *admin.Context, middleware *admin.Middleware) { middleware.Next(context) }

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

	serverConfig.DatabaseChckr.Name = os.Getenv("DB_CHCKR_NAME")
	serverConfig.DatabaseChckr.Password = os.Getenv("DB_CHCKR_PASSWORD")
	serverConfig.DatabaseChckr.Server = os.Getenv("DB_CHCKR_HOST")
	serverConfig.DatabaseChckr.User = os.Getenv("DB_CHCKR_USER")

	serverConfig.DatabaseCheckins.Name = os.Getenv("DB_CHECKINS_NAME")
	serverConfig.DatabaseCheckins.Password = os.Getenv("DB_CHECKINS_PASSWORD")
	serverConfig.DatabaseCheckins.Server = os.Getenv("DB_CHECKINS_HOST")
	serverConfig.DatabaseCheckins.User = os.Getenv("DB_CHECKINS_USER")
}
