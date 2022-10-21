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
			//-------------------------------------------------------------------------------------//xxxBook//-----------------------------------------------------------------------------------//
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
			//-------------------------------------------------------------------------------------//xxBook//-----------------------------------------------------------------------------------//

			//-------------------------------------------------------------------------------------//xxxUser//-----------------------------------------------------------------------------------//
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
			//-------------------------------------------------------------------------------------//xxUser//-----------------------------------------------------------------------------------//

			//-------------------------------------------------------------------------------------//xxxBorrow//---------------------------------------------------------------------------------//
			router.GET("/borrows", controller.ListBorrows)
			router.GET("/borrow/:id", controller.GetBorrows)
			router.POST("/borrows", controller.CreateBorrow)
			router.PATCH("/borrows", controller.UpdateBorrow)
			router.DELETE("/borrows/:id", controller.DeleteBorrow)
			//-------------------------------------------------------------------------------------//xxBorrow//---------------------------------------------------------------------------------//

			//-------------------------------------------------------------------------------------//xxxBill//---------------------------------------------------------------------------------//
			//bill route
			router.GET("/bills", controller.ListBills)
			router.GET("/bill/:id", controller.GetBill)
			router.POST("/bills", controller.CreateBill)
			router.PATCH("/bills", controller.UpdateBill)
			router.DELETE("/bills/:id", controller.DeleteBill)
			//-------------------------------------------------------------------------------------//xxBill//---------------------------------------------------------------------------------//
			
			//-------------------------------------------------------------------------------------//xxxReseachRoom//---------------------------------------------------------------------------------//
			// Computer_os Routes
			router.GET("/computer_oss", controller.ListComputer_oss)
			router.GET("/computer_os/:id", controller.GetComputer_os)
			router.POST("/computer_oss", controller.CreateComputer_os)
			router.PATCH("/computer_oss", controller.UpdateComputer_os)
			router.DELETE("/computer_oss/:id", controller.DeleteComputer_os)

			// Computer_reservation Routes
			router.GET("/computer_reservations", controller.ListComputer_reservations)
			router.GET("/computer_reservation/:id", controller.GetComputer_reservation)
			// api.GET("/playlist/watched/user/:id", controller.GetPlaylistWatchedByUser)
			router.POST("/computer_reservations", controller.CreateComputer_reservation)
			router.PATCH("/computer_reservations", controller.UpdateComputer_reservation)
			router.DELETE("/computer_reservations/:id", controller.DeleteComputer_reservation)

			// Computer Routes
			router.GET("/computers", controller.ListComputers)
			router.GET("/computer/:id", controller.GetComputer)
			router.POST("/computers", controller.CreateComputer)
			router.PATCH("/computers", controller.UpdateComputer)
			router.DELETE("/computers/:id", controller.DeleteComputer)

			// Time_com Routes
			router.GET("/time_coms", controller.ListTime_coms)
			router.GET("/time_com/:id", controller.GetTime_com)
			router.POST("/time_coms", controller.CreateTime_com)
			router.PATCH("/time_coms", controller.UpdateTime_com)
			router.DELETE("/time_coms/:id", controller.DeleteTime_com)
			//-------------------------------------------------------------------------------------//xxReseachRoom//---------------------------------------------------------------------------------//
			
			//-------------------------------------------------------------------------------------//xxxProblem//-------------------------------------------------------------//
			router.GET("/place_class", controller.ListPlace_Class)
			router.GET("/place_class/:id", controller.GetPlace_Class)
			router.POST("/place_class", controller.CreatePlace_Class)
			router.PATCH("/place_class", controller.UpdatePlace_Class)
			router.DELETE("/place_class/:id", controller.DeletePlace_Class)

			router.GET("/problem", controller.ListProblems)
			router.GET("/problem/:id", controller.GetProblem)
			router.POST("/problem", controller.CreateProblem)
			router.PATCH("/problem", controller.UpdateProblem)
			router.DELETE("/problem/:id", controller.DeleteProblem)

			router.GET("/problemreport", controller.ListProblemReports)
			router.GET("/problemreport/:id", controller.GetProblemReport)
			router.POST("/problemreport", controller.CreateProblemReport)
			router.PATCH("/problemreport", controller.UpdateProblemReport)
			router.DELETE("/problemreport/:id", controller.DeleteProblemReport)

			router.GET("/readingzone", controller.ListReadingZones)
			router.GET("/readingzone/:id", controller.GetReadingZone)
			router.POST("/readingzone", controller.CreateReadingZone)
			router.PATCH("/readingzone", controller.UpdateReadingZone)
			router.DELETE("/readingzone/:id", controller.DeleteReadingZone)

			router.GET("/researchroom", controller.ListResearchRooms)
			router.GET("/researchroom/:id", controller.GetResearchRoom)
			router.POST("/researchroom", controller.CreateResearchRoom)
			router.PATCH("/researchroom", controller.UpdateResearchRoom)
			router.DELETE("/researchroom/:id", controller.DeleteResearchRoom)

			router.GET("/toilet", controller.ListToilets)
			router.GET("/toilet/:id", controller.GetToilet)
			router.POST("/toilet", controller.CreateToilet)
			router.PATCH("/toilet", controller.UpdateToilet)
			router.DELETE("/toilet/:id", controller.DeleteToilet)
			//-------------------------------------------------------------------------------------//xxProblem//-------------------------------------------------------------//

			//-------------------------------------------------------------------------------------//xxReservation//-------------------------------------------------------------//
			// Research_Room Routes
			router.GET("/researchrooms", controller.ListResearchRooms)
			router.GET("/researchroom/:id", controller.GetResearchRoom)
			router.POST("/researchrooms", controller.CreateResearchRoom)
			router.PATCH("/researchrooms", controller.UpdateResearchRoom)
			router.DELETE("/researchrooms/:id", controller.DeleteResearchRoom)

			// Equipment Routes
			router.GET("/equipments", controller.ListEquipments)
			router.GET("/equipment/:id", controller.GetEquipment)
			router.POST("/equipments", controller.CreateEquipment)
			router.PATCH("/equipments", controller.UpdateEquipment)
			router.DELETE("/equipments/:id", controller.DeleteEquipment)

			// Room_Type Routes
			router.GET("/roomtypes", controller.ListRoomTypes)
			router.GET("/roomtype/:id", controller.GetRoomType)
			router.POST("/roomtypes", controller.CreateRoomType)
			router.PATCH("/roomtypes", controller.UpdateRoomType)
			router.DELETE("/roomtypes/:id", controller.DeleteRoomType)

			// AddOn Routes
			router.GET("/addons", controller.ListAddOns)
			router.GET("/addon/:id", controller.GetAddOn)
			router.POST("/addons", controller.CreateAddOn)
			router.PATCH("/addons", controller.UpdateAddOn)
			router.DELETE("/addons/:id", controller.DeleteAddOn)

			// Timeroom Routes
			router.GET("/timerooms", controller.ListTimes)
			router.GET("/timeroom/:id", controller.GetTime)
			router.POST("/timerooms", controller.CreateTime)
			router.PATCH("/timerooms", controller.UpdateTime)
			router.DELETE("/timerooms/:id", controller.DeleteTime)

			// Research_Room_Reservation_Record Routes
			router.GET("/researchroomreservationrecords", controller.ListResearchRoomReservationRecords)
			router.GET("/researchroomreservationrecord/:id", controller.GetResearchRoomReservationRecord)
			router.POST("researchroomreservationrecords", controller.CreateResearchRoomReservationRecord)
			router.PATCH("/researchroomreservationrecords", controller.UpdateResearchRoomReservationRecord)
			router.DELETE("/researchroomreservationrecords/:id", controller.DeleteResearchRoomReservationRecord)

			//-------------------------------------------------------------------------------------//xxReservation//-------------------------------------------------------------//
		
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
