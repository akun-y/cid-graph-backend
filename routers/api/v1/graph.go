package v1

import (
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"go-gin-example/m/v2/pkg/app"
	"go-gin-example/m/v2/pkg/e"
	"go-gin-example/m/v2/pkg/qrcode"
	"go-gin-example/m/v2/pkg/setting"
	"go-gin-example/m/v2/pkg/util"
	"go-gin-example/m/v2/service/graph_service"
)

// @Summary Get a single graph
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/graph/{id} [get]
func GetGraphByID(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	cidService := graph_service.GRAPH{ID: id}
	exists, err := cidService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_GRAPH_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_GRAPH, nil)
		return
	}

	CID, err := cidService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_GRAPH_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, CID)
}

// @Summary Get multiple graphs
// @Produce  json
// @Param token header string true "token"
// @Param tag_id query int false "TagID"
// @Param state query int false "State"
// @Param page query int false "Page"
// @Param created_by query int false "CreatedBy"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/graphs [get]
func GetGraphs(c *gin.Context) {
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
	page := com.StrTo(c.Query("page")).MustInt()
	service := graph_service.GRAPH{
		TagID:    tagId,
		State:    state,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

	total, err := service.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_GRAPH_FAIL, nil)
		return
	}

	ipfsCids, err := service.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_GRAPHS_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = ipfsCids
	data["total"] = total
	data["page"] = page
	data["page_size"] = service.PageSize
	data["tag_id"] = service.TagID
	data["state"] = service.State

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// {
// 	"name": "",
// 	"desc": "A green door",
// 	"size": 1,
// 	"length": 1,
// 	"type": "",
// 	"version": 1,
// 	"createTime": 1,
// 	"publishedTime": 1,
// 	"copyright": "",
// 	"metas": ["", ""],
// 	"owner": {
// 		"website": "",
// 		"github": "",
// 		"address": "",
// 		"linkin": "",
// 		"email": "",
// 		"cids": ["", "", ""]
// 	},
// 	"Author": {
// 		"website": "",
// 		"github": "",
// 		"address": "",
// 		"linkin": "",
// 		"email": "",
// 		"cids": ["", "", ""]
// 	}
// }
type AddGraphForm struct {
	Name        string `form:"name" valid:"Required;MaxSize(100)"`
	Desc        string `form:"desc" valid:"Required;MaxSize(255)"`
	Size        int    `form:"size"`
	Length      int    `form:"length"`
	Type        string `form:"type"`
	Version     int    `form:"version"`
	Copyright   string `form:"copyright"`
	Metas       string `form:"metas" valid:"MaxSize(65535)"`
	CID         string `form:"cid" valid:"Required;MaxSize(255)"`
	GraphAuthor int    `form:"graphAuthor" valid:"Min(0)"`
	Owner       int    `form:"owner"  valid:"Min(0)"`
	TagID       int    `form:"tag_id" valid:"Min(0)"`
	State       int    `form:"state" valid:"Range(0,1)"`
}

// @Summary Add graph
// @Accept json
// @Produce  json
// @Param token header string true "token"
// @Param body body AddGraphForm true "AddGraphForm"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/graph [post]
func AddGraph(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddGraphForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	// tagService := tag_service.Tag{ID: form.TagID}
	// exists, err := tagService.ExistByID()
	// if err != nil {
	// 	appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
	// 	return
	// }

	service := graph_service.GRAPH{
		TagID: form.TagID,
		Name:  form.Name,
		Desc:  form.Desc,

		Size:      form.Size,
		Length:    form.Size,
		Type:      form.Type,
		Version:   form.Version,
		Copyright: form.Copyright,
		Metas:     form.Metas,

		CID:         form.CID,
		State:       form.State,
		GraphAuthor: form.GraphAuthor,
		Owner:       form.Owner,
	}

	exists, err := service.ExistByCID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_GRAPH_FAIL, nil)
		return
	}
	if exists {
		log.Println("cid is exist:", form.CID)
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_GRAPH, nil)
		return
	}

	if err := service.Add(); err != nil {
		log.Println("cidService.Add error:", err)
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_GRAPH_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type EditGraphForm struct {
	AddGraphForm
	ID int `form:"id" valid:"Required;Min(0)"`
	// TagID  int    `form:"tag_id" valid:"Required;Min(1)"`
	// Name   string `form:"name" valid:"Required;MaxSize(100)"`
	// Desc   string `form:"desc" valid:"Required;MaxSize(255)"`
	// CID    string `form:"cid" valid:"Required;MaxSize(65535)"`
	// Author string `form:"author" valid:"Required;MaxSize(100)"`
	// Owner  string `form:"owner" valid:"Required;MaxSize(255)"`
	// State  int    `form:"state" valid:"Range(0,1)"`
}

// @Summary Update graph
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "ID"
// @Param body body AddGraphForm true "AddGraphForm"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/graph/{id} [put]
func EditGraphByID(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = EditGraphForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	cidService := graph_service.GRAPH{
		ID:    form.ID,
		TagID: form.TagID,
		Name:  form.Name,
		Desc:  form.Desc,
		//CID:   form.CID,
		State: form.State,

		Size:      form.Size,
		Length:    form.Size,
		Type:      form.Type,
		Version:   form.Version,
		Copyright: form.Copyright,
		Metas:     form.Metas,

		//GraphAuthor: form.GraphAuthor,
		//Owner:       form.Owner,
	}
	exists, err := cidService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_GRAPH_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_GRAPH, nil)
		return
	}

	// tagService := tag_service.Tag{ID: form.TagID}
	// exists, err = tagService.ExistByID()
	// if err != nil {
	// 	appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
	// 	return
	// }

	// if !exists {
	// 	appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
	// 	return
	// }

	err = cidService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_GRAPH_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary Delete graph
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/graph/{id} [delete]
func DeleteGraph(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	service := graph_service.GRAPH{ID: id}
	exists, err := service.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_GRAPH_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_GRAPH, nil)
		return
	}

	err = service.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_GRAPH_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

const (
	QRCODE_URL = "https://github.com/akun-y/blog/blob/master/github-qr.png"
)

func GenerateGraphPoster(c *gin.Context) {
	appG := app.Gin{C: c}
	CID := &graph_service.GRAPH{}
	qr := qrcode.NewQrCode(QRCODE_URL, 300, 300, qr.M, qr.Auto)
	posterName := graph_service.GetPosterFlag() + "-" + qrcode.GetQrCodeFileName(qr.URL) + qr.GetQrCodeExt()
	graphPoster := graph_service.NewGraphPoster(posterName, CID, qr)
	graphPosterBgService := graph_service.NewGraphPosterBg(
		"bg.jpg",
		graphPoster,
		&graph_service.Rect{
			X0: 0,
			Y0: 0,
			X1: 550,
			Y1: 700,
		},
		&graph_service.Pt{
			X: 125,
			Y: 298,
		},
	)

	_, filePath, err := graphPosterBgService.Generate()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GEN_GRAPH_POSTER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"poster_url":      qrcode.GetQrCodeFullUrl(posterName),
		"poster_save_url": filePath + posterName,
	})
}
