package response

import "github.com/defeng-hub/ByOfficeAutomatic/server/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
