package handler

import (
	"Movie-and-events/hash"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	Rtoken "github.com/flutter-amp/baking-api/baking/rtoken"
	"github.com/flutter-amp/baking-api/entity"
	"github.com/flutter-amp/baking-api/user"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	// "github.com/flutter-amp/Baking-API/form"
	// "github.com/flutter-amp/Baking-API/model"
	// "github.com/flutter-amp/Baking-API/rtoken"
	// "github.com/flutter-amp/Baking-API/session"
)

type UserHandler struct {
	UserService user.UserService
}

func NewUserHandler(us user.UserService) *UserHandler {
	return &UserHandler{UserService: us}
}

//func (uh *UserHandler) Authenticated(next http.HandlerFunc) http.HandlerFunc {
// 	// validate the token
// 	fn := func(w http.ResponseWriter, r *http.Request) {
// 		_token := r.Header.Get("Authorization")
// 		_token = strings.Replace(_token, "Bearer ", "", 1)
// 		valid, err := uh.tokenService.ValidateToken(_token)
// 		if err != nil && !valid {
// 			responses.ERROR(w, http.StatusUnauthorized, errors.New("unauthenticated: unauthorized to access the resource, log in again"))
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	}
// 	return http.HandlerFunc(fn)

// }

func (uh *UserHandler) SignUp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("user handelr")

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	user := &entity.User{}
	fmt.Println("in post user 2")

	err := json.Unmarshal(body, user)
	fmt.Println(user)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	pass, err2 := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err2 != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	user.Password = string(pass)
	user, errs := uh.UserService.StoreUser(user)

	if len(errs) > 0 {
		fmt.Println("HEEEEEEEEEEEEEEEE")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, _ := json.MarshalIndent(user, "", "\t\t")
	// p := fmt.Sprintf("/api/recipe/%d", recipe.ID)
	// w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	w.Write(output)
	return

}
func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("user handelr")

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	user := &entity.User{}
	fmt.Println("in post user 2")

	err := json.Unmarshal(body, user)
	fmt.Println(user)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user1, errs := uh.UserService.UserByEmail(user.Email)
	fmt.Println("sencond fffffffffffffff")
	fmt.Println(user1)
	if len(errs) > 0 || !hash.ArePasswordsSame(user1.Password, user.Password) {
		fmt.Println("HEEEEEEEEEEEEEEEE")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

		return
	}
	tokenString, err := Rtoken.GenerateJwtToken([]byte(Rtoken.GenerateRandomID(32)), Rtoken.CustomClaims{
		SessionId: "adonaTesfaye",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 1, 1).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	})
	output, _ := json.MarshalIndent(struct {
		Token string `json:"token" `
	}{
		Token: tokenString,
	}, "", "\t\t")
	// p := fmt.Sprintf("/api/recipe/%d", recipe.ID)
	// w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	w.Write(output)
	return
}

// func (uh *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
// 	var u entity.User
// 	err := json.NewDecoder(r.Body).Decode(&u)
// 	// if uh.service.EmailExists(u.Email) {
// 	// 	responses.ERROR(w, http.StatusNotAcceptable, errors.New("Email already occupied"))
// 	// 	json.NewEncoder(w).Encode(err)
// 	// }
// 	pass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		fmt.Println(err)
// 		http.Error(w, errors.New("Password Encryption  failed", http.StatusInternalServerError),

// 		json.NewEncoder(w).Encode(err)
// 	}

// 	u.Password = string(pass)

// 	newUser, errs := uh.service.StoreUser(&u)
// 	if len(errs) > 0 {
// 		responses.ERROR(w, http.StatusInternalServerError, errors.New("User Creation Failed"))
// 		return
// 	}

// 	json.NewEncoder(w).Encode(newUser)
// }

// func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
// 	var u model.User

// 	err := json.NewDecoder(r.Body).Decode(&u)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	user, errs := uh.service.UserByEmail(u.Email)
// 	if len(errs) > 0 {
// 		responses.ERROR(w, http.StatusUnauthorized, errors.New("authentication error"))
// 		return
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, errors.New("authentication error"))
// 		return
// 	}
// 	//claims := rtoken.NewClaims("adonaTesfaye", time.Now().AddDate(0, 1, 1).Unix())
// 	tokenString, err := rtoken.GenerateToken("adonaTesfaye", rtoken.CustomClaims{
// 		User: *user,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().AddDate(0, 1, 1).Unix(),
// 			IssuedAt:  time.Now().Unix(),
// 			NotBefore: time.Now().Unix(),
// 		},
// 	})
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	// responses.JSON(w, http.StatusOK, user)
// 	responses.JSON(w, http.StatusOK, struct {
// 		Token string `json:"token" `
// 	}{
// 		Token: tokenString,
// 	})
// 	log.Println(user.id + " has logged in!")

// }