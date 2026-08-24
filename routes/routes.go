package routes

import (
	"github.com/NandiniYeligeti/MediCarehms_backend/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(api *gin.RouterGroup) {
	api.POST("/auth/login", Login)
	api.POST("/users/:company_code", CreateUserHandler)
	api.GET("/users/:company_code", GetUsersHandler)
	api.PUT("/users/:company_code/:id/password", ChangePasswordHandler)

	// Super-admin routes
	superAdmin := api.Group("/super-admin")
	superAdmin.Use(middleware.AuthMiddleware())
	{
		superAdmin.GET("/companies", GetCompanies)
		superAdmin.POST("/company", CreateCompany)
	}

	// ================= DOCTORS =================
	doctor := api.Group("/doctor")
	doctor.Use(middleware.AuthMiddleware())
	{
		doctor.POST("/:company_code", CreateDoctor)

		doctor.GET("/:company_code", GetDoctors)

		doctor.GET("/:company_code/:id", GetDoctorByID)

		doctor.GET("/:company_code/entity/:entity_id", GetDoctorByEntityID)

		doctor.PUT("/:company_code/:id", UpdateDoctor)

		doctor.DELETE("/:company_code/:id", DeleteDoctor)
	}
	//========staff======

	staff := api.Group("/staff")
	staff.Use(middleware.AuthMiddleware())
	{
		staff.POST("/:company_code", CreateStaff)

		staff.GET("/:company_code", GetStaffs)

		staff.GET("/:company_code/:id", GetStaffByID)

		staff.GET("/:company_code/entity/:entity_id", GetStaffByEntityID)

		staff.PUT("/:company_code/:id", UpdateStaff)

		staff.DELETE("/:company_code/:id", DeleteHospitalStaff)
	}
	//===ward====
	ward := api.Group("/ward")
	ward.Use(middleware.AuthMiddleware())
	{
		ward.POST("/:company_code", CreateWard)

		ward.GET("/:company_code", GetWards)

		ward.GET("/:company_code/:id", GetWardByID)

		ward.GET("/:company_code/entity/:entity_id", GetWardByEntityID)

		ward.PUT("/:company_code/:id", UpdateWard)

		ward.DELETE("/:company_code/:id", DeleteWard)
	}
	//==========patient=======
	patient := api.Group("/patient")
	patient.Use(middleware.AuthMiddleware())
	{
		patient.POST("/:company_code", CreatePatient)

		patient.GET("/:company_code", GetPatients)

		patient.GET("/:company_code/:id", GetPatientByID)

		patient.GET("/:company_code/entity/:entity_id", GetPatientByEntityID)

		patient.PUT("/:company_code/:id", UpdatePatient)

		patient.DELETE("/:company_code/:id", DeletePatient)
	}
	//============ birth certificate=====
	birthcertificate := api.Group("/birthcertificate")
	birthcertificate.Use(middleware.AuthMiddleware())
	{
		birthcertificate.POST("/:company_code", CreateBirthCertificate)

		birthcertificate.GET("/:company_code", GetBirthCertificates)

		birthcertificate.GET("/:company_code/:id", GetBirthCertificateByID)

		birthcertificate.GET("/:company_code/entity/:entity_id", GetBirthCertificateByEntityID)

		birthcertificate.PUT("/:company_code/:id", UpdateBirthCertificate)

		birthcertificate.DELETE("/:company_code/:id", DeleteBirthCertificate)
	}
	//opd booking
	opdbooking := api.Group("/opdbooking")
	opdbooking.Use(middleware.AuthMiddleware())
	{
		opdbooking.POST("/:company_code", CreateOPDBooking)
		opdbooking.GET("/:company_code", GetOPDBookings)
		opdbooking.GET("/:company_code/:id", GetOPDBookingByID)
		opdbooking.GET("/:company_code/entity/:entity_id", GetOPDBookingByEntityID)
		opdbooking.PUT("/:company_code/:id", UpdateOPDBooking)
		opdbooking.DELETE("/:company_code/:id", DeleteOpdBooking)
	}
}
