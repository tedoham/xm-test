package service

import (
	"github.com/google/uuid"
	"github.com/tedoham/xm-test/errs"
	"github.com/tedoham/xm-test/internal/domain"
	"github.com/tedoham/xm-test/internal/dto"
)

type companyService struct {
	repo domain.CompanyRepository
}

func NewCompanyService(repo domain.CompanyRepository) domain.CompanyService {
	return &companyService{repo}
}

func (c companyService) GetCompany(companyId uuid.UUID) (*domain.Company, *errs.AppError) {
	data, err := c.repo.GetCompany(companyId)
	if err != nil {
		return nil, err
	}

	return data, nil
}
func (c companyService) CreateCompany(companyRequest *dto.CompanyRequestResponse) *errs.AppError {
	if len(companyRequest.Description) > 3000 {
		return errs.NewUnexpectedError("The description length should be under 3000 character")
	} else {
		err := c.repo.CreateCompany(companyRequest)
		if err != nil {
			return err
		}
	}

	return nil
}
func (c companyService) UpdateCompany(companyId uuid.UUID, companyRequest *dto.CompanyRequestResponse) *errs.AppError {
	_, err := c.GetCompany(companyId)

	if err != nil {
		return err
	} else {
		if len(companyRequest.Description) > 3000 {
			return errs.NewUnexpectedError("The description length should be under 3000 character")
		} else {
			err := c.repo.UpdateCompany(companyId, companyRequest)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (c companyService) DeleteCompany(companyId uuid.UUID) *errs.AppError {
	_, err := c.GetCompany(companyId)

	if err != nil {
		return err
	} else {
		err = c.repo.DeleteCompany(companyId)
		if err != nil {
			return err
		}
	}

	return nil
}
