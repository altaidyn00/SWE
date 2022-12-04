package patients

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/whym9/hospital/internal/admin"
)

type PatientInfo struct {
	DateOfBirth            string `json:"date_of_birth" db:"date_of_birth"`
	IIN                    string `json:"iin" db:"iin"`
	ID                     int    `json:"id" db:"government_id"`
	BloodGroup             string `json:"blood_group" db:"blood_group"`
	EmergencyContactNumber string `json:"emergency_contact_number" db:"emergency_contact_number"`
	Contactnumber          string `json:"contact_number" db:"contact_number"`
	Address                string `json:"address" db:"address"`
	MaritalStatus          string `json:"marital_status" db:"marital_status"`
	RegistrationDate       string `json:"registration_date" db:"registration_date"`
}

type PatientReg struct {
	admin.User  `json:"user"`
	PatientInfo `json:"patient"`
}

// Admin credentials:
// {
// 	"username": "user1",
// 	"password": "pass1"
// }
// Patient login:
// Patient refister example:
// {
// 	"user": {
//        "id": 4,
//        "email": "dedee",
//        "role": "patient",
//        "first_name": "Karl",
//        "last_name": "Ben",
//        "password": "passwprd"
//     },
// 	"patient": {
//         "date_of_birth": "2012-04-23",
//         "iin": "202",
//         "id": 4,
//         "blood_group": "one",
//         "emergency_contact_number": "14",
//         "contact_number": "45",
//         "address": "dsds",
//         "marital_status": "de",
//         "registration_date": "today"
//     }
// }

func RegisterPatient(w http.ResponseWriter, r *http.Request) {

	var newPatient PatientReg
	// if _, role, ok := admin.Verify(r); !ok || role != "Admin" {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Verifyin error"))
	// 	return
	// }
	err := json.NewDecoder(r.Body).Decode(&newPatient)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = admin.DB.Exec("insert into users value(?,?,?,?,?,?)",
		newPatient.ID, newPatient.Role, newPatient.Password, newPatient.First_name, newPatient.Last_name, newPatient.Email,
	)
	if err != nil {
		panic(err)
	}

	_, err = admin.DB.Exec("insert into patient value(?,?,?,?,?,?,?,?,?)",
		newPatient.IIN, newPatient.DateOfBirth, newPatient.ID, newPatient.BloodGroup,
		newPatient.EmergencyContactNumber, newPatient.Contactnumber, newPatient.Address, newPatient.MaritalStatus,
		newPatient.RegistrationDate,
	)

	if err != nil {
		panic(err)
	}

	res, err := json.Marshal(newPatient)
	if err != nil {
		log.Fatal(err)
	}
	err = admin.DB.MustBegin().Commit()
	if err != nil {
		panic(err)
	}
	w.Write(res)
}

func GetPatients(w http.ResponseWriter, r *http.Request) {
	// if _, role, ok := admin.Verify(r); !ok || role != "Admin" {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Verifyin error"))
	// 	return
	// }
	var patients []PatientReg
	err := admin.DB.Select(&patients, "select * from users, patient where patient.government_id=users.government_id;")
	if err != nil {
		panic(err)
	}
	res, err := json.Marshal(&patients)
	fmt.Println(patients)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(res)
}

func ViewPatient(w http.ResponseWriter, r *http.Request) {
	// if _, role, ok := admin.Verify(r); !ok || role != "Admin" {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Verifyin error"))
	// 	return
	// }
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(id)
	var patient []PatientReg
	admin.DB.Select(&patient, fmt.Sprintf("select * from users, patient where users.government_id=%d and patient.government_id=users.government_id;", id))
	if len(patient) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Patient does not exist!"))
		return
	}
	res, err := json.Marshal(patient[0])
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

	rows, err := admin.DB.Query(fmt.Sprintf("update patient set date_of_birth='%s', iin='%s', blood_group='%s', emergency_contact_number='%s', contact_number=%s, address='%s', marital_status='%s', registration_date='%s' where government_id=%d",
		oldPatient.DateOfBirth, oldPatient.IIN, oldPatient.BloodGroup, oldPatient.EmergencyContactNumber, oldPatient.Contactnumber, oldPatient.Address, oldPatient.MaritalStatus, oldPatient.RegistrationDate, oldPatient.ID))
	if err != nil {
		log.Fatal(err)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	res, err := json.Marshal(oldPatient)
	if err != nil {
		log.Fatal(err)
	}
	err = admin.DB.MustBegin().Commit()
	if err != nil {
		panic(err)
	}
	w.Write(res)
}
