package controllers

import "github.com/gofiber/fiber/v2"

type AuthControllers interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
}
