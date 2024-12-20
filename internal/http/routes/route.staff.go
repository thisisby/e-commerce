package routes

import (
	"ga_marketplace/internal/business/usecases"
	"ga_marketplace/internal/datasources/repositories/postgre"
	"ga_marketplace/internal/http/handlers"
	"ga_marketplace/internal/http/middlewares"
	"ga_marketplace/third_party/aws"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type StaffRoute struct {
	staffHandler    handlers.StaffHandler
	s3Client        *aws.S3Client
	router          *echo.Group
	db              *sqlx.DB
	authMiddleware  middlewares.AuthMiddleware
	adminMiddleware middlewares.AuthMiddleware
}

func NewStaffRoute(
	db *sqlx.DB,
	router *echo.Group,
	s3Client *aws.S3Client,
	authMiddleware middlewares.AuthMiddleware,
	adminMiddleware middlewares.AuthMiddleware,
) *StaffRoute {
	staffRepo := postgre.NewPostgreStaffRepository(db)
	appointmentRepo := postgre.NewPostgreAppointmentsRepository(db)
	staffUsecase := usecases.NewStaffUsecase(staffRepo, appointmentRepo)
	staffHandler := handlers.NewStaffHandler(staffUsecase, s3Client)

	return &StaffRoute{
		staffHandler:    staffHandler,
		router:          router,
		db:              db,
		s3Client:        s3Client,
		authMiddleware:  authMiddleware,
		adminMiddleware: adminMiddleware,
	}
}

func (r *StaffRoute) Register() {
	staff := r.router.Group("/staff")
	admin := r.router.Group("/admin/staff")
	services := r.router.Group("/services")
	serviceAddress := r.router.Group("/service-address")

	staff.Use(r.authMiddleware.Handle)
	staff.GET("", r.staffHandler.FindAll)
	staff.GET("/:staffId/time-slots", r.staffHandler.FindTimeSlotByStaffId)

	serviceAddress.Use(r.authMiddleware.Handle)
	serviceAddress.GET("/:service_address_id/staff", r.staffHandler.FindByServiceAddressId)

	services.Use(r.authMiddleware.Handle)
	services.GET("/:service_id/staff", r.staffHandler.FindByServiceId)

	admin.Use(r.adminMiddleware.Handle)
	admin.POST("", r.staffHandler.Save)
	admin.PATCH("/:id", r.staffHandler.Update)
	admin.DELETE("/:id", r.staffHandler.Delete)
}
