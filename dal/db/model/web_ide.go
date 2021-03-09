package model

import (
	"gorm.io/gorm"
)

// Container [...]
type Container struct {
	gorm.Model
	ContainerID string `gorm:"unique"` // docker容器id
	Status      int8   // Docker容器状态 1-Running 2-Paused 3-Restarting 4-OOMKilled 5-Dead 6-UNKNOWN
	ImageID     string `gorm:"index:idx_image_id"` // docker镜像id
	Name        string // docker容器名字
}

// Group [...]
type Group struct {
	gorm.Model
	OwnerID   uint32 `gorm:"uniqueIndex:uniq_idx_owner_id_group_name"`             // 组长的user_id
	GroupName string `gorm:"uniqueIndex:uniq_idx_owner_id_group_name;default:未命名"` // 组名
}

// Image [...]
type Image struct {
	gorm.Model
	ImageID   string `gorm:"unique"` // docker镜像id
	ImageSize int64  // 镜像大小，单位byte
	Author    string // 镜像作者
	RepoTags  string // 仓库标签，以json数组的形式存储
}

// Resource [...]
type Resource struct {
	gorm.Model
	Name string // 资源名称
}

// Role [...]
type Role struct {
	gorm.Model
	Name string // 角色名称
}

// RoleResource [...]
type RoleResource struct {
	gorm.Model
	RoleID     uint32 `gorm:"index:idx_user_id"` // role表的id
	ResourceID uint32 // resource表的id
}

// User [...]
type User struct {
	gorm.Model
	Username string `gorm:"unique;unique"` // 用户名
	Password string // 密码
}

// UserContainer [...]
type UserContainer struct {
	gorm.Model
	UserID      uint32 `gorm:"index:idx_user_id"` // user表的id
	ContainerID string // docker中的container_id
}

// UserExtra [...]
type UserExtra struct {
	gorm.Model
	UserID      uint32 `gorm:"unique"` // user表的id
	Nickname    string // 昵称
	PhoneNumber string // 电话号码
	Email       string // 邮箱
	AvatarURL   string // 头像链接
}

// UserGroup [...]
type UserGroup struct {
	gorm.Model
	UserID  uint32 `gorm:"index:idx_user_id"`  // user表的id
	GroupID string `gorm:"index:idx_group_id"` // docker中的image_id
}

// UserImage [...]
type UserImage struct {
	gorm.Model
	UserID  uint32 `gorm:"index:idx_user_id"` // user表的id
	ImageID string // docker中的image_id
}

// UserRole [...]
type UserRole struct {
	gorm.Model
	UserID uint32 `gorm:"index:idx_user_id"` // user表的id
	RoleID uint32 // role表的id
}
