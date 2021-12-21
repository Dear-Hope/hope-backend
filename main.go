package main

import (
	"HOPE-backend/ambulance"
	_ambulanceHandler "HOPE-backend/ambulance/handler"
	_ambulanceRepo "HOPE-backend/ambulance/repository"
	"HOPE-backend/auth"
	_authHandler "HOPE-backend/auth/handler"
	_authRepo "HOPE-backend/auth/repository"
	"HOPE-backend/chat"
	_chatHandler "HOPE-backend/chat/handler"
	_chatRepo "HOPE-backend/chat/repository"
	"HOPE-backend/hospital"
	_hospitalHandler "HOPE-backend/hospital/handler"
	_hospitalRepo "HOPE-backend/hospital/repository"
	"HOPE-backend/laboratory"
	_laboratoryHandler "HOPE-backend/laboratory/handler"
	_laboratoryRepo "HOPE-backend/laboratory/repository"
	"HOPE-backend/medicine"
	_medicineHandler "HOPE-backend/medicine/handler"
	_medicineRepo "HOPE-backend/medicine/repository"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// We'll need to check the origin of our connection
	// this will allow us to make requests from our React
	// development server to here.
	// For now, we'll do no checking and just allow any connection
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("X-Requested-With", "Accept", "X-XSRF-TOKEN", "Authorization")
	router.Use(cors.New(config))

	db := NewPostgreSQLDatabase()

	router.GET("/server/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Server is OK!")
	})

	v1 := router.Group("/api/v1")

	authRepo := _authRepo.NewPostgreSQLRepository(db)
	authSvc := auth.NewAuthService(authRepo)
	_authHandler.NewAuthHandler(v1, authSvc)

	medicineRepo := _medicineRepo.NewPostgreSQLRepository(db)
	medicineSvc := medicine.NewMedicineService(medicineRepo)
	_medicineHandler.NewMedicineHandler(v1, medicineSvc)

	ambulanceRepo := _ambulanceRepo.NewPostgreSQLRepository(db)
	ambulanceSvc := ambulance.NewAmbulanceService(ambulanceRepo)
	_ambulanceHandler.NewAmbulanceHandler(v1, ambulanceSvc)

	hospitalRepo := _hospitalRepo.NewPostgreSQLRepository(db)
	hospitalSvc := hospital.NewHospitalService(hospitalRepo)
	_hospitalHandler.NewHospitalHandler(v1, hospitalSvc)

	laboratoryRepo := _laboratoryRepo.NewPostgreSQLRepository(db)
	laboratorySvc := laboratory.NewLaboratoryService(laboratoryRepo)
	_laboratoryHandler.NewLaboratoryHandler(v1, laboratorySvc)

	chatRepo := _chatRepo.NewPostgreSQLRepository(db)
	chatSvc := chat.NewChatService(chatRepo, authRepo)
	_chatHandler.NewChatHandler(v1, chatSvc, upgrader)

	router.Run(":8000")
}
