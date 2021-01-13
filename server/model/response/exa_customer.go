package response

import "nginx-web/model"

type ExaCustomerResponse struct {
	Customer model.ExaCustomer `json:"customer"`
}
