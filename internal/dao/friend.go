package dao

import (
	"blog/internal/model"
	"fmt"
)

func (d *Dao) GetFriendList() (friends []*model.Friend, err error) {
	err = d.DB.Model(&model.Friend{}).Find(&friends).Error
	if err != nil {
		fmt.Println("err: ", err)
	}
	return
}

func (d *Dao) AddFriend(friend *model.Friend) error {
	err := d.DB.Create(friend).Error
	if err != nil {
		fmt.Println("err: ", err)
	}
	return err
}

func (d *Dao) UpdateFriend(friend *model.Friend) error {
	err := d.DB.Model(&model.Friend{}).Updates(friend).Error
	if err != nil {
		fmt.Println("err: ", err)
	}
	return err
}

func (d *Dao) DelFriend(id uint) error {
	err := d.DB.Delete(&model.Friend{}, id).Error
	if err != nil {
		fmt.Println("err: ", err)
	}
	return err
}
