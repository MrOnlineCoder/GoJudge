package auth

import (
	"net/http"
	"time"
	"log"
  "errors"
	"encoding/json"
	"crypto/sha512"
  "encoding/hex"

  jwt "github.com/dgrijalva/jwt-go"
  "github.com/gorilla/mux"

  "gojudge/db"
	"gojudge/api/utils"
)

const (
	ACCESS_NONE = -1
	ACCESS_PARTICIPANT = 0
	ACCESS_JURY = 1
	ACCESS_ADMIN = 2
)

type Token struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HashPassword(raw string) string {
	hash := sha512.New()
  hash.Write([]byte(raw))
  
  return hex.EncodeToString(hash.Sum(nil))
}

func GenerateToken(user_id int) (string, error) {
	tk := &Token{UserId: user_id};
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk);
	tokenString, err := token.SignedString([]byte("aabbcc"));
	
	if err != nil {
		return "", errors.New("Token generation failed.");
	} else {
		return tokenString, nil
	}
}

func DecodeToken(tokenIn string) (int, error) {
	tk := &Token{};

	token, err := jwt.ParseWithClaims(tokenIn, tk, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	        return nil, errors.New("Invalid signing method.")
	    }

			return []byte("aabbcc"), nil
	});

	if err != nil {
		return -1, errors.New("Malformed token.")
	}

	if !token.Valid {
    return -1, errors.New("Invalid token.")
	}

	return tk.UserId, nil
}

func ValidateAccess(r *http.Request, lvl int) (db.User, error) {
	cookie, err := r.Cookie("token");

	errorUser := db.User{
		Username: "error",
		Fullname: "error",
		Password: "error",
		Access: ACCESS_NONE,
	};

	if err != nil {
		return errorUser, errors.New("Missing access token.");
	}

	id, err := DecodeToken(cookie.Value);

	if err != nil {
		return errorUser, errors.New("Malformed access token.");
	}

	user, err := db.GetUser(id);

	if err != nil {
		return errorUser, errors.New("User not found.");
	}

	if (user.Access < lvl) {
		return errorUser, errors.New("Access denied.");
	}

	return user, nil
}

func clearTokenCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: "",
		Path: "/",
		Expires: time.Unix(0,0),
		MaxAge: -1,
		HttpOnly: true,
	});
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	loginData := &LoginData{};
	err := json.NewDecoder(r.Body).Decode(loginData);

	if err != nil {
		utils.SendError(w, "Invalid request data.")
		return
	}

	user, err := db.FindAuthUser(loginData.Username, HashPassword(loginData.Password));

	if err != nil {
		utils.SendError(w, "Incorrect username/password.")
		return
	}

	token, err := GenerateToken(user.Id);

	if err != nil {
		utils.SendError(w, "Token generation failed.");
		log.Println("Token generation failed for user", user.Id, ",", err);
		return;
	}

	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: token,
		Path: "/",
		Expires: time.Now().AddDate(0, 0, 3),
		MaxAge: 3 * 24 * 60 * 60,
		HttpOnly: true,
	});

	utils.SendSuccess(w, map[string]interface{} {});
}

func meHandler(w http.ResponseWriter, r *http.Request) {
	user, err := ValidateAccess(r, ACCESS_NONE);

	if err != nil {
		clearTokenCookie(w);
		utils.SendError(w, err.Error());
		return;
	}

	utils.SendSuccess(w, map[string]interface{} {
		"user": map[string]interface{} {
			"fullname": user.Fullname,
			"username": user.Username,
			"access": user.Access,
		},
	});
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("token");

	if err != nil {
		utils.SendError(w, "Missing access token.");
		return;
	}

	clearTokenCookie(w);
	utils.SendSuccess(w, map[string]interface{}{})
}

func InitAuthAPI(router *mux.Router) {
	router.HandleFunc("/login", loginHandler).Methods("POST");
	router.HandleFunc("/logout", logoutHandler).Methods("POST");
	router.HandleFunc("/me", meHandler);
}