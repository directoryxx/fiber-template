package http

import (
	logger2 "clean-arch-template/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

var logger = logger2.NewLogger()

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
