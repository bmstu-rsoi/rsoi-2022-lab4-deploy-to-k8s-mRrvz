package handlers

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/bonus", CreatePrivilegeHistoryHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/bonus/privilege", CreatePrivilegeHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/bonus/{username}", GetPrivilegeByUsernameHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/bonus/history/{privilegeId}", GetHistoryByIdHandler).Methods("GET", "OPTIONS")

	return router
}
