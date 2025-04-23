package service

import (
	"payso-internal-api/handler"
	"payso-internal-api/model"
	customerRepository "payso-internal-api/repository"

	log "github.com/sirupsen/logrus"
)

type CustomerService interface {
	GetcustomerService(mid string, page int, row int) (model.CustomerPagination, error)
	// CreatecustomerService(body model.CreateMerchantPayload, ipAddress string) (model.UpdateResponse, error)
	// DeletecustomerService(ReqMasterMerchantID string, ReqMerchantID string) (model.UpdateResponse, error)
}

type customerService struct {
	customerHandler handler.CustomerHandler
}

func NewcustomerService(customerHandler handler.CustomerHandler) CustomerService {
	return &customerService{customerHandler}
}

func (s *customerService) GetcustomerService(mid string, page int, row int) (model.CustomerPagination, error) {
	log.Infof("==-- GetcustomerService --==")

	var err error
	var CustomerList []model.CustomerPayload

	CustomerList, err = customerRepository.GetCustomerRepository(mid, page, row)
	if err != nil {
		log.Errorf("Error from GetCustomerRepository: %v", err)
		return model.CustomerPagination{}, err
	}

	// TotalPages, err := customerRepository.GetTotalCustomerRepository(mid)
	// if err != nil {
	// 	log.Errorf("Error from GetCustomerRepository: %v", err)
	// 	return model.CustomerPagination{}, err
	// }

	CustomerPagination := model.CustomerPagination{
		// TotalPages:   TotalPages,
		CustomerList: CustomerList}

	return CustomerPagination, err
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
