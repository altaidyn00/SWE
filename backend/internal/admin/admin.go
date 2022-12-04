package admin

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Id         int    `json:"id" db:"government_id"`
	Email      string `json:"email" db:"email"`
	Role       string `json:"role" db:"role"`
	First_name string `json:"first_name" db:"first_name"`
	Last_name  string `json:"last_name" db:"last_name"`
	Password   string `json:"password" db:"password"`
}

var DB *sqlx.DB

func Register(dsn string) {
	db, err := sqlx.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}
	DB = db

}

var All_users = []User{
	{
		Id:         0,
		First_name: "user1",
		Last_name:  "O",
		Email:      "some@some.kz",
		Role:       "Admin",
		Password:   "pass1",
	},
	{
		Id:         2,
		First_name: "user2",
		Last_name:  "_",
		Email:      "dw@sw.pop",
		Role:       "Patient",
		Password:   "pass2",
	},
}

func Find_user(name string) ([]User, bool) {
	var user []User
	DB.Select(&user, fmt.Sprintf("select * from users where first_name='%s';", name))

	if len(user) == 0 {
		return []User{}, false
	}

	return user, true
}

var jwt_key = []byte("secret_key")
var users = map[string][]string{
	"user1": {"pass1", "Admin"},
}

type Credentials struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	UserName string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

type Token struct {
	Name           string    `json:"name"`
	Value          string    `json:"value"`
	ExpirationTime time.Time `json:"expire"`
}

type Answer struct {
	T Token `json:"Acces_token"`
	U User  `json:"user"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users, ok := Find_user(credentials.UserName)

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var verify bool = false
	// creds, ok := users[credentials.UserName]
	var user User
	for _, u := range users {
		if credentials.Password == u.Password {
			verify = true
			user = u
		}
	}
	if !verify {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 30)

	claims := &Claims{
		UserName: credentials.UserName,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwt_key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	a := Answer{
		Token{
			"Admin_Token",
			tokenString,
			expirationTime,
		},
		user,
	}
	json.NewEncoder(w).Encode(a)
	// http.SetCookie(w,
	// 	&http.Cookie{
	// 		Name:    "token",
	// 		Value:   tokenString,
	// 		Expires: expirationTime,
	// 	},
	// )
	w.WriteHeader(http.StatusOK)
	//w.Write([]byte("Login completed successfully"))
}

func Home(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenString := cookie.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwt_key, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if claims.Role != "Admin" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Admin role reuired!"))
		return
	}

	w.Write([]byte(fmt.Sprintf("Hello user %s", claims.UserName)))
}

func Verify(r *http.Request) (string, string, bool) {
	var body Token
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return "", "", false
	}
	tokenString := body.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwt_key, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {

			return "", "", false
		}

		return "", "", false
	}

	if !tkn.Valid {

		return "", "", false
	}

	return claims.UserName, claims.Role, true
}

type Appointment struct {
	Doctor_id int    `json:"doctor_id" db:"doctor_id"`
	Date      string `json:"preferred_date" db:"preferred_date"`
	Time      string `json:"preferred_time" db:"preferred_time"`
	Name      string `json:"name" db:"name"`
	Surname   string `json:"surname" db:"surname"`
	Email     string `json:"email" db:"email"`
}

// {
// 	"doctor_id": 0
// 	"preferred_date": ""
// 	"preferred_time": ""
// 	"name": ""
// 	"suname": ""
// 	"email": ""
// }

func AppointmentReg(w http.ResponseWriter, r *http.Request) {
	var appointment Appointment
	// if _, role, ok := admin.Verify(r); !ok || role != "Admin" {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Verifyin error"))
	// 	return
	// }
	err := json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = DB.Exec("insert into appointment value(?,?,?,?,?,?)",
		appointment.Doctor_id, appointment.Date, appointment.Time, appointment.Name, appointment.Surname, appointment.Email,
	)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(appointment)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.MustBegin().Commit()
	if err != nil {
		panic(err)
	}
	w.Write(res)
}

func GetAppointments(w http.ResponseWriter, r *http.Request) {
	var appointments []Appointment
	err := DB.Select(&appointments, "select * from appointment;")
	if err != nil {
		panic(err)
	}
	res, err := json.Marshal(&appointments)
	fmt.Println(appointments)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(res)
}

type Spec struct {
	Id   int    `json:"id" db:"id"`
	Desc string `json:"description" db:"descript"`
}

type Dep struct {
	Id   int    `json:"id" db:"id"`
	Desc string `json:"description" db:"descript"`
}

func GetSpecializations(w http.ResponseWriter, r *http.Request) {
	var specs []Spec
	err := DB.Select(&specs, "select * from specialization;")
	if err != nil {
		panic(err)
	}
	res, err := json.Marshal(&specs)

	if err != nil {
		log.Fatal(err)
	}
	w.Write(res)

}

func GetDepartments(w http.ResponseWriter, r *http.Request) {
	var deps []Dep
	err := DB.Select(&deps, "select * from department;")
	if err != nil {
		panic(err)
	}
	res, err := json.Marshal(&deps)

	if err != nil {
		log.Fatal(err)
	}
	w.Write(res)

}
