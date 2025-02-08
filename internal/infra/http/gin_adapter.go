package httpserver

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "tech-challenge-hackaton/docs"
)

type GinHTTPServerAdapter struct {
	Engine   *gin.Engine
	basePath string
}

func NewGinHTTPServerAdapter() *GinHTTPServerAdapter {
	httpServer := &GinHTTPServerAdapter{
		Engine: gin.Default(),
	}

	httpServer.Engine.SetTrustedProxies(nil)
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"*"}
	config.AllowHeaders = []string{"Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"}
	httpServer.Engine.Use(cors.New(config))
	return httpServer
}

func (g *GinHTTPServerAdapter) SetBasePath(basePath string) {
	g.basePath = basePath
}

func (g *GinHTTPServerAdapter) Run(adds ...string) error {
	return g.Engine.Run(adds...)
}

func (g *GinHTTPServerAdapter) SetSwagger(path string) {
	g.Engine.GET(g.basePath+path, ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func createHandlerFuncs(callbacks ...HTTPHandlerFunc) []gin.HandlerFunc {
	var handlers []gin.HandlerFunc
	for _, c := range callbacks {
		handlers = append(handlers, func(ctx *gin.Context){
			c(ctx)
		})
	}
	return handlers
}

func (g *GinHTTPServerAdapter) GET(path string, callbacks ...HTTPHandlerFunc) {
	g.Engine.GET(
		g.basePath+path,
		createHandlerFuncs(callbacks...)...,
	)
}

func (g *GinHTTPServerAdapter) POST(path string, callbacks ...HTTPHandlerFunc) {
	g.Engine.POST(
		g.basePath+path,
		createHandlerFuncs(callbacks...)...,
	)
}

func (g *GinHTTPServerAdapter) PUT(path string, callbacks ...HTTPHandlerFunc) {
	g.Engine.PUT(
		g.basePath+path,
		createHandlerFuncs(callbacks...)...,
	)
}

func (g *GinHTTPServerAdapter) PATCH(path string, callbacks ...HTTPHandlerFunc) {
	g.Engine.PATCH(
		g.basePath+path,
		createHandlerFuncs(callbacks...)...,
	)
}

func (g *GinHTTPServerAdapter) DELETE(path string, callbacks ...HTTPHandlerFunc) {
	g.Engine.DELETE(
		g.basePath+path,
		createHandlerFuncs(callbacks...)...,
	)
}

func (g *GinHTTPServerAdapter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	g.Engine.ServeHTTP(w, req)

	
}

