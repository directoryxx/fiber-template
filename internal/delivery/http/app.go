package http

import (
	"clean-arch-template/internal/repository"
	"clean-arch-template/internal/usecase"
	"clean-arch-template/pkg/config"
	"clean-arch-template/pkg/database/sqlc"
	"clean-arch-template/pkg/infrastructure"
	"clean-arch-template/pkg/logger"
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/contrib/fibersentry"
	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"os"
)

type ContextData struct {
	// Store information for custom TracesSampler.
	request *fasthttp.Request
	// ...
}

var (
	newLogger = logger.NewLogger()
)

func Run() *fiber.App {
	/**
	Load Config
	*/
	envSource := config.LoadConfig()
	newLogger.Info("Loaded Config : " + envSource)

	fmt.Println("svc " + os.Getenv("SERVICE_NAME"))
	fmt.Println("url " + os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT"))
	fmt.Println("insecure " + os.Getenv("INSECURE_MODE"))

	cleanup := initTracer()
	defer cleanup(context.Background())

	/**
	Init Sentry
	*/
	_ = sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_URI_DSN"),
		Debug:            true,
		AttachStacktrace: true,
		EnableTracing:    true,
		// Specify a fixed sample rate:
		// We recommend adjusting this value in production
		TracesSampleRate: 1.0,
	})

	/**h
	Load Infrastructure
	*/
	newLogger.Info("Opening Connection Database")
	dbPool, err := infrastructure.OpenDB()
	if err != nil {
		newLogger.Error("Error Connection Database : " + err.Error())
		panic(err)
	}
	newLogger.Info("Success Connection Database")
	//defer dbPool.Close()

	/**
	SQLC Init
	*/
	sqlcQueries := sqlc.New(dbPool)

	/**
	Init framework
	*/
	app := fiber.New(fiber.Config{
		ServerHeader: "App",
		AppName:      "Clean Arch Template v0.0.1",
	})

	/**
	Implement Sentry to middleware
	*/
	app.Use(fibersentry.New(fibersentry.Config{
		Repanic:         true,
		WaitForDelivery: true,
	}))

	app.Use(otelfiber.Middleware())

	/**
	Pprof
	*/
	//app.Use(pprof.New())

	// Create repository instances
	roleRepository := repository.NewRoleRepository(sqlcQueries)

	// Create usecase instances
	roleUsecase := usecase.NewRoleService(roleRepository)

	// Create handler instances
	handlers := &Handlers{
		//UserHandler: http.NewUserHandler(),
		RoleHandler: NewRoleHandler(roleUsecase),
		// Initialize other handlers as neededd
	}

	// Create router instance and pass the handler
	router := NewRouter(app, handlers)

	// Set up routes
	router.SetupRoutes()

	return app
}

func initTracer() func(context.Context) error {

	//secureOption := otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	//if len(os.Getenv("INSECURE_MODE")) > 0 {
	secureOption := otlptracegrpc.WithInsecure()
	//}

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracegrpc.NewClient(
			secureOption,
			otlptracegrpc.WithEndpoint(os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")),
		),
	)

	if err != nil {
		newLogger.Error(err.Error())
	}
	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", os.Getenv("SERVICE_NAME")),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		newLogger.Error("Could not set resources: " + err.Error())
	}

	otel.SetTracerProvider(
		sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithBatcher(exporter),
			sdktrace.WithResource(resources),
		),
	)
	return exporter.Shutdown
}
