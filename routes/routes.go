package routes

import "github.com/gin-gonic/gin"

func Routes(api *gin.RouterGroup) {

	// ================= DOCTORS =================
	doctor := api.Group("/doctor")
	// doctor.Use(middleware.AuthMiddleware())
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
	// staff.Use(middleware.AuthMiddleware())
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
// ward.Use(middleware.AuthMiddleware())
{
	ward.POST("/:company_code", CreateWard)

	ward.GET("/:company_code", GetWards)

	ward.GET("/:company_code/:id", GetWardByID)

	ward.GET("/:company_code/entity/:entity_id", GetWardByEntityID)

	ward.PUT("/:company_code/:id", UpdateWard)

	ward.DELETE("/:company_code/:id", DeleteWard)
}
}