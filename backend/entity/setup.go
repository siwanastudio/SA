package entity

import (
	"time"
	//"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("sa5.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(
		//Book
		&Book{},
		&BookType{},
		&Shelf{},
		//User
		&Province{},
		&MemberClass{},
		&User{},
		&Role{},
		&Employee{},
		//Borrow
		&Borrow{},
		//reseachroom
		&RoomType{},
		&Equipment{},
		&ResearchRoom{},
		&TimeRoom{},
		&AddOn{},
		&ResearchRoomReservationRecord{},
		//Bill
		&Bill{},
		//Com_reser
		&Computer_os{},
		&Computer_reservation{},
		&Computer{},
		&Time_com{},
		//Problem
		&Place_Class{},
		&Relation{},
		&Toilet{},
		&ReadingZone{},
		&ProblemReport{},
	)

	db = database
	//-------------------------------------------------------------------------------------//xxxUser//-----------------------------------------------------------------------------------//
	//User
	//password1, err := bcrypt.GenerateFromPassword([]byte("zaq1@wsX"), 14)
	//password2, err := bcrypt.GenerateFromPassword([]byte("zxvseta"), 14)
	//password3, err := bcrypt.GenerateFromPassword([]byte("1111111111111"), 14)

	//add example data
	//emp
	
	db.Model(&Employee{}).Create(&Employee{
		Name:     "Sirinya",
		Email:    "sirinya@mail.com",
		Password: "zaq1@wsX",
	})

	db.Model(&Employee{}).Create(&Employee{
		Name:     "Attawit",
		Email:    "attawit@mail.com",
		Password: "zxvseta",
	})

	var sirin Employee
	db.Raw("SELECT * FROM employees WHERE email = ?", "sirinya@mail.com").Scan(&sirin)
	var Attawit Employee
	db.Raw("SELECT * FROM employees WHERE email = ?", "attawit@mail.com").Scan(&Attawit)
	//Role
	student := Role{
		Name:       "Student",
		BorrowDay:  3,
		BookRoomHR: 3,
		BookComHR:  4,
	}
	db.Model(&Role{}).Create(&student)

	teacher := Role{
		Name:       "Teacher",
		BorrowDay:  7,
		BookRoomHR: 12,
		BookComHR:  12,
	}
	db.Model(&Role{}).Create(&teacher)

	//province
	korat := Province{
		Name: "Nakhon Ratchasima",
	}
	db.Model(&Province{}).Create(&korat)

	chon := Province{
		Name: "Chonburi",
	}
	db.Model(&Province{}).Create(&chon)

	bangkok := Province{
		Name: "Bangkok",
	}
	db.Model(&Province{}).Create(&bangkok)

	//member
	classic := MemberClass{
		Name:     "classic",
		Discount: 0,
	}
	db.Model(&MemberClass{}).Create(&classic)

	silver := MemberClass{
		Name:     "silver",
		Discount: 5,
	}
	db.Model(&MemberClass{}).Create(&silver)

	gold := MemberClass{
		Name:     "gold",
		Discount: 10,
	}
	db.Model(&MemberClass{}).Create(&gold)

	plat := MemberClass{
		Name:     "platinum",
		Discount: 20,
	}
	db.Model(&MemberClass{}).Create(&plat)

	//user
	/*db.Model(&User{}).Create(&User{
		PIN:       "B6111111",
		FirstName: "preecha",
		LastName:  "anpa",
		CIV:       "1111111111111",
		PHONE:     "0811111111",
		Email:     "preechapat@mail.com",
		Password:  "1111111111111",
		ADDRESS:   "ถนน a อำเภอ v",
		//FK
		Employee:    sirin,
		Role:        student,
		Province:    korat,
		MemberClass: classic,
	})*/
	b6111111 := User{
		Pin:       "B6111111",
		FirstName: "preecha",
		LastName:  "anpa",
		Civ:       "1111111111111",
		Phone:     "0811111111",
		Email:     "preechapat@mail.com",
		Password:  "1111111111111",
		Address:   "ถนน a อำเภอ v",
		//FK
		Employee:    sirin,
		Role:        student,
		Province:    korat,
		MemberClass: classic,
	}
	db.Model(&User{}).Create(&b6111111)
	B6222222 := User{
		Pin:       "B6222222",
		FirstName: "kawin",
		LastName:  "l.pat",
		Civ:       "2222222222222",
		Phone:     "0922222222",
		Email:     "kawin@mail.com",
		Password:  "2222222222222",
		Address:   "หอ b อำเภอ r",
		//FK
		Employee:    sirin,
		Role:        student,
		Province:    korat,
		MemberClass: classic,
	}
	db.Model(&User{}).Create(&B6222222)
	//-------------------------------------------------------------------------------------//xxUser//-----------------------------------------------------------------------------------//

	//-------------------------------------------------------------------------------------//xxxBook//-----------------------------------------------------------------------------------//

	//Role
	R1 := Role{
		Name:        "Student",
		BorrowDay:  3,
		BookRoomHR: 3,
		BookComHR:  4,
	}
	db.Model(&Role{}).Create(&R1)
	R2 := Role{
		Name:        "Teacher",
		BorrowDay:  7,
		BookRoomHR: 12,
		BookComHR:  12,
	}
	db.Model(&Role{}).Create(&R2)

	//Shelf
	S1 := Shelf{
		Type:  "SCIENCE",
		Floor: 1,
	}
	db.Model(&Shelf{}).Create(&S1)
	S2 := Shelf{
		Type:  "ENGINEERING",
		Floor: 1,
	}
	db.Model(&Shelf{}).Create(&S2)
	S3 := Shelf{
		Type:  "ENVIRRONMENT",
		Floor: 1,
	}
	db.Model(&Shelf{}).Create(&S3)
	S4 := Shelf{
		Type:  "HISTORY",
		Floor: 1,
	}
	db.Model(&Shelf{}).Create(&S4)
	S5 := Shelf{
		Type:  "FICTION",
		Floor: 2,
	}
	db.Model(&Shelf{}).Create(&S5)
	S6 := Shelf{
		Type:  "FANTASY",
		Floor: 2,
	}
	db.Model(&Shelf{}).Create(&S6)
	S7 := Shelf{
		Type:  "HORROR",
		Floor: 2,
	}
	db.Model(&Shelf{}).Create(&S7)

	//Book Type
	BT1 := BookType{
		Type: "COMPUTER ENGINEERING",
	}
	db.Model(&BookType{}).Create(&BT1)

	BT2 := BookType{
		Type: "ELECTRIC ENGINEERING",
	}
	db.Model(&BookType{}).Create(&BT2)

	BT3 := BookType{
		Type: "SUPERHERO FANTASY",
	}
	db.Model(&BookType{}).Create(&BT3)

	BT4 := BookType{
		Type: "HORROR FICTION",
	}
	db.Model(&BookType{}).Create(&BT4)

	BT5 := BookType{
		Type: "DARK AND GRIMDARK FANTASY",
	}
	db.Model(&BookType{}).Create(&BT5)
	BT6 := BookType{
		Type: "CONTEMPORARY FANTASY",
	}
	db.Model(&BookType{}).Create(&BT6)

	//Book
	/*db.Model(&Book{}).Create(&Book{
		Name:     "Python 1",
		Employee: sirin,
		Booktype: BT1,
		Shelf:    S2,
		Role:     R1,
		Author:   "Sirin",
		Page:     500,
		Quantity: 20,
		Price:    300,
		Date:     time.Now(),
	})
	db.Model(&Book{}).Create(&Book{
		Name:     "Java",
		Employee: Attawit,
		Booktype: BT1,
		Shelf:    S2,
		Role:     R2,
		Author:   "AJ",
		Page:     350,
		Quantity: 10,
		Price:    200,
		Date:     time.Now(),
	})*/

	python := Book{
		Name:     "Python 1",
		Employee: sirin,
		Booktype: BT1,
		Shelf:    S2,
		Role:     student,
		Author:   "Sirin",
		Page:     500,
		Quantity: 20,
		Price:    300,
		Date:     time.Now(),
	}
	db.Model(&Book{}).Create(&python)

	//-------------------------------------------------------------------------------------//xxBook//-----------------------------------------------------------------------------------//

	//-------------------------------------------------------------------------------------//xxxBorrow//---------------------------------------------------------------------------------//
	var B6111111 User
	db.Raw("SELECT * FROM Users WHERE pin = ? ", "B6111111").Scan(&B6111111) //ดึง id

	var Python Book
	db.Raw("SELECT * FROM Books WHERE name = ? ", "Python 1").Scan(&Python) //ดึง id

	db.Model(&Borrow{}).Create(&Borrow{
		Employee:  sirin,
		User:      B6111111,
		Book:      Python,
		User_Name: B6111111.Pin,
		Book_Name: Python.Name,
		Status:    "Borrow",
		DueDate:   time.Now(),
		DateTime:  time.Now(),
		Fine:      0,
	})

	//-------------------------------------------------------------------------------------//xxBorrow//---------------------------------------------------------------------------------//
	//Equipment data
	monitor := Equipment{
		Name: "จอ monitor สำหรับการนำเสนอ",
	}
	db.Model(&Equipment{}).Create(&monitor)

	printer := Equipment{
		Name: "เครื่องปริ้นท์",
	}
	db.Model(&Equipment{}).Create(&printer)

	printerMoniter := Equipment{
		Name: "เครื่องปริ้นท์ + จอ monitor สำหรับการนำเสนอ",
	}
	db.Model(&Equipment{}).Create(&printerMoniter)

	//Room_type data
	single_room := RoomType{
		Type: "ห้องเดี่ยว",
	}
	db.Model(&RoomType{}).Create(&single_room)

	group_room := RoomType{
		Type: "ห้องกลุ่ม",
	}
	db.Model(&RoomType{}).Create(&group_room)

	tutor_room := RoomType{
		Type: "ห้องสำหรับติว",
	}
	db.Model(&RoomType{}).Create(&tutor_room)

	//Research_Room Data
	// RR001 := Research_Room{
	// 	Name:      "RR001",
	// 	RoomType:  group_room,
	// 	Equipment: monitor,
	// }
	// db.Model(&Research_Room{}).Create(&RR001)

	// RR002 := Research_Room{
	// 	Name:      "RR002",
	// 	RoomType:  group_room,
	// 	Equipment: printerMoniter,
	// }
	// db.Model(&Research_Room{}).Create(&RR002)

	// RR003 := Research_Room{
	// 	Name:      "RR003",
	// 	RoomType:  tutor_room,
	// 	Equipment: printer,
	// }
	// db.Model(&Research_Room{}).Create(&RR003)

	// RR004 := Research_Room{
	// 	Name:      "RR004",
	// 	RoomType:  single_room,
	// 	Equipment: monitor,
	// }
	// db.Model(&Research_Room{}).Create(&RR004)

	//Research room
	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		Name:      "RR001",
		RoomType:  group_room,
		Equipment: monitor,
	})

	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		Name:      "RR002",
		RoomType:  group_room,
		Equipment: printerMoniter,
	})

	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		Name:      "RR003",
		RoomType:  tutor_room,
		Equipment: printer,
	})

	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		Name:      "RR004",
		RoomType:  single_room,
		Equipment: monitor,
	})

	//ดึง Data ของ User มาเก็บไว้ในตัวแปรก่อน
	var preecha User
	db.Raw("SELECT * FROM users WHERE email = ?", "preechapat@mail.com").Scan(&preecha)

	var kawin User
	db.Raw("SELECT * FROM users WHERE email = ?", "kawin@mail.com").Scan(&kawin)

	//ดึง Data ของ researchroom มาเก็บไว้ในตัวแปรก่อน
	var Room1 ResearchRoom
	db.Raw("SELECT * FROM research_rooms WHERE name = ?", "RR001").Scan(&Room1)
	var Room2 ResearchRoom
	db.Raw("SELECT * FROM research_rooms WHERE name = ?", "RR002").Scan(&Room2)
	var Room3 ResearchRoom
	db.Raw("SELECT * FROM research_rooms WHERE name = ?", "RR003").Scan(&Room3)
	var Room4 ResearchRoom
	db.Raw("SELECT * FROM research_rooms WHERE name = ?", "RR004").Scan(&Room4)

	//Addon data
	powerPlug := AddOn{
		Name: "ปลั๊กพ่วง",
	}
	db.Model(&AddOn{}).Create(&powerPlug)

	Adapter := AddOn{
		Name: "สายชาร์จ",
	}
	db.Model(&AddOn{}).Create(&Adapter)

	Pillow := AddOn{
		Name: "หมอน",
	}
	db.Model(&AddOn{}).Create(&Pillow)

	powerPlugAndAdapter := AddOn{
		Name: "ปลั๊กพ่วง + สายชาร์จ",
	}
	db.Model(&AddOn{}).Create(&powerPlugAndAdapter)

	adapterAndPillow := AddOn{
		Name: "สายชาร์จ + หมอน",
	}
	db.Model(&AddOn{}).Create(&adapterAndPillow)

	powerPlugAndAdapterAndPillow := AddOn{
		Name: "ปลั๊กพ่วง + สายชาร์จ + หมอน",
	}
	db.Model(&AddOn{}).Create(&powerPlugAndAdapterAndPillow)

	//Time data
	timeMorning := TimeRoom{
		Period: "08:00 - 12:00",
	}
	db.Model(&TimeRoom{}).Create(&timeMorning)

	timeAfternoon := TimeRoom{
		Period: "13:00 - 17:00",
	}
	db.Model(&TimeRoom{}).Create(&timeAfternoon)

	//RRRR 1
	db.Model(&ResearchRoomReservationRecord{}).Create(&ResearchRoomReservationRecord{
		ResearchRoom: Room1,
		User:         preecha,
		AddOn:        powerPlugAndAdapter,
		BookDate:     time.Now(),
		TimeRoom:     timeMorning,
	})

	//RRRR 2
	db.Model(&ResearchRoomReservationRecord{}).Create(&ResearchRoomReservationRecord{
		ResearchRoom: Room2,
		User:         kawin,
		AddOn:        powerPlugAndAdapterAndPillow,
		BookDate:     time.Now(),
		TimeRoom:     timeAfternoon,
	})

	// //
	// // ===Query
	// //

	// var target User
	// db.Model(&User{}).Find(&target, db.Where("pin = ?", "B6111111"))

	//-------------------------------------------------------------------------------------//xxReseachRoom//---------------------------------------------------------------------------------//

	//-------------------------------------------------------------------------------------//xxxBill//---------------------------------------------------------------------------------//
	
	var User1 User
	db.Raw("SELECT * FROM Users WHERE pin = ? ", "B6111111").Scan(&User1)      //ดึง id
	db.Raw("SELECT * FROM Users WHERE civ = ? ", "1111111111111").Scan(&User1) //ดึง id

	db.Model(&Bill{}).Create(&Bill{
		Book_Name:        Python.Name, //ค้นหาจาก id
		Book_Price:       uint(Python.Price),
		Employee:         sirin, //ค้นหาจาก id
		Book:             Python,
		User:             User1,                  //ค้นหาจาก id
		MemberClass_Name: User1.MemberClass.Name, //ดึงไงวะ
		Discount:         20,
		Total:            480,
		BillTime:         time.Now(),
		
	})
	//-------------------------------------------------------------------------------------//xxBill//---------------------------------------------------------------------------------//

	//-------------------------------------------------------------------------------------//xxxComReservation//---------------------------------------------------------------------//
		//Computer_os data
	comp_os_name1 := Computer_os{
		Name: "Windows",
	}
	db.Model(&Computer_os{}).Create(&comp_os_name1)

	comp_os_name2 := Computer_os{
		Name: "macOS",
	}
	db.Model(&Computer_os{}).Create(&comp_os_name2)

	//Computer data
	db.Model(&Computer{}).Create(&Computer{
		Comp_name:   "W01A",
		Comp_room:   "ROOM A",
		Computer_os: comp_os_name1,
	})

	db.Model(&Computer{}).Create(&Computer{
		Comp_name:   "W02A",
		Comp_room:   "ROOM A",
		Computer_os: comp_os_name1,
	})

	db.Model(&Computer{}).Create(&Computer{
		Comp_name:   "W03A",
		Comp_room:   "ROOM A",
		Computer_os: comp_os_name1,
	})

	db.Model(&Computer{}).Create(&Computer{
		Comp_name:   "M01B",
		Comp_room:   "ROOM B",
		Computer_os: comp_os_name2,
	})

	db.Model(&Computer{}).Create(&Computer{
		Comp_name:   "M02B",
		Comp_room:   "ROOM B",
		Computer_os: comp_os_name2,
	})

	//TIME data
	time_p1 := Time_com{
		Time_com_period: "08:00 - 09:00",
	}
	db.Model(&Time_com{}).Create(&time_p1)

	time_p2 := Time_com{
		Time_com_period: "09:00 - 10:00",
	}
	db.Model(&Time_com{}).Create(&time_p2)

	time_p3 := Time_com{
		Time_com_period: "10:00 - 11:00",
	}
	db.Model(&Time_com{}).Create(&time_p3)

	time_p4 := Time_com{
		Time_com_period: "11:00 - 12:00",
	}
	db.Model(&Time_com{}).Create(&time_p4)

	time_p5 := Time_com{
		Time_com_period: "12:00 - 13:00",
	}
	db.Model(&Time_com{}).Create(&time_p5)

	time_p6 := Time_com{
		Time_com_period: "13:00 - 14:00",
	}
	db.Model(&Time_com{}).Create(&time_p6)

	time_p7 := Time_com{
		Time_com_period: "14:00 - 15:00",
	}
	db.Model(&Time_com{}).Create(&time_p7)

	time_p8 := Time_com{
		Time_com_period: "15:00 - 16:00",
	}
	db.Model(&Time_com{}).Create(&time_p8)

	//ดึง Data ของ COMPUTER มาเก็บไว้ในตัวแปรก่อน
	// cn = comp_name ที่มาจาก COMP_NAME ใน Entity COMPUTER
	var cn1 Computer
	db.Raw("SELECT * FROM computers WHERE COMP_NAME = ?", "W01A").Scan(&cn1)
	var cn2 Computer
	db.Raw("SELECT * FROM computers WHERE COMP_NAME = ?", "W02A").Scan(&cn2)
	var cn3 Computer
	db.Raw("SELECT * FROM computers WHERE COMP_NAME = ?", "W03A").Scan(&cn3)
	var cn4 Computer
	db.Raw("SELECT * FROM computers WHERE COMP_NAME = ?", "M01B").Scan(&cn4)
	var cn5 Computer
	db.Raw("SELECT * FROM computers WHERE COMP_NAME = ?", "M02B").Scan(&cn5)

	//COMPUTER_RESERVATION
	db.Model(&Computer_reservation{}).Create(&Computer_reservation{

		Date:     time.Now(),
		Computer: cn1,
		Time_com: time_p1,
		User:     preecha,
	})
	//-------------------------------------------------------------------------------------//xxComReservation//---------------------------------------------------------------------//

	
	//-------------------------------------------------------------------------------------//xxxProblem//-------------------------------------------------------------//
	pcRdZone := Place_Class{
		name: "Reading Zone",
	}
	db.Model(&Place_Class{}).Create(&pcRdZone)

	pcTlt := Place_Class{
		name: "Toilet",
	}
	db.Model(&Place_Class{}).Create(&pcTlt)

	pcReschRoom := Place_Class{
		name: "Research Room",
	}
	db.Model(&Place_Class{}).Create(&pcReschRoom)

	pcCom := Place_Class{
		name: "Computer",
	}
	db.Model(&Place_Class{}).Create(&pcCom)

	//=========================================================================== KOOL PART PLACE
	RdZone1 := ReadingZone{
		name: "Reading Zone 1",
	}
	db.Model(&ReadingZone{}).Create(&RdZone1)
	RdZone2 := ReadingZone{
		name: "Reading Zone 2",
	}
	db.Model(&ReadingZone{}).Create(&RdZone2)
	RdZone3 := ReadingZone{
		name: "Reading Zone 3",
	}
	db.Model(&ReadingZone{}).Create(&RdZone3)

	Tlt1 := Toilet{
		name: "Toilet 1",
	}
	db.Model(&Toilet{}).Create(&Tlt1)
	Tlt2 := Toilet{
		name: "Toilet 2",
	}
	db.Model(&Toilet{}).Create(&Tlt2)
	Tlt3 := Toilet{
		name: "Toilet 3",
	}
	db.Model(&Toilet{}).Create(&Tlt3)
	Tlt4 := Toilet{
		name: "Toilet 4",
	}
	db.Model(&Toilet{}).Create(&Tlt4)
	//=========================================================================== KOOL PART PROBLEM

	probLightBulb := Problem{
		name: "Light Bulb หลอดไฟขัดข้อง",
	}
	db.Model(&Problem{}).Create(&probLightBulb)
	probAirCon := Problem{
		name: "Air Condition แอร์",
	}
	db.Model(&Problem{}).Create(&probAirCon)
	probinTlt := Problem{
		name: "in Toilet ปัญหาในห้องน้ำ",
	}
	db.Model(Problem{}).Create(&probinTlt)
	probRRdevice := Problem{
		name: "Research Room Device อุปกรณ์ห้องค้นคว้า",
	}
	db.Model(&Problem{}).Create(&probRRdevice)
	probReschCom := Problem{
		name: " Research Computer (เกี่ยวกับคอมค้นคว้า)",
	}
	db.Model(&Problem{}).Create(&probReschCom)
	probBookshelf := Problem{
		name: "Bookshelf (เกี่ยวกับชั้นวางหนังสือ)",
	}
	db.Model(&Problem{}).Create(&probBookshelf)
	probDeskChair := Problem{
		name: "Desk & Chair (เก้าอี้และโต๊ะ)",
	}
	db.Model(&Problem{}).Create(&probDeskChair)
	//=========================================================================== KOOL PART RELATION
	db.Model(&Relation{}).Create(&Relation{
		Place_Class: pcRdZone,
		Problem:     probLightBulb,
	})
	db.Model(&Relation{}).Create(&Relation{
		Place_Class: pcRdZone,
		Problem:     probAirCon,
	})
	db.Model(&Relation{}).Create(&Relation{
		Place_Class: pcRdZone,
		Problem:     probBookshelf,
	})
	db.Model(&Relation{}).Create(&Relation{
		Place_Class: pcTlt,
		Problem:     probinTlt,
	})
	db.Model(&Relation{}).Create(&Relation{
		Place_Class: pcTlt,
		Problem:     probinTlt,
	})
	db.Model(&Relation{}).Create(&Relation{
		Place_Class: pcTlt,
		Problem:     probLightBulb,
	})
	db.Model(&Relation{}).Create(&Relation{
		Place_Class: pcReschRoom,
		Problem:     probLightBulb,
	})
	db.Model(&Relation{}).Create(&Relation{
		Place_Class: pcReschRoom,
		Problem:     probAirCon,
	})
	db.Model(&Relation{}).Create(&Relation{
		Place_Class: pcReschRoom,
		Problem:     probRRdevice,
	})
	db.Model(&Relation{}).Create(&Relation{
		Place_Class: pcReschRoom,
		Problem:     probDeskChair,
	})
	db.Model(&Relation{}).Create(&Relation{
		Place_Class: pcCom,
		Problem:     probReschCom,
	})

	//=========================================================================== KOOL PART ProbReport
	db.Model(&ProblemReport{}).Create(&ProblemReport{
		User:           preecha,
		RdZone_id:      nil,
		Tlt_id:         &Tlt2.ID,
		ReschRoom_id:   nil,
		Com_id:         nil,
		Problem_ID:     &probinTlt.ID,
		Place_Class_ID: &pcTlt.ID,
		Comment:        "ห้องน้ำกดส้วมไม่ลง มุแง TT",
	})
	//-------------------------------------------------------------------------------------//xxProblem//-------------------------------------------------------------//
}
