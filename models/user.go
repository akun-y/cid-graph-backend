package models

import "github.com/jinzhu/gorm"

type GraphUser struct {
	Model

	Name string `json:"name"`

	Website string `json:"website"`
	Github  string `json:"github"`
	Wallet  string `json:"wallet"`
	Linkin  string `json:"linkin"`
	Email   string `json:"email"`
	Cids    string `json:"cids"`

	State int `json:"state"`
}

// ExistUserByName checks if there is a user with the same name
func ExistUserByName(name string) (bool, error) {
	var user GraphUser
	err := db.Select("id").Where("name = ? AND deleted_on = ? ", name, 0).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

// AddUser Add a User
func AddUser(data map[string]interface{}) error {
	user := GraphUser{
		Name:    data["name"].(string),
		State:   data["state"].(int),
		Website: data["website"].(string),
		Github:  data["github"].(string),
		Wallet:  data["wallet"].(string),
		Linkin:  data["linkin"].(string),
		Email:   data["email"].(string),
		Cids:    data["cids"].(string),
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

// GetUser Get a single user based on ID
func GetUser(id int) (*GraphUser, error) {
	var record GraphUser
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

// GetUsers gets a list of user based on paging and constraints
func GetUsers(pageNum int, pageSize int, maps interface{}) ([]GraphUser, error) {
	var (
		user []GraphUser
		err  error
	)

	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Find(&user).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&user).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return user, nil
}

// GetUserTotal counts the total number of user based on the constraint
func GetUserTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&GraphUser{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// ExistUserByID determines whether a User exists based on the ID
func ExistUserByID(id int) (bool, error) {
	var user GraphUser
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

// DeleteUser delete a user
func DeleteUser(id int) error {
	if err := db.Where("id = ?", id).Delete(&GraphUser{}).Error; err != nil {
		return err
	}

	return nil
}

// EditUser modify a single user
func EditUser(id int, data interface{}) error {
	if err := db.Model(&GraphUser{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
