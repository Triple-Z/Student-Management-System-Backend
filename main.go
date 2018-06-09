package main

import (
	"./app"
	"github.com/gin-gonic/gin"
)

func main() {

	//gin.SetMode(gin.ReleaseMode)  // Gin release mode

	router := gin.Default()

	ApiV1 := router.Group("/api/v1")
	{
		studentsApi := ApiV1.Group("/students")
		{
			studentsApi.GET("", app.FetchAllStudents)
			studentsApi.POST("", app.CreateStudent)
			studentsApi.GET("/:id", app.FetchSingleStudent)
			studentsApi.PUT("/:id", app.UpdateStudent)
			studentsApi.DELETE("/:id", app.DeleteStudent)
		}

		departmentsApi := ApiV1.Group("/departments")
		{
			departmentsApi.GET("", app.FetchAllDepartments)
			departmentsApi.POST("", app.CreateDepartment)
			departmentsApi.GET("/:id", app.FetchSingleDepartment)
			departmentsApi.PUT("/:id", app.UpdateDeprtment)
			departmentsApi.DELETE("/:id", app.DeleteDepartment)
		}
	}

	router.Run(":3000")
}
