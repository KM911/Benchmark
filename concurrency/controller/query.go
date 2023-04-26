package controller

import (
	"app/dao"
	"github.com/gofiber/fiber/v2"
)

func Query(c *fiber.Ctx) error {
	// 解析参数 从url中获取参数
	id := c.Params("id")
	return c.SendString(dao.QueryLike(id))
}
