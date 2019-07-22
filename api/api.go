package api

import (
	"github.com/gorilla/mux"
	
	"gojudge/auth"
	"gojudge/admin"
	"gojudge/contest"
)

var router *mux.Router;

func Create() *mux.Router {
	router = mux.NewRouter();

	// /api
	apiRouter := router.PathPrefix("/api").Subrouter();

	// /api/auth
	authRouter := apiRouter.PathPrefix("/auth").Subrouter();
	auth.InitAuthAPI(authRouter);

	// /api/admin
	adminRouter := apiRouter.PathPrefix("/admin").Subrouter();
	admin.InitAdminAPI(adminRouter);

	// /api/contest
	contestRouter := apiRouter.PathPrefix("/contest").Subrouter();
	contest.InitContestAPI(contestRouter);

	return router;
}