package actions

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	csrf "github.com/gobuffalo/mw-csrf"
	forcessl "github.com/gobuffalo/mw-forcessl"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
	tokenauth "github.com/gobuffalo/mw-tokenauth"
	"github.com/gorilla/sessions"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"github.com/unrolled/secure"

	"management/models"

	"github.com/gobuffalo/buffalo-pop/pop/popmw"
	i18n "github.com/gobuffalo/mw-i18n"
	"github.com/gobuffalo/packr/v2"

	gwa "github.com/gobuffalo/gocraft-work-adapter"
	redisqueue "github.com/gomodule/redigo/redis"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var T *i18n.Translator

const (
	currentUserID = "current_user_id"
)

var (
	errNoTransaction = errors.New("No transaction")
)

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

		var (
			SessionName = "_api_session"

			store sessions.Store
		)

		var corsConfig []buffalo.PreWare

		if !IsProduction() {
			corsConfig = []buffalo.PreWare{
				cors.AllowAll().Handler,
			}
		}
		// Addr:     fmt.Sprintf("%s:%s", envy.Get("REDIS_SERVER", "localhost"), envy.Get("REDIS_PORT", "6309")),
		// Password: envy.Get("REDIS_PASSWORD", ""),
		app = buffalo.New(buffalo.Options{
			Env:          ENV,
			SessionName:  SessionName,
			SessionStore: store,
			PreWares:     corsConfig,
			Worker: gwa.New(gwa.Options{
				Pool: &redisqueue.Pool{
					MaxActive: 5,
					MaxIdle:   5,
					Wait:      true,
					Dial: func() (redisqueue.Conn, error) {
						return redisqueue.Dial("tcp", fmt.Sprintf("%s:%s", envy.Get("REDIS_SERVER", "localhost"), envy.Get("REDIS_PORT", "6379")), redisqueue.DialPassword(envy.Get("REDIS_PASSWORD", "")))
					},
				},
				Name:           "Worker",
				MaxConcurrency: 25,
				Logger:         buffalo.NewOptions().Logger,
			}),
		})

		if !IsProduction() {
			app.Use(paramlogger.ParameterLogger)
		}

		// Automatically redirect to SSL
		app.Use(forceSSL())

		if !IsProduction() {
			// Log request parameters (filters apply).
			app.Use(paramlogger.ParameterLogger)
		}

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		app.Use(csrf.New)

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.Connection)
		// Remove to disable this.
		app.Use(popmw.Transaction(models.DB))

		// Setup and use translations:
		app.Use(translations())

		api := app.Group("/api")

		//Middlewares
		TokenAuth := tokenauth.New(tokenauth.Options{
			SignMethod: jwt.SigningMethodHS256,
			GetKey: func(_ jwt.SigningMethod) (interface{}, error) {
				return secretKey, nil
			},
		})

		// api.Use(SetCurrentUser)
		// api.Use(Authorize)
		api.Use(TokenAuth)

		//Resources
		userResource := UsersResource{&buffalo.BaseResource{}}

		// api.Middleware.Skip(Authorize, AuthCreate, userResource.Create)

		api.Middleware.Skip(TokenAuth, AuthCreate, userResource.Create)

		//PLATFORMS
		platform := Platform{}

		app.GET("/github/callback", platform.GithubCallback)

		app.GET("/user/confirm", userResource.Confirm)

		api.POST("/login", AuthCreate)

		api.DELETE("/logout", AuthDestroy)

		api.GET("/refresh", AuthRefresh)

		api.Resource("/users", userResource)

		app.GET("/", HomeHandler)

		api.Resource("/projects", ProjectsResource{})

		api.Resource("/comments", CommentsResource{})

		api.Resource("/tasks", TasksResource{})

		api.Resource("/licenses", LicensesResource{})

		app.GET("/{path:.+}", VueHandler)

		go func() {
			app.ServeFiles("/", assetsBox) // serve files from the public directory
		}()

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
		SSLRedirect:     IsProduction(),
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}

func IsProduction() bool {
	return ENV == "production"
}
