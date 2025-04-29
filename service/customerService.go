package service

import (
	"payso-internal-api/handler"
	"payso-internal-api/model"
	customerRepository "payso-internal-api/repository"

	log "github.com/sirupsen/logrus"
)

type CustomerService interface {
	GetCustomerService(mid string, page int, row int) (model.CustomerPagination, error)
	CreateCustomerService(body model.CreateCustomerPayload, ipAddress string) (model.UpdateResponse, error)

	// DeletecustomerService(ReqCustomerID string) (model.UpdateResponse, error)
}

type customerService struct {
	customerHandler handler.CustomerHandler
}

func NewCustomerService(customerHandler handler.CustomerHandler) CustomerService {
	return &customerService{customerHandler}
}

func (s *customerService) GetCustomerService(mid string, page int, row int) (model.CustomerPagination, error) {
	log.Infof("==-- GetcustomerService --==")

	var err error
	var CustomerList []model.CustomerPayload

	CustomerList, err = customerRepository.GetCustomerRepository(mid, page, row)
	if err != nil {
		log.Errorf("Error from GetCustomerRepository: %v", err)
		return model.CustomerPagination{}, err
	}

	TotalPages, err := customerRepository.GetTotalCustomerRepository(row)
	if err != nil {
		log.Errorf("Error from GetCustomerRepository: %v", err)
		return model.CustomerPagination{}, err
	}

	CustomerPagination := model.CustomerPagination{
		TotalPages:   TotalPages,
		CustomerList: CustomerList}

	return CustomerPagination, err
}

func (s *customerService) CreateCustomerService(body model.CreateCustomerPayload, ipAddress string) (model.UpdateResponse, error) {
	log.Infof("==-- CreatecustomerService --==")

	var err error
	var Result model.UpdateResponse

	Result, err = customerRepository.CreateCustomerRepository(body)
	if err != nil {
		log.Errorf("Error from CreateCustomerRepository: %v", err)
		return model.UpdateResponse{}, err
	}
	return Result, err
}

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
