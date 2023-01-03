package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type API struct {
	FirmwareStore FirmwareStore
	LogOutput     io.Writer
}

func (a *API) logAndRespond(w http.ResponseWriter, message string, err error) {
	_, _ = fmt.Fprint(w, message)
	if err != nil {
		_, _ = fmt.Fprintf(a.LogOutput, "%s: %s", message, err.Error())
	} else {
		_, _ = fmt.Fprint(a.LogOutput, message)
	}
}

func (a *API) sendJSON(object interface{}, w http.ResponseWriter) {
	encoded, err := json.Marshal(object)
	if err != nil {
		a.logAndRespond(w, "failed to write the response", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(encoded)
}

func (a *API) getAllFirmware(w http.ResponseWriter, r *http.Request) {
	firmwareList, err := a.FirmwareStore.GetAllFirmware()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.logAndRespond(w, "failed to request firmware list from the firmware store", err)
		return
	}

	a.sendJSON(firmwareList, w)
}

func (a *API) getFirmwareTypes(w http.ResponseWriter, r *http.Request) {
	types, err := a.FirmwareStore.GetAllTypes()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.logAndRespond(w, "failed to request firmware types from the firmware store", err)
		return
	}

	a.sendJSON(types, w)
}

func (a *API) getAllFirmwareByType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	firmwareList, err := a.FirmwareStore.GetAllFirmwareByType(vars["type"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.logAndRespond(w, fmt.Sprintf("failed to request firmware for type %s from the firmware store", vars["type"]), err)
		return
	}

	if len(firmwareList) == 0 {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintf(w, "no firmware found for type %s", vars["type"])
		return
	}

	a.sendJSON(firmwareList, w)
}

func (a *API) getFirmware(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	firmware, err := a.FirmwareStore.GetFirmware(vars["type"], vars["version"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.logAndRespond(w, fmt.Sprintf("failed to request firmware %s %s from the firmware store", vars["type"], vars["version"]), err)
		return
	}

	if firmware == nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintf(w, "no firmware %s %s found", vars["type"], vars["version"])
		return
	}

	a.sendJSON(firmware, w)
}

func (a *API) addFirmware(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if IsInvalidType(vars["type"]) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "\"%s\" is not a valid firmware type", vars["type"])
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		a.logAndRespond(w, fmt.Sprintf("failed to read body of new firmware %s %s", vars["type"], vars["version"]), err)
		return
	}

	err = a.FirmwareStore.AddFirmware(vars["type"], vars["version"], body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.logAndRespond(w, fmt.Sprintf("failed to add new firmware %s %s to the firmware store", vars["type"], vars["version"]), err)
		return
	}
}

func (a *API) deleteFirmware(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := a.FirmwareStore.DeleteFirmware(vars["type"], vars["version"])
	if err != nil {
		if strings.Contains(err.Error(), "does not exist") {
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprintf(a.LogOutput, "attempt to delete missing firmware %s %s: %s", vars["type"], vars["version"], err.Error())
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			a.logAndRespond(w, fmt.Sprintf("failed to delete firmware %s %s from the firmware store", vars["type"], vars["version"]), err)
		}
		return
	}
}

func (a *API) getFirmwareData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data, err := a.FirmwareStore.GetFirmwareData(vars["type"], vars["version"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.logAndRespond(w, fmt.Sprintf("failed to request firmware data %s %s from the firmware store", vars["type"], vars["version"]), err)
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
	r.HandleFunc("/{type}/{version}", a.getFirmware).Methods("GET")
	r.HandleFunc("/{type}/{version}", a.addFirmware).Methods("PUT")
	r.HandleFunc("/{type}/{version}", a.deleteFirmware).Methods("DELETE")
	r.HandleFunc("/{type}/{version}/data", a.getFirmwareData).Methods("GET")
	return handlers.LoggingHandler(a.LogOutput, r)
}
