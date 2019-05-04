package actions

import (
	"net/http"
	"os"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/worker"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/events"
	forcessl "github.com/gobuffalo/mw-forcessl"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
	redisw "github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"github.com/unrolled/secure"
	"gopkg.in/boj/redistore.v1"

	"github.com/gobuffalo/buffalo-pop/pop/popmw"
	gwa "github.com/gobuffalo/gocraft-work-adapter"
	i18n "github.com/gobuffalo/mw-i18n"
	"github.com/gobuffalo/packr/v2"
	"github.com/myWebsite/golang/models"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
//
// Routing, middleware, groups, etc... are declared TOP -> DOWN.
// This means if you add a middleware to `app` *after* declaring a
// group, that group will NOT have that new middleware. The same
// is true of resource declarations as well.
//
// It also means that routes are checked in the order they are declared.
// `ServeFiles` is a CATCH-ALL route, so it should always be
// placed last in the route declarations, as it will prevent routes
// declared after it to never be called.
func App() *buffalo.App {
	if app == nil {
		store, _ := redistore.NewRediStoreWithDB(20, "tcp", os.Getenv("REDIS_SERVER")+":"+os.Getenv("REDIS_PORT"), os.Getenv("REDIS_PASSWORD"), os.Getenv("REDIS_DB_SESSION"), []byte("_baron_session"))
		app = buffalo.New(buffalo.Options{
			Env:          ENV,
			SessionName:  "_baron_session",
			SessionStore: store,
			Worker: gwa.New(gwa.Options{
				Pool: &redisw.Pool{
					MaxActive: 5,
					MaxIdle:   5,
					Wait:      true,
					Dial: func() (redisw.Conn, error) {
						c, err := redisw.Dial("tcp", os.Getenv("REDIS_SERVER")+":"+os.Getenv("REDIS_PORT"))
						if err != nil {
							return nil, err
						}
						if _, err := c.Do("AUTH", os.Getenv("REDIS_PASSWORD")); err != nil {
							c.Close()
							return nil, err
						}

						return c, err
					},
				},
				Name:           "platform_worker",
				MaxConcurrency: 25,
			}),
		})
		if ENV == "development" {
			// Log request parameters (filters apply).
			app.Use(paramlogger.ParameterLogger)
		}
		// Automatically redirect to SSL
		// app.Use(forceSSL())

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		//app.Use(csrf.New)

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.Connection)
		// Remove to disable this.
		app.Use(popmw.Transaction(models.DB))

		// Setup and use translations:
		app.Use(translations())

		app.Use(SetCurrentUser)
		app.Use(Authorize)

		app.GET("/test", Test)

		app.GET("/", HomeHandler)

		app.GET("/about", AboutHandler)

		app.GET("/projects", ProjectsHandler)

		app.GET("/welcome", RedirectPlatformHandler)

		app.GET("/dashboard", VueHandler)

		app.GET("/dashboard/{rest:.*}", VueHandler)

		// Redirects to VUE
		app.GET("/login", VueHandler)

		app.GET("/register", VueHandler)
		//END Redirects to VUE

		app.Middleware.Skip(Authorize, HomeHandler, AboutHandler, VueHandler)
		//PLATFORMS
		platform := Platform{}

		app.GET("/github/callback", platform.GithubCallback)

		//END PLATFORMS
		//API
		api := app.Group("/api")

		//AUTH
		auth := api.Group("/auth")

		auth.POST("/login", AuthCreate)
		auth.DELETE("/logout", AuthDestroy)
		auth.GET("/refresh", AuthRefresh)

		auth.Middleware.Skip(Authorize, AuthCreate)
		//AUTH
		var user buffalo.Resource
		user = &UsersResource{&buffalo.BaseResource{}}
		resourceUser := api.Resource("/user", user)
		resourceUser.Middleware.Skip(Authorize, user.Create)
		api.Resource("/project", ProjectsResource{})

		api.Resource("/comment", CommentsResource{})

		api.Resource("/task", TasksResource{})

		api.Resource("/license", LicensesResource{})
		//api.Middleware.Skip(Authorize, UsersResource{}.Create)

		//END API
		//Skip Authorize Middleware

		app.ServeFiles("/", assetsBox) // serve files from the public directory
	}

	return app
}

// translations will load locale files, set up the translator `actions.T`,
// and will return a middleware to use to load the correct locale for each
// request.
// for more information: https://gobuffalo.io/en/docs/localization
func translations() buffalo.MiddlewareFunc {
	var err error
	if T, err = i18n.New(packr.New("app:locales", "../locales"), "en-US"); err != nil {
		app.Stop(err)
	}
	return T.Middleware()
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/en/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
	return forcessl.Middleware(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}

//HTTP500 returns and message with internal server error
func HTTP500(c buffalo.Context) error {
	return c.Render(http.StatusInternalServerError, r.JSON(MessageData{Message: T.Translate(c, "message.server.error"), MessageType: "error"}))
}

//HTTP403 returns and message with errors
func HTTP403(c buffalo.Context, message string, errors ...interface{}) error {
	return c.Render(http.StatusForbidden, r.JSON(MessageData{Message: message, MessageType: "error", Errors: errors}))
}

// Error is used to responde to API Error
type Error struct {
	Message string `json:"message"`
}

// Errors is used to responde to multiple API Error
type Errors []Error

// Success is used to responde to API 200
type Success struct {
	Message string `json:"message"`
}

//MessageData is used to responde to API Requests
type MessageData struct {
	Message     string      `json:"message,omitempty"`
	MessageType string      `json:"message-type,omitempty"`
	Errors      interface{} `json:"errors,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

func initWorker(app *buffalo.App) error {
	w := app.Worker

	if w == nil {
		return errors.New("No Worker defined")
	}

	w.Register("send_email", func(args worker.Args) error {
		// do work to send an email
		return nil
	})

	w.Register("projects_github", func(args worker.Args) error {
		return nil
	})

	w.Register("projects_gitlab", func(args worker.Args) error {
		return nil
	})

	w.Register("update_projects_github", func(args worker.Args) error {
		return nil
	})

	w.Register("update_projects_gitlab", func(args worker.Args) error {
		return nil
	})

	return nil
}

func initEvents() error {
	_, err := events.Listen(func(e events.Event) {
		switch e.Kind {
		case "platform:take_projects":
			// data, err := e.Payload.Pluck("data")

			// if err == nil {
			// 	w.PerformIn(worker.Job{
			// 		Queue:   "github",
			// 		Handler: "projects_github",
			// 		Args: worker.Args{
			// 			"user_id": 123,
			// 		},
			// 	}, 5*time.Second)
			// }

		case "platform:update_projects":

		default:
			// do nothing
		}
	})
	return err
}
