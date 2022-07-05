package routers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "go-gin-example/m/v2/docs"

	"go-gin-example/m/v2/middleware/jwt"
	"go-gin-example/m/v2/pkg/export"
	"go-gin-example/m/v2/pkg/qrcode"
	"go-gin-example/m/v2/pkg/upload"
	"go-gin-example/m/v2/routers/api"
	v1 "go-gin-example/m/v2/routers/api/v1"

	"github.com/gin-contrib/cors"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	
	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH","GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
		  return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	  }))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.POST("/auth", api.GetAuth)
	r.GET("/total/info",api.GetTotalInfo)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		// apiv1.GET("/tags", v1.GetTags)
		// apiv1.POST("/tags", v1.AddTag)
		// apiv1.PUT("/tags/:id", v1.EditTag)
		// apiv1.DELETE("/tags/:id", v1.DeleteTag)
		// r.POST("/tags/export", v1.ExportTag)
		// r.POST("/tags/import", v1.ImportTag)
		apiv1.GET("/user/:id", v1.GetUser)
		apiv1.GET("/users", v1.GetUsers)
		apiv1.POST("/user", v1.AddUser)
		apiv1.PUT("/user/:id", v1.EditUser)
		apiv1.DELETE("/user/:id", v1.DeleteUser)

		apiv1.GET("/graphs", v1.GetGraphs)
		apiv1.GET("/graph/:id", v1.GetGraphByID)
		apiv1.POST("/graph", v1.AddGraph)
		apiv1.PUT("/graph/:id", v1.EditGraphByID)
		apiv1.DELETE("/graph/:id", v1.DeleteGraph)
		apiv1.POST("/graph/poster/generate", v1.GenerateGraphPoster)

		apiv1.GET("/cids", v1.GetCIDs)
		apiv1.GET("/cid/:id", v1.GetCIDByID)
		apiv1.POST("/cid", v1.AddCID)
	}

	return r
}
