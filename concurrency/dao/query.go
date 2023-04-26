package dao

import "strconv"

func QueryLike(id string) string {
	var like LIKE
	DB.Table("demo").Where("id = ?", id).First(&like)
	// 返回点赞数 string类型
	return strconv.Itoa(like.Like)
}
