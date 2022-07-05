package v1

import (
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"go-gin-example/m/v2/pkg/app"
	"go-gin-example/m/v2/pkg/e"
	"go-gin-example/m/v2/pkg/setting"
	"go-gin-example/m/v2/pkg/util"
	"go-gin-example/m/v2/service/cid_service"
)

// @Summary Get a single cid
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/cid/{id} [get]
func GetCIDByID(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	cidService := cid_service.CID{ID: id}
	exists, err := cidService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_CID_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_CID, nil)
		return
	}

	CID, err := cidService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_CID_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, CID)
}

// @Summary Get multiple cids
// @Produce  json
// @Param token header string true "token"
// @Param tag_id query int false "TagID"
// @Param state query int false "State"
// @Param page query int false "Page"
// @Param created_by query int false "CreatedBy"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/cids [get]
func GetCIDs(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state")
	}

	tagId := -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		valid.Min(tagId, 1, "tag_id")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	service := cid_service.CID{
		TagID:    tagId,
		State:    state,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

	total, err := service.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_CID_FAIL, nil)
		return
	}

	ipfsCids, err := service.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_CIDS_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = ipfsCids
	data["total"] = total

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

type AddCIDForm struct {
	Name        string `form:"name" valid:"Required;MaxSize(100)"`
	Desc        string `form:"desc" valid:"Required;MaxSize(255)"`
	Size        int    `form:"size"`
	Length      int    `form:"length"`
	Type        string `form:"type"`
	Version     int    `form:"version"`
	CID         string `form:"cid" valid:"Required;MaxSize(255)"`
	State       int    `form:"state" valid:"Range(0,1)"`
}

// @Summary Add CID
// @Accept json
// @Produce  json
// @Param token header string true "token"
// @Param body body AddCIDForm true "AddCIDForm"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/cid [post]
func AddCID(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddCIDForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	service := cid_service.CID{
		Name:  form.Name,
		Desc:  form.Desc,

		Size:      form.Size,
		Length:    form.Size,
		Type:      form.Type,
		Version:   form.Version,

		CID:         form.CID,
		State:       form.State,
	}

	exists, err := service.ExistByCID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_CID_FAIL, nil)
		return
	}
	if exists {
		log.Println("cid is exist:", form.CID)
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_CID, nil)
		return
	}

	if err := service.Add(); err != nil {
		log.Println("cidService.Add error:", err)
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_CID_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
