package patients

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var patients []PatientInfo

type PatientInfo struct {
	DateOfBirth            string `json:"dateofbirth"`
	IIN                    string `json:"iin"`
	ID                     string `json:"id"`
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
	// if _, role, ok := admin.Verify(r); !ok || role != "Admin" {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Verifyin error"))
	// 	return
	// }
	err := json.NewDecoder(r.Body).Decode(&newPatient)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	patients = append(patients, newPatient)

	res, err := json.Marshal(newPatient)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(res)
}

func GetPatients(w http.ResponseWriter, r *http.Request) {
	// if _, role, ok := admin.Verify(r); !ok || role != "Admin" {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Verifyin error"))
	// 	return
	// }

	res, err := json.Marshal(patients)

	if err != nil {
		log.Fatal(err)
	}
	w.Write(res)
}

func findPatient(id string) int {
	for i := range patients {
		if id == patients[i].ID {
			return i
		}
	}
	return -1
}

func ViewPatient(w http.ResponseWriter, r *http.Request) {
	// if _, role, ok := admin.Verify(r); !ok || role != "Admin" {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Verifyin error"))
	// 	return
	// }
	id := r.FormValue("id")

	i := findPatient(id)
	if i == -1 {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("Given patient does not exist"))
		return
	}

	res, err := json.Marshal(patients[i])
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func ModifyPatient(w http.ResponseWriter, r *http.Request) {
	var oldPatient PatientInfo
	// if _, role, ok := admin.Verify(r); !ok || role != "Admin" {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Verifyin error"))
	// 	return
	// }
	err := json.NewDecoder(r.Body).Decode(&oldPatient)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := oldPatient.ID

	i := findPatient(id)
	if i == -1 {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("Given patient does not exist"))
		return
	}

	patients[i] = oldPatient
	res, err := json.Marshal(patients[i])
	if err != nil {
		log.Fatal(err)
	}
	w.Write(res)
}
