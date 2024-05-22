package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/lipesalin/go-ecommerce-billing/internal/api/rest"
)

type UserHandler struct {
	// UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	// Instance user service & inject to handler
	handler := UserHandler{}

	// [PUBLIC ROUTES]
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	// [PRIVATE ROUTES]
	app.Get("/verify", handler.GetVerificationCode)
	app.Post("/verify", handler.Verify)

	// profile
	app.Get("/profile", handler.GetProfile)
	app.Post("/profile", handler.CreateProfile)

	// cart
	app.Get("/cart", handler.GetCart)
	app.Post("/cart", handler.AddToCart)

	// order
	app.Get("/order", handler.GetOrders)
	app.Post("/order", handler.CreateOrder)
	app.Get("/order/:id", handler.GetOrder)

	// seller
	app.Get("/become-seller", handler.BecomeSeller)
}

func (u *UserHandler) Register(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "register",
	})
}

func (u *UserHandler) Login(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "login",
	})
}

func (u *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get-verification-code",
	})
}

func (u *UserHandler) Verify(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "verify",
	})
}

func (u *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get-profile",
	})
}

func (u *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "create-profile",
	})
}

func (u *UserHandler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get-cart",
	})
}

func (u *UserHandler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "add-to-cart",
	})
}

func (u *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get-orders",
	})
}

func (u *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get-order-by-id",
	})
}

func (u *UserHandler) CreateOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "create-orders",
	})
}

func (u *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "become-seller",
	})
}
