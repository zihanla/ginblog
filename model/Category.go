package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// CheckCategory 查询分类是否存在
func CheckCategory(name string) int {
	var cate Category
	// 查询分类id 通过 name 这个条件 如果条件通过 把值写入users
	db.Select("id").Where("name = ?", name).First(&cate)
	// 如果 id 大于0 说明数据库 存在此分类
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED // 返回3001 分类已经存在
	}
	return errmsg.SUCCESS
}

// CreateCate 新增分类
func CreateCate(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetCate 查询分类列表
func GetCate(pageSize, pageNum int) ([]Category, int64) {
	// 创建一个切片做容器
	var cate []Category
	// pageSize 每页大小 pageNum 当前页码
	// limit 限制查询多少数据 offset 设置偏移多少 find 查询到的内容存放哪里
	var total int64
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Count(&total).Error
	// 添加limit条件 没有找到 返回ErrRecordNotFound错误
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cate, total
}

// todo 查询分类下的所有文章

// DeleteCate 删除分类
func DeleteCate(id int) int {
	var cate Category
	// 批量删除 将删除所有匹配的记录
	err = db.Where("id = ?", id).Delete(&cate).Error
	// err.Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// EditCate 编辑分类
func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	// 更新多列
	err = db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
