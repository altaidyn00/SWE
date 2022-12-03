package admin

import (
	"encoding/json"
	"fmt"
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
