package doctor

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
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
	DepID         int    `json:"department_id" db:"department_id"`
	SpecID        int    `json:"specialization_id" db:"specialization_id"`
	Expirience    int    `json:"experience_in_years" db:"experience_in_years"`
	Price         int    `json:"appointment_price" db:"appointment_price"`
	Degree        string `json:"education_degree" db:"education_degree"`
	Schedule      string `json:"schedule_detail" db:"schedule_detail"`
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
// 		"department_id": 2,
// 		"specialization_id": 8,
// 		"experience_in_years": 0,
// 		"appointment_price": 50,
// 		"education_degree": "cool",
// 		"schedule_detail": "now",
// 		"rating": 4,
//         "address": "dsds",
//         "photolocation": ":)"
//     }
// }

type DoctorReg struct {
	admin.User `json:"user"`
	DoctorInfo `json:"doctor"`
}

func RegisterDoctor(w http.ResponseWriter, r *http.Request) {
	var newDoctor DoctorReg

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		panic(err)
	}

	_, err = admin.DB.Exec("insert into users value(?,?,?,?,?,?)",
		id, r.FormValue("role"), r.FormValue("password"), r.FormValue("first_name"), r.FormValue("last_name"), r.FormValue("email"),
	)
	if err != nil {
		panic(err)
	}

	numPat, err := strconv.Atoi(r.FormValue("number_of_patients"))
	if err != nil {
		panic(err)
	}

	depID, err := strconv.Atoi(r.FormValue("department_id"))
	if err != nil {
		panic(err)
	}

	specID, err := strconv.Atoi(r.FormValue("specialization_id"))
	if err != nil {
		panic(err)
	}
	price, err := strconv.Atoi(r.FormValue("appointment_price"))
	if err != nil {
		panic(err)
	}
	rating, err := strconv.Atoi(r.FormValue("rating"))
	if err != nil {
		panic(err)
	}
	experience, err := strconv.Atoi(r.FormValue("experience_in_years"))
	if err != nil {
		panic(err)
	}

	newDoctor.ID = id
	newDoctor.Role = r.FormValue("role")
	newDoctor.Password = r.FormValue("password")
	newDoctor.First_name = r.FormValue("first_name")
	newDoctor.Last_name = r.FormValue("last_name")
	newDoctor.Email = r.FormValue("email")
	newDoctor.DateOfBirth = r.FormValue("date_of_birth")
	newDoctor.IIN = r.FormValue("iin")
	newDoctor.Contactnumber = r.FormValue("contact_number")
	newDoctor.Category = r.FormValue("category")
	newDoctor.DepID = depID
	newDoctor.SpecID = specID
	newDoctor.NumOfPatients = numPat
	newDoctor.Expirience = experience
	newDoctor.Price = price
	newDoctor.Rating = rating
	newDoctor.Degree = r.FormValue("education_degree")
	newDoctor.Schedule = r.FormValue("schedule_detail")
	newDoctor.Address = r.FormValue("address")

	photo, err := GetPhoto(r)
	if err != nil {
		w.WriteHeader(http.StatusSeeOther)
		w.Write([]byte("Could not save photo"))
		return
	}
	newDoctor.PhotoLocation = photo
	_, err = admin.DB.Exec("insert into doctor value(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		r.FormValue("date_of_birth"), r.FormValue("iin"), id, r.FormValue("contact_number"),
		numPat, r.FormValue("category"), depID, specID,
		experience, price, r.FormValue("education_degree"), r.FormValue("schedule_detail"), rating,
		r.FormValue("address"), photo,
	)

	if err != nil {
		panic(err)
	}
	res, err := json.Marshal(newDoctor)
	if err != nil {
		w.WriteHeader(http.StatusSeeOther)
		return
	}
	err = admin.DB.MustBegin().Commit()
	if err != nil {
		panic(err)
	}
	w.Write(res)
}

func GetPhoto(r *http.Request) (string, error) {

	if err := r.ParseMultipartForm(200 * 1024 * 1024); err != nil {
		fmt.Printf("could not parse multipart form: %v\n", err)
		return "", err
	}

	file, fileHeader, err := r.FormFile("photo")
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	fileName := fileHeader.Filename

	newPath := filepath.Join("../files", fileName)
	fmt.Printf("File: %s\n", newPath)

	err = saveFile(fileContent, newPath)
	if err != nil {
		return "", err
	}
	return "SWE/files/" + fileName, nil

}

func GetDoctors(w http.ResponseWriter, _ *http.Request) {

	var doctors []DoctorReg
	err := admin.DB.Select(&doctors, "select * from users, doctor where doctor.government_id=users.government_id;")
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
	var doctors []DoctorReg
	admin.DB.Select(&doctors, fmt.Sprintf("select * from users, doctor where users.government_id=%d and doctor.government_id=users.government_id", id))
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

	rows, err := admin.DB.Query(fmt.Sprintf("update doctor set date_of_birth='%s', iin='%s', contact_number='%s', number_of_patients=%d, category='%s', department_id=%d, specialization_id=%d, experience_in_years=%d, appointment_price=%d, education_degree='%s', schedule_detail='%s', rating=%d, address='%s', photolocation='%s' where government_id=%d",
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
	err = admin.DB.MustBegin().Commit()
	if err != nil {
		panic(err)
	}
	w.Write(res)
}

func GetDoctorBySpec(w http.ResponseWriter, r *http.Request) {
	spec, err := strconv.Atoi(r.FormValue("specialization"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Doctor does not exist!"))
		return
	}
	var doctors []DoctorReg
	admin.DB.Select(&doctors, fmt.Sprintf("select * from users, doctor where doctor.specialization_details_id=%d and doctor.government_id=users.government_id;", spec))
	if len(doctors) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Doctor does not exist!"))
		return
	}
	res, err := json.Marshal(doctors)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetDoctorByDep(w http.ResponseWriter, r *http.Request) {
	dep, err := strconv.Atoi(r.FormValue("department"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Doctor does not exist!"))
		return
	}

	var doctors []DoctorReg
	admin.DB.Select(&doctors, fmt.Sprintf("select * from users, doctor where doctor.department_id=%d and doctor.government_id=users.government_id;", dep))
	if len(doctors) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Doctor does not exist!"))
		return
	}
	res, err := json.Marshal(doctors)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetDoctorByName(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	var doctors []DoctorReg
	admin.DB.Select(&doctors, fmt.Sprintf("select * from users, doctor where users.first_name='%s' and doctor.government_id=users.government_id;", name))
	if len(doctors) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Doctor does not exist!"))
		return
	}
	res, err := json.Marshal(doctors)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func saveFile(content []byte, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("couldn't create file")
	}

	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return errors.New("counldn't write to file")
	}
	return nil
}
