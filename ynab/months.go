package ynab

type MonthsService service

type MonthSummariesResponse struct {
	Data MonthSummariesWrapper `json:"data"`
}

type MonthSummariesWrapper struct {
	Months []MonthSummary `json:"months"`
}

type MonthDetailResponse struct {
	Data MonthDetailWrapper `json:"data"`
}

type MonthDetailWrapper struct {
	Month MonthDetail `json:"month"`
}

type MonthSummary struct {
	Month string `json:"month"`
	Note *string `json:"note"`
	ToBeBudgeted *int `json:"to_be_budgeted"`
	AgeOfMoney *int `json:"age_of_money"`
}

type MonthDetail struct {
	MonthSummary
	Categories []Category `json:"categories"`
}

/*
https://api.youneedabudget.com/v1#/Months/getBudgetMonths
*/
func (ms *MonthsService) List(budgetId string) ([]MonthSummary, error) {
	var response MonthSummariesResponse
	if err := service(*ms).do("GET", "budgets/" + budgetId + "/months", nil, &response); err != nil {
		return nil, err
	}
	return response.Data.Months, nil
}

/*
https://api.youneedabudget.com/v1#/Months/getBudgetMonth
*/
func (ms *MonthsService) Get(budgetId string, month string) (MonthDetail, error) {
	var response MonthDetailResponse
	if err := service(*ms).do("GET", "budgets/" + budgetId + "/months/" + month, nil, &response); err != nil {
		return MonthDetail{}, err
	}
	return response.Data.Month, nil
}
