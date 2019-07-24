package config

import (
	"net/http"
 	"encoding/json"

	"gojudge/auth"
	"gojudge/api/utils"
	"gojudge/config"
)

type ConfigBody struct {
	Config config.Config `json:"config"`
}

func SaveConfigHandler(w http.ResponseWriter, r *http.Request) {
	_, err := auth.ValidateAccess(r, auth.ACCESS_ADMIN);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}
	parsedBody := &ConfigBody{};

	err = json.NewDecoder(r.Body).Decode(parsedBody);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	config.Set(parsedBody.Config);
	config.Save();

	utils.SendSuccess(w, map[string]interface{}{});
}

func GetConfigHandler(w http.ResponseWriter, r *http.Request) {
	_, err := auth.ValidateAccess(r, auth.ACCESS_ADMIN);

	if err != nil {
		utils.SendError(w, err.Error());
		return;
	}

	utils.SendSuccess(w, map[string]interface{} {
		"config": config.Get(),
	});
}
