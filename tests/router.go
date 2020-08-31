package tests

import (
	"go-sharp/constant"
	"github.com/gin-gonic/gin"
)

func InjectBusinessManagerTestFactory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		businessManagerFactory:=TestServiceContainer().GetBusinessManagerFactory(ctx)

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
		// access the status we are sending

	}
}
