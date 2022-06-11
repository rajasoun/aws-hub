package handlers

import (
	"github.com/gorilla/mux"
)

const HealthEndPoint = "/health"
const LocalProfilesEndPoint = "/aws/profiles"
const UsersCountEndPoint = "/aws/iam/users"
const UserIdentityEndPoint = "/aws/iam/account"
const AccountAliasEndPoint = "/aws/iam/alias"

func (handler *AWSHandler) SetUpRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(HealthEndPoint, handler.HealthCheckHandler)
	router.HandleFunc(LocalProfilesEndPoint, handler.ListProfilesHandler)
	router.HandleFunc(UsersCountEndPoint, handler.IAMGetUserCountHandler)
	router.HandleFunc(UserIdentityEndPoint, handler.IAMGetUserIdentityHandler)
	router.HandleFunc(AccountAliasEndPoint, handler.IAMGetAliasesHandler)
	return router
}
