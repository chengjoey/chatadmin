package router

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"runtime/debug"

	"github.com/chengjoey/chatadmin/app/api"
	"github.com/chengjoey/chatadmin/controller"
	"github.com/chengjoey/chatadmin/pkg/response"
	"github.com/gin-gonic/gin"
)

var (
// //go:embed template
// embededFiles embed.FS
// //go:embed ./template/js
// jsFiles embed.FS
// //go:embed ./template/css
// cssFiles embed.FS
// //go:embed ./template/static
// staticFiles embed.FS
)

func JSONAppErrorReporter() gin.HandlerFunc {
	return jsonAppErrorMiddleware(gin.ErrorTypeAny)
}

func jsonAppErrorMiddleware(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detedtedErrors := c.Errors.ByType(errType)
		if len(detedtedErrors) > 0 {
			err := detedtedErrors[0].Err
			var parsedError response.AppError
			switch err.(type) {
			case response.AppError:
				parsedError = err.(response.AppError)
			default:
				parsedError = response.AppError{
					Inner:      err,
					Msg:        err.Error(),
					StackTrace: string(debug.Stack()),
				}
			}
			c.IndentedJSON(http.StatusBadRequest, parsedError)
			return
		}
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Token")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func RegisterHandle(r *gin.Engine, embededFiles embed.FS) {
	go controller.Broadcaster.Start()

	r.Use(CORSMiddleware())
	r.Use(JSONAppErrorReporter())
	templ := template.Must(template.New("").ParseFS(embededFiles, "template/*.html"))
	r.SetHTMLTemplate(templ)
	jsFiles, err := fs.Sub(embededFiles, "template/js")
	if err != nil {
		panic(err)
	}
	cssFiles, err := fs.Sub(embededFiles, "template/css")
	if err != nil {
		panic(err)
	}
	staticFiles, err := fs.Sub(embededFiles, "template/static")
	if err != nil {
		panic(err)
	}
	r.StaticFS("/js", http.FS(jsFiles))
	r.StaticFS("/css", http.FS(cssFiles))
	r.StaticFS("/static", http.FS(staticFiles))
	r.GET("favicon.ico", func(c *gin.Context) {
		file, _ := embededFiles.ReadFile("template/favicon.ico")
		c.Data(
			http.StatusOK,
			"image/x-icon",
			file,
		)
	})
	// r.LoadHTMLGlob(global.RootDir + "/template/*.html")
	// r.Static("/js", global.RootDir+"/template/js")
	// r.Static("/css", global.RootDir+"/template/css")
	// r.Static("/static", global.RootDir+"/template/static")
	r.GET("/", api.HomeHandle)
	r.POST("/login", api.Login)

	gAuth := r.Group("/v1", api.AuthRequired())
	gAuth.GET("/chat", api.WebSocketHandle)
	gAuth.GET("/stream", api.StreamSSEvent)
	gAuth.GET("/users", api.AdminGetUserList)
	gAuth.GET("/msgs", api.GetHistoryMsg)
}
