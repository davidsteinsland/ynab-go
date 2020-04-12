package ynab

type PayeeLocationsService service

type PayeeLocationsResponse struct {
	Data PayeeLocationsWrapper `json:"data"`
}

type PayeeLocationsWrapper struct {
	PayeeLocations []PayeeLocation `json:"payee_locations"`
}

type PayeeLocationResponse struct {
	Data PayeeLocationWrapper `json:"data"`
}

type PayeeLocationWrapper struct {
	PayeeLocation PayeeLocation `json:"payee_location"`
}

type PayeeLocation struct {
	Id        string  `json:"id"`
	PayeeId   string  `json:"payee_id"`
	Latitude  *string `json:"latitude"`
	Longitude *string `json:"longitude"`
}

/*
https://api.youneedabudget.com/v1#/Payee_Locations/getPayeeLocations
*/
func (pls *PayeeLocationsService) List(budgetId string) ([]PayeeLocation, error) {
	var response PayeeLocationsResponse
	if err := service(*pls).do("GET", "budgets/"+budgetId+"/payee_locations", nil, &response); err != nil {
		return nil, err
	}
	return response.Data.PayeeLocations, nil
}

/*
https://api.youneedabudget.com/v1#/Payee_Locations/getPayeeLocationById
*/
func (pls *PayeeLocationsService) Get(budgetId string, payeeLocationId string) (PayeeLocation, error) {
	var response PayeeLocationResponse
	if err := service(*pls).do("GET", "budgets/"+budgetId+"/payee_locations/"+payeeLocationId, nil, &response); err != nil {
		return PayeeLocation{}, err
	}
	return response.Data.PayeeLocation, nil
}

/*
https://api.youneedabudget.com/v1#/Payee_Locations/getPayeeLocationsByPayee
*/
func (pls *PayeeLocationsService) GetByPayee(budgetId string, payeeId string) ([]PayeeLocation, error) {
	var response PayeeLocationsResponse
	if err := service(*pls).do("GET", "budgets/"+budgetId+"/payees/"+payeeId+"/locations", nil, &response); err != nil {
		return nil, err
	}
	return response.Data.PayeeLocations, nil
}
