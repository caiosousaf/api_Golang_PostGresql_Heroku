package models

type Error404Message struct {
	Message 	string	`json:"message" example:"Cannot BindJSON"`
}

type Error404Create struct {
	Message		string	`json:"message" example:"Loss of contact with the database"`
}

type Error404Delete struct {
	Message		string	`json:"message" example:"Loss of contact with the database"`
}

type Error404Get struct {
	Message		string	`json:"message" example:"Loss of contact with the database"`
}

type Error404Update struct {
	Message		string	`json:"message" example:"Loss of contact with the database"`
}