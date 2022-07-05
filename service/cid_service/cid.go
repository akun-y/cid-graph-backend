package cid_service

import (
	"encoding/json"

	"go-gin-example/m/v2/models"
	"go-gin-example/m/v2/pkg/gredis"
	"go-gin-example/m/v2/pkg/logging"
	"go-gin-example/m/v2/service/cache_service"
)

type CID struct {
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
	CIDAuthor int
	Owner       int

	PageNum  int
	PageSize int
	State    int
}

func (a *CID) Add() error {
	CID := map[string]interface{}{
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
		"author": a.CIDAuthor,
		"owner":        a.Owner,
	}

	if err := models.AddCID(CID); err != nil {
		return err
	}

	return nil
}

func (a *CID) Get() (*models.GRAPH_cid, error) {
	var cacheCID *models.GRAPH_cid

	cache := cache_service.CID{ID: a.ID}
	key := cache.GetCIDKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheCID)
			return cacheCID, nil
		}
	}

	CID, err := models.GetCID(a.ID)
	if err != nil {
		return nil, err
	}

	gredis.Set(key, CID, 3600)
	return CID, nil
}

func (a *CID) GetAll() ([]*models.GRAPH_cid, error) {
	var (
		CIDs, cacheCIDs []*models.GRAPH_cid
	)

	cache := cache_service.CID{
		TagID: a.TagID,
		//State: a.State,

		PageNum:  a.PageNum,
		PageSize: a.PageSize,
	}
	key := cache.GetCIDsKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheCIDs)
			return cacheCIDs, nil
		}
	}

	CIDs, err := models.GetCIDs(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}

	gredis.Set(key, CIDs, 3600)
	return CIDs, nil
}

func (a *CID) Delete() error {
	return models.DeleteCID(a.ID)
}

func (a *CID) ExistByID() (bool, error) {
	return models.ExistCIDByID(a.ID)
}
func (a *CID) ExistByCID() (bool, error) {
	return models.ExistCID(a.CID)
}
func (a *CID) Count() (int, error) {
	return models.GetCIDTotal(a.getMaps())
}

func (a *CID) getMaps() map[string]interface{} {
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
