package cache_service

import (
	"strconv"
	"strings"

	"go-gin-example/m/v2/pkg/e"
)

type CID struct {
	ID    int
	State int
	Total int

	PageNum  int
	PageSize int
}

func (a *CID) GetCIDKey() string {
	return e.CACHE_CID + "_" + strconv.Itoa(a.ID)
}

func (a *CID) GetCIDsKey() string {
	keys := []string{
		e.CACHE_CID,
		"LIST",
	}

	if a.ID > 0 {
		keys = append(keys, strconv.Itoa(a.ID))
	}
	if a.State >= 0 {
		keys = append(keys, strconv.Itoa(a.State))
	}
	if a.PageNum > 0 {
		keys = append(keys, strconv.Itoa(a.PageNum))
	}
	if a.PageSize > 0 {
		keys = append(keys, strconv.Itoa(a.PageSize))
	}
	if a.Total > 0 {
		keys = append(keys, strconv.Itoa(a.Total))
	}
	return strings.Join(keys, "_")
}
