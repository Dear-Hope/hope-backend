package main

import (
	"HOPE-backend/v1/ambulance"
	_ambulanceHandler "HOPE-backend/v1/ambulance/handler"
	_ambulanceRepo "HOPE-backend/v1/ambulance/repository"
	_chatHandler "HOPE-backend/v1/chat/handler"
	"HOPE-backend/v1/hospital"
	_hospitalHandler "HOPE-backend/v1/hospital/handler"
	_hospitalRepo "HOPE-backend/v1/hospital/repository"
	"HOPE-backend/v1/laboratory"
	_laboratoryHandler "HOPE-backend/v1/laboratory/handler"
	_laboratoryRepo "HOPE-backend/v1/laboratory/repository"
	"HOPE-backend/v1/medicine"
	_medicineHandler "HOPE-backend/v1/medicine/handler"
	_medicineRepo "HOPE-backend/v1/medicine/repository"
	"HOPE-backend/v1/newsletter"
	_newsletterHandler "HOPE-backend/v1/newsletter/handler"
	_newsletterRepo "HOPE-backend/v1/newsletter/repository"
	"HOPE-backend/v2/services/auth"
	_authHandler "HOPE-backend/v2/services/auth/handler"
	_authRepo "HOPE-backend/v2/services/auth/repository"
	"HOPE-backend/v2/services/moodtracker"
	_moodHandler "HOPE-backend/v2/services/moodtracker/handler"
	_moodRepo "HOPE-backend/v2/services/moodtracker/repository"
	"HOPE-backend/v2/services/selfcare"
	_selfCareHandler "HOPE-backend/v2/services/selfcare/handler"
	_selfCareRepo "HOPE-backend/v2/services/selfcare/repository"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	dbV2 "HOPE-backend/v2/db"

	sendblue "github.com/sendinblue/APIv3-go-library/lib"
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
	router.Use(cors.Default())

	db := NewPostgreSQLDatabase()
	db2 := dbV2.NewPostgreSQLDatabase()

	router.GET("/server/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Server is OK!")
	})

	v1 := router.Group("/api/v1")
	v2 := router.Group("/api/v2")

	mailer := sendblue.NewAPIClient(sendblue.NewConfiguration())

	authRepo := _authRepo.NewPostgreSQLRepository(db2)
	authSvc := auth.NewAuthService(authRepo, mailer)
	_authHandler.NewAuthHandler(v2, authSvc)

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

	pool := _chatHandler.NewPool()
	go pool.Start()

	// chatRepo := _chatRepo.NewPostgreSQLRepository(db)
	// chatSvc := chat.NewChatService(chatRepo, authRepo)
	// _chatHandler.NewChatHandler(v1, chatSvc, upgrader, pool)

	// recordRepo := _recordRepo.NewPostgreSQLRepository(db)
	// recordSvc := psychologicalrecord.NewPsychologicalRecordService(recordRepo, authRepo)
	// _recordHandler.NewPsychologicalRecordHandler(v1, recordSvc)

	moodRepo := _moodRepo.NewPostgreSQLRepository(db2)
	moodSvc := moodtracker.NewMoodTrackerService(moodRepo, authRepo)
	_moodHandler.NewMoodTrackerHandler(v2, moodSvc)

	selfCareRepo := _selfCareRepo.NewPostgreSQLRepository(db2)
	selfCareSvc := selfcare.NewSelfCareService(selfCareRepo)
	_selfCareHandler.NewSelfCareHandler(v2, selfCareSvc)

	// err := mailchimp.SetKey("eb5431057e55a836f23671a6c07c7643-us14")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	newsletterRepo := _newsletterRepo.NewPostgreSQLRepository(db)
	newsletterSvc := newsletter.NewNewsletterService(newsletterRepo)
	_newsletterHandler.NewNewsletterService(v1, newsletterSvc)

	router.Run(":8000")
}
