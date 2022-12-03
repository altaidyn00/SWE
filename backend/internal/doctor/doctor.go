package doctor

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/whym9/hospital/internal/admin"
)

type DoctorInfo struct {
	DateOfBirth   string `json:"date_of_birth" db:"date_of_birth"`
	IIN           string `json:"iin" db:"iin"`
	ID            int    `json:"id" db:"government_id"`
	Contactnumber string `json:"contact_number" db:"contact_number"`
	NumOfPatients int    `json:"number_of_patients" db:"number_of_patients"`
	Category      string `json:"category" db:"category"`
	DepID         string `json:"department_id" db:"department_id"`
	SpecID        string `json:"specialization_details_id" db:"specialization_details_id"`
	Expirience    int    `json:"experience_in_years" db:"experience_in_years"`
	Price         int    `json:"appointment_price" db:"appointment_price"`
	Degree        string `json:"education_degree" db:"education_degree"`
	Schedule      string `json:"schedule_detaile" db:"schedule_detaile"`
	Rating        int    `json:"rating" db:"rating"`
	Address       string `json:"address" db:"address"`
	PhotoLocation string `json:"photolocation" db:"photolocation"`
}

// {
// 	"user": {
//        "id": 15,
//        "email": "fewf",
//        "role": "doctor",
//        "first_name": "Taylor",
//        "last_name": "Smith",
//        "password": "password"
//     },
// 	"doctor": {
//         "date_of_birth": "2012-04-23",
//         "iin": "202",
//         "id": 15,
//         "contact_number": "45",
//         "number_of_patients": 14,
//         "category": "45",
// 		"department_id": "dswdw",
// 		"specialization_details_id": "toto",
// 		"experiance_in_years": 0,
// 		"appointment_price": 50,
// 		"education_degree": "cool",
// 		"schedule_detaile": "now",
// 		"rating": 4
//         "address": "dsds",
//         "photolocation": ":)"
//     }
// }

type DoctorReg struct {
	admin.User `json:"user"`
	DoctorInfo `json:"doctor"`
}

var doctors []DoctorInfo

var uploadDir = "files/"
var maxSize int64 = 200 * 1024 * 1024

func RegisterDoctor(w http.ResponseWriter, r *http.Request) {
	var newDoctor DoctorReg
	err := json.NewDecoder(r.Body).Decode(&newDoctor)
	if err != nil {
		panic(err)
	}
	_, err = admin.DB.Exec("insert into users value(?,?,?,?,?,?)",
		newDoctor.ID, newDoctor.Role, newDoctor.Password, newDoctor.First_name, newDoctor.Last_name, newDoctor.Email,
	)

	if err != nil {
		panic(err)
	}

	_, err = admin.DB.Exec("insert into doctor value(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		newDoctor.DateOfBirth, newDoctor.IIN, newDoctor.ID, newDoctor.Contactnumber,
		newDoctor.NumOfPatients, newDoctor.Category, newDoctor.DepID, newDoctor.SpecID,
		newDoctor.Expirience, newDoctor.Price, newDoctor.Degree, newDoctor.Schedule, newDoctor.Rating,
		newDoctor.Address, newDoctor.PhotoLocation,
	)

	if err != nil {
		panic(err)
	}

	res, err := json.Marshal(newDoctor)
	if err != nil {
		w.WriteHeader(http.StatusSeeOther)
		return
	}
	w.Write(res)
}

func GetDoctors(w http.ResponseWriter, r *http.Request) {

	var doctors []DoctorInfo
	err := admin.DB.Select(&doctors, "select * from doctor;")
	if err != nil {
		panic(err)
	}
	res, err := json.Marshal(&doctors)
	fmt.Println(doctors)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(res)
}

func ViewDoctor(w http.ResponseWriter, r *http.Request) {
	// if !admin.Verify(r) {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Verifying error"))
	// 	return
	// }

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(id)
	var doctors []DoctorInfo
	admin.DB.Select(&doctors, fmt.Sprintf("select * from doctor where government_id=%d", id))
	if len(doctors) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Doctor does not exist!"))
		return
	}
	res, err := json.Marshal(doctors[0])
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func ModifyDoctor(w http.ResponseWriter, r *http.Request) {
	var oldDoctor DoctorInfo
	// if !admin.Verify(r) {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Verifyin error"))
	// 	return
	// }

	err := json.NewDecoder(r.Body).Decode(&oldDoctor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rows, err := admin.DB.Query(fmt.Sprintf("update doctor set date_of_birth='%s', iin='%s', contact_number='%s', number_of_patients=%d, category='%s', department_id='%s', specialization_details_id='%s', experience_in_years=%d, appointment_price=%d, education_degree='%s', schedule_detaile='%s', rating=%d, address='%s', photolocation='%s' where government_id=%d",
		oldDoctor.DateOfBirth, oldDoctor.IIN, oldDoctor.Contactnumber, oldDoctor.NumOfPatients, oldDoctor.Category, oldDoctor.DepID, oldDoctor.SpecID, oldDoctor.Expirience, oldDoctor.Price, oldDoctor.Degree, oldDoctor.Schedule, oldDoctor.Rating, oldDoctor.Address, oldDoctor.PhotoLocation, oldDoctor.ID))
	if err != nil {
		log.Fatal(err)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	res, err := json.Marshal(oldDoctor)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(res)
}

func GetDoctorBySpec(w http.ResponseWriter, r *http.Request) {
	spec := r.FormValue("specialization")

	var doctors []DoctorInfo
	admin.DB.Select(&doctors, fmt.Sprintf("select * from doctor where specialization_details_id='%s'", spec))
	if len(doctors) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Doctor does not exist!"))
		return
	}
	res, err := json.Marshal(doctors[0])
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetDoctorByDep(w http.ResponseWriter, r *http.Request) {
	dep := r.FormValue("department")

	var doctors []DoctorInfo
	admin.DB.Select(&doctors, fmt.Sprintf("select * from doctor where department_id='%s'", dep))
	if len(doctors) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Doctor does not exist!"))
		return
	}
	res, err := json.Marshal(doctors[0])
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
