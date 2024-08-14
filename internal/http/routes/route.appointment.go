package routes

import (
	"ga_marketplace/internal/business/usecases"
	"ga_marketplace/internal/datasources/repositories/postgre"
	"ga_marketplace/internal/http/handlers"
	"ga_marketplace/internal/http/middlewares"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type AppointmentRoute struct {
	appointmentHandler handlers.AppointmentHandler
	router             *echo.Group
	db                 *sqlx.DB
	authMiddleware     middlewares.AuthMiddleware
	adminMiddleware    middlewares.AuthMiddleware
}

func NewAppointmentRoute(
	db *sqlx.DB,
	router *echo.Group,
	authMiddleware middlewares.AuthMiddleware,
	adminMiddleware middlewares.AuthMiddleware,
) *AppointmentRoute {
	serviceItemRepo := postgre.NewPostgreServiceItemRepository(db)
	serviceItemUsecase := usecases.NewServiceItemUsecase(serviceItemRepo)

	appointmentRepo := postgre.NewPostgreAppointmentsRepository(db)
	appointmentUsecase := usecases.NewAppointmentUsecase(appointmentRepo, serviceItemUsecase)

	appointmentHandler := handlers.NewAppointmentHandler(appointmentUsecase)

	return &AppointmentRoute{
		appointmentHandler: *appointmentHandler,
		router:             router,
		db:                 db,
		authMiddleware:     authMiddleware,
		adminMiddleware:    adminMiddleware,
	}
}

func (r *AppointmentRoute) Register() {
	appointments := r.router.Group("/appointments")
	appointmentsAdmin := r.router.Group("/admin/appointments")

	appointments.Use(r.authMiddleware.Handle)
	appointments.POST("", r.appointmentHandler.CreateAppointment)
	appointments.PATCH("/:id/change-time", r.appointmentHandler.ChangeTime)
	appointments.GET("", r.appointmentHandler.FindMyAppointments)

	appointmentsAdmin.Use(r.adminMiddleware.Handle)
	appointmentsAdmin.GET("", r.appointmentHandler.FindAllAppointments)
	appointmentsAdmin.PATCH("/:id", r.appointmentHandler.UpdateAppointment)
	appointmentsAdmin.DELETE("/:id", r.appointmentHandler.DeleteAppointment)
}
