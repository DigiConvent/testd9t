package router

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/DigiConvent/testd9t/core/log"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
	"github.com/gin-gonic/gin"
)

func serveFrontend(router *gin.Engine) {
	if sys_domain.ProgramVersion != "dev" {
		router.Static("/assets", "/home/testd9t/frontend/assets")
		router.StaticFile("/favicon.ico", "/home/testd9t/frontend/favicon.ico")
		router.StaticFile("/index.html", "/home/testd9t/frontend/index.html")
		router.StaticFile("/", "/home/testd9t/frontend/index.html")
	} else {
		router.Use(proxyHandler("http://localhost:5173"))
	}
}

func proxyHandler(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path)
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
