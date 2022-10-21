package main

import (
	"github.com/siwanastudio/SA-65-SW/controller"

	"github.com/siwanastudio/SA-65-SW/entity"

	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	// r.Use(CORSMiddleware())

	router := r.Group("/")
	{
		// router.Use(middlewares.Authorizes())
		{	
			//-------------------------------------------------------------------------------------//Book//-----------------------------------------------------------------------------------//
			// shelf Routes
			router.GET("/shelves", controller.ListShelf)
			router.GET("/shelf/:id", controller.GetShelf)
			router.POST("/shelves", controller.CreateShelf)
			router.PATCH("/shelves", controller.UpdateShelf)
			router.DELETE("/shelves/:id", controller.DeleteShelf)

			// BOOK_TYPE Routes
			router.GET("/book_types", controller.ListBookType)
			router.GET("/book_types/:id", controller.GetBookType)
			router.POST("/book_types", controller.CreateBookType)
			router.PATCH("/book_types", controller.UpdateBookType)
			router.DELETE("/book_types/:id", controller.DeleteBookType)

			// Role Routes
			router.GET("/roles", controller.ListRole)
			router.GET("/roles/:id", controller.GetRole)
			router.POST("/roles", controller.CreateRole)
			router.PATCH("/roles", controller.UpdateRole)
			router.DELETE("/roles/:id", controller.DeleteRole)

			/*// employee Routes
			router.GET("/employees", controller.ListEmp)
			router.GET("/employees/:id", controller.GetEmp)
			//router.POST("/employees", controller.CreateEmployee)
			router.PATCH("/employees", controller.UpdateEmp)
			router.DELETE("/employees/:id", controller.DeleteEmp)*/

			// book Routes
			router.GET("/books", controller.ListBook)
			router.GET("/book/:id", controller.GetBook)
			router.POST("/books", controller.CreateBook)
			router.PATCH("/books", controller.UpdateBook)
			router.DELETE("/books/:id", controller.DeleteBook)
			//-------------------------------------------------------------------------------------//Book//-----------------------------------------------------------------------------------//

			//-------------------------------------------------------------------------------------//User//-----------------------------------------------------------------------------------//
			//user routes
			router.GET("/users", controller.ListUser)
			router.GET("/user/:id", controller.GetUser)
			router.POST("/users", controller.CreateUser)
			router.PATCH("/users", controller.UpdateUser)
			router.DELETE("/users/:id", controller.DeleteUser)

			//employee routes
			router.GET("/employees", controller.ListEmployee)
			router.GET("/employee/:id", controller.GetEmployee)
			//r.POST("/employees", controller.CreateEmployee)
			router.PATCH("/employees", controller.UpdateEmployee)
			router.DELETE("/employees/:id", controller.DeleteEmployee)

			//memberClass routes
			router.GET("/memberclasses", controller.ListMemberClass)
			router.GET("/memberclass/:id", controller.GetMemberClass)
			router.POST("/memberclasses", controller.CreateMemberClass)
			router.PATCH("/memberclasses", controller.UpdateMemberclass)
			router.DELETE("/memberclasses/:id", controller.DeleteMemberClass)

			//province routes
			router.GET("/provinces", controller.ListProvince)
			router.GET("/province/:id", controller.GetProvince)
			router.POST("/provinces", controller.CreateProvince)
			router.PATCH("/provinces", controller.UpdateProvince)
			router.DELETE("/provinces/:id", controller.DeleteMemberClass)

			//role routes
			router.GET("/roles", controller.ListRole)
			router.GET("/role/:id", controller.GetRole)
			router.POST("/roles", controller.CreateUser)
			router.PATCH("/roles", controller.UpdateRole)
			router.DELETE("/roles/:id", controller.DeleteRole)
			//-------------------------------------------------------------------------------------//User//-----------------------------------------------------------------------------------//

			//-------------------------------------------------------------------------------------//Borrow//---------------------------------------------------------------------------------//
			router.GET("/borrows", controller.ListBorrows)
			router.GET("/borrow/:id", controller.GetBorrows)
			router.POST("/borrows", controller.CreateBorrow)
			router.PATCH("/borrows", controller.UpdateBorrow)
			router.DELETE("/borrows/:id", controller.DeleteBorrow)
			//-------------------------------------------------------------------------------------//Borrow//---------------------------------------------------------------------------------//
		}
	}

	// Signup User Route
	// r.POST("/signup", controller.CreateUser)
	// login User Route
	// r.POST("/login", controller.Login)

	// Run the server go run main.go
	r.Run("localhost: " + PORT)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
