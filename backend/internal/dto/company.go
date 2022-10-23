package dto

import (
	"encoding/json"
	"io"
)

// CompanyRequestResponse is an representation request body to create a new Company
type CompanyRequestResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	Registered  bool   `json:"registered"`
	Type        int    `json:"type"`
}

// FromJSONCompanyRequestResponse converts json body request to a CompanyRequestResponse struct
func FromJSONCompanyRequestResponse(body io.Reader) (*CompanyRequestResponse, error) {
	companyRequestResponse := CompanyRequestResponse{}
	if err := json.NewDecoder(body).Decode(&companyRequestResponse); err != nil {
		return nil, err
	}

	return &companyRequestResponse, nil
}
