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
}
