package handler

type CustomerHandler interface {
}

type customerHandler struct{}

func NewCustomerHandler() CustomerHandler {
	return &customerHandler{}
}
