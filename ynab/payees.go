package ynab

type PayeesService service

type PayeesResponse struct {
	Data PayeesWrapper `json:"data"`
}

type PayeesWrapper struct {
	Payees []Payee `json:"payees"`
}

type PayeeResponse struct {
	Data PayeeWrapper `json:"data"`
}

type PayeeWrapper struct {
	Payee Payee `json:"payee"`
}

type Payee struct {
	Id string `json:"id"`
	Name string `json:"name"`
	TransferAccountId *string `json:"transfer_account_id"`
}

/*
https://api.youneedabudget.com/v1#/Payees/getPayees
*/
func (ps *PayeesService) List(budgetId string) ([]Payee, error) {
	var response PayeesResponse
	if err := service(*ps).do("GET", "budgets/" + budgetId + "/payees", nil, &response); err != nil {
		return nil, err
	}
	return response.Data.Payees, nil
}

/*
https://api.youneedabudget.com/v1#/Payees/getPayeeById
*/
func (ps *PayeesService) Get(budgetId string, payeeId string) (Payee, error) {
	var response PayeeResponse
	if err := service(*ps).do("GET", "budgets/" + budgetId + "/payees/" + payeeId, nil, &response); err != nil {
		return Payee{}, err
	}
	return response.Data.Payee, nil
}
