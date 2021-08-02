package model

import (
	"github.com/satori/go.uuid"
)

// SysUser
//CREATE TABLE `sys_users` (
//  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
//  `created_at` datetime DEFAULT NULL,
//  `updated_at` datetime DEFAULT NULL,
//  `deleted_at` datetime DEFAULT NULL,
//  `uuid` varchar(191) DEFAULT NULL COMMENT '用户UUID',
//  `username` varchar(191) DEFAULT NULL COMMENT '用户登录名',
//  `password` varchar(191) DEFAULT NULL COMMENT '用户登录密码',
//  `nick_name` varchar(191) DEFAULT '系统用户' COMMENT '用户昵称',
//  `header_img` varchar(191) DEFAULT 'http://qmplusimg.henrongyi.top/head.png' COMMENT '用户头像',
//  `authority_id` varchar(191) DEFAULT '888' COMMENT '用户角色ID',
//  `side_mode` varchar(191) DEFAULT 'dark' COMMENT '用户角色ID',
//  `active_color` varchar(191) DEFAULT '#1890ff' COMMENT '用户角色ID',
//  `base_color` varchar(191) DEFAULT '#fff' COMMENT '用户角色ID',
//  PRIMARY KEY (`id`),
//  KEY `idx_sys_users_deleted_at` (`deleted_at`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

type SysUser struct {
	Model
	UUID        uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`                                                    // 用户UUID
	Username    string    `json:"userName" gorm:"comment:用户登录名"`                                                 // 用户登录名
	Password    string    `json:"-"  gorm:"comment:用户登录密码"`                                                      // 用户登录密码
	NickName    string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                     // 用户昵称
	HeaderImg   string    `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"` // 用户头像
	AuthorityId string    `json:"authorityId" gorm:"default:888;comment:用户角色ID"`                                 // 用户角色ID
	SideMode    string    `json:"sideMode" gorm:"default:dark;comment:用户角色ID"`                                   // 用户侧边主题
	ActiveColor string    `json:"activeColor" gorm:"default:#1890ff;comment:用户角色ID"`                             // 活跃颜色
	BaseColor   string    `json:"baseColor" gorm:"default:#fff;comment:用户角色ID"`                                  // 基础颜色
}
