package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/whym9/hospital/internal/admin"
	"github.com/whym9/hospital/internal/doctor"
	"github.com/whym9/hospital/internal/patients"
)

func StartServer(addr string) {
	http.HandleFunc("/", Get)
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

func Get(w http.ResponseWriter, r *http.Request) {
	name, _, ok := admin.Verify(r)
	if !ok {
		fmt.Println("1")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var user admin.User
	user, ok = admin.Find_user(name)
	if !ok {
		fmt.Println(name)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(user)
	res, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	return
}
