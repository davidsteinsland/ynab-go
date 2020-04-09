package ynab

import (
	"strings"
	"time"
)

type CategoriesService service

type CategoriesResponse struct {
	Data CategoryGroupsWrapper `json:"data"`
}

type CategoryGroupsWrapper struct {
	CategoryGroups []CategoryGroupWithCategories `json:"category_groups"`
}

type CategoryResponse struct {
	Data CategoryWrapper `json:"data"`
}

type CategoryWrapper struct {
	Category Category `json:"category"`
}

type CategoryGroup struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Hidden bool   `json:"hidden"`
}

type CategoryGroupWithCategories struct {
	CategoryGroup
	Categories []Category `json:"categories"`
}

type Category struct {
	Id              string  `json:"id"`
	CategoryGroupId string  `json:"category_group_id"`
	Name            string  `json:"name"`
	Hidden          bool    `json:"hidden"`
	Note            *string `json:"note"`
	Budgeted        int     `json:"budgeted"`
	Activity        int     `json:"activity"`
	Balance         int     `json:"balance"`
}

/*
https://api.youneedabudget.com/v1#/Categories/getCategories
*/
func (cs *CategoriesService) List(budgetId string) ([]CategoryGroupWithCategories, error) {
	var response CategoriesResponse
	if err := service(*cs).do("GET", "budgets/"+budgetId+"/categories", nil, &response); err != nil {
		return nil, err
	}
	return response.Data.CategoryGroups, nil
}

/*
https://api.youneedabudget.com/v1#/Categories/getCategoryById
*/
func (cs *CategoriesService) Get(budgetId string, categoryId string) (Category, error) {
	var response CategoryResponse
	if err := service(*cs).do("GET", "budgets/"+budgetId+"/categories/"+categoryId, nil, &response); err != nil {
		return Category{}, err
	}
	return response.Data.Category, nil
}

func (cs *CategoriesService) Patch(budgetId string, month time.Time, categoryId string, category Category) (Category, error) {
	monthStr := strings.Split(month.Format(time.RFC3339), "T")[0]

	var response CategoryResponse
	err := service(*cs).do("PATCH", "budgets/"+budgetId+"/months/"+monthStr+"/categories/"+categoryId, category, &response)
	if err != nil {
		return Category{}, err
	}
	return response.Data.Category, nil
}
