package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"

	"github.com/tedoham/xm-test/errs"
	"github.com/tedoham/xm-test/internal/domain"
	"github.com/tedoham/xm-test/internal/dto"
	"github.com/tedoham/xm-test/logger"
)

type companyRepository struct {
	client *sqlx.DB
}

func NewCompanyRepository(client *sqlx.DB) domain.CompanyRepository {
	return &companyRepository{
		client: client,
	}
}

func (c companyRepository) GetCompany(companyId uuid.UUID) (*domain.Company, *errs.AppError) {
	sqlGetCompany := `SELECT * FROM companies WHERE id = $1`

	var company domain.Company
	err := c.client.Get(&company, sqlGetCompany, companyId)

	if err != nil {
		logger.Error("Error while fetching company information: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &company, nil
}

func (c companyRepository) CreateCompany(companyRequest *dto.CompanyRequestResponse) *errs.AppError {
	sqlCreateCompany := `INSERT INTO companies (name, description, amount, registered, company_type_id) 
		VALUES ($1, $2, $3, $4, $5) RETURNING *`

	err := c.client.Get(&companyRequest, sqlCreateCompany)

	if err != nil {
		logger.Error("Error while creating company information: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	return nil
}

func (c companyRepository) UpdateCompany(companyId uuid.UUID, companyRequest *dto.CompanyRequestResponse) *errs.AppError {
	sqlUpdateCompany := `UPDATE companies SET name = $1, description = $2, amount = $3, registered = $4, company_type_id = $5 
		WHERE id = $1 RETURNING *`

	err := c.client.Get(&companyRequest, sqlUpdateCompany, companyId)

	if err != nil {
		logger.Error("Error while updating company information: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	return nil
}
func (c companyRepository) DeleteCompany(companyId uuid.UUID) *errs.AppError {
	sqlDeleteCompany := `DELETE FROM companies WHERE id = $1`

	_, err := c.client.Exec(sqlDeleteCompany, companyId)

	if err != nil {
		logger.Error("Error while deleting company information: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	return nil
}
