package dao

import (
	"sync"
)

type LIKE struct {
	Id   string
	Like int
}

var (
	// 互斥锁
	MutexLock sync.Mutex
)

/*
通过id对视频点赞数自增
*/
func Like(id string) {
	if id == "1" {
		println("default like")
		UpdateLikeAfterQuery()
	}
	if id == "2" {
		println("mutex like")
		UpdateLikeWithoutQuery()
	}
	if id == "3" {
		println("channel like")
		MutexLike()
	}
}

func UpdateLikeAfterQuery() {
	var like LIKE
	DB.Table("demo").Where("id = 1").First(&like)
	like.Like = like.Like + 1
	DB.Table("demo").Where("id = 1").Update("like", like.Like)
}
func UpdateLikeWithoutQuery() {
	DB.Exec("update demo set `like` = `like` + 1 where id = 2")
}
func MutexLike() {
	MutexLock.Lock()
	var like LIKE
	DB.Table("demo").Where("id = 3").First(&like).Update("like", like.Like+1)
	MutexLock.Unlock()
}
