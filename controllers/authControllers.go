package controllers

import "github.com/gofiber/fiber/v2"

type AuthControllers interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
	GetUserProfile(ctx *fiber.Ctx) error
	UpdateUserProfile(ctx *fiber.Ctx) error
	UploadUserPhoto(ctx *fiber.Ctx) error
}
