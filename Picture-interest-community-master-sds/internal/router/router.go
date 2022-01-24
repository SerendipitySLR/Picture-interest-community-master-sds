package router

import (
	"github.com/gin-gonic/gin"
	"ptc/api/v1/mainpage"
	"ptc/api/v1/personalpage"
	//"ptc/pkg/log"
)

func InitRouter() *gin.Engine {
	//gin.SetMode(.AppMode)

	r := gin.Default()
	//发送模块
	routerV1MainPage := r.Group("/v1/mainPage")
	{
		routerV1MainPage.GET("/page", mainpage.ShowPage)
		routerV1MainPage.POST("/like", mainpage.AddLike)
		routerV1MainPage.GET("/comment", mainpage.ShowComment)
		routerV1MainPage.GET("/moreComment", mainpage.ShowMoreComment)
		routerV1MainPage.POST("/insertComment", mainpage.InsertComment)
		routerV1MainPage.GET("/showCollection", mainpage.ShowCollection)
		routerV1MainPage.POST("/addCollection", mainpage.AddCollection)

		//routerV1MainPage.POST("/send", mainpage.Send)
		//routerV1MainPage.POST("/follow", mainpage.Forward)
	}

	//个人主页路由模块
	routerV1PersonalPage := r.Group("/v1/personalPage")
	{
		routerV1PersonalPage.GET("/showPersonInfo", personalpage.ShowPersonalInfo)
		routerV1PersonalPage.GET("/showPersonalPost", personalpage.ShowPersonalPost)
		routerV1PersonalPage.GET("/showCollection", mainpage.ShowCollection)
		routerV1PersonalPage.GET("/showCollectionPost", personalpage.ShowCollectionPost)
		routerV1PersonalPage.POST("/newCollection", personalpage.NewCollection)
		routerV1PersonalPage.GET("/showProfile", personalpage.ShowProfile)
		routerV1PersonalPage.POST("/modifyProfile", personalpage.ModifyProfile)
	}
	// 其他模块...

	//注册zap日志框架的中间件
	return r
}
