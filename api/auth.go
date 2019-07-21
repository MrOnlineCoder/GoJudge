package api

import (
	"gojudge/db"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"encoding/json"
	"crypto/sha512"
  "encoding/hex"
  jwt "github.com/dgrijalva/jwt-go"
  "log"
  "errors"
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

func clearTokenCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: "",
		MaxAge: 0,
		HttpOnly: true,
	});
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	loginData := &LoginData{};
	err := json.NewDecoder(r.Body).Decode(loginData);

	if err != nil {
		SendError(w, "Invalid request data.")
		return
	}

	user, err := db.FindAuthUser(loginData.Username, HashPassword(loginData.Password));

	if err != nil {
		SendError(w, "Incorrect username/password.")
		return
	}

	token, err := GenerateToken(user.Id);

	if err != nil {
		SendError(w, "Token generation failed.");
		log.Println("Token generation failed for user", user.Id, ",", err);
		return;
	}

	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: token,
		Expires: time.Now().AddDate(0, 0, 1),
		HttpOnly: true,
	});

	SendSuccess(w, map[string]interface{} {});
}

func meHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token");

	if err != nil {
		SendError(w, "Missing access token.");
		return;
	}

	id, err := DecodeToken(cookie.Value);

	if err != nil {
		clearTokenCookie(w);
		SendError(w, "Invalid token.");
		return;
	}

	user, err := db.GetUser(id);

	if err != nil {
		clearTokenCookie(w);
		SendError(w, "User not found.");
		return;
	}

	SendSuccess(w, map[string]interface{} {
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
		SendError(w, "Missing access token.");
		return;
	}

	clearTokenCookie(w);
	SendSuccess(w, map[string]interface{}{})
}

func InitAuthAPI(router *mux.Router) {
	router.HandleFunc("/login", loginHandler).Methods("POST");
	router.HandleFunc("/logout", logoutHandler).Methods("POST");
	router.HandleFunc("/me", meHandler);
}