package models

import (
	"github.com/jinzhu/gorm"
)

type GRAPH_cid struct {
	Model
	CID         string `json:"cid" gorm:"column:ipfs_cid"`
	Size        int    `json:"size"`
	Length      int    `json:"length"`
	Type        string `json:"type"`
	Version     int    `json:"version"`
	PublishTime int    `json:"publish_time"`
}

// ExistCIDByID checks if an CID exists based on ID
func ExistCIDByID(id int) (bool, error) {
	var record GRAPH_cid
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&record).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if record.ID > 0 {
		return true, nil
	}

	return false, nil
}
func ExistCID(cid string) (bool, error) {
	var record GRAPH_cid
	err := db.Select("id").Where("ipfs_cid = ? AND deleted_on = ? ", cid, 0).First(&record).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if record.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetCIDTotal gets the total number of CIDs based on the constraints
func GetCIDTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&GRAPH_cid{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
func CIDTotal() (int) {
	var count int
	if err := db.Model(&GRAPH_cid{}).Count(&count).Error; err != nil {
		return -1
	}

	return count
}

// GetCIDs gets a list of CIDs based on paging constraints
func GetCIDs(pageNum int, pageSize int, maps interface{}) ([]*GRAPH_cid, error) {
	var CIDs []*GRAPH_cid
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&CIDs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return CIDs, nil
}

// GetCID Get a single CID based on ID
func GetCID(id int) (*GRAPH_cid, error) {
	var record GRAPH_cid
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

// AddCID add a single CID
func AddCID(data map[string]interface{}) error {
	record := GRAPH_cid{
		CID:     data["cid"].(string),
		Size:    data["size"].(int),
		Length:  data["length"].(int),
		Type:    data["type"].(string),
		Version: data["version"].(int),
	}
	if err := db.Create(&record).Error; err != nil {
		return err
	}

	return nil
}

// DeleteCID delete a single CID
func DeleteCID(id int) error {
	if err := db.Where("id = ?", id).Delete(GRAPH_cid{}).Error; err != nil {
		return err
	}

	return nil
}

// CleanAllCID clear all CID
func CleanAllCID() error {
	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&GRAPH_cid{}).Error; err != nil {
		return err
	}

	return nil
}
