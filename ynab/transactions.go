package ynab

type TransactionsService service

type TransactionsResponse struct {
	Data TransactionsWrapper `json:"data"`
}

type TransactionsWrapper struct {
	Transactions []TransactionDetail `json:"transactions"`
}

type HybridTransactionsResponse struct {
	Data HybridTransactionsWrapper `json:"data"`
}

type HybridTransactionsWrapper struct {
	Transactions []HybridTransaction `json:"transactions"`
}

type TransactionResponse struct {
	Data TransactionWrapper `json:"data"`
}

type TransactionWrapper struct {
	Transaction TransactionDetail `json:"transaction"`
}

type SaveTransactionWrapper struct {
	Transaction *SaveTransaction `json:"transaction"`
}

type BulkResponse struct {
	Data BulkIdWrapper `json:"data"`
}

type BulkIdWrapper struct {
	Bulk BulkIds `json:"bulk"`
}

type BulkIds struct {
	TransactionIds []string `json:"transaction_ids"`
	DuplicateImportIds []string `json:"duplicate_import_ids"`
}

type BulkTransactions struct {
	Transactions []SaveTransaction `json:"transactions"`
}

type TransactionDetail struct {
	TransactionSummary
	AccountName string `json:"account_name"`
	PayeeName string `json:"payee_name"`
	CategoryName string `json:"category_name"`
	SubTransactions []SubTransaction `json:"subtransactions"`
}

type HybridTransaction struct {
	TransactionSummary
	Type string `json:"type"`
	ParentTransactionId string `json:"parent_transaction_id"`
	AccountName string `json:"account_name"`
	PayeeName string `json:"payee_name"`
	CategoryName string `json:"category_name"`
}

type TransactionSummary struct {
	Id string `json:"id"`
	Date string `json:"date"`
	Amount int `json:"amount"`
	Memo *string `json:"memo"`
	Cleared string `json:"cleared"`
	Approved bool `json:"approved"`
	FlagColor *string `json:"flag_color"`
	AccountId string `json:"account_id"`
	PayeeId *string `json:"payee_id"`
	CategoryId *string `json:"category_id"`
	TransferAccountId *string `json:"transfer_account_id"`
	ImportId *string `json:"import_id"`
}

type SaveTransaction struct {
	AccountId string `json:"account_id"`
	Date string `json:"date"`
	Amount int `json:"amount"`
	PayeeId string `json:"payee_id,omitempty"`
	PayeeName string `json:"payee_name,omitempty"`
	CategoryId string `json:"category_id,omitempty"`
	Memo string `json:"memo,omitempty"`
	Cleared string `json:"cleared,omitempty"`
	Approved bool `json:"approved,omitempty"`
	FlagColor string `json:"flag_color,omitempty"`
	ImportId string `json:"import_id,omitempty"`
}

type SubTransaction struct {
	Id string `json:"id"`
	TransactionId string `json:"transaction_id"`
	Amount int `json:"amount"`
	Memo *string `json:"memo"`
	PayeeId *string `json:"payee_id"`
	CategoryId *string `json:"category_id"`
	TransferAccountId *string `json:"transfer_account_id"`
}

func NewSaveTransaction(accountId string, date string, amount int) *SaveTransaction {
	return &SaveTransaction{
		AccountId: accountId,
		Date: date,
		Amount: amount,
	}
}

/*
https://api.youneedabudget.com/v1#/Transactions/getTransactions
*/
func (ts *TransactionsService) List(budgetId string) ([]TransactionDetail, error) {
	var response TransactionsResponse
	if err := service(*ts).do("GET", "budgets/" + budgetId + "/transactions", nil, &response); err != nil {
		return nil, err
	}
	return response.Data.Transactions, nil
}

/*
https://api.youneedabudget.com/v1#/Transactions/getTransactionsById
*/
func (ts *TransactionsService) Get(budgetId string, transactionId string) (TransactionDetail, error) {
	var response TransactionResponse
	if err := service(*ts).do("GET", "budgets/" + budgetId + "/transactions/" + transactionId, nil, &response); err != nil {
		return TransactionDetail{}, err
	}
	return response.Data.Transaction, nil
}

/*
https://api.youneedabudget.com/v1#/Transactions/createTransaction
*/
func (ts *TransactionsService) Create(budgetId string, transaction *SaveTransaction) (TransactionDetail, error) {
	var response TransactionResponse
	if err := service(*ts).do("POST", "budgets/" + budgetId + "/transactions", SaveTransactionWrapper{transaction}, &response); err != nil {
		return TransactionDetail{}, err
	}
	return response.Data.Transaction, nil
}

/*
https://api.youneedabudget.com/v1#/Transactions/bulkCreateTransactions
*/
func (ts *TransactionsService) CreateBulk(budgetId string, transactions []SaveTransaction) ([]TransactionDetail, error) {
	var response TransactionsResponse
	if err := service(*ts).do("POST", "budgets/" + budgetId + "/transactions/bulk", BulkTransactions{transactions}, &response); err != nil {
		return nil, err
	}
	return response.Data.Transactions, nil
}

/*
https://api.youneedabudget.com/v1#/Transactions/updateTransaction
*/
func (ts *TransactionsService) Edit(budgetId string, transactionId string, transaction *SaveTransaction) (TransactionDetail, error) {
	var response TransactionResponse
	if err := service(*ts).do("PUT", "budgets/" + budgetId + "/transactions/" + transactionId, SaveTransactionWrapper{transaction}, &response); err != nil {
		return TransactionDetail{}, err
	}
	return response.Data.Transaction, nil
}

/*
https://api.youneedabudget.com/v1#/Transactions/getTransactionsByAccount
*/
func (ts *TransactionsService) GetByAccount(budgetId string, accountId string) ([]TransactionDetail, error) {
	var response TransactionsResponse
	if err := service(*ts).do("GET", "budgets/" + budgetId + "/accounts/" + accountId + "/transactions", nil, &response); err != nil {
		return nil, err
	}
	return response.Data.Transactions, nil
}

/*
https://api.youneedabudget.com/v1#/Transactions/getTransactionsByCategory
*/
func (ts *TransactionsService) GetByCategory(budgetId string, categoryId string) ([]HybridTransaction, error) {
	var response HybridTransactionsResponse
	if err := service(*ts).do("GET", "budgets/" + budgetId + "/categories/" + categoryId + "/transactions", nil, &response); err != nil {
		return nil, err
	}
	return response.Data.Transactions, nil
}

/*
https://api.youneedabudget.com/v1#/Transactions/getTransactionsByPayee
*/
func (ts *TransactionsService) GetByPayee(budgetId string, payeeId string) ([]HybridTransaction, error) {
	var response HybridTransactionsResponse
	if err := service(*ts).do("GET", "budgets/" + budgetId + "/payees/" + payeeId + "/transactions", nil, &response); err != nil {
		return nil, err
	}
	return response.Data.Transactions, nil
}
