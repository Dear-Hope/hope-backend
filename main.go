package main

import (
	"HOPE-backend/ambulance"
	_ambulanceHandler "HOPE-backend/ambulance/handler"
	_ambulanceRepo "HOPE-backend/ambulance/repository"
	"HOPE-backend/auth"
	_authHandler "HOPE-backend/auth/handler"
	_authRepo "HOPE-backend/auth/repository"
	"HOPE-backend/hospital"
	_hospitalHandler "HOPE-backend/hospital/handler"
	_hospitalRepo "HOPE-backend/hospital/repository"
	"HOPE-backend/laboratory"
	_laboratoryHandler "HOPE-backend/laboratory/handler"
	_laboratoryRepo "HOPE-backend/laboratory/repository"
	"HOPE-backend/medicine"
	_medicineHandler "HOPE-backend/medicine/handler"
	_medicineRepo "HOPE-backend/medicine/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := NewPostgreSQLDatabase()

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

	router.Run(":80")
}
