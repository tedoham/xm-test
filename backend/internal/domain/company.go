package domain

import (
	"github.com/google/uuid"
	"github.com/tedoham/xm-test/errs"
	"github.com/tedoham/xm-test/internal/dto"
)

type Company struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Amount      int       `db:"amount"`
	Registered  bool      `db:"registered"`
	Type        int       `db:"company_type_id"`
}

type CompanyService interface {
	GetCompany(companyId uuid.UUID) (*Company, *errs.AppError)
	CreateCompany(companyRequest *dto.CompanyRequestResponse) *errs.AppError
	UpdateCompany(companyId uuid.UUID, companyRequest *dto.CompanyRequestResponse) *errs.AppError
	DeleteCompany(companyId uuid.UUID) *errs.AppError
}

type CompanyRepository interface {
	GetCompany(companyId uuid.UUID) (*Company, *errs.AppError)
	CreateCompany(companyRequest *dto.CompanyRequestResponse) *errs.AppError
	UpdateCompany(companyId uuid.UUID, companyRequest *dto.CompanyRequestResponse) *errs.AppError
	DeleteCompany(companyId uuid.UUID) *errs.AppError
}
