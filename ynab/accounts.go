package ynab

type AccountsService service

type AccountsResponse struct {
	Data AccountsWrapper `json:"data"`
}

type AccountsWrapper struct {
	Accounts []Account `json:"accounts"`
}

type AccountResponse struct {
	Data AccountWrapper `json:"data"`
}

type AccountWrapper struct {
	Account Account `json:"account"`
}

type Account struct {
	Id               string  `json:"id"`
	Name             string  `json:"name"`
	Type             string  `json:"type"`
	OnBudget         bool    `json:"on_budget"`
	Closed           bool    `json:"closed"`
	Note             *string `json:"note"`
	Balance          int     `json:"balance"`
	ClearedBalance   int     `json:"cleared_balance"`
	UnclearedBalance int     `json:"uncleared_balance"`
}

/*
https://api.youneedabudget.com/v1#/Accounts/getAccounts
*/
func (as *AccountsService) List(budgetId string) ([]Account, error) {
	var response AccountsResponse
	if err := service(*as).do("GET", "budgets/"+budgetId+"/accounts", nil, &response); err != nil {
		return nil, err
	}
	return response.Data.Accounts, nil
}

/*
https://api.youneedabudget.com/v1#/Accounts/getAccountById
*/
func (as *AccountsService) Get(budgetId string, accountId string) (Account, error) {
	var response AccountResponse
	if err := service(*as).do("GET", "budgets/"+budgetId+"/accounts/"+accountId, nil, &response); err != nil {
		return Account{}, err
	}
	return response.Data.Account, nil
}
