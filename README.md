Go YNAB Client
==============

Unofficial Golang client library for the YNAB API.

## Example Usage

```go

accessToken := "foobar"
client := ynab.NewDefaultClient(accessToken)

budgets, _ := client.BudgetService.List()
for _, budgetSummary := range budgets {
    fmt.Printf("Budget %v: %v\n", budgetSummary.Id, budgetSummary.Name)

    budget, _ := client.BudgetService.Get(budgetSummary.Id)

    fmt.Printf("Accounts:\n")
    for _, account := range budget.Accounts {
        fmt.Printf("\tAccount %v: %v\n", account.Id, account.Name)
        fmt.Printf("\t\tBalance: %v\n", account.Balance)

        transactions, _ := client.TransactionsService.GetByAccount(budgetSummary.Id, account.Id)

        fmt.Printf("\t\tTransactions:\n")
        for _, transaction := range transactions {
            fmt.Printf("\t\t\t%v: %v\n", transaction.Date, transaction.Amount)
        }
    }
}
```
