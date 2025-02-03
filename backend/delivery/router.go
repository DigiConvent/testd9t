package router

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"

	"github.com/DigiConvent/testd9t/core/log"
	"github.com/DigiConvent/testd9t/delivery/api"
	services "github.com/DigiConvent/testd9t/pkg"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
	"github.com/gin-gonic/gin"
)

func SetupRouter(services *services.Services) {
	mainRouter := gin.Default()
	mainRouter.RedirectTrailingSlash = true

	api.RegisterRoutes(mainRouter, services)
	mainRouter.NoRoute(handleFrontend())

	if sys_domain.ProgramVersion == "dev" {
		log.Info("Development mode")
		mainRouter.Use(gin.Logger())
		mainRouter.Use(gin.Recovery())
		go func() {
			err := mainRouter.Run(":8080")
			if err != nil {
				panic("failed to start dev server: " + err.Error())
			}
		}()
	} else {
		var waitGroup sync.WaitGroup
		go func() {
			defer waitGroup.Done()
			subRouter := gin.New()
			subRouter.Use(func(ctx *gin.Context) {
				log.Info("Redirecting http to https://" + ctx.Request.Host + ctx.Request.RequestURI)
				ctx.Redirect(http.StatusMovedPermanently, "https://"+ctx.Request.Host+ctx.Request.RequestURI)
			})
			if err := subRouter.Run(":80"); err != nil {
				log.Error("Could not start http redirect server")
			}
		}()

		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			err := mainRouter.RunTLS(":443", "/home/testd9t/certs/fullchain.pem", "/home/testd9t/certs/privkey.pem")
			if err != nil {
				panic("failed to start server: " + err.Error())
			}
		}()

		waitGroup.Wait()
	}
}

func handleFrontend() gin.HandlerFunc {
	if sys_domain.ProgramVersion != "dev" {
		return func(c *gin.Context) {
			path := c.Request.URL.Path
			if path == "/favicon.ico" {
				path = "/favicon.ico"
			} else if strings.HasPrefix(path, "/assets/") {

			} else {
				path = "/index.html"
			}
			c.File("/home/testd9t/frontend/" + path)
		}
	} else {
		return proxyHandler("http://localhost:5173")
	}
}

func proxyHandler(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		remote, err := url.Parse(target)
		if err != nil {
			log.Error(err.Error())
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
			req.URL.RawQuery = c.Request.URL.RawQuery
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
