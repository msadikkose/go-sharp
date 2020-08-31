package user

import (
	"go-sharp/business/manager"
	"go-sharp/constant"
	"go-sharp/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strconv"
)


// GetUserById godoc
// @Summary Gets user
// @Description Gets user by id
// @Tags User
// @Produce json
// @Param  idUser path string true "used for user id"
// @Success 200 {object} model.User
// @Failure 400 {object} string
// @Failure 404 {string} string
// @Router /user/get/{idUser} [get]
func Get(ctx *gin.Context) {

	idUser,err:= strconv.Atoi( ctx.Param("idUser"))
	if err != nil{
		ctx.JSON(400,"Bad request")
	}
	businessManagerFactory:=ctx.MustGet(constant.Factory).(*manager.Factory)
	usr,err:=businessManagerFactory.UserManager.GetById(idUser)
	if err != nil{
		ctx.JSON(500,"Internal Server Error"+err.Error())
		return
	}

	ctx.JSON(200,usr)
}

// GetUsers godoc
// @Summary Gets all users
// @Description Gets all users
// @Tags User
// @Produce json
// @Success 200 {object} []model.User
// @Failure 500 {object} string
// @Router /user/get-all/ [get]
func GetAll(ctx *gin.Context) {
	businessManagerFactory:=ctx.MustGet(constant.Factory).(*manager.Factory)
	result,err:=businessManagerFactory.UserManager.GetAll()
	if err != nil{
		ctx.JSON(500,"Internal Server Error"+err.Error())
		return
	}

	ctx.JSON(200, result)
}


// GetUsers godoc
// @Summary Creates new user
// @Description Creates new user
// @Tags User
// @Accept  json
// @Produce json
// @Param  Request body model.User true "user"
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /user/create/ [post]
func Create(ctx *gin.Context) {
	var userModel model.User
	err:=ctx.MustBindWith(&userModel,binding.JSON)
	if err != nil{
		ctx.JSON(400,"Bad request")
	}
	businessManagerFactory:=ctx.MustGet(constant.Factory).(*manager.Factory)
	err=businessManagerFactory.UserManager.Create(&userModel)
	if err != nil{
		ctx.JSON(500,"Internal Server Error"+err.Error())
		return
	}

	ctx.JSON(200, "OK")
}


// GetUsers godoc
// @Summary Updates existing user
// @Description Updates existing user
// @Tags User
// @Accept  json
// @Produce json
// @Param  Request body model.User true "user"
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /user/update/ [put]
func Update(ctx *gin.Context) {
	var userModel model.User
	err:=ctx.MustBindWith(&userModel,binding.JSON)
	if err != nil{
		ctx.JSON(400,"Bad request")
	}
	businessManagerFactory:=ctx.MustGet(constant.Factory).(*manager.Factory)
	err=businessManagerFactory.UserManager.Update(&userModel)
	if err != nil{
		ctx.JSON(500,"Internal Server Error"+err.Error())
		return
	}

	ctx.JSON(200, "OK")
}

// GetUserById godoc
// @Summary Deletes user
// @Description Deletes user by id
// @Tags User
// @Produce json
// @Param  idUser path string true "used for user id"
// @Success 200 {string} string
// @Failure 400 {object} string
// @Failure 404 {string} string
// @Router /user/delete/{idUser} [delete]
func Delete(ctx *gin.Context) {
	idUser,err:= strconv.Atoi( ctx.Param("idUser"))
	if err != nil{
		ctx.JSON(400,"Bad request")
	}
	businessManagerFactory:=ctx.MustGet(constant.Factory).(*manager.Factory)
	err=businessManagerFactory.UserManager.Delete(idUser)
	if err != nil{
		ctx.JSON(500,"Internal Server Error"+err.Error())
		return
	}
	ctx.JSON(200, "OK")
}