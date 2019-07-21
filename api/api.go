package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

var router *mux.Router;

func SendError(w http.ResponseWriter, msg string) {
	w.Header().Add("Content-Type", "application/json")

	respMap := map[string] interface{} {
		"success": false, 
		"message": msg,
	}

	json.NewEncoder(w).Encode(respMap);
}

func SendSuccess(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")

	respMap := map[string] interface{} {
		"success": true,
	}

	for k, v := range data {
    respMap[k] = v
  }

	json.NewEncoder(w).Encode(respMap);
}

func Create() *mux.Router {
	router = mux.NewRouter();

	// /api
	apiRouter := router.PathPrefix("/api").Subrouter();

	// /api/auth
	authRouter := apiRouter.PathPrefix("/auth").Subrouter();
	InitAuthAPI(authRouter);

	return router;
}