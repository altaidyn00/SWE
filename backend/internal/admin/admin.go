package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type User struct {
	id         int    `json:"id"`
	email      string `json:"email"`
	role       string `json:"role"`
	first_name string `json:"first_name"`
	last_name  string `json:"last_name"`
}

var All_users = []User{
	{
		id:         0,
		first_name: "Admin",
		last_name:  "O",
		email:      "some@some.kz",
		role:       "Admin",
	},
}

func Find_user(name string) (User, bool) {
	for _, u := range All_users {
		if u.first_name == name {
			return u, true
		}
	}
	return User{}, false
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
	Name           string
	Value          string
	ExpirationTime time.Time
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	creds, ok := users[credentials.UserName]
	expectedpassword := creds[0]
	if !ok || expectedpassword != credentials.Password {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 30)

	claims := &Claims{
		UserName: credentials.UserName,
		Role:     creds[1],
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

	json.NewEncoder(w).Encode(Token{
		"Admin_Token",
		tokenString,
		expirationTime,
	})
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
