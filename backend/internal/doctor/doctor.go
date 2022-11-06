package doctor

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/whym9/hospital/internal/admin"
)

type DoctorInfo struct {
	DateOfBirth   string  `json:"dateofbirth"`
	IIN           string  `json:"iin"`
	ID            int     `json:"id"`
	FullName      string  `json:"fullname"`
	Contactnumber string  `json:"contactnumber"`
	DepID         int     `json:"departmentID"`
	SpecID        string  `json:"specID"`
	Expirience    int     `json:"expirience"`
	PhotoLocation string  `json:"photo"`
	Category      string  `json:"category"`
	Degree        string  `json:"degree"`
	Rating        float64 `json:"rating"`
	Address       string  `json:"address"`
}

var doctors []DoctorInfo

var uploadDir = "files/"
var maxSize int64 = 200 * 1024 * 1024

func RegisterDoctor(w http.ResponseWriter, r *http.Request) {
	if !admin.Verify(r) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Verifyin error"))
		return
	}

	if err := r.ParseMultipartForm(maxSize); err != nil {
		fmt.Printf("could not parse multipart form: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("CANT_PARSE_FORM"))
		return
	}

	file, fileHeader, err := r.FormFile("uploadFile")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("INVALID_FILE"))
		return
	}
	defer file.Close()

	fileSize := fileHeader.Size

	if fileSize > maxSize {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("FILE_TOO_BIG"))
		return
	}
	fileContent, err := io.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("INVALID_FILE"))
		return
	}

	fileType := http.DetectContentType(fileContent)
	if fileType != "image/jpeg" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Incorrect format"))

		return
	}
	fileLocation := uploadDir + fileHeader.Filename
	f, err := os.Create(fileLocation)
	if err != nil {
		log.Fatal(err)
	}
	f.Write(fileContent)
	doc := DoctorInfo{}

	doc.DateOfBirth = r.FormValue("dateofbirth")
	doc.IIN = r.FormValue("iin")
	intg, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please check the formats of the forms again!"))
		fmt.Println("1")
		log.Fatal(err)
		return
	}
	doc.ID = intg
	doc.FullName = r.FormValue("fullname")
	intg, err = strconv.Atoi(r.FormValue("depID"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please check the formats of the forms again!"))
		fmt.Println("2")
		log.Fatal(err)
		return
	}
	doc.Contactnumber = r.FormValue("contactnumber")
	doc.DepID = intg
	doc.SpecID = r.FormValue("specID")
	intg, err = strconv.Atoi(r.FormValue("expirience"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please check the formats of the forms again!"))
		fmt.Println("3")
		log.Fatal(err)
		return
	}
	doc.Expirience = intg
	doc.PhotoLocation = fileLocation
	doc.Category = r.FormValue("category")
	doc.Degree = r.FormValue("degree")

	fltg, err := strconv.ParseFloat(r.FormValue("rating"), 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please check the formats of the forms again!"))
		fmt.Println("4")
		log.Fatal(err)
		return
	}
	doc.Rating = fltg
	doc.Address = r.FormValue("address")

	doctors = append(doctors, doc)

	w.Write([]byte(fmt.Sprintf("User %s has been registered successfully", doc.FullName)))
}

func GetDoctors(w http.ResponseWriter, r *http.Request) {
	if !admin.Verify(r) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Verifyin error"))
		return
	}
	// var ds struct {
	// 	doctors []int `json: doctorsID`
	// }
	// for _, d := range doctors {
	// 	ds.doctors = append(ds.doctors, d.ID)
	// }

	res, err := json.Marshal(doctors)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(res)
}

func ViewDoctor(w http.ResponseWriter, r *http.Request) {
	if !admin.Verify(r) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Verifying error"))
		return
	}
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Fatal(err)
	}
	i := findDoctor(id)
	if i == -1 {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("Given doctor does not exist"))
		return
	}

	res, err := json.Marshal(&doctors[i])
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func findDoctor(id int) int {
	for i := range doctors {
		if id == doctors[i].ID {
			return i
		}
	}
	return -1
}

func ModifyDoctor(w http.ResponseWriter, r *http.Request) {
	if !admin.Verify(r) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Verifyin error"))
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Fatal(err)
	}
	i := findDoctor(id)
	if i == -1 {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("Given doctor does not exist"))
		return
	}

	field := r.FormValue("field")

	switch field {
	case "dateofbirth":
		doctors[i].DateOfBirth = r.FormValue("modify")

	case "iin":
		doctors[i].IIN = r.FormValue("modify")
	case "id":
		x, err := strconv.Atoi(r.FormValue("modify"))
		if err != nil {
			log.Fatal(err)
		}
		doctors[i].ID = x
	case "fullname":
		doctors[i].FullName = r.FormValue("modify")
	case "depID":
		x, err := strconv.Atoi(r.FormValue("modify"))
		if err != nil {
			log.Fatal(err)
		}
		doctors[i].DepID = x

	case "specID":
		doctors[i].SpecID = r.FormValue("modify")

	case "contactnumber":
		doctors[i].Contactnumber = r.FormValue("modify")

	case "expirience":
		x, err := strconv.Atoi(r.FormValue("modify"))
		if err != nil {
			log.Fatal(err)
		}
		doctors[i].Expirience = x

	case "photo":
		location := saveFile(w, r, "modify")

		if location == "" {
			return
		}
		doctors[i].PhotoLocation = location

	case "category":
		doctors[i].Category = r.FormValue("modify")

	case "rating":
		x, err := strconv.ParseFloat(r.FormValue("modify"), 64)
		if err != nil {
			log.Fatal(err)
		}
		doctors[i].Rating = x

	case "degree":
		doctors[i].Degree = r.FormValue("modify")

	case "address":
		doctors[i].Address = r.FormValue("modify")

	}

	w.Write([]byte("Success!"))
}

func saveFile(w http.ResponseWriter, r *http.Request, name string) string {
	if err := r.ParseMultipartForm(maxSize); err != nil {
		fmt.Printf("could not parse multipart form: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("CANT_PARSE_FORM"))
		return ""
	}

	file, fileHeader, err := r.FormFile(name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("INVALID_FILE"))
		return ""
	}
	defer file.Close()

	fileSize := fileHeader.Size

	if fileSize > maxSize {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("FILE_TOO_BIG"))
		return ""
	}
	fileContent, err := io.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("INVALID_FILE"))
		return ""
	}

	fileType := http.DetectContentType(fileContent)
	if fileType != "image/jpeg" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Incorrect format"))

		return ""
	}
	fileLocation := uploadDir + fileHeader.Filename
	f, err := os.Create(fileLocation)
	if err != nil {
		log.Fatal(err)
	}
	f.Write(fileContent)
	return fileLocation
}
