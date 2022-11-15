package doctor

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
	//PhotoLocation string  `json:"photo"`
	Category      string  `json:"category"`
	Degree        string  `json:"degree"`
	Rating        string `json:"rating"`
	Address       string  `json:"address"`
}

var doctors []DoctorInfo

var uploadDir = "files/"
var maxSize int64 = 200 * 1024 * 1024

func RegisterDoctor(w http.ResponseWriter, r *http.Request) {
	// if !admin.Verify(r) {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Verifyin error"))
	// 	return
	// }

// 	if err := r.ParseMultipartForm(maxSize); err != nil {
// 		fmt.Printf("could not parse multipart form: %v\n", err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("CANT_PARSE_FORM"))
// 		return
// 	}

// 	file, fileHeader, err := r.FormFile("photo")
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("INVALID_FILE"))
// 		return
// 	}
// 	defer file.Close()

// 	fileSize := fileHeader.Size

// 	if fileSize > maxSize {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("FILE_TOO_BIG"))
// 		return
// 	}
// 	fileContent, err := io.ReadAll(file)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("INVALID_FILE"))
// 		return
// 	}

// 	fileLocation := uploadDir + fileHeader.Filename
// 	f, err := os.Create(fileLocation)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	f.Write(fileContent)
	doc := DoctorInfo{}

	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	doctors = append(doctors, doc)

	res, err := json.Marshal(doc)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(res)
}

func GetDoctors(w http.ResponseWriter, r *http.Request) {
	// if !admin.Verify(r) {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Verifyin error"))
	// 	return
	// }

	res, err := json.Marshal(doctors)
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
	id := oldDoctor.ID

	i := findDoctor(id)
	if i == -1 {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("Given patient does not exist"))
		return
	}

	doctors[i] = oldDoctor
	res, err := json.Marshal(doctors[i])
	if err != nil {
		log.Fatal(err)
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
