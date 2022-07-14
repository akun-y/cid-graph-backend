package models

import (
	"github.com/jinzhu/gorm"
)

type GRAPH_DATA struct {
	Model

	TagID int      `json:"tag_id"`
	Tag   GraphTag `json:"tag"`

	Name string `json:"name"`
	Desc string `json:"desc"`

	Size        int    `json:"size"`
	Length      int    `json:"length"`
	Type        string `json:"type"`
	Version     int    `json:"version"`
	PublishTime int    `json:"publish_time"`

	Copyright     string    `json:"copyright"`
	Metas         string    `json:"metas"`
	IpfsCid       string    `json:"ipfs_cid" gorm:"index"`
	GraphAuthorID int       `json:"graph_author_id"`
	Author        GraphUser `gorm:"references:graph_author_id" json:"author"`
	//GraphAuthorID        GraphUser `gorm:"references:ID"`

	OwnerID int       `json:"owner_id"`
	Owner   GraphUser `json:"owner"`

	State int `json:"state"`
}

// ExistGraphByID checks if an CID exists based on ID
func ExistGraphByID(id int) (bool, error) {
	var record GRAPH_DATA
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&record).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if record.ID > 0 {
		return true, nil
	}

	return false, nil
}
func ExistGraphByCID(cid string) (bool, error) {
	var record GRAPH_DATA
	err := db.Select("id").Where("ipfs_cid = ? AND deleted_on = ? ", cid, 0).First(&record).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if record.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetGraphTotal gets the total number of graphs based on the constraints
func GetGraphTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&GRAPH_DATA{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetGraphs gets a list of graphs based on paging constraints
func GetGraphs(pageNum int, pageSize int, maps interface{}) ([]*GRAPH_DATA, error) {
	var graphs []*GRAPH_DATA
	err := db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&graphs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return graphs, nil
}

// GetGraph Get a single CID based on ID
func GetGraph(id int) (*GRAPH_DATA, error) {
	var record GRAPH_DATA
	err := db.Where("id = ? AND deleted_on = ? ", id, 0).First(&record).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	// err = db.Model(&CID).Related(&CID.Tag).Error
	// if err != nil && err != gorm.ErrRecordNotFound {
	// 	return nil, err
	// }

	return &record, nil
}
func GraphTotal() (int) {
	var count int
	if err := db.Model(&GRAPH_DATA{}).Count(&count).Error; err != nil {
		return -1
	}

	return count
}
// EditGraph modify a single CID
func EditGraph(id int, data interface{}) error {
	if err := db.Model(&GRAPH_DATA{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// AddGraph add a single CID
func AddGraph(data map[string]interface{}) error {
	record := GRAPH_DATA{
		Name: data["name"].(string),
		Desc: data["desc"].(string),

		Size:      data["size"].(int),
		Length:    data["length"].(int),
		Type:      data["type"].(string),
		Version:   data["version"].(int),
		Copyright: data["copyright"].(string),
		Metas:     data["metas"].(string),

		IpfsCid:       data["cid"].(string),
		//GraphAuthorID: data["graph_author"].(int),
		GraphAuthorID: data["graph_author"].(int),
		OwnerID:       data["owner"].(int),
		TagID:         data["tag_id"].(int),
	}
	if err := db.Create(&record).Error; err != nil {
		return err
	}

	return nil
}

// DeleteGraph delete a single CID
func DeleteGraph(id int) error {
	if err := db.Where("id = ?", id).Delete(GRAPH_DATA{}).Error; err != nil {
		return err
	}

	return nil
}

// CleanAllGraph clear all CID
func CleanAllGraph() error {
	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&GRAPH_DATA{}).Error; err != nil {
		return err
	}

	return nil
}
