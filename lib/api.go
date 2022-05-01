package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type API struct {
	FirmwareStore FirmwareStore
	LogOutput     io.Writer
}

func sendJSON(object interface{}, w http.ResponseWriter) {
	encoded, err := json.Marshal(object)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprint(w, "failed to write the response")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(encoded)
}

func (a *API) getAllFirmware(w http.ResponseWriter, r *http.Request) {
	firmwareList, err := a.FirmwareStore.GetAllFirmware()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprint(w, "failed to request firmware list from the firmware store")
		return
	}

	sendJSON(firmwareList, w)
}

func (a *API) getFirmwareTypes(w http.ResponseWriter, r *http.Request) {
	types, err := a.FirmwareStore.GetAllTypes()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprint(w, "failed to request firmware types from the firmware store")
		return
	}

	sendJSON(types, w)
}

func (a *API) getAllFirmwareByType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	firmwareList, err := a.FirmwareStore.GetAllFirmwareByType(vars["type"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "failed to request firmware for type %s from the firmware store", vars["type"])
		return
	}

	if len(firmwareList) == 0 {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintf(w, "no firmware found for type %s", vars["type"])
		return
	}

	sendJSON(firmwareList, w)
}

func (a *API) getFirmware(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	firmware, err := a.FirmwareStore.GetFirmware(vars["type"], vars["version"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "failed to request firmware %s %s from the firmware store", vars["type"], vars["version"])
		return
	}

	if firmware == nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintf(w, "no firmware %s %s found", vars["type"], vars["version"])
		return
	}

	sendJSON(firmware, w)
}

func (a *API) addFirmware(w http.ResponseWriter, r *http.Request) {
}

func (a *API) deleteFirmware(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := a.FirmwareStore.DeleteFirmware(vars["type"], vars["version"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "failed to delete firmware %s %s from the firmware store", vars["type"], vars["version"])
		return
	}
}

func (a *API) getFirmwareData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data, err := a.FirmwareStore.GetFirmwareData(vars["type"], vars["version"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "failed to request firmware data %s %s from the firmware store", vars["type"], vars["version"])
		return
	}

	if len(data) == 0 {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintf(w, "no firmware data %s %s found", vars["type"], vars["version"])
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s-%s.bin\"", vars["type"], vars["version"]))
	w.Header().Set("Content-Type", "application/octet-stream")
	_, _ = w.Write(data)
}

func (a *API) GetMux() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", a.getAllFirmware).Methods("GET")
	r.HandleFunc("/types", a.getFirmwareTypes).Methods("GET")
	r.HandleFunc("/{type}", a.getAllFirmwareByType).Methods("GET")
	// r.HandleFunc("/{type}/latest", a.getLatestFirmwareForType).Methods("GET")
	r.HandleFunc("/{type}/{version}", a.getFirmware).Methods("GET")
	r.HandleFunc("/{type}/{version}", a.addFirmware).Methods("PUT")
	r.HandleFunc("/{type}/{version}", a.deleteFirmware).Methods("DELETE")
	r.HandleFunc("/{type}/{version}/data", a.getFirmwareData).Methods("GET")
	return handlers.LoggingHandler(a.LogOutput, r)
}
