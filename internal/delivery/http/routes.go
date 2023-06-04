package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// Handlers represents the collection of HTTP handler functions.
type Handlers struct {
	//UserHandler *UserHandler
	RoleHandler *RoleHandler
	//productHandler *ProductHandler
	// Add more handler fields as needed
}

// Router represents the HTTP router.
type Router struct {
	Router   *fiber.App
	handlers *Handlers
}

// NewRouter creates a new instance of the HTTP router.
func NewRouter(app *fiber.App, handlers *Handlers) *Router {
	return &Router{
		Router:   app,
		handlers: handlers,
	}
}

// SetupRoutes sets up the application routes.
func (r *Router) SetupRoutes() {

	r.Router.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("pong")
	})

	/**
	Swagger Route
	*/
	r.Router.Get("/docs/*", swagger.HandlerDefault)

	/**
	API Prefix
	*/
	api := r.Router.Group("/api")

	/**
	v1 Prefix
	*/
	v1 := api.Group("/v1")

	/**
	Role Feat
	*/
	v1.Get("/roles", r.handlers.RoleHandler.ListRoles)
	v1.Get("/roles/:id", r.handlers.RoleHandler.GetRole)
	v1.Post("/roles", r.handlers.RoleHandler.CreateRole)
	v1.Delete("/roles/:id", r.handlers.RoleHandler.DeleteRole)

	// Add more routes and associated handlers as needed
}
