package controllers

import "github.com/gofiber/fiber/v2"

type MenuControllers interface {
	GetAll(ctx *fiber.Ctx) error
}
