package ynab

type ScheduledTransactionsService service

type ScheduledTransactionsResponse struct {
	Data ScheduledTransactionsWrapper `json:"data"`
}

type ScheduledTransactionsWrapper struct {
	ScheduledTransactions []ScheduledTransactionDetail `json:"scheduled_transactions"`
}

type ScheduledTransactionResponse struct {
	Data ScheduledTransactionWrapper `json:"data"`
}

type ScheduledTransactionWrapper struct {
	ScheduledTransaction ScheduledTransactionDetail `json:"scheduled_transaction"`
}

type ScheduledTransactionSummary struct {
	Id string `json:"id"`
	DateFirst string `json:"date_first"`
	DateNext string `json:"date_next"`
	Frequency string `json:"frequency"`
	Amount int `json:"amount"`
	Memo *string `json:"memo"`
	FlagColor *string `json:"flag_color"`
	AccountId string `json:"account_id"`
	PayeeId *string `json:"payee_id"`
	CategoryId *string `json:"category_id"`
	TransferAccountId *string `json:"transfer_account_id"`
}

type ScheduledTransactionDetail struct {
	ScheduledTransactionSummary

	AccountName string `json:"account_name"`
	PayeeName string `json:"payee_name"`
	CategoryName string `json:"category_name"`

	SubTransactions []ScheduledSubTransaction `json:"subtransactions"`
}

type ScheduledSubTransaction struct {
	Id string `json:"id"`
	ScheduledTransactionId string `json:"scheduled_transaction_id"`
	Amount int `json:"amount"`
	Memo *string `json:"memo"`
	PayeeId *string `json:"payee_id"`
	CategoryId *string `json:"category_id"`
	TransferAccountId *string `json:"transfer_account_id"`
}

/*
https://api.youneedabudget.com/v1#/Scheduled_Transactions/getScheduledTransactions
*/
func (sts *ScheduledTransactionsService) List(budgetId string) ([]ScheduledTransactionDetail, error) {
	var response ScheduledTransactionsResponse
	if err := service(*sts).do("GET", "budgets/" + budgetId + "/scheduled_transactions", nil, &response); err != nil {
		return nil, err
	}
	return response.Data.ScheduledTransactions, nil
}

/*
https://api.youneedabudget.com/v1#/Scheduled_Transactions/getScheduledTransactionById
*/
func (sts *ScheduledTransactionsService) Get(budgetId string, scheduledTransactionId string) (ScheduledTransactionDetail, error) {
	var response ScheduledTransactionResponse
	if err := service(*sts).do("GET", "budgets/" + budgetId + "/scheduled_transactions/" + scheduledTransactionId, nil, &response); err != nil {
		return ScheduledTransactionDetail{}, err
	}
	return response.Data.ScheduledTransaction, nil
}
