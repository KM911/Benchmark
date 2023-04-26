package controller

import (
	"app/dao"

	"github.com/gofiber/fiber/v2"
)

func Like(c *fiber.Ctx) error {
	// 解析参数 从url中获取参数
	id := c.Params("id")
	println("params info ", id)
	dao.Like(id)
	return c.SendString("ok")
}
