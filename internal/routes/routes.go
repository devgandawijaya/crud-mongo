package routes

import (
	"crud-mongo/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	company := r.Group("/companies")
	{
		company.POST("", controllers.CreateCompany)
		company.GET("", controllers.GetCompanies)
		company.GET("/:id", controllers.GetCompanyByID)
		company.PUT("/:id", controllers.UpdateCompany)
		company.DELETE("/:id", controllers.DeleteCompany)
	}

	product := r.Group("/products")
	{
		product.POST("", controllers.CreateProduct)
		product.GET("", controllers.GetProducts)
		product.GET("/:id", controllers.GetProductByID)
		product.PUT("/:id", controllers.UpdateProduct)
		product.DELETE("/:id", controllers.DeleteProduct)
	}

}
