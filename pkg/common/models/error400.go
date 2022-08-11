package models

type Error400Get struct {
	Message 	string		`json:"message" example:"Data not found with the passed parameters"`
}

type Error400Create struct {
	Message 	string		`json:"message" example:"Could not create. Parameters were not passed correctly"`
}

type Error400Delete struct {
	Message		string		`json:"message" example:"Unable to delete. non-existent ID "`
}

type Error400Update struct {
	Message		string		`json:"message" example:"could not be modified. The parameters do not meet the requirements "`
}