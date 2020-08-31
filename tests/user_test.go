package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"go-sharp/app"
	"go-sharp/dal/db"
	"go-sharp/ent/user"
	"go-sharp/model"
	"go-sharp/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var dbHandler db.IDbHandler

func initTest(){
	app.InitConfig()
	dbHandler = &TestDbHandler{}
	dbHandler.Migrate()
}

func getTestEngine() *gin.Engine {
	return router.GinRouter().InitRouter(InjectBusinessManagerTestFactory)
}

func TestUserGetAll(t *testing.T){
	initTest()
	engine :=getTestEngine()
	w := performRequest(engine, "GET", "/api/v1/user/get-all/",nil)
	assert.Equal(t, http.StatusOK, w.Code, "Response code is not Ok")
	decoder := json.NewDecoder(w.Body)
	userList := make([]model.User,0)
	err := decoder.Decode(&userList)
	assert.Nil(t,err,"Error while parsing response to user list")
	assert.True(t, len(userList)>0,"User list does not contains any record")
}

func TestUserGet(t *testing.T){
	initTest()
	engine :=getTestEngine()
	w := performRequest(engine, "GET", "/api/v1/user/get/1",nil)
	assert.Equal(t, http.StatusOK, w.Code, "Response code is not Ok")
	decoder := json.NewDecoder(w.Body)
	user := model.User{}
	err := decoder.Decode(&user)
	assert.Nil(t,err,"Error while parsing response to user")
	assert.ObjectsAreEqualValues("Ali",user.Name)
}

func TestUserCreate(t *testing.T)  {
	initTest()
	engine:=getTestEngine()
	usr := &model.User{}
	usr.Name ="MSK"
	usr.Age =32
	b, _ := json.Marshal(usr)
	fmt.Println(string(b))
	w := performRequest(engine, "POST", "/api/v1/user/create/",bytes.NewBuffer(b))
	assert.Equal(t, http.StatusOK, w.Code, "Response code is not Ok")

	ctx:=context.Background()
	client := dbHandler.GetDbClient()
	users,err:=client.User.Query().Where(user.NameEQ("MSK")).All(ctx)
	assert.Nil(t,err,"Error while getting users from db")
	assert.True(t, len(users)>0,"No user found")
}

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}