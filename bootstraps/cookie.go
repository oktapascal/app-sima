package bootstraps

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

// Cookie is an interface for managing auth tokens stored in cookies.
type Cookie interface {
	// GetCookieToken returns the name of the cookie that stores the auth token.
	GetCookieToken() string
	// CreateTokenAndCookie stores the given auth token in a cookie.
	CreateTokenAndCookie(data string, expiration time.Time, ctx *fiber.Ctx)
	// DeleteCookie removes the auth token from the cookie.
	DeleteCookie(ctx *fiber.Ctx)
}

// CookieImpl is a concrete implementation of the Cookie interface.
type CookieImpl struct {
	// Config is a Config interface used to get the name of the cookie for storing the auth token.
	Config Config
}

// NewCookieImpl creates a new CookieImpl object using the given Config object.
func NewCookieImpl(config Config) *CookieImpl {
	// Return a new CookieImpl object initialized with the given Config object.
	return &CookieImpl{Config: config}
}

// GetCookieToken returns the name of the cookie that stores the auth token.
func (cookie *CookieImpl) GetCookieToken() string {
	// Use the Config object to get the name of the cookie.
	return cookie.Config.Get("COOKIE_ACCESS")
}

// CreateTokenAndCookie stores the given auth token in a cookie.
func (cookie *CookieImpl) CreateTokenAndCookie(data string, expiration time.Time, ctx *fiber.Ctx) {
	// Create a new cookie with the given data and expiration.
	cookies := cookie.setTokenCookie(cookie.GetCookieToken(), data, expiration)
	// Set the cookie in the HTTP response.
	ctx.Cookie(cookies)
}

// DeleteCookie removes the auth token from the cookie.
func (cookie *CookieImpl) DeleteCookie(ctx *fiber.Ctx) {
	// Create a new cookie with an expired expiration date and the value "deleted".
	cookies := cookie.setTokenCookie(cookie.GetCookieToken(), "deleted", time.Now().Add(-3*time.Second))
	// Set the cookie in the HTTP response.
	ctx.Cookie(cookies)
}

// setTokenCookie creates a new fiber.Cookie object with the given name, data, and expiration date.
func (cookie *CookieImpl) setTokenCookie(name string, data string, expiration time.Time) *fiber.Cookie {
	// Initialize a new fiber.Cookie object.
	cookies := new(fiber.Cookie)
	cookies.Name = name
	cookies.Value = data
	cookies.Expires = expiration
	cookies.HTTPOnly = true

	return cookies
}
