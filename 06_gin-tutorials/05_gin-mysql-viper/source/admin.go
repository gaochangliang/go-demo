package source

import (
	"gin/05/global"
	"gin/05/model"
	"github.com/gookit/color"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type admin struct{}

var Admin = new(admin)

var admins = []model.SysUser{
	{
		Model: model.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{},
		},
		UUID:        uuid.NewV4(),
		Username:    "admin",
		Password:    "e10adc3949ba59abbe56e057f20f883e",
		NickName:    "超级管理员",
		HeaderImg:   "http://qmplusimg.henrongyi.top/gva_header.jpg",
		AuthorityId: "888",
	},
	{
		Model: model.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{},
		},
		UUID:        uuid.NewV4(),
		Username:    "admin2",
		Password:    "e10adc3949ba59abbe56e057f20f8832",
		NickName:    "2号超级管理员",
		HeaderImg:   "http://qmplusimg.henrongyi.top/gva_header.jpg",
		AuthorityId: "777",
	},
}

func (a *admin) Init() error {
	return global.GLOBAL_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]model.SysUser{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_users 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&admins).Error; err != nil {
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_users 表初始数据成功!")
		return nil
	})
}
