package cache_service

import (
	"strconv"
	"strings"

	"go-gin-example/m/v2/pkg/e"
)

type GRAPH struct {
	ID    int
	TagID int
	State int

	PageNum  int
	PageSize int

	Total int
}

func (a *GRAPH) GetGraphKey() string {
	return e.CACHE_GRAPH + "_" + strconv.Itoa(a.ID)
}

func (a *GRAPH) GetGraphsKey() string {
	keys := []string{
		e.CACHE_GRAPH,
		"LIST",
	}

	if a.ID > 0 {
		keys = append(keys, strconv.Itoa(a.ID))
	}
	if a.TagID > 0 {
		keys = append(keys, strconv.Itoa(a.TagID))
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
