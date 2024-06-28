package controllers

type AddUserRequest struct {
	UserInfo
}

type UserInfo struct {
	ID       string `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
}
