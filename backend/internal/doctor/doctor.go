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
	ID            string  `json:"id"`
	FullName      string  `json:"fullname"`
	Contactnumber string  `json:"contactnumber"`
	DepID         string  `json:"departmentID"`
	SpecID        string  `json:"specID"`
	Expirience    string  `json:"expirience"`
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

	file, fileHeader, err := r.FormFile("photo")
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

	fileLocation := uploadDir + fileHeader.Filename
	f, err := os.Create(fileLocation)
	if err != nil {
		log.Fatal(err)
	}
	f.Write(fileContent)
	doc := DoctorInfo{}

	doc.DateOfBirth = r.FormValue("dateofbirth")
	doc.IIN = r.FormValue("iin")
	doc.ID = r.FormValue("id")
	doc.FullName = r.FormValue("fullname")

	doc.Contactnumber = r.FormValue("contactnumber")
	doc.DepID = r.FormValue("depID")
	doc.SpecID = r.FormValue("specID")
	doc.Expirience = r.FormValue("expirience")
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

	res, err := json.Marshal(doc)
	if err != nil {
		w.WriteHeader(http.StatusSeeOther)
		return
	}
	w.Write(res)
}

func GetDoctors(w http.ResponseWriter, r *http.Request) {
	if !admin.Verify(r) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Verifyin error"))
		return
	}

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
	id := r.FormValue("id")
	i := findDoctor(id)
	if i == -1 {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("Given doctor does not exist"))
		return
	}

	res, err := json.Marshal(doctors[i])
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func findDoctor(id string) int {
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

	id := r.FormValue("id")

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
		doctors[i].ID = r.FormValue("modify")
	case "fullname":
		doctors[i].FullName = r.FormValue("modify")
	case "depID":
		doctors[i].DepID = r.FormValue("modify")

	case "specID":
		doctors[i].SpecID = r.FormValue("modify")

	case "contactnumber":
		doctors[i].Contactnumber = r.FormValue("modify")

	case "expirience":
		doctors[i].Expirience = r.FormValue("modify")

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
	res, err := json.Marshal(doctors[i])
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.Write(res)
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

	fileLocation := uploadDir + fileHeader.Filename
	f, err := os.Create(fileLocation)
	if err != nil {
		log.Fatal(err)
	}
	f.Write(fileContent)
	return fileLocation
}
