package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bmstu-rsoi/rsoi-2022-lab2-microservices-mRrvz/src/bonus-service/internal/models"
	"github.com/bmstu-rsoi/rsoi-2022-lab2-microservices-mRrvz/src/bonus-service/internal/repository"
	"github.com/gorilla/mux"
)

func CreatePrivilegeHistoryHandler(w http.ResponseWriter, r *http.Request) {
	var record models.PrivilegeHistory

	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bonusRepo := repository.BonusRepository{}
	if err := bonusRepo.CreateHistoryRecord(&record); err != nil {
		log.Printf("Failed to create ticket: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func CreatePrivilegeHandler(w http.ResponseWriter, r *http.Request) {
	var record models.Privilege

	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bonusRepo := repository.BonusRepository{}
	if err := bonusRepo.CreatePrivilege(&record); err != nil {
		log.Printf("Failed to create ticket: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetPrivilegeByUsernameHandler(w http.ResponseWriter, r *http.Request) {
	bonusRepo := repository.BonusRepository{}
	params := mux.Vars(r)
	privilege, err := bonusRepo.GetPrvilegeByUsername(params["username"])
	if err != nil {
		log.Printf("failed to get flghts: %s", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(privilege); err != nil {
		log.Printf("Failed to encode flight: %s", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetHistoryByIdHandler(w http.ResponseWriter, r *http.Request) {
	bonusRepo := repository.BonusRepository{}
	params := mux.Vars(r)
	history, err := bonusRepo.GetHistoryById(params["privilegeId"])
	if err != nil {
		log.Printf("failed to get flghts: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(history); err != nil {
		log.Printf("Failed to encode flight: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
