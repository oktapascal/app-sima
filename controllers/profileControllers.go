package controllers

import "github.com/gofiber/fiber/v2"

type ProfileControllers interface {
	GetUser(ctx *fiber.Ctx) error
	UpdateUser(ctx *fiber.Ctx) error
}
