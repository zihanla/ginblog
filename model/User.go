package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	// 自动生成 创建 更新 删除 时间
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

// CheckUser 查询用户是否存在
func CheckUser(name string) int {
	var users User
	// 查询用户id 通过 username 这个条件 如果条件通过 把值写入users
	db.Select("id").Where("username = ?", name).First(&users)
	// 如果 id 大于0 说明数据库 存在此用户
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED // 返回1001 用户已经存在
	}
	return errmsg.SUCCESS
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	// 加密密码 使用钩子函数 实现相同调用效果
	//data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetUsers 查询用户列表
func GetUsers(pageSize, pageNum int) ([]User, int64) {
	// 创建一个切片做容器
	var users []User
	// pageSize 每页大小 pageNum 当前页码
	// limit 限制查询多少数据 offset 设置偏移多少 find 查询到的内容存放哪里
	var total int64
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total).Error
	// 添加limit条件 没有找到 返回ErrRecordNotFound错误
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	// 批量删除 将删除所有匹配的记录
	err = db.Where("id = ?", id).Delete(&user).Error
	// err.Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// EditUser 编辑用户
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	// 更新多列
	err = db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// TODO 密码加密方式这一块不是太懂，整个流程跑通后再研究

// BeforeSave gorm的保存、删除操作会默认运行在事务上，过程不可见，返回了错误会自动回滚，以下为gorm约定的返回方式
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	return nil
}

// ScryptPw 密码加密
func ScryptPw(password string) string {
	const KeyLen = 10
	// 设置一个盐
	salt := make([]byte, 8)
	salt = []byte{6, 28, 38, 77, 25, 225, 27, 11}

	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		// 失败后输出日志再关闭
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

// CheckLogin 登录验证
func CheckLogin(username, password string) int {
	// 创建一个模型容器
	var user User

	// 查找用户名
	db.Where("username = ?", username).First(&user)
	// 验证用户
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	// 验证密码
	if ScryptPw(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	// 验证角色
	if user.Role != 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCESS
}
