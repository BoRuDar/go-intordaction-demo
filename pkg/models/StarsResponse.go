package models

type StarsResponse struct {
	ID    int `json:"id"`
	Stars int `json:"stargazers_count"`
	Forks int `json:"forks_count"`
}
