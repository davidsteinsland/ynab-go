package ynab

type BudgetsService service

type BudgetSummaryResponse struct {
	Data BudgetSummaryWrapper `json:"data"`
}

type BudgetSummaryWrapper struct {
	Budgets []BudgetSummary `json:"budgets"`
}

type BudgetDetailResponse struct {
	Data BudgetDetailWrapper `json:"data"`
}

type BudgetDetailWrapper struct {
	Budget          BudgetDetail `json:"budget"`
	ServerKnowledge int          `json:"server_knowledge"`
}

type BudgetSummary struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	LastModifiedOn *string        `json:"last_modified_on"`
	DateFormat     DateFormat     `json:"date_format"`
	CurrencyFormat CurrencyFormat `json:"currency_format"`
}

type BudgetDetail struct {
	BudgetSummary
	Accounts                 []Account                     `json:"accounts"`
	Payees                   []Payee                       `json:"payees"`
	PayeeLocations           []PayeeLocation               `json:"payee_locations"`
	CategoryGroups           []CategoryGroup               `json:"category_groups"`
	Categories               []Category                    `json:"categories"`
	Months                   []MonthDetail                 `json:"months"`
	Transactions             []TransactionSummary          `json:"transactions"`
	Subtransactions          []SubTransaction              `json:"subtransactions"`
	ScheduledTransactions    []ScheduledTransactionSummary `json:"scheduled_transactions"`
	ScheduledSubtransactions []ScheduledSubTransaction     `json:"scheduled_subtransactions"`
}

type DateFormat struct {
	Format string `json:"format"`
}

type CurrencyFormat struct {
	IsoCode          string `json:"iso_code"`
	ExampleFormat    string `json:"example_format"`
	DecimalDigits    int    `json:"decimal_digits"`
	DecimalSeparator string `json:"decimal_separator"`
	SymbolFirst      bool   `json:"symbol_first"`
	GroupSeparator   string `json:"group_separator"`
	CurrencySymbol   string `json:"currency_symbol"`
	DisplaySymbol    bool   `json:"display_symbol"`
}

/*
https://api.youneedabudget.com/v1#/Budgets/getBudgets
*/
func (bs *BudgetsService) List() ([]BudgetSummary, error) {
	var response BudgetSummaryResponse
	if err := service(*bs).do("GET", "budgets", nil, &response); err != nil {
		return nil, err
	}
	return response.Data.Budgets, nil
}

/*
https://api.youneedabudget.com/v1#/Budgets/getBudgetById
*/

func (bs *BudgetsService) Get(id string) (BudgetDetail, error) {
	var response BudgetDetailResponse
	if err := service(*bs).do("GET", "budgets/"+id, nil, &response); err != nil {
		return BudgetDetail{}, err
	}
	return response.Data.Budget, nil
}
