package main

import (
	"einvite/backend/repositories"
	"einvite/backend/repositories/mongo"
	servicesimpl "einvite/backend/services"
	services "einvite/common/services"
	"einvite/framework"
	"einvite/frontend/auth"
	rest "einvite/frontend/rest"
	"einvite/frontend/web"
	wssockjs "einvite/frontend/ws/sockjs"
	"fmt"
	"github.com/fzzy/sockjs-go/sockjs"
	"log"
	"net/http"
)

type wsConfig struct {
	Port string
	SSL  bool
}

func bootstrapHttp() {

	fmt.Println("Bootstrapping Http")

	_configure()

	_registerRepositories()
	_registerServices()

	_registerRestControllers()
	_registerAuthControllers()

	server := _createHttpServer()

	_registerRestRoutes(server)

	_registerAuthRoutes(server)

	server.ListenAndServe()
}

func bootstrapWebsocket() {

	fmt.Println("Bootstrapping Websocket")

	_configure()

	_registerRepositories()
	_registerServices()

	_registerWSRoutes()

	_startWebsocketServer()
}

func _configure() {
	framework.Config.Init("config.json")
}

func _registerRepositories() {
	framework.SetFactory("repositories.UserRepository", func() interface{} {
		return mongo.NewUserRepository()
	})

	framework.SetFactory("repositories.EventRepository", func() interface{} {
		return mongo.NewEventRepository()
	})

	framework.SetFactory("repositories.SessionRepository", func() interface{} {
		return mongo.NewSessionRepository()
	})
}

func _registerServices() {
	framework.SetFactory("services.UserService", func() interface{} {

		repository := framework.GetFactory("repositories.UserRepository").(repositories.UserRepository)
		return servicesimpl.NewUserService(repository)
	})

	framework.SetFactory("services.EventService", func() interface{} {

		repository := framework.GetFactory("repositories.EventRepository").(repositories.EventRepository)
		return servicesimpl.NewEventService(repository)
	})

	framework.SetFactory("services.SessionService", func() interface{} {

		repository := framework.GetFactory("repositories.SessionRepository").(repositories.SessionRepository)
		return servicesimpl.NewSessionService(repository)
	})
}

func _registerRestControllers() {

	framework.SetFactory("controllers.UserController", func() interface{} {
		service := framework.GetFactory("services.UserService").(services.UserService)
		return rest.NewUserController(service)
	})

	framework.SetFactory("controllers.EventController", func() interface{} {
		service := framework.GetFactory("services.EventService").(services.EventService)
		return rest.NewEventController(service)
	})
}

func _registerAuthControllers() {

	framework.SetFactory("controllers.GoogleController", func() interface{} {
		service := framework.GetFactory("services.UserService").(services.UserService)
		return auth.NewGoogleController(service)
	})

	framework.SetFactory("controllers.FacebookController", func() interface{} {
		service := framework.GetFactory("services.UserService").(services.UserService)
		return auth.NewFacebookController(service)
	})
}

func _registerWSRoutes() {

}

func _registerRestRoutes(server framework.WebServer) {

	userController := framework.GetFactory("controllers.UserController").(*rest.UserController)

	//user routes
	server.Get("/users", userController.ListUsers)
	server.Get("/user", userController.GetUser)
	server.Post("/user", userController.CreateUser)

	server.Get("/who", userController.Who)

	//event routes
	eventController := framework.GetFactory("controllers.EventController").(*rest.EventController)

	server.Get("/test", eventController.Test)
}

func _registerAuthRoutes(server framework.WebServer) {

	var google = framework.GetFactory("controllers.GoogleController").(*auth.GoogleController)
	var facebook = framework.GetFactory("controllers.FacebookController").(*auth.FacebookController)

	server.Get("/", google.HandleRoot)

	server.Get("/load", google.TestHighLoad)

	server.Post("/auth/facebook", facebook.Auth)
	server.Get("/auth/facebook/callback", facebook.AuthCallback)

	server.Post("/auth/google", google.Auth)
	server.Get("/auth/google/callback", google.AuthCallback)
}

func _createHttpServer() framework.WebServer {
	httpCfg := &web.HttpConfig{}
	framework.Config.ReadInto("http", &httpCfg)

	sessionCfg := &web.SessionConfig{}
	framework.Config.ReadInto("session", &sessionCfg)

	securityCfg := &web.SecurityKeys{}
	framework.Config.ReadInto("security", &securityCfg)
	//TODO: setting public properties is not cool.. should be automatically set from the config file
	securityCfg.RawSignKey = []byte(securityCfg.SignKey)
	securityCfg.RawEncryptionKey = []byte(securityCfg.EncryptionKey)

	webCfg := &web.WebConfig{Http: httpCfg, Session: sessionCfg, Security: securityCfg}

	sessionService := framework.GetFactory("services.SessionService").(services.SessionService)

	return web.NewWebServer(webCfg, sessionService)
}

func _startWebsocketServer() {

	wsCfg := &wsConfig{}
	framework.Config.ReadInto("websocket", &wsCfg)

	log.Println("Starting websocket server on port", wsCfg.Port)

	server := sockjs.NewServeMux(http.DefaultServeMux)
	conf := sockjs.NewConfig()

	server.Handle("/ws", wssockjs.HandleSockjsSession, conf)

	http.ListenAndServe(":"+wsCfg.Port, server)

}
