package user_service

import (
	"encoding/json"
	"go-gin-example/m/v2/pkg/export"
	"go-gin-example/m/v2/pkg/file"
	"go-gin-example/m/v2/pkg/gredis"
	"go-gin-example/m/v2/pkg/logging"
	"go-gin-example/m/v2/service/cache_service"
	"strconv"
	"time"

	"go-gin-example/m/v2/models"

	"github.com/tealeg/xlsx"
)

type User struct {
	ID   int
	Name string

	Website string
	Github  string
	Wallet  string
	Phone   string
	Linkin  string
	Email   string
	Cids    string
	Meta    string

	State int

	PageNum  int
	PageSize int
}

func (u *User) ExistByName() (bool, error) {
	return models.ExistUserByName(u.Name)
}

func (t *User) ExistByID() (bool, error) {
	return models.ExistUserByID(t.ID)
}

func (user *User) Add() error {
	data := map[string]interface{}{
		"name": user.Name,

		"website": user.Website,
		"github":  user.Github,
		"wallet":  user.Wallet,
		"linkin":  user.Linkin,
		"email":   user.Email,
		"cids":    user.Cids,

		"state": user.State,
		"meta":  user.Meta,
	}
	return models.AddUser(data)
}

func (user *User) Edit() error {
	data := make(map[string]interface{})
	data["name"] = user.Name
	if user.State >= 0 {
		data["state"] = user.State
	}
	data["website"] = user.Website
	data["github"] = user.Github
	data["wallet"] = user.Wallet
	data["linkin"] = user.Linkin
	data["email"] = user.Email
	data["cids"] = user.Cids

	data["state"] = user.State
	data["meta"] = user.Meta

	return models.EditUser(user.ID, data)
}

func (t *User) Delete() error {
	return models.DeleteUser(t.ID)
}

func (t *User) Count() (int, error) {
	return models.GetUserTotal(t.getMaps())
}
func (a *User) Get() (*models.GraphUser, error) {
	var cacheUser *models.GraphUser

	cache := cache_service.UserCache{ID: a.ID}
	key := cache.GetUserKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheUser)
			return cacheUser, nil
		}
	}

	user, err := models.GetUser(a.ID)
	if err != nil {
		return nil, err
	}

	gredis.Set(key, user, 3600)
	return user, nil
}
func (t *User) GetAll() ([]models.GraphUser, error) {
	var (
		Users, cacheUsers []models.GraphUser
	)

	cache := cache_service.UserCache{
		Name:  t.Name,
		State: t.State,

		PageNum:  t.PageNum,
		PageSize: t.PageSize,
	}
	key := cache.GetUsersKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheUsers)
			return cacheUsers, nil
		}
	}

	Users, err := models.GetUsers(t.PageNum, t.PageSize, t.getMaps())
	if err != nil {
		return nil, err
	}

	gredis.Set(key, Users, 3600)
	return Users, nil
}

func (t *User) Export() (string, error) {
	Users, err := t.GetAll()
	if err != nil {
		return "", err
	}

	xlsFile := xlsx.NewFile()
	sheet, err := xlsFile.AddSheet("标签信息")
	if err != nil {
		return "", err
	}

	titles := []string{"ID", "名称", "创建人", "创建时间", "修改人", "修改时间"}
	row := sheet.AddRow()

	var cell *xlsx.Cell
	for _, title := range titles {
		cell = row.AddCell()
		cell.Value = title
	}

	for _, v := range Users {
		values := []string{
			strconv.Itoa(v.ID),
			v.Name,
			strconv.Itoa(v.CreatedOn),
			strconv.Itoa(v.ModifiedOn),
		}

		row = sheet.AddRow()
		for _, value := range values {
			cell = row.AddCell()
			cell.Value = value
		}
	}

	time := strconv.Itoa(int(time.Now().Unix()))
	filename := "Users-" + time + export.EXT

	dirFullPath := export.GetExcelFullPath()
	err = file.IsNotExistMkDir(dirFullPath)
	if err != nil {
		return "", err
	}

	err = xlsFile.Save(dirFullPath + filename)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func (t *User) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0

	if t.Name != "" {
		maps["name"] = t.Name
	}
	if t.State >= 0 {
		maps["state"] = t.State
	}

	return maps
}
