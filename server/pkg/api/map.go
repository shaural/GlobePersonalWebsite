package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shaural/GlobePersonalWebsite/server/pkg/db"
)

func AddMapHandler(router *mux.Router) {
	mapRouter := router.PathPrefix("/api/map").Subrouter()
	mapRouter.HandleFunc("/country", handleGetCountry).Methods("GET")
	mapRouter.HandleFunc("/state", handleGetState).Methods("GET")
}

func handleGetCountry(w http.ResponseWriter, r *http.Request) {
	ldb, err := db.NewDatabase(r.Context())
	if CheckAndWriteError(err, http.StatusInternalServerError, w) {
		return
	}
	defer ldb.Close()
	countries, err := ldb.GetCountries()
	if CheckAndWriteError(err, http.StatusInternalServerError, w) {
		return
	}
	SendJSONResponse(countries, w)
}

func handleGetState(w http.ResponseWriter, r *http.Request) {
	ldb, err := db.NewDatabase(r.Context())
	if CheckAndWriteError(err, http.StatusInternalServerError, w) {
		return
	}
	defer ldb.Close()
	states, err := ldb.GetStates()
	if CheckAndWriteError(err, http.StatusInternalServerError, w) {
		return
	}
	stateMap := make(map[string][]*db.State)
	for _, state := range states {
		stateMap[state.CountryID] = append(stateMap[state.CountryID], state)
	}
	SendJSONResponse(stateMap, w)
}
