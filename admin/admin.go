package admin

import (
	"net/http"
 	"encoding/json"
	"gojudge/auth"
	"gojudge/api/utils"
	"github.com/gorilla/mux"
	"gojudge/db"
)

type CreateUserBodyContents struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Password string `json:"password"`
	Access int `json:"access"`
}

type CreateUserBody struct {
	User CreateUserBodyContents `json:"user"`
}

type DeleteUserBody struct {
	UserID int `json:"user_id"`
}

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	_, err := auth.ValidateAccess(r, auth.ACCESS_ADMIN);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	users, err := db.GetAllUsers();

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	utils.SendSuccess(w, map[string]interface{} {
		"users": users,
	});
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	_, err := auth.ValidateAccess(r, auth.ACCESS_ADMIN);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	parsedBody := &CreateUserBody{};

	err = json.NewDecoder(r.Body).Decode(parsedBody);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	parsedBody.User.Password = auth.HashPassword(parsedBody.User.Password);

	target := db.User{
		Username: parsedBody.User.Username,
		Password: parsedBody.User.Password,
		Fullname: parsedBody.User.Fullname,
		Access: parsedBody.User.Access,
	};

	ok := db.CreateUser(target);

	if !ok {
		utils.SendError(w, "Database write error.");
		return;
	}

	utils.SendSuccess(w, map[string]interface{}{});
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	_, err := auth.ValidateAccess(r, auth.ACCESS_ADMIN);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	parsedBody := &DeleteUserBody{};

	err = json.NewDecoder(r.Body).Decode(parsedBody);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	ok := db.DeleteUser(parsedBody.UserID);

	if !ok {
		utils.SendError(w, "Database write error.");
		return;
	}

	utils.SendSuccess(w, map[string]interface{}{});
}


func InitAdminAPI(router *mux.Router) {
	router.HandleFunc("/users", listUsersHandler).Methods("GET");
	router.HandleFunc("/createUser", createUserHandler).Methods("POST");
	router.HandleFunc("/deleteUser", deleteUserHandler).Methods("POST");
}