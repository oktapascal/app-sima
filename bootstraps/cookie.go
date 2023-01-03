package bootstraps

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type Cookie interface {
	GetCookieToken() string
	CreateTokenAndCookie(data string, expiration time.Time, ctx *fiber.Ctx)
	DeleteCookie(ctx *fiber.Ctx)
}

type CookieImpl struct {
	Config Config
}

func NewCookieImpl(config Config) *CookieImpl {
	return &CookieImpl{Config: config}
}

func (cookie *CookieImpl) GetCookieToken() string {
	return cookie.Config.Get("COOKIE_ACCESS")
}

func (cookie *CookieImpl) CreateTokenAndCookie(data string, expiration time.Time, ctx *fiber.Ctx) {
	cookies := cookie.setTokenCookie(cookie.GetCookieToken(), data, expiration)

	ctx.Cookie(cookies)
}

func (cookie *CookieImpl) DeleteCookie(ctx *fiber.Ctx) {
	cookies := cookie.setTokenCookie(cookie.GetCookieToken(), "deleted", time.Now().Add(-3*time.Second))

	ctx.Cookie(cookies)
}

func (cookie *CookieImpl) setTokenCookie(name string, data string, expiration time.Time) *fiber.Cookie {
	cookies := new(fiber.Cookie)
	cookies.Name = name
	cookies.Value = data
	cookies.Expires = expiration
	cookies.HTTPOnly = true

	return cookies
}
