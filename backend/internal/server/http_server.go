package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/whym9/hospital/internal/admin"
	"github.com/whym9/hospital/internal/doctor"
	"github.com/whym9/hospital/internal/patients"
)

func StartServer(addr string) {
	http.HandleFunc("/login", admin.Login)
	http.HandleFunc("/registerDoctor", doctor.RegisterDoctor)
	http.HandleFunc("/registerPatient", patients.RegisterPatient)
	http.HandleFunc("/viewPatientDetails", patients.ViewPatient)
	http.HandleFunc("/viewDoctorDetails", doctor.ViewDoctor)
	http.HandleFunc("/modifyPatientInfo", patients.ModifyPatient)
	http.HandleFunc("/modifyDoctorInfo", doctor.ModifyDoctor)
	http.HandleFunc("/getDoctors", doctor.GetDoctors)
	http.HandleFunc("/getPatients", patients.GetPatients)
	fmt.Println("Server started on localhost:8080")
	log.Fatal(http.ListenAndServe(addr, nil))
}
