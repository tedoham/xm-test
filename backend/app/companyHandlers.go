package app

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/tedoham/xm-test/internal/domain"
	"github.com/tedoham/xm-test/internal/dto"
)

type CompanyHandlers struct {
	service domain.CompanyService
}

func (ch CompanyHandlers) GetCompany(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["company_id"]
	companyId, errId := uuid.Parse(id)

	if errId != nil {
		writeResponse(w, http.StatusBadRequest, errId)
	}

	company, err := ch.service.GetCompany(companyId)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, company)
	}
}
func (ch CompanyHandlers) CreateCompany(w http.ResponseWriter, r *http.Request) {

	var company *dto.CompanyRequestResponse
	errConvert := json.NewDecoder(r.Body).Decode(&company)

	if errConvert != nil {
		writeResponse(w, http.StatusBadRequest, errConvert.Error())
	} else {
		err := ch.service.CreateCompany(company)

		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, company)
		}
	}
}

func (ch CompanyHandlers) UpdateCompany(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["company_id"]
	companyId, errId := uuid.Parse(id)

	if errId != nil {
		writeResponse(w, http.StatusBadRequest, errId)
	}

	var company *dto.CompanyRequestResponse
	errConvert := json.NewDecoder(r.Body).Decode(&company)

	if errConvert != nil {
		writeResponse(w, http.StatusBadRequest, errConvert.Error())
	} else {
		err := ch.service.UpdateCompany(companyId, company)

		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, company)
		}
	}
}
func (ch CompanyHandlers) DeleteCompany(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["company_id"]
	companyId, errId := uuid.Parse(id)

	if errId != nil {
		writeResponse(w, http.StatusBadRequest, errId)
	}

	err := ch.service.DeleteCompany(companyId)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, companyId)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
