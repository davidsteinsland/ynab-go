package ynab

type UserService service

type User struct {
	Id string `json:"id"`
}

type UserResponse struct {
	Data UserWrapper `json:"data"`
}

type UserWrapper struct {
	User User `json:"user"`
}

/*
https://api.youneedabudget.com/v1#/User/getUser
*/
func (us *UserService) Get() (User, error) {
	var response UserResponse
	if err := service(*us).do("GET", "user", nil, &response); err != nil {
		return User{}, err
	}
	return response.Data.User, nil
}
