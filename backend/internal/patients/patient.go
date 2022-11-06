package patients

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	//"github.com/whym9/hospital/internal/admin"
)

var patients []PatientInfo

type PatientInfo struct {
	DateOfBirth            string `json:"dateofbirth"`
	IIN                    string `json:"iin"`
	ID                     int    `json:"id"`
	FullName               string `json:"fullname"`
	BloodGroup             string `json:"blooodgroup"`
	EmergencyContactNumber string `json:"emergencynumber"`
	Contactnumber          string `json:"contactnumber"`
	Email                  string `json:"email"`
	Address                string `json:"address"`
	MaritalStatus          string `json:"martialstatus"`
	RegistrationDate       string `json:"registrationdate"`
}

// Admin credentials:
// {
// 	"username": "user1",
// 	"password": "pass1"
// }
// Patient refister example:
// {
//     "dateofbirth": "1.01.2020",
//     "iin": 2,
//     "id": 0,
//     "fullname": ["some", "some"],
//     "bloodgroup": "one",
//     "emergencynumber": 14,
//     "contactnumber": 45,
//     "email": "email",
//     "address": "dsds",
//     "maritalstatus": "de",
//     "registrationdate": "today"
// }

func RegisterPatient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Patients")
	var newPatient PatientInfo
	// if !admin.Verify(r) {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Verifyin error"))
	// 	return
	// }
	err := json.NewDecoder(r.Body).Decode(&newPatient)
	if err != nil {
		fmt.Println("Patients")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	patients = append(patients, newPatient)

	w.Write([]byte(fmt.Sprintf("User %s has been registered successfully", newPatient.FullName)))
}

func GetPatients(w http.ResponseWriter, r *http.Request) {
	// if !admin.Verify(r) {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Verifyin error"))
	// 	return
	// }
	// var ps struct {
	// 	patients []int `json: patientsID`
	// }
	// pats := []int
	// for _, d := range patients {
	// 	ps.patients = append(ps.patients, d.ID)
	// }

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	fmt.Println(patients)
	w.Write(res)
}

func findPatient(id int) int {
	for i, _ := range patients {
		if id == patients[i].ID {
			return i
		}
	}
	return -1
}

func ViewPatient(w http.ResponseWriter, r *http.Request) {
	// if !admin.Verify(r) {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Verifyin error"))
	// 	return
	// }
	id, err := strconv.Atoi(r.FormValue("ID"))
	if err != nil {
		log.Fatal(err)
	}
	i := findPatient(id)
	if i == -1 {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("Given patient does not exist"))
		return
	}

	res, err := json.Marshal(&patients[i])
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func ModifyPatient(w http.ResponseWriter, r *http.Request) {
	// if !admin.Verify(r) {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Verifyin error"))
	// 	return
	// }

	id, err := strconv.Atoi(r.FormValue("ID"))
	if err != nil {
		log.Fatal(err)
	}
	i := findPatient(id)
	if i == -1 {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("Given patient does not exist"))
		return
	}

	field := r.FormValue("field")

	switch field {
	case "dateofbirth":
		patients[i].DateOfBirth = r.FormValue("modify")

	case "iin":
		patients[i].IIN = r.FormValue("modify")
	case "id":
		x, err := strconv.Atoi(r.FormValue("modify"))
		if err != nil {
			log.Fatal(err)
		}
		patients[i].ID = x
	case "fullname":
		patients[i].FullName = r.FormValue("modify")
	case "blooodgroup":
		patients[i].BloodGroup = r.FormValue("modify")

	case "emergencynumber":
		patients[i].EmergencyContactNumber = r.FormValue("modify")

	case "contactnumber":
		patients[i].Contactnumber = r.FormValue("modify")

	case "email":
		patients[i].Email = r.FormValue("modify")

	case "address":
		patients[i].Address = r.FormValue("modify")

	case "martialstatus":
		patients[i].MaritalStatus = r.FormValue("modify")

	case "registrationdate":
		patients[i].RegistrationDate = r.FormValue("modify")

	}

	w.Write([]byte("Success!"))
}
