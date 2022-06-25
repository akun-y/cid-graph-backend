package api

import (
	"go-gin-example/m/v2/pkg/app"
	"go-gin-example/m/v2/pkg/e"
	"go-gin-example/m/v2/pkg/util"
	"go-gin-example/m/v2/service/user_service"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type AUTH struct {
	Username string  `valid:"Required; MaxSize(50)"`
	Password string  `valid:"Required; MaxSize(50)"`
	Kind     string  `json:"kind,omitempty"`
	Name     *string `json:"name"`
}

// type AuthForm struct {
// 	username      string `json:"username" valid:"Required;MaxSize(100)"`
// 	password      string `json:"password" valid:"Required;MaxSize(255)"`
// }

// @Summary Post Auth
// @Accept json
// @Produce  json
// @Param json body AUTH true "Auth Get Token"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [post]
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	user := AUTH{}
	c.BindJSON(&user)

	valid := validation.Validation{}

	username := user.Username
	password := user.Password
	a := AUTH{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := user_service.Auth{Username: username, Password: password}
	isExist, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}
