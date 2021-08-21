package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/amirghedira/go-rest-api/routes"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var passwordCrypted = "$2a$11$lR1UEiV4t.J3KrKyLN24j.707QNMy25KS7EXbanAHy5Di2K2bp7EW"

func testRequestHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	println(params["password"])
	bytes, err := bcrypt.GenerateFromPassword([]byte(params["password"]), 11)
	fmt.Println(string(bytes))
	if err != nil {

		fmt.Println(string(bytes))
	}
}
func CheckPasswordHash(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	println(params["password"])
	err := bcrypt.CompareHashAndPassword([]byte(passwordCrypted), []byte(params["password"]))
	fmt.Println(err == nil)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

func GenerateToken(w http.ResponseWriter, r *http.Request) {

	var hmacSampleSecret []byte
	secretString, _ := GenerateRandomString(1000)
	hmacSampleSecret = []byte(secretString)
	// println(secretString)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"foo": "bar"})

	// Sign and get the complete encoded token as a string using the secret
	token.SignedString(hmacSampleSecret)

	// fmt.Println(tokenString)
}

func main() {

	app := mux.NewRouter()
	app.HandleFunc("/hash/{password}", testRequestHandler).Methods("GET")
	app.HandleFunc("/check/{password}", CheckPasswordHash).Methods("GET")
	app.Handle("/token", loggingMiddleware(http.HandlerFunc(GenerateToken))).Methods("GET")

	routes.BookApi(app.PathPrefix("/book").Subrouter())
	log.Fatal(http.ListenAndServe(":5000", app))
}
