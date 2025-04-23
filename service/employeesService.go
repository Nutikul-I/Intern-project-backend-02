package service

import (
	"payso-internal-api/handler"
	"payso-internal-api/model"
	EmployeesRepository "payso-internal-api/repository"

	log "github.com/sirupsen/logrus"
)

type EmployeesService interface {
	GetEmployeesService(mid string, page int, row int) (model.EmployeesPagination, error)
	// CreateEmployeesService(body model.CreateMerchantPayload, ipAddress string) (model.UpdateResponse, error)
	// DeleteEmployeesService(ReqMasterMerchantID string, ReqMerchantID string) (model.UpdateResponse, error)
}
type employeesService struct {
	employeesHandler handler.EmployeesHandler
}

func NewEmployeesService(employeesHandler handler.EmployeesHandler) EmployeesService {
	return &employeesService{employeesHandler}
}

func (s *employeesService) GetEmployeesService(mid string, page int, row int) (model.EmployeesPagination, error) {
	log.Infof("==-- GetemployeesService --==")

	var err error
	var EmployeesList []model.EmployeesPayload

	EmployeesList, err = EmployeesRepository.GetEmployeesRepository(mid, page, row)
	if err != nil {
		log.Errorf("Error from GetEmployeesRepository: %v", err)
		return model.EmployeesPagination{}, err
	}

	// TotalPages, err := employeesRepository.GetTotalEmployeesRepository(mid)
	// if err != nil {
	// 	log.Errorf("Error from GetEmployeesRepository: %v", err)
	// 	return model.EmployeesPagination{}, err
	// }

	EmployeesPagination := model.EmployeesPagination{
		// TotalPages:   TotalPages,
		EmployeesList: EmployeesList}

	return EmployeesPagination, err
}

// func (s *customerService) CreatecustomerService(body model.CreateMerchantPayload, ipAddress string) (model.UpdateResponse, error) {
// 	log.Infof("==-- CreatecustomerService --==")

// 	var err error
// 	var Result model.UpdateResponse

// 	Result, err = merchantRepository.CreateMerchantRepository(body)
// 	if err != nil {
// 		log.Errorf("Error from CreateMerchantRepository: %v", err)
// 		return model.UpdateResponse{}, err
// 	}
// 	return Result, err
// }

// func (s *customerService) DeletecustomerService(ReqMasterMerchantID string, ReqMerchantID string) (model.UpdateResponse, error) {
// 	log.Infof("==-- DeletecustomerService --==")

// 	var err error
// 	var Result model.UpdateResponse

// 	Result, err = merchantRepository.DeleteMerchantRepository(ReqMasterMerchantID, ReqMerchantID)
// 	if err != nil {
// 		log.Errorf("Error from DeleteMerchantRepository: %v", err)
// 		return model.UpdateResponse{}, err
// 	}

// 	return Result, err
// }
