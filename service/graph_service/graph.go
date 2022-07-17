package graph_service

import (
	"encoding/json"

	"go-gin-example/m/v2/models"
	"go-gin-example/m/v2/pkg/gredis"
	"go-gin-example/m/v2/pkg/logging"
	"go-gin-example/m/v2/service/cache_service"
)

//type GRAPH = graphtype.GRAPH

// type GRAPH struct {
// 	graphtype.GRAPH
// }
type GRAPH struct {
	ID    int
	TagID int
	Name  string
	Desc  string

	Size      int
	Length    int
	Type      string
	Version   int
	Copyright string
	Metas     string

	CID         string
	GraphAuthor int
	Owner       int

	PageNum  int
	PageSize int
	State    int
}

func (a *GRAPH) Add() error {
	graph := map[string]interface{}{
		"tag_id": a.TagID,
		"name":   a.Name,
		"desc":   a.Desc,

		"size":      a.Size,
		"length":    a.Length,
		"type":      a.Type,
		"version":   a.Version,
		"copyright": a.Copyright,
		"metas":     a.Metas,

		"cid":          a.CID,
		"graph_author": a.GraphAuthor,
		"owner":        a.Owner,
	}

	if err := models.AddGraph(graph); err != nil {
		return err
	}

	return nil
}

func (a *GRAPH) Edit() error {
	cache := cache_service.GRAPH{ID: a.ID}
	key := cache.GetGraphKey()
	gredis.Delete(key)

	return models.EditGraph(a.ID, map[string]interface{}{
		"tag_id": a.TagID,
		"name":   a.Name,
		"desc":   a.Desc,

		"size":      a.Size,
		"length":    a.Length,
		"type":      a.Type,
		"version":   a.Version,
		"copyright": a.Copyright,
		"metas":     a.Metas,

		//"cid":    a.CID,
		//"author": a.GraphAuthor,
		//"owner":  a.Owner,
	})
}

func (a *GRAPH) Get() (*models.GRAPH_DATA, error) {
	var cacheGraph *models.GRAPH_DATA

	cache := cache_service.GRAPH{ID: a.ID}
	key := cache.GetGraphKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheGraph)
			return cacheGraph, nil
		}
	}

	graph, err := models.GetGraph(a.ID)
	if err != nil {
		return nil, err
	}

	gredis.Set(key, graph, 3600)
	return graph, nil
}

func (a *GRAPH) GetAll() ([]*models.GRAPH_DATA, error) {
	var (
		graphs, cacheGraphs []*models.GRAPH_DATA
	)
	total := models.GraphTotal()
	cache := cache_service.GRAPH{
		TagID: a.TagID,
		State: a.State,
		Total: total,

		PageNum:  a.PageNum,
		PageSize: a.PageSize,
	}
	key := cache.GetGraphsKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheGraphs)
			return cacheGraphs, nil
		}
	}

	graphs, err := models.GetGraphs(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}

	gredis.Set(key, graphs, 3600)
	return graphs, nil
}

func (a *GRAPH) Delete() error {
	return models.DeleteGraph(a.ID)
}

func (a *GRAPH) ExistByID() (bool, error) {
	return models.ExistGraphByID(a.ID)
}
func (a *GRAPH) ExistByCID() (bool, error) {
	return models.ExistGraphByCID(a.CID)
}
func (a *GRAPH) Count() (int, error) {
	return models.GetGraphTotal(a.getMaps())
}
func (a *GRAPH) Total() (int) {
	return models.GraphTotal()
}
func (a *GRAPH) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	// if a.State != -1 {
	// 	maps["state"] = a.State
	// }
	if a.TagID != -1 {
		maps["tag_id"] = a.TagID
	}

	return maps
}
