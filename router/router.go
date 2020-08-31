package router

import (
	"go-sharp/app"
	"go-sharp/constant"
	"go-sharp/controller/user"
	"go-sharp/docs"
	"go-sharp/ioc"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"sync"
)

type IGinRouter interface {
	InitRouter(businessManagerFactoryInjector func () gin.HandlerFunc) *gin.Engine
}

type router struct{

}

func (r *router) InitRouter(businessManagerFactoryInjector func () gin.HandlerFunc) *gin.Engine {

	engine := gin.New()
	engine.Use(businessManagerFactoryInjector())
	if app.Config.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}



	v1 := engine.Group("/api/v1")
	{
		users := v1.Group("/user")
		{
			users.GET("get/:idUser", user.Get)
			users.GET("get-all/", user.GetAll)
			users.POST("create/", user.Create)
			users.PUT("update/", user.Update)
			users.DELETE("delete/:idUser", user.Delete)
		}
	}
	docs.SwaggerInfo.Host = "localhost" + app.Config.Port
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	return engine

}

var (
	m          *router
	routerOnce sync.Once
)

func GinRouter() IGinRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}

func InjectBusinessManagerFactory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		businessManagerFactory:=ioc.ServiceContainer().GetBusinessManagerFactory(ctx)

		ctx.Set(constant.Factory,businessManagerFactory)
		// before request

		ctx.Next()

		var err error
		err = nil
		if status := ctx.Writer.Status(); status ==200{
			err=businessManagerFactory.Uow.CommitTransaction()
		}else{
			err=businessManagerFactory.Uow.RollbackTransaction()
		}

		if err != nil{
			ctx.JSON(500,err.Error())
		}
	}
}