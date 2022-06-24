package cache_service

import (
	"strconv"
	"strings"

	"go-gin-example/m/v2/pkg/e"
)

type UserCache struct {
	ID       int
	Name     string
	State    int
	PageNum  int
	PageSize int
}

func (a *UserCache) GetUserKey() string {
	return e.CACHE_USER + "_" + strconv.Itoa(a.ID)
}
func (t *UserCache) GetUsersKey() string {
	keys := []string{
		e.CACHE_USER,
		"LIST",
	}

	if t.Name != "" {
		keys = append(keys, t.Name)
	}
	if t.State >= 0 {
		keys = append(keys, strconv.Itoa(t.State))
	}
	if t.PageNum > 0 {
		keys = append(keys, strconv.Itoa(t.PageNum))
	}
	if t.PageSize > 0 {
		keys = append(keys, strconv.Itoa(t.PageSize))
	}

	return strings.Join(keys, "_")
}
