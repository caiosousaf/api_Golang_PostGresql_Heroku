package models

type Error400Get struct {
	Message 	string		`json:"message" example:"Data not found with the passed parameters"`
}

type Error400Create struct {
	Message 	string		`json:"message" example:"Could not create. Parameters were not passed correctly"`
}